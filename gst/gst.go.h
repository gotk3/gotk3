#include <stdlib.h>
// Type Casting
static GstObject *
toGstObject(void *p)
{
	return (GST_OBJECT(p));
}

static GstElement *
toGstElement(void *p)
{
	return (GST_ELEMENT(p));
}


// Definitions
typedef enum
{
	EMPTY = 0
} _CONSTS;
