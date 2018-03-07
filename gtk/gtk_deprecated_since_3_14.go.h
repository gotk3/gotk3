
/*
 * deprecated since version 3.14
 */

// Wrapper for gtk_menu_popup to allow calling gtk_status_icon_position_menu as callback from go code
// Used in func (v *Menu) PopupAtStatusIcon
static void
gotk_menu_popup_at_status_icon(GtkMenu *menu, GtkStatusIcon *status_icon, guint button, guint32 activate_time)
{
	gtk_menu_popup(menu, NULL, NULL, gtk_status_icon_position_menu, status_icon, button, activate_time);
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

static GtkMisc *
toGtkMisc(void *p)
{
	return (GTK_MISC(p));
}

static GtkStatusIcon *
toGtkStatusIcon(void *p)
{
	return (GTK_STATUS_ICON(p));
}

static GdkPixbuf *
toGdkPixbuf(void *p)
{
	return (GDK_PIXBUF(p));
}

/*
 * End deprecated since version 3.14
 */
