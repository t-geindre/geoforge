package main

import (
	"geoforge/cam"
	"geoforge/noise"
	"geoforge/render"
	"geoforge/ui"
	"geoforge/world"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	cam      cam.Camera
	world    *world.World
	renderer *render.Renderer
}

func NewGame() *Game {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	seed := rand.Int64N(math.MaxInt64)

	// ---- Warp field (shared) ----
	warpX := noise.NewOpenSimplex(seed + 100)
	warpX = noise.NewFbm(warpX, 2, 0.5, 2.0)
	warpX = noise.NewScale(warpX, 0.001)
	warpX = noise.NewSigned(warpX) // [-1..1]

	warpY := noise.NewOpenSimplex(seed + 101)
	warpY = noise.NewFbm(warpY, 2, 0.5, 2.0)
	warpY = noise.NewScale(warpY, 0.001)
	warpY = noise.NewSigned(warpY) // [-1..1]

	const warpAmount = 200.0 // tune: 100..400

	// ---- Continents mask (0..1) ----
	continent := noise.NewOpenSimplex(seed + 1)
	continent = noise.NewFbm(continent, 3, 0.5, 2.0)
	continent = noise.NewScale(continent, 0.0005)                  // very large shapes
	continent = noise.NewWarp(continent, warpX, warpY, warpAmount) // organic coasts

	landMask := noise.NewSmoothstep(continent, 0.45, 0.55)

	// inland mask (mountains mostly in the interior)
	inland := noise.NewSmoothstep(landMask, 0.85, 0.98)
	inland = noise.NewPow(inland, 2.5)

	// ---- Ocean floor (base around sea level) ----
	// We target a sea level around 0.50 so that rendering thresholds stay intuitive.
	// Ocean range ~ [0.30..0.48]
	ocean := noise.NewOpenSimplex(seed + 2)
	ocean = noise.NewFbm(ocean, 2, 0.5, 2.0)
	ocean = noise.NewScale(ocean, 0.002)
	ocean = noise.NewMulConst(ocean, 0.18) // amplitude
	ocean = noise.NewAddConst(ocean, 0.30) // base

	// ---- Land base relief ----
	// Land range ~ [0.50..0.95] before mountains.
	land := noise.NewOpenSimplex(seed + 3)
	land = noise.NewFbm(land, 6, 0.5, 2.0)
	land = noise.NewScale(land, 0.005)
	land = noise.NewPow(land, 1.8)       // more plains
	land = noise.NewMulConst(land, 0.45) // amplitude
	land = noise.NewAddConst(land, 0.50) // base above sea level

	// ---- Mountains (ridged), gated by inland ----
	mount := noise.NewOpenSimplex(seed + 4)
	mount = noise.NewFbm(mount, 4, 0.5, 2.0)
	mount = noise.NewScale(mount, 0.008)
	mount = noise.NewRidge(mount)
	mount = noise.NewPow(mount, 3.0)       // sharper ridges
	mount = noise.NewMulConst(mount, 0.25) // strength
	mount = noise.NewMul(mount, inland)    // mostly inland

	landPlusMount := noise.NewAdd(land, mount)
	landPlusMount = noise.NewClamp(landPlusMount, 0, 1)

	// ---- Final mix: oceans vs land by continent mask ----
	height := noise.NewMix(ocean, landPlusMount, landMask)
	height = noise.NewClamp(height, 0, 1)

	return &Game{
		cam:      cam.NewWheelZoom(cam.NewMousePan(cam.NewCamera())),
		world:    world.NewWorld(256, 10, 1, 1024, height),
		renderer: render.NewRenderer(),
	}
}

func (g *Game) Update() error {
	g.world.Update(g.cam.WorldRect())
	g.cam.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.Draw(g.world, g.cam, screen)
	ui.DrawPanel(screen, ui.TopLeft,
		"%.0f FPS %.0f TPS\nChunks: %d/%d",
		ebiten.ActualFPS(),
		ebiten.ActualTPS(),
		g.renderer.DrawnChunks(),
		len(g.world.Chunks()),
	)
}

func (g *Game) Layout(x, y int) (int, int) {
	g.cam.SetViewport(x, y)
	return x, y
}
