package data

import (
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"lega-bridge/util"
	"os"
)

func SetUpCron() (gocron.Scheduler, error) {

	s, e := gocron.NewScheduler()
	//defer func() { _ = s.Shutdown() }()

	if e != nil {
		return nil, e
	}

	crontab := getEnv("FETCH_CRON_TAB", "*/10 * * * *")

	fmt.Println("Setting up Cron with ", crontab, "")

	_, _ = s.NewJob(
		gocron.CronJob(
			// standard cron tab parsing
			crontab,
			false,
		),
		gocron.NewTask(
			func() {
				fmt.Println("Updating Courses")
				UpdateCourses(util.Scrape())
			},
		),
	)

	return s, nil

}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
