package data

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"log"
	"net/http"
	"regexp"
)

func Scrape() []Course {
	res, err := http.Get("https://lega.sfs-bayern.de/cgi-perl/lega-display.pl")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	utfBody, err := iconv.NewReader(res.Body, "iso-8859-1", "utf-8")
	if err != nil {
		log.Fatal(err)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		log.Fatal(err)
	}

	// Link zur Kursübersicht, nicht zur Anmeldung
	linkpattern := "https://www.sfsg.de/lehrgaenge/lehrgangsangebot/detailansicht/%s/"

	courses := make([]Course, 0)
	// Finde Lehrgänge
	doc.Find("table.klein>tbody>tr>td[width=\"40%\"]").Each(func(i int, s *goquery.Selection) {

		coursenumber, exists := s.Find("input[name=\"LgNr\"]").Attr("value")

		if exists {
			// Gehen wir mal davon aus das alles mit Lehrgangsnummer auch eine valide Zeile ist.
			// Mögliche UTF-8 Zeichen ^^
			coursename, _ := s.Find("input[name=\"Lehrgang\"]").Attr("value")
			pattern := regexp.MustCompile(`^.+-.+-(.+)-.+-.+$`)
			ctype := pattern.ReplaceAllString(coursenumber, "$1")
			start, _ := s.Find("input[name=\"Beginn\"]").Attr("value")
			end, _ := s.Find("input[name=\"Ende\"]").Attr("value")
			free, _ := s.Find("input[name=\"Gesamt\"]").Attr("value")
			link := fmt.Sprintf(linkpattern, ctype)

			c := Course{coursenumber, coursename, ctype, start, end, free, link}
			courses = append(courses, c)
		}

	})
	return courses
}
