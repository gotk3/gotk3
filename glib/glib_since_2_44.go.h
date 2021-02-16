// Same copyright and license as the rest of the files in this project

#include <stdlib.h>

#include <glib-object.h>
#include <glib.h>

static GListModel *toGListModel(void *p) { return (G_LIST_MODEL(p)); }

static GListStore *toGListStore(void *p) { return (G_LIST_STORE(p)); }

static inline void _g_list_store_insert_sorted(GListStore *model, gpointer item,
                                               gpointer user_data) {
  g_list_store_insert_sorted(model, item,
                             (GCompareDataFunc)(goCompareDataFuncs), user_data);
}
