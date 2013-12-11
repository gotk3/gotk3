/*
 * Copyright (c) 2013 Conformal Systems <info@conformal.com>
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

#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static GtkGrid *
toGtkGrid(void *p)
{
	return (GTK_GRID(p));
}

static GtkWidget *
toGtkWidget(void *p)
{
	return (GTK_WIDGET(p));
}

static GtkContainer *
toGtkContainer(void *p)
{
	return (GTK_CONTAINER(p));
}

static GtkProgressBar *
toGtkProgressBar(void *p)
{
	return (GTK_PROGRESS_BAR(p));
}

static GtkBin *
toGtkBin(void *p)
{
	return (GTK_BIN(p));
}

static GtkWindow *
toGtkWindow(void *p)
{
	return (GTK_WINDOW(p));
}

static GtkBox *
toGtkBox(void *p)
{
	return (GTK_BOX(p));
}

static GtkStatusbar *
toGtkStatusbar(void *p)
{
	return (GTK_STATUSBAR(p));
}

static GtkMisc *
toGtkMisc(void *p)
{
	return (GTK_MISC(p));
}

static GtkLabel *
toGtkLabel(void *p)
{
	return (GTK_LABEL(p));
}

static GtkNotebook *
toGtkNotebook(void *p)
{
	return (GTK_NOTEBOOK(p));
}

static GtkEntry *
toGtkEntry(void *p)
{
	return (GTK_ENTRY(p));
}

static GtkEntryBuffer *
toGtkEntryBuffer(void *p)
{
	return (GTK_ENTRY_BUFFER(p));
}

static GtkEntryCompletion *
toGtkEntryCompletion(void *p)
{
	return (GTK_ENTRY_COMPLETION(p));
}

static GtkAdjustment *
toGtkAdjustment(void *p)
{
	return (GTK_ADJUSTMENT(p));
}

static GtkImage *
toGtkImage(void *p)
{
	return (GTK_IMAGE(p));
}

static GtkButton *
toGtkButton(void *p)
{
	return (GTK_BUTTON(p));
}

static GtkScrolledWindow *
toGtkScrolledWindow(void *p)
{
	return (GTK_SCROLLED_WINDOW(p));
}

static GtkMenuItem *
toGtkMenuItem(void *p)
{
	return (GTK_MENU_ITEM(p));
}

static GtkMenu *
toGtkMenu(void *p)
{
	return (GTK_MENU(p));
}

static GtkMenuShell *
toGtkMenuShell(void *p)
{
	return (GTK_MENU_SHELL(p));
}

static GtkMenuBar *
toGtkMenuBar(void *p)
{
	return (GTK_MENU_BAR(p));
}

static GtkSpinButton *
toGtkSpinButton(void *p)
{
	return (GTK_SPIN_BUTTON(p));
}

static GtkComboBox *
toGtkComboBox(void *p)
{
	return (GTK_COMBO_BOX(p));
}

static GtkListStore *
toGtkListStore(void *p)
{
	return (GTK_LIST_STORE(p));
}

static GtkTextView *
toGtkTextView(void *p)
{
	return (GTK_TEXT_VIEW(p));
}

static GtkTextTagTable *
toGtkTextTagTable(void *p)
{
	return (GTK_TEXT_TAG_TABLE(p));
}

static GtkTextBuffer *
toGtkTextBuffer(void *p)
{
	return (GTK_TEXT_BUFFER(p));
}

static GtkTreeModel *
toGtkTreeModel(void *p)
{
	return (GTK_TREE_MODEL(p));
}

static GtkCellRenderer *
toGtkCellRenderer(void *p)
{
	return (GTK_CELL_RENDERER(p));
}

static GtkCellRendererText *
toGtkCellRendererText(void *p)
{
	return (GTK_CELL_RENDERER_TEXT(p));
}

static GtkCellLayout *
toGtkCellLayout(void *p)
{
	return (GTK_CELL_LAYOUT(p));
}

static GtkOrientable *
toGtkOrientable(void *p)
{
	return (GTK_ORIENTABLE(p));
}

static GtkTreeView *
toGtkTreeView(void *p)
{
	return (GTK_TREE_VIEW(p));
}

static GtkTreeViewColumn *
toGtkTreeViewColumn(void *p)
{
	return (GTK_TREE_VIEW_COLUMN(p));
}

static GtkTreeSelection *
toGtkTreeSelection(void *p)
{
	return (GTK_TREE_SELECTION(p));
}

static GtkClipboard *
toGtkClipboard(void *p)
{
	return (GTK_CLIPBOARD(p));
}

static GtkDialog *
toGtkDialog(void *p)
{
	return (GTK_DIALOG(p));
}

static GtkMessageDialog *
toGtkMessageDialog(void *p)
{
	return (GTK_MESSAGE_DIALOG(p));
}

static GtkBuilder *
toGtkBuilder(void *p)
{
	return (GTK_BUILDER(p));
}

static GtkSeparatorMenuItem *
toGtkSeparatorMenuItem(void *p)
{
	return (GTK_SEPARATOR_MENU_ITEM(p));
}

static GtkCheckButton *
toGtkCheckButton(void *p)
{
	return (GTK_CHECK_BUTTON(p));
}

static GtkToggleButton *
toGtkToggleButton(void *p)
{
	return (GTK_TOGGLE_BUTTON(p));
}

static GtkFrame *
toGtkFrame(void *p)
{
	return (GTK_FRAME(p));
}

static GtkSeparator *
toGtkSeparator(void *p)
{
	return (GTK_SEPARATOR(p));
}

static GtkScrollbar *
toGtkScrollbar(void *p)
{
	return (GTK_SCROLLBAR(p));
}

static GtkRange *
toGtkRange(void *p)
{
	return (GTK_RANGE(p));
}

static GtkSearchEntry *
toGtkSearchEntry(void *p)
{
	return (GTK_SEARCH_ENTRY(p));
}

static GType * 
alloc_types(int n) {
	return ((GType *)g_new0(GType, n));
}

static void
set_type(GType *types, int n, GType t)
{
	types[n] = t;
}

static GtkTreeViewColumn *
_gtk_tree_view_column_new_with_attributes_one(const gchar *title,
    GtkCellRenderer *renderer, const gchar *attribute, gint column)
{
	GtkTreeViewColumn	*tvc;

	tvc = gtk_tree_view_column_new_with_attributes(title, renderer,
	    attribute, column, NULL);
	return (tvc);
}

static GtkWidget *
_gtk_message_dialog_new(GtkWindow *parent, GtkDialogFlags flags,
    GtkMessageType type, GtkButtonsType buttons, char *msg)
{
	GtkWidget		*w;

	w = gtk_message_dialog_new(parent, flags, type, buttons, "%s", msg);
	return (w);
}

static gchar *
error_get_message(GError *error)
{
	return error->message;
}

static const gchar *
object_get_class_name(GObject *object)
{
	return G_OBJECT_CLASS_NAME(G_OBJECT_GET_CLASS(object));
}
