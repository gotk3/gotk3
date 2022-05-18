package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"

type MainLoop C.GMainLoop

// native returns a pointer to the underlying GMainContext.
func (v *MainLoop) native() *C.GMainLoop {
	if v == nil {
		return nil
	}
	return (*C.GMainLoop)(v)
}

// MainLoopNew is a wrapper around g_main_loop_new().
func MainLoopNew(ctx *MainContext, isRunning bool) *MainLoop {
	c := C.g_main_loop_new(ctx.native(), gbool(isRunning))
	if c == nil {
		return nil
	}
	return (*MainLoop)(c)
}

// IsRunning is a wrapper around g_main_loop_is_running()
func (v *MainLoop) IsRunning() bool {
	return gobool(C.g_main_loop_is_running(v.native()))
}

// Run is a wrapper around g_main_loop_run()
func (v *MainLoop) Run() {
	C.g_main_loop_run(v.native())
}

// Quit is a wrapper around g_main_loop_quit()
func (v *MainLoop) Quit() {
	C.g_main_loop_quit(v.native())
}
