#pragma once

#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <glib.h>
#include <gio/gio.h>

static inline char** next_charptr(char** s) { return (s+1); }

static inline void char_g_strfreev(char** s) {
    g_strfreev((gchar**) s);
}

static GIcon *
toGIcon(void *p)
{
	return (G_ICON(p));
}

static GFileIcon *
toGFileIcon(void *p)
{
	return (G_FILE_ICON(p));
}

static GFile *
toGFile(void *p)
{
	return (G_FILE(p));
}

