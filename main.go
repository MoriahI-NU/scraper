package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"strings"
)

// Create Article Structure
type Article struct {
	Url   string
	Title string
	Text  string
	Tags  string
}

func main() {

	// Websites of Interest
	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	//iterate through each website
	for _, scraperUrl := range urls {

		//Empty container for each article
		article := Article{}

		c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"))

		//define url
		article.Url = scraperUrl

		//define title
		c.OnHTML("span.mw-page-title-main", func(e *colly.HTMLElement) {
			article.Title = e.Text
		})

		//define text/content
		c.OnHTML("div.mw-parser-output", func(h *colly.HTMLElement) {
			article.Text = h.Text
		})

		//define tags - and separate them with a comma and a space for readability
		c.OnHTML("div.mw-normal-catlinks", func(t *colly.HTMLElement) {
			tags := strings.Split(t.Text, "\n")
			article.Tags = strings.Join(tags, ", ")
		})

		//create output file
		c.OnScraped(func(r *colly.Response) {
			writeToJSONLinesFile("output3.jsonl", article)
		})

		c.Visit(scraperUrl)
	}
}

// ///////////HELPER TO WRITE TO JSON LINES
func writeToJSONLinesFile(filename string, data Article) {
	//file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	//Write each article to a separate line to fit Json Lines format
	if err := encoder.Encode(data); err != nil {
		fmt.Println("Error:", err)
	}
}
