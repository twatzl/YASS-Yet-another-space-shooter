package physicsSystem

var physicsSystem PhysicsSystemControl = nil

func Provide(service PhysicsSystemControl) {
	physicsSystem = service
}

func GetControl() PhysicsSystemControl {
	return physicsSystem
}

func Get() PhysicsSystem {
	return physicsSystem
}
