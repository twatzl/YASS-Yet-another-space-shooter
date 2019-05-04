package gameObjects

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/twatzl/pixel-test/game"
	"github.com/twatzl/pixel-test/services/simulationService"
	"github.com/twatzl/pixel-test/systems/inputSystem"
	"github.com/twatzl/pixel-test/systems/physicsSystem"
	"github.com/twatzl/pixel-test/util"
)

type ship struct {
	game.GameObjectBase
	*physicsSystem.PhysicComponent
}

func (s *ship) Destroy() {

}

func (s *ship) Disable() {
	simulationService.Get().UnregisterGameObject(s)
}

func (s *ship) Enable() {
	simulationService.Get().RegisterGameObject(s)
}

func (s *ship) Update() {
	s.UpdateComponents()
}

func CreateShip() *ship {
	s := &ship{}

	sprite := util.LoadSprite("sprites/shuttle.png")
	renderer := game.CreateSpriteRenderer(sprite)
	s.InitBase(renderer)
	s.PhysicComponent = physicsSystem.NewPhysicComponent(200)

	physicsSystem.Get().RegisterPhysicComponent(s.PhysicComponent, s.GetTransform())

	rotSpeed := 120.0
	rocketWumms := 20000.0

	inputSystem.Get().RegisterKeyEventHandler(inputSystem.KeyPressed, pixelgl.KeyLeft, func() {
		deltaT := simulationService.Get().GetElapsedTime().Seconds()
		s.GetTransform().Rotate(rotSpeed * deltaT)

	})

	inputSystem.Get().RegisterKeyEventHandler(inputSystem.KeyPressed, pixelgl.KeyRight, func() {
		deltaT := simulationService.Get().GetElapsedTime().Seconds()
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
