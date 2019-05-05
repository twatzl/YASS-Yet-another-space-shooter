package game

import (
	"github.com/faiface/pixel"
	"github.com/twatzl/pixel-test/engine/services/renderService"
)

type spriteRenderer struct {
	sprite *pixel.Sprite
}

func (r *spriteRenderer) Create() {
	panic("implement me")
}

func (r *spriteRenderer) Update() {
	panic("implement me")
}

func (r *spriteRenderer) Enable() {
	panic("implement me")
}

func (r *spriteRenderer) Disable() {
	panic("implement me")
}

func (r *spriteRenderer) Destroy() {
	panic("implement me")
}

func CreateSpriteRenderer(sprite *pixel.Sprite) *spriteRenderer {
	sr := &spriteRenderer{}
	sr.SetSprite(sprite)
	return sr
}

func (r *spriteRenderer) SetSprite(sprite *pixel.Sprite) {
	r.sprite = sprite
}

func (r *spriteRenderer) Render() {
	context := renderService.Get().GetContext()
	transform := context.GetTransform()
	view := context.GetViewMatrix()
	transform = transform.Chained(view)
	r.sprite.Draw(context.GetWin(), transform)
}
