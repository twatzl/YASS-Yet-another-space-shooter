package gameObjects

import (
	"github.com/faiface/pixel"
	"github.com/twatzl/pixel-test/engine/component"
	"github.com/twatzl/pixel-test/engine/game"
	"github.com/twatzl/pixel-test/engine/util"
)

type background struct {
	game.GameObjectBase
	component.RenderableComponent
	component.GraphicComponent
}

func (background) Destroy() {

}

func (background) Disable() {

}

func (background) Enable() {

}

func (background) Update() {

}

func CreateBackground() *background {
	s := &background{}

	s.InitBase()
	sprite := util.LoadSprite("sprites/background.jpg")
	renderer := game.CreateSpriteRenderer(sprite)
	s.GraphicComponent = component.NewGraphicComponent(renderer, s.GetTransform())

	// TODO: fix scaling
	s.GetTransform().Scale(pixel.V(2,2))

	return s
}