package collisionSystem

import "github.com/faiface/pixel"

func isPointInsideBoundingBox(p pixel.Vec, bb pixel.Rect) bool {
	return p.X >= bb.Min.X && p.X < bb.Max.X && p.Y >= bb.Min.Y && p.Y < bb.Max.X
}