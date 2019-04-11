package sg

import (
	"github.com/faiface/pixel"
)

type SpriteNode struct {
	sgnode
	Sprite *pixel.Sprite
}

func (node *SpriteNode) render(context *renderContext)  {
	node.Sprite.Draw(context.win, context.transform)
}
