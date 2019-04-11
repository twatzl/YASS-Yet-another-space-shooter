package gameObjects

import (
	"github.com/faiface/pixel"
	"github.com/twatzl/pixel-test/game"
	"github.com/twatzl/pixel-test/util"
)

type background struct {
	game.GameObjectBase
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

	sprite := util.LoadSprite("sprites/background.jpg")
	renderer := game.CreateSpriteRenderer(sprite)
	s.InitBase(renderer)

	// TODO: fix scaling
	s.GetTransform().Scale(pixel.V(2,2))

	return s
}