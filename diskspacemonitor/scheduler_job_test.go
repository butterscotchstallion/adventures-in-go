package main

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/jonboulle/clockwork"
	"sync"
	"testing"
	"time"
)

func TestScheduleSpaceCheck(t *testing.T) {
	fakeClock := clockwork.NewFakeClock()
	s, _ := gocron.NewScheduler(
		gocron.WithClock(fakeClock),
	)
	var wg sync.WaitGroup
	wg.Add(1)

	ScheduleSpaceCheck(s)

	fakeClock.BlockUntil(1)
	fakeClock.Advance(time.Second * 5)
	wg.Wait()
	_ = s.StopJobs()
}
