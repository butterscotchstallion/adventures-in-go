package main

import (
	"github.com/go-co-op/gocron/v2"
	"go.uber.org/zap"
	"time"
)

// ScheduleSpaceCheck /**
func ScheduleSpaceCheck(scheduler gocron.Scheduler) {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()

	//scheduler, _ := gocron.NewScheduler()
	defer func() { _ = scheduler.Shutdown() }()

	job, err := scheduler.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(
				gocron.NewAtTime(1, 6, 0),
			),
		),
		gocron.NewTask(func() {
			CheckLowSpaceAndNotify()
		}),
	)

	if err != nil {
		sugar.Errorf("Error running job: %v", err)
	} else {
		sugar.Debugf("Job %v created", job.Name())
		scheduler.Start()
		sugar.Debug("Scheduler started")

		select {
		case <-time.After(time.Minute):
		}

		err = scheduler.Shutdown()
		if err != nil {
			sugar.Errorf("Error shutting down scheduler: %v", err)
		}
	}
}
