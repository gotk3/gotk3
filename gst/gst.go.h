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

static GstBin *
toGstBin(void *p)
{
	return (GST_BIN(p));
}

static GstPluginFeature *
toGstPluginFeature(void *p)
{
	return (GST_PLUGIN_FEATURE(p));
}

static GstElementFactory *
toGstElementFactory(void *p)
{
	return (GST_ELEMENT_FACTORY(p));
}

static GstPad *
toGstPad(void *p)
{
	return (GST_PAD(p));
}

static GstGhostPad *
toGstGhostPad(void *p)
{
	return (GST_GHOST_PAD(p));
}

static gint
toGstMessageType(void *p) {
	return (GST_MESSAGE_TYPE(p));
}

// Other

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

static const char*
messageTypeName(void *p)
{
	return (GST_MESSAGE_TYPE_NAME(p));
}

// Definitions
typedef enum
{
	___EMPTY = 0
} _CONSTS;


