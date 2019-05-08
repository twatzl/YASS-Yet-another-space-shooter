package game

import (
	"github.com/faiface/pixel"
	"github.com/twatzl/pixel-test/engine/services/renderService"
)

type spriteRenderer struct {
	enabled         bool
	sprite          *pixel.Sprite
	spriteTransform pixel.Matrix
}

func (r *spriteRenderer) Enable() {
	r.enabled = true
}

func (r *spriteRenderer) Disable() {
	r.enabled = false
}

func (r *spriteRenderer) Destroy() {
	r.Disable()
	r.sprite = nil
}

func CreateSpriteRenderer(sprite *pixel.Sprite) *spriteRenderer {
	sr := &spriteRenderer{}
	sr.SetSprite(sprite)
	sr.spriteTransform = pixel.IM.Moved(sr.sprite.Picture().Bounds().Max.Scaled(0.5))
	sr.Enable()
	return sr
}

func (r *spriteRenderer) SetSprite(sprite *pixel.Sprite) {
	r.sprite = sprite
}

func (r *spriteRenderer) Render() {
	if !r.enabled {
		return
	}
	context := renderService.Get().GetContext()
	transform := context.GetTransform()
	view := context.GetViewMatrix()
	transform = transform.Chained(view)
	r.sprite.Draw(context.GetWin(), transform)
}
