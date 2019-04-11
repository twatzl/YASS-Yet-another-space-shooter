package physicsSystem

import "github.com/twatzl/pixel-test/components"

type physicSystemImpl struct {
	config PhysicsConfig
	physicComponents []physicTransformPair
}

type physicTransformPair struct {
	physic *PhysicComponent
	transform components.Transform
}

func NewPhysicsSystem(config PhysicsConfig) PhysicsSystemControl {
	return &physicSystemImpl{
		config: config,
		physicComponents: []physicTransformPair{},
	}
}

func (ps *physicSystemImpl) RegisterPhysicComponent(c *PhysicComponent, t components.Transform) {
	ps.physicComponents = append(ps.physicComponents, physicTransformPair{
		physic: c,
		transform: t,
	})
}

func (ps *physicSystemImpl) Update(deltaT float64) {
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