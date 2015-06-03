// +build !gst_1_0,!gst_1_2

package gst

// #include <gst/gst.h>
// #include "gst.go.h"
import "C"

// Since 1.4
const (
	MESSAGE_EXTENDED       MessageType = C.GST_MESSAGE_EXTENDED
	MESSAGE_DEVICE_ADDED   MessageType = C.GST_MESSAGE_DEVICE_ADDED
	MESSAGE_DEVICE_REMOVED MessageType = C.GST_MESSAGE_DEVICE_REMOVED
)
