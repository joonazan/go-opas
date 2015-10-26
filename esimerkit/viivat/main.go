package main

import (
	"github.com/go-gl/v2.1/gl"
	"github.com/joonazan/closedgl"
)

func main() {
	closedgl.Run(func(dt float64) {
		gl.Begin(gl.LINES)
		gl.End()
	})
}
