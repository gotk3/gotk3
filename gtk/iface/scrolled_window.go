package iface


type ScrolledWindow interface {
    Bin

    GetHAdjustment() Adjustment
    GetVAdjustment() Adjustment
    SetHAdjustment(Adjustment)
    SetPolicy(PolicyType, PolicyType)
    SetVAdjustment(Adjustment)
} // end of ScrolledWindow

func AssertScrolledWindow(_ ScrolledWindow) {}
