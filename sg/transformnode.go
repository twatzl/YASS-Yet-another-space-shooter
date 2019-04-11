package sg

import (
	"github.com/faiface/pixel"
	"math"
)

type TransformNode struct {
	sgnode
	transformation pixel.Matrix
	rotation float64
	position pixel.Vec
}

func (node *TransformNode) render(context *renderContext) {
	oldTransform := context.transform
	context.transform = node.transformation

	for _, c := range node.children {
		c.render(context)
	}

	context.transform = oldTransform
}

func (node *TransformNode) updateTransform() {
	node.transformation = pixel.IM.Moved(node.position).Rotated(node.position, node.rotation)
}

func (node *TransformNode) Rotate(angle float64) {
	angle = angle/360 * 2 * math.Pi
	node.rotation += angle
	node.updateTransform()
}

func (node *TransformNode) Move(position pixel.Vec) {
	node.position = node.position.Add(position)
	node.updateTransform()
}


