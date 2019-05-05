package inputSystem

var inputSystem InputSystemControl = nil

func Get() InputSystem {
	return inputSystem
}

func GetControl() InputSystemControl {
	return inputSystem
}

func Provide(system InputSystemControl) {
	inputSystem = system
}
