package scrapper

import (
	URL "net/url"

	"github.com/gocolly/colly"
)

func GetImageUrls(url string) []string {
	if !tryParseURL(url) {
		panic("The request url is not valid")
	}

	images := make([]string, 0)
	c := colly.NewCollector()

	signalChan := make(chan bool, 1)
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		link := e.Attr("src")

		images = append(images, link)

	})

	c.OnScraped(func(e *colly.Response) {

		signalChan <- true
	})

	c.OnError(func(e *colly.Response, er error) {
		signalChan <- true
	})
	c.Visit(url)
	if <-signalChan {
		return images
	} else {
		panic("The Request Failed")
	}

}

func tryParseURL(url string) bool {
	_, err := URL.ParseRequestURI(url)
	if err != nil {
		return false
	}
	return true
}
