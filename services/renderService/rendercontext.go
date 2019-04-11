package renderService

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type RenderContext interface {
	GetWin() *pixelgl.Window
	GetTransform() pixel.Matrix
	SetTransform(matrix pixel.Matrix)
	GetViewMatrix() pixel.Matrix
	SetViewMatrix(matrix pixel.Matrix)
}

type renderContextImpl struct {
	win *pixelgl.Window
	transform pixel.Matrix
	viewMatrix pixel.Matrix
}

var _ RenderContext = &renderContextImpl{}

func (ctx *renderContextImpl) GetWin() *pixelgl.Window {
	return ctx.win
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