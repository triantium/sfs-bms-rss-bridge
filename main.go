package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type course struct {
	coursenumber string
	course       string
	start        string
	end          string
	free         string
}

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://lega.sfs-bayern.de/cgi-perl/lega-display.pl")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find Courses
	doc.Find("table.klein>tbody>tr>td[width=\"40%\"]").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		coursenumber, exists := s.Find("input[name=\"LgNr\"]").Attr("value")
		if exists {
			fmt.Printf("Review %d: %s\n", i, coursenumber)
			coursename, _ := s.Find("input[name=\"Lehrgang\"]").Attr("value")
			start, _ := s.Find("input[name=\"Beginn\"]").Attr("value")
			end, _ := s.Find("input[name=\"Ende\"]").Attr("value")
			free, _ := s.Find("input[name=\"Gesamt\"]").Attr("value")

			c := course{coursenumber, coursename, start, end, free}

			fmt.Printf("Course %d: %s\n", i, c)
		}

	})
}

func main() {
	ExampleScrape()
}
