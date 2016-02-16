package gtk

type MenuButton interface {
	ToggleButton

	GetAlignWidget() Widget
	GetDirection() ArrowType
	GetPopup() Menu
	SetAlignWidget(Widget)
	SetDirection(ArrowType)
	SetPopup(Menu)
} // end of MenuButton

func AssertMenuButton(_ MenuButton) {}
