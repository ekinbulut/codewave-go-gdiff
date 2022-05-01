package srv

import (
	"io/ioutil"
	"net/http"
)

type Crawler struct {
	url string
}

func NewCrawler(url string) *Crawler {
	return &Crawler{
		url: url,
	}
}

func (c *Crawler) Crawl() (string, error) {

	var response string
	resp, err := http.Get(c.url)
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
