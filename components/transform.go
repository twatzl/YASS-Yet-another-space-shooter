package components

import (
	"github.com/faiface/pixel"
	"math"
)

type Transform interface {
	GetTranslation() pixel.Vec
	GetRotation() float64
	GetScale() pixel.Vec
	GetLocalMatrix() pixel.Matrix

	Rotate(angle float64)
	Move(translation pixel.Vec)
	MoveTo(position pixel.Vec)
	Scale(scale pixel.Vec)
}

type transformImpl struct {
	translation pixel.Vec
	rotation    float64 /* radians */
	scale       pixel.Vec
	localMatrix pixel.Matrix
}

var _ Transform = &transformImpl{}

func NewTransform() Transform {
	return &transformImpl{
		scale: pixel.V(1,1),
		rotation: 0,
		translation: pixel.ZV,
		localMatrix: pixel.IM,
	}
}

func (t *transformImpl) GetTranslation() pixel.Vec {
	return t.translation
}

func (t *transformImpl) GetRotation() float64 {
	return t.rotation
}

func (t *transformImpl) GetScale() pixel.Vec {
	return t.scale
}

func (t *transformImpl) GetLocalMatrix() pixel.Matrix {
	return t.localMatrix
}

func (t *transformImpl) updateTransform() {
	t.localMatrix = pixel.IM.Moved(t.translation).Rotated(t.translation, t.rotation)
}

func (t *transformImpl) Rotate(angle float64) {
	angle = angle/360 * 2 * math.Pi
	t.rotation += angle
	t.updateTransform()
}

func (t *transformImpl) Move(translation pixel.Vec) {
	t.translation = t.translation.Add(translation)
	t.updateTransform()
}

func (t *transformImpl) MoveTo(position pixel.Vec) {
	t.translation = position
	t.updateTransform()
}

func (t *transformImpl) Scale(scale pixel.Vec) {
	t.scale = scale
}
