package gtk

type TreePath interface {
	String() string
} // end of TreePath

func AssertTreePath(_ TreePath) {}
