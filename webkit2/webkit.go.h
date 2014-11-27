#include <stdlib.h>

static WebKitWebView *
toWebKitWebView(void *p)
{
	return (WEBKIT_WEB_VIEW(p));
}

static WebKitSettings *
toWebKitSettings(void *p)
{
	return (WEBKIT_SETTINGS(p));
}

static WebKitWebContext *
toWebKitWebContext(void *p)
{
	return (WEBKIT_WEB_CONTEXT(p));
}

static WebKitWindowProperties *
toWebKitWindowProperties(void *p)
{
	return (WEBKIT_WINDOW_PROPERTIES(p));
}

// Definitions and macros
typedef enum
{
	_WEBKIT_MAJOR_VERSION = WEBKIT_MAJOR_VERSION,
	_WEBKIT_MINOR_VERSION = WEBKIT_MINOR_VERSION,
	_WEBKIT_MICRO_VERSION = WEBKIT_MICRO_VERSION
} _CONSTS;

static gboolean
_WEBKIT_CHECK_VERSION(unsigned int major, unsigned int minor, unsigned int micro)
{
	return WEBKIT_CHECK_VERSION(major, minor, micro);
}
