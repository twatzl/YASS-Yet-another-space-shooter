package util

import (
	"github.com/faiface/pixel"
	"image"
	"math"
	"os"
)

func LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func LoadSprite(path string) *pixel.Sprite {
	pic, err := LoadPicture(path)
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())
	return sprite
}

func Deg2Rad(angle float64) float64 {
	angle = angle/360 * 2 * math.Pi
	return angle
}

func Rad2Deg(angle float64) float64 {
	angle = angle/(2 * math.Pi) * 360
	return angle
}