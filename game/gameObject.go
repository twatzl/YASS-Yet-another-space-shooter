package game

import (
	"github.com/twatzl/pixel-test/components"
	"github.com/twatzl/pixel-test/services/renderService"
)

type GameObject interface {
	Enable()
	Disable()
	Update()
	Render()
	Destroy()
}

type GameObjectBase struct {
	components []Component
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
	oldTransform := renderService.Get().GetContext().GetTransform()
	ct := oldTransform.Chained(g.GetTransform().GetLocalMatrix())
	renderService.Get().GetContext().SetTransform(ct)

	/* render */
	g.renderer.Render()
	g.renderChildren()

	/* reset transform */
	renderService.Get().GetContext().SetTransform(oldTransform)
}

func (g *GameObjectBase) AddBehavior(behavior Component) {
	g.components = append(g.components, behavior)
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
