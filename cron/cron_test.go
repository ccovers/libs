package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	err := Cron([]Job{
		{
			Spec: SpecEverySec,
			JobFunc: func() {
				fmt.Println(time.Now())
			},
		},
		{
			Spec: SpecEvery1m,
			JobFunc: func() {
				fmt.Println("=====", time.Now())
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(time.Second * 8)
}
