package main

import (
	"geoforge/camera"
	"geoforge/noise"
	"geoforge/render"
	"geoforge/world"
)

func noiseSetup() (camera.Camera, *world.World, *render.Renderer) {
	cam := camera.NewWheelZoom(camera.NewMousePan(camera.NewCamera()))
	wrld := world.NewWorld(1, cam)
	nse := noise.NewFastNoise()
	wrld.SetNoise(nse)
	rdr := render.NewRenderer(wrld, cam)

	return cam, wrld, rdr
}
