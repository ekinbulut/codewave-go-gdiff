package main

import (
	"bufio"

	"fmt"

	"io/ioutil"
	"net/http"
	"os"
)

type Crawler struct {
	
}

func (c *Crawler) Crawl(url string) (string, error) {

	var response string
	resp, err := http.Get(url)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	response = string(body)
	return response, nil
}

type FileWriter struct {
	fileName string
	file     *os.File
}

// write string to file
func (f *FileWriter) Write(s string) error {

	f.file, _ = os.Create(f.fileName)
	defer f.file.Close()
	writer := bufio.NewWriter(f.file)
	_, err := writer.WriteString(s)
	if err != nil {
		return err
	}
	writer.Flush()

	return nil
}

func main() {

	c := &Crawler{}
	resp, err := c.Crawl("https://site.com")
	if err != nil {
		panic(err)
	}

	// get href attribute

	fileWriter := &FileWriter{
		fileName: ".\\site.html",
		file:     &os.File{},
	}

	err = fileWriter.Write(resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
