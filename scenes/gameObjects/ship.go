package gameObjects

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/twatzl/pixel-test/engine/component"
	"github.com/twatzl/pixel-test/engine/game"
	"github.com/twatzl/pixel-test/engine/renderer"
	"github.com/twatzl/pixel-test/engine/services/timeService"
	"github.com/twatzl/pixel-test/engine/systems/simulationSystem"
	"github.com/twatzl/pixel-test/engine/systems/collisionSystem"
	"github.com/twatzl/pixel-test/engine/systems/inputSystem"
	"github.com/twatzl/pixel-test/engine/systems/physicsSystem"
	"github.com/twatzl/pixel-test/engine/util"
	"time"
)

type ship struct {
	game.GameObjectBase
	*physicsSystem.PhysicComponent
	collisionSystem.CircularCollider
	shipRenderer               component.GraphicComponent
	explosionAnimationRenderer component.GraphicComponent
	explosionAnimation         renderer.SpriteSheetAnimation
	onShipDestroyed            func()
}

func (s *ship) Destroy() {
	s.DestroyComponents()
	s.DestroyChildren()
}

func (s *ship) Disable() {
	simulationSystem.Get().UnregisterGameObject(s)
}

func (s *ship) Enable() {
	simulationSystem.Get().RegisterGameObject(s)
}

func (s *ship) Update() {
	s.explosionAnimation.Update()
	s.UpdateComponents()
	s.UpdateChildren()
}

func (s *ship) Render() {
	s.RenderChildren()
	s.shipRenderer.Render()
	s.explosionAnimationRenderer.Render()
}

func (s *ship) onCollide(other game.GameObject) {
	//TODO: sound + animation
	println("ship colliding")
	s.shipRenderer.Disable()
	s.explosionAnimation.Start()
	s.explosionAnimationRenderer.Enable()
	s.CircularCollider.Disable()
}

func (s *ship) onExplosionAnimationFinished() {
	s.explosionAnimationRenderer.Disable()
	s.Disable() // disable ship
	s.onShipDestroyed()
}

func CreateShip(pos pixel.Vec, onShipDestroyed func()) *ship {
	s := &ship{
		onShipDestroyed: onShipDestroyed,
	}
	s.InitBase()

	// initial position
	s.GetTransform().Move(pos)

	// ship graphic
	sprite := util.LoadSprite("assets/sprites/shuttle.png")
	shipRenderer :=  game.CreateSpriteRenderer(sprite)
	s.shipRenderer = component.NewGraphicComponent(shipRenderer, s.GetTransform())

	// explosion animation
	s.explosionAnimation = renderer.NewSpriteSheetAnimation("assets/sprites/explosion.png", 64,64,1*time.Second, 15, s.onExplosionAnimationFinished)
	animRenderer := renderer.NewAnimationRenderer(s.explosionAnimation)
	s.explosionAnimationRenderer = component.NewGraphicComponent(animRenderer, s.GetTransform())
	s.explosionAnimationRenderer.Disable()

	// physics
	s.PhysicComponent = physicsSystem.NewPhysicComponent(200)
	physicsSystem.Get().RegisterPhysicComponent(s.PhysicComponent, s.GetTransform())
	rotSpeed := 120.0
	rocketWumms := 20000.0

	// collision
	s.CircularCollider = collisionSystem.NewCircularCollider(
		s,
		(sprite.Frame().H() + sprite.Frame().W())/4,
		s.GetTransform().GetTranslation,
		s.onCollide)
	collisionSystem.Get().RegisterCircularCollider(s)

	// inputs
	// todo: refactor this
	inputSystem.Get().RegisterKeyEventHandler(inputSystem.KeyPressed, pixelgl.KeyLeft, func() {
		if s.IsEnabled() {
			deltaT := timeService.Get().GetDeltaTime().Seconds()
			s.GetTransform().Rotate(rotSpeed * deltaT)
		}

	})

	inputSystem.Get().RegisterKeyEventHandler(inputSystem.KeyPressed, pixelgl.KeyRight, func() {
		if s.IsEnabled() {
			deltaT := timeService.Get().GetDeltaTime().Seconds()
			s.GetTransform().Rotate(-rotSpeed * deltaT)
		}
	})

	inputSystem.Get().RegisterKeyEventHandler(inputSystem.KeyPressed, pixelgl.KeyUp, func() {
		if s.IsEnabled() {
			rotvec := s.GetTransform().GetRotationVec()
			asdf := rotvec.Scaled(rocketWumms)
			physicsSystem.ApplyForce(s.PhysicComponent, asdf)
		}
	})

	s.Enable()
	return s
}
