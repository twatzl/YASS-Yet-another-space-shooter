package game

type Scene interface {
	AddGameObject(o GameObject)
	GetGameObjects() []GameObject
	RenderScene()
}

type SceneBase struct {
	gameObjects []GameObject
}

func (sb *SceneBase) AddGameObject(o GameObject) {
	sb.gameObjects = append(sb.gameObjects, o)
}

func (sb *SceneBase) GetGameObjects() []GameObject {
	return sb.gameObjects
}

func (sb *SceneBase) RenderScene() {
	for _,o := range sb.GetGameObjects() {
		o.Render()
	}
}