package renderer

import "github.com/twatzl/pixel-test/engine/services/renderService"

type animationRenderer struct {
	anim    SpriteSheetAnimation
	enabled bool
}

func NewAnimationRenderer(anim SpriteSheetAnimation) Renderer {
	return &animationRenderer{
		anim: anim,
		enabled: true,
	}
}

func (ar *animationRenderer) Render() {
	if !ar.enabled {
		return
	}

	context := renderService.Get().GetContext()
	transform := context.GetTransform()
	view := context.GetViewMatrix()
	transform = transform.Chained(view)
	ar.anim.GetSprite().Draw(context.GetTarget(), transform)
}

func (ar *animationRenderer) Enable() {
	ar.enabled = true
}

func (ar *animationRenderer) Disable() {
	ar.enabled = false
}

func (ar *animationRenderer) Destroy() {

}
