package main

import (
	"awesomeProject/cam"
	"awesomeProject/render"
	"awesomeProject/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	cam      *cam.Camera
	world    *world.World
	renderer *render.Renderer
}

func NewGame() *Game {
	return &Game{
		cam:      cam.NewCamera(),
		world:    world.NewWorld(256, 2, 1024),
		renderer: render.NewRenderer(),
	}
}

func (g *Game) Update() error {
	g.world.Update(g.cam)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.Draw(g.world, g.cam, screen)
}

func (g *Game) Layout(x, y int) (int, int) {
	g.cam.SetViewport(x, y)
	return x, y
}
