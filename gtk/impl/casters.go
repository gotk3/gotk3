package impl

import "github.com/gotk3/gotk3/gtk"

func castToApplication(s gtk.Application) *Application {
	if s == nil {
		return nil
	}
	return s.(*Application)
}

func castToAdjustment(s gtk.Adjustment) *Adjustment {
	if s == nil {
		return nil
	}
	return s.(*Adjustment)
}

func castToEntryBuffer(s gtk.EntryBuffer) *EntryBuffer {
	if s == nil {
		return nil
	}
	return s.(*EntryBuffer)
}

func castToEntryCompletion(s gtk.EntryCompletion) *EntryCompletion {
	if s == nil {
		return nil
	}
	return s.(*EntryCompletion)
}

func castToFileFilter(s gtk.FileFilter) *FileFilter {
	if s == nil {
		return nil
	}
	return s.(*FileFilter)
}

func castToRadioButton(s gtk.RadioButton) *RadioButton {
	if s == nil {
		return nil
	}
	return s.(*RadioButton)
}

func castToRadioMenuItem(s gtk.RadioMenuItem) *RadioMenuItem {
	if s == nil {
		return nil
	}
	return s.(*RadioMenuItem)
}

func castToRecentFilter(s gtk.RecentFilter) *RecentFilter {
	if s == nil {
		return nil
	}
	return s.(*RecentFilter)
}

func castToTextBuffer(s gtk.TextBuffer) *TextBuffer {
	if s == nil {
		return nil
	}
	return s.(*TextBuffer)
}

func castToTextIter(s gtk.TextIter) *TextIter {
	if s == nil {
		return nil
	}
	return s.(*TextIter)
}

func castToTextTag(s gtk.TextTag) *TextTag {
	if s == nil {
		return nil
	}
	return s.(*TextTag)
}

func castToTextTagTable(s gtk.TextTagTable) *TextTagTable {
	if s == nil {
		return nil
	}
	return s.(*TextTagTable)
}

func castToTreeIter(s gtk.TreeIter) *TreeIter {
	if s == nil {
		return nil
	}
	return s.(*TreeIter)
}

func castToTreePath(s gtk.TreePath) *TreePath {
	if s == nil {
		return nil
	}
	return s.(*TreePath)
}

func castToAccelGroup(s gtk.AccelGroup) *AccelGroup {
	if s == nil {
		return nil
	}
	return s.(*AccelGroup)
}

func castToStyleContext(s gtk.StyleContext) *StyleContext {
	if s == nil {
		return nil
	}
	return s.(*StyleContext)
}

func castToTreeViewColumn(s gtk.TreeViewColumn) *TreeViewColumn {
	if s == nil {
		return nil
	}
	return s.(*TreeViewColumn)
}

func castToCellRenderer(s gtk.CellRenderer) *CellRenderer {
	if s == nil {
		return nil
	}
	return s.(*CellRenderer)
}

func castToEntry(s gtk.Entry) *Entry {
	if s == nil {
		return nil
	}
	return s.(*Entry)
}

func castToAllocation(s gtk.Allocation) *Allocation {
	if s == nil {
		return nil
	}
	return s.(*Allocation)
}
