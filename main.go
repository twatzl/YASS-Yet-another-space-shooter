package main

import (
	"github.com/twatzl/pixel-test/game"
	"github.com/twatzl/pixel-test/scenes"
	"github.com/twatzl/pixel-test/services/renderService"
	"github.com/twatzl/pixel-test/services/simulationService"
	"github.com/twatzl/pixel-test/services/windowService"
	"github.com/twatzl/pixel-test/systems/inputSystem"
	"github.com/twatzl/pixel-test/systems/physicsSystem"
	_ "image/png"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var bgcolor = colornames.Skyblue

func main() {
	pixelgl.Run(run)
}

func run() {
	win := initServicesAndSystems()

	mainScene := scenes.InitMainScene(win.Bounds())
	g := game.InitGame(mainScene)

	lastTime := time.Now()
	deltaT := time.Since(lastTime)

	for !win.Closed() {
		deltaT = time.Since(lastTime)
		lastTime = time.Now()
		simulationService.GetControl().SetElapsedTime(deltaT)

		/* update physics */
		physicsSystem.GetControl().Update()

		/* handle inputs */
		inputSystem.GetControl().HandleInputs()

		/* update the game logic */
		simulationService.GetControl().UpdateGameObjects()

		/* render */
		win.Clear(bgcolor)

		g.Render()

		win.Update()
	}
}

func initServicesAndSystems() *pixelgl.Window {
	windowService.Get().CreateWindow(pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	})
	win := windowService.Get().GetWindow()

	rService := renderService.NewSimpleRenderService(win, win.Bounds())
	renderService.ProvideRenderService(rService)

	sService := simulationService.New()
	simulationService.Provide(sService)

	iSystem := inputSystem.New()
	inputSystem.Provide(iSystem)

	ps := physicsSystem.NewPhysicsSystem(physicsSystem.PhysicsConfig{
		Gravity:          9.81,
		GravityDirection: pixel.V(0, -1),
	})
	physicsSystem.Provide(ps)
	return win
}


/*
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
	if win.JustPressed(pixelgl.KeySpace) {
		_ = streamer.Seek(0)
		speaker.Play(streamer)
		go bgcolorBlink()
	}
	rotSpeed := 120.0
	if win.Pressed(pixelgl.KeyLeft) {
		spaceshipTransform.Rotate(rotSpeed * deltaTime)
	}
	if win.Pressed(pixelgl.KeyRight) {
		spaceshipTransform.Rotate(-rotSpeed * deltaTime)
	}

}*/

func bgcolorBlink() {
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
