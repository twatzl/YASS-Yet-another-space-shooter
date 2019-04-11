package physicsSystem

import (
	"github.com/faiface/pixel"
)

type PhysicComponent struct {
	speed pixel.Vec
	mass  float64
}

func NewPhysicComponent(mass float64) *PhysicComponent {
	return &PhysicComponent{
		speed: pixel.ZV,
	}
}
