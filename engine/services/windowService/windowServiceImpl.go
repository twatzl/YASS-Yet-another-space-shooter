package windowService

import "github.com/faiface/pixel/pixelgl"

type windowServiceImpl struct {
	win *pixelgl.Window
}

func (w *windowServiceImpl) GetWindow() *pixelgl.Window {
	return w.win
}

func (w *windowServiceImpl) CreateWindow(cfg pixelgl.WindowConfig) {
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	w.win = win
}


