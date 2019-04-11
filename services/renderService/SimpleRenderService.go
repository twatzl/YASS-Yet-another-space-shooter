package renderService

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type simpleRenderService struct {
	 ctx *renderContextImpl
}

func (s *simpleRenderService) GetContext() RenderContext {
	return s.ctx
}

func NewSimpleRenderService(win *pixelgl.Window) RenderService {
	return &simpleRenderService{
		ctx: &renderContextImpl{
			win: win,
			transform: pixel.IM,
		},
	}
}

