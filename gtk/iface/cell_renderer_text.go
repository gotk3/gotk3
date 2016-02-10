package iface


type CellRendererText interface {
    CellRenderer
} // end of CellRendererText

func AssertCellRendererText(_ CellRendererText) {}
