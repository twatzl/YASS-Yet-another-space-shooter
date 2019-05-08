package simulationSystem

import (
	"github.com/twatzl/pixel-test/engine/game"
)

type SimulationSystem interface {
	RegisterGameObject(g game.GameObject)
	UnregisterGameObject(g game.GameObject)
}

type SimulationSystemControl interface {
	SimulationSystem
	UpdateGameObjects()
}

type simulationSystemImpl struct {
	gameObjects map[game.GameObject] game.GameObject
}

func New() SimulationSystemControl {
	return &simulationSystemImpl{
		gameObjects: make(map[game.GameObject] game.GameObject),
	}
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
