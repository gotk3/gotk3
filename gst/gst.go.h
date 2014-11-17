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


static GstBus *
toGstBus(void *p)
{
	return (GST_BUS(p));
}

static gint
toGstMessageType(void *p) {
	return (GST_MESSAGE_TYPE(p));
}

static guint64
messageTimestamp(void *p)
{
	return (GST_MESSAGE_TIMESTAMP(p));
}

static guint64
messageSeqnum(void *p)
{
	return (GST_MESSAGE_SEQNUM(p));
}

// Definitions
typedef enum
{
	EMPTY = 0
} _CONSTS;


