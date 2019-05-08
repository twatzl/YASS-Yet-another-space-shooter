package component

import (
	"github.com/twatzl/pixel-test/engine/renderer"
	"github.com/twatzl/pixel-test/engine/services/renderService"
)

type GraphicComponent interface {
	Component
	RenderableComponent
}

type graphicComponent struct {
	renderer   renderer.Renderer
	enabled    bool
	transform Transform
}

func (gc *graphicComponent) Init() {
	// do nothing
}

func (gc *graphicComponent) initGraphicComponent(renderer renderer.Renderer, transform Transform) {
	gc.transform = transform
	gc.enabled = true
	gc.renderer = renderer
}

func (gc *graphicComponent) Update() {
	// do nothing
}

func (gc *graphicComponent) Enable() {
	gc.enabled = true
	gc.renderer.Enable()
}

func (gc *graphicComponent) Disable() {
	gc.enabled = false
	gc.renderer.Disable()
}

func (gc *graphicComponent) Destroy() {
	gc.renderer.Destroy()
}

func (gc *graphicComponent) Render() {
	if !gc.enabled {
		return
	}

	/* transform */
	renderContext := renderService.Get().GetContext()
	oldTransform := renderContext.GetTransform()
	ct := oldTransform.Chained(gc.transform.GetLocalMatrix())
	renderContext.SetTransform(ct)

	/* render */
	gc.renderer.Render()

	/* reset transform */
	renderContext.SetTransform(oldTransform)
}

func NewGraphicComponent(renderer renderer.Renderer, transform Transform) GraphicComponent {
	gc := &graphicComponent{}
	gc.initGraphicComponent(renderer, transform)
	return gc
}