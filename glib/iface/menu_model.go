package iface


type MenuModel interface {
    Object

    GetItemLink(int, string) MenuModel
    GetNItems() int
    IsMutable() bool
    ItemsChanged(int, int, int)
} // end of MenuModel

func AssertMenuModel(_ MenuModel) {}
