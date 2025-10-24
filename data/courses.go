package data

import (
	"fmt"
	"strings"
)

type Course struct {
	CourseNumber string
	CourseName   string
	CourseType   string
	Start        string
	End          string
	Free         string
	Link         string
	Place        string
}

var courses = Scrape()

func GetCourses() []Course {
	return courses
}

func UpdateCourses(update []Course) {
	var deletedCourses = make([]Course, 0)
	var updatedCourses = make([]Course, 0)
	var newCourses = make([]Course, 0)
	for _, c := range courses {
		// if missing -> deleted
		contains := containsCourse(update, c)

		if contains {
			// if in slice -> updated
			updatedCourses = append(updatedCourses, c)
		} else {
			// if missing -> deleted-candidate
			deletedCourses = append(deletedCourses, c)
		}

	}
	for _, c := range update {
		// if not in courses -> new
		contains := containsCourse(courses, c)
		if !contains {
			newCourses = append(newCourses, c)
			fmt.Println("New Course found: ", c.CourseName, ".")
		} else {
			// if missing -> deleted
			deletedCourses = removeCourse(deletedCourses, c)
		}
	}
	courses = update
}

func removeCourse(courses []Course, c Course) []Course {
	returnCourses := make([]Course, 0)
	for _, course := range courses {
		if !sameCourse(course, c) {
			returnCourses = append(returnCourses, course)
		}
	}
	return returnCourses
}

func containsCourse(courses []Course, course Course) bool {
	for _, c := range courses {
		if sameCourse(c, course) {
			return true
		}
	}

	return false
}

func sameCourse(c1, c2 Course) bool {
	return strings.EqualFold(c1.CourseNumber, c2.CourseNumber)
}
