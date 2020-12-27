#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

#include <gio/gio.h>

static GFileInputStream *toGFileInputStream(void *p) {
  return (G_FILE_INPUT_STREAM(p));
}

static GFileOutputStream *toGFileOutputStream(void *p) {
  return (G_FILE_OUTPUT_STREAM(p));
}
