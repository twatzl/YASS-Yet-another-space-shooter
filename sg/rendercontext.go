package sg

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type renderContext struct {
	win *pixelgl.Window
	transform pixel.Matrix
}

