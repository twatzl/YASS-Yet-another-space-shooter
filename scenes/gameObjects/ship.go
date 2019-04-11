package gameObjects

import (
	"github.com/twatzl/pixel-test/game"
	"github.com/twatzl/pixel-test/util"
)

type ship struct {
	b *game.GameObjectBase
}

func (ship) Destroy() {

}

func (ship) Disable() {

}

func (ship) Enable() {

}

func (s *ship) Render() {
	s.b.Render()
}

func (ship) Update() {

}

func CreateShip() *ship {
	s := &ship{}

	sprite := util.LoadSprite("sprites/shuttle.png")
	renderer := game.CreateSpriteRenderer(sprite)
	s.b = game.InitGameObjectBase(renderer)

	return s
}