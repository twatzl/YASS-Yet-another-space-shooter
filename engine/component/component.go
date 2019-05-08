package component

type Component interface {
	Init()
	Update()
	Enable()
	Disable()
	Destroy()
}