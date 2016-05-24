package closedgl

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"log"
	"runtime"
)

func NewWindow(width, height int, title string) *glfw.Window {

	// OpenGL kaatuu jos sitä kutsutaan eri CPUista
	runtime.LockOSThread()

	complain(glfw.Init(), "Initializing GLFW:")

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Samples, 8)
	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	complain(err, "Creating window:")

	// Tämä saa GetKey:n sanomaan että nappula on pohjassa, jos se kävi pohjassa.
	window.SetInputMode(glfw.StickyKeysMode, glfw.True)

	window.MakeContextCurrent()
	complain(gl.Init(), "Initializing OpenGL:")

	return window
}

func complain(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func destroyWindow(w *glfw.Window) {
	w.Destroy()
	glfw.Terminate()
}

func Run(render func(time float64), width, height int, title string) {
	RunInWindow(render, NewWindow(width, height, title))
}

func RunInWindow(render func(time float64), window *glfw.Window) {

	for !window.ShouldClose() {

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		render(glfw.GetTime())

		window.SwapBuffers()
		glfw.PollEvents()
	}

	destroyWindow(window)
}
