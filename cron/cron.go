package cron

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

// second minute hour day_of_month month day_of_week
var (
	specEvery8Am         string = "0 0 8 * * *"
	specEvery12pm        string = "0 0 12 * * *"
	specEvery10m         string = "0 0/10 * * *"
	specHourly           string = "@hourly"
	specWeekly           string = "@weekly"
	specDaily            string = "@daily"
	specEvery2h          string = "0 0 0/2 * *"
	specEvery4h          string = "0 0 0/4 * *"
	specEvery12h         string = "0 0 0/12 * *"
	specEveryTuesday12pm string = "0 0 12 * * 2"
	specEveryFriday5pm   string = "0 0 17 * * 5"
	specFirstDayInAMonth string = "0 0 0 1 * *"
	specEveryTheusday    string = "0 0 0 * * 2"
)
var spec string = "0 1 * * * *"

func Cron() {
	c := cron.New()

	c.AddFunc("@every 1s", func() {
		fmt.Println(time.Now())
	})
	c.Start()
	time.Sleep(time.Second * 800)
}
