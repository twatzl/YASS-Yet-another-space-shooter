package scenes

import (
	"github.com/twatzl/pixel-test/camera"
	"github.com/twatzl/pixel-test/components"
	"github.com/twatzl/pixel-test/game"
)

type updateCameraLookAtBehavior struct {
	cam       camera.Camera
	transform components.Transform
}

func NewUpdateCameraLookAtBehavior(cam camera.Camera, transform components.Transform) game.Behavior {
	return &updateCameraLookAtBehavior{
		cam:       cam,
		transform: transform,
	}
}

func (b *updateCameraLookAtBehavior) Init() {}

func (b *updateCameraLookAtBehavior) Update() {
	v := b.transform.GetTranslation()
	b.cam.SetLookAt(v)
	//println("look at behavior")
}
