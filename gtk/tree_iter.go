package gtk

type TreeIter interface {
	Copy() (TreeIter, error)
} // end of TreeIter

func AssertTreeIter(_ TreeIter) {}