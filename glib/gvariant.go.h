//GVariant : GVariant — strongly typed value datatype
// https://developer.gnome.org/glib/2.26/glib-GVariant.html

#ifndef __GVARIANT_GO_H__
#define __GVARIANT_GO_H__

#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>

// Type Casting
static GVariant *
toGVariant(void *p)
{
	return (GVariant(p));
}
#endif
