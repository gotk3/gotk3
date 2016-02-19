package gtk

type Grid interface {
	Container

	Attach(Widget, int, int, int, int)
	AttachNextTo(Widget, Widget, PositionType, int, int)
	GetChildAt(int, int) (Widget, error)
	GetColumnHomogeneous() bool
	GetColumnSpacing() uint
	GetRowHomogeneous() bool
	GetRowSpacing() uint
	InsertColumn(int)
	InsertNextTo(Widget, PositionType)
	InsertRow(int)
	SetColumnHomogeneous(bool)
	SetColumnSpacing(uint)
	SetRowHomogeneous(bool)
	SetRowSpacing(uint)
} // end of Grid

func AssertGrid(_ Grid) {}
