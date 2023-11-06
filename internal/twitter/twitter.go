package twitter

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

type FeedParser interface {
	ParseURL(url string) (*gofeed.Feed, error)
}

type TextChecker interface {
	CheckText(text string) string
}

// CheckNitterRSS function to fetch and parse tweets from Nitter RSS feeds
func CheckNitterRSS(fp FeedParser, tc TextChecker, nitterUsernames []string) {
	for _, username := range nitterUsernames {
		url := fmt.Sprintf("https://nitter.net/%s/rss", username)
		feed, err := fp.ParseURL(url)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		for _, item := range feed.Items {
			tweetText := item.Title
			tc.CheckText(tweetText)
		}
	}
}
