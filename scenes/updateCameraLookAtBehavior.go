package scenes

import (
	"github.com/twatzl/pixel-test/engine/camera"
	"github.com/twatzl/pixel-test/engine/component"
)

type updateCameraLookAtBehavior struct {
	cam       camera.Camera
	transform component.Transform
	enabled   bool
}

func NewUpdateCameraLookAtBehavior(cam camera.Camera, transform component.Transform) component.Component {
	return &updateCameraLookAtBehavior{
		enabled: true,
		cam:       cam,
		transform: transform,
	}
}

func (b *updateCameraLookAtBehavior) Init() {}

func (b *updateCameraLookAtBehavior) Enable() {
	b.enabled = true
}

func (b *updateCameraLookAtBehavior) Disable() {
	b.enabled = false
}

func (b *updateCameraLookAtBehavior) Destroy() {}

func (b *updateCameraLookAtBehavior) Update() {
	if !b.enabled {
		return
	}

	v := b.transform.GetTranslation()
	b.cam.SetLookAt(v)
	//println("look at behavior")
}
