package scenes

import (
	"github.com/faiface/pixel"
	"github.com/twatzl/pixel-test/engine/camera"
	"github.com/twatzl/pixel-test/engine/game"
	"github.com/twatzl/pixel-test/scenes/gameObjects"
	"github.com/twatzl/pixel-test/engine/services/renderService"
	"github.com/twatzl/pixel-test/terrain"
)

type mainScene struct {
	b               game.SceneBase
	camera1         camera.Camera
	camera2         camera.Camera
	c1Pos           pixel.Vec
	c2Pos           pixel.Vec
	terrain         terrain.Terrain
	terrainRenderer game.Renderable
}

func (s *mainScene) GetGameObjects() []game.GameObject {
	return s.b.GetGameObjects()
}

func (s *mainScene) AddGameObject(o game.GameObject) {
	s.b.AddGameObject(o)
}

func (s *mainScene) RenderScene() {
	target := renderService.Get().GetContext().GetTarget()

	s.camera1.RenderImage(s)
	s.camera2.RenderImage(s)

	im1 := s.camera1.GetRenderedImage()
	im2 := s.camera2.GetRenderedImage()

	s1 := pixel.NewSprite(im1, im1.Bounds())
	s2 := pixel.NewSprite(im2, im2.Bounds())

	s1.Draw(target, pixel.IM.Moved(s.c1Pos))
	s2.Draw(target, pixel.IM.Moved(s.c2Pos))
}

func (s *mainScene) Render() {
	s.terrainRenderer.Render()
	s.b.RenderScene()
}

func InitMainScene(targetBounds pixel.Rect) *mainScene {
	ms := &mainScene{}

	camBounds := pixel.R(0,0, targetBounds.Max.X/2, targetBounds.Max.Y)
	ms.camera1 = camera.New(camBounds)
	ms.camera2 = camera.New(camBounds)

	q := targetBounds.Max.X / 4
	ms.c1Pos = pixel.V(q, targetBounds.Max.Y/2)
	ms.c2Pos = pixel.V(targetBounds.Max.X - q, targetBounds.Max.Y/2)

	//ms.camera1.SetLookAt(pixel.V(50,50))
	//ms.camera2.SetLookAt(pixel.V(-50,-50))

	ms.terrain = terrain.New()
	err := ms.terrain.LoadFromImage("assets/maps/map01.png")
	if err != nil {
		println("error loading the map " + err.Error())
		return ms
	}
	ms.terrainRenderer = game.CreateSpriteRenderer(ms.terrain.GetSprite())

	ship := gameObjects.CreateShip()

	//ms.AddGameObject(gameObjects.CreateBackground())
	ms.AddGameObject(ship)

	b := NewUpdateCameraLookAtBehavior(ms.camera1, ship.GetTransform())
	ship.AddComponent(b)

	return ms
}