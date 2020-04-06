//+build gtk_3_6 gtk_3_8 gtk_3_10 gtk_3_12 gtk_3_14 gtk_3_16 gtk_3_18 gtk_3_20 gtk_3_22 gtk_deprecated

package gtk

// #include <gtk/gtk.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

// GetFocusChain is a wrapper around gtk_container_get_focus_chain().
func (v *Container) GetFocusChain() ([]IWidget, bool, error) {
	var cwlist *C.GList
	c := C.gtk_container_get_focus_chain(v.native(), &cwlist)

	if cwlist == nil {
		return nil, gobool(c), nil
	}

	var widgets []IWidget
	wlist := glib.WrapList(uintptr(unsafe.Pointer(cwlist)))
	for ; wlist.Data() != nil; wlist = wlist.Next() {
		w, ok := wlist.Data().(*Widget)
		if !ok {
			return nil, gobool(c), fmt.Errorf("element is not of type *Widget, got %T", w)
		}
		widget, err := castWidget(w.toWidget())
		if err != nil {
			return nil, gobool(c), err
		}
		widgets = append(widgets, widget)
	}
	return widgets, gobool(c), nil
}

/*
 * GtkContainer
 */

// SetFocusChain is a wrapper around gtk_container_set_focus_chain().
func (v *Container) SetFocusChain(focusableWidgets []IWidget) {
	var list *glib.List
	for _, w := range focusableWidgets {
		data := uintptr(unsafe.Pointer(w.toWidget()))
		list = list.Append(data)
	}
	glist := (*C.GList)(unsafe.Pointer(list))
	C.gtk_container_set_focus_chain(v.native(), glist)
}

// TODO:
// gtk_container_unset_focus_chain

// CssProviderGetDefault is a wrapper around gtk_css_provider_get_default().
func CssProviderGetDefault() (*CssProvider, error) {
	c := C.gtk_css_provider_get_default()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCssProvider(obj), nil
}
