package impl

import "github.com/gotk3/gotk3/gtk"

func castToApplication(s gtk.Application) *application {
	if s == nil {
		return nil
	}
	return s.(*application)
}

func castToAdjustment(s gtk.Adjustment) *adjustment {
	if s == nil {
		return nil
	}
	return s.(*adjustment)
}

func castToEntryBuffer(s gtk.EntryBuffer) *entryBuffer {
	if s == nil {
		return nil
	}
	return s.(*entryBuffer)
}

func castToEntryCompletion(s gtk.EntryCompletion) *entryCompletion {
	if s == nil {
		return nil
	}
	return s.(*entryCompletion)
}

func castToFileFilter(s gtk.FileFilter) *fileFilter {
	if s == nil {
		return nil
	}
	return s.(*fileFilter)
}

func castToRadioButton(s gtk.RadioButton) *radioButton {
	if s == nil {
		return nil
	}
	return s.(*radioButton)
}

func castToRadioMenuItem(s gtk.RadioMenuItem) *radioMenuItem {
	if s == nil {
		return nil
	}
	return s.(*radioMenuItem)
}

func castToRecentFilter(s gtk.RecentFilter) *recentFilter {
	if s == nil {
		return nil
	}
	return s.(*recentFilter)
}

func castToTextBuffer(s gtk.TextBuffer) *textBuffer {
	if s == nil {
		return nil
	}
	return s.(*textBuffer)
}

func castToTextIter(s gtk.TextIter) *textIter {
	if s == nil {
		return nil
	}
	return s.(*textIter)
}

func castToTextTag(s gtk.TextTag) *textTag {
	if s == nil {
		return nil
	}
	return s.(*textTag)
}

func castToTextTagTable(s gtk.TextTagTable) *textTagTable {
	if s == nil {
		return nil
	}
	return s.(*textTagTable)
}

func castToTreeIter(s gtk.TreeIter) *treeIter {
	if s == nil {
		return nil
	}
	return s.(*treeIter)
}

func castToTreePath(s gtk.TreePath) *treePath {
	if s == nil {
		return nil
	}
	return s.(*treePath)
}

func castToAccelGroup(s gtk.AccelGroup) *accelGroup {
	if s == nil {
		return nil
	}
	return s.(*accelGroup)
}

func castToStyleContext(s gtk.StyleContext) *styleContext {
	if s == nil {
		return nil
	}
	return s.(*styleContext)
}

func castToTreeViewColumn(s gtk.TreeViewColumn) *treeViewColumn {
	if s == nil {
		return nil
	}
	return s.(*treeViewColumn)
}

func castToCellRenderer(s gtk.CellRenderer) *cellRenderer {
	if s == nil {
		return nil
	}
	return s.(*cellRenderer)
}

func castToEntry(s gtk.Entry) *entry {
	if s == nil {
		return nil
	}
	return s.(*entry)
}

func castToAllocation(s gtk.Allocation) *allocation {
	if s == nil {
		return nil
	}
	return s.(*allocation)
}
