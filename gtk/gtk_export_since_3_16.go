// Same copyright and license as the rest of the files in this project
// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14

package gtk

// #include <gtk/gtk.h>
import "C"
import "github.com/gotk3/gotk3/internal/callback"

//export goListBoxCreateWidgetFuncs
func goListBoxCreateWidgetFuncs(item, userData C.gpointer) {
	fn := callback.Get(uintptr(userData)).(ListBoxCreateWidgetFunc)
	fn(uintptr(item))
}
