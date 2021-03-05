package gtk

/*
 #include <gtk/gtk.h>
*/
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/internal/callback"
)

//export goBuilderConnect
func goBuilderConnect(
	builder *C.GtkBuilder,
	object *C.GObject,
	signal_name *C.gchar,
	handler_name *C.gchar,
	connect_object *C.GObject,
	flags C.GConnectFlags,
	user_data C.gpointer) {

	builderSignals.Lock()
	signals, ok := builderSignals.m[builder]
	builderSignals.Unlock()

	if !ok {
		panic("no signal mapping defined for this GtkBuilder")
	}

	h := C.GoString((*C.char)(handler_name))
	s := C.GoString((*C.char)(signal_name))

	handler, ok := signals[h]
	if !ok {
		return
	}

	if object == nil {
		panic("unexpected nil object from builder")
	}

	//TODO: figure out a better way to get a glib.Object from a *C.GObject
	gobj := glib.Object{glib.ToGObject(unsafe.Pointer(object))}
	gobj.Connect(s, handler)
}

//export goTreeViewSearchEqualFunc
func goTreeViewSearchEqualFunc(model *C.GtkTreeModel, column C.gint, key *C.gchar, iter *C.GtkTreeIter, data C.gpointer) C.gboolean {
	fn := callback.Get(uintptr(data)).(TreeViewSearchEqualFunc)
	return gbool(fn(
		wrapTreeModel(glib.Take(unsafe.Pointer(model))),
		int(column),
		C.GoString(key),
		&TreeIter{(C.GtkTreeIter)(*iter)},
	))
}

//export goTreeModelFilterVisibleFunc
func goTreeModelFilterVisibleFunc(model *C.GtkTreeModel, iter *C.GtkTreeIter, data C.gpointer) C.gboolean {
	goIter := &TreeIter{(C.GtkTreeIter)(*iter)}
	fn := callback.Get(uintptr(data)).(TreeModelFilterVisibleFunc)
	return gbool(fn(
		wrapTreeModel(glib.Take(unsafe.Pointer(model))),
		goIter,
	))
}

//export goTreeSortableSortFunc
func goTreeSortableSortFunc(model *C.GtkTreeModel, a, b *C.GtkTreeIter, data C.gpointer) C.gint {
	fn := callback.Get(uintptr(data)).(TreeIterCompareFunc)
	return C.gint(fn(
		wrapTreeModel(glib.Take(unsafe.Pointer(model))),
		&TreeIter{(C.GtkTreeIter)(*a)},
		&TreeIter{(C.GtkTreeIter)(*b)},
	))
}

//export goTreeModelForeachFunc
func goTreeModelForeachFunc(model *C.GtkTreeModel, path *C.GtkTreePath, iter *C.GtkTreeIter, data C.gpointer) C.gboolean {
	fn := callback.Get(uintptr(data)).(TreeModelForeachFunc)
	return gbool(fn(
		wrapTreeModel(glib.Take(unsafe.Pointer(model))),
		&TreePath{(*C.GtkTreePath)(path)},
		&TreeIter{(C.GtkTreeIter)(*iter)},
	))
}

//export goTreeSelectionForeachFunc
func goTreeSelectionForeachFunc(model *C.GtkTreeModel, path *C.GtkTreePath, iter *C.GtkTreeIter, data C.gpointer) {
	fn := callback.Get(uintptr(data)).(TreeSelectionForeachFunc)
	fn(
		wrapTreeModel(glib.Take(unsafe.Pointer(model))),
		&TreePath{(*C.GtkTreePath)(path)},
		&TreeIter{(C.GtkTreeIter)(*iter)},
	)
}

//export goTreeSelectionFunc
func goTreeSelectionFunc(selection *C.GtkTreeSelection, model *C.GtkTreeModel, path *C.GtkTreePath, selected C.gboolean, data C.gpointer) C.gboolean {
	fn := callback.Get(uintptr(data)).(TreeSelectionFunc)
	return gbool(fn(
		wrapTreeSelection(glib.Take(unsafe.Pointer(selection))),
		wrapTreeModel(glib.Take(unsafe.Pointer(model))),
		&TreePath{(*C.GtkTreePath)(path)},
		gobool(selected),
	))
}
