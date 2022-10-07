package worldmap

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/likheketo/raycaster/pkg/observer"
)

const (
	MapWidth, MapHeight = 24, 24
)

var colorMap map[int]color.Color = map[int]color.Color{
	0: color.RGBA{255, 0, 0, 255},
	1: color.RGBA{0, 255, 0, 255},
	2: color.RGBA{0, 0, 255, 255},
	3: color.RGBA{255, 255, 0, 255},
	4: color.RGBA{0, 255, 255, 255},
	5: color.RGBA{255, 0, 255, 255},
}

type GameMap [MapHeight][MapWidth]int

var GenericMap GameMap = [24][24]int{
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 2, 2, 2, 2, 2, 0, 0, 0, 0, 3, 0, 3, 0, 3, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, 0, 3, 0, 0, 0, 3, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 2, 2, 0, 2, 2, 0, 0, 0, 0, 3, 0, 3, 0, 3, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 4, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 0, 4, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 0, 0, 0, 0, 5, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 0, 4, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 0, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 4, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}

func (m *GameMap) RenderMap2D(screen *ebiten.Image, observer *observer.Observer, width, height int) {
	mapBox := ebiten.NewImage(width, height)

	tileWidth := float64(width) / MapWidth
	tileHeight := float64(height) / MapHeight

	for y := 0; y < MapHeight; y++ {
		for x := 0; x < MapWidth; x++ {
			clr := colorMap[m[y][x]]
			ebitenutil.DrawRect(mapBox, float64(x)*tileWidth, float64(y)*tileHeight, float64(tileWidth), float64(tileHeight), clr)
		}
	}

	ebitenutil.DrawCircle(mapBox, observer.Position.X*tileWidth, observer.Position.Y*tileHeight, 2, color.White)
	ebitenutil.DrawLine(mapBox, observer.Position.X*tileWidth, observer.Position.Y*tileHeight,
		observer.Position.X*tileWidth+5*observer.Direction.X, observer.Position.Y*tileHeight+5*observer.Direction.Y, color.White)

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(mapBox, op)
}

func (m *GameMap) Rendermap3D(screen *ebiten.Image, obs *observer.Observer, width, height int) {
	for x := 0; x < width; x++ {
		cameraX := 2*float64(x)/float64(width) - 1
		rayDirX := obs.Direction.X + obs.Plane.X*cameraX
		rayDirY := obs.Direction.Y + obs.Plane.Y*cameraX

		// which box we are in
		mapX := int(obs.Position.X)
		mapY := int(obs.Position.Y)

		// length of ray from current position to next x or y side
		var sideDistX, sideDistY float64

		// length of ray from one x or y side to next x or y side
		deltaDistX, deltaDistY := 1e30, 1e30
		if rayDirX != 0 {
			deltaDistX = math.Abs(1 / rayDirX)
		}
		if rayDirY != 0 {
			deltaDistY = math.Abs(1 / rayDirY)
		}

		var perpWallDist float64 // perpendicular distance to wall

		// what direction to step in x or y direction
		var stepX, stepY int
		hit := 0     // was there a wall hit?
		var side int //was North, South or a East, West wall hit?
		if rayDirX < 0 {
			stepX = -1
			sideDistX = (obs.Position.X - float64(mapX)) * deltaDistX
		} else {
			stepX = 1
			sideDistX = (float64(mapX) + 1 - obs.Position.X) * deltaDistX
		}
		if rayDirY < 0 {
			stepY = -1
			sideDistY = (obs.Position.Y - float64(mapY)) * deltaDistY
		} else {
			stepY = 1
			sideDistY = (float64(mapY) + 1 - obs.Position.Y) * deltaDistY
		}

		//perform DDA
		for hit == 0 {
			if sideDistX < sideDistY {
				sideDistX += deltaDistX
				mapX += stepX
				side = 0
			} else {
				sideDistY += deltaDistY
				mapY += stepY
				side = 1
			}
			if m[mapY][mapX] > 0 {
				hit = 1
			}
		}

		// calculate distance projected on camera direction
		if side == 0 {
			perpWallDist = sideDistX - deltaDistX
		} else {
			perpWallDist = sideDistY - deltaDistY
		}

		// calculate height of line to draw on screen
		lineHeight := int(float64(height) / perpWallDist)

		// calculate lowest and highest pixel to fill in the current stripe
		drawStart := -lineHeight/2 + height/2
		if drawStart < 0 {
			drawStart = 0
		}
		drawEnd := lineHeight/2 + height/2
		if drawEnd >= height {
			drawEnd = height - 1
		}
		clr := colorMap[m[mapY][mapX]]
		if side == 1 {
			r, g, b, a := clr.RGBA()
			clr = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a) / 2}
		}
		ebitenutil.DrawLine(screen, float64(x), float64(drawStart), float64(x), float64(drawEnd), clr)
	}
}
