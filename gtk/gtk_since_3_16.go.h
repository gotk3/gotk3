/*
 * Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
 *
 * This file originated from: http://opensource.conformal.com/
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

#pragma once

#include <stdlib.h>

static GListModel *toGListModel(void *p) { return (G_LIST_MODEL(p)); }

static GtkModelButton *toGtkModelButton(void *mb) {
  return (GTK_MODEL_BUTTON(mb));
}

static GtkPopoverMenu *toGtkPopoverMenu(void *p) {
  return (GTK_POPOVER_MENU(p));
}

static GtkStackSidebar *toGtkStackSidebar(void *p) {
  return (GTK_STACK_SIDEBAR(p));
}

static GtkGLArea *toGtkGLArea(void *p) { return (GTK_GL_AREA(p)); }

extern void goListBoxCreateWidgetFuncs(gpointer item, gpointer user_data);

static inline void _gtk_list_box_bind_model(GtkListBox *box, GListModel *model,
                                            gpointer user_data) {
  gtk_list_box_bind_model(
      box, model, (GtkListBoxCreateWidgetFunc)(goListBoxCreateWidgetFuncs),
      user_data, NULL);
}
