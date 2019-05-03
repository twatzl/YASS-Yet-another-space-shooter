package camera

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/twatzl/pixel-test/game"
	"github.com/twatzl/pixel-test/services/renderService"
	"github.com/twatzl/pixel-test/services/windowService"
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
}

func (c *cameraImpl) SetBounds(bounds pixel.Rect) {
	// TODO: improve if newCanvas is too heavyweight for performance
	c.canvas = pixelgl.NewCanvas(bounds)
	c.ctx = renderService.NewRenderContext(c.canvas, bounds)
}

func (c *cameraImpl) SetLookAt(pos pixel.Vec) {
	c.lookAt = pos
}

func (c *cameraImpl) RenderImage(renderable game.Renderable) {
	lookAtMat := lookAt(c.lookAt, 0)
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

func lookAt(pos pixel.Vec, rot float64) pixel.Matrix {
	center := windowService.Get().GetWindow().Bounds().Center()
	offset := center.Scaled(-1)
	offset = offset.Add(pos)
	return calculateViewMatrix(center, 0)
}

func calculateViewMatrix(translation pixel.Vec, rotation float64) pixel.Matrix {
	return pixel.IM.Rotated(pixel.ZV, rotation).Moved(translation)
}
