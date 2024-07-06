package scheduler

import (
    "github.com/robfig/cron/v3"
)

func Schedule(schedule string, task func()) {
    c := cron.New()
    c.AddFunc(schedule, task)
    c.Start()
}
