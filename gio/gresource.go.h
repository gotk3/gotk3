#pragma once

#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <glib.h>

static inline char** next_charptr(char** s) { return (s+1); }

static inline void char_g_strfreev(char** s) {
    g_strfreev((gchar**) s);
}
