package renderService

import (
	"github.com/faiface/pixel"
)

type RenderContext interface {
	// for debugging
	GetName() string
	/* deprecated: use GetTarget instead */
	GetWin() pixel.Target
	GetTarget() pixel.Target
	GetTargetBounds() pixel.Rect
	GetTransform() pixel.Matrix
	SetTransform(matrix pixel.Matrix)
	GetViewMatrix() pixel.Matrix
	SetViewMatrix(matrix pixel.Matrix)
}

type renderContextImpl struct {
	// for debugging
	name         string
	target       pixel.Target
	targetBounds pixel.Rect
	transform    pixel.Matrix
	viewMatrix   pixel.Matrix
}

// NewRenderContext instantiates a new render context with a target bounds and a given name.
//	The name can be used for debugging purposes.
func NewRenderContext(name string, target pixel.Target, bounds pixel.Rect) RenderContext {
	return &renderContextImpl{
		name:         name,
		target:       target,
		targetBounds: bounds,
		transform:    pixel.IM,
	}
}

func (ctx *renderContextImpl) GetName() string {
	return ctx.name
}

func (ctx *renderContextImpl) GetWin() pixel.Target {
	return ctx.target
}

func (ctx *renderContextImpl) GetTarget() pixel.Target {
	return ctx.target
}

func (ctx *renderContextImpl) GetTargetBounds() pixel.Rect {
	return ctx.targetBounds
}

func (ctx *renderContextImpl) GetTransform() pixel.Matrix {
	return ctx.transform
}

func (ctx *renderContextImpl) SetTransform(matrix pixel.Matrix) {
	ctx.transform = matrix
}

func (ctx *renderContextImpl) GetViewMatrix() pixel.Matrix {
	return ctx.viewMatrix
}

func (ctx *renderContextImpl) SetViewMatrix(matrix pixel.Matrix) {
	ctx.viewMatrix = matrix
}
