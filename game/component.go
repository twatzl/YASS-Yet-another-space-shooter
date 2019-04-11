package game

type Component interface {
	Create(o GameObject)
	Update()
	Enable()
	Disable()
	Destroy()
}