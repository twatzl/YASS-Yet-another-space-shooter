package physicsSystem

import (
	"github.com/faiface/pixel"
	"github.com/twatzl/pixel-test/engine/component"
)

type PhysicsSystem interface {
	RegisterPhysicComponent(c *PhysicComponent, t component.Transform)
}

/**
 * The part of the physics system that is called by the game
 * in order to update physics.
 */
type PhysicsSystemControl interface {
	PhysicsSystem
	Update()
}

type PhysicsConfig struct {
	Gravity          float64
	GravityDirection pixel.Vec
}
