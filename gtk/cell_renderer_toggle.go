package gtk

type CellRendererToggle interface {
	CellRenderer

	GetActivatable() bool
	GetActive() bool
	GetRadio() bool
	SetActivatable(bool)
	SetActive(bool)
	SetRadio(bool)
} // end of CellRendererToggle

func AssertCellRendererToggle(_ CellRendererToggle) {}
