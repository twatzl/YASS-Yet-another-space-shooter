package renderService

import (
	"github.com/faiface/pixel"
)

type RenderContext interface {
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
	target       pixel.Target
	targetBounds pixel.Rect
	transform    pixel.Matrix
	viewMatrix   pixel.Matrix
}

var _ RenderContext = &renderContextImpl{}

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
