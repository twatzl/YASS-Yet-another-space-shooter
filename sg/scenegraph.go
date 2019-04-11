package sg

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Scenegraph struct {
	children []Node
	Window *pixelgl.Window
}

func (sg *Scenegraph) Render() {
	c := &renderContext{
		win: sg.Window,
		transform: pixel.IM,
	}

	for _,n := range sg.children {
		n.render(c)
	}
}

func (sg *Scenegraph) Add(n Node) {
	sg.children = append(sg.children, n)
}