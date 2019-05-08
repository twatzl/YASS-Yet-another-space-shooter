package physicsSystem

import (
	"github.com/faiface/pixel"
	"github.com/twatzl/pixel-test/engine/component"
	"github.com/twatzl/pixel-test/engine/services/timeService"
)

type physicSystemImpl struct {
	config           PhysicsConfig
	physicComponents []physicTransformPair
}

type physicTransformPair struct {
	physic    *PhysicComponent
	transform component.Transform
}

func NewPhysicsSystem(config PhysicsConfig) PhysicsSystemControl {
	return &physicSystemImpl{
		config:           config,
		physicComponents: []physicTransformPair{},
	}
}

func (ps *physicSystemImpl) RegisterPhysicComponent(c *PhysicComponent, t component.Transform) {
	ps.physicComponents = append(ps.physicComponents, physicTransformPair{
		physic:    c,
		transform: t,
	})
}

func (ps *physicSystemImpl) Update() {
	deltaT := timeService.Get().GetDeltaTime().Seconds()
	for _, c := range ps.physicComponents {
		ps.handleGravity(deltaT, c.physic)
		ps.updateMovement(deltaT, c)
	}
}

func (ps *physicSystemImpl) handleGravity(deltaT float64, c *PhysicComponent) {
	deltaV := ps.config.GravityDirection.Scaled(ps.config.Gravity)
	deltaV = deltaV.Scaled(deltaT)
	c.speed = c.speed.Add(deltaV)
}

func (ps *physicSystemImpl) updateMovement(deltaT float64, pair physicTransformPair) {
	dist := pair.physic.speed.Scaled(deltaT)
	pair.transform.Move(dist)
}

var _ PhysicsSystemControl = &physicSystemImpl{}

func ApplyForce(c *PhysicComponent, force pixel.Vec) {
	if (c.mass == 0) {
		return;
	}

	deltaT := timeService.Get().GetDeltaTime().Seconds()
	/* speed += F/m * deltaT */
	acc := force.Scaled(1 / c.mass)
	deltaV := acc.Scaled(deltaT)
	c.speed = c.speed.Add(deltaV)
}
