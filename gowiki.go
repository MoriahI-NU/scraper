package main

import (
	"encoding/json"
	"fmt"
	gowiki "github.com/trietmn/go-wiki"
	"os"
)

// Article structure
type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	//Create output file - JSON LINES
	file, err := os.Create("output2.jsonl")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//For this package, only the title is needed and not the whole url
	titles := []string{
		"Robotics",
		"Robot",
		"Reinforcement_learning",
		"Robot_Operating_System",
		"Intelligent_agent",
		"Software_agent",
		"Robotic_process_automation",
		"Chatbot",
		"Applications_of_artificial_intelligence",
		"Android_(robot)",
	}

	//for loop to iterate through each article
	for _, scraperUrl := range titles {
		page, err := gowiki.GetPage(scraperUrl, -1, false, true)
		if err != nil {
			fmt.Println(err)
			continue
		}

		//Get text information from page
		content, err := page.GetContent()
		if err != nil {
			fmt.Println(err)
			continue
		}

		//Fit into Article structure
		data := Article{
			Title:   page.Title,
			Content: content,
		}

		//Encode into json - each iteration will be on a new line to match
		//json lines format
		encoder := json.NewEncoder(file)
		if err := encoder.Encode(data); err != nil {
			fmt.Println(err)
		}
	}
}
