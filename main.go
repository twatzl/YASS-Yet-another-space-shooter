package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/twatzl/pixel-test/game"
	"github.com/twatzl/pixel-test/scenes"
	"github.com/twatzl/pixel-test/services/renderService"
	"github.com/twatzl/pixel-test/services/windowService"
	_ "image/png"
	"log"
	"os"
	"time"

	"github.com/faiface/beep/wav"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

	"github.com/twatzl/pixel-test/sg"
)

var streamer beep.StreamSeekCloser = nil
var bgcolor = colornames.Skyblue
var spaceshipTransform *sg.TransformNode = nil

func main() {
	pixelgl.Run(run)
}

func run() {
	windowService.Get().CreateWindow(pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	})

	win := windowService.Get().GetWindow()
	rService := renderService.NewSimpleRenderService(win)
	renderService.ProvideRenderService(rService)
	renderService.Get().GetContext().SetViewMatrix(lookAt(pixel.V(0,0), 0))

	mainScene := scenes.InitMainScene()
	g := game.InitGame(mainScene)


	loadSounds();

	lastTime := time.Now()
	deltaT := time.Since(lastTime).Seconds()

	for !win.Closed() {
		deltaT = time.Since(lastTime).Seconds()
		lastTime = time.Now()

		g.HandleInput(deltaT)

		win.Clear(bgcolor)
		g.Render()
		win.Update()
	}
}

func loadSounds() {
	f, err := os.Open("sound/laser_shooting_sfx.wav")
	if err != nil {
		log.Fatal(err)
	}

	st, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	streamer = st
	_ = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
}

func handleInput(win *pixelgl.Window, deltaTime float64) {
	rotSpeed := 120.0
	if win.JustPressed(pixelgl.KeySpace) {
		_ = streamer.Seek(0);
		speaker.Play(streamer)
		go bgcolorBlink()
	}
	if win.Pressed(pixelgl.KeyLeft) {
		spaceshipTransform.Rotate(rotSpeed * deltaTime)
	}
	if win.Pressed(pixelgl.KeyRight) {
		spaceshipTransform.Rotate(-rotSpeed * deltaTime)
	}
}

func bgcolorBlink(){
	blinkdur := 50 * time.Millisecond
	bgcolor = colornames.Orange
	time.Sleep(blinkdur)
	bgcolor = colornames.Skyblue
	time.Sleep(blinkdur)
	bgcolor = colornames.Orange
	time.Sleep(blinkdur)
	bgcolor = colornames.Skyblue
	time.Sleep(blinkdur)
	bgcolor = colornames.Orange
	time.Sleep(blinkdur)
	bgcolor = colornames.Skyblue
	time.Sleep(blinkdur)

}

func lookAt(pos pixel.Vec, rot float64) pixel.Matrix {
	center := windowService.Get().GetWindow().Bounds().Center()
	offset := center.Scaled(-1)
	offset = offset.Add(pos)
	return calculateViewMatrix(center, 0)
}

func calculateViewMatrix(translation pixel.Vec, rotation float64) pixel.Matrix {
	return pixel.IM.Rotated(pixel.ZV, rotation).Moved(translation)
}