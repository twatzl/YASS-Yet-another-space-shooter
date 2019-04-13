package physicsSystem

import (
	"github.com/faiface/pixel"
	"github.com/twatzl/pixel-test/components"
)

type PhysicsSystem interface {
	RegisterPhysicComponent(c *PhysicComponent, t components.Transform)
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
