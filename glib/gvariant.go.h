
#ifndef __GLIB_GO_H__
#define __GLIB_GO_H__

#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>

#include <glib/gi18n.h>
#include <locale.h>

static GVariant *
toGVariant(void *p)
{
	return (GVariant(p));
}

#endif
