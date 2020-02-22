//+build gtk_3_6 gtk_3_8 gtk_3_10 gtk_3_12 gtk_3_14 gtk_3_16 gtk_3_18 gtk_deprecated

package gtk

// #include <gtk/gtk.h>
import "C"
import (
	"github.com/gotk3/gotk3/gdk"
)

/*
GtkWidget *
gtk_test_create_simple_window (const gchar *window_title,
							   const gchar *dialog_text);

GtkWidget *
gtk_test_create_widget (GType widget_type,
                        const gchar *first_property_name,
						...);

GtkWidget *
gtk_test_display_button_window (const gchar *window_title,
                                const gchar *dialog_text,
								...);

double
gtk_test_slider_get_value (GtkWidget *widget);

void
gtk_test_slider_set_perc (GtkWidget *widget,
						  double percentage);

gboolean
gtk_test_spin_button_click (GtkSpinButton *spinner,
                            guint button,
							gboolean upwards);

gchar *
gtk_test_text_get (GtkWidget *widget);

void
gtk_test_text_set (GtkWidget *widget,
				   const gchar *string);
*/

// TestWidgetClick is a wrapper around gtk_test_widget_click()
// Deprecated since 3.20
//
// This function will generate a button click (button press and button release event)
// in the middle of the first GdkWindow found that belongs to widget.
// For windowless widgets like GtkButton (which returns FALSE from gtk_widget_get_has_window()),
// this will often be an input-only event window.
// For other widgets, this is usually widget->window.
//
// widget: Widget to generate a button click on.
// button: Number of the pointer button for the event, usually 1, 2 or 3.
// modifiers: Keyboard modifiers the event is setup with.
//
// returns: whether all actions neccessary for the button click simulation were carried out successfully.
func TestWidgetClick(widget IWidget, button gdk.Button, modifiers gdk.ModifierType) bool {
	return gobool(C.gtk_test_widget_click(widget.toWidget(), C.guint(button), C.GdkModifierType(modifiers)))
}
