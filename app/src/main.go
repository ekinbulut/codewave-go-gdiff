package main

import (
	"fmt"

	"flag"

	service "github.com/ekinbulut/go-http-crawler/app/srv"
)

// flag -c=https://lego.storeturkey.com.tr/technic?ps=4
// flag -c=https://lego.storeturkey.com.tr/technic?ps=4 -o=.\lego.html

var site string
var outputFile string

type App struct {
	Name    string
	Version string
}

func NewApp() *App {
	return &App{
		Name:    "go-http-crawler",
		Version: "1.0.0",
	}
}

func (a *App) Run() {

	a.printAppInfo()
	parseFlags()
	// print flags
	fmt.Println("site:", site)
	fmt.Println("outputFile:", outputFile)

	// print progress
	fmt.Println("crawling...")

	resp, err := crawlsite(site)
	if err != nil {
		fmt.Println(err)
	}

	// print progress
	fmt.Println("creating output...")

	err = createOutput(resp)
	if err != nil {
		fmt.Println(err)
	}
	// print progress
	fmt.Println("done")

}

// print help
func printHelp() {
	fmt.Println("go-http-crawler -u=https://sample.com -o=output.html")
}

// print App info
func (a *App) printAppInfo() {
	fmt.Printf("%s %s\n", a.Name, a.Version)
}

func crawlsite(site string) (string, error) {

	c := service.NewCrawler(site)
	resp, err := c.Crawl()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return resp, nil
}

func createOutput(resp string) error {
	f := service.NewFileWriter(outputFile)
	err := f.Write(resp)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	app := NewApp()
	app.Run()

}

// parse flags
func parseFlags() {
	flag.StringVar(&site, "u", "", "u=https://sample.com")
	flag.StringVar(&outputFile, "o", "", "o=output.html")

	flag.Parse()
}
