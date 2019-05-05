package gameObjects

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/twatzl/pixel-test/engine/game"
	"github.com/twatzl/pixel-test/engine/systems/simulationSystem"
	"github.com/twatzl/pixel-test/engine/systems/collisionSystem"
	"github.com/twatzl/pixel-test/engine/systems/inputSystem"
	"github.com/twatzl/pixel-test/engine/systems/physicsSystem"
	"github.com/twatzl/pixel-test/engine/util"
)

type ship struct {
	game.GameObjectBase
	*physicsSystem.PhysicComponent
	collisionSystem.CircularCollider
}

func (s *ship) Destroy() {

}

func (s *ship) Disable() {
	simulationSystem.Get().UnregisterGameObject(s)
}

func (s *ship) Enable() {
	simulationSystem.Get().RegisterGameObject(s)
}

func (s *ship) Update() {
	s.UpdateComponents()
}

func (s *ship) onCollide(other game.GameObject) {
	//TODO: sound + animation
	println("ship colliding")
}

func CreateShip() *ship {
	s := &ship{}

	sprite := util.LoadSprite("assets/sprites/shuttle.png")
	renderer := game.CreateSpriteRenderer(sprite)
	s.InitBase(renderer)
	s.PhysicComponent = physicsSystem.NewPhysicComponent(200)

	physicsSystem.Get().RegisterPhysicComponent(s.PhysicComponent, s.GetTransform())

	s.CircularCollider = collisionSystem.NewCircularCollider(
		s,
		(sprite.Frame().H() + sprite.Frame().W())/4,
		s.GetTransform().GetTranslation,
		s.onCollide)
	collisionSystem.Get().RegisterCircularCollider(s)

	rotSpeed := 120.0
	rocketWumms := 20000.0

	inputSystem.Get().RegisterKeyEventHandler(inputSystem.KeyPressed, pixelgl.KeyLeft, func() {
		deltaT := simulationSystem.Get().GetElapsedTime().Seconds()
		s.GetTransform().Rotate(rotSpeed * deltaT)

	})

	inputSystem.Get().RegisterKeyEventHandler(inputSystem.KeyPressed, pixelgl.KeyRight, func() {
		deltaT := simulationSystem.Get().GetElapsedTime().Seconds()
		s.GetTransform().Rotate(-rotSpeed * deltaT)

	})

	inputSystem.Get().RegisterKeyEventHandler(inputSystem.KeyPressed, pixelgl.KeyUp, func() {
		rotvec := s.GetTransform().GetRotationVec()
		asdf := rotvec.Scaled(rocketWumms)
		physicsSystem.ApplyForce(s.PhysicComponent, asdf)
	})

	s.Enable()
	return s
}
