package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"errors"
	"os"
	"unsafe"

	"github.com/gotk3/gotk3/gdk"

	"github.com/gotk3/gotk3/glib"
)

// TestFindLabel is a wrapper around gtk_test_find_label().
// This function will search widget and all its descendants for a GtkLabel widget with a text string matching label_pattern.
// The labelPattern may contain asterisks “*” and question marks “?” as placeholders, g_pattern_match() is used for the matching.
func TestFindLabel(widget IWidget, labelPattern string) (IWidget, error) {
	cstr := C.CString(labelPattern)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_test_find_label(widget.toWidget(), (*C.gchar)(cstr))
	if c == nil {
		return nil, errors.New("no label with pattern '" + labelPattern + "' found")
	}
	obj, err := castWidget(c)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// TestFindSibling is a wrapper around gtk_test_find_sibling().
// This function will search siblings of base_widget and siblings of its ancestors for all widgets matching widgetType.
// Of the matching widgets, the one that is geometrically closest to base_widget will be returned.
func TestFindSibling(baseWidget IWidget, widgetType glib.Type) (IWidget, error) {
	c := C.gtk_test_find_sibling(baseWidget.toWidget(), C.GType(widgetType))
	if c == nil {
		return nil, errors.New("no widget of type '" + widgetType.Name() + "' found")
	}
	obj, err := castWidget(c)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// TestFindWidget is a wrapper around gtk_test_find_widget().
// This function will search the descendants of widget for a widget of type widget_type that has a label matching labelPattern next to it.
// This is most useful for automated GUI testing, e.g. to find the “OK” button in a dialog and synthesize clicks on it.
// However see TestFindLabel(), TestFindSibling() and TestWidgetClick() (and their GTK documentation)
// for possible caveats involving the search of such widgets and synthesizing widget events.
func TestFindWidget(widget IWidget, labelPattern string, widgetType glib.Type) (IWidget, error) {
	cstr := C.CString(labelPattern)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_test_find_widget(widget.toWidget(), (*C.gchar)(cstr), C.GType(widgetType))
	if c == nil {
		return nil, errors.New("no widget with label pattern '" + labelPattern + "' and type '" + widgetType.Name() + "' found")
	}
	obj, err := castWidget(c)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

/*
TestInit is a wrapper around gtk_test_init().
This function is used to initialize a GTK+ test program.
It will in turn call g_test_init() and gtk_init() to properly initialize the testing framework and graphical toolkit.
It’ll also set the program’s locale to “C” and prevent loading of rc files and Gtk+ modules.
This is done to make tets program environments as deterministic as possible.

Like gtk_init() and g_test_init(), any known arguments will be processed and stripped from argc and argv.
*/
func TestInit(args *[]string) {
	if args != nil {
		argc := C.int(len(*args))
		argv := C.make_strings(argc)
		defer C.destroy_strings(argv)

		for i, arg := range *args {
			cstr := C.CString(arg)
			C.set_string(argv, C.int(i), (*C.gchar)(cstr))
		}

		C._gtk_test_init((*C.int)(unsafe.Pointer(&argc)),
			(***C.char)(unsafe.Pointer(&argv)))

		unhandled := make([]string, argc)
		for i := 0; i < int(argc); i++ {
			cstr := C.get_string(argv, C.int(i))
			unhandled[i] = goString(cstr)
			C.free(unsafe.Pointer(cstr))
		}
		*args = unhandled
	} else {
		// gtk_test_init does not take nil, we have to use an empty argument list
		// (only containing the first arg, which is the executable name)
		argc := C.int(1)
		argv := C.make_strings(argc)
		defer C.destroy_strings(argv)

		// Add first argument
		cstr := C.CString(os.Args[0])
		C.set_string(argv, C.int(0), (*C.gchar)(cstr))

		C._gtk_test_init((*C.int)(unsafe.Pointer(&argc)),
			(***C.char)(unsafe.Pointer(&argv)))
	}
}

// TestListAllTypes is a wrapper around gtk_test_list_all_types().
// Return the type ids that have been registered after calling TestRegisterAllTypes().
func TestListAllTypes() []glib.Type {
	var types *C.GType
	var clen C.guint

	types = C.gtk_test_list_all_types(&clen)
	defer C.free(unsafe.Pointer(types))

	length := uint(clen)

	typeReturn := make([]glib.Type, length)
	for i := uint(0); i < length; i++ {
		current := (*C.GType)(pointerAtOffset(unsafe.Pointer(types), unsafe.Sizeof(*types), i))
		typeReturn[i] = glib.Type(*current)
	}
	return typeReturn
}

// pointerAtOffset adjusts `arrayPointer` (pointer to the first element of a C array)
// to point at the offset `i`,
// to be able to read the value there without having to go through cgo.
func pointerAtOffset(arrayPointer unsafe.Pointer, elementSize uintptr, offset uint) unsafe.Pointer {
	return unsafe.Pointer(uintptr(arrayPointer) + elementSize*uintptr(offset))
}

// TestRegisterAllTypes is a wrapper around gtk_test_register_all_types().
// Force registration of all core Gtk+ and Gdk object types.
// This allowes to refer to any of those object types via g_type_from_name() after calling this function.
func TestRegisterAllTypes() {
	C.gtk_test_register_all_types()
}

// TestWidgetSendKey is a wrapper around gtk_test_widget_send_key()
//
// This function will generate keyboard press and release events
// in the middle of the first GdkWindow found that belongs to widget.
// For windowless widgets like GtkButton (which returns FALSE from gtk_widget_get_has_window()),
// this will often be an input-only event window.
// For other widgets, this is usually widget->window.
//
// widget: Widget to generate a key press and release on.
// keyval: A Gdk keyboard value.
// modifiers: Keyboard modifiers the event is setup with.
//
// returns: whether all actions neccessary for the key event simulation were carried out successfully.
func TestWidgetSendKey(widget IWidget, keyval uint, modifiers gdk.ModifierType) bool {
	return gobool(C.gtk_test_widget_send_key(widget.toWidget(), C.guint(keyval), C.GdkModifierType(modifiers)))
}
