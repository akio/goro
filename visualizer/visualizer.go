package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/application"
	"github.com/g3n/engine/window"
)

type VisualizerApp struct {
	application.Application
	geom  geometry.IGeometry
	clock int
}

func (app *VisualizerApp) onKeyDown(eventName string, event interface{}) {
}

func (app *VisualizerApp) onBeforeRender(eventName string, event interface{}) {
	var rot math32.Matrix4
	app.clock = app.clock + 1
	euler := math32.Vector3{0, float32(0.01), 0}
	rot.MakeRotationFromEuler(&euler)
	app.geom.GetGeometry().ApplyMatrix(&rot)
}

func main() {
	app, _ := application.Create(application.Options{
		Title:  "Hello G3N",
		Width:  800,
		Height: 600,
	})

	var vis VisualizerApp
	vis.Application = *app
	vis.clock = 0

	vis.Window().Subscribe(window.OnKeyDown, vis.onKeyDown)
	vis.Subscribe(application.OnBeforeRender, vis.onBeforeRender)

	dataDir := "/home/akio/Projects/rosgo_ws2/src/github.com/akio/goro/hsr_meshes/meshes/"
	test := obj.DecodeObj(dataDir+"head_v1/head_light.obj", dataDir+"head_v1/head_light.mtl")

	// Create a blue torus and add it to the scene
	geom := geometry.NewTorus(1, .4, 12, 32, math32.Pi*2)
	mat := material.NewPhong(math32.NewColor("DarkBlue"))
	torusMesh := graphic.NewMesh(geom, mat)
	app.Scene().Add(torusMesh)
	vis.geom = geom

	// Add lights to the scene
	ambientLight := light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.8)
	app.Scene().Add(ambientLight)
	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 5.0)
	pointLight.SetPosition(1, 0, 2)
	app.Scene().Add(pointLight)

	// Add an axis helper to the scene
	axis := graphic.NewAxisHelper(0.5)
	app.Scene().Add(axis)

	app.CameraPersp().SetPosition(0, 0, 3)
	app.Run()
}
