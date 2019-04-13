package physicsSystem

import (
	"github.com/faiface/pixel"
	"github.com/twatzl/pixel-test/components"
	"github.com/twatzl/pixel-test/services/simulationService"
)

type physicSystemImpl struct {
	config           PhysicsConfig
	physicComponents []physicTransformPair
}

type physicTransformPair struct {
	physic    *PhysicComponent
	transform components.Transform
}

func NewPhysicsSystem(config PhysicsConfig) PhysicsSystemControl {
	return &physicSystemImpl{
		config:           config,
		physicComponents: []physicTransformPair{},
	}
}

func (ps *physicSystemImpl) RegisterPhysicComponent(c *PhysicComponent, t components.Transform) {
	ps.physicComponents = append(ps.physicComponents, physicTransformPair{
		physic:    c,
		transform: t,
	})
}

func (ps *physicSystemImpl) Update() {
	deltaT := simulationService.Get().GetElapsedTime().Seconds()
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

	deltaT := simulationService.Get().GetElapsedTime().Seconds()
	/* speed += F/m * deltaT */
	acc := force.Scaled(1 / c.mass)
	deltaV := acc.Scaled(deltaT)
	c.speed = c.speed.Add(deltaV)
}
