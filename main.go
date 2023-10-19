package main

import (
	"fmt"
	"github.com/trietmn/go-wiki"
)

func main() {
	page, err := gowiki.GetPage("Robotics", -1, false, true)
	if err != nil {
		fmt.Println(err)
	}

	content, err := page.GetContent()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", content)
}
