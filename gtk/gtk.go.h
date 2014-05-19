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

#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static GtkAboutDialog *
toGtkAboutDialog(void *p)
{
	return (GTK_ABOUT_DIALOG(p));
}

static GtkAlignment *
toGtkAlignment(void *p)
{
	return (GTK_ALIGNMENT(p));
}

static GtkArrow *
toGtkArrow(void *p)
{
	return (GTK_ARROW(p));
}

static GtkAssistant *
toGtkAssistant(void *p)
{
	return (GTK_ASSISTANT(p));
}

static GtkCalendar *
toGtkCalendar(void *p)
{
	return (GTK_CALENDAR(p));
}

static GtkDrawingArea *
toGtkDrawingArea(void *p)
{
	return (GTK_DRAWING_AREA(p));
}

static GtkEventBox *
toGtkEventBox(void *p)
{
	return (GTK_EVENT_BOX(p));
}

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

static GtkStatusIcon *
toGtkStatusIcon(void *p)
{
	return (GTK_STATUS_ICON(p));
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

static GtkSpinner *
toGtkSpinner(void *p)
{
	return (GTK_SPINNER(p));
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

static GtkSwitch *
toGtkSwitch(void *p)
{
	return (GTK_SWITCH(p));
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

static GtkCellRendererToggle *
toGtkCellRendererToggle(void *p)
{
	return (GTK_CELL_RENDERER_TOGGLE(p));
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

static GtkOffscreenWindow *
toGtkOffscreenWindow(void *p)
{
	return (GTK_OFFSCREEN_WINDOW(p));
}

static GtkFileChooser *
toGtkFileChooser(void *p)
{
	return (GTK_FILE_CHOOSER(p));
}

static GtkFileChooserButton *
toGtkFileChooserButton(void *p)
{
	return (GTK_FILE_CHOOSER_BUTTON(p));
}

static GtkFileChooserWidget *
toGtkFileChooserWidget(void *p)
{
	return (GTK_FILE_CHOOSER_WIDGET(p));
}

static GtkMenuButton *
toGtkMenuButton(void *p)
{
	return (GTK_MENU_BUTTON(p));
}

static GtkRadioButton *
toGtkRadioButton(void *p)
{
	return (GTK_RADIO_BUTTON(p));
}

static GtkCheckMenuItem *
toGtkCheckMenuItem(void *p)
{
	return (GTK_CHECK_MENU_ITEM(p));
}

static GtkRadioMenuItem *
toGtkRadioMenuItem(void *p)
{
	return (GTK_RADIO_MENU_ITEM(p));
}

static GtkToolItem *
toGtkToolItem(void *p)
{
	return (GTK_TOOL_ITEM(p));
}

static GtkToolbar *
toGtkToolbar(void *p)
{
	return (GTK_TOOLBAR(p));
}

static GtkEditable *
toGtkEditable(void *p)
{
	return (GTK_EDITABLE(p));
}

static GtkToolButton *
toGtkToolButton(void *p)
{
	return (GTK_TOOL_BUTTON(p));
}

static GtkSeparatorToolItem *
toGtkSeparatorToolItem(void *p)
{
	return (GTK_SEPARATOR_TOOL_ITEM(p));
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

static GtkWidget *
_gtk_message_dialog_new_with_markup(GtkWindow *parent, GtkDialogFlags flags,
    GtkMessageType type, GtkButtonsType buttons, char *msg)
{
	GtkWidget		*w;

	w = gtk_message_dialog_new_with_markup(parent, flags, type, buttons,
	    "%s", msg);
	return (w);
}

void
_gtk_message_dialog_format_secondary_text(GtkMessageDialog *message_dialog,
    const gchar *msg)
{
	gtk_message_dialog_format_secondary_text(message_dialog, "%s", msg);
}

void
_gtk_message_dialog_format_secondary_markup(GtkMessageDialog *message_dialog,
    const gchar *msg)
{
	gtk_message_dialog_format_secondary_markup(message_dialog, "%s", msg);
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
