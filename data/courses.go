package data

import "lega-bridge/util"

var courses = util.Scrape()

func GetCourses() []util.Course {
	return courses
}

func UpdateCourses(newCourses []util.Course) {
	courses = newCourses
}
