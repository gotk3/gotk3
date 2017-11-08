// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

#pragma once

#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static GtkShortcutsWindow *
toGtkShortcutsWindow(void *p)
{
	return (GTK_SHORTCUTS_WINDOW(p));
}

static GtkShortcutsSection *
toGtkShortcutsSection(void *p)
{
	return (GTK_SHORTCUTS_SECTION(p));
}

static GtkShortcutsGroup *
toGtkShortcutsGroup(void *p)
{
	return (GTK_SHORTCUTS_GROUP(p));
}

static GtkShortcutsShortcut *
toGtkShortcutsShortcut(void *p)
{
	return (GTK_SHORTCUTS_SHORTCUT(p));
}

