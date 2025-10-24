package util

import (
	"fmt"
	"lega-bridge/data"
	"log"
	"time"

	"github.com/gorilla/feeds"
)

func GenerateAtom(courses []data.Course) string {
	feed := generateFeed(courses)
	atom, err := feed.ToAtom()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(atom)
	return atom
}

func GenerateRSS(courses []data.Course) string {
	feed := generateFeed(courses)
	rss, err := feed.ToRss()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(rss)
	return rss
}

func GenerateJSON(courses []data.Course) string {
	feed := generateFeed(courses)
	json, err := feed.ToJSON()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(json)
	return json
}

func generateFeed(courses []data.Course) *feeds.Feed {
	now := time.Now()
	feed := &feeds.Feed{
		Title:   "Lehrgänge der SFS Bayern",
		Link:    &feeds.Link{Href: "https://www.bms-fw.bayern.de/Navigation/Public/LastMinute.aspx"},
		Id:      "https://www.bms-fw.bayern.de/", // -> https://datatracker.ietf.org/doc/html/rfc3987
		Created: now,
	}

	for _, c := range courses {
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       c.CourseName,
			Link:        &feeds.Link{Href: c.Link},
			Id:          c.CourseNumber, // -> https://datatracker.ietf.org/doc/html/rfc3987
			Content:     fmt.Sprintf("Lehrgang: %s<br/>Beginn: %s<br/>Ende: %s<br/>Schulungsort: %s</br>Freie Plätze: %s", c.CourseName, c.Start, c.End, c.Place, c.Free),
			Description: c.CourseType,
		})
	}
	return feed
}
