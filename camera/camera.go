package camera

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/twatzl/pixel-test/game"
	"github.com/twatzl/pixel-test/services/renderService"
	"golang.org/x/image/colornames"
)

type Camera interface {
	SetBounds(bounds pixel.Rect)
	SetLookAt(pos pixel.Vec)
	RenderImage(renderable game.Renderable)
	GetRenderedImage() pixel.Picture
}

type cameraImpl struct {
	lookAt pixel.Vec
	ctx    renderService.RenderContext
	canvas *pixelgl.Canvas
	bounds pixel.Rect
}

func (c *cameraImpl) SetBounds(bounds pixel.Rect) {
	// TODO: improve if newCanvas is too heavyweight for performance
	c.bounds = bounds
	c.canvas = pixelgl.NewCanvas(bounds)
	c.ctx = renderService.NewRenderContext("camera render context", c.canvas, bounds)
}

func (c *cameraImpl) SetLookAt(pos pixel.Vec) {
	c.lookAt = pos
}

func (c *cameraImpl) RenderImage(renderable game.Renderable) {
	c.canvas.Clear(colornames.Black)
	lookAtMat := lookAt(c.bounds, c.lookAt, 0)
	c.ctx.SetViewMatrix(lookAtMat)
	oldCtx := renderService.Get().GetContext()
	renderService.Get().SetContext(c.ctx)
	renderable.Render()
	renderService.Get().SetContext(oldCtx)
}

func (c *cameraImpl) GetRenderedImage() pixel.Picture {
	return pixel.PictureDataFromPicture(c.canvas)
}

func New(bounds pixel.Rect) Camera {
	cam := &cameraImpl{}
	cam.lookAt = pixel.ZV
	cam.SetBounds(bounds)
	return cam
}

func lookAt(camBounds pixel.Rect, pos pixel.Vec, rot float64) pixel.Matrix {
	center := camBounds.Center()
	offset := center.Sub(pos)
	return calculateViewMatrix(offset, 0)
}

func calculateViewMatrix(translation pixel.Vec, rotation float64) pixel.Matrix {
	return pixel.IM.Moved(translation).Rotated(translation, rotation)
}
