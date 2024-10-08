package main

import (
	"lega-bridge/data"
	"lega-bridge/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	s, e := data.SetUpCron()

	if e != nil {
		panic(e)
	}
	defer func() { _ = s.Shutdown() }()
	s.Start()

	r := gin.Default()
	r.GET("/feed/atom", func(c *gin.Context) {
		atom := util.GenerateAtom(data.GetCourses())
		c.Header("Content-Type", "application/atom+xml; charset=utf-8")
		c.String(http.StatusOK, atom)
	})
	r.GET("/feed/rss", func(c *gin.Context) {
		rss := util.GenerateRSS(data.GetCourses())
		c.Header("Content-Type", "application/rss+xml; charset=utf-8")
		c.String(http.StatusOK, rss)
	})
	r.GET("/feed/json", func(c *gin.Context) {
		json := util.GenerateJSON(data.GetCourses())
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.String(http.StatusOK, json)
	})
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
