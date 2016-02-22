package gtk

type ComboBoxText interface {
	ComboBox

	Append(string, string)
	AppendText(string)
	GetActiveText() string
	Insert(int, string, string)
	InsertText(int, string)
	Prepend(string, string)
	PrependText(string)
	Remove2(int)
	RemoveAll()
} // end of ComboBoxText

func AssertComboBoxText(_ ComboBoxText) {}
