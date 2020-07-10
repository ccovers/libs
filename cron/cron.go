package cron

import (
	"github.com/robfig/cron"
)

// minute hour day_of_month month day_of_week
var (
	SpecEverySec         string = "@every 1s"
	SpecHourly           string = "@hourly"
	SpecDaily            string = "@daily"
	SpecWeekly           string = "@weekly"
	SpecEvery1thm        string = "1 * * * *"
	SpecEvery1m          string = "0/1 * * * *"
	SpecEvery8Am         string = "0 8 * * *"
	SpecEvery8h          string = "0 0/8 * * *"
	SpecEveryTuesday18pm string = "0 18 * * 2"
	SpecFirstDayInAMonth string = "0 0 1 * *"
	SpecEveryTheusday    string = "0 0 * * 2"
)

type JobFunc func()
type Job struct {
	Spec    string
	JobFunc JobFunc
}

func Cron(jobs []Job) error {
	if len(jobs) == 0 {
		return nil
	}

	c := cron.New()
	for _, job := range jobs {
		_, err := c.AddFunc(job.Spec, job.JobFunc)
		if err != nil {
			return err
		}
	}

	c.Start()
	return nil
}
