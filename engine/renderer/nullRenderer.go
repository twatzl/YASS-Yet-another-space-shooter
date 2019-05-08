package renderer

type nullRenderer struct {

}

func (*nullRenderer) Render() {
}

func (*nullRenderer) Enable() {
}

func (*nullRenderer) Disable() {
}

func (*nullRenderer) Destroy() {
}

// NewNullRenderer returns a renderer which does nothing.
func NewNullRenderer() Renderer {
	return &nullRenderer{}
}