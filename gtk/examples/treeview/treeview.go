package main

import (
	"github.com/andre-hub/gotk3/glib"
	"github.com/andre-hub/gotk3/gtk"
	"log"
)

// IDs to access the tree view columns by
const (
	COLUMN_VERSION = iota
	COLUMN_FEATURE
)

// Add a column to the tree view (during the initialization of the tree view)
func createColumn(title string, id int) *gtk.TreeViewColumn {
	cellRenderer, err := gtk.CellRendererTextNew()
	if err != nil {
		log.Fatal("Unable to create text cell renderer:", err)
	}

	column, err := gtk.TreeViewColumnNewWithAttribute(title, cellRenderer, "text", id)
	if err != nil {
		log.Fatal("Unable to create cell column:", err)
	}

	return column
}

// Creates a tree view and the list store that holds its data
func setupTreeView() (*gtk.TreeView, *gtk.ListStore) {
	treeView, err := gtk.TreeViewNew()
	if err != nil {
		log.Fatal("Unable to create tree view:", err)
	}

	treeView.AppendColumn(createColumn("Version", COLUMN_VERSION))
	treeView.AppendColumn(createColumn("Feature", COLUMN_FEATURE))

	// Creating a list store. This is what holds the data that will be shown on our tree view.
	listStore, err := gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		log.Fatal("Unable to create list store:", err)
	}
	treeView.SetModel(listStore)

	return treeView, listStore
}

// Append a row to the list store for the tree view
func addRow(listStore *gtk.ListStore, version, feature string) {
	// Get an iterator for a new row at the end of the list store
	iter := listStore.Append()

	// Set the contents of the list store row that the iterator represents
	err := listStore.Set(iter,
		[]int{COLUMN_VERSION, COLUMN_FEATURE},
		[]interface{}{version, feature})

	if err != nil {
		log.Fatal("Unable to add row:", err)
	}
}

// Create and initialize the window
func setupWindow(title string) *gtk.Window {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.SetTitle(title)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetDefaultSize(600, 300)
	return win
}

func main() {
	gtk.Init(nil)

	win := setupWindow("Go Feature Timeline")

	treeView, listStore := setupTreeView()
	win.Add(treeView)

	// Add some rows to the list store
	addRow(listStore, "r57", "Gofix command added for rewriting code for new APIs")
	addRow(listStore, "r60", "URL parsing moved to new \"url\" package")
	addRow(listStore, "go1.0", "Rune type introduced as alias for int32")
	addRow(listStore, "go1.1", "Race detector added to tools")
	addRow(listStore, "go1.2", "Limit for number of threads added")
	addRow(listStore, "go1.3", "Support for various BSD's, Plan 9 and Solaris")

	win.ShowAll()
	gtk.Main()
}
