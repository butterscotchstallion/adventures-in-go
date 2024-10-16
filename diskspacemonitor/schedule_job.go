package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
	"go.uber.org/zap"
)

const jobCheckHour uint = 9
const jobCheckMinute uint = 0

// ScheduleSpaceCheck /**
func ScheduleSpaceCheck(scheduler gocron.Scheduler, logger *zap.SugaredLogger) {
	defer func() { _ = scheduler.Shutdown() }()

	job, err := scheduler.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(
				gocron.NewAtTime(jobCheckHour, jobCheckMinute, 0),
			),
		),
		gocron.NewTask(func() {
			CheckLowSpaceAndNotify()
		}),
	)

	if err != nil {
		logger.Error(fmt.Sprintf("Error running job: %v", err))
	} else {
		logger.Debugf("Job %v created", job.ID())
		scheduler.Start()
		logger.Debug("Scheduler started")

		select {
		case <-time.After(time.Minute):
		}

		err = scheduler.Shutdown()
		if err != nil {
			logger.Errorf("Error shutting down scheduler: %v", err)
		}
	}
}
