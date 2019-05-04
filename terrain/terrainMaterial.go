package terrain

import "image/color"

type TerrainMaterial  struct {
	name      string // mostly for internal purpose
	hardness  int
	color     color.RGBA
	collision bool
}

