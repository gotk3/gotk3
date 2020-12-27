// Same copyright and license as the rest of the files in this project

#include <stdlib.h>

static GdkPixbuf *toGdkPixbuf(void *p) { return (GDK_PIXBUF(p)); }

static GdkPixbufAnimation *toGdkPixbufAnimation(void *p) {
  return (GDK_PIXBUF_ANIMATION(p));
}
static gboolean

_gdk_pixbuf_save_png(GdkPixbuf *pixbuf, const char *filename, GError **err,
                     const char *compression) {
  return gdk_pixbuf_save(pixbuf, filename, "png", err, "compression",
                         compression, NULL);
}

static gboolean _gdk_pixbuf_save_jpeg(GdkPixbuf *pixbuf, const char *filename,
                                      GError **err, const char *quality) {
  return gdk_pixbuf_save(pixbuf, filename, "jpeg", err, "quality", quality,
                         NULL);
}

static GdkPixbufLoader *toGdkPixbufLoader(void *p) {
  return (GDK_PIXBUF_LOADER(p));
}
