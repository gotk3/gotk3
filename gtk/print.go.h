#pragma once

#include <stdint.h>
#include <stdlib.h>
#include <string.h>

extern void goPrintSettings(gchar *key, gchar *value, gpointer user_data);

static inline void _gtk_print_settings_foreach(GtkPrintSettings *ps,
                                               gpointer user_data) {
  gtk_print_settings_foreach(ps, (GtkPrintSettingsFunc)(goPrintSettings),
                             user_data);
}

extern void goPageSetupDone(GtkPageSetup *setup, gpointer data);

static inline void
_gtk_print_run_page_setup_dialog_async(GtkWindow *parent, GtkPageSetup *setup,
                                       GtkPrintSettings *settings,
                                       gpointer data) {
  gtk_print_run_page_setup_dialog_async(
      parent, setup, settings, (GtkPageSetupDoneFunc)(goPageSetupDone), data);
}
