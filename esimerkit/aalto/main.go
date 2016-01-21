package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/joonazan/closedgl"
	"math"
)

func main() {
	closedgl.Run(piirrä, 640, 640, "Animoitu siniaalto")
}

func piirrä(aika float64) {

	const pisteitä = 200

	gl.Begin(gl.LINE_STRIP)

	gl.Color3d(
		nollastaYhteen(aika*10),
		nollastaYhteen(aika*7),
		nollastaYhteen(aika*3),
	)

	for i := 0.0; i < pisteitä; i++ {

		x := i / (pisteitä - 1)

		gl.Vertex2d(x*2-1, 0.8*math.Sin(x*8+aika))
	}

	gl.End()
}

func nollastaYhteen(x float64) float64 {
	return (math.Sin(x) + 1) * 0.5
}
