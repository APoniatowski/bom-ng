package main

import (
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/APoniatowski/bom-ng/internal/news"
	"github.com/APoniatowski/bom-ng/internal/twitter"
)

func main() {
	nitterUsernames := []string{"disclosetv", "warmonitors", "thewarmonitor"}
	newsSites := []string{
		"https://www.bbc.co.uk/news",
		"https://www.disclose.tv/news",
		"https://www.rt.com/",
		"https://www.reuters.com/",
		"https://www.forbes.com/",
		"https://www.aljazeera.com/",
		"https://www.theguardian.com/international",
		"https://news.yahoo.com",
	}
	feedParser := gofeed.NewParser()
	textChecker := news.TextChecker{}
	for {
		for _, newsSite := range newsSites {
			news.CheckNews(textChecker, newsSite)
			time.Sleep(60 * time.Second)
		}
		twitter.CheckNitterRSS(feedParser, textChecker, nitterUsernames)
		time.Sleep(2700 * time.Second)
	}
}
