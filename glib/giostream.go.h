#pragma once

#include <gio/gio.h>
#include <glib.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static GIOStream *toGIOStream(void *p) { return (G_IO_STREAM(p)); }

static GInputStream *toGInputStream(void *p) { return (G_INPUT_STREAM(p)); }

static GOutputStream *toGOutputStream(void *p) { return (G_OUTPUT_STREAM(p)); }
