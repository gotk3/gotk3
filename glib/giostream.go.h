#pragma once

#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <glib.h>
#include <gio/gio.h>

static GIOStream *
toGIOStream(void *p)
{
	return (G_IO_STREAM(p));
}

static GInputStream *
toGInputStream(void *p)
{
	return (G_INPUT_STREAM(p));
}

static GOutputStream *
toGOutputStream(void *p)
{
	return (G_OUTPUT_STREAM(p));
}
