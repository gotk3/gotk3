package iface


type VolumeButton interface {
    ScaleButton
} // end of VolumeButton

func AssertVolumeButton(_ VolumeButton) {}
