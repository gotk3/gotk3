package iface


type MessageDialog interface {
    Dialog

    FormatSecondaryMarkup(string, ...interface{})
    FormatSecondaryText(string, ...interface{})
    SetMarkup(string)
} // end of MessageDialog

func AssertMessageDialog(_ MessageDialog) {}
