#include <stdlib.h>

static WebKitUserContentManager *
toWebKitUserContentManager(void *p)
{
	return (WEBKIT_USER_CONTENT_MANAGER(p));
}
