// Same copyright and license as the rest of the files in this project

#include <stdlib.h>

#include <glib-object.h>
#include <glib.h>

static inline void _g_list_store_sort(GListStore *model, gpointer user_data) {
  g_list_store_sort(model, (GCompareDataFunc)(goCompareDataFuncs), user_data);
}
