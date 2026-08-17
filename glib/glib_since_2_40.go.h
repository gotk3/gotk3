// Same copyright and license as the rest of the files in this project

#include <stdlib.h>

#include <glib-object.h>
#include <glib.h>

static GPropertyAction *toGPropertyAction(void *p) {
  return (G_PROPERTY_ACTION(p));
}

static GNotification *toGNotification(void *p) { return (G_NOTIFICATION(p)); }

/*
 * GVariantDict
 */
static GVariantDict *toGVariantDict(void *p) { return (GVariantDict *)p; }

