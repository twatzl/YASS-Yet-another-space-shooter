package windowService

import "github.com/faiface/pixel/pixelgl"

type WindowService interface {
	GetWindow() *pixelgl.Window
	CreateWindow(cfg pixelgl.WindowConfig)
}
