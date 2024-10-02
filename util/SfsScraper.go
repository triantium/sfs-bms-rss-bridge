package util

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"regexp"
)

type Course struct {
	coursenumber string
	course       string
	coursetype   string
	start        string
	end          string
	free         string
	link         string
}

func Scrape() []Course {
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

	linkpattern := "https://www.sfsg.de/lehrgaenge/lehrgangsangebot/detailansicht/%s/"

	courses := make([]Course, 0)
	// Find Courses
	doc.Find("table.klein>tbody>tr>td[width=\"40%\"]").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		coursenumber, exists := s.Find("input[name=\"LgNr\"]").Attr("value")

		if exists {
			//fmt.Printf("Review %d: %s\n", i, coursenumber)
			coursename, _ := s.Find("input[name=\"Lehrgang\"]").Attr("value")
			pattern := regexp.MustCompile(`^.+-.+-(.+)-.+-.+$`)
			ctype := pattern.ReplaceAllString(coursenumber, "$1")
			start, _ := s.Find("input[name=\"Beginn\"]").Attr("value")
			end, _ := s.Find("input[name=\"Ende\"]").Attr("value")
			free, _ := s.Find("input[name=\"Gesamt\"]").Attr("value")
			link := fmt.Sprintf(linkpattern, ctype)
			c := Course{coursenumber, coursename, ctype, start, end, free, link}
			courses = append(courses, c)

			// fmt.Printf("Course %d: %s\n", i, c)
		}

	})
	return courses
}
