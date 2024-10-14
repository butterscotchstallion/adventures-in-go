package main

import (
	gocronmocks "github.com/go-co-op/gocron/mocks/v2"
	"github.com/go-co-op/gocron/v2"
	"go.uber.org/mock/gomock"
	"testing"
)

func myFunc(s gocron.Scheduler) {
	s.Start()
	_ = s.Shutdown()
}

func TestMyFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	s := gocronmocks.NewMockScheduler(ctrl)
	s.EXPECT().Start().Times(1)
	s.EXPECT().Shutdown().Times(1).Return(nil)

	myFunc(s)
}

func TestScheduleSpaceCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	s := gocronmocks.NewMockScheduler(ctrl)
	s.EXPECT().Start().Times(1)
	s.EXPECT().Shutdown().Times(1).Return(nil)

	//ScheduleSpaceCheck(s)
}
