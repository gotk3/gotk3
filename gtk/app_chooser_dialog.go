package gtk

type AppChooserDialog interface {
	Dialog

	GetHeading() (string, error)
	GetWidget() AppChooserWidget
	SetHeading(string)
} // end of AppChooserDialog

func AssertAppChooserDialog(_ AppChooserDialog) {}
