package game

import (
	"github.com/twatzl/pixel-test/engine/components"
	"github.com/twatzl/pixel-test/engine/services/renderService"
)

type GameObject interface {
	Enable()
	Disable()
	Update()
	Render()
	Destroy()
}

type GameObjectBase struct {
	components []Behavior
	children   []GameObject
	renderer   Renderable
	transform  components.Transform
}

func (g *GameObjectBase) InitBase(renderer Renderable) {
	g.renderer = renderer
	g.transform = components.NewTransform()
}

func (g *GameObjectBase) Render() {
	/* transform */
	renderContext := renderService.Get().GetContext()
	oldTransform := renderContext.GetTransform()
	ct := oldTransform.Chained(g.GetTransform().GetLocalMatrix())
	renderContext.SetTransform(ct)

	/* render */
	g.renderer.Render()
	g.renderChildren()

	/* reset transform */
	renderContext.SetTransform(oldTransform)
}

func (g *GameObjectBase) AddComponent(component Behavior) {
	g.components = append(g.components, component)
}

func (g *GameObjectBase) UpdateComponents() {
	for _, b := range g.components {
		b.Update()
	}
}

func (g *GameObjectBase) renderChildren() {
	for _, o := range g.children {
		o.Render()
	}
}

func (g *GameObjectBase) GetTransform() components.Transform {
	return g.transform
}

func (g *GameObjectBase) GetRenderer() Renderable {
	return g.renderer
}
