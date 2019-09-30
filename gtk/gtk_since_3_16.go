// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14

// See: https://developer.gnome.org/gtk3/3.16/api-index-3-16.html

package gtk

// #include <gtk/gtk.h>
// #include "gtk_since_3_16.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

const(
	POLICY_EXTERNAL  PolicyType = C.GTK_POLICY_EXTERNAL
)

func init() {
	tm := []glib.TypeMarshaler{

		// Objects/Interfaces
		{glib.Type(C.gtk_popover_menu_get_type()), marshalPopoverMenu},
		{glib.Type(C.gtk_model_button_get_type()), marshalModelButton},
		{glib.Type(C.gtk_stack_sidebar_get_type()), marshalStackSidebar},
	}
	glib.RegisterGValueMarshalers(tm)

	//Contribute to casting
	for k, v := range map[string]WrapFn{
		"GtkPopoverMenu":  wrapPopoverMenu,
		"GtkModelButton":  wrapModelButton,
		"GtkStackSidebar": wrapStackSidebar,
	} {
		WrapMap[k] = v
	}
}

// SetOverlayScrolling is a wrapper around gtk_scrolled_window_set_overlay_scrolling().
func (v *ScrolledWindow) SetOverlayScrolling(scrolling bool) {
	C.gtk_scrolled_window_set_overlay_scrolling(v.native(), gbool(scrolling))
}

// GetOverlayScrolling is a wrapper around gtk_scrolled_window_get_overlay_scrolling().
func (v *ScrolledWindow) GetOverlayScrolling() bool {
	return gobool(C.gtk_scrolled_window_get_overlay_scrolling(v.native()))
}

// SetWideHandle is a wrapper around gtk_paned_set_wide_handle().
func (v *Paned) SetWideHandle(wide bool) {
	C.gtk_paned_set_wide_handle(v.native(), gbool(wide))
}

// GetWideHandle is a wrapper around gtk_paned_get_wide_handle().
func (v *Paned) GetWideHandle() bool {
	return gobool(C.gtk_paned_get_wide_handle(v.native()))
}

// GetXAlign is a wrapper around gtk_label_get_xalign().
func (v *Label) GetXAlign() float64 {
	c := C.gtk_label_get_xalign(v.native())
	return float64(c)
}

// GetYAlign is a wrapper around gtk_label_get_yalign().
func (v *Label) GetYAlign() float64 {
	c := C.gtk_label_get_yalign(v.native())
	return float64(c)
}

// SetXAlign is a wrapper around gtk_label_set_xalign().
func (v *Label) SetXAlign(n float64) {
	C.gtk_label_set_xalign(v.native(), C.gfloat(n))
}

// SetYAlign is a wrapper around gtk_label_set_yalign().
func (v *Label) SetYAlign(n float64) {
	C.gtk_label_set_yalign(v.native(), C.gfloat(n))
}

/*
* GtkModelButton
*/

// ModelButton is a representation of GTK's GtkModelButton.
type ModelButton struct {
	Button
 }
 
 func (v *ModelButton) native() *C.GtkModelButton {
	 if v == nil || v.GObject == nil {
		 return nil
	 }
 
	 p := unsafe.Pointer(v.GObject)
	 return C.toGtkModelButton(p)
 }
 
 func marshalModelButton(p uintptr) (interface{}, error) {
	 c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	 return wrapModelButton(glib.Take(unsafe.Pointer(c))), nil
 }
 
 func wrapModelButton(obj *glib.Object) *ModelButton {
	 actionable := wrapActionable(obj)
	 return &ModelButton{Button{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}, actionable}}
 }
 
 // ModelButtonNew is a wrapper around gtk_model_button_new
 func ModelButtonNew() (*ModelButton, error) {
	 c := C.gtk_model_button_new()
	 if c == nil {
		 return nil, nilPtrErr
	 }
	 return wrapModelButton(glib.Take(unsafe.Pointer(c))), nil
 }

/*
 * GtkPopoverMenu
 */

// PopoverMenu is a representation of GTK's GtkPopoverMenu.
type PopoverMenu struct {
	Popover
}

func (v *PopoverMenu) native() *C.GtkPopoverMenu {
	if v == nil || v.GObject == nil {
		return nil
	}

	p := unsafe.Pointer(v.GObject)
	return C.toGtkPopoverMenu(p)
}

func marshalPopoverMenu(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapPopoverMenu(glib.Take(unsafe.Pointer(c))), nil
}

func wrapPopoverMenu(obj *glib.Object) *PopoverMenu {
	return &PopoverMenu{Popover{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}}
}

// PopoverMenuNew is a wrapper around gtk_popover_menu_new
func PopoverMenuNew() (*PopoverMenu, error) {
	c := C.gtk_popover_menu_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapPopoverMenu(glib.Take(unsafe.Pointer(c))), nil
}

// OpenSubmenu is a wrapper around gtk_popover_menu_open_submenu
func (v *PopoverMenu) OpenSubmenu(name string) {
	cstr1 := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(cstr1))

	C.gtk_popover_menu_open_submenu(v.native(), cstr1)
}

/*
 * GtkStackSidebar
 */

// StackSidebar is a representation of GTK's GtkStackSidebar.
type StackSidebar struct {
	Bin
}

// native returns a pointer to the underlying GtkStack.
func (v *StackSidebar) native() *C.GtkStackSidebar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkStackSidebar(p)
}

func marshalStackSidebar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapStackSidebar(obj), nil
}

func wrapStackSidebar(obj *glib.Object) *StackSidebar {
	return &StackSidebar{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// StackSidebarNew is a wrapper around gtk_stack_sidebar_new().
func StackSidebarNew() (*StackSidebar, error) {
	c := C.gtk_stack_sidebar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapStackSidebar(glib.Take(unsafe.Pointer(c))), nil
}

func (v *StackSidebar) SetStack(stack *Stack) {
	C.gtk_stack_sidebar_set_stack(v.native(), stack.native())
}

func (v *StackSidebar) GetStack() *Stack {
	c := C.gtk_stack_sidebar_get_stack(v.native())
	if c == nil {
		return nil
	}
	return wrapStack(glib.Take(unsafe.Pointer(c)))
}

// GrabFocusWithoutSelecting is a wrapper for gtk_entry_grab_focus_without_selecting()
func (v *Entry) GrabFocusWithoutSelecting() {
	C.gtk_entry_grab_focus_without_selecting(v.native())
}

// InsertMarkup() is a wrapper around  gtk_text_buffer_insert_markup()
func (v *TextBuffer) InsertMarkup(start *TextIter, text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_insert_markup(v.native(), (*C.GtkTextIter)(start), (*C.gchar)(cstr), C.gint(len(text)))
}
