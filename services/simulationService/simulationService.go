package simulationService

import (
	"github.com/twatzl/pixel-test/game"
	"time"
)

type SimulationService interface {
	GetElapsedTime() time.Duration
	RegisterGameObject(g game.GameObject)
	UnregisterGameObject(g game.GameObject)
}

type SimulationServiceControl interface {
	SimulationService
	SetElapsedTime(elapsed time.Duration)
	UpdateGameObjects()
}

type simulationServiceImpl struct {
	elapsed time.Duration
	gameObjects map[game.GameObject] game.GameObject
}

func New() SimulationServiceControl {
	return &simulationServiceImpl{
		elapsed: 0,
		gameObjects: make(map[game.GameObject] game.GameObject),
	}
}

func (s *simulationServiceImpl) GetElapsedTime() time.Duration {
	return s.elapsed
}

func (s *simulationServiceImpl) SetElapsedTime(elapsed time.Duration) {
	//logrus.Infoln("elapsed time: " + elapsed.String())
	s.elapsed = elapsed
}

func (s *simulationServiceImpl) UpdateGameObjects() {
	for _, g := range s.gameObjects {
		g.Update()
	}
}

func (s *simulationServiceImpl) RegisterGameObject(g game.GameObject) {
	s.gameObjects[g] = g
}

func (s *simulationServiceImpl) UnregisterGameObject(g game.GameObject) {
	delete(s.gameObjects, g)
}
