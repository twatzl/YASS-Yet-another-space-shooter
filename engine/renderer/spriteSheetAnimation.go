package renderer

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/twatzl/pixel-test/engine/services/timeService"
	"github.com/twatzl/pixel-test/engine/util"
	"time"
)

type SpriteSheetAnimation interface {
	GetSprite() *pixel.Sprite
	Start()
	Stop()
	Reset()
	Update()
}

type spriteSheetAnimationImpl struct {
	framesPerRow int
	currentFrame int
	start        time.Duration
	duration     time.Duration
	endFrame     int
	frameWidth   int
	frameHeight  int
	running      bool
	sprites      []*pixel.Sprite
	onFinished   func()
}

// NewSpriteSheetAnimation creates a new sprite sheet animation from a sprite sheet file.
// file - the path to the sprite sheet image
// frameWidth - the width of a single animation frame in a spritesheet
// frameHeight - the height of a single animation frame in a spritesheet
// duration - the duration of the animation in milliseconds
// endFrame - the last frame of the animation in the sprite sheet
func NewSpriteSheetAnimation(file string, frameWidth, frameHeight int, duration time.Duration, endFrame int, onAnimationFinished func()) SpriteSheetAnimation {
	anim := &spriteSheetAnimationImpl{
		duration:    duration,
		endFrame:    endFrame,
		frameWidth:  frameWidth,
		frameHeight: frameHeight,
		running:     false,
		sprites:     make([]*pixel.Sprite, endFrame+1),
		onFinished: onAnimationFinished,
	}

	img, err := util.LoadPicture(file)
	if err != nil {
		println("error loading animation image. " + err.Error())
	}

	framesPerRow := int(img.Bounds().Max.X) / frameWidth

	for i := 0; i <= endFrame; i++ {
		row := i / framesPerRow
		frame := i % framesPerRow
		x := frame * frameWidth
		y := row * frameHeight

		sprite := pixel.NewSprite(img, pixel.R(float64(x), float64(y), float64(x+frameWidth), float64(y+frameHeight)))
		anim.sprites[i] = sprite
	}

	anim.Reset()

	return anim
}

func (anim *spriteSheetAnimationImpl) GetSprite() *pixel.Sprite {
	return anim.sprites[anim.currentFrame]
}

func (anim *spriteSheetAnimationImpl) Start() {
	fmt.Println("anim start")
	anim.start = timeService.Get().GetElapsedTime()
	anim.running = true
}

func (anim *spriteSheetAnimationImpl) Stop() {
	anim.running = false
}

func (anim *spriteSheetAnimationImpl) Reset() {
	anim.currentFrame = 0
}

func (anim *spriteSheetAnimationImpl) Destroy() {
	// TODO
}

func (anim *spriteSheetAnimationImpl) Update() {
	if !anim.running {
		return
	}

	duration := anim.duration
	elapsedTime := timeService.Get().GetElapsedTime() - anim.start
	fmt.Printf("globalElapsed: %f, anim start: %f, elapsed: %f\n", timeService.Get().GetElapsedTime().Seconds(), anim.start.Seconds(), elapsedTime.Seconds())
	anim.currentFrame = int(elapsedTime.Seconds()/duration.Seconds() * float64(anim.endFrame))

	fmt.Printf("current frame: %d\n", anim.currentFrame)

	if anim.currentFrame > anim.endFrame {
		anim.currentFrame = anim.endFrame
		anim.onFinished()
		anim.Stop()
	}
}
