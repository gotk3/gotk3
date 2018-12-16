// Same copyright and license as the rest of the files in this project

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"

/*
 * GtkTextChildAnchor
 */

// TextChildAnchor is a representation of GTK's GtkTextChildAnchor
type TextChildAnchor C.GtkTextChildAnchor

// native returns a pointer to the underlying GtkTextChildAnchor.
func (v *TextChildAnchor) native() *C.GtkTextChildAnchor {
	return (*C.GtkTextChildAnchor)(v)
}
