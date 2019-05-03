package renderService

import (
	"github.com/faiface/pixel"
)

type simpleRenderService struct {
	 ctx RenderContext
}

func (s *simpleRenderService) GetContext() RenderContext {
	return s.ctx
}

func (s *simpleRenderService) SetContext(ctx RenderContext) {
	s.ctx = ctx
}

func NewSimpleRenderService(target pixel.Target, bounds pixel.Rect) RenderService {
	return &simpleRenderService{
		ctx: NewRenderContext(target, bounds),
	}
}

func NewRenderContext(target pixel.Target, bounds pixel.Rect) RenderContext {
	return &renderContextImpl{
		target:    target,
		targetBounds: bounds,
		transform: pixel.IM,
	}
}