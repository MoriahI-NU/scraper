package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
)

type Sections struct {
	Header  string
	Content string
}

type Article struct {
	Title    string
	Sections []Sections
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

	for _, scraperUrl := range urls {

		article := Article{}
		sections := []Sections{}

		c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"))

		c.OnHTML("span.mw-page-title-main", func(e *colly.HTMLElement) {
			article.Title = e.Text
		})

		currentHeader := ""

		c.OnHTML("p, h2", func(h *colly.HTMLElement) {
			if h.Name == "h2" {
				currentHeader = h.Text
			} else if h.Name == "p" {
				sections = append(sections, Sections{Header: currentHeader, Content: h.Text})
			}
		})

		c.OnScraped(func(r *colly.Response) {
			article.Sections = groupSections(sections)
			writeToJSONLinesFile("output.jsonl", article)
		})

		c.Visit(scraperUrl)
	}
}

// /////////HELPER TO AGGREGATE PARAGRAPHS
func groupSections(sections []Sections) []Sections {
	var groupedSections []Sections
	if len(sections) == 0 {
		return groupedSections
	}

	currentSection := sections[0]
	for i := 1; i < len(sections); i++ {
		if sections[i].Header == currentSection.Header {
			currentSection.Content += "\n" + sections[i].Content
		} else {
			groupedSections = append(groupedSections, currentSection)
			currentSection = sections[i]
		}
	}
	groupedSections = append(groupedSections, currentSection)
	return groupedSections
}

// ///////////HELPER TO WRITE TO JSON
func writeToJSONLinesFile(filename string, data Article) {
	//file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		fmt.Println("Error:", err)
	}
}

//////////////////////////////////////////////////////
// encoder := json.NewEncoder(file)
// for _, article := range data.Article {
// 	if err := encoder.Encode(article); err != nil {
// 		fmt.Println("Error:", err)
//////////////////////////////////////////////////////

//instead of :
// c.OnHTML("p, h2", func(h *colly.HTMLElement) {
// 	if h.Name == "h2" {
// 		currentHeader = h.Text
// 	} else if h.Name == "p" {
// 		sections = append(sections, Sections{Header: currentHeader, Content: h.Text})
// 	}
// })

//Try this?
// c.OnHTML("p, h2, ul, ol", func(h *colly.HTMLElement) {
// 	if h.Name == "h2" {
// 		currentHeader = h.Text
// 	} else if h.Name == "p" {
// 		sections = append(sections, Sections{Header: currentHeader, Content: h.Text})
// 	}
// })
