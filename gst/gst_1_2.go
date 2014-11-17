// +build !gst_1_0

package gst

// #include <gst/gst.h>
// #include "gst.go.h"
import "C"

// Since 1.2
const (
	MESSAGE_NEED_CONTEXT MessageType = C.GST_MESSAGE_NEED_CONTEXT
	MESSAGE_HAVE_CONTEXT MessageType = C.GST_MESSAGE_HAVE_CONTEXT
)
