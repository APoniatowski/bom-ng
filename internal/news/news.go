package news

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type TextChecker struct{}

// List of countries, verbs, and extra words to look for.
var (
	countries = []string{"Belarus", "Russia", "Union State", "Poland", "Lithuania", "Latvia", "Estonia", "Israel", "Iran"}
	verbs     = []string{"attacks", "invades", "assaults", "advancing towards", "advancing to"}
	extra     = []string{"Wagner", "Breaking News", "PMC", "mercenaries"}
)

// CheckNews function to fetch and parse news from a given URL
func CheckNews(tc TextChecker, url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	tags := []string{"h1", "h2", "h3", "p", "a", "span"}
	for _, tag := range tags {
		doc.Find(tag).Each(func(index int, item *goquery.Selection) {
			text := item.Text()
			if len(strings.Fields(text)) > 4 {
				tc.CheckText(text)
			}
		})
	}
}

// CheckText function to analyze the text for specified keywords
func (tc TextChecker) CheckText(text string) string {
	lowerText := strings.ToLower(text) // convert text to lowercase once for efficiency
	voiceMessage := "news message:"
	for _, verb := range verbs {
		verbIndex := strings.Index(lowerText, strings.ToLower(verb))
		if verbIndex != -1 {
			for _, country1 := range countries {
				country1Index := strings.Index(lowerText, strings.ToLower(country1))
				if country1Index != -1 {
					for _, country2 := range countries {
						if country1 != country2 {
							country2Index := strings.Index(lowerText, strings.ToLower(country2))
							if country2Index != -1 {
								// Determine the order of countries based on their position in the text
								if country1Index < country2Index {
									voiceMessage += fmt.Sprintf(" %s %s %s", strings.ToLower(country1), strings.ToLower(verb), strings.ToLower(country2))
								} else {
									voiceMessage += fmt.Sprintf(" %s %s %s", strings.ToLower(country2), strings.ToLower(verb), strings.ToLower(country1))
								}
								for _, word := range extra {
									if strings.Contains(lowerText, strings.ToLower(word)) {
										output, _ := exec.Command("sh", "-c", "pactl list sinks | grep Mute").Output()
										if strings.Contains(string(output), "yes") {
											err := exec.Command("pactl", "set-sink-mute", "0", "0").Run()
											if err != nil {
												fmt.Println(err)
											}
											err = exec.Command("pactl", "set-sink-volume", "0", "10%").Run()
											if err != nil {
												fmt.Println(err)
											}
										}
										voiceMessage += " " + word
										// Display a notification and play a sound
										err := exec.Command("notify-send", "News Alert", voiceMessage).Run()
										if err != nil {
											fmt.Println(err)
										}
										Speak(voiceMessage)
										if strings.Contains(string(output), "yes") {
											err := exec.Command("pactl", "set-sink-mute", "0", "1").Run()
											if err != nil {
												fmt.Println(err)
											}
										}
										fmt.Println(voiceMessage)
										return voiceMessage
									} else {
										err := exec.Command("notify-send", "News Warning", voiceMessage).Run()
										if err != nil {
											fmt.Println(err)
										}
										fmt.Println(voiceMessage)
										return voiceMessage
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return voiceMessage
}

func Speak(message string) {
	cmd := exec.Command("espeak", message)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
