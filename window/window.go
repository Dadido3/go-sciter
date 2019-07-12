package window

import (
	"runtime"

	"github.com/sciter-sdk/go-sciter"
)

type Window struct {
	*sciter.Sciter
	creationFlags sciter.WindowCreationFlag
}

func init() {
	// Lock main function to main thread
	runtime.LockOSThread()
}

func (w *Window) run() {
	// runtime.LockOSThread()
}
