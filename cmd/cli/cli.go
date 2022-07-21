package cli

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	service "github.com/ekinbulut/gdiff/internal"
)

type App struct {
	Name        string
	Version     string
	crawler     *service.Crawler
	fileWriter  *service.FileWriter
	diffChecker *service.DiffChecker
	site        string
	outputFile  string
	interval    int
	keyword     string
	username    string
}

// set flags
func (app *App) SetInterval(i int) {
	app.interval = i
}

func (app *App) SetKeyword(k string) {
	app.keyword = k
}

func (app *App) SetUsername(u string) {
	app.username = u
}

func NewApp(site string, outputFile string) *App {

	app := &App{

		Name:        "go-http-crawler",
		Version:     "1.0.0",
		crawler:     service.NewCrawler(site),
		fileWriter:  service.NewFileWriter(outputFile),
		diffChecker: service.NewDiffChecker(),
		site:        site,
		outputFile:  outputFile,
	}
	return app
}

// check keyword
func (app *App) checkKeyword(old string, new string) bool {
	return strings.Contains(old, app.keyword)
}

func (app *App) Run() {

	app.printAppInfo()

	// print flags
	log.Printf("site: %s", app.site)
	log.Printf("outputFile: %s", app.outputFile)

	// execute in given interval
	if app.interval > 0 {
		go func() {
			for {
				app.execute()
				log.Printf("Waiting for next cycle in %d seconds", app.interval)
				time.Sleep(time.Duration(app.interval) * time.Second)
			}
		}()
		app.gracefulShutdown()
	} else {
		app.execute()
	}

}

func (app *App) gracefulShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	for {
		select {
		case <-c:
			log.Printf("Application exit")
			return
		case <-time.After(time.Duration(app.interval) * time.Second):
			log.Printf("Application exit")
			return
		}
	}
}

func (app *App) execute() {
	// print progress
	log.Printf("Crawling %s", app.site)

	resp, err := app.crawlsite(app.site)
	if err != nil {
		log.Printf("Error Occured: %s", err)
	}

	fileExists := app.fileWriter.Exists()
	if fileExists {
		// read file
		old, err := app.fileWriter.Read()
		if err != nil {
			log.Printf("Error Occured: %s", err)
		}

		if app.keyword != "" {
			if app.checkKeyword(old, resp) {
				log.Printf("Crawling finished %s", app.site)
				if app.username != "" {
					n := service.NewNotification()
					n.SendEmail(app.username, app.username, "Notificaton", "Keyword found")
				}
				return
			}
		}

		// check diff
		b, err := app.diffChecker.Check(old, resp)
		if err != nil {
			log.Printf("Error Occured: %s", err)
		}
		if b {
			log.Printf("No changes found")
		} else {
			log.Printf("Changes found")
			// print diffs
			html, err := app.diffChecker.PrintDiffsToHtml(old, resp)
			if err != nil {
				log.Printf("Error Occured: %s", err)
			}

			err = app.createOutput(resp)
			if err != nil {
				log.Printf("Error Occured: %s", err)
			}

			if app.username != "" {
				n := service.NewNotification()
				n.SendEmail(app.username, app.username, "Notification", html)
			}
		}

	} else {
		err := app.createOutput(resp)
		if err != nil {
			log.Printf("Error Occured: %s", err)
		}

	}

	// print progress
	log.Printf("Crawling finished %s", app.site)

}

// print App info
func (a *App) printAppInfo() {
	log.Printf("%s %s\n", a.Name, a.Version)
}

func (app *App) crawlsite(site string) (string, error) {

	resp, err := app.crawler.Crawl()
	if err != nil {
		log.Printf("Error Occured: %s", err)
		return "", err
	}
	return resp, nil
}

func (app *App) createOutput(resp string) error {
	f := app.fileWriter
	if err := f.Write(resp); err != nil {
		log.Printf("Error Occured: %s", err)
		return err
	}

	if err := f.Rename(app.outputFile); err != nil {
		log.Printf("Error Occured: %s", err)
		return err
	}
	return nil
}
