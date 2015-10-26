package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/joonazan/closedgl"
	"math"
)

func main() {
	closedgl.Run(draw, 640, 640, "Animoitu siniaalto")
}

func draw(aika float64) {

	const pisteitä = 200

	gl.Begin(gl.LINE_STRIP)

	gl.Color3d(siniväri(aika*10), siniväri(aika*7), siniväri(aika*3))

	kerroin := 1.0 / (pisteitä - 1)
	for i := 0.0; i < pisteitä; i++ {
		x := i * kerroin
		gl.Vertex2d(x*2-1, 0.8*math.Sin(x*8+aika))
	}

	gl.End()
}

func siniväri(x float64) float64 {
	return (math.Sin(x) + 1) * 0.5
}
