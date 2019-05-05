package collisionSystem

import (
	"github.com/faiface/pixel"
	"github.com/twatzl/pixel-test/game"
)

type ColliderCenterFunc func() pixel.Vec
type CollisionCallback func(other game.GameObject)

type Collider interface {
}

type CircularCollider interface {
	Collider
	GetRadius() float64
	GetCenter() pixel.Vec
	callCollisionCallback(other game.GameObject)
	getGameObject() game.GameObject
}

type TerrainCollider interface {
	Collider
	CollidesAt(x int, y int) bool
}

type circularCollider struct {
	gameObject  game.GameObject
	radius      float64
	getCenter   ColliderCenterFunc
	boundingBox pixel.Rect
	onCollide   CollisionCallback
}

func NewCircularCollider(gameObject game.GameObject,
	radius float64,
	getCenter ColliderCenterFunc,
	onCollide CollisionCallback) CircularCollider {
	return &circularCollider{
		gameObject: gameObject,
		radius:    radius,
		getCenter: getCenter,
		onCollide: onCollide,
	}
}

func (cc *circularCollider) GetRadius() float64 {
	return cc.radius
}

func (cc *circularCollider) GetCenter() pixel.Vec {
	return cc.getCenter()
}

func (cc *circularCollider) update() {
	center := cc.getCenter()
	bb := pixel.R(center.X-cc.radius, center.Y-cc.radius, center.X+cc.radius, center.Y+cc.radius)
	cc.boundingBox = bb
}

func (cc *circularCollider) GetBoundingBox() pixel.Rect {
	return cc.boundingBox
}

func (cc *circularCollider) CollidesWithPoint(x, y int) {
	panic("method not implemented for circular collider")
}

func (cc *circularCollider) getGameObject() game.GameObject {
	return cc.gameObject
}

func (cc *circularCollider) callCollisionCallback(other game.GameObject) {
	cc.onCollide(other)
}
