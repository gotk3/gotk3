package gtk

import "github.com/gotk3/gotk3/gdk"

type TreeView interface {
	Container

	AppendColumn(TreeViewColumn) int
	CollapseAll()
	CollapseRow(TreePath) bool
	ColumnsAutosize()
	ExpandAll()
	ExpandRow(TreePath, bool) bool
	ExpandToPath(TreePath)
	GetActivateOnSingleClick() bool
	GetBinWindow() gdk.Window
	GetColumn(int) TreeViewColumn
	GetCursor() (TreePath, TreeViewColumn)
	GetEnableSearch() bool
	GetEnableTreeLines() bool
	GetExpanderColumn() TreeViewColumn
	GetFixedHeightMode() bool
	GetHeadersClickable() bool
	GetHeadersVisible() bool
	GetHoverExpand() bool
	GetHoverSelection() bool
	GetLevelIndentation() int
	GetModel() (TreeModel, error)
	GetNColumns() uint
	GetPathAtPos(int, int, TreePath, TreeViewColumn, *int, *int) bool
	GetReorderable() bool
	GetRubberBanding() bool
	GetSearchColumn() int
	GetSearchEntry() Entry
	GetSelection() (TreeSelection, error)
	GetShowExpanders() bool
	GetTooltipColumn() int
	InsertColumn(TreeViewColumn, int) int
	IsRubberBandingActive() bool
	MoveColumnAfter(TreeViewColumn, TreeViewColumn)
	RemoveColumn(TreeViewColumn) int
	RowActivated(TreePath, TreeViewColumn)
	RowExpanded(TreePath) bool
	ScrollToPoint(int, int)
	SetActivateOnSingleClick(bool)
	SetCursor(TreePath, TreeViewColumn, bool)
	SetCursorOnCell(TreePath, TreeViewColumn, CellRenderer, bool)
	SetEnableSearch(bool)
	SetEnableTreeLines(bool)
	SetExpanderColumn(TreeViewColumn)
	SetFixedHeightMode(bool)
	SetHeadersClickable(bool)
	SetHeadersVisible(bool)
	SetHoverExpand(bool)
	SetHoverSelection(bool)
	SetLevelIndentation(int)
	SetModel(TreeModel)
	SetReorderable(bool)
	SetRubberBanding(bool)
	SetSearchColumn(int)
	SetSearchEntry(Entry)
	SetShowExpanders(bool)
	SetTooltipColumn(int)
} // end of TreeView

func AssertTreeView(_ TreeView) {}
