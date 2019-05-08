package renderer

type Renderer interface {
	Render()
	Enable()
	Disable()
	Destroy()
}