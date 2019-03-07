//+build gtk_3_6 gtk_3_8 gtk_3_10 gtk_3_12 gtk_3_14 gtk_3_16 gtk_3_18 gtk_3_20 gtk_3_22

package gtk

// #include <gtk/gtk.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

// GetFocusChain is a wrapper around gtk_container_get_focus_chain().
func (v *Container) GetFocusChain() ([]*Widget, bool) {
	var cwlist *C.GList
	c := C.gtk_container_get_focus_chain(v.native(), &cwlist)

	var widgets []*Widget
	wlist := glib.WrapList(uintptr(unsafe.Pointer(cwlist)))
	for ; wlist.Data() != nil; wlist = wlist.Next() {
		widgets = append(widgets, wrapWidget(glib.Take(wlist.Data().(unsafe.Pointer))))
	}
	return widgets, gobool(c)
}

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

// CssProviderGetDefault is a wrapper around gtk_css_provider_get_default().
func CssProviderGetDefault() (*CssProvider, error) {
	c := C.gtk_css_provider_get_default()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCssProvider(obj), nil
}
