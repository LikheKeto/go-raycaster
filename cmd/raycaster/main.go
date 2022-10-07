package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/likheketo/raycaster/pkg/observer"
	"github.com/likheketo/raycaster/pkg/worldmap"
)

const (
	screenWidth  = 1024
	screenHeight = 512
)

type Game struct {
	gameMap                  worldmap.GameMap
	obs                      observer.Observer
	moveSpeed, rotationSpeed float64
	showMap                  bool
}

func NewGame() *Game {
	return &Game{
		gameMap: worldmap.GenericMap,
		obs: observer.Observer{
			Position: observer.Vector{
				X: 22,
				Y: 12,
			},
			Direction: observer.Vector{
				X: -1,
				Y: 0,
			},
			Plane: observer.Vector{
				X: 0,
				Y: 0.66,
			},
		},
	}
}

func (g *Game) Update() error {
	frameTime := float64(1) / ebiten.ActualFPS()
	if frameTime == math.Inf(1) {
		return nil
	}
	g.moveSpeed = frameTime * 3
	g.rotationSpeed = frameTime * 2
	g.HandleMovement()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.gameMap.Rendermap3D(screen, &g.obs, screenWidth, screenHeight)
	if g.showMap {
		g.gameMap.RenderMap2D(screen, &g.obs, 120, 120)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Raycaster")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) HandleMovement() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if g.gameMap[int(g.obs.Position.Y)][int(g.obs.Position.X+g.obs.Direction.X*g.moveSpeed)] == 0 {
			g.obs.Position.X += g.obs.Direction.X * g.moveSpeed
		}
		if g.gameMap[int(g.obs.Position.Y+g.obs.Direction.Y*g.moveSpeed)][int(g.obs.Position.X)] == 0 {
			g.obs.Position.Y += g.obs.Direction.Y * g.moveSpeed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		if g.gameMap[int(g.obs.Position.Y)][int(g.obs.Position.X-g.obs.Direction.X*g.moveSpeed)] == 0 {
			g.obs.Position.X -= g.obs.Direction.X * g.moveSpeed
		}
		if g.gameMap[int(g.obs.Position.Y-g.obs.Direction.Y*g.moveSpeed)][int(g.obs.Position.X)] == 0 {
			g.obs.Position.Y -= g.obs.Direction.Y * g.moveSpeed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.obs.Rotate(g.rotationSpeed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.obs.Rotate(-g.rotationSpeed)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		g.showMap = !g.showMap
	}
}
