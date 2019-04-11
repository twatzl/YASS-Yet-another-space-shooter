package gameObjects

import (
	"github.com/twatzl/pixel-test/game"
	"github.com/twatzl/pixel-test/systems/physicsSystem"
	"github.com/twatzl/pixel-test/util"
)

type ship struct {
	game.GameObjectBase
	*physicsSystem.PhysicComponent
}

func (ship) Destroy() {

}

func (ship) Disable() {

}

func (ship) Enable() {

}

func (ship) Update() {

}

func CreateShip() *ship {
	s := &ship{}

	sprite := util.LoadSprite("sprites/shuttle.png")
	renderer := game.CreateSpriteRenderer(sprite)
	s.InitBase(renderer)
	s.PhysicComponent = physicsSystem.NewPhysicComponent(2000)

	physicsSystem.Get().RegisterPhysicComponent(s.PhysicComponent, s.GetTransform())

	return s
}