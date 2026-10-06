#pragma once

#include <stdint.h>
#include <stdlib.h>
#include <string.h>

// Callback IDs are guintptr, not gpointer: see gtk.go.h.

extern void goPrintSettings(gchar *key, gchar *value, guintptr user_data);

static inline void _gotk3_goPrintSettings(const gchar *key, const gchar *value,
                                          gpointer user_data) {
  goPrintSettings((gchar *)key, (gchar *)value, (guintptr)user_data);
}

static inline void _gtk_print_settings_foreach(GtkPrintSettings *ps,
                                               guintptr user_data) {
  gtk_print_settings_foreach(ps, _gotk3_goPrintSettings, (gpointer)user_data);
}

extern void goPageSetupDone(GtkPageSetup *setup, guintptr data);

static inline void _gotk3_goPageSetupDone(GtkPageSetup *setup, gpointer data) {
  goPageSetupDone(setup, (guintptr)data);
}

static inline void
_gtk_print_run_page_setup_dialog_async(GtkWindow *parent, GtkPageSetup *setup,
                                       GtkPrintSettings *settings,
                                       guintptr data) {
  gtk_print_run_page_setup_dialog_async(parent, setup, settings,
                                        _gotk3_goPageSetupDone,
                                        (gpointer)data);
}
