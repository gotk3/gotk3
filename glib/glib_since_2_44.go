// Same copyright and license as the rest of the files in this project

// +build !glib_2_40,!glib_2_42

package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
// #include "glib_since_2_44.go.h"
import "C"

/*
 * Application
 */

// GetIsBusy is a wrapper around g_application_get_is_busy().
func (v *Application) GetIsBusy() bool {
	return gobool(C.g_application_get_is_busy(v.native()))
}

/*
 * SimpleAction
 */

// SetStateHint is a wrapper around g_simple_action_set_state_hint
func (v *SimpleAction) SetStateHint(stateHint *Variant) {
	C.g_simple_action_set_state_hint(v.native(), stateHint.native())
}
