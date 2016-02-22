package gtk

type AppChooserButton interface {
	ComboBox

	AppendSeparator()
	GetHeading() (string, error)
	GetShowDefaultItem() bool
	GetShowDialogItem() bool
	SetActiveCustomItem(string)
	SetHeading(string)
	SetShowDefaultItem(bool)
	SetShowDialogItem(bool)
} // end of AppChooserButton

func AssertAppChooserButton(_ AppChooserButton) {}
