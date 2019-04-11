package game

type game struct {
	currentScene Scene
}

func (g *game) setCurrentScene(scene Scene) {
	g.currentScene = scene
}

func (g *game) Render() {

	g.currentScene.Render()
}

func (g *game) HandleInput(f float64) {
	// TODO: fixme
}

func InitGame(s Scene) *game {
	g := &game{}
	g.setCurrentScene(s)
	return g
}