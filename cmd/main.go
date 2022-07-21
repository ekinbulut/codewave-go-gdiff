package main

import (
	"flag"

	cli "github.com/ekinbulut/gdiff/cmd/cli"
)

// flag -c=https://lego.storeturkey.com.tr/technic?ps=4
// flag -c=https://lego.storeturkey.com.tr/technic?ps=4 -o=.\lego.html

var (
	site       string
	outputFile string
	interval   int
	keyword    string
	username   string
)

// parse flags
func parseFlags() {
	flag.StringVar(&site, "u", "", "u=https://sample.com")
	flag.StringVar(&site, "url", "", "url=https://sample.com")
	flag.StringVar(&outputFile, "o", "output.html", "o=output.html")
	flag.StringVar(&keyword, "w", "", "w=keyword")
	flag.IntVar(&interval, "i", 0, "i=1")
	flag.StringVar(&username, "user", "", "user=username")

	flag.Parse()
}

func main() {
	parseFlags()
	app := cli.NewApp(site, outputFile)
	app.SetInterval(interval)
	app.SetKeyword(keyword)
	app.SetUsername(username)
	app.Run()
}
