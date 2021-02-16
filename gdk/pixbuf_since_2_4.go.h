// Same copyright and license as the rest of the files in this project

#include <stdlib.h>

extern gboolean goPixbufSaveCallback(gchar *buf, gsize count, GError **error,
                                     gpointer data);

static inline gboolean _gdk_pixbuf_save_png_writer(GdkPixbuf *pixbuf,
                                                   gpointer callback_id,
                                                   GError **err,
                                                   const char *compression) {
  return gdk_pixbuf_save_to_callback(
      pixbuf, (GdkPixbufSaveFunc)(goPixbufSaveCallback), callback_id, "png",
      err, "compression", compression, NULL);
}

static inline gboolean _gdk_pixbuf_save_jpeg_writer(GdkPixbuf *pixbuf,
                                                    gpointer callback_id,
                                                    GError **err,
                                                    const char *quality) {
  return gdk_pixbuf_save_to_callback(
      pixbuf, (GdkPixbufSaveFunc)(goPixbufSaveCallback), callback_id, "jpeg",
      err, "quality", quality, NULL);
}

static inline void _pixbuf_error_set_callback_not_found(GError **err) {
  GQuark domain = g_quark_from_static_string("go error");
  g_set_error_literal(err, domain, 1, "pixbuf callback not found");
}

static inline void _pixbuf_error_set(GError **err, char *message) {
  GQuark domain = g_quark_from_static_string("go error");
  g_set_error_literal(err, domain, 1, message);
}
