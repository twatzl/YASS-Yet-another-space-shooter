package timeService

import (
	"time"
)

type TimeService interface {
	GetDeltaTime() time.Duration
	GetElapsedTime() time.Duration
}

type TimeServiceControl interface {
	TimeService
	SetDeltaTime(elapsed time.Duration)
	StartNow()
}

type TimeServiceImpl struct {
	elapsed time.Duration
	start time.Time
}

func New() TimeServiceControl {
	return &TimeServiceImpl{
		elapsed: 0,
	}
}

func (s *TimeServiceImpl) GetDeltaTime() time.Duration {
	return s.elapsed
}

func (s *TimeServiceImpl) GetElapsedTime() time.Duration {
	return time.Now().Sub(s.start)
}

func (s *TimeServiceImpl) SetDeltaTime(elapsed time.Duration) {
	s.elapsed = elapsed
}

func (s *TimeServiceImpl) StartNow() {
	s.start = time.Now()
}