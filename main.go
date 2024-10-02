package main

import (
	"fmt"
	"lega-bridge/util"
)

func main() {
	courses := util.Scrape()
	util.GenerateRSS(courses)
	util.GenerateAtom(courses)
	util.GenerateJSON(courses)
	fmt.Println("Done!")
}
