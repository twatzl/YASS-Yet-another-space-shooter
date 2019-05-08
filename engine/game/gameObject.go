package game

import (
	"github.com/twatzl/pixel-test/engine/component"
)

type GameObject interface {
	Enable()
	Disable()
	Update()
	Render()
	Destroy()
	GetTransform() component.Transform
}

type GameObjectBase struct {
	components []component.Component
	children   []GameObject
	transform  component.Transform
	enabled    bool
	//TODO: maybe also allow to remove children?
}

func (g *GameObjectBase) InitBase() {
	g.transform = component.NewTransform()
	g.enabled = true
}

func (g *GameObjectBase) EnableChildren() {
	for _, c := range g.children {
		c.Enable()
	}
}

func (g *GameObjectBase) DisableChildren() {
	for _, c := range g.children {
		c.Disable()
	}
}

func (g *GameObjectBase) EnableComponents() {
	for _, c := range g.components {
		c.Enable()
	}
}

func (g *GameObjectBase) DisableComponents() {
	for _, c := range g.components {
		c.Disable()
	}
}

func (g *GameObjectBase) AddComponent(component component.Component) {
	g.components = append(g.components, component)
}

func (g *GameObjectBase) AddChild(gameObject GameObject) {
	g.children = append(g.children, gameObject)
}

func (g *GameObjectBase) UpdateComponents() {
	for _, b := range g.components {
		b.Update()
	}
}

func (g *GameObjectBase) UpdateChildren() {
	for _, o := range g.children {
		o.Update()
	}
}

func (g *GameObjectBase) DestroyComponents() {
	for _, b := range g.components {
		b.Destroy()
	}
}

func (g *GameObjectBase) DestroyChildren() {
	for _, o := range g.children {
		o.Destroy()
	}
}

func (g *GameObjectBase) RenderChildren() {
	for _, o := range g.children {
		o.Render()
	}
}

func (g *GameObjectBase) GetTransform() component.Transform {
	return g.transform
}