package sg

type Node interface {
	render(context *renderContext)
	Add(n Node)
}

type sgnode struct {
	children []Node
}

func (node *sgnode) render(context *renderContext) {
	for _, c := range node.children {
		c.render(context)
	}
}

func (node *sgnode) Add(n Node) {
	node.children = append(node.children, n)
}
