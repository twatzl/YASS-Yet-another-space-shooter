package simulationSystem

import (
	"github.com/twatzl/pixel-test/engine/game"
	"time"
)

type SimulationSystem interface {
	GetElapsedTime() time.Duration
	RegisterGameObject(g game.GameObject)
	UnregisterGameObject(g game.GameObject)
}

type SimulationSystemControl interface {
	SimulationSystem
	SetElapsedTime(elapsed time.Duration)
	UpdateGameObjects()
}

type simulationSystemImpl struct {
	elapsed time.Duration
	gameObjects map[game.GameObject] game.GameObject
}

func New() SimulationSystemControl {
	return &simulationSystemImpl{
		elapsed: 0,
		gameObjects: make(map[game.GameObject] game.GameObject),
	}
}

func (s *simulationSystemImpl) GetElapsedTime() time.Duration {
	return s.elapsed
}

func (s *simulationSystemImpl) SetElapsedTime(elapsed time.Duration) {
	//logrus.Infoln("elapsed time: " + elapsed.String())
	s.elapsed = elapsed
}

func (s *simulationSystemImpl) UpdateGameObjects() {
	for _, g := range s.gameObjects {
		g.Update()
	}
}

func (s *simulationSystemImpl) RegisterGameObject(g game.GameObject) {
	s.gameObjects[g] = g
}

func (s *simulationSystemImpl) UnregisterGameObject(g game.GameObject) {
	delete(s.gameObjects, g)
}
