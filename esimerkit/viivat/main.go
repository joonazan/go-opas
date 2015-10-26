package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/joonazan/closedgl"
)

func main() {
	closedgl.Run(draw, 640, 640, "Kaksi viivaa")
}

func draw(aika float64) {

	gl.Begin(gl.LINES)

	gl.Color3f(1, 1, 1)
	gl.Vertex2d(0, 0)
	gl.Vertex2d(0.5, 0)

	gl.Color3f(1, 0, 0)
	gl.Vertex2d(0, 0)
	gl.Vertex2d(0, 0.5)

	gl.End()
}
