// +build !pango_1_42

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
	attr.pangoAttribute = c
	return attr
}
