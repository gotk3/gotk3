// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

package impl

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	gdk_impl "github.com/gotk3/gotk3/gdk/impl"
	glib_impl "github.com/gotk3/gotk3/glib/impl"
	"github.com/gotk3/gotk3/gtk"
)

/*
 * GtkTreeView
 */

// TreeView is a representation of GTK's GtkTreeView.
type treeView struct {
	container
}

// native returns a pointer to the underlying GtkTreeView.
func (v *treeView) native() *C.GtkTreeView {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTreeView(p)
}

func marshalTreeView(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapTreeView(obj), nil
}

func wrapTreeView(obj *glib_impl.Object) *treeView {
	return &treeView{container{widget{glib_impl.InitiallyUnowned{obj}}}}
}

func setupTreeView(c unsafe.Pointer) (*treeView, error) {
	if c == nil {
		return nil, nilPtrErr
	}

	return wrapTreeView(wrapObject(c)), nil
}

// TreeViewNew() is a wrapper around gtk_tree_view_new().
func TreeViewNew() (*treeView, error) {
	return setupTreeView(unsafe.Pointer(C.gtk_tree_view_new()))
}

// TreeViewNewWithModel() is a wrapper around gtk_tree_view_new_with_model().
func TreeViewNewWithModel(model ITreeModel) (*treeView, error) {
	return setupTreeView(unsafe.Pointer(C.gtk_tree_view_new_with_model(model.toTreeModel())))
}

// GetModel() is a wrapper around gtk_tree_view_get_model().
func (v *treeView) GetModel() (gtk.TreeModel, error) {
	c := C.gtk_tree_view_get_model(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapTreeModel(wrapObject(unsafe.Pointer(c))), nil
}

// SetModel() is a wrapper around gtk_tree_view_set_model().
func (v *treeView) SetModel(model gtk.TreeModel) {
	C.gtk_tree_view_set_model(v.native(), model.(ITreeModel).toTreeModel())
}

// GetSelection() is a wrapper around gtk_tree_view_get_selection().
func (v *treeView) GetSelection() (gtk.TreeSelection, error) {
	c := C.gtk_tree_view_get_selection(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapTreeSelection(wrapObject(unsafe.Pointer(c))), nil
}

// AppendColumn() is a wrapper around gtk_tree_view_append_column().
func (v *treeView) AppendColumn(column gtk.TreeViewColumn) int {
	c := C.gtk_tree_view_append_column(v.native(), castToTreeViewColumn(column).native())
	return int(c)
}

// GetPathAtPos() is a wrapper around gtk_tree_view_get_path_at_pos().
func (v *treeView) GetPathAtPos(x, y int, path gtk.TreePath, column gtk.TreeViewColumn, cellX, cellY *int) bool {
	var ctp **C.GtkTreePath
	if path != nil {
		ctp = (**C.GtkTreePath)(unsafe.Pointer(&castToTreePath(path).GtkTreePath))
	} else {
		ctp = nil
	}

	var pctvcol **C.GtkTreeViewColumn
	if column != nil {
		ctvcol := castToTreeViewColumn(column).native()
		pctvcol = &ctvcol
	} else {
		pctvcol = nil
	}

	return 0 != C.gtk_tree_view_get_path_at_pos(
		v.native(),
		(C.gint)(x),
		(C.gint)(y),
		ctp,
		pctvcol,
		(*C.gint)(unsafe.Pointer(cellX)),
		(*C.gint)(unsafe.Pointer(cellY)))
}

// GetLevelIndentation is a wrapper around gtk_tree_view_get_level_indentation().
func (v *treeView) GetLevelIndentation() int {
	return int(C.gtk_tree_view_get_level_indentation(v.native()))
}

// GetShowExpanders is a wrapper around gtk_tree_view_get_show_expanders().
func (v *treeView) GetShowExpanders() bool {
	return gobool(C.gtk_tree_view_get_show_expanders(v.native()))
}

// SetLevelIndentation is a wrapper around gtk_tree_view_set_level_indentation().
func (v *treeView) SetLevelIndentation(indent int) {
	C.gtk_tree_view_set_level_indentation(v.native(), C.gint(indent))
}

// SetShowExpanders is a wrapper around gtk_tree_view_set_show_expanders().
func (v *treeView) SetShowExpanders(show bool) {
	C.gtk_tree_view_set_show_expanders(v.native(), gbool(show))
}

// GetHeadersVisible is a wrapper around gtk_tree_view_get_headers_visible().
func (v *treeView) GetHeadersVisible() bool {
	return gobool(C.gtk_tree_view_get_headers_visible(v.native()))
}

// SetHeadersVisible is a wrapper around gtk_tree_view_set_headers_visible().
func (v *treeView) SetHeadersVisible(show bool) {
	C.gtk_tree_view_set_headers_visible(v.native(), gbool(show))
}

// ColumnsAutosize is a wrapper around gtk_tree_view_columns_autosize().
func (v *treeView) ColumnsAutosize() {
	C.gtk_tree_view_columns_autosize(v.native())
}

// GetHeadersClickable is a wrapper around gtk_tree_view_get_headers_clickable().
func (v *treeView) GetHeadersClickable() bool {
	return gobool(C.gtk_tree_view_get_headers_clickable(v.native()))
}

// SetHeadersClickable is a wrapper around gtk_tree_view_set_headers_clickable().
func (v *treeView) SetHeadersClickable(show bool) {
	C.gtk_tree_view_set_headers_clickable(v.native(), gbool(show))
}

// GetActivateOnSingleClick is a wrapper around gtk_tree_view_get_activate_on_single_click().
func (v *treeView) GetActivateOnSingleClick() bool {
	return gobool(C.gtk_tree_view_get_activate_on_single_click(v.native()))
}

// SetActivateOnSingleClick is a wrapper around gtk_tree_view_set_activate_on_single_click().
func (v *treeView) SetActivateOnSingleClick(show bool) {
	C.gtk_tree_view_set_activate_on_single_click(v.native(), gbool(show))
}

// RemoveColumn() is a wrapper around gtk_tree_view_remove_column().
func (v *treeView) RemoveColumn(column gtk.TreeViewColumn) int {
	return int(C.gtk_tree_view_remove_column(v.native(), castToTreeViewColumn(column).native()))
}

// InsertColumn() is a wrapper around gtk_tree_view_insert_column().
func (v *treeView) InsertColumn(column gtk.TreeViewColumn, pos int) int {
	return int(C.gtk_tree_view_insert_column(v.native(), castToTreeViewColumn(column).native(), C.gint(pos)))
}

// GetNColumns() is a wrapper around gtk_tree_view_get_n_columns().
func (v *treeView) GetNColumns() uint {
	return uint(C.gtk_tree_view_get_n_columns(v.native()))
}

// GetColumn() is a wrapper around gtk_tree_view_get_column().
func (v *treeView) GetColumn(n int) gtk.TreeViewColumn {
	c := C.gtk_tree_view_get_column(v.native(), C.gint(n))
	if c == nil {
		return nil
	}
	return wrapTreeViewColumn(wrapObject(unsafe.Pointer(c)))
}

// MoveColumnAfter() is a wrapper around gtk_tree_view_move_column_after().
func (v *treeView) MoveColumnAfter(column gtk.TreeViewColumn, baseColumn gtk.TreeViewColumn) {
	C.gtk_tree_view_move_column_after(v.native(), castToTreeViewColumn(column).native(), castToTreeViewColumn(baseColumn).native())
}

// SetExpanderColumn() is a wrapper around gtk_tree_view_set_expander_column().
func (v *treeView) SetExpanderColumn(column gtk.TreeViewColumn) {
	C.gtk_tree_view_set_expander_column(v.native(), castToTreeViewColumn(column).native())
}

// GetExpanderColumn() is a wrapper around gtk_tree_view_get_expander_column().
func (v *treeView) GetExpanderColumn() gtk.TreeViewColumn {
	c := C.gtk_tree_view_get_expander_column(v.native())
	if c == nil {
		return nil
	}
	return wrapTreeViewColumn(wrapObject(unsafe.Pointer(c)))
}

// ScrollToPoint() is a wrapper around gtk_tree_view_scroll_to_point().
func (v *treeView) ScrollToPoint(treeX, treeY int) {
	C.gtk_tree_view_scroll_to_point(v.native(), C.gint(treeX), C.gint(treeY))
}

// SetCursor() is a wrapper around gtk_tree_view_set_cursor().
func (v *treeView) SetCursor(path gtk.TreePath, focusColumn gtk.TreeViewColumn, startEditing bool) {
	C.gtk_tree_view_set_cursor(v.native(), castToTreePath(path).native(), castToTreeViewColumn(focusColumn).native(), gbool(startEditing))
}

// SetCursorOnCell() is a wrapper around gtk_tree_view_set_cursor_on_cell().
func (v *treeView) SetCursorOnCell(path gtk.TreePath, focusColumn gtk.TreeViewColumn, focusCell gtk.CellRenderer, startEditing bool) {
	C.gtk_tree_view_set_cursor_on_cell(v.native(), castToTreePath(path).native(), castToTreeViewColumn(focusColumn).native(), castToCellRenderer(focusCell).native(), gbool(startEditing))
}

// GetCursor() is a wrapper around gtk_tree_view_get_cursor().
func (v *treeView) GetCursor() (p gtk.TreePath, c gtk.TreeViewColumn) {
	var path *C.GtkTreePath
	var col *C.GtkTreeViewColumn

	C.gtk_tree_view_get_cursor(v.native(), &path, &col)

	if path != nil {
		p = &treePath{path}
		runtime.SetFinalizer(p, (*treePath).free)
	}

	if col != nil {
		c = wrapTreeViewColumn(wrapObject(unsafe.Pointer(col)))
	}

	return
}

// RowActivated() is a wrapper around gtk_tree_view_row_activated().
func (v *treeView) RowActivated(path gtk.TreePath, column gtk.TreeViewColumn) {
	C.gtk_tree_view_row_activated(v.native(), castToTreePath(path).native(), castToTreeViewColumn(column).native())
}

// ExpandAll() is a wrapper around gtk_tree_view_expand_all().
func (v *treeView) ExpandAll() {
	C.gtk_tree_view_expand_all(v.native())
}

// CollapseAll() is a wrapper around gtk_tree_view_collapse_all().
func (v *treeView) CollapseAll() {
	C.gtk_tree_view_collapse_all(v.native())
}

// ExpandToPath() is a wrapper around gtk_tree_view_expand_to_path().
func (v *treeView) ExpandToPath(path gtk.TreePath) {
	C.gtk_tree_view_expand_to_path(v.native(), castToTreePath(path).native())
}

// ExpandRow() is a wrapper around gtk_tree_view_expand_row().
func (v *treeView) ExpandRow(path gtk.TreePath, openAll bool) bool {
	return gobool(C.gtk_tree_view_expand_row(v.native(), castToTreePath(path).native(), gbool(openAll)))
}

// CollapseRow() is a wrapper around gtk_tree_view_collapse_row().
func (v *treeView) CollapseRow(path gtk.TreePath) bool {
	return gobool(C.gtk_tree_view_collapse_row(v.native(), castToTreePath(path).native()))
}

// RowExpanded() is a wrapper around gtk_tree_view_row_expanded().
func (v *treeView) RowExpanded(path gtk.TreePath) bool {
	return gobool(C.gtk_tree_view_row_expanded(v.native(), castToTreePath(path).native()))
}

// SetReorderable is a wrapper around gtk_tree_view_set_reorderable().
func (v *treeView) SetReorderable(b bool) {
	C.gtk_tree_view_set_reorderable(v.native(), gbool(b))
}

// GetReorderable() is a wrapper around gtk_tree_view_get_reorderable().
func (v *treeView) GetReorderable() bool {
	return gobool(C.gtk_tree_view_get_reorderable(v.native()))
}

// GetBinWindow() is a wrapper around gtk_tree_view_get_bin_window().
func (v *treeView) GetBinWindow() gdk.Window {
	c := C.gtk_tree_view_get_bin_window(v.native())
	if c == nil {
		return nil
	}

	w := &gdk_impl.Window{wrapObject(unsafe.Pointer(c))}
	return w
}

// SetEnableSearch is a wrapper around gtk_tree_view_set_enable_search().
func (v *treeView) SetEnableSearch(b bool) {
	C.gtk_tree_view_set_enable_search(v.native(), gbool(b))
}

// GetEnableSearch() is a wrapper around gtk_tree_view_get_enable_search().
func (v *treeView) GetEnableSearch() bool {
	return gobool(C.gtk_tree_view_get_enable_search(v.native()))
}

// SetSearchColumn is a wrapper around gtk_tree_view_set_search_column().
func (v *treeView) SetSearchColumn(c int) {
	C.gtk_tree_view_set_search_column(v.native(), C.gint(c))
}

// GetSearchColumn() is a wrapper around gtk_tree_view_get_search_column().
func (v *treeView) GetSearchColumn() int {
	return int(C.gtk_tree_view_get_search_column(v.native()))
}

// GetSearchEntry() is a wrapper around gtk_tree_view_get_search_entry().
func (v *treeView) GetSearchEntry() gtk.Entry {
	c := C.gtk_tree_view_get_search_entry(v.native())
	if c == nil {
		return nil
	}
	return wrapEntry(wrapObject(unsafe.Pointer(c)))
}

// SetSearchEntry() is a wrapper around gtk_tree_view_set_search_entry().
func (v *treeView) SetSearchEntry(e gtk.Entry) {
	C.gtk_tree_view_set_search_entry(v.native(), castToEntry(e).native())
}

// SetFixedHeightMode is a wrapper around gtk_tree_view_set_fixed_height_mode().
func (v *treeView) SetFixedHeightMode(b bool) {
	C.gtk_tree_view_set_fixed_height_mode(v.native(), gbool(b))
}

// GetFixedHeightMode() is a wrapper around gtk_tree_view_get_fixed_height_mode().
func (v *treeView) GetFixedHeightMode() bool {
	return gobool(C.gtk_tree_view_get_fixed_height_mode(v.native()))
}

// SetHoverSelection is a wrapper around gtk_tree_view_set_hover_selection().
func (v *treeView) SetHoverSelection(b bool) {
	C.gtk_tree_view_set_hover_selection(v.native(), gbool(b))
}

// GetHoverSelection() is a wrapper around gtk_tree_view_get_hover_selection().
func (v *treeView) GetHoverSelection() bool {
	return gobool(C.gtk_tree_view_get_hover_selection(v.native()))
}

// SetHoverExpand is a wrapper around gtk_tree_view_set_hover_expand().
func (v *treeView) SetHoverExpand(b bool) {
	C.gtk_tree_view_set_hover_expand(v.native(), gbool(b))
}

// GetHoverExpand() is a wrapper around gtk_tree_view_get_hover_expand().
func (v *treeView) GetHoverExpand() bool {
	return gobool(C.gtk_tree_view_get_hover_expand(v.native()))
}

// SetRubberBanding is a wrapper around gtk_tree_view_set_rubber_banding().
func (v *treeView) SetRubberBanding(b bool) {
	C.gtk_tree_view_set_rubber_banding(v.native(), gbool(b))
}

// GetRubberBanding() is a wrapper around gtk_tree_view_get_rubber_banding().
func (v *treeView) GetRubberBanding() bool {
	return gobool(C.gtk_tree_view_get_rubber_banding(v.native()))
}

// IsRubberBandingActive() is a wrapper around gtk_tree_view_is_rubber_banding_active().
func (v *treeView) IsRubberBandingActive() bool {
	return gobool(C.gtk_tree_view_is_rubber_banding_active(v.native()))
}

// SetEnableTreeLines is a wrapper around gtk_tree_view_set_enable_tree_lines().
func (v *treeView) SetEnableTreeLines(b bool) {
	C.gtk_tree_view_set_enable_tree_lines(v.native(), gbool(b))
}

// GetEnableTreeLines() is a wrapper around gtk_tree_view_get_enable_tree_lines().
func (v *treeView) GetEnableTreeLines() bool {
	return gobool(C.gtk_tree_view_get_enable_tree_lines(v.native()))
}

// GetTooltipColumn() is a wrapper around gtk_tree_view_get_tooltip_column().
func (v *treeView) GetTooltipColumn() int {
	return int(C.gtk_tree_view_get_tooltip_column(v.native()))
}

// SetTooltipColumn() is a wrapper around gtk_tree_view_set_tooltip_column().
func (v *treeView) SetTooltipColumn(c int) {
	C.gtk_tree_view_set_tooltip_column(v.native(), C.gint(c))
}

// void 	gtk_tree_view_set_tooltip_row ()
// void 	gtk_tree_view_set_tooltip_cell ()
// gboolean 	gtk_tree_view_get_tooltip_context ()
// void 	gtk_tree_view_set_grid_lines ()
// GtkTreeViewGridLines 	gtk_tree_view_get_grid_lines ()
// void 	(*GtkTreeDestroyCountFunc) ()
// void 	gtk_tree_view_set_destroy_count_func ()
// gboolean 	(*GtkTreeViewRowSeparatorFunc) ()
// GtkTreeViewRowSeparatorFunc 	gtk_tree_view_get_row_separator_func ()
// void 	gtk_tree_view_set_row_separator_func ()
// void 	(*GtkTreeViewSearchPositionFunc) ()
// GtkTreeViewSearchPositionFunc 	gtk_tree_view_get_search_position_func ()
// void 	gtk_tree_view_set_search_position_func ()
// void 	gtk_tree_view_set_search_equal_func ()
// GtkTreeViewSearchEqualFunc 	gtk_tree_view_get_search_equal_func ()
// void 	gtk_tree_view_map_expanded_rows ()
// GList * 	gtk_tree_view_get_columns ()
// gint 	gtk_tree_view_insert_column_with_attributes ()
// gint 	gtk_tree_view_insert_column_with_data_func ()
// void 	gtk_tree_view_set_column_drag_function ()
// void 	gtk_tree_view_scroll_to_cell ()
// gboolean 	gtk_tree_view_is_blank_at_pos ()
// void 	gtk_tree_view_get_cell_area ()
// void 	gtk_tree_view_get_background_area ()
// void 	gtk_tree_view_get_visible_rect ()
// gboolean 	gtk_tree_view_get_visible_range ()
// void 	gtk_tree_view_convert_bin_window_to_tree_coords ()
// void 	gtk_tree_view_convert_bin_window_to_widget_coords ()
// void 	gtk_tree_view_convert_tree_to_bin_window_coords ()
// void 	gtk_tree_view_convert_tree_to_widget_coords ()
// void 	gtk_tree_view_convert_widget_to_bin_window_coords ()
// void 	gtk_tree_view_convert_widget_to_tree_coords ()
// void 	gtk_tree_view_enable_model_drag_dest ()
// void 	gtk_tree_view_enable_model_drag_source ()
// void 	gtk_tree_view_unset_rows_drag_source ()
// void 	gtk_tree_view_unset_rows_drag_dest ()
// void 	gtk_tree_view_set_drag_dest_row ()
// void 	gtk_tree_view_get_drag_dest_row ()
// gboolean 	gtk_tree_view_get_dest_row_at_pos ()
// cairo_surface_t * 	gtk_tree_view_create_row_drag_icon ()
