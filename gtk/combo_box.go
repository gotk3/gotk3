package gtk

type ComboBox interface {
	Bin

	GetActive() int
	GetActiveID() string
	GetActiveIter() (TreeIter, error)
	GetModel() (TreeModel, error)
	SetActive(int)
	SetActiveID(string) bool
	SetActiveIter(TreeIter)
	SetModel(TreeModel)
} // end of ComboBox

func AssertComboBox(_ ComboBox) {}
