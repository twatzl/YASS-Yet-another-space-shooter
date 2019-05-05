package terrain

import (
	"github.com/faiface/pixel"
	"github.com/pkg/errors"
	"github.com/twatzl/pixel-test/engine/systems/collisionSystem"
	"github.com/twatzl/pixel-test/engine/util"
	"image"
	"image/color"
	"os"
	"strings"
)

var collisionBits = 1
var terrainMaterialIndexBits = 4

// golang uses 16 bit colors
var maxColor = 65535

var terrainMaterials = []*TerrainMaterial{
	{
		name:      "empty space",
		color:     color.RGBA{0, 0, 0, 255},
		hardness:  0,
		collision: false,
	},
	{
		name:      "impenetrable stone",
		color:     color.RGBA{127, 127, 127, 255},
		hardness:  10000000000,
		collision: true,
	},
	{
		name:      "grass",
		color:     color.RGBA{70, 128, 70, 255},
		hardness:  10,
		collision: true,
	},
}

type Terrain interface {
	LoadFromImage(path string) error
	GetSprite() *pixel.Sprite
}

type terrainImpl struct {
	bounds pixel.Rect
	// contains 1 if the terrain pixel collides
	collisionBitmap util.BitBoard
	// contains the index of the terrain type
	terrainMaterialType util.BitBoard
	sprite              *pixel.Sprite
}

func (t *terrainImpl) GetBoundingBox() pixel.Rect {
	panic("implement me")
}

func (t *terrainImpl) CollidesWithPoint(x, y int) {
	panic("implement me")
}

func (t *terrainImpl) update() {
	panic("implement me")
}

func (t *terrainImpl) CollidesAt(x int, y int) bool {
	return t.collisionBitmap.GetBool(x, y)
}

func New() Terrain {
	t := &terrainImpl{}

	// register terrain as collider
	collisionSystem.Get().RegisterTerrainCollider(t)

	return t
}

var _ collisionSystem.TerrainCollider = &terrainImpl{}

func (t *terrainImpl) LoadFromImage(path string) error {
	if !strings.HasSuffix(path, ".png") {
		return errors.New("only png files supported for maps")
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	bounds := img.Bounds()

	t.bounds = pixel.R(float64(bounds.Min.X), float64(bounds.Min.Y), float64(bounds.Max.X), float64(bounds.Max.Y))
	t.collisionBitmap = util.NewBitBoard(collisionBits, bounds.Max.X, bounds.Max.Y)
	t.terrainMaterialType = util.NewBitBoard(terrainMaterialIndexBits, bounds.Max.X, bounds.Max.Y)

	for i := bounds.Min.Y; i < bounds.Max.Y; i++ {
		for j := bounds.Min.X; j < bounds.Max.X; j++ {
			lowestDist := maxColor * maxColor * 3 // maximum distance is white to black..
			var pixelMatIndex = -1
			pixel := img.At(j, i)

			for index, mat := range terrainMaterials {
				dist := euclideanDist(pixel, mat.color)
				if dist < lowestDist {
					lowestDist = dist
					pixelMatIndex = index
				}
			}

			pixelMat := terrainMaterials[pixelMatIndex]
			if pixelMat.collision {
				t.collisionBitmap.Set(j, i, 1)
			}

			t.terrainMaterialType.Set(j, i, pixelMatIndex)
		}
	}

	pic := pixel.PictureDataFromImage(img)
	sprite := pixel.NewSprite(pic, pic.Bounds())

	t.sprite = sprite

	return nil
}

func (t *terrainImpl) GetSprite() *pixel.Sprite {
	return t.sprite
}

func euclideanDist(c1, c2 color.Color) int {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()

	// convert values to int, because in the calculation they could become negative
	r1i := int(r1)
	g1i := int(g1)
	b1i := int(b1)
	r2i := int(r2)
	g2i := int(g2)
	b2i := int(b2)

	return (r1i-r2i)*(r1i-r2i) + (g1i-g2i)*(g1i-g2i) + (b1i-b2i)*(b1i-b2i)
}
