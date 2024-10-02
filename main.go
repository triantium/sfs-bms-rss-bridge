package main

import (
	"fmt"
	"lega-bridge/util"
)

func main() {
	courses := util.Scrape()
	for i, c := range courses {
		fmt.Printf("Courses %d: %s\n", i, c)
	}
}
