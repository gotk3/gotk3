package iface


type EventMotion interface {
    Event

    MotionVal() (float64, float64)
    MotionValRoot() (float64, float64)
} // end of EventMotion

func AssertEventMotion(_ EventMotion) {}
