// +build !pango_1_36,!pango_1_38,!pango_1_40,!pango_1_42

package pango

// #include <pango/pango.h>
// #include "pango.go.h"
import "C"

var (
	ATTR_INSERT_HYPHENS AttrType = C.PANGO_ATTR_INSERT_HYPHENS
)

func AttrInsertHyphensNew(insertHyphens bool) *Attribute {
	c := C.pango_attr_insert_hyphens_new(gbool(insertHyphens))
	attr := new(Attribute)
	attr.internal = c
	return attr
}
