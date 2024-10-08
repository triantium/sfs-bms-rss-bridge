package util

import (
	"fmt"
	"github.com/gorilla/feeds"
	"log"
	"time"
)

func GenerateAtom(courses []Course) string {
	feed := generateFeed(courses)
	atom, err := feed.ToAtom()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(atom)
	return atom
}

func GenerateRSS(courses []Course) string {
	feed := generateFeed(courses)
	rss, err := feed.ToRss()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(rss)
	return rss
}

func GenerateJSON(courses []Course) string {
	feed := generateFeed(courses)
	json, err := feed.ToJSON()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(json)
	return json
}

func generateFeed(courses []Course) *feeds.Feed {
	now := time.Now()
	feed := &feeds.Feed{
		Title: "Lehrgänge der SFS Bayern",
		Link:  &feeds.Link{Href: "https://lega.sfs-bayern.de/"},
		Id:    "https://lega.sfs-bayern.de/", // -> https://datatracker.ietf.org/doc/html/rfc3987

		Created: now,
	}

	for _, c := range courses {
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       c.course,
			Link:        &feeds.Link{Href: c.link},
			Id:          c.coursenumber, // -> https://datatracker.ietf.org/doc/html/rfc3987
			Content:     fmt.Sprintf("Lehrgang: %s<br/>Beginn: %s<br/>Ende: %s<br/>Freie Plätze: %s", c.course, c.start, c.end, c.free),
			Description: c.coursetype,
		})
	}
	return feed
}
