package main

import (
	"fmt"
	"strings"
	"time"

	"flag"

	service "github.com/ekinbulut/go-http-crawler/app/srv"
)

// flag -c=https://lego.storeturkey.com.tr/technic?ps=4
// flag -c=https://lego.storeturkey.com.tr/technic?ps=4 -o=.\lego.html

var site string
var outputFile string
var interval int
var keyword string

type App struct {
	Name        string
	Version     string
	crawler     *service.Crawler
	fileWriter  *service.FileWriter
	diffChecker *service.DiffChecker
}

func NewApp() *App {
	return &App{
		Name:        "go-http-crawler",
		Version:     "1.0.0",
		crawler:     service.NewCrawler(site),
		fileWriter:  service.NewFileWriter(outputFile),
		diffChecker: service.NewDiffChecker(),
	}
}

func (a *App) Run() {

	a.printAppInfo()

	// print flags
	fmt.Println("site:", site)
	fmt.Println("outputFile:", outputFile)

	// execute in given interval
	if interval > 0 {
		for {
			a.execute()
			fmt.Println("sleeping...")
			time.Sleep(time.Duration(interval) * time.Second)
		}
	} else {
		a.execute()
	}

}

func (app *App) execute() {
	// print progress
	fmt.Println("crawling...")

	resp, err := app.crawlsite(site)
	if err != nil {
		fmt.Println(err)
	}

	b := app.fileWriter.Exists()
	if b {
		// read file
		old, err := app.fileWriter.Read()
		if err != nil {
			fmt.Println(err)
		}

		// search for keywords in string
		// e := checkKeyword(old, keyword)
		// if e {
		// 	fmt.Println("keyword found")
		// 	return
		// } else {
		// 	fmt.Println("keyword not found")
		// 	n := service.NewNotification("Value has changed")
		// 	n.SendEmail()
		// }

		// check diff
		b, err := app.diffChecker.Check(old, resp)
		if err != nil {
			fmt.Println(err)
		}
		if b {
			fmt.Println("no changes")
		} else {
			fmt.Println("changes found")
			// print diffs
			html, err := app.diffChecker.PrintDiffsToHtml(old, resp)
			if err != nil {
				fmt.Println(err)
			}

			//fmt.Println(html)
			err = app.createOutput(resp)
			if err != nil {
				fmt.Println(err)
			}

			n := service.NewNotification()
			n.SendEmail("", "", "", html)
		}
	} else {
		err := app.createOutput(resp)
		if err != nil {
			fmt.Println(err)
		}
	}

	// print progress
	fmt.Println("done")

}

// check keyword
func checkKeyword(old string, new string) bool {
	return strings.Contains(old, keyword)
}

// print App info
func (a *App) printAppInfo() {
	fmt.Printf("%s %s\n", a.Name, a.Version)
}

func (app *App) crawlsite(site string) (string, error) {

	resp, err := app.crawler.Crawl()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return resp, nil
}

func (app *App) createOutput(resp string) error {
	f := app.fileWriter
	err := f.Write(resp)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = f.Rename(outputFile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	parseFlags()
	app := NewApp()
	app.Run()

}

// parse flags
func parseFlags() {
	flag.StringVar(&site, "u", "https://lego.storeturkey.com.tr/10300-lego-icons-gelecege-donus-zaman-makinesi", "u=https://sample.com")
	flag.StringVar(&site, "url", "https://lego.storeturkey.com.tr/10300-lego-icons-gelecege-donus-zaman-makinesi", "url=https://sample.com")
	flag.StringVar(&outputFile, "o", "output.html", "o=output.html")
	flag.StringVar(&keyword, "w", "", "w=keyword")
	flag.IntVar(&interval, "i", 0, "i=1")

	flag.Parse()
}
