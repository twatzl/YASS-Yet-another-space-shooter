package scenes

import (
	"github.com/twatzl/pixel-test/game"
	"github.com/twatzl/pixel-test/scenes/gameObjects"
)

type mainScene struct {
	b game.SceneBase

}

func (s *mainScene) GetGameObjects() []game.GameObject {
	return s.b.GetGameObjects()
}

func (s *mainScene) AddGameObject(o game.GameObject) {
	s.b.AddGameObject(o)
}

func (s *mainScene) Render() {
	s.b.Render()
}

func InitMainScene() *mainScene {
	ms := &mainScene{}

	ms.AddGameObject(gameObjects.CreateShip())

	return ms
}