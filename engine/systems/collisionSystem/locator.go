package collisionSystem

var collisionSystem CollisionSystemControl = nil

func Provide(service CollisionSystemControl) {
	collisionSystem = service
}

func GetControl() CollisionSystemControl {
	return collisionSystem
}

func Get() CollisionSystem {
	return collisionSystem
}
