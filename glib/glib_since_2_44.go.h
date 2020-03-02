// Same copyright and license as the rest of the files in this project

#include <stdlib.h>

#include <glib.h>
#include <glib-object.h>

static GListModel *
toGListModel(void *p)
{
	return (G_LIST_MODEL(p));
}

static GListStore *
toGListStore(void *p)
{
	return (G_LIST_STORE(p));
}
