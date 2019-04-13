package simulationService

import "time"

type SimulationService interface {
	GetElapsedTime() time.Duration
}

type SimulationServiceControl interface {
	SimulationService
	SetElapsedTime(elapsed time.Duration)
}

type simulationServiceImpl struct {
	elapsed time.Duration
}

func New() SimulationServiceControl {
	return &simulationServiceImpl{}
}

func (s *simulationServiceImpl) GetElapsedTime() time.Duration {
	return s.elapsed
}

func (s *simulationServiceImpl) SetElapsedTime(elapsed time.Duration) {
	s.elapsed = elapsed
}
