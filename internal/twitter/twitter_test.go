package twitter

import (
	"testing"

	"github.com/mmcdole/gofeed"
)

type MockFeedParser struct {
	feed *gofeed.Feed
	err  error
}

func (m *MockFeedParser) ParseURL(url string) (*gofeed.Feed, error) {
	return m.feed, m.err
}

type MockTextChecker struct {
	texts []string
}

func (m *MockTextChecker) CheckText(text string) string {
	m.texts = append(m.texts, text)
	return ""
}

func TestCheckNitterRSS(t *testing.T) {
	mockParser := &MockFeedParser{
		feed: &gofeed.Feed{
			Items: []*gofeed.Item{
				{Title: "Test Tweet 1"},
				{Title: "Test Tweet 2"},
			},
		},
	}
	mockChecker := &MockTextChecker{}
	usernames := []string{"testuser"}
	CheckNitterRSS(mockParser, mockChecker, usernames)
	if len(mockChecker.texts) != 2 {
		t.Errorf("CheckNitterRSS() = %d; want 2", len(mockChecker.texts))
	}
}
