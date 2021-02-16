// +build !gtk_3_6,!gtk_3_8

package gtk

// #include <gtk/gtk.h>
import "C"

// GetPreferredHeightAndBaselineForWidth is a wrapper around gtk_widget_get_preferred_height_and_baseline_for_width().
func (v *Widget) GetPreferredHeightAndBaselineForWidth(height int) (int, int, int, int) {

	var minimum, natural, minimum_baseline, natural_baseline C.gint

	C.gtk_widget_get_preferred_height_and_baseline_for_width(
		v.native(),
		C.gint(height),
		&minimum,
		&natural,
		&minimum_baseline,
		&natural_baseline)
	return int(minimum),
		int(natural),
		int(minimum_baseline),
		int(natural_baseline)
}

// TODO:
// gtk_widget_get_valign_with_baseline().
// gtk_widget_init_template().
// gtk_widget_class_set_template().
// gtk_widget_class_set_template_from_resource().
// gtk_widget_get_template_child().
// gtk_widget_class_bind_template_child().
// gtk_widget_class_bind_template_child_internal().
// gtk_widget_class_bind_template_child_private().
// gtk_widget_class_bind_template_child_internal_private().
// gtk_widget_class_bind_template_child_full().
// gtk_widget_class_bind_template_callback().
// gtk_widget_class_bind_template_callback_full().
// gtk_widget_class_set_connect_func().

// GetScaleFactor is a wrapper around gtk_widget_get_scale_factor().
func (v *Widget) GetScaleFactor() int {
	return int(C.gtk_widget_get_scale_factor(v.native()))
}
