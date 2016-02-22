package gtk

type IconView interface {
	Container

	GetModel() (TreeModel, error)
	ScrollToPath(TreePath, bool, float64, float64)
	SelectPath(TreePath)
	SetModel(TreeModel)
} // end of IconView

func AssertIconView(_ IconView) {}
