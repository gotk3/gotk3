// Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
//
// This file originated from: http://opensource.conformal.com/
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// Go bindings for GTK+ 3.  Supports version 3.6 and later.
//
// Functions use the same names as the native C function calls, but use
// CamelCase.  In cases where native GTK uses pointers to values to
// simulate multiple return values, Go's native multiple return values
// are used instead.  Whenever a native GTK call could return an
// unexpected NULL pointer, an additonal error is returned in the Go
// binding.
//
// GTK's C API documentation can be very useful for understanding how the
// functions in this package work and what each type is for.  This
// documentation can be found at https://developer.gnome.org/gtk3/.
//
// In addition to Go versions of the C GTK functions, every struct type
// includes a method named Native (either by direct implementation, or
// by means of struct embedding).  These methods return a uintptr of the
// native C object the binding type represents.  These pointers may be
// type switched to a native C pointer using unsafe and used with cgo
// function calls outside this package.
//
// Memory management is handled in proper Go fashion, using runtime
// finalizers to properly free memory when it is no longer needed.  Each
// time a Go type is created with a pointer to a GObject, a reference is
// added for Go, sinking the floating reference when necessary.  After
// going out of scope and the next time Go's garbage collector is run, a
// finalizer is run to remove Go's reference to the GObject.  When this
// reference count hits zero (when neither Go nor GTK holds ownership)
// the object will be freed internally by GTK.
package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/pango"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.gtk_align_get_type()), marshalAlign},
		{glib.Type(C.gtk_accel_flags_get_type()), marshalAccelFlags},
		{glib.Type(C.gtk_accel_group_get_type()), marshalAccelGroup},
		{glib.Type(C.gtk_accel_map_get_type()), marshalAccelMap},
		{glib.Type(C.gtk_arrow_placement_get_type()), marshalArrowPlacement},
		{glib.Type(C.gtk_arrow_type_get_type()), marshalArrowType},
		{glib.Type(C.gtk_assistant_page_type_get_type()), marshalAssistantPageType},
		{glib.Type(C.gtk_buttons_type_get_type()), marshalButtonsType},
		{glib.Type(C.gtk_calendar_display_options_get_type()), marshalCalendarDisplayOptions},
		{glib.Type(C.gtk_dest_defaults_get_type()), marshalDestDefaults},
		{glib.Type(C.gtk_dialog_flags_get_type()), marshalDialogFlags},
		{glib.Type(C.gtk_entry_icon_position_get_type()), marshalEntryIconPosition},
		{glib.Type(C.gtk_file_chooser_action_get_type()), marshalFileChooserAction},
		{glib.Type(C.gtk_icon_lookup_flags_get_type()), marshalSortType},
		{glib.Type(C.gtk_icon_size_get_type()), marshalIconSize},
		{glib.Type(C.gtk_image_type_get_type()), marshalImageType},
		{glib.Type(C.gtk_input_hints_get_type()), marshalInputHints},
		{glib.Type(C.gtk_input_purpose_get_type()), marshalInputPurpose},
		{glib.Type(C.gtk_justification_get_type()), marshalJustification},
		{glib.Type(C.gtk_license_get_type()), marshalLicense},
		{glib.Type(C.gtk_message_type_get_type()), marshalMessageType},
		{glib.Type(C.gtk_orientation_get_type()), marshalOrientation},
		{glib.Type(C.gtk_pack_type_get_type()), marshalPackType},
		{glib.Type(C.gtk_path_type_get_type()), marshalPathType},
		{glib.Type(C.gtk_policy_type_get_type()), marshalPolicyType},
		{glib.Type(C.gtk_position_type_get_type()), marshalPositionType},
		{glib.Type(C.gtk_relief_style_get_type()), marshalReliefStyle},
		{glib.Type(C.gtk_response_type_get_type()), marshalResponseType},
		{glib.Type(C.gtk_selection_mode_get_type()), marshalSelectionMode},
		{glib.Type(C.gtk_shadow_type_get_type()), marshalShadowType},
		{glib.Type(C.gtk_sort_type_get_type()), marshalSortType},
		{glib.Type(C.gtk_state_flags_get_type()), marshalStateFlags},
		{glib.Type(C.gtk_target_flags_get_type()), marshalTargetFlags},
		{glib.Type(C.gtk_toolbar_style_get_type()), marshalToolbarStyle},
		{glib.Type(C.gtk_tree_model_flags_get_type()), marshalTreeModelFlags},
		{glib.Type(C.gtk_window_position_get_type()), marshalWindowPosition},
		{glib.Type(C.gtk_window_type_get_type()), marshalWindowType},
		{glib.Type(C.gtk_wrap_mode_get_type()), marshalWrapMode},

		// Objects/Interfaces
		{glib.Type(C.gtk_about_dialog_get_type()), marshalAboutDialog},
		{glib.Type(C.gtk_accel_group_get_type()), marshalAccelGroup},
		{glib.Type(C.gtk_accel_map_get_type()), marshalAccelMap},
		{glib.Type(C.gtk_adjustment_get_type()), marshalAdjustment},
		{glib.Type(C.gtk_assistant_get_type()), marshalAssistant},
		{glib.Type(C.gtk_bin_get_type()), marshalBin},
		{glib.Type(C.gtk_builder_get_type()), marshalBuilder},
		{glib.Type(C.gtk_button_get_type()), marshalButton},
		{glib.Type(C.gtk_box_get_type()), marshalBox},
		{glib.Type(C.gtk_calendar_get_type()), marshalCalendar},
		{glib.Type(C.gtk_cell_layout_get_type()), marshalCellLayout},
		{glib.Type(C.gtk_cell_renderer_get_type()), marshalCellRenderer},
		{glib.Type(C.gtk_cell_renderer_spinner_get_type()), marshalCellRendererSpinner},
		{glib.Type(C.gtk_cell_renderer_pixbuf_get_type()), marshalCellRendererPixbuf},
		{glib.Type(C.gtk_cell_renderer_text_get_type()), marshalCellRendererText},
		{glib.Type(C.gtk_cell_renderer_toggle_get_type()), marshalCellRendererToggle},
		{glib.Type(C.gtk_check_button_get_type()), marshalCheckButton},
		{glib.Type(C.gtk_check_menu_item_get_type()), marshalCheckMenuItem},
		{glib.Type(C.gtk_clipboard_get_type()), marshalClipboard},
		{glib.Type(C.gtk_combo_box_get_type()), marshalComboBox},
		{glib.Type(C.gtk_combo_box_text_get_type()), marshalComboBoxText},
		{glib.Type(C.gtk_container_get_type()), marshalContainer},
		{glib.Type(C.gtk_dialog_get_type()), marshalDialog},
		{glib.Type(C.gtk_drawing_area_get_type()), marshalDrawingArea},
		{glib.Type(C.gtk_editable_get_type()), marshalEditable},
		{glib.Type(C.gtk_entry_get_type()), marshalEntry},
		{glib.Type(C.gtk_entry_buffer_get_type()), marshalEntryBuffer},
		{glib.Type(C.gtk_entry_completion_get_type()), marshalEntryCompletion},
		{glib.Type(C.gtk_event_box_get_type()), marshalEventBox},
		{glib.Type(C.gtk_expander_get_type()), marshalExpander},
		{glib.Type(C.gtk_file_chooser_get_type()), marshalFileChooser},
		{glib.Type(C.gtk_file_chooser_button_get_type()), marshalFileChooserButton},
		{glib.Type(C.gtk_file_chooser_dialog_get_type()), marshalFileChooserDialog},
		{glib.Type(C.gtk_file_chooser_widget_get_type()), marshalFileChooserWidget},
		{glib.Type(C.gtk_font_button_get_type()), marshalFontButton},
		{glib.Type(C.gtk_frame_get_type()), marshalFrame},
		{glib.Type(C.gtk_grid_get_type()), marshalGrid},
		{glib.Type(C.gtk_icon_view_get_type()), marshalIconView},
		{glib.Type(C.gtk_image_get_type()), marshalImage},
		{glib.Type(C.gtk_label_get_type()), marshalLabel},
		{glib.Type(C.gtk_link_button_get_type()), marshalLinkButton},
		{glib.Type(C.gtk_layout_get_type()), marshalLayout},
		{glib.Type(C.gtk_list_store_get_type()), marshalListStore},
		{glib.Type(C.gtk_menu_get_type()), marshalMenu},
		{glib.Type(C.gtk_menu_bar_get_type()), marshalMenuBar},
		{glib.Type(C.gtk_menu_button_get_type()), marshalMenuButton},
		{glib.Type(C.gtk_menu_item_get_type()), marshalMenuItem},
		{glib.Type(C.gtk_menu_shell_get_type()), marshalMenuShell},
		{glib.Type(C.gtk_message_dialog_get_type()), marshalMessageDialog},
		{glib.Type(C.gtk_notebook_get_type()), marshalNotebook},
		{glib.Type(C.gtk_offscreen_window_get_type()), marshalOffscreenWindow},
		{glib.Type(C.gtk_orientable_get_type()), marshalOrientable},
		{glib.Type(C.gtk_paned_get_type()), marshalPaned},
		{glib.Type(C.gtk_progress_bar_get_type()), marshalProgressBar},
		{glib.Type(C.gtk_radio_button_get_type()), marshalRadioButton},
		{glib.Type(C.gtk_radio_menu_item_get_type()), marshalRadioMenuItem},
		{glib.Type(C.gtk_range_get_type()), marshalRange},
		{glib.Type(C.gtk_scale_button_get_type()), marshalScaleButton},
		{glib.Type(C.gtk_scale_get_type()), marshalScale},
		{glib.Type(C.gtk_scrollbar_get_type()), marshalScrollbar},
		{glib.Type(C.gtk_scrolled_window_get_type()), marshalScrolledWindow},
		{glib.Type(C.gtk_search_entry_get_type()), marshalSearchEntry},
		{glib.Type(C.gtk_selection_data_get_type()), marshalSelectionData},
		{glib.Type(C.gtk_separator_get_type()), marshalSeparator},
		{glib.Type(C.gtk_separator_menu_item_get_type()), marshalSeparatorMenuItem},
		{glib.Type(C.gtk_separator_tool_item_get_type()), marshalSeparatorToolItem},
		{glib.Type(C.gtk_spin_button_get_type()), marshalSpinButton},
		{glib.Type(C.gtk_spinner_get_type()), marshalSpinner},
		{glib.Type(C.gtk_statusbar_get_type()), marshalStatusbar},
		{glib.Type(C.gtk_switch_get_type()), marshalSwitch},
		{glib.Type(C.gtk_text_view_get_type()), marshalTextView},
		{glib.Type(C.gtk_text_tag_get_type()), marshalTextTag},
		{glib.Type(C.gtk_text_tag_table_get_type()), marshalTextTagTable},
		{glib.Type(C.gtk_text_buffer_get_type()), marshalTextBuffer},
		{glib.Type(C.gtk_toggle_button_get_type()), marshalToggleButton},
		{glib.Type(C.gtk_toolbar_get_type()), marshalToolbar},
		{glib.Type(C.gtk_tool_button_get_type()), marshalToolButton},
		{glib.Type(C.gtk_tool_item_get_type()), marshalToolItem},
		{glib.Type(C.gtk_tree_model_get_type()), marshalTreeModel},
		{glib.Type(C.gtk_tree_selection_get_type()), marshalTreeSelection},
		{glib.Type(C.gtk_tree_store_get_type()), marshalTreeStore},
		{glib.Type(C.gtk_tree_view_get_type()), marshalTreeView},
		{glib.Type(C.gtk_tree_view_column_get_type()), marshalTreeViewColumn},
		{glib.Type(C.gtk_volume_button_get_type()), marshalVolumeButton},
		{glib.Type(C.gtk_widget_get_type()), marshalWidget},
		{glib.Type(C.gtk_window_get_type()), marshalWindow},

		// Boxed
		{glib.Type(C.gtk_target_entry_get_type()), marshalTargetEntry},
		{glib.Type(C.gtk_text_iter_get_type()), marshalTextIter},
		{glib.Type(C.gtk_tree_iter_get_type()), marshalTreeIter},
		{glib.Type(C.gtk_tree_path_get_type()), marshalTreePath},
	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * Type conversions
 */

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

func gobool(b C.gboolean) bool {
	return b != C.FALSE
}

// Wrapper function for new objects with reference management.
func wrapObject(ptr unsafe.Pointer) *glib.Object {
	obj := &glib.Object{glib.ToGObject(ptr)}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return obj
}

// Wrapper function for TestBoolConvs since cgo can't be used with
// testing package
func testBoolConvs() error {
	b := gobool(gbool(true))
	if b != true {
		return errors.New("Unexpected bool conversion result")
	}

	cb := gbool(gobool(C.gboolean(0)))
	if cb != C.gboolean(0) {
		return errors.New("Unexpected bool conversion result")
	}

	return nil
}

/*
 * Unexported vars
 */

var nilPtrErr = errors.New("cgo returned unexpected nil pointer")

/*
 * Constants
 */

// Align is a representation of GTK's GtkAlign.
type Align int

const (
	ALIGN_FILL   Align = C.GTK_ALIGN_FILL
	ALIGN_START  Align = C.GTK_ALIGN_START
	ALIGN_END    Align = C.GTK_ALIGN_END
	ALIGN_CENTER Align = C.GTK_ALIGN_CENTER
)

func marshalAlign(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Align(c), nil
}

// ArrowPlacement is a representation of GTK's GtkArrowPlacement.
type ArrowPlacement int

const (
	ARROWS_BOTH  ArrowPlacement = C.GTK_ARROWS_BOTH
	ARROWS_START ArrowPlacement = C.GTK_ARROWS_START
	ARROWS_END   ArrowPlacement = C.GTK_ARROWS_END
)

func marshalArrowPlacement(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return ArrowPlacement(c), nil
}

// ArrowType is a representation of GTK's GtkArrowType.
type ArrowType int

const (
	ARROW_UP    ArrowType = C.GTK_ARROW_UP
	ARROW_DOWN  ArrowType = C.GTK_ARROW_DOWN
	ARROW_LEFT  ArrowType = C.GTK_ARROW_LEFT
	ARROW_RIGHT ArrowType = C.GTK_ARROW_RIGHT
	ARROW_NONE  ArrowType = C.GTK_ARROW_NONE
)

func marshalArrowType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return ArrowType(c), nil
}

// AssistantPageType is a representation of GTK's GtkAssistantPageType.
type AssistantPageType int

const (
	ASSISTANT_PAGE_CONTENT  AssistantPageType = C.GTK_ASSISTANT_PAGE_CONTENT
	ASSISTANT_PAGE_INTRO    AssistantPageType = C.GTK_ASSISTANT_PAGE_INTRO
	ASSISTANT_PAGE_CONFIRM  AssistantPageType = C.GTK_ASSISTANT_PAGE_CONFIRM
	ASSISTANT_PAGE_SUMMARY  AssistantPageType = C.GTK_ASSISTANT_PAGE_SUMMARY
	ASSISTANT_PAGE_PROGRESS AssistantPageType = C.GTK_ASSISTANT_PAGE_PROGRESS
	ASSISTANT_PAGE_CUSTOM   AssistantPageType = C.GTK_ASSISTANT_PAGE_CUSTOM
)

func marshalAssistantPageType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return AssistantPageType(c), nil
}

// ButtonsType is a representation of GTK's GtkButtonsType.
type ButtonsType int

const (
	BUTTONS_NONE      ButtonsType = C.GTK_BUTTONS_NONE
	BUTTONS_OK        ButtonsType = C.GTK_BUTTONS_OK
	BUTTONS_CLOSE     ButtonsType = C.GTK_BUTTONS_CLOSE
	BUTTONS_CANCEL    ButtonsType = C.GTK_BUTTONS_CANCEL
	BUTTONS_YES_NO    ButtonsType = C.GTK_BUTTONS_YES_NO
	BUTTONS_OK_CANCEL ButtonsType = C.GTK_BUTTONS_OK_CANCEL
)

func marshalButtonsType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return ButtonsType(c), nil
}

// CalendarDisplayOptions is a representation of GTK's GtkCalendarDisplayOptions
type CalendarDisplayOptions int

const (
	CALENDAR_SHOW_HEADING      CalendarDisplayOptions = C.GTK_CALENDAR_SHOW_HEADING
	CALENDAR_SHOW_DAY_NAMES    CalendarDisplayOptions = C.GTK_CALENDAR_SHOW_DAY_NAMES
	CALENDAR_NO_MONTH_CHANGE   CalendarDisplayOptions = C.GTK_CALENDAR_NO_MONTH_CHANGE
	CALENDAR_SHOW_WEEK_NUMBERS CalendarDisplayOptions = C.GTK_CALENDAR_SHOW_WEEK_NUMBERS
	CALENDAR_SHOW_DETAILS      CalendarDisplayOptions = C.GTK_CALENDAR_SHOW_DETAILS
)

func marshalCalendarDisplayOptions(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return CalendarDisplayOptions(c), nil
}

// DestDefaults is a representation of GTK's GtkDestDefaults.
type DestDefaults int

const (
	DEST_DEFAULT_MOTION    DestDefaults = C.GTK_DEST_DEFAULT_MOTION
	DEST_DEFAULT_HIGHLIGHT DestDefaults = C.GTK_DEST_DEFAULT_HIGHLIGHT
	DEST_DEFAULT_DROP      DestDefaults = C.GTK_DEST_DEFAULT_DROP
	DEST_DEFAULT_ALL       DestDefaults = C.GTK_DEST_DEFAULT_ALL
)

func marshalDestDefaults(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return DestDefaults(c), nil
}

// DialogFlags is a representation of GTK's GtkDialogFlags.
type DialogFlags int

const (
	DIALOG_MODAL               DialogFlags = C.GTK_DIALOG_MODAL
	DIALOG_DESTROY_WITH_PARENT DialogFlags = C.GTK_DIALOG_DESTROY_WITH_PARENT
)

func marshalDialogFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return DialogFlags(c), nil
}

// EntryIconPosition is a representation of GTK's GtkEntryIconPosition.
type EntryIconPosition int

const (
	ENTRY_ICON_PRIMARY   EntryIconPosition = C.GTK_ENTRY_ICON_PRIMARY
	ENTRY_ICON_SECONDARY EntryIconPosition = C.GTK_ENTRY_ICON_SECONDARY
)

func marshalEntryIconPosition(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return EntryIconPosition(c), nil
}

// FileChooserAction is a representation of GTK's GtkFileChooserAction.
type FileChooserAction int

const (
	FILE_CHOOSER_ACTION_OPEN          FileChooserAction = C.GTK_FILE_CHOOSER_ACTION_OPEN
	FILE_CHOOSER_ACTION_SAVE          FileChooserAction = C.GTK_FILE_CHOOSER_ACTION_SAVE
	FILE_CHOOSER_ACTION_SELECT_FOLDER FileChooserAction = C.GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER
	FILE_CHOOSER_ACTION_CREATE_FOLDER FileChooserAction = C.GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER
)

func marshalFileChooserAction(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return FileChooserAction(c), nil
}

// IconLookupFlags is a representation of GTK's GtkIconLookupFlags.
type IconLookupFlags int

const (
	ICON_LOOKUP_NO_SVG           IconLookupFlags = C.GTK_ICON_LOOKUP_NO_SVG
	ICON_LOOKUP_FORCE_SVG                        = C.GTK_ICON_LOOKUP_FORCE_SVG
	ICON_LOOKUP_USE_BUILTIN                      = C.GTK_ICON_LOOKUP_USE_BUILTIN
	ICON_LOOKUP_GENERIC_FALLBACK                 = C.GTK_ICON_LOOKUP_GENERIC_FALLBACK
	ICON_LOOKUP_FORCE_SIZE                       = C.GTK_ICON_LOOKUP_FORCE_SIZE
)

func marshalIconLookupFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return IconLookupFlags(c), nil
}

// IconSize is a representation of GTK's GtkIconSize.
type IconSize int

const (
	ICON_SIZE_INVALID       IconSize = C.GTK_ICON_SIZE_INVALID
	ICON_SIZE_MENU          IconSize = C.GTK_ICON_SIZE_MENU
	ICON_SIZE_SMALL_TOOLBAR IconSize = C.GTK_ICON_SIZE_SMALL_TOOLBAR
	ICON_SIZE_LARGE_TOOLBAR IconSize = C.GTK_ICON_SIZE_LARGE_TOOLBAR
	ICON_SIZE_BUTTON        IconSize = C.GTK_ICON_SIZE_BUTTON
	ICON_SIZE_DND           IconSize = C.GTK_ICON_SIZE_DND
	ICON_SIZE_DIALOG        IconSize = C.GTK_ICON_SIZE_DIALOG
)

func marshalIconSize(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return IconSize(c), nil
}

// ImageType is a representation of GTK's GtkImageType.
type ImageType int

const (
	IMAGE_EMPTY     ImageType = C.GTK_IMAGE_EMPTY
	IMAGE_PIXBUF    ImageType = C.GTK_IMAGE_PIXBUF
	IMAGE_STOCK     ImageType = C.GTK_IMAGE_STOCK
	IMAGE_ICON_SET  ImageType = C.GTK_IMAGE_ICON_SET
	IMAGE_ANIMATION ImageType = C.GTK_IMAGE_ANIMATION
	IMAGE_ICON_NAME ImageType = C.GTK_IMAGE_ICON_NAME
	IMAGE_GICON     ImageType = C.GTK_IMAGE_GICON
)

func marshalImageType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return ImageType(c), nil
}

// InputHints is a representation of GTK's GtkInputHints.
type InputHints int

const (
	INPUT_HINT_NONE                InputHints = C.GTK_INPUT_HINT_NONE
	INPUT_HINT_SPELLCHECK          InputHints = C.GTK_INPUT_HINT_SPELLCHECK
	INPUT_HINT_NO_SPELLCHECK       InputHints = C.GTK_INPUT_HINT_NO_SPELLCHECK
	INPUT_HINT_WORD_COMPLETION     InputHints = C.GTK_INPUT_HINT_WORD_COMPLETION
	INPUT_HINT_LOWERCASE           InputHints = C.GTK_INPUT_HINT_LOWERCASE
	INPUT_HINT_UPPERCASE_CHARS     InputHints = C.GTK_INPUT_HINT_UPPERCASE_CHARS
	INPUT_HINT_UPPERCASE_WORDS     InputHints = C.GTK_INPUT_HINT_UPPERCASE_WORDS
	INPUT_HINT_UPPERCASE_SENTENCES InputHints = C.GTK_INPUT_HINT_UPPERCASE_SENTENCES
	INPUT_HINT_INHIBIT_OSK         InputHints = C.GTK_INPUT_HINT_INHIBIT_OSK
)

func marshalInputHints(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return InputHints(c), nil
}

// InputPurpose is a representation of GTK's GtkInputPurpose.
type InputPurpose int

const (
	INPUT_PURPOSE_FREE_FORM InputPurpose = C.GTK_INPUT_PURPOSE_FREE_FORM
	INPUT_PURPOSE_ALPHA     InputPurpose = C.GTK_INPUT_PURPOSE_ALPHA
	INPUT_PURPOSE_DIGITS    InputPurpose = C.GTK_INPUT_PURPOSE_DIGITS
	INPUT_PURPOSE_NUMBER    InputPurpose = C.GTK_INPUT_PURPOSE_NUMBER
	INPUT_PURPOSE_PHONE     InputPurpose = C.GTK_INPUT_PURPOSE_PHONE
	INPUT_PURPOSE_URL       InputPurpose = C.GTK_INPUT_PURPOSE_URL
	INPUT_PURPOSE_EMAIL     InputPurpose = C.GTK_INPUT_PURPOSE_EMAIL
	INPUT_PURPOSE_NAME      InputPurpose = C.GTK_INPUT_PURPOSE_NAME
	INPUT_PURPOSE_PASSWORD  InputPurpose = C.GTK_INPUT_PURPOSE_PASSWORD
	INPUT_PURPOSE_PIN       InputPurpose = C.GTK_INPUT_PURPOSE_PIN
)

func marshalInputPurpose(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return InputPurpose(c), nil
}

// Justify is a representation of GTK's GtkJustification.
type Justification int

const (
	JUSTIFY_LEFT   Justification = C.GTK_JUSTIFY_LEFT
	JUSTIFY_RIGHT  Justification = C.GTK_JUSTIFY_RIGHT
	JUSTIFY_CENTER Justification = C.GTK_JUSTIFY_CENTER
	JUSTIFY_FILL   Justification = C.GTK_JUSTIFY_FILL
)

func marshalJustification(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Justification(c), nil
}

// License is a representation of GTK's GtkLicense.
type License int

const (
	LICENSE_UNKNOWN      License = C.GTK_LICENSE_UNKNOWN
	LICENSE_CUSTOM       License = C.GTK_LICENSE_CUSTOM
	LICENSE_GPL_2_0      License = C.GTK_LICENSE_GPL_2_0
	LICENSE_GPL_3_0      License = C.GTK_LICENSE_GPL_3_0
	LICENSE_LGPL_2_1     License = C.GTK_LICENSE_LGPL_2_1
	LICENSE_LGPL_3_0     License = C.GTK_LICENSE_LGPL_3_0
	LICENSE_BSD          License = C.GTK_LICENSE_BSD
	LICENSE_MIT_X11      License = C.GTK_LICENSE_MIT_X11
	LICENSE_GTK_ARTISTIC License = C.GTK_LICENSE_ARTISTIC
)

func marshalLicense(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return License(c), nil
}

// MessageType is a representation of GTK's GtkMessageType.
type MessageType int

const (
	MESSAGE_INFO     MessageType = C.GTK_MESSAGE_INFO
	MESSAGE_WARNING  MessageType = C.GTK_MESSAGE_WARNING
	MESSAGE_QUESTION MessageType = C.GTK_MESSAGE_QUESTION
	MESSAGE_ERROR    MessageType = C.GTK_MESSAGE_ERROR
	MESSAGE_OTHER    MessageType = C.GTK_MESSAGE_OTHER
)

func marshalMessageType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return MessageType(c), nil
}

// Orientation is a representation of GTK's GtkOrientation.
type Orientation int

const (
	ORIENTATION_HORIZONTAL Orientation = C.GTK_ORIENTATION_HORIZONTAL
	ORIENTATION_VERTICAL   Orientation = C.GTK_ORIENTATION_VERTICAL
)

func marshalOrientation(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Orientation(c), nil
}

// PackType is a representation of GTK's GtkPackType.
type PackType int

const (
	PACK_START PackType = C.GTK_PACK_START
	PACK_END   PackType = C.GTK_PACK_END
)

func marshalPackType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PackType(c), nil
}

// PathType is a representation of GTK's GtkPathType.
type PathType int

const (
	PATH_WIDGET       PathType = C.GTK_PATH_WIDGET
	PATH_WIDGET_CLASS PathType = C.GTK_PATH_WIDGET_CLASS
	PATH_CLASS        PathType = C.GTK_PATH_CLASS
)

func marshalPathType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PathType(c), nil
}

// PolicyType is a representation of GTK's GtkPolicyType.
type PolicyType int

const (
	POLICY_ALWAYS    PolicyType = C.GTK_POLICY_ALWAYS
	POLICY_AUTOMATIC PolicyType = C.GTK_POLICY_AUTOMATIC
	POLICY_NEVER     PolicyType = C.GTK_POLICY_NEVER
)

func marshalPolicyType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PolicyType(c), nil
}

// PositionType is a representation of GTK's GtkPositionType.
type PositionType int

const (
	POS_LEFT   PositionType = C.GTK_POS_LEFT
	POS_RIGHT  PositionType = C.GTK_POS_RIGHT
	POS_TOP    PositionType = C.GTK_POS_TOP
	POS_BOTTOM PositionType = C.GTK_POS_BOTTOM
)

func marshalPositionType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PositionType(c), nil
}

// ReliefStyle is a representation of GTK's GtkReliefStyle.
type ReliefStyle int

const (
	RELIEF_NORMAL ReliefStyle = C.GTK_RELIEF_NORMAL
	RELIEF_HALF   ReliefStyle = C.GTK_RELIEF_HALF
	RELIEF_NONE   ReliefStyle = C.GTK_RELIEF_NONE
)

func marshalReliefStyle(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return ReliefStyle(c), nil
}

// ResponseType is a representation of GTK's GtkResponseType.
type ResponseType int

const (
	RESPONSE_NONE         ResponseType = C.GTK_RESPONSE_NONE
	RESPONSE_REJECT       ResponseType = C.GTK_RESPONSE_REJECT
	RESPONSE_ACCEPT       ResponseType = C.GTK_RESPONSE_ACCEPT
	RESPONSE_DELETE_EVENT ResponseType = C.GTK_RESPONSE_DELETE_EVENT
	RESPONSE_OK           ResponseType = C.GTK_RESPONSE_OK
	RESPONSE_CANCEL       ResponseType = C.GTK_RESPONSE_CANCEL
	RESPONSE_CLOSE        ResponseType = C.GTK_RESPONSE_CLOSE
	RESPONSE_YES          ResponseType = C.GTK_RESPONSE_YES
	RESPONSE_NO           ResponseType = C.GTK_RESPONSE_NO
	RESPONSE_APPLY        ResponseType = C.GTK_RESPONSE_APPLY
	RESPONSE_HELP         ResponseType = C.GTK_RESPONSE_HELP
)

func marshalResponseType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return ResponseType(c), nil
}

// SelectionMode is a representation of GTK's GtkSelectionMode.
type SelectionMode int

const (
	SELECTION_NONE     SelectionMode = C.GTK_SELECTION_NONE
	SELECTION_SINGLE   SelectionMode = C.GTK_SELECTION_SINGLE
	SELECTION_BROWSE   SelectionMode = C.GTK_SELECTION_BROWSE
	SELECTION_MULTIPLE SelectionMode = C.GTK_SELECTION_MULTIPLE
)

func marshalSelectionMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SelectionMode(c), nil
}

// ShadowType is a representation of GTK's GtkShadowType.
type ShadowType int

const (
	SHADOW_NONE       ShadowType = C.GTK_SHADOW_NONE
	SHADOW_IN         ShadowType = C.GTK_SHADOW_IN
	SHADOW_OUT        ShadowType = C.GTK_SHADOW_OUT
	SHADOW_ETCHED_IN  ShadowType = C.GTK_SHADOW_ETCHED_IN
	SHADOW_ETCHED_OUT ShadowType = C.GTK_SHADOW_ETCHED_OUT
)

func marshalShadowType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return ShadowType(c), nil
}

// SortType is a representation of GTK's GtkSortType.
type SortType int

const (
	SORT_ASCENDING  SortType = C.GTK_SORT_ASCENDING
	SORT_DESCENDING          = C.GTK_SORT_DESCENDING
)

func marshalSortType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SortType(c), nil
}

// StateFlags is a representation of GTK's GtkStateFlags.
type StateFlags int

const (
	STATE_FLAG_NORMAL       StateFlags = C.GTK_STATE_FLAG_NORMAL
	STATE_FLAG_ACTIVE       StateFlags = C.GTK_STATE_FLAG_ACTIVE
	STATE_FLAG_PRELIGHT     StateFlags = C.GTK_STATE_FLAG_PRELIGHT
	STATE_FLAG_SELECTED     StateFlags = C.GTK_STATE_FLAG_SELECTED
	STATE_FLAG_INSENSITIVE  StateFlags = C.GTK_STATE_FLAG_INSENSITIVE
	STATE_FLAG_INCONSISTENT StateFlags = C.GTK_STATE_FLAG_INCONSISTENT
	STATE_FLAG_FOCUSED      StateFlags = C.GTK_STATE_FLAG_FOCUSED
	STATE_FLAG_BACKDROP     StateFlags = C.GTK_STATE_FLAG_BACKDROP
)

func marshalStateFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return StateFlags(c), nil
}

// TargetFlags is a representation of GTK's GtkTargetFlags.
type TargetFlags int

const (
	TARGET_SAME_APP     TargetFlags = C.GTK_TARGET_SAME_APP
	TARGET_SAME_WIDGET  TargetFlags = C.GTK_TARGET_SAME_WIDGET
	TARGET_OTHER_APP    TargetFlags = C.GTK_TARGET_OTHER_APP
	TARGET_OTHER_WIDGET TargetFlags = C.GTK_TARGET_OTHER_WIDGET
)

func marshalTargetFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return TargetFlags(c), nil
}

// ToolbarStyle is a representation of GTK's GtkToolbarStyle.
type ToolbarStyle int

const (
	TOOLBAR_ICONS      ToolbarStyle = C.GTK_TOOLBAR_ICONS
	TOOLBAR_TEXT       ToolbarStyle = C.GTK_TOOLBAR_TEXT
	TOOLBAR_BOTH       ToolbarStyle = C.GTK_TOOLBAR_BOTH
	TOOLBAR_BOTH_HORIZ ToolbarStyle = C.GTK_TOOLBAR_BOTH_HORIZ
)

func marshalToolbarStyle(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return ToolbarStyle(c), nil
}

// TreeModelFlags is a representation of GTK's GtkTreeModelFlags.
type TreeModelFlags int

const (
	TREE_MODEL_ITERS_PERSIST TreeModelFlags = C.GTK_TREE_MODEL_ITERS_PERSIST
	TREE_MODEL_LIST_ONLY     TreeModelFlags = C.GTK_TREE_MODEL_LIST_ONLY
)

func marshalTreeModelFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return TreeModelFlags(c), nil
}

// WindowPosition is a representation of GTK's GtkWindowPosition.
type WindowPosition int

const (
	WIN_POS_NONE             WindowPosition = C.GTK_WIN_POS_NONE
	WIN_POS_CENTER           WindowPosition = C.GTK_WIN_POS_CENTER
	WIN_POS_MOUSE            WindowPosition = C.GTK_WIN_POS_MOUSE
	WIN_POS_CENTER_ALWAYS    WindowPosition = C.GTK_WIN_POS_CENTER_ALWAYS
	WIN_POS_CENTER_ON_PARENT WindowPosition = C.GTK_WIN_POS_CENTER_ON_PARENT
)

func marshalWindowPosition(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return WindowPosition(c), nil
}

// WindowType is a representation of GTK's GtkWindowType.
type WindowType int

const (
	WINDOW_TOPLEVEL WindowType = C.GTK_WINDOW_TOPLEVEL
	WINDOW_POPUP    WindowType = C.GTK_WINDOW_POPUP
)

func marshalWindowType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return WindowType(c), nil
}

// WrapMode is a representation of GTK's GtkWrapMode.
type WrapMode int

const (
	WRAP_NONE      WrapMode = C.GTK_WRAP_NONE
	WRAP_CHAR      WrapMode = C.GTK_WRAP_CHAR
	WRAP_WORD      WrapMode = C.GTK_WRAP_WORD
	WRAP_WORD_CHAR WrapMode = C.GTK_WRAP_WORD_CHAR
)

func marshalWrapMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return WrapMode(c), nil
}

/*
 * Init and main event loop
 */

/*
Init() is a wrapper around gtk_init() and must be called before any
other GTK calls and is used to initialize everything necessary.

In addition to setting up GTK for usage, a pointer to a slice of
strings may be passed in to parse standard GTK command line arguments.
args will be modified to remove any flags that were handled.
Alternatively, nil may be passed in to not perform any command line
parsing.
*/
func Init(args *[]string) {
	if args != nil {
		argc := C.int(len(*args))
		argv := make([]*C.char, argc)
		for i, arg := range *args {
			argv[i] = C.CString(arg)
		}
		C.gtk_init((*C.int)(unsafe.Pointer(&argc)),
			(***C.char)(unsafe.Pointer(&argv)))
		unhandled := make([]string, argc)
		for i := 0; i < int(argc); i++ {
			unhandled[i] = C.GoString(argv[i])
			C.free(unsafe.Pointer(argv[i]))
		}
		*args = unhandled
	} else {
		C.gtk_init(nil, nil)
	}
}

// Main() is a wrapper around gtk_main() and runs the GTK main loop,
// blocking until MainQuit() is called.
func Main() {
	C.gtk_main()
}

// MainIterationDo is a wrapper around gtk_main_iteration_do
func MainIterationDo(blocking bool) bool {
	return gobool(C.gtk_main_iteration_do(gbool(blocking)))
}

// MainQuit() is a wrapper around gtk_main_quit() is used to terminate
// the GTK main loop (started by Main()).
func MainQuit() {
	C.gtk_main_quit()
}

/*
 * GtkAboutDialog
 */

// AboutDialog is a representation of GTK's GtkAboutDialog.
type AboutDialog struct {
	Dialog
}

// native returns a pointer to the underlying GtkAboutDialog.
func (v *AboutDialog) native() *C.GtkAboutDialog {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkAboutDialog(p)
}

func marshalAboutDialog(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapAboutDialog(obj), nil
}

func wrapAboutDialog(obj *glib.Object) *AboutDialog {
	return &AboutDialog{Dialog{Window{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}}}
}

// AboutDialogNew is a wrapper around gtk_about_dialog_new().
func AboutDialogNew() (*AboutDialog, error) {
	c := C.gtk_about_dialog_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAboutDialog(obj), nil
}

// GetComments is a wrapper around gtk_about_dialog_get_comments().
func (v *AboutDialog) GetComments() string {
	c := C.gtk_about_dialog_get_comments(v.native())
	return C.GoString((*C.char)(c))
}

// SetComments is a wrapper around gtk_about_dialog_set_comments().
func (v *AboutDialog) SetComments(comments string) {
	cstr := C.CString(comments)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_comments(v.native(), (*C.gchar)(cstr))
}

// GetCopyright is a wrapper around gtk_about_dialog_get_copyright().
func (v *AboutDialog) GetCopyright() string {
	c := C.gtk_about_dialog_get_copyright(v.native())
	return C.GoString((*C.char)(c))
}

// SetCopyright is a wrapper around gtk_about_dialog_set_copyright().
func (v *AboutDialog) SetCopyright(copyright string) {
	cstr := C.CString(copyright)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_copyright(v.native(), (*C.gchar)(cstr))
}

// GetLicense is a wrapper around gtk_about_dialog_get_license().
func (v *AboutDialog) GetLicense() string {
	c := C.gtk_about_dialog_get_license(v.native())
	return C.GoString((*C.char)(c))
}

// SetLicense is a wrapper around gtk_about_dialog_set_license().
func (v *AboutDialog) SetLicense(license string) {
	cstr := C.CString(license)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_license(v.native(), (*C.gchar)(cstr))
}

// GetLicenseType is a wrapper around gtk_about_dialog_get_license_type().
func (v *AboutDialog) GetLicenseType() License {
	c := C.gtk_about_dialog_get_license_type(v.native())
	return License(c)
}

// SetLicenseType is a wrapper around gtk_about_dialog_set_license_type().
func (v *AboutDialog) SetLicenseType(license License) {
	C.gtk_about_dialog_set_license_type(v.native(), C.GtkLicense(license))
}

// SetLogo is a wrapper around gtk_about_dialog_set_logo().
func (v *AboutDialog) SetLogo(logo *gdk.Pixbuf) {
	logoPtr := (*C.GdkPixbuf)(unsafe.Pointer(logo.Native()))
	C.gtk_about_dialog_set_logo(v.native(), logoPtr)
}

// GetLogoIconName is a wrapper around gtk_about_dialog_get_logo_icon_name().
func (v *AboutDialog) GetLogoIconName() string {
	c := C.gtk_about_dialog_get_logo_icon_name(v.native())
	return C.GoString((*C.char)(c))
}

// SetLogoIconName is a wrapper around gtk_about_dialog_set_logo_icon_name().
func (v *AboutDialog) SetLogoIconName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_logo_icon_name(v.native(), (*C.gchar)(cstr))
}

// GetProgramName is a wrapper around gtk_about_dialog_get_program_name().
func (v *AboutDialog) GetProgramName() string {
	c := C.gtk_about_dialog_get_program_name(v.native())
	return C.GoString((*C.char)(c))
}

// SetProgramName is a wrapper around gtk_about_dialog_set_program_name().
func (v *AboutDialog) SetProgramName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_program_name(v.native(), (*C.gchar)(cstr))
}

// GetAuthors is a wrapper around gtk_about_dialog_get_authors().
func (v *AboutDialog) GetAuthors() []string {
	var authors []string
	cauthors := C.gtk_about_dialog_get_authors(v.native())
	for {
		authors = append(authors, C.GoString((*C.char)(*cauthors)))
		cauthors = C.next_gcharptr(cauthors)
		if *cauthors == nil {
			break
		}
	}
	return authors
}

// SetAuthors is a wrapper around gtk_about_dialog_set_authors().
func (v *AboutDialog) SetAuthors(authors []string) {
	cauthors := C.make_strings(C.int(len(authors) + 1))
	for i, author := range authors {
		cstr := C.CString(author)
		defer C.free(unsafe.Pointer(cstr))
		C.set_string(cauthors, C.int(i), (*C.gchar)(cstr))
	}

	C.set_string(cauthors, C.int(len(authors)), nil)
	C.gtk_about_dialog_set_authors(v.native(), cauthors)
	C.destroy_strings(cauthors)
}

// GetTranslatorCredits is a wrapper around gtk_about_dialog_get_translator_credits().
func (v *AboutDialog) GetTranslatorCredits() string {
	c := C.gtk_about_dialog_get_translator_credits(v.native())
	return C.GoString((*C.char)(c))
}

// SetTranslatorCredits is a wrapper around gtk_about_dialog_set_translator_credits().
func (v *AboutDialog) SetTranslatorCredits(translatorCredits string) {
	cstr := C.CString(translatorCredits)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_translator_credits(v.native(), (*C.gchar)(cstr))
}

// GetVersion is a wrapper around gtk_about_dialog_get_version().
func (v *AboutDialog) GetVersion() string {
	c := C.gtk_about_dialog_get_version(v.native())
	return C.GoString((*C.char)(c))
}

// SetVersion is a wrapper around gtk_about_dialog_set_version().
func (v *AboutDialog) SetVersion(version string) {
	cstr := C.CString(version)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_version(v.native(), (*C.gchar)(cstr))
}

// GetWebsite is a wrapper around gtk_about_dialog_get_website().
func (v *AboutDialog) GetWebsite() string {
	c := C.gtk_about_dialog_get_website(v.native())
	return C.GoString((*C.char)(c))
}

// SetWebsite is a wrapper around gtk_about_dialog_set_website().
func (v *AboutDialog) SetWebsite(website string) {
	cstr := C.CString(website)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_website(v.native(), (*C.gchar)(cstr))
}

// GetWebsiteLabel is a wrapper around gtk_about_dialog_get_website_label().
func (v *AboutDialog) GetWebsiteLabel() string {
	c := C.gtk_about_dialog_get_website_label(v.native())
	return C.GoString((*C.char)(c))
}

// SetWebsiteLabel is a wrapper around gtk_about_dialog_set_website_label().
func (v *AboutDialog) SetWebsiteLabel(websiteLabel string) {
	cstr := C.CString(websiteLabel)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_website_label(v.native(), (*C.gchar)(cstr))
}

// GetWrapLicense is a wrapper around gtk_about_dialog_get_wrap_license().
func (v *AboutDialog) GetWrapLicense() bool {
	return gobool(C.gtk_about_dialog_get_wrap_license(v.native()))
}

// SetWrapLicense is a wrapper around gtk_about_dialog_set_wrap_license().
func (v *AboutDialog) SetWrapLicense(wrapLicense bool) {
	C.gtk_about_dialog_set_wrap_license(v.native(), gbool(wrapLicense))
}

/*
 * GtkAdjustment
 */

// Adjustment is a representation of GTK's GtkAdjustment.
type Adjustment struct {
	glib.InitiallyUnowned
}

// native returns a pointer to the underlying GtkAdjustment.
func (v *Adjustment) native() *C.GtkAdjustment {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkAdjustment(p)
}

func marshalAdjustment(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapAdjustment(obj), nil
}

func wrapAdjustment(obj *glib.Object) *Adjustment {
	return &Adjustment{glib.InitiallyUnowned{obj}}
}

// AdjustmentNew is a wrapper around gtk_adjustment_new().
func AdjustmentNew(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) (*Adjustment, error) {
	c := C.gtk_adjustment_new(C.gdouble(value),
		C.gdouble(lower),
		C.gdouble(upper),
		C.gdouble(stepIncrement),
		C.gdouble(pageIncrement),
		C.gdouble(pageSize))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAdjustment(obj), nil
}

// GetValue is a wrapper around gtk_adjustment_get_value().
func (v *Adjustment) GetValue() float64 {
	c := C.gtk_adjustment_get_value(v.native())
	return float64(c)
}

// SetValue is a wrapper around gtk_adjustment_set_value().
func (v *Adjustment) SetValue(value float64) {
	C.gtk_adjustment_set_value(v.native(), C.gdouble(value))
}

// GetLower is a wrapper around gtk_adjustment_get_lower().
func (v *Adjustment) GetLower() float64 {
	c := C.gtk_adjustment_get_lower(v.native())
	return float64(c)
}

// GetPageSize is a wrapper around gtk_adjustment_get_page_size().
func (v *Adjustment) GetPageSize() float64 {
	return float64(C.gtk_adjustment_get_page_size(v.native()))
}

// SetPageSize is a wrapper around gtk_adjustment_set_page_size().
func (v *Adjustment) SetPageSize(value float64) {
	C.gtk_adjustment_set_page_size(v.native(), C.gdouble(value))
}

// Configure is a wrapper around gtk_adjustment_configure().
func (v *Adjustment) Configure(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) {
	C.gtk_adjustment_configure(v.native(), C.gdouble(value),
		C.gdouble(lower), C.gdouble(upper), C.gdouble(stepIncrement),
		C.gdouble(pageIncrement), C.gdouble(pageSize))
}

// SetLower is a wrapper around gtk_adjustment_set_lower().
func (v *Adjustment) SetLower(value float64) {
	C.gtk_adjustment_set_lower(v.native(), C.gdouble(value))
}

// GetUpper is a wrapper around gtk_adjustment_get_upper().
func (v *Adjustment) GetUpper() float64 {
	c := C.gtk_adjustment_get_upper(v.native())
	return float64(c)
}

// SetUpper is a wrapper around gtk_adjustment_set_upper().
func (v *Adjustment) SetUpper(value float64) {
	C.gtk_adjustment_set_upper(v.native(), C.gdouble(value))
}

// GetPageIncrement is a wrapper around gtk_adjustment_get_page_increment().
func (v *Adjustment) GetPageIncrement() float64 {
	c := C.gtk_adjustment_get_page_increment(v.native())
	return float64(c)
}

// SetPageIncrement is a wrapper around gtk_adjustment_set_page_increment().
func (v *Adjustment) SetPageIncrement(value float64) {
	C.gtk_adjustment_set_page_increment(v.native(), C.gdouble(value))
}

// GetStepIncrement is a wrapper around gtk_adjustment_get_step_increment().
func (v *Adjustment) GetStepIncrement() float64 {
	c := C.gtk_adjustment_get_step_increment(v.native())
	return float64(c)
}

// SetStepIncrement is a wrapper around gtk_adjustment_set_step_increment().
func (v *Adjustment) SetStepIncrement(value float64) {
	C.gtk_adjustment_set_step_increment(v.native(), C.gdouble(value))
}

// GetMinimumIncrement is a wrapper around gtk_adjustment_get_minimum_increment().
func (v *Adjustment) GetMinimumIncrement() float64 {
	c := C.gtk_adjustment_get_minimum_increment(v.native())
	return float64(c)
}

/*
void	gtk_adjustment_clamp_page ()
void	gtk_adjustment_changed ()
void	gtk_adjustment_value_changed ()
void	gtk_adjustment_configure ()
*/

/*
 * GtkAssistant
 */

// Assistant is a representation of GTK's GtkAssistant.
type Assistant struct {
	Window
}

// native returns a pointer to the underlying GtkAssistant.
func (v *Assistant) native() *C.GtkAssistant {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkAssistant(p)
}

func marshalAssistant(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapAssistant(obj), nil
}

func wrapAssistant(obj *glib.Object) *Assistant {
	return &Assistant{Window{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}}
}

// AssistantNew is a wrapper around gtk_assistant_new().
func AssistantNew() (*Assistant, error) {
	c := C.gtk_assistant_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAssistant(obj), nil
}

// GetCurrentPage is a wrapper around gtk_assistant_get_current_page().
func (v *Assistant) GetCurrentPage() int {
	c := C.gtk_assistant_get_current_page(v.native())
	return int(c)
}

// SetCurrentPage is a wrapper around gtk_assistant_set_current_page().
func (v *Assistant) SetCurrentPage(pageNum int) {
	C.gtk_assistant_set_current_page(v.native(), C.gint(pageNum))
}

// GetNPages is a wrapper around gtk_assistant_get_n_pages().
func (v *Assistant) GetNPages() int {
	c := C.gtk_assistant_get_n_pages(v.native())
	return int(c)
}

// GetNthPage is a wrapper around gtk_assistant_get_nth_page().
func (v *Assistant) GetNthPage(pageNum int) *Widget {
	c := C.gtk_assistant_get_nth_page(v.native(), C.gint(pageNum))
	if c == nil {
		return nil
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapWidget(obj)
}

// PrependPage is a wrapper around gtk_assistant_prepend_page().
func (v *Assistant) PrependPage(page IWidget) int {
	c := C.gtk_assistant_prepend_page(v.native(), page.toWidget())
	return int(c)
}

// AppendPage is a wrapper around gtk_assistant_append_page().
func (v *Assistant) AppendPage(page IWidget) int {
	c := C.gtk_assistant_append_page(v.native(), page.toWidget())
	return int(c)
}

// InsertPage is a wrapper around gtk_assistant_insert_page().
func (v *Assistant) InsertPage(page IWidget, position int) int {
	c := C.gtk_assistant_insert_page(v.native(), page.toWidget(),
		C.gint(position))
	return int(c)
}

// RemovePage is a wrapper around gtk_assistant_remove_page().
func (v *Assistant) RemovePage(pageNum int) {
	C.gtk_assistant_remove_page(v.native(), C.gint(pageNum))
}

// TODO: gtk_assistant_set_forward_page_func

// SetPageType is a wrapper around gtk_assistant_set_page_type().
func (v *Assistant) SetPageType(page IWidget, ptype AssistantPageType) {
	C.gtk_assistant_set_page_type(v.native(), page.toWidget(),
		C.GtkAssistantPageType(ptype))
}

// GetPageType is a wrapper around gtk_assistant_get_page_type().
func (v *Assistant) GetPageType(page IWidget) AssistantPageType {
	c := C.gtk_assistant_get_page_type(v.native(), page.toWidget())
	return AssistantPageType(c)
}

// SetPageTitle is a wrapper around gtk_assistant_set_page_title().
func (v *Assistant) SetPageTitle(page IWidget, title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_assistant_set_page_title(v.native(), page.toWidget(),
		(*C.gchar)(cstr))
}

// GetPageTitle is a wrapper around gtk_assistant_get_page_title().
func (v *Assistant) GetPageTitle(page IWidget) string {
	c := C.gtk_assistant_get_page_title(v.native(), page.toWidget())
	return C.GoString((*C.char)(c))
}

// SetPageComplete is a wrapper around gtk_assistant_set_page_complete().
func (v *Assistant) SetPageComplete(page IWidget, complete bool) {
	C.gtk_assistant_set_page_complete(v.native(), page.toWidget(),
		gbool(complete))
}

// GetPageComplete is a wrapper around gtk_assistant_get_page_complete().
func (v *Assistant) GetPageComplete(page IWidget) bool {
	c := C.gtk_assistant_get_page_complete(v.native(), page.toWidget())
	return gobool(c)
}

// AddActionWidget is a wrapper around gtk_assistant_add_action_widget().
func (v *Assistant) AddActionWidget(child IWidget) {
	C.gtk_assistant_add_action_widget(v.native(), child.toWidget())
}

// RemoveActionWidget is a wrapper around gtk_assistant_remove_action_widget().
func (v *Assistant) RemoveActionWidget(child IWidget) {
	C.gtk_assistant_remove_action_widget(v.native(), child.toWidget())
}

// UpdateButtonsState is a wrapper around gtk_assistant_update_buttons_state().
func (v *Assistant) UpdateButtonsState() {
	C.gtk_assistant_update_buttons_state(v.native())
}

// Commit is a wrapper around gtk_assistant_commit().
func (v *Assistant) Commit() {
	C.gtk_assistant_commit(v.native())
}

// NextPage is a wrapper around gtk_assistant_next_page().
func (v *Assistant) NextPage() {
	C.gtk_assistant_next_page(v.native())
}

// PreviousPage is a wrapper around gtk_assistant_previous_page().
func (v *Assistant) PreviousPage() {
	C.gtk_assistant_previous_page(v.native())
}

/*
 * GtkBin
 */

// Bin is a representation of GTK's GtkBin.
type Bin struct {
	Container
}

// native returns a pointer to the underlying GtkBin.
func (v *Bin) native() *C.GtkBin {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkBin(p)
}

func marshalBin(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapBin(obj), nil
}

func wrapBin(obj *glib.Object) *Bin {
	return &Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// GetChild is a wrapper around gtk_bin_get_child().
func (v *Bin) GetChild() (*Widget, error) {
	c := C.gtk_bin_get_child(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

/*
 * GtkBuilder
 */

// Builder is a representation of GTK's GtkBuilder.
type Builder struct {
	*glib.Object
}

// native() returns a pointer to the underlying GtkBuilder.
func (b *Builder) native() *C.GtkBuilder {
	if b == nil || b.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(b.GObject)
	return C.toGtkBuilder(p)
}

func marshalBuilder(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &Builder{obj}, nil
}

// BuilderNew is a wrapper around gtk_builder_new().
func BuilderNew() (*Builder, error) {
	c := C.gtk_builder_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return &Builder{obj}, nil
}

// AddFromFile is a wrapper around gtk_builder_add_from_file().
func (b *Builder) AddFromFile(filename string) error {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	res := C.gtk_builder_add_from_file(b.native(), (*C.gchar)(cstr), &err)
	if res == 0 {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

// AddFromResource is a wrapper around gtk_builder_add_from_resource().
func (b *Builder) AddFromResource(path string) error {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	res := C.gtk_builder_add_from_resource(b.native(), (*C.gchar)(cstr), &err)
	if res == 0 {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

// AddFromString is a wrapper around gtk_builder_add_from_string().
func (b *Builder) AddFromString(str string) error {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	length := (C.gsize)(len(str))
	var err *C.GError = nil
	res := C.gtk_builder_add_from_string(b.native(), (*C.gchar)(cstr), length, &err)
	if res == 0 {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

// GetObject is a wrapper around gtk_builder_get_object(). The returned result
// is an IObject, so it will need to be type-asserted to the appropriate type before
// being used. For example, to get an object and type assert it as a window:
//
//   obj, err := builder.GetObject("window")
//   if err != nil {
//       // object not found
//       return
//   }
//   if w, ok := obj.(*gtk.Window); ok {
//       // do stuff with w here
//   } else {
//       // not a *gtk.Window
//   }
//
func (b *Builder) GetObject(name string) (glib.IObject, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_builder_get_object(b.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil, errors.New("object '" + name + "' not found")
	}
	obj, err := cast(c)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

var (
	builderSignals = struct {
		sync.RWMutex
		m map[*C.GtkBuilder]map[string]interface{}
	}{
		m: make(map[*C.GtkBuilder]map[string]interface{}),
	}
)

// ConnectSignals is a wrapper around gtk_builder_connect_signals_full().
func (b *Builder) ConnectSignals(signals map[string]interface{}) {
	builderSignals.Lock()
	builderSignals.m[b.native()] = signals
	builderSignals.Unlock()

	C._gtk_builder_connect_signals_full(b.native())
}

/*
 * GtkButton
 */

// Button is a representation of GTK's GtkButton.
type Button struct {
	Bin
}

// native() returns a pointer to the underlying GtkButton.
func (v *Button) native() *C.GtkButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkButton(p)
}

func marshalButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapButton(obj), nil
}

func wrapButton(obj *glib.Object) *Button {
	return &Button{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// ButtonNew() is a wrapper around gtk_button_new().
func ButtonNew() (*Button, error) {
	c := C.gtk_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapButton(obj), nil
}

// ButtonNewWithLabel() is a wrapper around gtk_button_new_with_label().
func ButtonNewWithLabel(label string) (*Button, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_button_new_with_label((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapButton(obj), nil
}

// ButtonNewWithMnemonic() is a wrapper around gtk_button_new_with_mnemonic().
func ButtonNewWithMnemonic(label string) (*Button, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_button_new_with_mnemonic((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapButton(obj), nil
}

// Clicked() is a wrapper around gtk_button_clicked().
func (v *Button) Clicked() {
	C.gtk_button_clicked(v.native())
}

// SetRelief() is a wrapper around gtk_button_set_relief().
func (v *Button) SetRelief(newStyle ReliefStyle) {
	C.gtk_button_set_relief(v.native(), C.GtkReliefStyle(newStyle))
}

// GetRelief() is a wrapper around gtk_button_get_relief().
func (v *Button) GetRelief() ReliefStyle {
	c := C.gtk_button_get_relief(v.native())
	return ReliefStyle(c)
}

// SetLabel() is a wrapper around gtk_button_set_label().
func (v *Button) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_button_set_label(v.native(), (*C.gchar)(cstr))
}

// GetLabel() is a wrapper around gtk_button_get_label().
func (v *Button) GetLabel() (string, error) {
	c := C.gtk_button_get_label(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// SetUseUnderline() is a wrapper around gtk_button_set_use_underline().
func (v *Button) SetUseUnderline(useUnderline bool) {
	C.gtk_button_set_use_underline(v.native(), gbool(useUnderline))
}

// GetUseUnderline() is a wrapper around gtk_button_get_use_underline().
func (v *Button) GetUseUnderline() bool {
	c := C.gtk_button_get_use_underline(v.native())
	return gobool(c)
}

// SetFocusOnClick() is a wrapper around gtk_button_set_focus_on_click().
func (v *Button) SetFocusOnClick(focusOnClick bool) {
	C.gtk_button_set_focus_on_click(v.native(), gbool(focusOnClick))
}

// GetFocusOnClick() is a wrapper around gtk_button_get_focus_on_click().
func (v *Button) GetFocusOnClick() bool {
	c := C.gtk_button_get_focus_on_click(v.native())
	return gobool(c)
}

// SetImage() is a wrapper around gtk_button_set_image().
func (v *Button) SetImage(image IWidget) {
	C.gtk_button_set_image(v.native(), image.toWidget())
}

// GetImage() is a wrapper around gtk_button_get_image().
func (v *Button) GetImage() (*Widget, error) {
	c := C.gtk_button_get_image(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// SetImagePosition() is a wrapper around gtk_button_set_image_position().
func (v *Button) SetImagePosition(position PositionType) {
	C.gtk_button_set_image_position(v.native(), C.GtkPositionType(position))
}

// GetImagePosition() is a wrapper around gtk_button_get_image_position().
func (v *Button) GetImagePosition() PositionType {
	c := C.gtk_button_get_image_position(v.native())
	return PositionType(c)
}

// SetAlwaysShowImage() is a wrapper around gtk_button_set_always_show_image().
func (v *Button) SetAlwaysShowImage(alwaysShow bool) {
	C.gtk_button_set_always_show_image(v.native(), gbool(alwaysShow))
}

// GetAlwaysShowImage() is a wrapper around gtk_button_get_always_show_image().
func (v *Button) GetAlwaysShowImage() bool {
	c := C.gtk_button_get_always_show_image(v.native())
	return gobool(c)
}

// GetEventWindow() is a wrapper around gtk_button_get_event_window().
func (v *Button) GetEventWindow() (*gdk.Window, error) {
	c := C.gtk_button_get_event_window(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := &gdk.Window{obj}
	w.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

/*
 * GtkColorButton
 */

// ColorButton is a representation of GTK's GtkColorButton.
type ColorButton struct {
	Button
}

// Native returns a pointer to the underlying GtkColorButton.
func (v *ColorButton) native() *C.GtkColorButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkColorButton(p)
}

func wrapColorButton(obj *glib.Object) *ColorButton {
	return &ColorButton{Button{Bin{Container{Widget{
		glib.InitiallyUnowned{obj}}}}}}
}

// ColorButtonNew is a wrapper around gtk_color_button_new().
func ColorButtonNew() (*ColorButton, error) {
	c := C.gtk_color_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	tb := wrapColorButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return tb, nil
}

// ColorButtonNewWithRGBA is a wrapper around gtk_color_button_new_with_rgba().
func ColorButtonNewWithRGBA(gdkColor *gdk.RGBA) (*ColorButton, error) {
	c := C.gtk_color_button_new_with_rgba((*C.GdkRGBA)(unsafe.Pointer(gdkColor.Native())))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	tb := wrapColorButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return tb, nil
}

// GetRGBA is a wrapper around gtk_color_chooser_get_rgba().
func (v *ColorButton) GetRGBA() *gdk.RGBA {
	gdkColor := gdk.NewRGBA()
	C.gtk_color_chooser_get_rgba(C.toGtkColorChooser(unsafe.Pointer(v.native())), (*C.GdkRGBA)(unsafe.Pointer(gdkColor.Native())))
	return gdkColor
}

// SetRGBA is a wrapper around gtk_color_chooser_set_rgba().
func (v *ColorButton) SetRGBA(gdkColor *gdk.RGBA) {
	C.gtk_color_chooser_set_rgba(C.toGtkColorChooser(unsafe.Pointer(v.native())), (*C.GdkRGBA)(unsafe.Pointer(gdkColor.Native())))
}

/*
 * GtkBox
 */

// Box is a representation of GTK's GtkBox.
type Box struct {
	Container
}

// native() returns a pointer to the underlying GtkBox.
func (v *Box) native() *C.GtkBox {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkBox(p)
}

func marshalBox(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapBox(obj), nil
}

func wrapBox(obj *glib.Object) *Box {
	return &Box{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// BoxNew() is a wrapper around gtk_box_new().
func BoxNew(orientation Orientation, spacing int) (*Box, error) {
	c := C.gtk_box_new(C.GtkOrientation(orientation), C.gint(spacing))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapBox(obj), nil
}

// PackStart() is a wrapper around gtk_box_pack_start().
func (v *Box) PackStart(child IWidget, expand, fill bool, padding uint) {
	C.gtk_box_pack_start(v.native(), child.toWidget(), gbool(expand),
		gbool(fill), C.guint(padding))
}

// PackEnd() is a wrapper around gtk_box_pack_end().
func (v *Box) PackEnd(child IWidget, expand, fill bool, padding uint) {
	C.gtk_box_pack_end(v.native(), child.toWidget(), gbool(expand),
		gbool(fill), C.guint(padding))
}

// GetHomogeneous() is a wrapper around gtk_box_get_homogeneous().
func (v *Box) GetHomogeneous() bool {
	c := C.gtk_box_get_homogeneous(v.native())
	return gobool(c)
}

// SetHomogeneous() is a wrapper around gtk_box_set_homogeneous().
func (v *Box) SetHomogeneous(homogeneous bool) {
	C.gtk_box_set_homogeneous(v.native(), gbool(homogeneous))
}

// GetSpacing() is a wrapper around gtk_box_get_spacing().
func (v *Box) GetSpacing() int {
	c := C.gtk_box_get_spacing(v.native())
	return int(c)
}

// SetSpacing() is a wrapper around gtk_box_set_spacing()
func (v *Box) SetSpacing(spacing int) {
	C.gtk_box_set_spacing(v.native(), C.gint(spacing))
}

// ReorderChild() is a wrapper around gtk_box_reorder_child().
func (v *Box) ReorderChild(child IWidget, position int) {
	C.gtk_box_reorder_child(v.native(), child.toWidget(), C.gint(position))
}

// QueryChildPacking() is a wrapper around gtk_box_query_child_packing().
func (v *Box) QueryChildPacking(child IWidget) (expand, fill bool, padding uint, packType PackType) {
	var cexpand, cfill C.gboolean
	var cpadding C.guint
	var cpackType C.GtkPackType

	C.gtk_box_query_child_packing(v.native(), child.toWidget(), &cexpand,
		&cfill, &cpadding, &cpackType)
	return gobool(cexpand), gobool(cfill), uint(cpadding), PackType(cpackType)
}

// SetChildPacking() is a wrapper around gtk_box_set_child_packing().
func (v *Box) SetChildPacking(child IWidget, expand, fill bool, padding uint, packType PackType) {
	C.gtk_box_set_child_packing(v.native(), child.toWidget(), gbool(expand),
		gbool(fill), C.guint(padding), C.GtkPackType(packType))
}

/*
 * GtkCalendar
 */

// Calendar is a representation of GTK's GtkCalendar.
type Calendar struct {
	Widget
}

// native() returns a pointer to the underlying GtkCalendar.
func (v *Calendar) native() *C.GtkCalendar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCalendar(p)
}

func marshalCalendar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapCalendar(obj), nil
}

func wrapCalendar(obj *glib.Object) *Calendar {
	return &Calendar{Widget{glib.InitiallyUnowned{obj}}}
}

// CalendarNew is a wrapper around gtk_calendar_new().
func CalendarNew() (*Calendar, error) {
	c := C.gtk_calendar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCalendar(obj), nil
}

// SelectMonth is a wrapper around gtk_calendar_select_month().
func (v *Calendar) SelectMonth(month, year uint) {
	C.gtk_calendar_select_month(v.native(), C.guint(month), C.guint(year))
}

// SelectDay is a wrapper around gtk_calendar_select_day().
func (v *Calendar) SelectDay(day uint) {
	C.gtk_calendar_select_day(v.native(), C.guint(day))
}

// MarkDay is a wrapper around gtk_calendar_mark_day().
func (v *Calendar) MarkDay(day uint) {
	C.gtk_calendar_mark_day(v.native(), C.guint(day))
}

// UnmarkDay is a wrapper around gtk_calendar_unmark_day().
func (v *Calendar) UnmarkDay(day uint) {
	C.gtk_calendar_unmark_day(v.native(), C.guint(day))
}

// GetDayIsMarked is a wrapper around gtk_calendar_get_day_is_marked().
func (v *Calendar) GetDayIsMarked(day uint) bool {
	c := C.gtk_calendar_get_day_is_marked(v.native(), C.guint(day))
	return gobool(c)
}

// ClearMarks is a wrapper around gtk_calendar_clear_marks().
func (v *Calendar) ClearMarks() {
	C.gtk_calendar_clear_marks(v.native())
}

// GetDisplayOptions is a wrapper around gtk_calendar_get_display_options().
func (v *Calendar) GetDisplayOptions() CalendarDisplayOptions {
	c := C.gtk_calendar_get_display_options(v.native())
	return CalendarDisplayOptions(c)
}

// SetDisplayOptions is a wrapper around gtk_calendar_set_display_options().
func (v *Calendar) SetDisplayOptions(flags CalendarDisplayOptions) {
	C.gtk_calendar_set_display_options(v.native(),
		C.GtkCalendarDisplayOptions(flags))
}

// GetDate is a wrapper around gtk_calendar_get_date().
func (v *Calendar) GetDate() (year, month, day uint) {
	var cyear, cmonth, cday C.guint
	C.gtk_calendar_get_date(v.native(), &cyear, &cmonth, &cday)
	return uint(cyear), uint(cmonth), uint(cday)
}

// TODO gtk_calendar_set_detail_func

// GetDetailWidthChars is a wrapper around gtk_calendar_get_detail_width_chars().
func (v *Calendar) GetDetailWidthChars() int {
	c := C.gtk_calendar_get_detail_width_chars(v.native())
	return int(c)
}

// SetDetailWidthChars is a wrapper around gtk_calendar_set_detail_width_chars().
func (v *Calendar) SetDetailWidthChars(chars int) {
	C.gtk_calendar_set_detail_width_chars(v.native(), C.gint(chars))
}

// GetDetailHeightRows is a wrapper around gtk_calendar_get_detail_height_rows().
func (v *Calendar) GetDetailHeightRows() int {
	c := C.gtk_calendar_get_detail_height_rows(v.native())
	return int(c)
}

// SetDetailHeightRows is a wrapper around gtk_calendar_set_detail_height_rows().
func (v *Calendar) SetDetailHeightRows(rows int) {
	C.gtk_calendar_set_detail_height_rows(v.native(), C.gint(rows))
}

/*
 * GtkCellLayout
 */

// CellLayout is a representation of GTK's GtkCellLayout GInterface.
type CellLayout struct {
	*glib.Object
}

// ICellLayout is an interface type implemented by all structs
// embedding a CellLayout.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkCellLayout.
type ICellLayout interface {
	toCellLayout() *C.GtkCellLayout
}

// native() returns a pointer to the underlying GObject as a GtkCellLayout.
func (v *CellLayout) native() *C.GtkCellLayout {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCellLayout(p)
}

func marshalCellLayout(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapCellLayout(obj), nil
}

func wrapCellLayout(obj *glib.Object) *CellLayout {
	return &CellLayout{obj}
}

func (v *CellLayout) toCellLayout() *C.GtkCellLayout {
	if v == nil {
		return nil
	}
	return v.native()
}

// PackStart() is a wrapper around gtk_cell_layout_pack_start().
func (v *CellLayout) PackStart(cell ICellRenderer, expand bool) {
	C.gtk_cell_layout_pack_start(v.native(), cell.toCellRenderer(),
		gbool(expand))
}

// AddAttribute() is a wrapper around gtk_cell_layout_add_attribute().
func (v *CellLayout) AddAttribute(cell ICellRenderer, attribute string, column int) {
	cstr := C.CString(attribute)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_cell_layout_add_attribute(v.native(), cell.toCellRenderer(),
		(*C.gchar)(cstr), C.gint(column))
}

/*
 * GtkCellRenderer
 */

// CellRenderer is a representation of GTK's GtkCellRenderer.
type CellRenderer struct {
	glib.InitiallyUnowned
}

// ICellRenderer is an interface type implemented by all structs
// embedding a CellRenderer.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkCellRenderer.
type ICellRenderer interface {
	toCellRenderer() *C.GtkCellRenderer
}

// native returns a pointer to the underlying GtkCellRenderer.
func (v *CellRenderer) native() *C.GtkCellRenderer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCellRenderer(p)
}

func (v *CellRenderer) toCellRenderer() *C.GtkCellRenderer {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalCellRenderer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapCellRenderer(obj), nil
}

func wrapCellRenderer(obj *glib.Object) *CellRenderer {
	return &CellRenderer{glib.InitiallyUnowned{obj}}
}

/*
 * GtkCellRendererSpinner
 */

// CellRendererSpinner is a representation of GTK's GtkCellRendererSpinner.
type CellRendererSpinner struct {
	CellRenderer
}

// native returns a pointer to the underlying GtkCellRendererSpinner.
func (v *CellRendererSpinner) native() *C.GtkCellRendererSpinner {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCellRendererSpinner(p)
}

func marshalCellRendererSpinner(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapCellRendererSpinner(obj), nil
}

func wrapCellRendererSpinner(obj *glib.Object) *CellRendererSpinner {
	return &CellRendererSpinner{CellRenderer{glib.InitiallyUnowned{obj}}}
}

// CellRendererSpinnerNew is a wrapper around gtk_cell_renderer_text_new().
func CellRendererSpinnerNew() (*CellRendererSpinner, error) {
	c := C.gtk_cell_renderer_spinner_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCellRendererSpinner(obj), nil
}

/*
 * GtkCellRendererPixbuf
 */

// CellRendererPixbuf is a representation of GTK's GtkCellRendererPixbuf.
type CellRendererPixbuf struct {
	CellRenderer
}

// native returns a pointer to the underlying GtkCellRendererPixbuf.
func (v *CellRendererPixbuf) native() *C.GtkCellRendererPixbuf {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCellRendererPixbuf(p)
}

func marshalCellRendererPixbuf(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapCellRendererPixbuf(obj), nil
}

func wrapCellRendererPixbuf(obj *glib.Object) *CellRendererPixbuf {
	return &CellRendererPixbuf{CellRenderer{glib.InitiallyUnowned{obj}}}
}

// CellRendererPixbufNew is a wrapper around gtk_cell_renderer_pixbuf_new().
func CellRendererPixbufNew() (*CellRendererPixbuf, error) {
	c := C.gtk_cell_renderer_pixbuf_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCellRendererPixbuf(obj), nil
}

/*
 * GtkCellRendererText
 */

// CellRendererText is a representation of GTK's GtkCellRendererText.
type CellRendererText struct {
	CellRenderer
}

// native returns a pointer to the underlying GtkCellRendererText.
func (v *CellRendererText) native() *C.GtkCellRendererText {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCellRendererText(p)
}

func marshalCellRendererText(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapCellRendererText(obj), nil
}

func wrapCellRendererText(obj *glib.Object) *CellRendererText {
	return &CellRendererText{CellRenderer{glib.InitiallyUnowned{obj}}}
}

// CellRendererTextNew is a wrapper around gtk_cell_renderer_text_new().
func CellRendererTextNew() (*CellRendererText, error) {
	c := C.gtk_cell_renderer_text_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCellRendererText(obj), nil
}

/*
 * GtkCellRendererToggle
 */

// CellRendererToggle is a representation of GTK's GtkCellRendererToggle.
type CellRendererToggle struct {
	CellRenderer
}

// native returns a pointer to the underlying GtkCellRendererToggle.
func (v *CellRendererToggle) native() *C.GtkCellRendererToggle {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCellRendererToggle(p)
}

func (v *CellRendererToggle) toCellRenderer() *C.GtkCellRenderer {
	if v == nil {
		return nil
	}
	return v.CellRenderer.native()
}

func marshalCellRendererToggle(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapCellRendererToggle(obj), nil
}

func wrapCellRendererToggle(obj *glib.Object) *CellRendererToggle {
	return &CellRendererToggle{CellRenderer{glib.InitiallyUnowned{obj}}}
}

// CellRendererToggleNew is a wrapper around gtk_cell_renderer_toggle_new().
func CellRendererToggleNew() (*CellRendererToggle, error) {
	c := C.gtk_cell_renderer_toggle_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCellRendererToggle(obj), nil
}

// SetRadio is a wrapper around gtk_cell_renderer_toggle_set_radio().
func (v *CellRendererToggle) SetRadio(set bool) {
	C.gtk_cell_renderer_toggle_set_radio(v.native(), gbool(set))
}

// GetRadio is a wrapper around gtk_cell_renderer_toggle_get_radio().
func (v *CellRendererToggle) GetRadio() bool {
	c := C.gtk_cell_renderer_toggle_get_radio(v.native())
	return gobool(c)
}

// SetActive is a wrapper arround gtk_cell_renderer_set_active().
func (v *CellRendererToggle) SetActive(active bool) {
	C.gtk_cell_renderer_toggle_set_active(v.native(), gbool(active))
}

// GetActive is a wrapper around gtk_cell_renderer_get_active().
func (v *CellRendererToggle) GetActive() bool {
	c := C.gtk_cell_renderer_toggle_get_active(v.native())
	return gobool(c)
}

// SetActivatable is a wrapper around gtk_cell_renderer_set_activatable().
func (v *CellRendererToggle) SetActivatable(activatable bool) {
	C.gtk_cell_renderer_toggle_set_activatable(v.native(),
		gbool(activatable))
}

// GetActivatable is a wrapper around gtk_cell_renderer_get_activatable().
func (v *CellRendererToggle) GetActivatable() bool {
	c := C.gtk_cell_renderer_toggle_get_activatable(v.native())
	return gobool(c)
}

/*
 * GtkCheckButton
 */

// CheckButton is a wrapper around GTK's GtkCheckButton.
type CheckButton struct {
	ToggleButton
}

// native returns a pointer to the underlying GtkCheckButton.
func (v *CheckButton) native() *C.GtkCheckButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCheckButton(p)
}

func marshalCheckButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapCheckButton(obj), nil
}

func wrapCheckButton(obj *glib.Object) *CheckButton {
	return &CheckButton{ToggleButton{Button{Bin{Container{Widget{
		glib.InitiallyUnowned{obj}}}}}}}
}

// CheckButtonNew is a wrapper around gtk_check_button_new().
func CheckButtonNew() (*CheckButton, error) {
	c := C.gtk_check_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCheckButton(obj), nil
}

// CheckButtonNewWithLabel is a wrapper around
// gtk_check_button_new_with_label().
func CheckButtonNewWithLabel(label string) (*CheckButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_check_button_new_with_label((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	cb := wrapCheckButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return cb, nil
}

// CheckButtonNewWithMnemonic is a wrapper around
// gtk_check_button_new_with_mnemonic().
func CheckButtonNewWithMnemonic(label string) (*CheckButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_check_button_new_with_mnemonic((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCheckButton(obj), nil
}

/*
 * GtkCheckMenuItem
 */

type CheckMenuItem struct {
	MenuItem
}

// native returns a pointer to the underlying GtkCheckMenuItem.
func (v *CheckMenuItem) native() *C.GtkCheckMenuItem {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCheckMenuItem(p)
}

func marshalCheckMenuItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapCheckMenuItem(obj), nil
}

func wrapCheckMenuItem(obj *glib.Object) *CheckMenuItem {
	return &CheckMenuItem{MenuItem{Bin{Container{Widget{
		glib.InitiallyUnowned{obj}}}}}}
}

// CheckMenuItemNew is a wrapper around gtk_check_menu_item_new().
func CheckMenuItemNew() (*CheckMenuItem, error) {
	c := C.gtk_check_menu_item_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCheckMenuItem(obj), nil
}

// CheckMenuItemNewWithLabel is a wrapper around
// gtk_check_menu_item_new_with_label().
func CheckMenuItemNewWithLabel(label string) (*CheckMenuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_check_menu_item_new_with_label((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCheckMenuItem(obj), nil
}

// CheckMenuItemNewWithMnemonic is a wrapper around
// gtk_check_menu_item_new_with_mnemonic().
func CheckMenuItemNewWithMnemonic(label string) (*CheckMenuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_check_menu_item_new_with_mnemonic((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCheckMenuItem(obj), nil
}

// GetActive is a wrapper around gtk_check_menu_item_get_active().
func (v *CheckMenuItem) GetActive() bool {
	c := C.gtk_check_menu_item_get_active(v.native())
	return gobool(c)
}

// SetActive is a wrapper around gtk_check_menu_item_set_active().
func (v *CheckMenuItem) SetActive(isActive bool) {
	C.gtk_check_menu_item_set_active(v.native(), gbool(isActive))
}

// Toggled is a wrapper around gtk_check_menu_item_toggled().
func (v *CheckMenuItem) Toggled() {
	C.gtk_check_menu_item_toggled(v.native())
}

// GetInconsistent is a wrapper around gtk_check_menu_item_get_inconsistent().
func (v *CheckMenuItem) GetInconsistent() bool {
	c := C.gtk_check_menu_item_get_inconsistent(v.native())
	return gobool(c)
}

// SetInconsistent is a wrapper around gtk_check_menu_item_set_inconsistent().
func (v *CheckMenuItem) SetInconsistent(setting bool) {
	C.gtk_check_menu_item_set_inconsistent(v.native(), gbool(setting))
}

// SetDrawAsRadio is a wrapper around gtk_check_menu_item_set_draw_as_radio().
func (v *CheckMenuItem) SetDrawAsRadio(drawAsRadio bool) {
	C.gtk_check_menu_item_set_draw_as_radio(v.native(), gbool(drawAsRadio))
}

// GetDrawAsRadio is a wrapper around gtk_check_menu_item_get_draw_as_radio().
func (v *CheckMenuItem) GetDrawAsRadio() bool {
	c := C.gtk_check_menu_item_get_draw_as_radio(v.native())
	return gobool(c)
}

/*
 * GtkClipboard
 */

// Clipboard is a wrapper around GTK's GtkClipboard.
type Clipboard struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkClipboard.
func (v *Clipboard) native() *C.GtkClipboard {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkClipboard(p)
}

func marshalClipboard(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapClipboard(obj), nil
}

func wrapClipboard(obj *glib.Object) *Clipboard {
	return &Clipboard{obj}
}

// Store is a wrapper around gtk_clipboard_store
func (v *Clipboard) Store() {
	C.gtk_clipboard_store(v.native())
}

// ClipboardGet() is a wrapper around gtk_clipboard_get().
func ClipboardGet(atom gdk.Atom) (*Clipboard, error) {
	c := C.gtk_clipboard_get(C.GdkAtom(unsafe.Pointer(atom)))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	cb := &Clipboard{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return cb, nil
}

// ClipboardGetForDisplay() is a wrapper around gtk_clipboard_get_for_display().
func ClipboardGetForDisplay(display *gdk.Display, atom gdk.Atom) (*Clipboard, error) {
	displayPtr := (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	c := C.gtk_clipboard_get_for_display(displayPtr,
		C.GdkAtom(unsafe.Pointer(atom)))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	cb := &Clipboard{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return cb, nil
}

// WaitIsTextAvailable is a wrapper around gtk_clipboard_wait_is_text_available
func (v *Clipboard) WaitIsTextAvailable() bool {
	c := C.gtk_clipboard_wait_is_text_available(v.native())
	return gobool(c)
}

// WaitForText is a wrapper around gtk_clipboard_wait_for_text
func (v *Clipboard) WaitForText() (string, error) {
	c := C.gtk_clipboard_wait_for_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	defer C.g_free(C.gpointer(c))
	return C.GoString((*C.char)(c)), nil
}

// SetText() is a wrapper around gtk_clipboard_set_text().
func (v *Clipboard) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_clipboard_set_text(v.native(), (*C.gchar)(cstr),
		C.gint(len(text)))
}

// WaitIsRichTextAvailable is a wrapper around gtk_clipboard_wait_is_rich_text_available
func (v *Clipboard) WaitIsRichTextAvailable(buf *TextBuffer) bool {
	c := C.gtk_clipboard_wait_is_rich_text_available(v.native(), buf.native())
	return gobool(c)
}

// WaitIsUrisAvailable is a wrapper around gtk_clipboard_wait_is_uris_available
func (v *Clipboard) WaitIsUrisAvailable() bool {
	c := C.gtk_clipboard_wait_is_uris_available(v.native())
	return gobool(c)
}

// WaitIsImageAvailable is a wrapper around gtk_clipboard_wait_is_image_available
func (v *Clipboard) WaitIsImageAvailable() bool {
	c := C.gtk_clipboard_wait_is_image_available(v.native())
	return gobool(c)
}

// SetImage is a wrapper around gtk_clipboard_set_image
func (v *Clipboard) SetImage(pixbuf *gdk.Pixbuf) {
	C.gtk_clipboard_set_image(v.native(), (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native())))
}

// WaitForImage is a wrapper around gtk_clipboard_wait_for_image
func (v *Clipboard) WaitForImage() (*gdk.Pixbuf, error) {
	c := C.gtk_clipboard_wait_for_image(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &gdk.Pixbuf{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// WaitIsTargetAvailable is a wrapper around gtk_clipboard_wait_is_target_available
func (v *Clipboard) WaitIsTargetAvailable(target gdk.Atom) bool {
	c := C.gtk_clipboard_wait_is_target_available(v.native(), C.GdkAtom(unsafe.Pointer(target)))
	return gobool(c)
}

// WaitForContents is a wrapper around gtk_clipboard_wait_for_contents
func (v *Clipboard) WaitForContents(target gdk.Atom) (*SelectionData, error) {
	c := C.gtk_clipboard_wait_for_contents(v.native(), C.GdkAtom(unsafe.Pointer(target)))
	if c == nil {
		return nil, nilPtrErr
	}
	p := &SelectionData{c}
	runtime.SetFinalizer(p, (*SelectionData).free)
	return p, nil
}

/*
 * GtkComboBox
 */

// ComboBox is a representation of GTK's GtkComboBox.
type ComboBox struct {
	Bin

	// Interfaces
	CellLayout
}

// native returns a pointer to the underlying GtkComboBox.
func (v *ComboBox) native() *C.GtkComboBox {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkComboBox(p)
}

func (v *ComboBox) toCellLayout() *C.GtkCellLayout {
	if v == nil {
		return nil
	}
	return C.toGtkCellLayout(unsafe.Pointer(v.GObject))
}

func marshalComboBox(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapComboBox(obj), nil
}

func wrapComboBox(obj *glib.Object) *ComboBox {
	cl := wrapCellLayout(obj)
	return &ComboBox{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}, *cl}
}

// ComboBoxNew() is a wrapper around gtk_combo_box_new().
func ComboBoxNew() (*ComboBox, error) {
	c := C.gtk_combo_box_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapComboBox(obj), nil
}

// ComboBoxNewWithEntry() is a wrapper around gtk_combo_box_new_with_entry().
func ComboBoxNewWithEntry() (*ComboBox, error) {
	c := C.gtk_combo_box_new_with_entry()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapComboBox(obj), nil
}

// ComboBoxNewWithModel() is a wrapper around gtk_combo_box_new_with_model().
func ComboBoxNewWithModel(model ITreeModel) (*ComboBox, error) {
	c := C.gtk_combo_box_new_with_model(model.toTreeModel())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapComboBox(obj), nil
}

// GetActive() is a wrapper around gtk_combo_box_get_active().
func (v *ComboBox) GetActive() int {
	c := C.gtk_combo_box_get_active(v.native())
	return int(c)
}

// SetActive() is a wrapper around gtk_combo_box_set_active().
func (v *ComboBox) SetActive(index int) {
	C.gtk_combo_box_set_active(v.native(), C.gint(index))
}

// GetActiveIter is a wrapper around gtk_combo_box_get_active_iter().
func (v *ComboBox) GetActiveIter() (*TreeIter, error) {
	var cIter C.GtkTreeIter
	c := C.gtk_combo_box_get_active_iter(v.native(), &cIter)
	if !gobool(c) {
		return nil, errors.New("unable to get active iter")
	}
	return &TreeIter{cIter}, nil
}

// SetActiveIter is a wrapper around gtk_combo_box_set_active_iter().
func (v *ComboBox) SetActiveIter(iter *TreeIter) {
	var cIter *C.GtkTreeIter
	if iter != nil {
		cIter = &iter.GtkTreeIter
	}
	C.gtk_combo_box_set_active_iter(v.native(), cIter)
}

// GetActiveID is a wrapper around gtk_combo_box_get_active_id().
func (v *ComboBox) GetActiveID() string {
	c := C.gtk_combo_box_get_active_id(v.native())
	return C.GoString((*C.char)(c))
}

// SetModel is a wrapper around gtk_combo_box_set_model().
func (v *ComboBox) SetModel(model ITreeModel) {
	C.gtk_combo_box_set_model(v.native(), model.toTreeModel())
}

/*
 * GtkComboBoxText
 */

// ComboBoxText is a representation of GTK's GtkComboBoxText.
type ComboBoxText struct {
	ComboBox
}

// native returns a pointer to the underlying GtkComboBoxText.
func (v *ComboBoxText) native() *C.GtkComboBoxText {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkComboBoxText(p)
}

func marshalComboBoxText(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapComboBoxText(obj), nil
}

func wrapComboBoxText(obj *glib.Object) *ComboBoxText {
	return &ComboBoxText{*wrapComboBox(obj)}
}

// ComboBoxTextNew is a wrapper around gtk_combo_box_text_new().
func ComboBoxTextNew() (*ComboBoxText, error) {
	c := C.gtk_combo_box_text_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapComboBoxText(obj), nil
}

// ComboBoxTextNewWithEntry is a wrapper around gtk_combo_box_text_new_with_entry().
func ComboBoxTextNewWithEntry() (*ComboBoxText, error) {
	c := C.gtk_combo_box_text_new_with_entry()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapComboBoxText(obj), nil
}

// AppendText is a wrapper around gtk_combo_box_text_append_text().
func (v *ComboBoxText) AppendText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_combo_box_text_append_text(v.native(), (*C.gchar)(cstr))
}

// GetActiveText is a wrapper around gtk_combo_box_text_get_active_text().
func (v *ComboBoxText) GetActiveText() string {
	c := (*C.char)(C.gtk_combo_box_text_get_active_text(v.native()))
	defer C.free(unsafe.Pointer(c))
	return C.GoString(c)
}

/*
 * GtkContainer
 */

// Container is a representation of GTK's GtkContainer.
type Container struct {
	Widget
}

// native returns a pointer to the underlying GtkContainer.
func (v *Container) native() *C.GtkContainer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkContainer(p)
}

func marshalContainer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapContainer(obj), nil
}

func wrapContainer(obj *glib.Object) *Container {
	return &Container{Widget{glib.InitiallyUnowned{obj}}}
}

// Add is a wrapper around gtk_container_add().
func (v *Container) Add(w IWidget) {
	C.gtk_container_add(v.native(), w.toWidget())
}

// Remove is a wrapper around gtk_container_remove().
func (v *Container) Remove(w IWidget) {
	C.gtk_container_remove(v.native(), w.toWidget())
}

// TODO: gtk_container_add_with_properties

// CheckResize is a wrapper around gtk_container_check_resize().
func (v *Container) CheckResize() {
	C.gtk_container_check_resize(v.native())
}

// TODO: gtk_container_foreach
// TODO: gtk_container_get_children
// TODO: gtk_container_get_path_for_child

// GetFocusChild is a wrapper around gtk_container_get_focus_child().
func (v *Container) GetFocusChild() *Widget {
	c := C.gtk_container_get_focus_child(v.native())
	if c == nil {
		return nil
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapWidget(obj)
}

// SetFocusChild is a wrapper around gtk_container_set_focus_child().
func (v *Container) SetFocusChild(child IWidget) {
	C.gtk_container_set_focus_child(v.native(), child.toWidget())
}

// GetFocusVAdjustment is a wrapper around
// gtk_container_get_focus_vadjustment().
func (v *Container) GetFocusVAdjustment() *Adjustment {
	c := C.gtk_container_get_focus_vadjustment(v.native())
	if c == nil {
		return nil
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAdjustment(obj)
}

// SetFocusVAdjustment is a wrapper around
// gtk_container_set_focus_vadjustment().
func (v *Container) SetFocusVAdjustment(adjustment *Adjustment) {
	C.gtk_container_set_focus_vadjustment(v.native(), adjustment.native())
}

// GetFocusHAdjustment is a wrapper around
// gtk_container_get_focus_hadjustment().
func (v *Container) GetFocusHAdjustment() *Adjustment {
	c := C.gtk_container_get_focus_hadjustment(v.native())
	if c == nil {
		return nil
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAdjustment(obj)
}

// SetFocusHAdjustment is a wrapper around
// gtk_container_set_focus_hadjustment().
func (v *Container) SetFocusHAdjustment(adjustment *Adjustment) {
	C.gtk_container_set_focus_hadjustment(v.native(), adjustment.native())
}

// ChildType is a wrapper around gtk_container_child_type().
func (v *Container) ChildType() glib.Type {
	c := C.gtk_container_child_type(v.native())
	return glib.Type(c)
}

// TODO: gtk_container_child_get_valist
// TODO: gtk_container_child_set_valist

// ChildNotify is a wrapper around gtk_container_child_notify().
func (v *Container) ChildNotify(child IWidget, childProperty string) {
	cstr := C.CString(childProperty)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_container_child_notify(v.native(), child.toWidget(),
		(*C.gchar)(cstr))
}

// ChildSetProperty is a wrapper around gtk_container_child_set_property().
func (v *Container) ChildSetProperty(child IWidget, name string, value interface{}) error {
	gv, e := glib.GValue(value)
	if e != nil {
		return e
	}
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_container_child_set_property(v.native(), child.toWidget(), (*C.gchar)(cstr), (*C.GValue)(unsafe.Pointer(gv)))
	return nil
}

// TODO: gtk_container_forall

// GetBorderWidth is a wrapper around gtk_container_get_border_width().
func (v *Container) GetBorderWidth() uint {
	c := C.gtk_container_get_border_width(v.native())
	return uint(c)
}

// SetBorderWidth is a wrapper around gtk_container_set_border_width().
func (v *Container) SetBorderWidth(borderWidth uint) {
	C.gtk_container_set_border_width(v.native(), C.guint(borderWidth))
}

// PropagateDraw is a wrapper around gtk_container_propagate_draw().
func (v *Container) PropagateDraw(child IWidget, cr *cairo.Context) {
	context := (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	C.gtk_container_propagate_draw(v.native(), child.toWidget(), context)
}

// GdkCairoSetSourcePixBuf() is a wrapper around gdk_cairo_set_source_pixbuf().
func GdkCairoSetSourcePixBuf(cr *cairo.Context, pixbuf *gdk.Pixbuf, pixbufX, pixbufY float64) {
	context := (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	ptr := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	C.gdk_cairo_set_source_pixbuf(context, ptr, C.gdouble(pixbufX), C.gdouble(pixbufY))
}

// GetFocusChain is a wrapper around gtk_container_get_focus_chain().
func (v *Container) GetFocusChain() ([]*Widget, bool) {
	var cwlist *C.GList
	c := C.gtk_container_get_focus_chain(v.native(), (**C.GList)(&cwlist))

	var widgets []*Widget
	wlist := glib.WrapList(uintptr(unsafe.Pointer(cwlist)))
	for ; wlist.Data() != uintptr(unsafe.Pointer(nil)); wlist = wlist.Next() {
		obj := &glib.Object{glib.ToGObject(unsafe.Pointer(wlist.Data()))}
		w := wrapWidget(obj)
		obj.RefSink()
		runtime.SetFinalizer(obj, (*glib.Object).Unref)
		widgets = append(widgets, w)
	}
	return widgets, gobool(c)
}

// SetFocusChain is a wrapper around gtk_container_set_focus_chain().
func (v *Container) SetFocusChain(focusableWidgets []IWidget) {
	var list *glib.List
	for _, w := range focusableWidgets {
		data := uintptr(unsafe.Pointer(w.toWidget()))
		list = list.Append(data)
	}
	glist := (*C.GList)(unsafe.Pointer(list))
	C.gtk_container_set_focus_chain(v.native(), (*C.GList)(glist))
}

/*
 * GtkCssProvider
 */

// CssProvider is a representation of GTK's GtkCssProvider.
type CssProvider struct {
	*glib.Object
}

func (v *CssProvider) toStyleProvider() *C.GtkStyleProvider {
	if v == nil {
		return nil
	}
	return C.toGtkStyleProvider(unsafe.Pointer(v.native()))
}

// native returns a pointer to the underlying GtkCssProvider.
func (v *CssProvider) native() *C.GtkCssProvider {
	if v == nil || v.Object == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCssProvider(p)
}

func wrapCssProvider(obj *glib.Object) *CssProvider {
	return &CssProvider{obj}
}

// CssProviderNew is a wrapper around gtk_css_provider_new().
func CssProviderNew() (*CssProvider, error) {
	c := C.gtk_css_provider_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	v := wrapCssProvider(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return v, nil
}

// LoadFromPath is a wrapper around gtk_css_provider_load_from_path().
func (v *CssProvider) LoadFromPath(path string) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	var gerr *C.GError
	if C.gtk_css_provider_load_from_path(v.native(), (*C.gchar)(cpath), &gerr) == 0 {
		defer C.g_error_free(gerr)
		return errors.New(C.GoString((*C.char)(gerr.message)))
	}
	return nil
}

// LoadFromData is a wrapper around gtk_css_provider_load_from_data().
func (v *CssProvider) LoadFromData(data string) error {
	cdata := C.CString(data)
	defer C.free(unsafe.Pointer(cdata))
	var gerr *C.GError
	if C.gtk_css_provider_load_from_data(v.native(), (*C.gchar)(unsafe.Pointer(cdata)), C.gssize(len(data)), &gerr) == 0 {
		defer C.g_error_free(gerr)
		return errors.New(C.GoString((*C.char)(gerr.message)))
	}
	return nil
}

// ToString is a wrapper around gtk_css_provider_to_string().
func (v *CssProvider) ToString() (string, error) {
	c := C.gtk_css_provider_to_string(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString(c), nil
}

// CssProviderGetDefault is a wrapper around gtk_css_provider_get_default().
func CssProviderGetDefault() (*CssProvider, error) {
	c := C.gtk_css_provider_get_default()
	if c == nil {
		return nil, nilPtrErr
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapCssProvider(obj), nil
}

// GetNamed is a wrapper around gtk_css_provider_get_named().
func (v *CssProvider) GetNamed(name string, variant string) (*CssProvider, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cvariant := C.CString(variant)
	defer C.free(unsafe.Pointer(cvariant))

	c := C.gtk_css_provider_get_named((*C.gchar)(cname), (*C.gchar)(cvariant))
	if c == nil {
		return nil, nilPtrErr
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapCssProvider(obj), nil
}

/*
 * GtkDialog
 */

// Dialog is a representation of GTK's GtkDialog.
type Dialog struct {
	Window
}

// native returns a pointer to the underlying GtkDialog.
func (v *Dialog) native() *C.GtkDialog {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkDialog(p)
}

func marshalDialog(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapDialog(obj), nil
}

func wrapDialog(obj *glib.Object) *Dialog {
	return &Dialog{Window{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}}
}

// DialogNew() is a wrapper around gtk_dialog_new().
func DialogNew() (*Dialog, error) {
	c := C.gtk_dialog_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapDialog(obj), nil
}

// Run() is a wrapper around gtk_dialog_run().
func (v *Dialog) Run() int {
	c := C.gtk_dialog_run(v.native())
	return int(c)
}

// Response() is a wrapper around gtk_dialog_response().
func (v *Dialog) Response(response ResponseType) {
	C.gtk_dialog_response(v.native(), C.gint(response))
}

// AddButton() is a wrapper around gtk_dialog_add_button().  text may
// be either the literal button text, or if using GTK 3.8 or earlier, a
// Stock type converted to a string.
func (v *Dialog) AddButton(text string, id ResponseType) (*Button, error) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_dialog_add_button(v.native(), (*C.gchar)(cstr), C.gint(id))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return &Button{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}, nil
}

// AddActionWidget() is a wrapper around gtk_dialog_add_action_widget().
func (v *Dialog) AddActionWidget(child IWidget, id ResponseType) {
	C.gtk_dialog_add_action_widget(v.native(), child.toWidget(), C.gint(id))
}

// SetDefaultResponse() is a wrapper around gtk_dialog_set_default_response().
func (v *Dialog) SetDefaultResponse(id ResponseType) {
	C.gtk_dialog_set_default_response(v.native(), C.gint(id))
}

// SetResponseSensitive() is a wrapper around
// gtk_dialog_set_response_sensitive().
func (v *Dialog) SetResponseSensitive(id ResponseType, setting bool) {
	C.gtk_dialog_set_response_sensitive(v.native(), C.gint(id),
		gbool(setting))
}

// GetResponseForWidget() is a wrapper around
// gtk_dialog_get_response_for_widget().
func (v *Dialog) GetResponseForWidget(widget IWidget) ResponseType {
	c := C.gtk_dialog_get_response_for_widget(v.native(), widget.toWidget())
	return ResponseType(c)
}

// GetWidgetForResponse() is a wrapper around
// gtk_dialog_get_widget_for_response().
func (v *Dialog) GetWidgetForResponse(id ResponseType) (*Widget, error) {
	c := C.gtk_dialog_get_widget_for_response(v.native(), C.gint(id))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// GetContentArea() is a wrapper around gtk_dialog_get_content_area().
func (v *Dialog) GetContentArea() (*Box, error) {
	c := C.gtk_dialog_get_content_area(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	b := &Box{Container{Widget{glib.InitiallyUnowned{obj}}}}
	return b, nil
}

// TODO(jrick)
/*
func (v *gdk.Screen) AlternativeDialogButtonOrder() bool {
	c := C.gtk_alternative_dialog_button_order(v.native())
	return gobool(c)
}
*/

// TODO(jrick)
/*
func SetAlternativeButtonOrder(ids ...ResponseType) {
}
*/

/*
 * GtkDrawingArea
 */

// DrawingArea is a representation of GTK's GtkDrawingArea.
type DrawingArea struct {
	Widget
}

// native returns a pointer to the underlying GtkDrawingArea.
func (v *DrawingArea) native() *C.GtkDrawingArea {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkDrawingArea(p)
}

func marshalDrawingArea(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapDrawingArea(obj), nil
}

func wrapDrawingArea(obj *glib.Object) *DrawingArea {
	return &DrawingArea{Widget{glib.InitiallyUnowned{obj}}}
}

// DrawingAreaNew is a wrapper around gtk_drawing_area_new().
func DrawingAreaNew() (*DrawingArea, error) {
	c := C.gtk_drawing_area_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapDrawingArea(obj), nil
}

/*
 * GtkEditable
 */

// Editable is a representation of GTK's GtkEditable GInterface.
type Editable struct {
	*glib.Object
}

// IEditable is an interface type implemented by all structs
// embedding an Editable.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkEditable.
type IEditable interface {
	toEditable() *C.GtkEditable
}

// native() returns a pointer to the underlying GObject as a GtkEditable.
func (v *Editable) native() *C.GtkEditable {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkEditable(p)
}

func marshalEditable(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapEditable(obj), nil
}

func wrapEditable(obj *glib.Object) *Editable {
	return &Editable{obj}
}

func (v *Editable) toEditable() *C.GtkEditable {
	if v == nil {
		return nil
	}
	return v.native()
}

// SelectRegion is a wrapper around gtk_editable_select_region().
func (v *Editable) SelectRegion(startPos, endPos int) {
	C.gtk_editable_select_region(v.native(), C.gint(startPos),
		C.gint(endPos))
}

// GetSelectionBounds is a wrapper around gtk_editable_get_selection_bounds().
func (v *Editable) GetSelectionBounds() (start, end int, nonEmpty bool) {
	var cstart, cend C.gint
	c := C.gtk_editable_get_selection_bounds(v.native(), &cstart, &cend)
	return int(cstart), int(cend), gobool(c)
}

// InsertText is a wrapper around gtk_editable_insert_text(). The returned
// int is the position after the inserted text.
func (v *Editable) InsertText(newText string, position int) int {
	cstr := C.CString(newText)
	defer C.free(unsafe.Pointer(cstr))
	pos := new(C.gint)
	*pos = C.gint(position)
	C.gtk_editable_insert_text(v.native(), (*C.gchar)(cstr),
		C.gint(len(newText)), pos)
	return int(*pos)
}

// DeleteText is a wrapper around gtk_editable_delete_text().
func (v *Editable) DeleteText(startPos, endPos int) {
	C.gtk_editable_delete_text(v.native(), C.gint(startPos), C.gint(endPos))
}

// GetChars is a wrapper around gtk_editable_get_chars().
func (v *Editable) GetChars(startPos, endPos int) string {
	c := C.gtk_editable_get_chars(v.native(), C.gint(startPos),
		C.gint(endPos))
	defer C.free(unsafe.Pointer(c))
	return C.GoString((*C.char)(c))
}

// CutClipboard is a wrapper around gtk_editable_cut_clipboard().
func (v *Editable) CutClipboard() {
	C.gtk_editable_cut_clipboard(v.native())
}

// CopyClipboard is a wrapper around gtk_editable_copy_clipboard().
func (v *Editable) CopyClipboard() {
	C.gtk_editable_copy_clipboard(v.native())
}

// PasteClipboard is a wrapper around gtk_editable_paste_clipboard().
func (v *Editable) PasteClipboard() {
	C.gtk_editable_paste_clipboard(v.native())
}

// DeleteSelection is a wrapper around gtk_editable_delete_selection().
func (v *Editable) DeleteSelection() {
	C.gtk_editable_delete_selection(v.native())
}

// SetPosition is a wrapper around gtk_editable_set_position().
func (v *Editable) SetPosition(position int) {
	C.gtk_editable_set_position(v.native(), C.gint(position))
}

// GetPosition is a wrapper around gtk_editable_get_position().
func (v *Editable) GetPosition() int {
	c := C.gtk_editable_get_position(v.native())
	return int(c)
}

// SetEditable is a wrapper around gtk_editable_set_editable().
func (v *Editable) SetEditable(isEditable bool) {
	C.gtk_editable_set_editable(v.native(), gbool(isEditable))
}

// GetEditable is a wrapper around gtk_editable_get_editable().
func (v *Editable) GetEditable() bool {
	c := C.gtk_editable_get_editable(v.native())
	return gobool(c)
}

/*
 * GtkEntry
 */

// Entry is a representation of GTK's GtkEntry.
type Entry struct {
	Widget

	// Interfaces
	Editable
}

type IEntry interface {
	toEntry() *C.GtkEntry
}

func (v *Entry) toEntry() *C.GtkEntry {
	return v.native()
}

// native returns a pointer to the underlying GtkEntry.
func (v *Entry) native() *C.GtkEntry {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkEntry(p)
}

func marshalEntry(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapEntry(obj), nil
}

func wrapEntry(obj *glib.Object) *Entry {
	e := wrapEditable(obj)
	return &Entry{Widget{glib.InitiallyUnowned{obj}}, *e}
}

// EntryNew() is a wrapper around gtk_entry_new().
func EntryNew() (*Entry, error) {
	c := C.gtk_entry_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapEntry(obj), nil
}

// EntryNewWithBuffer() is a wrapper around gtk_entry_new_with_buffer().
func EntryNewWithBuffer(buffer *EntryBuffer) (*Entry, error) {
	c := C.gtk_entry_new_with_buffer(buffer.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapEntry(obj), nil
}

// GetBuffer() is a wrapper around gtk_entry_get_buffer().
func (v *Entry) GetBuffer() (*EntryBuffer, error) {
	c := C.gtk_entry_get_buffer(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return &EntryBuffer{obj}, nil
}

// SetBuffer() is a wrapper around gtk_entry_set_buffer().
func (v *Entry) SetBuffer(buffer *EntryBuffer) {
	C.gtk_entry_set_buffer(v.native(), buffer.native())
}

// SetText() is a wrapper around gtk_entry_set_text().
func (v *Entry) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_set_text(v.native(), (*C.gchar)(cstr))
}

// GetText() is a wrapper around gtk_entry_get_text().
func (v *Entry) GetText() (string, error) {
	c := C.gtk_entry_get_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// GetTextLength() is a wrapper around gtk_entry_get_text_length().
func (v *Entry) GetTextLength() uint16 {
	c := C.gtk_entry_get_text_length(v.native())
	return uint16(c)
}

// TODO(jrick) GdkRectangle
/*
func (v *Entry) GetTextArea() {
}
*/

// SetVisibility() is a wrapper around gtk_entry_set_visibility().
func (v *Entry) SetVisibility(visible bool) {
	C.gtk_entry_set_visibility(v.native(), gbool(visible))
}

// SetInvisibleChar() is a wrapper around gtk_entry_set_invisible_char().
func (v *Entry) SetInvisibleChar(ch rune) {
	C.gtk_entry_set_invisible_char(v.native(), C.gunichar(ch))
}

// UnsetInvisibleChar() is a wrapper around gtk_entry_unset_invisible_char().
func (v *Entry) UnsetInvisibleChar() {
	C.gtk_entry_unset_invisible_char(v.native())
}

// SetMaxLength() is a wrapper around gtk_entry_set_max_length().
func (v *Entry) SetMaxLength(len int) {
	C.gtk_entry_set_max_length(v.native(), C.gint(len))
}

// GetActivatesDefault() is a wrapper around gtk_entry_get_activates_default().
func (v *Entry) GetActivatesDefault() bool {
	c := C.gtk_entry_get_activates_default(v.native())
	return gobool(c)
}

// GetHasFrame() is a wrapper around gtk_entry_get_has_frame().
func (v *Entry) GetHasFrame() bool {
	c := C.gtk_entry_get_has_frame(v.native())
	return gobool(c)
}

// GetWidthChars() is a wrapper around gtk_entry_get_width_chars().
func (v *Entry) GetWidthChars() int {
	c := C.gtk_entry_get_width_chars(v.native())
	return int(c)
}

// SetActivatesDefault() is a wrapper around gtk_entry_set_activates_default().
func (v *Entry) SetActivatesDefault(setting bool) {
	C.gtk_entry_set_activates_default(v.native(), gbool(setting))
}

// SetHasFrame() is a wrapper around gtk_entry_set_has_frame().
func (v *Entry) SetHasFrame(setting bool) {
	C.gtk_entry_set_has_frame(v.native(), gbool(setting))
}

// SetWidthChars() is a wrapper around gtk_entry_set_width_chars().
func (v *Entry) SetWidthChars(nChars int) {
	C.gtk_entry_set_width_chars(v.native(), C.gint(nChars))
}

// GetInvisibleChar() is a wrapper around gtk_entry_get_invisible_char().
func (v *Entry) GetInvisibleChar() rune {
	c := C.gtk_entry_get_invisible_char(v.native())
	return rune(c)
}

// SetAlignment() is a wrapper around gtk_entry_set_alignment().
func (v *Entry) SetAlignment(xalign float32) {
	C.gtk_entry_set_alignment(v.native(), C.gfloat(xalign))
}

// GetAlignment() is a wrapper around gtk_entry_get_alignment().
func (v *Entry) GetAlignment() float32 {
	c := C.gtk_entry_get_alignment(v.native())
	return float32(c)
}

// SetPlaceholderText() is a wrapper around gtk_entry_set_placeholder_text().
func (v *Entry) SetPlaceholderText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_set_placeholder_text(v.native(), (*C.gchar)(cstr))
}

// GetPlaceholderText() is a wrapper around gtk_entry_get_placeholder_text().
func (v *Entry) GetPlaceholderText() (string, error) {
	c := C.gtk_entry_get_placeholder_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// SetOverwriteMode() is a wrapper around gtk_entry_set_overwrite_mode().
func (v *Entry) SetOverwriteMode(overwrite bool) {
	C.gtk_entry_set_overwrite_mode(v.native(), gbool(overwrite))
}

// GetOverwriteMode() is a wrapper around gtk_entry_get_overwrite_mode().
func (v *Entry) GetOverwriteMode() bool {
	c := C.gtk_entry_get_overwrite_mode(v.native())
	return gobool(c)
}

// TODO(jrick) Pangolayout
/*
func (v *Entry) GetLayout() {
}
*/

// GetLayoutOffsets() is a wrapper around gtk_entry_get_layout_offsets().
func (v *Entry) GetLayoutOffsets() (x, y int) {
	var gx, gy C.gint
	C.gtk_entry_get_layout_offsets(v.native(), &gx, &gy)
	return int(gx), int(gy)
}

// LayoutIndexToTextIndex() is a wrapper around
// gtk_entry_layout_index_to_text_index().
func (v *Entry) LayoutIndexToTextIndex(layoutIndex int) int {
	c := C.gtk_entry_layout_index_to_text_index(v.native(),
		C.gint(layoutIndex))
	return int(c)
}

// TextIndexToLayoutIndex() is a wrapper around
// gtk_entry_text_index_to_layout_index().
func (v *Entry) TextIndexToLayoutIndex(textIndex int) int {
	c := C.gtk_entry_text_index_to_layout_index(v.native(),
		C.gint(textIndex))
	return int(c)
}

// TODO(jrick) PandoAttrList
/*
func (v *Entry) SetAttributes() {
}
*/

// TODO(jrick) PandoAttrList
/*
func (v *Entry) GetAttributes() {
}
*/

// GetMaxLength() is a wrapper around gtk_entry_get_max_length().
func (v *Entry) GetMaxLength() int {
	c := C.gtk_entry_get_max_length(v.native())
	return int(c)
}

// GetVisibility() is a wrapper around gtk_entry_get_visibility().
func (v *Entry) GetVisibility() bool {
	c := C.gtk_entry_get_visibility(v.native())
	return gobool(c)
}

// SetCompletion() is a wrapper around gtk_entry_set_completion().
func (v *Entry) SetCompletion(completion *EntryCompletion) {
	C.gtk_entry_set_completion(v.native(), completion.native())
}

// GetCompletion() is a wrapper around gtk_entry_get_completion().
func (v *Entry) GetCompletion() (*EntryCompletion, error) {
	c := C.gtk_entry_get_completion(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	e := &EntryCompletion{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return e, nil
}

// SetCursorHAdjustment() is a wrapper around
// gtk_entry_set_cursor_hadjustment().
func (v *Entry) SetCursorHAdjustment(adjustment *Adjustment) {
	C.gtk_entry_set_cursor_hadjustment(v.native(), adjustment.native())
}

// GetCursorHAdjustment() is a wrapper around
// gtk_entry_get_cursor_hadjustment().
func (v *Entry) GetCursorHAdjustment() (*Adjustment, error) {
	c := C.gtk_entry_get_cursor_hadjustment(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return &Adjustment{glib.InitiallyUnowned{obj}}, nil
}

// SetProgressFraction() is a wrapper around gtk_entry_set_progress_fraction().
func (v *Entry) SetProgressFraction(fraction float64) {
	C.gtk_entry_set_progress_fraction(v.native(), C.gdouble(fraction))
}

// GetProgressFraction() is a wrapper around gtk_entry_get_progress_fraction().
func (v *Entry) GetProgressFraction() float64 {
	c := C.gtk_entry_get_progress_fraction(v.native())
	return float64(c)
}

// SetProgressPulseStep() is a wrapper around
// gtk_entry_set_progress_pulse_step().
func (v *Entry) SetProgressPulseStep(fraction float64) {
	C.gtk_entry_set_progress_pulse_step(v.native(), C.gdouble(fraction))
}

// GetProgressPulseStep() is a wrapper around
// gtk_entry_get_progress_pulse_step().
func (v *Entry) GetProgressPulseStep() float64 {
	c := C.gtk_entry_get_progress_pulse_step(v.native())
	return float64(c)
}

// ProgressPulse() is a wrapper around gtk_entry_progress_pulse().
func (v *Entry) ProgressPulse() {
	C.gtk_entry_progress_pulse(v.native())
}

// TODO(jrick) GdkEventKey
/*
func (v *Entry) IMContextFilterKeypress() {
}
*/

// ResetIMContext() is a wrapper around gtk_entry_reset_im_context().
func (v *Entry) ResetIMContext() {
	C.gtk_entry_reset_im_context(v.native())
}

// TODO(jrick) GdkPixbuf
/*
func (v *Entry) SetIconFromPixbuf() {
}
*/

// SetIconFromIconName() is a wrapper around
// gtk_entry_set_icon_from_icon_name().
func (v *Entry) SetIconFromIconName(iconPos EntryIconPosition, name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_set_icon_from_icon_name(v.native(),
		C.GtkEntryIconPosition(iconPos), (*C.gchar)(cstr))
}

// TODO(jrick) GIcon
/*
func (v *Entry) SetIconFromGIcon() {
}
*/

// GetIconStorageType() is a wrapper around gtk_entry_get_icon_storage_type().
func (v *Entry) GetIconStorageType(iconPos EntryIconPosition) ImageType {
	c := C.gtk_entry_get_icon_storage_type(v.native(),
		C.GtkEntryIconPosition(iconPos))
	return ImageType(c)
}

// TODO(jrick) GdkPixbuf
/*
func (v *Entry) GetIconPixbuf() {
}
*/

// GetIconName() is a wrapper around gtk_entry_get_icon_name().
func (v *Entry) GetIconName(iconPos EntryIconPosition) (string, error) {
	c := C.gtk_entry_get_icon_name(v.native(),
		C.GtkEntryIconPosition(iconPos))
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// TODO(jrick) GIcon
/*
func (v *Entry) GetIconGIcon() {
}
*/

// SetIconActivatable() is a wrapper around gtk_entry_set_icon_activatable().
func (v *Entry) SetIconActivatable(iconPos EntryIconPosition, activatable bool) {
	C.gtk_entry_set_icon_activatable(v.native(),
		C.GtkEntryIconPosition(iconPos), gbool(activatable))
}

// GetIconActivatable() is a wrapper around gtk_entry_get_icon_activatable().
func (v *Entry) GetIconActivatable(iconPos EntryIconPosition) bool {
	c := C.gtk_entry_get_icon_activatable(v.native(),
		C.GtkEntryIconPosition(iconPos))
	return gobool(c)
}

// SetIconSensitive() is a wrapper around gtk_entry_set_icon_sensitive().
func (v *Entry) SetIconSensitive(iconPos EntryIconPosition, sensitive bool) {
	C.gtk_entry_set_icon_sensitive(v.native(),
		C.GtkEntryIconPosition(iconPos), gbool(sensitive))
}

// GetIconSensitive() is a wrapper around gtk_entry_get_icon_sensitive().
func (v *Entry) GetIconSensitive(iconPos EntryIconPosition) bool {
	c := C.gtk_entry_get_icon_sensitive(v.native(),
		C.GtkEntryIconPosition(iconPos))
	return gobool(c)
}

// GetIconAtPos() is a wrapper around gtk_entry_get_icon_at_pos().
func (v *Entry) GetIconAtPos(x, y int) int {
	c := C.gtk_entry_get_icon_at_pos(v.native(), C.gint(x), C.gint(y))
	return int(c)
}

// SetIconTooltipText() is a wrapper around gtk_entry_set_icon_tooltip_text().
func (v *Entry) SetIconTooltipText(iconPos EntryIconPosition, tooltip string) {
	cstr := C.CString(tooltip)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_set_icon_tooltip_text(v.native(),
		C.GtkEntryIconPosition(iconPos), (*C.gchar)(cstr))
}

// GetIconTooltipText() is a wrapper around gtk_entry_get_icon_tooltip_text().
func (v *Entry) GetIconTooltipText(iconPos EntryIconPosition) (string, error) {
	c := C.gtk_entry_get_icon_tooltip_text(v.native(),
		C.GtkEntryIconPosition(iconPos))
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// SetIconTooltipMarkup() is a wrapper around
// gtk_entry_set_icon_tooltip_markup().
func (v *Entry) SetIconTooltipMarkup(iconPos EntryIconPosition, tooltip string) {
	cstr := C.CString(tooltip)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_set_icon_tooltip_markup(v.native(),
		C.GtkEntryIconPosition(iconPos), (*C.gchar)(cstr))
}

// GetIconTooltipMarkup() is a wrapper around
// gtk_entry_get_icon_tooltip_markup().
func (v *Entry) GetIconTooltipMarkup(iconPos EntryIconPosition) (string, error) {
	c := C.gtk_entry_get_icon_tooltip_markup(v.native(),
		C.GtkEntryIconPosition(iconPos))
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// TODO(jrick) GdkDragAction
/*
func (v *Entry) SetIconDragSource() {
}
*/

// GetCurrentIconDragSource() is a wrapper around
// gtk_entry_get_current_icon_drag_source().
func (v *Entry) GetCurrentIconDragSource() int {
	c := C.gtk_entry_get_current_icon_drag_source(v.native())
	return int(c)
}

// TODO(jrick) GdkRectangle
/*
func (v *Entry) GetIconArea() {
}
*/

// SetInputPurpose() is a wrapper around gtk_entry_set_input_purpose().
func (v *Entry) SetInputPurpose(purpose InputPurpose) {
	C.gtk_entry_set_input_purpose(v.native(), C.GtkInputPurpose(purpose))
}

// GetInputPurpose() is a wrapper around gtk_entry_get_input_purpose().
func (v *Entry) GetInputPurpose() InputPurpose {
	c := C.gtk_entry_get_input_purpose(v.native())
	return InputPurpose(c)
}

// SetInputHints() is a wrapper around gtk_entry_set_input_hints().
func (v *Entry) SetInputHints(hints InputHints) {
	C.gtk_entry_set_input_hints(v.native(), C.GtkInputHints(hints))
}

// GetInputHints() is a wrapper around gtk_entry_get_input_hints().
func (v *Entry) GetInputHints() InputHints {
	c := C.gtk_entry_get_input_hints(v.native())
	return InputHints(c)
}

/*
 * GtkEntryBuffer
 */

// EntryBuffer is a representation of GTK's GtkEntryBuffer.
type EntryBuffer struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkEntryBuffer.
func (v *EntryBuffer) native() *C.GtkEntryBuffer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkEntryBuffer(p)
}

func marshalEntryBuffer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapEntryBuffer(obj), nil
}

func wrapEntryBuffer(obj *glib.Object) *EntryBuffer {
	return &EntryBuffer{obj}
}

// EntryBufferNew() is a wrapper around gtk_entry_buffer_new().
func EntryBufferNew(initialChars string, nInitialChars int) (*EntryBuffer, error) {
	cstr := C.CString(initialChars)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_entry_buffer_new((*C.gchar)(cstr), C.gint(nInitialChars))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	e := wrapEntryBuffer(obj)
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return e, nil
}

// GetText() is a wrapper around gtk_entry_buffer_get_text().  A
// non-nil error is returned in the case that gtk_entry_buffer_get_text
// returns NULL to differentiate between NULL and an empty string.
func (v *EntryBuffer) GetText() (string, error) {
	c := C.gtk_entry_buffer_get_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// SetText() is a wrapper around gtk_entry_buffer_set_text().
func (v *EntryBuffer) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_buffer_set_text(v.native(), (*C.gchar)(cstr),
		C.gint(len(text)))
}

// GetBytes() is a wrapper around gtk_entry_buffer_get_bytes().
func (v *EntryBuffer) GetBytes() uint {
	c := C.gtk_entry_buffer_get_bytes(v.native())
	return uint(c)
}

// GetLength() is a wrapper around gtk_entry_buffer_get_length().
func (v *EntryBuffer) GetLength() uint {
	c := C.gtk_entry_buffer_get_length(v.native())
	return uint(c)
}

// GetMaxLength() is a wrapper around gtk_entry_buffer_get_max_length().
func (v *EntryBuffer) GetMaxLength() int {
	c := C.gtk_entry_buffer_get_max_length(v.native())
	return int(c)
}

// SetMaxLength() is a wrapper around gtk_entry_buffer_set_max_length().
func (v *EntryBuffer) SetMaxLength(maxLength int) {
	C.gtk_entry_buffer_set_max_length(v.native(), C.gint(maxLength))
}

// InsertText() is a wrapper around gtk_entry_buffer_insert_text().
func (v *EntryBuffer) InsertText(position uint, text string) uint {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_entry_buffer_insert_text(v.native(), C.guint(position),
		(*C.gchar)(cstr), C.gint(len(text)))
	return uint(c)
}

// DeleteText() is a wrapper around gtk_entry_buffer_delete_text().
func (v *EntryBuffer) DeleteText(position uint, nChars int) uint {
	c := C.gtk_entry_buffer_delete_text(v.native(), C.guint(position),
		C.gint(nChars))
	return uint(c)
}

// EmitDeletedText() is a wrapper around gtk_entry_buffer_emit_deleted_text().
func (v *EntryBuffer) EmitDeletedText(pos, nChars uint) {
	C.gtk_entry_buffer_emit_deleted_text(v.native(), C.guint(pos),
		C.guint(nChars))
}

// EmitInsertedText() is a wrapper around gtk_entry_buffer_emit_inserted_text().
func (v *EntryBuffer) EmitInsertedText(pos uint, text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_buffer_emit_inserted_text(v.native(), C.guint(pos),
		(*C.gchar)(cstr), C.guint(len(text)))
}

/*
 * GtkEntryCompletion
 */

// EntryCompletion is a representation of GTK's GtkEntryCompletion.
type EntryCompletion struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkEntryCompletion.
func (v *EntryCompletion) native() *C.GtkEntryCompletion {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkEntryCompletion(p)
}

func marshalEntryCompletion(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapEntryCompletion(obj), nil
}

func wrapEntryCompletion(obj *glib.Object) *EntryCompletion {
	return &EntryCompletion{obj}
}

/*
 * GtkEventBox
 */

// EventBox is a representation of GTK's GtkEventBox.
type EventBox struct {
	Bin
}

// native returns a pointer to the underlying GtkEventBox.
func (v *EventBox) native() *C.GtkEventBox {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkEventBox(p)
}

func marshalEventBox(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapEventBox(obj), nil
}

func wrapEventBox(obj *glib.Object) *EventBox {
	return &EventBox{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// EventBoxNew is a wrapper around gtk_event_box_new().
func EventBoxNew() (*EventBox, error) {
	c := C.gtk_event_box_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapEventBox(obj), nil
}

// SetAboveChild is a wrapper around gtk_event_box_set_above_child().
func (v *EventBox) SetAboveChild(aboveChild bool) {
	C.gtk_event_box_set_above_child(v.native(), gbool(aboveChild))
}

// GetAboveChild is a wrapper around gtk_event_box_get_above_child().
func (v *EventBox) GetAboveChild() bool {
	c := C.gtk_event_box_get_above_child(v.native())
	return gobool(c)
}

// SetVisibleWindow is a wrapper around gtk_event_box_set_visible_window().
func (v *EventBox) SetVisibleWindow(visibleWindow bool) {
	C.gtk_event_box_set_visible_window(v.native(), gbool(visibleWindow))
}

// GetVisibleWindow is a wrapper around gtk_event_box_get_visible_window().
func (v *EventBox) GetVisibleWindow() bool {
	c := C.gtk_event_box_get_visible_window(v.native())
	return gobool(c)
}

/*
 * GtkExpander
 */

// Expander is a representation of GTK's GtkExpander.
type Expander struct {
	Bin
}

// native returns a pointer to the underlying GtkExpander.
func (v *Expander) native() *C.GtkExpander {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkExpander(p)
}

func marshalExpander(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapExpander(obj), nil
}

func wrapExpander(obj *glib.Object) *Expander {
	return &Expander{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// ExpanderNew is a wrapper around gtk_expander_new().
func ExpanderNew(label string) (*Expander, error) {
	var cstr *C.gchar
	if label != "" {
		cstr := C.CString(label)
		defer C.free(unsafe.Pointer(cstr))
	}
	c := C.gtk_expander_new((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapExpander(obj), nil
}

// SetExpanded is a wrapper around gtk_expander_set_expanded().
func (v *Expander) SetExpanded(expanded bool) {
	C.gtk_expander_set_expanded(v.native(), gbool(expanded))
}

// GetExpanded is a wrapper around gtk_expander_get_expanded().
func (v *Expander) GetExpanded() bool {
	c := C.gtk_expander_get_expanded(v.native())
	return gobool(c)
}

// SetLabel is a wrapper around gtk_expander_set_label().
func (v *Expander) SetLabel(label string) {
	var cstr *C.char
	if label != "" {
		cstr = C.CString(label)
		defer C.free(unsafe.Pointer(cstr))
	}
	C.gtk_expander_set_label(v.native(), (*C.gchar)(cstr))
}

// GetLabel is a wrapper around gtk_expander_get_label().
func (v *Expander) GetLabel() string {
	c := C.gtk_expander_get_label(v.native())
	return C.GoString((*C.char)(c))
}

// SetLabelWidget is a wrapper around gtk_expander_set_label_widget().
func (v *Expander) SetLabelWidget(widget IWidget) {
	C.gtk_expander_set_label_widget(v.native(), widget.toWidget())
}

/*
 * GtkFileChooser
 */

// FileChoser is a representation of GTK's GtkFileChooser GInterface.
type FileChooser struct {
	*glib.Object
}

// native returns a pointer to the underlying GObject as a GtkFileChooser.
func (v *FileChooser) native() *C.GtkFileChooser {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFileChooser(p)
}

func marshalFileChooser(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapFileChooser(obj), nil
}

func wrapFileChooser(obj *glib.Object) *FileChooser {
	return &FileChooser{obj}
}

// GetFilename is a wrapper around gtk_file_chooser_get_filename().
func (v *FileChooser) GetFilename() string {
	c := C.gtk_file_chooser_get_filename(v.native())
	s := C.GoString((*C.char)(c))
	defer C.g_free((C.gpointer)(c))
	return s
}

// SetCurrentFolder is a wrapper around gtk_file_chooser_set_current_folder().
func (v *FileChooser) SetCurrentFolder(folder string) bool {
	cstr := C.CString(folder)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_file_chooser_set_current_folder(v.native(), (*C.gchar)(cstr))
	return gobool(c)
}

// GetCurrentFolder is a wrapper around gtk_file_chooser_get_current_folder().
func (v *FileChooser) GetCurrentFolder() (string, error) {
	c := C.gtk_file_chooser_get_current_folder(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	defer C.free(unsafe.Pointer(c))
	return C.GoString((*C.char)(c)), nil
}

// SetPreviewWidget is a wrapper around gtk_file_chooser_set_preview_widget().
func (v *FileChooser) SetPreviewWidget(widget IWidget) {
	C.gtk_file_chooser_set_preview_widget(v.native(), widget.toWidget())
}

// SetPreviewWidgetActive is a wrapper around gtk_file_chooser_set_preview_widget_active().
func (v *FileChooser) SetPreviewWidgetActive(active bool) {
	C.gtk_file_chooser_set_preview_widget_active(v.native(), gbool(active))
}

// GetPreviewFilename is a wrapper around gtk_file_chooser_get_preview_filename().
func (v *FileChooser) GetPreviewFilename() string {
	c := C.gtk_file_chooser_get_preview_filename(v.native())
	defer C.free(unsafe.Pointer(c))
	return C.GoString(c)
}

// AddFilter is a wrapper around gtk_file_chooser_add_filter().
func (v *FileChooser) AddFilter(filter *FileFilter) {
	C.gtk_file_chooser_add_filter(v.native(), filter.native())
}

// GetURI is a wrapper around gtk_file_chooser_get_uri().
func (v *FileChooser) GetURI() string {
	c := C.gtk_file_chooser_get_uri(v.native())
	s := C.GoString((*C.char)(c))
	defer C.g_free((C.gpointer)(c))
	return s
}

// AddShortcutFolder is a wrapper around gtk_file_chooser_add_shortcut_folder().
func (v *FileChooser) AddShortcutFolder(folder string) bool {
	cstr := C.CString(folder)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_file_chooser_add_shortcut_folder(v.native(), cstr, nil)
	return gobool(c)
}

/*
 * GtkFileChooserButton
 */

// FileChooserButton is a representation of GTK's GtkFileChooserButton.
type FileChooserButton struct {
	Box

	// Interfaces
	FileChooser
}

// native returns a pointer to the underlying GtkFileChooserButton.
func (v *FileChooserButton) native() *C.GtkFileChooserButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFileChooserButton(p)
}

func marshalFileChooserButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapFileChooserButton(obj), nil
}

func wrapFileChooserButton(obj *glib.Object) *FileChooserButton {
	fc := wrapFileChooser(obj)
	return &FileChooserButton{Box{Container{Widget{glib.InitiallyUnowned{obj}}}}, *fc}
}

// FileChooserButtonNew is a wrapper around gtk_file_chooser_button_new().
func FileChooserButtonNew(title string, action FileChooserAction) (*FileChooserButton, error) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_file_chooser_button_new((*C.gchar)(cstr),
		(C.GtkFileChooserAction)(action))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapFileChooserButton(obj), nil
}

/*
 * GtkFileChooserDialog
 */

// FileChooserDialog is a representation of GTK's GtkFileChooserDialog.
type FileChooserDialog struct {
	Dialog

	// Interfaces
	FileChooser
}

// native returns a pointer to the underlying GtkFileChooserDialog.
func (v *FileChooserDialog) native() *C.GtkFileChooserDialog {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFileChooserDialog(p)
}

func marshalFileChooserDialog(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapFileChooserDialog(obj), nil
}

func wrapFileChooserDialog(obj *glib.Object) *FileChooserDialog {
	fc := wrapFileChooser(obj)
	return &FileChooserDialog{Dialog{Window{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}}, *fc}
}

// FileChooserDialogNewWith1Button is a wrapper around gtk_file_chooser_dialog_new() with one button.
func FileChooserDialogNewWith1Button(
	title string,
	parent *Window,
	action FileChooserAction,
	first_button_text string,
	first_button_id ResponseType) (*FileChooserDialog, error) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))
	c_first_button_text := C.CString(first_button_text)
	defer C.free(unsafe.Pointer(c_first_button_text))
	c := C.gtk_file_chooser_dialog_new_1(
		(*C.gchar)(c_title), parent.native(), C.GtkFileChooserAction(action),
		(*C.gchar)(c_first_button_text), C.int(first_button_id))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapFileChooserDialog(obj), nil
}

// FileChooserDialogNewWith2Buttons is a wrapper around gtk_file_chooser_dialog_new() with two buttons.
func FileChooserDialogNewWith2Buttons(
	title string,
	parent *Window,
	action FileChooserAction,
	first_button_text string,
	first_button_id ResponseType,
	second_button_text string,
	second_button_id ResponseType) (*FileChooserDialog, error) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))
	c_first_button_text := C.CString(first_button_text)
	defer C.free(unsafe.Pointer(c_first_button_text))
	c_second_button_text := C.CString(second_button_text)
	defer C.free(unsafe.Pointer(c_second_button_text))
	c := C.gtk_file_chooser_dialog_new_2(
		(*C.gchar)(c_title), parent.native(), C.GtkFileChooserAction(action),
		(*C.gchar)(c_first_button_text), C.int(first_button_id),
		(*C.gchar)(c_second_button_text), C.int(second_button_id))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapFileChooserDialog(obj), nil
}

/*
 * GtkFileChooserWidget
 */

// FileChooserWidget is a representation of GTK's GtkFileChooserWidget.
type FileChooserWidget struct {
	Box

	// Interfaces
	FileChooser
}

// native returns a pointer to the underlying GtkFileChooserWidget.
func (v *FileChooserWidget) native() *C.GtkFileChooserWidget {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFileChooserWidget(p)
}

func marshalFileChooserWidget(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapFileChooserWidget(obj), nil
}

func wrapFileChooserWidget(obj *glib.Object) *FileChooserWidget {
	fc := wrapFileChooser(obj)
	return &FileChooserWidget{Box{Container{Widget{glib.InitiallyUnowned{obj}}}}, *fc}
}

// FileChooserWidgetNew is a wrapper around gtk_gtk_file_chooser_widget_new().
func FileChooserWidgetNew(action FileChooserAction) (*FileChooserWidget, error) {
	c := C.gtk_file_chooser_widget_new((C.GtkFileChooserAction)(action))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapFileChooserWidget(obj), nil
}

/*
 * GtkFileFilter
 */

// FileChoser is a representation of GTK's GtkFileFilter GInterface.
type FileFilter struct {
	*glib.Object
}

// native returns a pointer to the underlying GObject as a GtkFileFilter.
func (v *FileFilter) native() *C.GtkFileFilter {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFileFilter(p)
}

func marshalFileFilter(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapFileFilter(obj), nil
}

func wrapFileFilter(obj *glib.Object) *FileFilter {
	return &FileFilter{obj}
}

// FileFilterNew is a wrapper around gtk_file_filter_new().
func FileFilterNew() (*FileFilter, error) {
	c := C.gtk_file_filter_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapFileFilter(obj), nil
}

// SetName is a wrapper around gtk_file_filter_set_name().
func (v *FileFilter) SetName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_file_filter_set_name(v.native(), (*C.gchar)(cstr))
}

// AddPattern is a wrapper around gtk_file_filter_add_pattern().
func (v *FileFilter) AddPattern(pattern string) {
	cstr := C.CString(pattern)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_file_filter_add_pattern(v.native(), (*C.gchar)(cstr))
}

// AddPixbufFormats is a wrapper around gtk_file_filter_add_pixbuf_formats().
func (v *FileFilter) AddPixbufFormats() {
	C.gtk_file_filter_add_pixbuf_formats(v.native())
}

/*
 * GtkFontButton
 */

// FontButton is a representation of GTK's GtkFontButton.
type FontButton struct {
	Button
}

// native returns a pointer to the underlying GtkFontButton.
func (v *FontButton) native() *C.GtkFontButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFontButton(p)
}

func marshalFontButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapFontButton(obj), nil
}

func wrapFontButton(obj *glib.Object) *FontButton {
	return &FontButton{Button{Bin{Container{Widget{
		glib.InitiallyUnowned{obj}}}}}}
}

// FontButtonNew is a wrapper around gtk_font_button_new().
func FontButtonNew() (*FontButton, error) {
	c := C.gtk_font_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapFontButton(obj), nil
}

// FontButtonNewWithFont is a wrapper around gtk_font_button_new_with_font().
func FontButtonNewWithFont(fontname string) (*FontButton, error) {
	cstr := C.CString(fontname)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_font_button_new_with_font((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapFontButton(obj), nil
}

// GetFontName is a wrapper around gtk_font_button_get_font_name().
func (v *FontButton) GetFontName() string {
	c := C.gtk_font_button_get_font_name(v.native())
	return C.GoString((*C.char)(c))
}

// SetFontName is a wrapper around gtk_font_button_set_font_name().
func (v *FontButton) SetFontName(fontname string) bool {
	cstr := C.CString(fontname)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_font_button_set_font_name(v.native(), (*C.gchar)(cstr))
	return gobool(c)
}

/*
 * GtkFrame
 */

// Frame is a representation of GTK's GtkFrame.
type Frame struct {
	Bin
}

// native returns a pointer to the underlying GtkFrame.
func (v *Frame) native() *C.GtkFrame {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFrame(p)
}

func marshalFrame(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapFrame(obj), nil
}

func wrapFrame(obj *glib.Object) *Frame {
	return &Frame{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// FrameNew is a wrapper around gtk_frame_new().
func FrameNew(label string) (*Frame, error) {
	var cstr *C.char
	if label != "" {
		cstr = C.CString(label)
		defer C.free(unsafe.Pointer(cstr))
	}
	c := C.gtk_frame_new((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapFrame(obj), nil
}

// SetLabel is a wrapper around gtk_frame_set_label().
func (v *Frame) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_frame_set_label(v.native(), (*C.gchar)(cstr))
}

// SetLabelWidget is a wrapper around gtk_frame_set_label_widget().
func (v *Frame) SetLabelWidget(labelWidget IWidget) {
	C.gtk_frame_set_label_widget(v.native(), labelWidget.toWidget())
}

// SetLabelAlign is a wrapper around gtk_frame_set_label_align().
func (v *Frame) SetLabelAlign(xAlign, yAlign float32) {
	C.gtk_frame_set_label_align(v.native(), C.gfloat(xAlign),
		C.gfloat(yAlign))
}

// SetShadowType is a wrapper around gtk_frame_set_shadow_type().
func (v *Frame) SetShadowType(t ShadowType) {
	C.gtk_frame_set_shadow_type(v.native(), C.GtkShadowType(t))
}

// GetLabel is a wrapper around gtk_frame_get_label().
func (v *Frame) GetLabel() string {
	c := C.gtk_frame_get_label(v.native())
	return C.GoString((*C.char)(c))
}

// GetLabelAlign is a wrapper around gtk_frame_get_label_align().
func (v *Frame) GetLabelAlign() (xAlign, yAlign float32) {
	var x, y C.gfloat
	C.gtk_frame_get_label_align(v.native(), &x, &y)
	return float32(x), float32(y)
}

// GetLabelWidget is a wrapper around gtk_frame_get_label_widget().
func (v *Frame) GetLabelWidget() (*Widget, error) {
	c := C.gtk_frame_get_label_widget(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// GetShadowType is a wrapper around gtk_frame_get_shadow_type().
func (v *Frame) GetShadowType() ShadowType {
	c := C.gtk_frame_get_shadow_type(v.native())
	return ShadowType(c)
}

/*
 * GtkGrid
 */

// Grid is a representation of GTK's GtkGrid.
type Grid struct {
	Container

	// Interfaces
	Orientable
}

// native returns a pointer to the underlying GtkGrid.
func (v *Grid) native() *C.GtkGrid {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkGrid(p)
}

func (v *Grid) toOrientable() *C.GtkOrientable {
	if v == nil {
		return nil
	}
	return C.toGtkOrientable(unsafe.Pointer(v.GObject))
}

func marshalGrid(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapGrid(obj), nil
}

func wrapGrid(obj *glib.Object) *Grid {
	o := wrapOrientable(obj)
	return &Grid{Container{Widget{glib.InitiallyUnowned{obj}}}, *o}
}

// GridNew() is a wrapper around gtk_grid_new().
func GridNew() (*Grid, error) {
	c := C.gtk_grid_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapGrid(obj), nil
}

// Attach() is a wrapper around gtk_grid_attach().
func (v *Grid) Attach(child IWidget, left, top, width, height int) {
	C.gtk_grid_attach(v.native(), child.toWidget(), C.gint(left),
		C.gint(top), C.gint(width), C.gint(height))
}

// AttachNextTo() is a wrapper around gtk_grid_attach_next_to().
func (v *Grid) AttachNextTo(child, sibling IWidget, side PositionType, width, height int) {
	C.gtk_grid_attach_next_to(v.native(), child.toWidget(),
		sibling.toWidget(), C.GtkPositionType(side), C.gint(width),
		C.gint(height))
}

// GetChildAt() is a wrapper around gtk_grid_get_child_at().
func (v *Grid) GetChildAt(left, top int) (*Widget, error) {
	c := C.gtk_grid_get_child_at(v.native(), C.gint(left), C.gint(top))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// InsertRow() is a wrapper around gtk_grid_insert_row().
func (v *Grid) InsertRow(position int) {
	C.gtk_grid_insert_row(v.native(), C.gint(position))
}

// InsertColumn() is a wrapper around gtk_grid_insert_column().
func (v *Grid) InsertColumn(position int) {
	C.gtk_grid_insert_column(v.native(), C.gint(position))
}

// InsertNextTo() is a wrapper around gtk_grid_insert_next_to()
func (v *Grid) InsertNextTo(sibling IWidget, side PositionType) {
	C.gtk_grid_insert_next_to(v.native(), sibling.toWidget(),
		C.GtkPositionType(side))
}

// SetRowHomogeneous() is a wrapper around gtk_grid_set_row_homogeneous().
func (v *Grid) SetRowHomogeneous(homogeneous bool) {
	C.gtk_grid_set_row_homogeneous(v.native(), gbool(homogeneous))
}

// GetRowHomogeneous() is a wrapper around gtk_grid_get_row_homogeneous().
func (v *Grid) GetRowHomogeneous() bool {
	c := C.gtk_grid_get_row_homogeneous(v.native())
	return gobool(c)
}

// SetRowSpacing() is a wrapper around gtk_grid_set_row_spacing().
func (v *Grid) SetRowSpacing(spacing uint) {
	C.gtk_grid_set_row_spacing(v.native(), C.guint(spacing))
}

// GetRowSpacing() is a wrapper around gtk_grid_get_row_spacing().
func (v *Grid) GetRowSpacing() uint {
	c := C.gtk_grid_get_row_spacing(v.native())
	return uint(c)
}

// SetColumnHomogeneous() is a wrapper around gtk_grid_set_column_homogeneous().
func (v *Grid) SetColumnHomogeneous(homogeneous bool) {
	C.gtk_grid_set_column_homogeneous(v.native(), gbool(homogeneous))
}

// GetColumnHomogeneous() is a wrapper around gtk_grid_get_column_homogeneous().
func (v *Grid) GetColumnHomogeneous() bool {
	c := C.gtk_grid_get_column_homogeneous(v.native())
	return gobool(c)
}

// SetColumnSpacing() is a wrapper around gtk_grid_set_column_spacing().
func (v *Grid) SetColumnSpacing(spacing uint) {
	C.gtk_grid_set_column_spacing(v.native(), C.guint(spacing))
}

// GetColumnSpacing() is a wrapper around gtk_grid_get_column_spacing().
func (v *Grid) GetColumnSpacing() uint {
	c := C.gtk_grid_get_column_spacing(v.native())
	return uint(c)
}

/*
 * GtkIconTheme
 */

// IconTheme is a representation of GTK's GtkIconTheme
type IconTheme struct {
	Theme *C.GtkIconTheme
}

// IconThemeGetDefault is a wrapper around gtk_icon_theme_get_default().
func IconThemeGetDefault() (*IconTheme, error) {
	c := C.gtk_icon_theme_get_default()
	if c == nil {
		return nil, nilPtrErr
	}
	return &IconTheme{c}, nil
}

// IconThemeGetForScreen is a wrapper around gtk_icon_theme_get_for_screen().
func IconThemeGetForScreen(screen gdk.Screen) (*IconTheme, error) {
	cScreen := (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	c := C.gtk_icon_theme_get_for_screen(cScreen)
	if c == nil {
		return nil, nilPtrErr
	}
	return &IconTheme{c}, nil
}

// LoadIcon is a wrapper around gtk_icon_theme_load_icon().
func (v *IconTheme) LoadIcon(iconName string, size int, flags IconLookupFlags) (*gdk.Pixbuf, error) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	c := C.gtk_icon_theme_load_icon(v.Theme, (*C.gchar)(cstr), C.gint(size), C.GtkIconLookupFlags(flags), &err)
	if c == nil {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	pb := &gdk.Pixbuf{obj}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return pb, nil
}

/*
 * GtkIconView
 */

// IconView is a representation of GTK's GtkIconView.
type IconView struct {
	Container
}

// native returns a pointer to the underlying GtkIconView.
func (v *IconView) native() *C.GtkIconView {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkIconView(p)
}

func marshalIconView(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapIconView(obj), nil
}

func wrapIconView(obj *glib.Object) *IconView {
	return &IconView{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// IconViewNew is a wrapper around gtk_icon_view_new().
func IconViewNew() (*IconView, error) {
	c := C.gtk_icon_view_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	iv := wrapIconView(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return iv, nil
}

// IconViewNewWithModel is a wrapper around gtk_icon_view_new_with_model().
func IconViewNewWithModel(model ITreeModel) (*IconView, error) {
	c := C.gtk_icon_view_new_with_model(model.toTreeModel())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapIconView(obj), nil
}

// GetModel is a wrapper around gtk_icon_view_get_model().
func (v *IconView) GetModel() (*TreeModel, error) {
	c := C.gtk_icon_view_get_model(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapTreeModel(obj), nil
}

// SetModel is a wrapper around gtk_icon_view_set_model().
func (v *IconView) SetModel(model ITreeModel) {
	C.gtk_icon_view_set_model(v.native(), model.toTreeModel())
}

// SelectPath is a wrapper around gtk_icon_view_select_path().
func (v *IconView) SelectPath(path *TreePath) {
	C.gtk_icon_view_select_path(v.native(), path.native())
}

// ScrollToPath is a wrapper around gtk_icon_view_scroll_to_path().
func (v *IconView) ScrollToPath(path *TreePath, useAlign bool, rowAlign, colAlign float64) {
	C.gtk_icon_view_scroll_to_path(v.native(), path.native(), gbool(useAlign),
		C.gfloat(rowAlign), C.gfloat(colAlign))
}

/*
 * GtkImage
 */

// Image is a representation of GTK's GtkImage.
type Image struct {
	Widget
}

// native returns a pointer to the underlying GtkImage.
func (v *Image) native() *C.GtkImage {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkImage(p)
}

func marshalImage(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapImage(obj), nil
}

func wrapImage(obj *glib.Object) *Image {
	return &Image{Widget{glib.InitiallyUnowned{obj}}}
}

// ImageNew() is a wrapper around gtk_image_new().
func ImageNew() (*Image, error) {
	c := C.gtk_image_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapImage(obj), nil
}

// ImageNewFromFile() is a wrapper around gtk_image_new_from_file().
func ImageNewFromFile(filename string) (*Image, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_image_new_from_file((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapImage(obj), nil
}

// ImageNewFromResource() is a wrapper around gtk_image_new_from_resource().
func ImageNewFromResource(resourcePath string) (*Image, error) {
	cstr := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_image_new_from_resource((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapImage(obj), nil
}

// ImageNewFromPixbuf is a wrapper around gtk_image_new_from_pixbuf().
func ImageNewFromPixbuf(pixbuf *gdk.Pixbuf) (*Image, error) {
	ptr := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	c := C.gtk_image_new_from_pixbuf(ptr)
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapImage(obj), nil
}

// TODO(jrick) GtkIconSet
/*
func ImageNewFromIconSet() {
}
*/

// TODO(jrick) GdkPixbufAnimation
/*
func ImageNewFromAnimation() {
}
*/

// ImageNewFromIconName() is a wrapper around gtk_image_new_from_icon_name().
func ImageNewFromIconName(iconName string, size IconSize) (*Image, error) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_image_new_from_icon_name((*C.gchar)(cstr),
		C.GtkIconSize(size))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapImage(obj), nil
}

// TODO(jrick) GIcon
/*
func ImageNewFromGIcon() {
}
*/

// Clear() is a wrapper around gtk_image_clear().
func (v *Image) Clear() {
	C.gtk_image_clear(v.native())
}

// SetFromFile() is a wrapper around gtk_image_set_from_file().
func (v *Image) SetFromFile(filename string) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_image_set_from_file(v.native(), (*C.gchar)(cstr))
}

// SetFromResource() is a wrapper around gtk_image_set_from_resource().
func (v *Image) SetFromResource(resourcePath string) {
	cstr := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_image_set_from_resource(v.native(), (*C.gchar)(cstr))
}

// SetFromFixbuf is a wrapper around gtk_image_set_from_pixbuf().
func (v *Image) SetFromPixbuf(pixbuf *gdk.Pixbuf) {
	pbptr := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	C.gtk_image_set_from_pixbuf(v.native(), pbptr)
}

// TODO(jrick) GtkIconSet
/*
func (v *Image) SetFromIconSet() {
}
*/

// TODO(jrick) GdkPixbufAnimation
/*
func (v *Image) SetFromAnimation() {
}
*/

// SetFromIconName() is a wrapper around gtk_image_set_from_icon_name().
func (v *Image) SetFromIconName(iconName string, size IconSize) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_image_set_from_icon_name(v.native(), (*C.gchar)(cstr),
		C.GtkIconSize(size))
}

// TODO(jrick) GIcon
/*
func (v *Image) SetFromGIcon() {
}
*/

// SetPixelSize() is a wrapper around gtk_image_set_pixel_size().
func (v *Image) SetPixelSize(pixelSize int) {
	C.gtk_image_set_pixel_size(v.native(), C.gint(pixelSize))
}

// GetStorageType() is a wrapper around gtk_image_get_storage_type().
func (v *Image) GetStorageType() ImageType {
	c := C.gtk_image_get_storage_type(v.native())
	return ImageType(c)
}

// GetPixbuf() is a wrapper around gtk_image_get_pixbuf().
func (v *Image) GetPixbuf() *gdk.Pixbuf {
	c := C.gtk_image_get_pixbuf(v.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	pb := &gdk.Pixbuf{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return pb
}

// TODO(jrick) GtkIconSet
/*
func (v *Image) GetIconSet() {
}
*/

// TODO(jrick) GdkPixbufAnimation
/*
func (v *Image) GetAnimation() {
}
*/

// GetIconName() is a wrapper around gtk_image_get_icon_name().
func (v *Image) GetIconName() (string, IconSize) {
	var iconName *C.gchar
	var size C.GtkIconSize
	C.gtk_image_get_icon_name(v.native(), &iconName, &size)
	return C.GoString((*C.char)(iconName)), IconSize(size)
}

// TODO(jrick) GIcon
/*
func (v *Image) GetGIcon() {
}
*/

// GetPixelSize() is a wrapper around gtk_image_get_pixel_size().
func (v *Image) GetPixelSize() int {
	c := C.gtk_image_get_pixel_size(v.native())
	return int(c)
}

/*
 * GtkLabel
 */

// Label is a representation of GTK's GtkLabel.
type Label struct {
	Widget
}

// native returns a pointer to the underlying GtkLabel.
func (v *Label) native() *C.GtkLabel {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkLabel(p)
}

func marshalLabel(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapLabel(obj), nil
}

func wrapLabel(obj *glib.Object) *Label {
	return &Label{Widget{glib.InitiallyUnowned{obj}}}
}

// LabelNew is a wrapper around gtk_label_new().
func LabelNew(str string) (*Label, error) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_label_new((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapLabel(obj), nil
}

// SetText is a wrapper around gtk_label_set_text().
func (v *Label) SetText(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_text(v.native(), (*C.gchar)(cstr))
}

// SetMarkup is a wrapper around gtk_label_set_markup().
func (v *Label) SetMarkup(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_markup(v.native(), (*C.gchar)(cstr))
}

// SetMarkupWithMnemonic is a wrapper around
// gtk_label_set_markup_with_mnemonic().
func (v *Label) SetMarkupWithMnemonic(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_markup_with_mnemonic(v.native(), (*C.gchar)(cstr))
}

// SetPattern is a wrapper around gtk_label_set_pattern().
func (v *Label) SetPattern(patern string) {
	cstr := C.CString(patern)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_pattern(v.native(), (*C.gchar)(cstr))
}

// SetJustify is a wrapper around gtk_label_set_justify().
func (v *Label) SetJustify(jtype Justification) {
	C.gtk_label_set_justify(v.native(), C.GtkJustification(jtype))
}

// SetEllipsize is a wrapper around gtk_label_set_ellipsize().
func (v *Label) SetEllipsize(mode pango.EllipsizeMode) {
	C.gtk_label_set_ellipsize(v.native(), C.PangoEllipsizeMode(mode))
}

// GetWidthChars is a wrapper around gtk_label_get_width_chars().
func (v *Label) GetWidthChars() int {
	c := C.gtk_label_get_width_chars(v.native())
	return int(c)
}

// SetWidthChars is a wrapper around gtk_label_set_width_chars().
func (v *Label) SetWidthChars(nChars int) {
	C.gtk_label_set_width_chars(v.native(), C.gint(nChars))
}

// GetMaxWidthChars is a wrapper around gtk_label_get_max_width_chars().
func (v *Label) GetMaxWidthChars() int {
	c := C.gtk_label_get_max_width_chars(v.native())
	return int(c)
}

// SetMaxWidthChars is a wrapper around gtk_label_set_max_width_chars().
func (v *Label) SetMaxWidthChars(nChars int) {
	C.gtk_label_set_max_width_chars(v.native(), C.gint(nChars))
}

// GetLineWrap is a wrapper around gtk_label_get_line_wrap().
func (v *Label) GetLineWrap() bool {
	c := C.gtk_label_get_line_wrap(v.native())
	return gobool(c)
}

// SetLineWrap is a wrapper around gtk_label_set_line_wrap().
func (v *Label) SetLineWrap(wrap bool) {
	C.gtk_label_set_line_wrap(v.native(), gbool(wrap))
}

// SetLineWrapMode is a wrapper around gtk_label_set_line_wrap_mode().
func (v *Label) SetLineWrapMode(wrapMode pango.WrapMode) {
	C.gtk_label_set_line_wrap_mode(v.native(), C.PangoWrapMode(wrapMode))
}

// GetSelectable is a wrapper around gtk_label_get_selectable().
func (v *Label) GetSelectable() bool {
	c := C.gtk_label_get_selectable(v.native())
	return gobool(c)
}

// GetText is a wrapper around gtk_label_get_text().
func (v *Label) GetText() (string, error) {
	c := C.gtk_label_get_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// GetJustify is a wrapper around gtk_label_get_justify().
func (v *Label) GetJustify() Justification {
	c := C.gtk_label_get_justify(v.native())
	return Justification(c)
}

// GetEllipsize is a wrapper around gtk_label_get_ellipsize().
func (v *Label) GetEllipsize() pango.EllipsizeMode {
	c := C.gtk_label_get_ellipsize(v.native())
	return pango.EllipsizeMode(c)
}

// GetCurrentUri is a wrapper around gtk_label_get_current_uri().
func (v *Label) GetCurrentUri() string {
	c := C.gtk_label_get_current_uri(v.native())
	return C.GoString((*C.char)(c))
}

// GetTrackVisitedLinks is a wrapper around gtk_label_get_track_visited_links().
func (v *Label) GetTrackVisitedLinks() bool {
	c := C.gtk_label_get_track_visited_links(v.native())
	return gobool(c)
}

// SetTrackVisitedLinks is a wrapper around gtk_label_set_track_visited_links().
func (v *Label) SetTrackVisitedLinks(trackLinks bool) {
	C.gtk_label_set_track_visited_links(v.native(), gbool(trackLinks))
}

// GetAngle is a wrapper around gtk_label_get_angle().
func (v *Label) GetAngle() float64 {
	c := C.gtk_label_get_angle(v.native())
	return float64(c)
}

// SetAngle is a wrapper around gtk_label_set_angle().
func (v *Label) SetAngle(angle float64) {
	C.gtk_label_set_angle(v.native(), C.gdouble(angle))
}

// GetSelectionBounds is a wrapper around gtk_label_get_selection_bounds().
func (v *Label) GetSelectionBounds() (start, end int, nonEmpty bool) {
	var cstart, cend C.gint
	c := C.gtk_label_get_selection_bounds(v.native(), &cstart, &cend)
	return int(cstart), int(cend), gobool(c)
}

// GetSingleLineMode is a wrapper around gtk_label_get_single_line_mode().
func (v *Label) GetSingleLineMode() bool {
	c := C.gtk_label_get_single_line_mode(v.native())
	return gobool(c)
}

// SetSingleLineMode is a wrapper around gtk_label_set_single_line_mode().
func (v *Label) SetSingleLineMode(mode bool) {
	C.gtk_label_set_single_line_mode(v.native(), gbool(mode))
}

// GetUseMarkup is a wrapper around gtk_label_get_use_markup().
func (v *Label) GetUseMarkup() bool {
	c := C.gtk_label_get_use_markup(v.native())
	return gobool(c)
}

// SetUseMarkup is a wrapper around gtk_label_set_use_markup().
func (v *Label) SetUseMarkup(use bool) {
	C.gtk_label_set_use_markup(v.native(), gbool(use))
}

// GetUseUnderline is a wrapper around gtk_label_get_use_underline().
func (v *Label) GetUseUnderline() bool {
	c := C.gtk_label_get_use_underline(v.native())
	return gobool(c)
}

// SetUseUnderline is a wrapper around gtk_label_set_use_underline().
func (v *Label) SetUseUnderline(use bool) {
	C.gtk_label_set_use_underline(v.native(), gbool(use))
}

// LabelNewWithMnemonic is a wrapper around gtk_label_new_with_mnemonic().
func LabelNewWithMnemonic(str string) (*Label, error) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_label_new_with_mnemonic((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapLabel(obj), nil
}

// SelectRegion is a wrapper around gtk_label_select_region().
func (v *Label) SelectRegion(startOffset, endOffset int) {
	C.gtk_label_select_region(v.native(), C.gint(startOffset),
		C.gint(endOffset))
}

// SetSelectable is a wrapper around gtk_label_set_selectable().
func (v *Label) SetSelectable(setting bool) {
	C.gtk_label_set_selectable(v.native(), gbool(setting))
}

// SetLabel is a wrapper around gtk_label_set_label().
func (v *Label) SetLabel(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_label(v.native(), (*C.gchar)(cstr))
}

// added by terrak
/*
 * GtkLayout
 */

// Layout is a representation of GTK's GtkLayout.
type Layout struct {
	Container
}

// native returns a pointer to the underlying GtkDrawingArea.
func (v *Layout) native() *C.GtkLayout {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkLayout(p)
}

func marshalLayout(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapLayout(obj), nil
}

func wrapLayout(obj *glib.Object) *Layout {
	return &Layout{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// LayoutNew is a wrapper around gtk_layout_new().
func LayoutNew(hadjustment, vadjustment *Adjustment) (*Layout, error) {
	c := C.gtk_layout_new(hadjustment.native(), vadjustment.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapLayout(obj), nil
}

// Layout.Put is a wrapper around gtk_layout_put().
func (v *Layout) Put(w IWidget, x, y int) {
	C.gtk_layout_put(v.native(), w.toWidget(), C.gint(x), C.gint(y))
}

// Layout.Move is a wrapper around gtk_layout_move().
func (v *Layout) Move(w IWidget, x, y int) {
	C.gtk_layout_move(v.native(), w.toWidget(), C.gint(x), C.gint(y))
}

// Layout.SetSize is a wrapper around gtk_layout_set_size
func (v *Layout) SetSize(width, height uint) {
	C.gtk_layout_set_size(v.native(), C.guint(width), C.guint(height))
}

// Layout.GetSize is a wrapper around gtk_layout_get_size
func (v *Layout) GetSize() (width, height uint) {
	var w, h C.guint
	C.gtk_layout_get_size(v.native(), &w, &h)
	return uint(w), uint(h)
}

/*
 * GtkLinkButton
 */

// LinkButton is a representation of GTK's GtkLinkButton.
type LinkButton struct {
	Button
}

// native returns a pointer to the underlying GtkLinkButton.
func (v *LinkButton) native() *C.GtkLinkButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkLinkButton(p)
}

func marshalLinkButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapLinkButton(obj), nil
}

func wrapLinkButton(obj *glib.Object) *LinkButton {
	return &LinkButton{Button{Bin{Container{Widget{
		glib.InitiallyUnowned{obj}}}}}}
}

// LinkButtonNew is a wrapper around gtk_link_button_new().
func LinkButtonNew(label string) (*LinkButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_link_button_new((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	lb := wrapLinkButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return lb, nil
}

// LinkButtonNewWithLabel is a wrapper around gtk_link_button_new_with_label().
func LinkButtonNewWithLabel(uri, label string) (*LinkButton, error) {
	curi := C.CString(uri)
	defer C.free(unsafe.Pointer(curi))
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	c := C.gtk_link_button_new_with_label((*C.gchar)(curi), (*C.gchar)(clabel))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	lb := wrapLinkButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return lb, nil
}

// GetUri is a wrapper around gtk_link_button_get_uri().
func (v *LinkButton) GetUri() string {
	c := C.gtk_link_button_get_uri(v.native())
	return C.GoString((*C.char)(c))
}

// SetUri is a wrapper around gtk_link_button_set_uri().
func (v *LinkButton) SetUri(uri string) {
	cstr := C.CString(uri)
	C.gtk_link_button_set_uri(v.native(), (*C.gchar)(cstr))
}

/*
 * GtkListStore
 */

// ListStore is a representation of GTK's GtkListStore.
type ListStore struct {
	*glib.Object

	// Interfaces
	TreeModel
}

// native returns a pointer to the underlying GtkListStore.
func (v *ListStore) native() *C.GtkListStore {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkListStore(p)
}

func marshalListStore(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapListStore(obj), nil
}

func wrapListStore(obj *glib.Object) *ListStore {
	tm := wrapTreeModel(obj)
	return &ListStore{obj, *tm}
}

func (v *ListStore) toTreeModel() *C.GtkTreeModel {
	if v == nil {
		return nil
	}
	return C.toGtkTreeModel(unsafe.Pointer(v.GObject))
}

// ListStoreNew is a wrapper around gtk_list_store_newv().
func ListStoreNew(types ...glib.Type) (*ListStore, error) {
	gtypes := C.alloc_types(C.int(len(types)))
	for n, val := range types {
		C.set_type(gtypes, C.int(n), C.GType(val))
	}
	defer C.g_free(C.gpointer(gtypes))
	c := C.gtk_list_store_newv(C.gint(len(types)), gtypes)
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	ls := wrapListStore(obj)
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return ls, nil
}

// Remove is a wrapper around gtk_list_store_remove().
func (v *ListStore) Remove(iter *TreeIter) bool {
	c := C.gtk_list_store_remove(v.native(), iter.native())
	return gobool(c)
}

// TODO(jrick)
/*
func (v *ListStore) SetColumnTypes(types ...glib.Type) {
}
*/

// Set() is a wrapper around gtk_list_store_set_value() but provides
// a function similar to gtk_list_store_set() in that multiple columns
// may be set by one call.  The length of columns and values slices must
// match, or Set() will return a non-nil error.
//
// As an example, a call to:
//  store.Set(iter, []int{0, 1}, []interface{}{"Foo", "Bar"})
// is functionally equivalent to calling the native C GTK function:
//  gtk_list_store_set(store, iter, 0, "Foo", 1, "Bar", -1);
func (v *ListStore) Set(iter *TreeIter, columns []int, values []interface{}) error {
	if len(columns) != len(values) {
		return errors.New("columns and values lengths do not match")
	}
	for i, val := range values {
		v.SetValue(iter, columns[i], val)
	}
	return nil
}

// SetValue is a wrapper around gtk_list_store_set_value().
func (v *ListStore) SetValue(iter *TreeIter, column int, value interface{}) error {
	switch value.(type) {
	case *gdk.Pixbuf:
		pix := value.(*gdk.Pixbuf)
		C._gtk_list_store_set(v.native(), iter.native(), C.gint(column), unsafe.Pointer(pix.Native()))

	default:
		gv, err := glib.GValue(value)
		if err != nil {
			return err
		}

		C.gtk_list_store_set_value(v.native(), iter.native(),
			C.gint(column),
			(*C.GValue)(unsafe.Pointer(gv.Native())))
	}

	return nil
}

// func (v *ListStore) Model(model ITreeModel) {
// 	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(model.toTreeModel()))}
//	v.TreeModel = *wrapTreeModel(obj)
//}

// SetSortColumnId() is a wrapper around gtk_tree_sortable_set_sort_column_id().
func (v *ListStore) SetSortColumnId(column int, order SortType) {
	sort := C.toGtkTreeSortable(unsafe.Pointer(v.Native()))
	C.gtk_tree_sortable_set_sort_column_id(sort, C.gint(column), C.GtkSortType(order))
}

func (v *ListStore) SetCols(iter *TreeIter, cols Cols) error {
	for key, value := range cols {
		err := v.SetValue(iter, key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

// Convenient map for Columns and values (See ListStore, TreeStore)
type Cols map[int]interface{}

// TODO(jrick)
/*
func (v *ListStore) InsertWithValues(iter *TreeIter, position int, columns []int, values []glib.Value) {
		var ccolumns *C.gint
		var cvalues *C.GValue

		C.gtk_list_store_insert_with_values(v.native(), iter.native(),
			C.gint(position), columns, values, C.gint(len(values)))
}
*/

// InsertBefore() is a wrapper around gtk_list_store_insert_before().
func (v *ListStore) InsertBefore(sibling *TreeIter) *TreeIter {
	var ti C.GtkTreeIter
	C.gtk_list_store_insert_before(v.native(), &ti, sibling.native())
	iter := &TreeIter{ti}
	return iter
}

// InsertAfter() is a wrapper around gtk_list_store_insert_after().
func (v *ListStore) InsertAfter(sibling *TreeIter) *TreeIter {
	var ti C.GtkTreeIter
	C.gtk_list_store_insert_after(v.native(), &ti, sibling.native())
	iter := &TreeIter{ti}
	return iter
}

// Prepend() is a wrapper around gtk_list_store_prepend().
func (v *ListStore) Prepend() *TreeIter {
	var ti C.GtkTreeIter
	C.gtk_list_store_prepend(v.native(), &ti)
	iter := &TreeIter{ti}
	return iter
}

// Append() is a wrapper around gtk_list_store_append().
func (v *ListStore) Append() *TreeIter {
	var ti C.GtkTreeIter
	C.gtk_list_store_append(v.native(), &ti)
	iter := &TreeIter{ti}
	return iter
}

// Clear() is a wrapper around gtk_list_store_clear().
func (v *ListStore) Clear() {
	C.gtk_list_store_clear(v.native())
}

// IterIsValid() is a wrapper around gtk_list_store_iter_is_valid().
func (v *ListStore) IterIsValid(iter *TreeIter) bool {
	c := C.gtk_list_store_iter_is_valid(v.native(), iter.native())
	return gobool(c)
}

// TODO(jrick)
/*
func (v *ListStore) Reorder(newOrder []int) {
}
*/

// Swap() is a wrapper around gtk_list_store_swap().
func (v *ListStore) Swap(a, b *TreeIter) {
	C.gtk_list_store_swap(v.native(), a.native(), b.native())
}

// MoveBefore() is a wrapper around gtk_list_store_move_before().
func (v *ListStore) MoveBefore(iter, position *TreeIter) {
	C.gtk_list_store_move_before(v.native(), iter.native(),
		position.native())
}

// MoveAfter() is a wrapper around gtk_list_store_move_after().
func (v *ListStore) MoveAfter(iter, position *TreeIter) {
	C.gtk_list_store_move_after(v.native(), iter.native(),
		position.native())
}

/*
 * GtkMenu
 */

// Menu is a representation of GTK's GtkMenu.
type Menu struct {
	MenuShell
}

// IMenu is an interface type implemented by all structs embedding
// a Menu.  It is meant to be used as an argument type for wrapper
// functions that wrap around a C GTK function taking a
// GtkMenu.
type IMenu interface {
	toMenu() *C.GtkMenu
	toWidget() *C.GtkWidget
}

// native() returns a pointer to the underlying GtkMenu.
func (v *Menu) native() *C.GtkMenu {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenu(p)
}

func (v *Menu) toMenu() *C.GtkMenu {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalMenu(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapMenu(obj), nil
}

func wrapMenu(obj *glib.Object) *Menu {
	return &Menu{MenuShell{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// MenuNew() is a wrapper around gtk_menu_new().
func MenuNew() (*Menu, error) {
	c := C.gtk_menu_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	m := wrapMenu(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return m, nil
}

// PopupAtMouse() is a wrapper for gtk_menu_popup(), without the option for a custom positioning function.
func (v *Menu) PopupAtMouseCursor(parentMenuShell IMenu, parentMenuItem IMenuItem, button int, activateTime uint32) {
	wshell := nullableWidget(parentMenuShell)
	witem := nullableWidget(parentMenuItem)

	C.gtk_menu_popup(v.native(),
		wshell,
		witem,
		nil,
		nil,
		C.guint(button),
		C.guint32(activateTime))
}

// Popdown() is a wrapper around gtk_menu_popdown().
func (v *Menu) Popdown() {
	C.gtk_menu_popdown(v.native())
}

// ReorderChild() is a wrapper around gtk_menu_reorder_child().
func (v *Menu) ReorderChild(child IWidget, position int) {
	C.gtk_menu_reorder_child(v.native(), child.toWidget(), C.gint(position))
}

/*
 * GtkMenuBar
 */

// MenuBar is a representation of GTK's GtkMenuBar.
type MenuBar struct {
	MenuShell
}

// native() returns a pointer to the underlying GtkMenuBar.
func (v *MenuBar) native() *C.GtkMenuBar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenuBar(p)
}

func marshalMenuBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapMenuBar(obj), nil
}

func wrapMenuBar(obj *glib.Object) *MenuBar {
	return &MenuBar{MenuShell{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// MenuBarNew() is a wrapper around gtk_menu_bar_new().
func MenuBarNew() (*MenuBar, error) {
	c := C.gtk_menu_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	m := wrapMenuBar(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return m, nil
}

/*
 * GtkMenuButton
 */

// MenuButton is a representation of GTK's GtkMenuButton.
type MenuButton struct {
	ToggleButton
}

// native returns a pointer to the underlying GtkMenuButton.
func (v *MenuButton) native() *C.GtkMenuButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenuButton(p)
}

func marshalMenuButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapMenuButton(obj), nil
}

func wrapMenuButton(obj *glib.Object) *MenuButton {
	return &MenuButton{ToggleButton{Button{Bin{Container{Widget{
		glib.InitiallyUnowned{obj}}}}}}}
}

// MenuButtonNew is a wrapper around gtk_menu_button_new().
func MenuButtonNew() (*MenuButton, error) {
	c := C.gtk_menu_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	m := wrapMenuButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return m, nil
}

// SetPopup is a wrapper around gtk_menu_button_set_popup().
func (v *MenuButton) SetPopup(menu IMenu) {
	C.gtk_menu_button_set_popup(v.native(), menu.toWidget())
}

// GetPopup is a wrapper around gtk_menu_button_get_popup().
func (v *MenuButton) GetPopup() *Menu {
	c := C.gtk_menu_button_get_popup(v.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	m := wrapMenu(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return m
}

// TODO: gtk_menu_button_set_menu_model
// TODO: gtk_menu_button_get_menu_model

// SetDirection is a wrapper around gtk_menu_button_set_direction().
func (v *MenuButton) SetDirection(direction ArrowType) {
	C.gtk_menu_button_set_direction(v.native(), C.GtkArrowType(direction))
}

// GetDirection is a wrapper around gtk_menu_button_get_direction().
func (v *MenuButton) GetDirection() ArrowType {
	c := C.gtk_menu_button_get_direction(v.native())
	return ArrowType(c)
}

// SetAlignWidget is a wrapper around gtk_menu_button_set_align_widget().
func (v *MenuButton) SetAlignWidget(alignWidget IWidget) {
	C.gtk_menu_button_set_align_widget(v.native(), alignWidget.toWidget())
}

// GetAlignWidget is a wrapper around gtk_menu_button_get_align_widget().
func (v *MenuButton) GetAlignWidget() *Widget {
	c := C.gtk_menu_button_get_align_widget(v.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w
}

/*
 * GtkMenuItem
 */

// MenuItem is a representation of GTK's GtkMenuItem.
type MenuItem struct {
	Bin
}

// IMenuItem is an interface type implemented by all structs
// embedding a MenuItem.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkMenuItem.
type IMenuItem interface {
	toMenuItem() *C.GtkMenuItem
	toWidget() *C.GtkWidget
}

// native returns a pointer to the underlying GtkMenuItem.
func (v *MenuItem) native() *C.GtkMenuItem {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenuItem(p)
}

func (v *MenuItem) toMenuItem() *C.GtkMenuItem {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalMenuItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapMenuItem(obj), nil
}

func wrapMenuItem(obj *glib.Object) *MenuItem {
	return &MenuItem{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// MenuItemNew() is a wrapper around gtk_menu_item_new().
func MenuItemNew() (*MenuItem, error) {
	c := C.gtk_menu_item_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	m := wrapMenuItem(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return m, nil
}

// MenuItemNewWithLabel() is a wrapper around gtk_menu_item_new_with_label().
func MenuItemNewWithLabel(label string) (*MenuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_menu_item_new_with_label((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	m := wrapMenuItem(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return m, nil
}

// MenuItemNewWithMnemonic() is a wrapper around
// gtk_menu_item_new_with_mnemonic().
func MenuItemNewWithMnemonic(label string) (*MenuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_menu_item_new_with_mnemonic((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	m := wrapMenuItem(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return m, nil
}

// SetSubmenu() is a wrapper around gtk_menu_item_set_submenu().
func (v *MenuItem) SetSubmenu(submenu IWidget) {
	C.gtk_menu_item_set_submenu(v.native(), submenu.toWidget())
}

// Sets text on the menu_item label
func (v *MenuItem) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_menu_item_set_label(v.native(), (*C.gchar)(cstr))
}

/*
 * GtkMenuShell
 */

// MenuShell is a representation of GTK's GtkMenuShell.
type MenuShell struct {
	Container
}

// native returns a pointer to the underlying GtkMenuShell.
func (v *MenuShell) native() *C.GtkMenuShell {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenuShell(p)
}

func marshalMenuShell(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapMenuShell(obj), nil
}

func wrapMenuShell(obj *glib.Object) *MenuShell {
	return &MenuShell{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// Append is a wrapper around gtk_menu_shell_append().
func (v *MenuShell) Append(child IMenuItem) {
	C.gtk_menu_shell_append(v.native(), child.toWidget())
}

/*
 * GtkMessageDialog
 */

// MessageDialog is a representation of GTK's GtkMessageDialog.
type MessageDialog struct {
	Dialog
}

// native returns a pointer to the underlying GtkMessageDialog.
func (v *MessageDialog) native() *C.GtkMessageDialog {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMessageDialog(p)
}

func marshalMessageDialog(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapMessageDialog(obj), nil
}

func wrapMessageDialog(obj *glib.Object) *MessageDialog {
	return &MessageDialog{Dialog{Window{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}}}
}

// MessageDialogNew() is a wrapper around gtk_message_dialog_new().
// The text is created and formatted by the format specifier and any
// additional arguments.
func MessageDialogNew(parent IWindow, flags DialogFlags, mType MessageType, buttons ButtonsType, format string, a ...interface{}) *MessageDialog {
	s := fmt.Sprintf(format, a...)
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	var w *C.GtkWindow = nil
	if parent != nil {
		w = parent.toWindow()
	}
	c := C._gtk_message_dialog_new(w,
		C.GtkDialogFlags(flags), C.GtkMessageType(mType),
		C.GtkButtonsType(buttons), cstr)
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	m := wrapMessageDialog(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return m
}

// MessageDialogNewWithMarkup is a wrapper around
// gtk_message_dialog_new_with_markup().
func MessageDialogNewWithMarkup(parent IWindow, flags DialogFlags, mType MessageType, buttons ButtonsType, format string, a ...interface{}) *MessageDialog {
	s := fmt.Sprintf(format, a...)
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	var w *C.GtkWindow = nil
	if parent != nil {
		w = parent.toWindow()
	}
	c := C._gtk_message_dialog_new_with_markup(w,
		C.GtkDialogFlags(flags), C.GtkMessageType(mType),
		C.GtkButtonsType(buttons), cstr)
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	m := wrapMessageDialog(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return m
}

// SetMarkup is a wrapper around gtk_message_dialog_set_markup().
func (v *MessageDialog) SetMarkup(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_message_dialog_set_markup(v.native(), (*C.gchar)(cstr))
}

// FormatSecondaryText is a wrapper around
// gtk_message_dialog_format_secondary_text().
func (v *MessageDialog) FormatSecondaryText(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C._gtk_message_dialog_format_secondary_text(v.native(),
		(*C.gchar)(cstr))
}

// FormatSecondaryMarkup is a wrapper around
// gtk_message_dialog_format_secondary_text().
func (v *MessageDialog) FormatSecondaryMarkup(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C._gtk_message_dialog_format_secondary_markup(v.native(),
		(*C.gchar)(cstr))
}

/*
 * GtkNotebook
 */

// Notebook is a representation of GTK's GtkNotebook.
type Notebook struct {
	Container
}

// native returns a pointer to the underlying GtkNotebook.
func (v *Notebook) native() *C.GtkNotebook {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkNotebook(p)
}

func marshalNotebook(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapNotebook(obj), nil
}

func wrapNotebook(obj *glib.Object) *Notebook {
	return &Notebook{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// NotebookNew() is a wrapper around gtk_notebook_new().
func NotebookNew() (*Notebook, error) {
	c := C.gtk_notebook_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	n := wrapNotebook(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return n, nil
}

// AppendPage() is a wrapper around gtk_notebook_append_page().
func (v *Notebook) AppendPage(child IWidget, tabLabel IWidget) int {
	cTabLabel := nullableWidget(tabLabel)
	c := C.gtk_notebook_append_page(v.native(), child.toWidget(), cTabLabel)
	return int(c)
}

// AppendPageMenu() is a wrapper around gtk_notebook_append_page_menu().
func (v *Notebook) AppendPageMenu(child IWidget, tabLabel IWidget, menuLabel IWidget) int {
	c := C.gtk_notebook_append_page_menu(v.native(), child.toWidget(),
		tabLabel.toWidget(), menuLabel.toWidget())
	return int(c)
}

// PrependPage() is a wrapper around gtk_notebook_prepend_page().
func (v *Notebook) PrependPage(child IWidget, tabLabel IWidget) int {
	cTabLabel := nullableWidget(tabLabel)
	c := C.gtk_notebook_prepend_page(v.native(), child.toWidget(), cTabLabel)
	return int(c)
}

// PrependPageMenu() is a wrapper around gtk_notebook_prepend_page_menu().
func (v *Notebook) PrependPageMenu(child IWidget, tabLabel IWidget, menuLabel IWidget) int {
	c := C.gtk_notebook_prepend_page_menu(v.native(), child.toWidget(),
		tabLabel.toWidget(), menuLabel.toWidget())
	return int(c)
}

// InsertPage() is a wrapper around gtk_notebook_insert_page().
func (v *Notebook) InsertPage(child IWidget, tabLabel IWidget, position int) int {
	label := nullableWidget(tabLabel)
	c := C.gtk_notebook_insert_page(v.native(), child.toWidget(), label, C.gint(position))

	return int(c)
}

// InsertPageMenu() is a wrapper around gtk_notebook_insert_page_menu().
func (v *Notebook) InsertPageMenu(child IWidget, tabLabel IWidget, menuLabel IWidget, position int) int {
	c := C.gtk_notebook_insert_page_menu(v.native(), child.toWidget(),
		tabLabel.toWidget(), menuLabel.toWidget(), C.gint(position))
	return int(c)
}

// RemovePage() is a wrapper around gtk_notebook_remove_page().
func (v *Notebook) RemovePage(pageNum int) {
	C.gtk_notebook_remove_page(v.native(), C.gint(pageNum))
}

// PageNum() is a wrapper around gtk_notebook_page_num().
func (v *Notebook) PageNum(child IWidget) int {
	c := C.gtk_notebook_page_num(v.native(), child.toWidget())
	return int(c)
}

// NextPage() is a wrapper around gtk_notebook_next_page().
func (v *Notebook) NextPage() {
	C.gtk_notebook_next_page(v.native())
}

// PrevPage() is a wrapper around gtk_notebook_prev_page().
func (v *Notebook) PrevPage() {
	C.gtk_notebook_prev_page(v.native())
}

// ReorderChild() is a wrapper around gtk_notebook_reorder_child().
func (v *Notebook) ReorderChild(child IWidget, position int) {
	C.gtk_notebook_reorder_child(v.native(), child.toWidget(),
		C.gint(position))
}

// SetTabPos() is a wrapper around gtk_notebook_set_tab_pos().
func (v *Notebook) SetTabPos(pos PositionType) {
	C.gtk_notebook_set_tab_pos(v.native(), C.GtkPositionType(pos))
}

// SetShowTabs() is a wrapper around gtk_notebook_set_show_tabs().
func (v *Notebook) SetShowTabs(showTabs bool) {
	C.gtk_notebook_set_show_tabs(v.native(), gbool(showTabs))
}

// SetShowBorder() is a wrapper around gtk_notebook_set_show_border().
func (v *Notebook) SetShowBorder(showBorder bool) {
	C.gtk_notebook_set_show_border(v.native(), gbool(showBorder))
}

// SetScrollable() is a wrapper around gtk_notebook_set_scrollable().
func (v *Notebook) SetScrollable(scrollable bool) {
	C.gtk_notebook_set_scrollable(v.native(), gbool(scrollable))
}

// PopupEnable() is a wrapper around gtk_notebook_popup_enable().
func (v *Notebook) PopupEnable() {
	C.gtk_notebook_popup_enable(v.native())
}

// PopupDisable() is a wrapper around gtk_notebook_popup_disable().
func (v *Notebook) PopupDisable() {
	C.gtk_notebook_popup_disable(v.native())
}

// GetCurrentPage() is a wrapper around gtk_notebook_get_current_page().
func (v *Notebook) GetCurrentPage() int {
	c := C.gtk_notebook_get_current_page(v.native())
	return int(c)
}

// GetMenuLabel() is a wrapper around gtk_notebook_get_menu_label().
func (v *Notebook) GetMenuLabel(child IWidget) (*Widget, error) {
	c := C.gtk_notebook_get_menu_label(v.native(), child.toWidget())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// GetNthPage() is a wrapper around gtk_notebook_get_nth_page().
func (v *Notebook) GetNthPage(pageNum int) (*Widget, error) {
	c := C.gtk_notebook_get_nth_page(v.native(), C.gint(pageNum))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// GetNPages() is a wrapper around gtk_notebook_get_n_pages().
func (v *Notebook) GetNPages() int {
	c := C.gtk_notebook_get_n_pages(v.native())
	return int(c)
}

// GetTabLabel() is a wrapper around gtk_notebook_get_tab_label().
func (v *Notebook) GetTabLabel(child IWidget) (*Widget, error) {
	c := C.gtk_notebook_get_tab_label(v.native(), child.toWidget())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// SetMenuLabel() is a wrapper around gtk_notebook_set_menu_label().
func (v *Notebook) SetMenuLabel(child, menuLabel IWidget) {
	C.gtk_notebook_set_menu_label(v.native(), child.toWidget(),
		menuLabel.toWidget())
}

// SetMenuLabelText() is a wrapper around gtk_notebook_set_menu_label_text().
func (v *Notebook) SetMenuLabelText(child IWidget, menuText string) {
	cstr := C.CString(menuText)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_notebook_set_menu_label_text(v.native(), child.toWidget(),
		(*C.gchar)(cstr))
}

// SetTabLabel() is a wrapper around gtk_notebook_set_tab_label().
func (v *Notebook) SetTabLabel(child, tabLabel IWidget) {
	C.gtk_notebook_set_tab_label(v.native(), child.toWidget(),
		tabLabel.toWidget())
}

// SetTabLabelText() is a wrapper around gtk_notebook_set_tab_label_text().
func (v *Notebook) SetTabLabelText(child IWidget, tabText string) {
	cstr := C.CString(tabText)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_notebook_set_tab_label_text(v.native(), child.toWidget(),
		(*C.gchar)(cstr))
}

// SetTabReorderable() is a wrapper around gtk_notebook_set_tab_reorderable().
func (v *Notebook) SetTabReorderable(child IWidget, reorderable bool) {
	C.gtk_notebook_set_tab_reorderable(v.native(), child.toWidget(),
		gbool(reorderable))
}

// SetTabDetachable() is a wrapper around gtk_notebook_set_tab_detachable().
func (v *Notebook) SetTabDetachable(child IWidget, detachable bool) {
	C.gtk_notebook_set_tab_detachable(v.native(), child.toWidget(),
		gbool(detachable))
}

// GetMenuLabelText() is a wrapper around gtk_notebook_get_menu_label_text().
func (v *Notebook) GetMenuLabelText(child IWidget) (string, error) {
	c := C.gtk_notebook_get_menu_label_text(v.native(), child.toWidget())
	if c == nil {
		return "", errors.New("No menu label for widget")
	}
	return C.GoString((*C.char)(c)), nil
}

// GetScrollable() is a wrapper around gtk_notebook_get_scrollable().
func (v *Notebook) GetScrollable() bool {
	c := C.gtk_notebook_get_scrollable(v.native())
	return gobool(c)
}

// GetShowBorder() is a wrapper around gtk_notebook_get_show_border().
func (v *Notebook) GetShowBorder() bool {
	c := C.gtk_notebook_get_show_border(v.native())
	return gobool(c)
}

// GetShowTabs() is a wrapper around gtk_notebook_get_show_tabs().
func (v *Notebook) GetShowTabs() bool {
	c := C.gtk_notebook_get_show_tabs(v.native())
	return gobool(c)
}

// GetTabLabelText() is a wrapper around gtk_notebook_get_tab_label_text().
func (v *Notebook) GetTabLabelText(child IWidget) (string, error) {
	c := C.gtk_notebook_get_tab_label_text(v.native(), child.toWidget())
	if c == nil {
		return "", errors.New("No tab label for widget")
	}
	return C.GoString((*C.char)(c)), nil
}

// GetTabPos() is a wrapper around gtk_notebook_get_tab_pos().
func (v *Notebook) GetTabPos() PositionType {
	c := C.gtk_notebook_get_tab_pos(v.native())
	return PositionType(c)
}

// GetTabReorderable() is a wrapper around gtk_notebook_get_tab_reorderable().
func (v *Notebook) GetTabReorderable(child IWidget) bool {
	c := C.gtk_notebook_get_tab_reorderable(v.native(), child.toWidget())
	return gobool(c)
}

// GetTabDetachable() is a wrapper around gtk_notebook_get_tab_detachable().
func (v *Notebook) GetTabDetachable(child IWidget) bool {
	c := C.gtk_notebook_get_tab_detachable(v.native(), child.toWidget())
	return gobool(c)
}

// SetCurrentPage() is a wrapper around gtk_notebook_set_current_page().
func (v *Notebook) SetCurrentPage(pageNum int) {
	C.gtk_notebook_set_current_page(v.native(), C.gint(pageNum))
}

// SetGroupName() is a wrapper around gtk_notebook_set_group_name().
func (v *Notebook) SetGroupName(groupName string) {
	cstr := C.CString(groupName)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_notebook_set_group_name(v.native(), (*C.gchar)(cstr))
}

// GetGroupName() is a wrapper around gtk_notebook_get_group_name().
func (v *Notebook) GetGroupName() (string, error) {
	c := C.gtk_notebook_get_group_name(v.native())
	if c == nil {
		return "", errors.New("No group name")
	}
	return C.GoString((*C.char)(c)), nil
}

// SetActionWidget() is a wrapper around gtk_notebook_set_action_widget().
func (v *Notebook) SetActionWidget(widget IWidget, packType PackType) {
	C.gtk_notebook_set_action_widget(v.native(), widget.toWidget(),
		C.GtkPackType(packType))
}

// GetActionWidget() is a wrapper around gtk_notebook_get_action_widget().
func (v *Notebook) GetActionWidget(packType PackType) (*Widget, error) {
	c := C.gtk_notebook_get_action_widget(v.native(),
		C.GtkPackType(packType))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

/*
 * GtkOffscreenWindow
 */

// OffscreenWindow is a representation of GTK's GtkOffscreenWindow.
type OffscreenWindow struct {
	Window
}

// native returns a pointer to the underlying GtkOffscreenWindow.
func (v *OffscreenWindow) native() *C.GtkOffscreenWindow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkOffscreenWindow(p)
}

func marshalOffscreenWindow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapOffscreenWindow(obj), nil
}

func wrapOffscreenWindow(obj *glib.Object) *OffscreenWindow {
	return &OffscreenWindow{Window{Bin{Container{Widget{
		glib.InitiallyUnowned{obj}}}}}}
}

// OffscreenWindowNew is a wrapper around gtk_offscreen_window_new().
func OffscreenWindowNew() (*OffscreenWindow, error) {
	c := C.gtk_offscreen_window_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	o := wrapOffscreenWindow(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return o, nil
}

// GetSurface is a wrapper around gtk_offscreen_window_get_surface().
// The returned surface is safe to use over window resizes.
func (v *OffscreenWindow) GetSurface() (*cairo.Surface, error) {
	c := C.gtk_offscreen_window_get_surface(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	cairoPtr := (uintptr)(unsafe.Pointer(c))
	s := cairo.NewSurface(cairoPtr, true)
	return s, nil
}

// GetPixbuf is a wrapper around gtk_offscreen_window_get_pixbuf().
func (v *OffscreenWindow) GetPixbuf() (*gdk.Pixbuf, error) {
	c := C.gtk_offscreen_window_get_pixbuf(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	pb := &gdk.Pixbuf{obj}
	// Pixbuf is returned with ref count of 1, so don't increment.
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return pb, nil
}

/*
 * GtkOrientable
 */

// Orientable is a representation of GTK's GtkOrientable GInterface.
type Orientable struct {
	*glib.Object
}

// IOrientable is an interface type implemented by all structs
// embedding an Orientable.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkOrientable.
type IOrientable interface {
	toOrientable() *C.GtkOrientable
}

// native returns a pointer to the underlying GObject as a GtkOrientable.
func (v *Orientable) native() *C.GtkOrientable {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkOrientable(p)
}

func marshalOrientable(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapOrientable(obj), nil
}

func wrapOrientable(obj *glib.Object) *Orientable {
	return &Orientable{obj}
}

// GetOrientation() is a wrapper around gtk_orientable_get_orientation().
func (v *Orientable) GetOrientation() Orientation {
	c := C.gtk_orientable_get_orientation(v.native())
	return Orientation(c)
}

// SetOrientation() is a wrapper around gtk_orientable_set_orientation().
func (v *Orientable) SetOrientation(orientation Orientation) {
	C.gtk_orientable_set_orientation(v.native(),
		C.GtkOrientation(orientation))
}

/*
 * GtkPaned
 */

// Paned is a representation of GTK's GtkPaned.
type Paned struct {
	Bin
}

// native returns a pointer to the underlying GtkPaned.
func (v *Paned) native() *C.GtkPaned {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkPaned(p)
}

func marshalPaned(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapPaned(obj), nil
}

func wrapPaned(obj *glib.Object) *Paned {
	return &Paned{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// PanedNew() is a wrapper around gtk_scrolled_window_new().
func PanedNew(orientation Orientation) (*Paned, error) {
	c := C.gtk_paned_new(C.GtkOrientation(orientation))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapPaned(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

// Add1() is a wrapper around gtk_paned_add1().
func (v *Paned) Add1(child IWidget) {
	C.gtk_paned_add1(v.native(), child.toWidget())
}

// Add2() is a wrapper around gtk_paned_add2().
func (v *Paned) Add2(child IWidget) {
	C.gtk_paned_add2(v.native(), child.toWidget())
}

// Pack1() is a wrapper around gtk_paned_pack1().
func (v *Paned) Pack1(child IWidget, resize, shrink bool) {
	C.gtk_paned_pack1(v.native(), child.toWidget(), gbool(resize), gbool(shrink))
}

// Pack2() is a wrapper around gtk_paned_pack2().
func (v *Paned) Pack2(child IWidget, resize, shrink bool) {
	C.gtk_paned_pack2(v.native(), child.toWidget(), gbool(resize), gbool(shrink))
}

// SetPosition() is a wrapper around gtk_paned_set_position().
func (v *Paned) SetPosition(position int) {
	C.gtk_paned_set_position(v.native(), C.gint(position))
}

// GetChild1() is a wrapper around gtk_paned_get_child1().
func (v *Paned) GetChild1() (*Widget, error) {
	c := C.gtk_paned_get_child1(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// GetChild2() is a wrapper around gtk_paned_get_child2().
func (v *Paned) GetChild2() (*Widget, error) {
	c := C.gtk_paned_get_child2(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// GetHandleWindow() is a wrapper around gtk_paned_get_handle_window().
func (v *Paned) GetHandleWindow() (*Window, error) {
	c := C.gtk_paned_get_handle_window(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWindow(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// GetPosition() is a wrapper around gtk_paned_get_position().
func (v *Paned) GetPosition() int {
	return int(C.gtk_paned_get_position(v.native()))
}

/*
 * GtkProgressBar
 */

// ProgressBar is a representation of GTK's GtkProgressBar.
type ProgressBar struct {
	Widget
}

// native returns a pointer to the underlying GtkProgressBar.
func (v *ProgressBar) native() *C.GtkProgressBar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkProgressBar(p)
}

func marshalProgressBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapProgressBar(obj), nil
}

func wrapProgressBar(obj *glib.Object) *ProgressBar {
	return &ProgressBar{Widget{glib.InitiallyUnowned{obj}}}
}

// ProgressBarNew() is a wrapper around gtk_progress_bar_new().
func ProgressBarNew() (*ProgressBar, error) {
	c := C.gtk_progress_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := wrapProgressBar(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// SetFraction() is a wrapper around gtk_progress_bar_set_fraction().
func (v *ProgressBar) SetFraction(fraction float64) {
	C.gtk_progress_bar_set_fraction(v.native(), C.gdouble(fraction))
}

// GetFraction() is a wrapper around gtk_progress_bar_get_fraction().
func (v *ProgressBar) GetFraction() float64 {
	c := C.gtk_progress_bar_get_fraction(v.native())
	return float64(c)
}

// SetShowText is a wrapper around gtk_progress_bar_set_show_text().
func (v *ProgressBar) SetShowText(showText bool) {
	C.gtk_progress_bar_set_show_text(v.native(), gbool(showText))
}

// GetShowText is a wrapper around gtk_progress_bar_get_show_text().
func (v *ProgressBar) GetShowText() bool {
	c := C.gtk_progress_bar_get_show_text(v.native())
	return gobool(c)
}

// SetText() is a wrapper around gtk_progress_bar_set_text().
func (v *ProgressBar) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_progress_bar_set_text(v.native(), (*C.gchar)(cstr))
}

/*
 * GtkRadioButton
 */

// RadioButton is a representation of GTK's GtkRadioButton.
type RadioButton struct {
	CheckButton
}

// native returns a pointer to the underlying GtkRadioButton.
func (v *RadioButton) native() *C.GtkRadioButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkRadioButton(p)
}

func marshalRadioButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapRadioButton(obj), nil
}

func wrapRadioButton(obj *glib.Object) *RadioButton {
	return &RadioButton{CheckButton{ToggleButton{Button{Bin{Container{
		Widget{glib.InitiallyUnowned{obj}}}}}}}}
}

// RadioButtonNew is a wrapper around gtk_radio_button_new().
func RadioButtonNew(group *glib.SList) (*RadioButton, error) {
	gslist := (*C.GSList)(unsafe.Pointer(group.Native()))
	c := C.gtk_radio_button_new(gslist)
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	r := wrapRadioButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return r, nil
}

// RadioButtonNewFromWidget is a wrapper around
// gtk_radio_button_new_from_widget().
func RadioButtonNewFromWidget(radioGroupMember *RadioButton) (*RadioButton, error) {
	c := C.gtk_radio_button_new_from_widget(radioGroupMember.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	r := wrapRadioButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return r, nil
}

// RadioButtonNewWithLabel is a wrapper around
// gtk_radio_button_new_with_label().
func RadioButtonNewWithLabel(group *glib.SList, label string) (*RadioButton, error) {
	gslist := (*C.GSList)(unsafe.Pointer(group.Native()))
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_button_new_with_label(gslist, (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	r := wrapRadioButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return r, nil
}

// RadioButtonNewWithLabelFromWidget is a wrapper around
// gtk_radio_button_new_with_label_from_widget().
func RadioButtonNewWithLabelFromWidget(radioGroupMember *RadioButton, label string) (*RadioButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_button_new_with_label_from_widget(radioGroupMember.native(),
		(*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	r := wrapRadioButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return r, nil
}

// RadioButtonNewWithMnemonic is a wrapper around
// gtk_radio_button_new_with_mnemonic()
func RadioButtonNewWithMnemonic(group *glib.SList, label string) (*RadioButton, error) {
	gslist := (*C.GSList)(unsafe.Pointer(group.Native()))
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_button_new_with_mnemonic(gslist, (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	r := wrapRadioButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return r, nil
}

// RadioButtonNewWithMnemonicFromWidget is a wrapper around
// gtk_radio_button_new_with_mnemonic_from_widget().
func RadioButtonNewWithMnemonicFromWidget(radioGroupMember *RadioButton, label string) (*RadioButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_button_new_with_mnemonic_from_widget(radioGroupMember.native(),
		(*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	r := wrapRadioButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return r, nil
}

// SetGroup is a wrapper around gtk_radio_button_set_group().
func (v *RadioButton) SetGroup(group *glib.SList) {
	gslist := (*C.GSList)(unsafe.Pointer(group.Native()))
	C.gtk_radio_button_set_group(v.native(), gslist)
}

// GetGroup is a wrapper around gtk_radio_button_get_group().
func (v *RadioButton) GetGroup() (*glib.SList, error) {
	c := C.gtk_radio_button_get_group(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return glib.WrapSList(uintptr(unsafe.Pointer(c))), nil
}

// JoinGroup is a wrapper around gtk_radio_button_join_group().
func (v *RadioButton) JoinGroup(groupSource *RadioButton) {
	C.gtk_radio_button_join_group(v.native(), groupSource.native())
}

/*
 * GtkRadioMenuItem
 */

// RadioMenuItem is a representation of GTK's GtkRadioMenuItem.
type RadioMenuItem struct {
	CheckMenuItem
}

// native returns a pointer to the underlying GtkRadioMenuItem.
func (v *RadioMenuItem) native() *C.GtkRadioMenuItem {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkRadioMenuItem(p)
}

func marshalRadioMenuItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapRadioMenuItem(obj), nil
}

func wrapRadioMenuItem(obj *glib.Object) *RadioMenuItem {
	return &RadioMenuItem{CheckMenuItem{MenuItem{Bin{Container{
		Widget{glib.InitiallyUnowned{obj}}}}}}}
}

// RadioMenuItemNew is a wrapper around gtk_radio_menu_item_new().
func RadioMenuItemNew(group *glib.SList) (*RadioMenuItem, error) {
	gslist := (*C.GSList)(unsafe.Pointer(group.Native()))
	c := C.gtk_radio_menu_item_new(gslist)
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	r := wrapRadioMenuItem(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return r, nil
}

// RadioMenuItemNewWithLabel is a wrapper around
// gtk_radio_menu_item_new_with_label().
func RadioMenuItemNewWithLabel(group *glib.SList, label string) (*RadioMenuItem, error) {
	gslist := (*C.GSList)(unsafe.Pointer(group.Native()))
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_menu_item_new_with_label(gslist, (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	r := wrapRadioMenuItem(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return r, nil
}

// RadioMenuItemNewWithMnemonic is a wrapper around
// gtk_radio_menu_item_new_with_mnemonic().
func RadioMenuItemNewWithMnemonic(group *glib.SList, label string) (*RadioMenuItem, error) {
	gslist := (*C.GSList)(unsafe.Pointer(group.Native()))
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_menu_item_new_with_mnemonic(gslist, (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	r := wrapRadioMenuItem(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return r, nil
}

// RadioMenuItemNewFromWidget is a wrapper around
// gtk_radio_menu_item_new_from_widget().
func RadioMenuItemNewFromWidget(group *RadioMenuItem) (*RadioMenuItem, error) {
	c := C.gtk_radio_menu_item_new_from_widget(group.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	r := wrapRadioMenuItem(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return r, nil
}

// RadioMenuItemNewWithLabelFromWidget is a wrapper around
// gtk_radio_menu_item_new_with_label_from_widget().
func RadioMenuItemNewWithLabelFromWidget(group *RadioMenuItem, label string) (*RadioMenuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_menu_item_new_with_label_from_widget(group.native(),
		(*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	r := wrapRadioMenuItem(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return r, nil
}

// RadioMenuItemNewWithMnemonicFromWidget is a wrapper around
// gtk_radio_menu_item_new_with_mnemonic_from_widget().
func RadioMenuItemNewWithMnemonicFromWidget(group *RadioMenuItem, label string) (*RadioMenuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_menu_item_new_with_mnemonic_from_widget(group.native(),
		(*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	r := wrapRadioMenuItem(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return r, nil
}

// SetGroup is a wrapper around gtk_radio_menu_item_set_group().
func (v *RadioMenuItem) SetGroup(group *glib.SList) {
	gslist := (*C.GSList)(unsafe.Pointer(group.Native()))
	C.gtk_radio_menu_item_set_group(v.native(), gslist)
}

// GetGroup is a wrapper around gtk_radio_menu_item_get_group().
func (v *RadioMenuItem) GetGroup() (*glib.SList, error) {
	c := C.gtk_radio_menu_item_get_group(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return glib.WrapSList(uintptr(unsafe.Pointer(c))), nil
}

/*
 * GtkRange
 */

// Range is a representation of GTK's GtkRange.
type Range struct {
	Widget
}

// native returns a pointer to the underlying GtkRange.
func (v *Range) native() *C.GtkRange {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkRange(p)
}

func marshalRange(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapRange(obj), nil
}

func wrapRange(obj *glib.Object) *Range {
	return &Range{Widget{glib.InitiallyUnowned{obj}}}
}

// GetValue is a wrapper around gtk_range_get_value().
func (v *Range) GetValue() float64 {
	c := C.gtk_range_get_value(v.native())
	return float64(c)
}

// SetValue is a wrapper around gtk_range_set_value().
func (v *Range) SetValue(value float64) {
	C.gtk_range_set_value(v.native(), C.gdouble(value))
}

// SetIncrements() is a wrapper around gtk_range_set_increments().
func (v *Range) SetIncrements(step, page float64) {
	C.gtk_range_set_increments(v.native(), C.gdouble(step), C.gdouble(page))
}

// SetRange() is a wrapper around gtk_range_set_range().
func (v *Range) SetRange(min, max float64) {
	C.gtk_range_set_range(v.native(), C.gdouble(min), C.gdouble(max))
}

// IRecentChooser is an interface type implemented by all structs
// embedding a RecentChooser.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkWidget.
type IRecentChooser interface {
	toRecentChooser() *C.GtkRecentChooser
}

/*
 * GtkRecentChooser
 */

// RecentChooser is a representation of GTK's GtkRecentChooser.
type RecentChooser struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkRecentChooser.
func (v *RecentChooser) native() *C.GtkRecentChooser {
	if v == nil || v.Object == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkRecentChooser(p)
}

func wrapRecentChooser(obj *glib.Object) *RecentChooser {
	return &RecentChooser{obj}
}

func (v *RecentChooser) toRecentChooser() *C.GtkRecentChooser {
	return v.native()
}

func (v *RecentChooser) GetCurrentUri() string {
	curi := C.gtk_recent_chooser_get_current_uri(v.native())
	uri := C.GoString((*C.char)(curi))
	return uri
}

func (v *RecentChooser) AddFilter(filter *RecentFilter) {
	C.gtk_recent_chooser_add_filter(v.native(), filter.native())
}

func (v *RecentChooser) RemoveFilter(filter *RecentFilter) {
	C.gtk_recent_chooser_remove_filter(v.native(), filter.native())
}

/*
 * GtkRecentChooserMenu
 */

// RecentChooserMenu is a representation of GTK's GtkRecentChooserMenu.
type RecentChooserMenu struct {
	Menu
	RecentChooser
}

// native returns a pointer to the underlying GtkRecentManager.
func (v *RecentChooserMenu) native() *C.GtkRecentChooserMenu {
	if v == nil || v.Object == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkRecentChooserMenu(p)
}

func wrapRecentChooserMenu(obj *glib.Object) *RecentChooserMenu {
	return &RecentChooserMenu{
		Menu{MenuShell{Container{Widget{glib.InitiallyUnowned{obj}}}}},
		RecentChooser{obj},
	}
}

/*
 * GtkRecentFilter
 */

// RecentFilter is a representation of GTK's GtkRecentFilter.
type RecentFilter struct {
	glib.InitiallyUnowned
}

// native returns a pointer to the underlying GtkRecentFilter.
func (v *RecentFilter) native() *C.GtkRecentFilter {
	if v == nil || v.Object == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkRecentFilter(p)
}

func wrapRecentFilter(obj *glib.Object) *RecentFilter {
	return &RecentFilter{glib.InitiallyUnowned{obj}}
}

// RecentFilterNew is a wrapper around gtk_recent_filter_new().
func RecentFilterNew() (*RecentFilter, error) {
	c := C.gtk_recent_filter_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	f := wrapRecentFilter(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return f, nil
}

/*
 * GtkRecentManager
 */

// RecentManager is a representation of GTK's GtkRecentManager.
type RecentManager struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkRecentManager.
func (v *RecentManager) native() *C.GtkRecentManager {
	if v == nil || v.Object == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkRecentManager(p)
}

func marshalRecentManager(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapRecentManager(obj), nil
}

func wrapRecentManager(obj *glib.Object) *RecentManager {
	return &RecentManager{obj}
}

// RecentManagerGetDefault is a wrapper around gtk_recent_manager_get_default().
func RecentManagerGetDefault() (*RecentManager, error) {
	c := C.gtk_recent_manager_get_default()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	v := wrapRecentManager(obj)
	return v, nil
}

// AddItem is a wrapper around gtk_recent_manager_add_item().
func (v *RecentManager) AddItem(fileURI string) bool {
	cstr := C.CString(fileURI)
	defer C.free(unsafe.Pointer(cstr))
	cok := C.gtk_recent_manager_add_item(v.native(), (*C.gchar)(cstr))
	return gobool(cok)
}

/*
 * GtkScale
 */

// Scale is a representation of GTK's GtkScale.
type Scale struct {
	Range
}

// native returns a pointer to the underlying GtkScale.
func (v *Scale) native() *C.GtkScale {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkScale(p)
}

func marshalScale(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapScale(obj), nil
}

func wrapScale(obj *glib.Object) *Scale {
	return &Scale{Range{Widget{glib.InitiallyUnowned{obj}}}}
}

// ScaleNew is a wrapper around gtk_scale_new().
func ScaleNew(orientation Orientation, adjustment *Adjustment) (*Scale, error) {
	c := C.gtk_scale_new(C.GtkOrientation(orientation), adjustment.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapScale(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

// ScaleNewWithRange is a wrapper around gtk_scale_new_with_range().
func ScaleNewWithRange(orientation Orientation, min, max, step float64) (*Scale, error) {
	c := C.gtk_scale_new_with_range(C.GtkOrientation(orientation),
		C.gdouble(min), C.gdouble(max), C.gdouble(step))

	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapScale(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

/*
 * GtkScaleButton
 */

// ScaleButton is a representation of GTK's GtkScaleButton.
type ScaleButton struct {
	Button
}

// native() returns a pointer to the underlying GtkScaleButton.
func (v *ScaleButton) native() *C.GtkScaleButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkScaleButton(p)
}

func marshalScaleButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapScaleButton(obj), nil
}

func wrapScaleButton(obj *glib.Object) *ScaleButton {
	return &ScaleButton{Button{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}}
}

// ScaleButtonNew() is a wrapper around gtk_scale_button_new().
func ScaleButtonNew(size IconSize, min, max, step float64, icons []string) (*ScaleButton, error) {
	cicons := make([]*C.gchar, len(icons))
	for i, icon := range icons {
		cicons[i] = (*C.gchar)(C.CString(icon))
		defer C.free(unsafe.Pointer(cicons[i]))
	}
	cicons = append(cicons, nil)

	c := C.gtk_scale_button_new(C.GtkIconSize(size),
		C.gdouble(min),
		C.gdouble(max),
		C.gdouble(step),
		&cicons[0])
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapScaleButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return b, nil
}

// GetAdjustment() is a wrapper around gtk_scale_button_get_adjustment().
func (v *ScaleButton) GetAdjustment() *Adjustment {
	c := C.gtk_scale_button_get_adjustment(v.native())
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &Adjustment{glib.InitiallyUnowned{obj}}
}

// GetPopup() is a wrapper around gtk_scale_button_get_popup().
func (v *ScaleButton) GetPopup() (*Widget, error) {
	c := C.gtk_scale_button_get_popup(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// GetValue() is a wrapper around gtk_scale_button_get_value().
func (v *ScaleButton) GetValue() float64 {
	return float64(C.gtk_scale_button_get_value(v.native()))
}

// SetAdjustment() is a wrapper around gtk_scale_button_set_adjustment().
func (v *ScaleButton) SetAdjustment(adjustment *Adjustment) {
	C.gtk_scale_button_set_adjustment(v.native(), adjustment.native())
}

// SetValue() is a wrapper around gtk_scale_button_set_value().
func (v *ScaleButton) SetValue(value float64) {
	C.gtk_scale_button_set_value(v.native(), C.gdouble(value))
}

/*
 * GtkScrollable
 */

// IScrollable is an interface type implemented by all structs
// embedding a Scrollable.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkScrollable.
type IScrollable interface {
	toScrollable() *C.GtkScrollable
}

// Scrollable is a representation of GTK's GtkScrollable GInterface.
type Scrollable struct {
	*glib.Object
}

// native() returns a pointer to the underlying GObject as a GtkScrollable.
func (v *Scrollable) native() *C.GtkScrollable {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkScrollable(p)
}

func wrapScrollable(obj *glib.Object) *Scrollable {
	return &Scrollable{obj}
}

func (v *Scrollable) toScrollable() *C.GtkScrollable {
	if v == nil {
		return nil
	}
	return v.native()
}

// SetHAdjustment is a wrapper around gtk_scrollable_set_hadjustment().
func (v *Scrollable) SetHAdjustment(adjustment *Adjustment) {
	C.gtk_scrollable_set_hadjustment(v.native(), adjustment.native())
}

// GetHAdjustment is a wrapper around gtk_scrollable_get_hadjustment().
func (v *Scrollable) GetHAdjustment() (*Adjustment, error) {
	c := C.gtk_scrollable_get_hadjustment(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	a := wrapAdjustment(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return a, nil
}

// SetVAdjustment is a wrapper around gtk_scrollable_set_vadjustment().
func (v *Scrollable) SetVAdjustment(adjustment *Adjustment) {
	C.gtk_scrollable_set_vadjustment(v.native(), adjustment.native())
}

// GetVAdjustment is a wrapper around gtk_scrollable_get_vadjustment().
func (v *Scrollable) GetVAdjustment() (*Adjustment, error) {
	c := C.gtk_scrollable_get_vadjustment(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	a := wrapAdjustment(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return a, nil
}

/*
 * GtkScrollbar
 */

// Scrollbar is a representation of GTK's GtkScrollbar.
type Scrollbar struct {
	Range
}

// native returns a pointer to the underlying GtkScrollbar.
func (v *Scrollbar) native() *C.GtkScrollbar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkScrollbar(p)
}

func marshalScrollbar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapScrollbar(obj), nil
}

func wrapScrollbar(obj *glib.Object) *Scrollbar {
	return &Scrollbar{Range{Widget{glib.InitiallyUnowned{obj}}}}
}

// ScrollbarNew is a wrapper around gtk_scrollbar_new().
func ScrollbarNew(orientation Orientation, adjustment *Adjustment) (*Scrollbar, error) {
	c := C.gtk_scrollbar_new(C.GtkOrientation(orientation), adjustment.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapScrollbar(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

/*
 * GtkScrolledWindow
 */

// ScrolledWindow is a representation of GTK's GtkScrolledWindow.
type ScrolledWindow struct {
	Bin
}

// native returns a pointer to the underlying GtkScrolledWindow.
func (v *ScrolledWindow) native() *C.GtkScrolledWindow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkScrolledWindow(p)
}

func marshalScrolledWindow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapScrolledWindow(obj), nil
}

func wrapScrolledWindow(obj *glib.Object) *ScrolledWindow {
	return &ScrolledWindow{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// ScrolledWindowNew() is a wrapper around gtk_scrolled_window_new().
func ScrolledWindowNew(hadjustment, vadjustment *Adjustment) (*ScrolledWindow, error) {
	c := C.gtk_scrolled_window_new(hadjustment.native(),
		vadjustment.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapScrolledWindow(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

// SetPolicy() is a wrapper around gtk_scrolled_window_set_policy().
func (v *ScrolledWindow) SetPolicy(hScrollbarPolicy, vScrollbarPolicy PolicyType) {
	C.gtk_scrolled_window_set_policy(v.native(),
		C.GtkPolicyType(hScrollbarPolicy),
		C.GtkPolicyType(vScrollbarPolicy))
}

// GetHAdjustment() is a wrapper around gtk_scrolled_window_get_hadjustment().
func (v *ScrolledWindow) GetHAdjustment() *Adjustment {
	c := C.gtk_scrolled_window_get_hadjustment(v.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	a := wrapAdjustment(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return a
}

// SetHAdjustment is a wrapper around gtk_scrolled_window_set_hadjustment().
func (v *ScrolledWindow) SetHAdjustment(adjustment *Adjustment) {
	C.gtk_scrolled_window_set_hadjustment(v.native(), adjustment.native())
}

// GetVAdjustment() is a wrapper around gtk_scrolled_window_get_vadjustment().
func (v *ScrolledWindow) GetVAdjustment() *Adjustment {
	c := C.gtk_scrolled_window_get_vadjustment(v.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	a := wrapAdjustment(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return a
}

// SetVAdjustment is a wrapper around gtk_scrolled_window_set_vadjustment().
func (v *ScrolledWindow) SetVAdjustment(adjustment *Adjustment) {
	C.gtk_scrolled_window_set_vadjustment(v.native(), adjustment.native())
}

/*
 * GtkSearchEntry
 */

// SearchEntry is a reprensentation of GTK's GtkSearchEntry.
type SearchEntry struct {
	Entry
}

// native returns a pointer to the underlying GtkSearchEntry.
func (v *SearchEntry) native() *C.GtkSearchEntry {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSearchEntry(p)
}

func marshalSearchEntry(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapSearchEntry(obj), nil
}

func wrapSearchEntry(obj *glib.Object) *SearchEntry {
	e := wrapEditable(obj)
	return &SearchEntry{Entry{Widget{glib.InitiallyUnowned{obj}}, *e}}
}

// SearchEntryNew is a wrapper around gtk_search_entry_new().
func SearchEntryNew() (*SearchEntry, error) {
	c := C.gtk_search_entry_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapSearchEntry(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

/*
* GtkSelectionData
 */
type SelectionData struct {
	GtkSelectionData *C.GtkSelectionData
}

func marshalSelectionData(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*SelectionData)(unsafe.Pointer(c)), nil
}

// native returns a pointer to the underlying GtkSelectionData.
func (v *SelectionData) native() *C.GtkSelectionData {
	if v == nil {
		return nil
	}
	return v.GtkSelectionData
}

// GetLength is a wrapper around gtk_selection_data_get_length
func (v *SelectionData) GetLength() int {
	return int(C.gtk_selection_data_get_length(v.native()))
}

// GetData is a wrapper around gtk_selection_data_get_data_with_length.
// It returns a slice of the correct size with the selection's data.
func (v *SelectionData) GetData() (data []byte) {
	var length C.gint
	c := C.gtk_selection_data_get_data_with_length(v.native(), &length)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sliceHeader.Data = uintptr(unsafe.Pointer(c))
	sliceHeader.Len = int(length)
	sliceHeader.Cap = int(length)
	return
}

func (v *SelectionData) free() {
	C.gtk_selection_data_free(v.native())
}

/*
 * GtkSeparator
 */

// Separator is a representation of GTK's GtkSeparator.
type Separator struct {
	Widget
}

// native returns a pointer to the underlying GtkSeperator.
func (v *Separator) native() *C.GtkSeparator {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSeparator(p)
}

func marshalSeparator(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapSeparator(obj), nil
}

func wrapSeparator(obj *glib.Object) *Separator {
	return &Separator{Widget{glib.InitiallyUnowned{obj}}}
}

// SeparatorNew is a wrapper around gtk_separator_new().
func SeparatorNew(orientation Orientation) (*Separator, error) {
	c := C.gtk_separator_new(C.GtkOrientation(orientation))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapSeparator(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

/*
 * GtkSeparatorMenuItem
 */

// SeparatorMenuItem is a representation of GTK's GtkSeparatorMenuItem.
type SeparatorMenuItem struct {
	MenuItem
}

// native returns a pointer to the underlying GtkSeparatorMenuItem.
func (v *SeparatorMenuItem) native() *C.GtkSeparatorMenuItem {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSeparatorMenuItem(p)
}

func marshalSeparatorMenuItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapSeparatorMenuItem(obj), nil
}

func wrapSeparatorMenuItem(obj *glib.Object) *SeparatorMenuItem {
	return &SeparatorMenuItem{MenuItem{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}}
}

// SeparatorMenuItemNew is a wrapper around gtk_separator_menu_item_new().
func SeparatorMenuItemNew() (*SeparatorMenuItem, error) {
	c := C.gtk_separator_menu_item_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapSeparatorMenuItem(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

/*
 * GtkSeparatorToolItem
 */

// SeparatorToolItem is a representation of GTK's GtkSeparatorToolItem.
type SeparatorToolItem struct {
	ToolItem
}

// native returns a pointer to the underlying GtkSeparatorToolItem.
func (v *SeparatorToolItem) native() *C.GtkSeparatorToolItem {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSeparatorToolItem(p)
}

func marshalSeparatorToolItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapSeparatorToolItem(obj), nil
}

func wrapSeparatorToolItem(obj *glib.Object) *SeparatorToolItem {
	return &SeparatorToolItem{ToolItem{Bin{Container{Widget{
		glib.InitiallyUnowned{obj}}}}}}
}

// SeparatorToolItemNew is a wrapper around gtk_separator_tool_item_new().
func SeparatorToolItemNew() (*SeparatorToolItem, error) {
	c := C.gtk_separator_tool_item_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapSeparatorToolItem(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

// SetDraw is a wrapper around gtk_separator_tool_item_set_draw().
func (v *SeparatorToolItem) SetDraw(draw bool) {
	C.gtk_separator_tool_item_set_draw(v.native(), gbool(draw))
}

// GetDraw is a wrapper around gtk_separator_tool_item_get_draw().
func (v *SeparatorToolItem) GetDraw() bool {
	c := C.gtk_separator_tool_item_get_draw(v.native())
	return gobool(c)
}

/*
 * GtkSpinButton
 */

// SpinButton is a representation of GTK's GtkSpinButton.
type SpinButton struct {
	Entry
}

// native returns a pointer to the underlying GtkSpinButton.
func (v *SpinButton) native() *C.GtkSpinButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSpinButton(p)
}

func marshalSpinButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapSpinButton(obj), nil
}

func wrapSpinButton(obj *glib.Object) *SpinButton {
	e := wrapEditable(obj)
	return &SpinButton{Entry{Widget{glib.InitiallyUnowned{obj}}, *e}}
}

// Configure() is a wrapper around gtk_spin_button_configure().
func (v *SpinButton) Configure(adjustment *Adjustment, climbRate float64, digits uint) {
	C.gtk_spin_button_configure(v.native(), adjustment.native(),
		C.gdouble(climbRate), C.guint(digits))
}

// SpinButtonNew() is a wrapper around gtk_spin_button_new().
func SpinButtonNew(adjustment *Adjustment, climbRate float64, digits uint) (*SpinButton, error) {
	c := C.gtk_spin_button_new(adjustment.native(),
		C.gdouble(climbRate), C.guint(digits))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapSpinButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

// SpinButtonNewWithRange() is a wrapper around
// gtk_spin_button_new_with_range().
func SpinButtonNewWithRange(min, max, step float64) (*SpinButton, error) {
	c := C.gtk_spin_button_new_with_range(C.gdouble(min), C.gdouble(max),
		C.gdouble(step))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapSpinButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

// GetValueAsInt() is a wrapper around gtk_spin_button_get_value_as_int().
func (v *SpinButton) GetValueAsInt() int {
	c := C.gtk_spin_button_get_value_as_int(v.native())
	return int(c)
}

// SetValue() is a wrapper around gtk_spin_button_set_value().
func (v *SpinButton) SetValue(value float64) {
	C.gtk_spin_button_set_value(v.native(), C.gdouble(value))
}

// GetValue() is a wrapper around gtk_spin_button_get_value().
func (v *SpinButton) GetValue() float64 {
	c := C.gtk_spin_button_get_value(v.native())
	return float64(c)
}

// GetAdjustment() is a wrapper around gtk_spin_button_get_adjustment
func (v *SpinButton) GetAdjustment() *Adjustment {
	c := C.gtk_spin_button_get_adjustment(v.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	a := wrapAdjustment(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return a
}

// SetRange is a wrapper around gtk_spin_button_set_range().
func (v *SpinButton) SetRange(min, max float64) {
	C.gtk_spin_button_set_range(v.native(), C.gdouble(min), C.gdouble(max))
}

// SetIncrements() is a wrapper around gtk_spin_button_set_increments().
func (v *SpinButton) SetIncrements(step, page float64) {
	C.gtk_spin_button_set_increments(v.native(), C.gdouble(step), C.gdouble(page))
}

/*
 * GtkSpinner
 */

// Spinner is a representation of GTK's GtkSpinner.
type Spinner struct {
	Widget
}

// native returns a pointer to the underlying GtkSpinner.
func (v *Spinner) native() *C.GtkSpinner {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSpinner(p)
}

func marshalSpinner(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapSpinner(obj), nil
}

func wrapSpinner(obj *glib.Object) *Spinner {
	return &Spinner{Widget{glib.InitiallyUnowned{obj}}}
}

// SpinnerNew is a wrapper around gtk_spinner_new().
func SpinnerNew() (*Spinner, error) {
	c := C.gtk_spinner_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapSpinner(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

// Start is a wrapper around gtk_spinner_start().
func (v *Spinner) Start() {
	C.gtk_spinner_start(v.native())
}

// Stop is a wrapper around gtk_spinner_stop().
func (v *Spinner) Stop() {
	C.gtk_spinner_stop(v.native())
}

/*
 * GtkStatusbar
 */

// Statusbar is a representation of GTK's GtkStatusbar
type Statusbar struct {
	Box
}

// native returns a pointer to the underlying GtkStatusbar
func (v *Statusbar) native() *C.GtkStatusbar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkStatusbar(p)
}

func marshalStatusbar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapStatusbar(obj), nil
}

func wrapStatusbar(obj *glib.Object) *Statusbar {
	return &Statusbar{Box{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// StatusbarNew() is a wrapper around gtk_statusbar_new().
func StatusbarNew() (*Statusbar, error) {
	c := C.gtk_statusbar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapStatusbar(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

// GetContextId() is a wrapper around gtk_statusbar_get_context_id().
func (v *Statusbar) GetContextId(contextDescription string) uint {
	cstr := C.CString(contextDescription)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_statusbar_get_context_id(v.native(), (*C.gchar)(cstr))
	return uint(c)
}

// Push() is a wrapper around gtk_statusbar_push().
func (v *Statusbar) Push(contextID uint, text string) uint {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_statusbar_push(v.native(), C.guint(contextID),
		(*C.gchar)(cstr))
	return uint(c)
}

// Pop() is a wrapper around gtk_statusbar_pop().
func (v *Statusbar) Pop(contextID uint) {
	C.gtk_statusbar_pop(v.native(), C.guint(contextID))
}

// GetMessageArea() is a wrapper around gtk_statusbar_get_message_area().
func (v *Statusbar) GetMessageArea() (*Box, error) {
	c := C.gtk_statusbar_get_message_area(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &Box{Container{Widget{glib.InitiallyUnowned{obj}}}}, nil
}

/*
 * GtkSwitch
 */

// Switch is a representation of GTK's GtkSwitch.
type Switch struct {
	Widget
}

// native returns a pointer to the underlying GtkSwitch.
func (v *Switch) native() *C.GtkSwitch {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSwitch(p)
}

func marshalSwitch(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapSwitch(obj), nil
}

func wrapSwitch(obj *glib.Object) *Switch {
	return &Switch{Widget{glib.InitiallyUnowned{obj}}}
}

// SwitchNew is a wrapper around gtk_switch_new().
func SwitchNew() (*Switch, error) {
	c := C.gtk_switch_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapSwitch(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

// GetActive is a wrapper around gtk_switch_get_active().
func (v *Switch) GetActive() bool {
	c := C.gtk_switch_get_active(v.native())
	return gobool(c)
}

// SetActive is a wrapper around gtk_switch_set_active().
func (v *Switch) SetActive(isActive bool) {
	C.gtk_switch_set_active(v.native(), gbool(isActive))
}

/*
 * GtkTargetEntry
 */

// TargetEntry is a representation of GTK's GtkTargetEntry
type TargetEntry C.GtkTargetEntry

func marshalTargetEntry(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*TargetEntry)(unsafe.Pointer(c)), nil
}

func (v *TargetEntry) native() *C.GtkTargetEntry {
	return (*C.GtkTargetEntry)(unsafe.Pointer(v))
}

// TargetEntryNew is a wrapper aroud gtk_target_entry_new().
func TargetEntryNew(target string, flags TargetFlags, info uint) (*TargetEntry, error) {
	cstr := C.CString(target)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_target_entry_new((*C.gchar)(cstr), C.guint(flags), C.guint(info))
	if c == nil {
		return nil, nilPtrErr
	}
	t := (*TargetEntry)(unsafe.Pointer(c))
	runtime.SetFinalizer(t, (*TargetEntry).free)
	return t, nil
}

func (v *TargetEntry) free() {
	C.gtk_target_entry_free(v.native())
}

/*
 * GtkTextView
 */

// TextView is a representation of GTK's GtkTextView
type TextView struct {
	Container
}

// native returns a pointer to the underlying GtkTextView.
func (v *TextView) native() *C.GtkTextView {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTextView(p)
}

func marshalTextView(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapTextView(obj), nil
}

func wrapTextView(obj *glib.Object) *TextView {
	return &TextView{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// TextViewNew is a wrapper around gtk_text_view_new().
func TextViewNew() (*TextView, error) {
	c := C.gtk_text_view_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	t := wrapTextView(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return t, nil
}

// TextViewNewWithBuffer is a wrapper around gtk_text_view_new_with_buffer().
func TextViewNewWithBuffer(buf *TextBuffer) (*TextView, error) {
	cbuf := buf.native()
	c := C.gtk_text_view_new_with_buffer(cbuf)
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	t := wrapTextView(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return t, nil
}

// GetBuffer is a wrapper around gtk_text_view_get_buffer().
func (v *TextView) GetBuffer() (*TextBuffer, error) {
	c := C.gtk_text_view_get_buffer(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	t := wrapTextBuffer(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return t, nil
}

// SetBuffer is a wrapper around gtk_text_view_set_buffer().
func (v *TextView) SetBuffer(buffer *TextBuffer) {
	C.gtk_text_view_set_buffer(v.native(), buffer.native())
}

// SetEditable is a wrapper around gtk_text_view_set_editable().
func (v *TextView) SetEditable(editable bool) {
	C.gtk_text_view_set_editable(v.native(), gbool(editable))
}

// GetEditable is a wrapper around gtk_text_view_get_editable().
func (v *TextView) GetEditable() bool {
	c := C.gtk_text_view_get_editable(v.native())
	return gobool(c)
}

// SetWrapMode is a wrapper around gtk_text_view_set_wrap_mode().
func (v *TextView) SetWrapMode(wrapMode WrapMode) {
	C.gtk_text_view_set_wrap_mode(v.native(), C.GtkWrapMode(wrapMode))
}

// GetWrapMode is a wrapper around gtk_text_view_get_wrap_mode().
func (v *TextView) GetWrapMode() WrapMode {
	return WrapMode(C.gtk_text_view_get_wrap_mode(v.native()))
}

// SetCursorVisible is a wrapper around gtk_text_view_set_cursor_visible().
func (v *TextView) SetCursorVisible(visible bool) {
	C.gtk_text_view_set_cursor_visible(v.native(), gbool(visible))
}

// GetCursorVisible is a wrapper around gtk_text_view_get_cursor_visible().
func (v *TextView) GetCursorVisible() bool {
	c := C.gtk_text_view_get_cursor_visible(v.native())
	return gobool(c)
}

// SetOverwrite is a wrapper around gtk_text_view_set_overwrite().
func (v *TextView) SetOverwrite(overwrite bool) {
	C.gtk_text_view_set_overwrite(v.native(), gbool(overwrite))
}

// GetOverwrite is a wrapper around gtk_text_view_get_overwrite().
func (v *TextView) GetOverwrite() bool {
	c := C.gtk_text_view_get_overwrite(v.native())
	return gobool(c)
}

// SetJustification is a wrapper around gtk_text_view_set_justification().
func (v *TextView) SetJustification(justify Justification) {
	C.gtk_text_view_set_justification(v.native(), C.GtkJustification(justify))
}

// GetJustification is a wrapper around gtk_text_view_get_justification().
func (v *TextView) GetJustification() Justification {
	c := C.gtk_text_view_get_justification(v.native())
	return Justification(c)
}

// SetAcceptsTab is a wrapper around gtk_text_view_set_accepts_tab().
func (v *TextView) SetAcceptsTab(acceptsTab bool) {
	C.gtk_text_view_set_accepts_tab(v.native(), gbool(acceptsTab))
}

// GetAcceptsTab is a wrapper around gtk_text_view_get_accepts_tab().
func (v *TextView) GetAcceptsTab() bool {
	c := C.gtk_text_view_get_accepts_tab(v.native())
	return gobool(c)
}

// SetPixelsAboveLines is a wrapper around gtk_text_view_set_pixels_above_lines().
func (v *TextView) SetPixelsAboveLines(px int) {
	C.gtk_text_view_set_pixels_above_lines(v.native(), C.gint(px))
}

// GetPixelsAboveLines is a wrapper around gtk_text_view_get_pixels_above_lines().
func (v *TextView) GetPixelsAboveLines() int {
	c := C.gtk_text_view_get_pixels_above_lines(v.native())
	return int(c)
}

// SetPixelsBelowLines is a wrapper around gtk_text_view_set_pixels_below_lines().
func (v *TextView) SetPixelsBelowLines(px int) {
	C.gtk_text_view_set_pixels_below_lines(v.native(), C.gint(px))
}

// GetPixelsBelowLines is a wrapper around gtk_text_view_get_pixels_below_lines().
func (v *TextView) GetPixelsBelowLines() int {
	c := C.gtk_text_view_get_pixels_below_lines(v.native())
	return int(c)
}

// SetPixelsInsideWrap is a wrapper around gtk_text_view_set_pixels_inside_wrap().
func (v *TextView) SetPixelsInsideWrap(px int) {
	C.gtk_text_view_set_pixels_inside_wrap(v.native(), C.gint(px))
}

// GetPixelsInsideWrap is a wrapper around gtk_text_view_get_pixels_inside_wrap().
func (v *TextView) GetPixelsInsideWrap() int {
	c := C.gtk_text_view_get_pixels_inside_wrap(v.native())
	return int(c)
}

// SetLeftMargin is a wrapper around gtk_text_view_set_left_margin().
func (v *TextView) SetLeftMargin(margin int) {
	C.gtk_text_view_set_left_margin(v.native(), C.gint(margin))
}

// GetLeftMargin is a wrapper around gtk_text_view_get_left_margin().
func (v *TextView) GetLeftMargin() int {
	c := C.gtk_text_view_get_left_margin(v.native())
	return int(c)
}

// SetRightMargin is a wrapper around gtk_text_view_set_right_margin().
func (v *TextView) SetRightMargin(margin int) {
	C.gtk_text_view_set_right_margin(v.native(), C.gint(margin))
}

// GetRightMargin is a wrapper around gtk_text_view_get_right_margin().
func (v *TextView) GetRightMargin() int {
	c := C.gtk_text_view_get_right_margin(v.native())
	return int(c)
}

// SetIndent is a wrapper around gtk_text_view_set_indent().
func (v *TextView) SetIndent(indent int) {
	C.gtk_text_view_set_indent(v.native(), C.gint(indent))
}

// GetIndent is a wrapper around gtk_text_view_get_indent().
func (v *TextView) GetIndent() int {
	c := C.gtk_text_view_get_indent(v.native())
	return int(c)
}

// SetInputHints is a wrapper around gtk_text_view_set_input_hints().
func (v *TextView) SetInputHints(hints InputHints) {
	C.gtk_text_view_set_input_hints(v.native(), C.GtkInputHints(hints))
}

// GetInputHints is a wrapper around gtk_text_view_get_input_hints().
func (v *TextView) GetInputHints() InputHints {
	c := C.gtk_text_view_get_input_hints(v.native())
	return InputHints(c)
}

// SetInputPurpose is a wrapper around gtk_text_view_set_input_purpose().
func (v *TextView) SetInputPurpose(purpose InputPurpose) {
	C.gtk_text_view_set_input_purpose(v.native(),
		C.GtkInputPurpose(purpose))
}

// GetInputPurpose is a wrapper around gtk_text_view_get_input_purpose().
func (v *TextView) GetInputPurpose() InputPurpose {
	c := C.gtk_text_view_get_input_purpose(v.native())
	return InputPurpose(c)
}

/*
 * GtkTextTag
 */

type TextTag struct {
	*glib.Object
}

// native returns a pointer to the underlying GObject as a GtkTextTag.
func (v *TextTag) native() *C.GtkTextTag {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTextTag(p)
}

func marshalTextTag(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapTextTag(obj), nil
}

func wrapTextTag(obj *glib.Object) *TextTag {
	return &TextTag{obj}
}

func TextTagNew(name string) (*TextTag, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	c := C.gtk_text_tag_new((*C.gchar)(cname))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	t := wrapTextTag(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return t, nil
}

// GetPriority() is a wrapper around gtk_text_tag_get_priority().
func (v *TextTag) GetPriority() int {
	return int(C.gtk_text_tag_get_priority(v.native()))
}

// SetPriority() is a wrapper around gtk_text_tag_set_priority().
func (v *TextTag) SetPriority(priority int) {
	C.gtk_text_tag_set_priority(v.native(), C.gint(priority))
}

// Event() is a wrapper around gtk_text_tag_event().
func (v *TextTag) Event(eventObject *glib.Object, event *gdk.Event, iter *TextIter) bool {
	ok := C.gtk_text_tag_event(v.native(),
		(*C.GObject)(unsafe.Pointer(eventObject.Native())),
		(*C.GdkEvent)(unsafe.Pointer(event.Native())),
		(*C.GtkTextIter)(iter),
	)
	return gobool(ok)
}

/*
 * GtkTextTagTable
 */

type TextTagTable struct {
	*glib.Object
}

// native returns a pointer to the underlying GObject as a GtkTextTagTable.
func (v *TextTagTable) native() *C.GtkTextTagTable {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTextTagTable(p)
}

func marshalTextTagTable(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapTextTagTable(obj), nil
}

func wrapTextTagTable(obj *glib.Object) *TextTagTable {
	return &TextTagTable{obj}
}

func TextTagTableNew() (*TextTagTable, error) {
	c := C.gtk_text_tag_table_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	t := wrapTextTagTable(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return t, nil
}

// Add() is a wrapper around gtk_text_tag_table_add().
func (v *TextTagTable) Add(tag *TextTag) {
	C.gtk_text_tag_table_add(v.native(), tag.native())
	//return gobool(c) // TODO version-separate
}

// Lookup() is a wrapper around gtk_text_tag_table_lookup().
func (v *TextTagTable) Lookup(name string) (*TextTag, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	c := C.gtk_text_tag_table_lookup(v.native(), (*C.gchar)(cname))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapTextTag(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// Remove() is a wrapper around gtk_text_tag_table_remove().
func (v *TextTagTable) Remove(tag *TextTag) {
	C.gtk_text_tag_table_remove(v.native(), tag.native())
}

/*
 * GtkTextBuffer
 */

// TextBuffer is a representation of GTK's GtkTextBuffer.
type TextBuffer struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkTextBuffer.
func (v *TextBuffer) native() *C.GtkTextBuffer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTextBuffer(p)
}

func marshalTextBuffer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapTextBuffer(obj), nil
}

func wrapTextBuffer(obj *glib.Object) *TextBuffer {
	return &TextBuffer{obj}
}

// TextBufferNew() is a wrapper around gtk_text_buffer_new().
func TextBufferNew(table *TextTagTable) (*TextBuffer, error) {
	c := C.gtk_text_buffer_new(table.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	e := wrapTextBuffer(obj)
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return e, nil
}

// ApplyTag() is a wrapper around gtk_text_buffer_apply_tag().
func (v *TextBuffer) ApplyTag(tag *TextTag, start, end *TextIter) {
	C.gtk_text_buffer_apply_tag(v.native(), tag.native(), (*C.GtkTextIter)(start), (*C.GtkTextIter)(end))
}

// ApplyTagByName() is a wrapper around gtk_text_buffer_apply_tag_by_name().
func (v *TextBuffer) ApplyTagByName(name string, start, end *TextIter) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_apply_tag_by_name(v.native(), (*C.gchar)(cstr),
		(*C.GtkTextIter)(start), (*C.GtkTextIter)(end))
}

// Delete() is a wrapper around gtk_text_buffer_delete().
func (v *TextBuffer) Delete(start, end *TextIter) {
	C.gtk_text_buffer_delete(v.native(), (*C.GtkTextIter)(start), (*C.GtkTextIter)(end))
}

func (v *TextBuffer) GetBounds() (start, end *TextIter) {
	start, end = new(TextIter), new(TextIter)
	C.gtk_text_buffer_get_bounds(v.native(), (*C.GtkTextIter)(start), (*C.GtkTextIter)(end))
	return
}

// GetCharCount() is a wrapper around gtk_text_buffer_get_char_count().
func (v *TextBuffer) GetCharCount() int {
	return int(C.gtk_text_buffer_get_char_count(v.native()))
}

// GetIterAtOffset() is a wrapper around gtk_text_buffer_get_iter_at_offset().
func (v *TextBuffer) GetIterAtOffset(charOffset int) *TextIter {
	var iter C.GtkTextIter
	C.gtk_text_buffer_get_iter_at_offset(v.native(), &iter, C.gint(charOffset))
	return (*TextIter)(&iter)
}

// GetStartIter() is a wrapper around gtk_text_buffer_get_start_iter().
func (v *TextBuffer) GetStartIter() *TextIter {
	var iter C.GtkTextIter
	C.gtk_text_buffer_get_start_iter(v.native(), &iter)
	return (*TextIter)(&iter)
}

// GetEndIter() is a wrapper around gtk_text_buffer_get_end_iter().
func (v *TextBuffer) GetEndIter() *TextIter {
	var iter C.GtkTextIter
	C.gtk_text_buffer_get_end_iter(v.native(), &iter)
	return (*TextIter)(&iter)
}

// GetLineCount() is a wrapper around gtk_text_buffer_get_line_count().
func (v *TextBuffer) GetLineCount() int {
	return int(C.gtk_text_buffer_get_line_count(v.native()))
}

// GetModified() is a wrapper around gtk_text_buffer_get_modified().
func (v *TextBuffer) GetModified() bool {
	return gobool(C.gtk_text_buffer_get_modified(v.native()))
}

// GetTagTable() is a wrapper around gtk_text_buffer_get_tag_table().
func (v *TextBuffer) GetTagTable() (*TextTagTable, error) {
	c := C.gtk_text_buffer_get_tag_table(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapTextTagTable(obj), nil
}

func (v *TextBuffer) GetText(start, end *TextIter, includeHiddenChars bool) (string, error) {
	c := C.gtk_text_buffer_get_text(
		v.native(), (*C.GtkTextIter)(start), (*C.GtkTextIter)(end), gbool(includeHiddenChars),
	)
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// Insert() is a wrapper around gtk_text_buffer_insert().
func (v *TextBuffer) Insert(iter *TextIter, text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_insert(v.native(), (*C.GtkTextIter)(iter), (*C.gchar)(cstr), C.gint(len(text)))
}

// InsertAtCursor() is a wrapper around gtk_text_buffer_insert_at_cursor().
func (v *TextBuffer) InsertAtCursor(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_insert_at_cursor(v.native(), (*C.gchar)(cstr), C.gint(len(text)))
}

// RemoveTag() is a wrapper around gtk_text_buffer_remove_tag().
func (v *TextBuffer) RemoveTag(tag *TextTag, start, end *TextIter) {
	C.gtk_text_buffer_remove_tag(v.native(), tag.native(), (*C.GtkTextIter)(start), (*C.GtkTextIter)(end))
}

// SetModified() is a wrapper around gtk_text_buffer_set_modified().
func (v *TextBuffer) SetModified(setting bool) {
	C.gtk_text_buffer_set_modified(v.native(), gbool(setting))
}

func (v *TextBuffer) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_set_text(v.native(), (*C.gchar)(cstr),
		C.gint(len(text)))
}

/*
 * GtkTextIter
 */

// TextIter is a representation of GTK's GtkTextIter
type TextIter C.GtkTextIter

func marshalTextIter(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*TextIter)(unsafe.Pointer(c)), nil
}

/*
 * GtkToggleButton
 */

// ToggleButton is a representation of GTK's GtkToggleButton.
type ToggleButton struct {
	Button
}

// native returns a pointer to the underlying GtkToggleButton.
func (v *ToggleButton) native() *C.GtkToggleButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkToggleButton(p)
}

func marshalToggleButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapToggleButton(obj), nil
}

func wrapToggleButton(obj *glib.Object) *ToggleButton {
	return &ToggleButton{Button{Bin{Container{Widget{
		glib.InitiallyUnowned{obj}}}}}}
}

// ToggleButtonNew is a wrapper around gtk_toggle_button_new().
func ToggleButtonNew() (*ToggleButton, error) {
	c := C.gtk_toggle_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	tb := wrapToggleButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return tb, nil
}

// ToggleButtonNewWithLabel is a wrapper around
// gtk_toggle_button_new_with_label().
func ToggleButtonNewWithLabel(label string) (*ToggleButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_toggle_button_new_with_label((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	tb := wrapToggleButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return tb, nil
}

// ToggleButtonNewWithMnemonic is a wrapper around
// gtk_toggle_button_new_with_mnemonic().
func ToggleButtonNewWithMnemonic(label string) (*ToggleButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_toggle_button_new_with_mnemonic((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	tb := wrapToggleButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return tb, nil
}

// GetActive is a wrapper around gtk_toggle_button_get_active().
func (v *ToggleButton) GetActive() bool {
	c := C.gtk_toggle_button_get_active(v.native())
	return gobool(c)
}

// SetActive is a wrapper around gtk_toggle_button_set_active().
func (v *ToggleButton) SetActive(isActive bool) {
	C.gtk_toggle_button_set_active(v.native(), gbool(isActive))
}

/*
 * GtkToolbar
 */

// Toolbar is a representation of GTK's GtkToolbar.
type Toolbar struct {
	Container
}

// native returns a pointer to the underlying GtkToolbar.
func (v *Toolbar) native() *C.GtkToolbar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkToolbar(p)
}

func marshalToolbar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapToolbar(obj), nil
}

func wrapToolbar(obj *glib.Object) *Toolbar {
	return &Toolbar{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// ToolbarNew is a wrapper around gtk_toolbar_new().
func ToolbarNew() (*Toolbar, error) {
	c := C.gtk_toolbar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	tb := wrapToolbar(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return tb, nil
}

// Insert is a wrapper around gtk_toolbar_insert().
func (v *Toolbar) Insert(item IToolItem, pos int) {
	C.gtk_toolbar_insert(v.native(), item.toToolItem(), C.gint(pos))
}

// GetItemIndex is a wrapper around gtk_toolbar_get_item_index().
func (v *Toolbar) GetItemIndex(item IToolItem) int {
	c := C.gtk_toolbar_get_item_index(v.native(), item.toToolItem())
	return int(c)
}

// GetNItems is a wrapper around gtk_toolbar_get_n_items().
func (v *Toolbar) GetNItems() int {
	c := C.gtk_toolbar_get_n_items(v.native())
	return int(c)
}

// GetNthItem is a wrapper around gtk_toolbar_get_nth_item().
func (v *Toolbar) GetNthItem(n int) *ToolItem {
	c := C.gtk_toolbar_get_nth_item(v.native(), C.gint(n))
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	ti := wrapToolItem(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return ti
}

// GetDropIndex is a wrapper around gtk_toolbar_get_drop_index().
func (v *Toolbar) GetDropIndex(x, y int) int {
	c := C.gtk_toolbar_get_drop_index(v.native(), C.gint(x), C.gint(y))
	return int(c)
}

// SetDropHighlightItem is a wrapper around
// gtk_toolbar_set_drop_highlight_item().
func (v *Toolbar) SetDropHighlightItem(toolItem IToolItem, index int) {
	C.gtk_toolbar_set_drop_highlight_item(v.native(),
		toolItem.toToolItem(), C.gint(index))
}

// SetShowArrow is a wrapper around gtk_toolbar_set_show_arrow().
func (v *Toolbar) SetShowArrow(showArrow bool) {
	C.gtk_toolbar_set_show_arrow(v.native(), gbool(showArrow))
}

// UnsetIconSize is a wrapper around gtk_toolbar_unset_icon_size().
func (v *Toolbar) UnsetIconSize() {
	C.gtk_toolbar_unset_icon_size(v.native())
}

// GetShowArrow is a wrapper around gtk_toolbar_get_show_arrow().
func (v *Toolbar) GetShowArrow() bool {
	c := C.gtk_toolbar_get_show_arrow(v.native())
	return gobool(c)
}

// GetStyle is a wrapper around gtk_toolbar_get_style().
func (v *Toolbar) GetStyle() ToolbarStyle {
	c := C.gtk_toolbar_get_style(v.native())
	return ToolbarStyle(c)
}

// GetIconSize is a wrapper around gtk_toolbar_get_icon_size().
func (v *Toolbar) GetIconSize() IconSize {
	c := C.gtk_toolbar_get_icon_size(v.native())
	return IconSize(c)
}

// GetReliefStyle is a wrapper around gtk_toolbar_get_relief_style().
func (v *Toolbar) GetReliefStyle() ReliefStyle {
	c := C.gtk_toolbar_get_relief_style(v.native())
	return ReliefStyle(c)
}

// SetStyle is a wrapper around gtk_toolbar_set_style().
func (v *Toolbar) SetStyle(style ToolbarStyle) {
	C.gtk_toolbar_set_style(v.native(), C.GtkToolbarStyle(style))
}

// SetIconSize is a wrapper around gtk_toolbar_set_icon_size().
func (v *Toolbar) SetIconSize(iconSize IconSize) {
	C.gtk_toolbar_set_icon_size(v.native(), C.GtkIconSize(iconSize))
}

// UnsetStyle is a wrapper around gtk_toolbar_unset_style().
func (v *Toolbar) UnsetStyle() {
	C.gtk_toolbar_unset_style(v.native())
}

/*
 * GtkToolButton
 */

// ToolButton is a representation of GTK's GtkToolButton.
type ToolButton struct {
	ToolItem
}

// native returns a pointer to the underlying GtkToolButton.
func (v *ToolButton) native() *C.GtkToolButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkToolButton(p)
}

func marshalToolButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapToolButton(obj), nil
}

func wrapToolButton(obj *glib.Object) *ToolButton {
	return &ToolButton{ToolItem{Bin{Container{Widget{
		glib.InitiallyUnowned{obj}}}}}}
}

// ToolButtonNew is a wrapper around gtk_tool_button_new().
func ToolButtonNew(iconWidget IWidget, label string) (*ToolButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	w := nullableWidget(iconWidget)
	c := C.gtk_tool_button_new(w, (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	tb := wrapToolButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return tb, nil
}

// SetLabel is a wrapper around gtk_tool_button_set_label().
func (v *ToolButton) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tool_button_set_label(v.native(), (*C.gchar)(cstr))
}

// GetLabel is a wrapper aroud gtk_tool_button_get_label().
func (v *ToolButton) GetLabel() string {
	c := C.gtk_tool_button_get_label(v.native())
	return C.GoString((*C.char)(c))
}

// SetUseUnderline is a wrapper around gtk_tool_button_set_use_underline().
func (v *ToolButton) SetGetUnderline(useUnderline bool) {
	C.gtk_tool_button_set_use_underline(v.native(), gbool(useUnderline))
}

// GetUseUnderline is a wrapper around gtk_tool_button_get_use_underline().
func (v *ToolButton) GetuseUnderline() bool {
	c := C.gtk_tool_button_get_use_underline(v.native())
	return gobool(c)
}

// SetIconName is a wrapper around gtk_tool_button_set_icon_name().
func (v *ToolButton) SetIconName(iconName string) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tool_button_set_icon_name(v.native(), (*C.gchar)(cstr))
}

// GetIconName is a wrapper around gtk_tool_button_get_icon_name().
func (v *ToolButton) GetIconName() string {
	c := C.gtk_tool_button_get_icon_name(v.native())
	return C.GoString((*C.char)(c))
}

// SetIconWidget is a wrapper around gtk_tool_button_set_icon_widget().
func (v *ToolButton) SetIconWidget(iconWidget IWidget) {
	C.gtk_tool_button_set_icon_widget(v.native(), iconWidget.toWidget())
}

// GetIconWidget is a wrapper around gtk_tool_button_get_icon_widget().
func (v *ToolButton) GetIconWidget() *Widget {
	c := C.gtk_tool_button_get_icon_widget(v.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w
}

// SetLabelWidget is a wrapper around gtk_tool_button_set_label_widget().
func (v *ToolButton) SetLabelWidget(labelWidget IWidget) {
	C.gtk_tool_button_set_label_widget(v.native(), labelWidget.toWidget())
}

// GetLabelWidget is a wrapper around gtk_tool_button_get_label_widget().
func (v *ToolButton) GetLabelWidget() *Widget {
	c := C.gtk_tool_button_get_label_widget(v.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w
}

/*
 * GtkToolItem
 */

// ToolItem is a representation of GTK's GtkToolItem.
type ToolItem struct {
	Bin
}

// IToolItem is an interface type implemented by all structs embedding
// a ToolItem.  It is meant to be used as an argument type for wrapper
// functions that wrap around a C GTK function taking a GtkToolItem.
type IToolItem interface {
	toToolItem() *C.GtkToolItem
}

// native returns a pointer to the underlying GtkToolItem.
func (v *ToolItem) native() *C.GtkToolItem {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkToolItem(p)
}

func (v *ToolItem) toToolItem() *C.GtkToolItem {
	return v.native()
}

func marshalToolItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapToolItem(obj), nil
}

func wrapToolItem(obj *glib.Object) *ToolItem {
	return &ToolItem{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// ToolItemNew is a wrapper around gtk_tool_item_new().
func ToolItemNew() (*ToolItem, error) {
	c := C.gtk_tool_item_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	ti := wrapToolItem(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return ti, nil
}

// SetHomogeneous is a wrapper around gtk_tool_item_set_homogeneous().
func (v *ToolItem) SetHomogeneous(homogeneous bool) {
	C.gtk_tool_item_set_homogeneous(v.native(), gbool(homogeneous))
}

// GetHomogeneous is a wrapper around gtk_tool_item_get_homogeneous().
func (v *ToolItem) GetHomogeneous() bool {
	c := C.gtk_tool_item_get_homogeneous(v.native())
	return gobool(c)
}

// SetExpand is a wrapper around gtk_tool_item_set_expand().
func (v *ToolItem) SetExpand(expand bool) {
	C.gtk_tool_item_set_expand(v.native(), gbool(expand))
}

// GetExpand is a wrapper around gtk_tool_item_get_expand().
func (v *ToolItem) GetExpand() bool {
	c := C.gtk_tool_item_get_expand(v.native())
	return gobool(c)
}

// SetTooltipText is a wrapper around gtk_tool_item_set_tooltip_text().
func (v *ToolItem) SetTooltipText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tool_item_set_tooltip_text(v.native(), (*C.gchar)(cstr))
}

// SetTooltipMarkup is a wrapper around gtk_tool_item_set_tooltip_markup().
func (v *ToolItem) SetTooltipMarkup(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tool_item_set_tooltip_markup(v.native(), (*C.gchar)(cstr))
}

// SetUseDragWindow is a wrapper around gtk_tool_item_set_use_drag_window().
func (v *ToolItem) SetUseDragWindow(useDragWindow bool) {
	C.gtk_tool_item_set_use_drag_window(v.native(), gbool(useDragWindow))
}

// GetUseDragWindow is a wrapper around gtk_tool_item_get_use_drag_window().
func (v *ToolItem) GetUseDragWindow() bool {
	c := C.gtk_tool_item_get_use_drag_window(v.native())
	return gobool(c)
}

// SetVisibleHorizontal is a wrapper around
// gtk_tool_item_set_visible_horizontal().
func (v *ToolItem) SetVisibleHorizontal(visibleHorizontal bool) {
	C.gtk_tool_item_set_visible_horizontal(v.native(),
		gbool(visibleHorizontal))
}

// GetVisibleHorizontal is a wrapper around
// gtk_tool_item_get_visible_horizontal().
func (v *ToolItem) GetVisibleHorizontal() bool {
	c := C.gtk_tool_item_get_visible_horizontal(v.native())
	return gobool(c)
}

// SetVisibleVertical is a wrapper around gtk_tool_item_set_visible_vertical().
func (v *ToolItem) SetVisibleVertical(visibleVertical bool) {
	C.gtk_tool_item_set_visible_vertical(v.native(), gbool(visibleVertical))
}

// GetVisibleVertical is a wrapper around gtk_tool_item_get_visible_vertical().
func (v *ToolItem) GetVisibleVertical() bool {
	c := C.gtk_tool_item_get_visible_vertical(v.native())
	return gobool(c)
}

// SetIsImportant is a wrapper around gtk_tool_item_set_is_important().
func (v *ToolItem) SetIsImportant(isImportant bool) {
	C.gtk_tool_item_set_is_important(v.native(), gbool(isImportant))
}

// GetIsImportant is a wrapper around gtk_tool_item_get_is_important().
func (v *ToolItem) GetIsImportant() bool {
	c := C.gtk_tool_item_get_is_important(v.native())
	return gobool(c)
}

// TODO: gtk_tool_item_get_ellipsize_mode

// GetIconSize is a wrapper around gtk_tool_item_get_icon_size().
func (v *ToolItem) GetIconSize() IconSize {
	c := C.gtk_tool_item_get_icon_size(v.native())
	return IconSize(c)
}

// GetOrientation is a wrapper around gtk_tool_item_get_orientation().
func (v *ToolItem) GetOrientation() Orientation {
	c := C.gtk_tool_item_get_orientation(v.native())
	return Orientation(c)
}

// GetToolbarStyle is a wrapper around gtk_tool_item_get_toolbar_style().
func (v *ToolItem) gtk_tool_item_get_toolbar_style() ToolbarStyle {
	c := C.gtk_tool_item_get_toolbar_style(v.native())
	return ToolbarStyle(c)
}

// GetReliefStyle is a wrapper around gtk_tool_item_get_relief_style().
func (v *ToolItem) GetReliefStyle() ReliefStyle {
	c := C.gtk_tool_item_get_relief_style(v.native())
	return ReliefStyle(c)
}

// GetTextAlignment is a wrapper around gtk_tool_item_get_text_alignment().
func (v *ToolItem) GetTextAlignment() float32 {
	c := C.gtk_tool_item_get_text_alignment(v.native())
	return float32(c)
}

// GetTextOrientation is a wrapper around gtk_tool_item_get_text_orientation().
func (v *ToolItem) GetTextOrientation() Orientation {
	c := C.gtk_tool_item_get_text_orientation(v.native())
	return Orientation(c)
}

// RetrieveProxyMenuItem is a wrapper around
// gtk_tool_item_retrieve_proxy_menu_item()
func (v *ToolItem) RetrieveProxyMenuItem() *MenuItem {
	c := C.gtk_tool_item_retrieve_proxy_menu_item(v.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	m := wrapMenuItem(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return m
}

// SetProxyMenuItem is a wrapper around gtk_tool_item_set_proxy_menu_item().
func (v *ToolItem) SetProxyMenuItem(menuItemId string, menuItem IMenuItem) {
	cstr := C.CString(menuItemId)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tool_item_set_proxy_menu_item(v.native(), (*C.gchar)(cstr),
		C.toGtkWidget(unsafe.Pointer(menuItem.toMenuItem())))
}

// RebuildMenu is a wrapper around gtk_tool_item_rebuild_menu().
func (v *ToolItem) RebuildMenu() {
	C.gtk_tool_item_rebuild_menu(v.native())
}

// ToolbarReconfigured is a wrapper around gtk_tool_item_toolbar_reconfigured().
func (v *ToolItem) ToolbarReconfigured() {
	C.gtk_tool_item_toolbar_reconfigured(v.native())
}

// TODO: gtk_tool_item_get_text_size_group

/*
 * GtkTreeIter
 */

// TreeIter is a representation of GTK's GtkTreeIter.
type TreeIter struct {
	GtkTreeIter C.GtkTreeIter
}

// native returns a pointer to the underlying GtkTreeIter.
func (v *TreeIter) native() *C.GtkTreeIter {
	if v == nil {
		return nil
	}
	return &v.GtkTreeIter
}

func marshalTreeIter(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*TreeIter)(unsafe.Pointer(c)), nil
}

func (v *TreeIter) free() {
	C.gtk_tree_iter_free(v.native())
}

// Copy() is a wrapper around gtk_tree_iter_copy().
func (v *TreeIter) Copy() (*TreeIter, error) {
	c := C.gtk_tree_iter_copy(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	t := &TreeIter{*c}
	runtime.SetFinalizer(t, (*TreeIter).free)
	return t, nil
}

/*
 * GtkTreeModel
 */

// TreeModel is a representation of GTK's GtkTreeModel GInterface.
type TreeModel struct {
	*glib.Object
}

// ITreeModel is an interface type implemented by all structs
// embedding a TreeModel.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkTreeModel.
type ITreeModel interface {
	toTreeModel() *C.GtkTreeModel
}

// native returns a pointer to the underlying GObject as a GtkTreeModel.
func (v *TreeModel) native() *C.GtkTreeModel {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTreeModel(p)
}

func (v *TreeModel) toTreeModel() *C.GtkTreeModel {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalTreeModel(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapTreeModel(obj), nil
}

func wrapTreeModel(obj *glib.Object) *TreeModel {
	return &TreeModel{obj}
}

// GetFlags() is a wrapper around gtk_tree_model_get_flags().
func (v *TreeModel) GetFlags() TreeModelFlags {
	c := C.gtk_tree_model_get_flags(v.native())
	return TreeModelFlags(c)
}

// GetNColumns() is a wrapper around gtk_tree_model_get_n_columns().
func (v *TreeModel) GetNColumns() int {
	c := C.gtk_tree_model_get_n_columns(v.native())
	return int(c)
}

// GetColumnType() is a wrapper around gtk_tree_model_get_column_type().
func (v *TreeModel) GetColumnType(index int) glib.Type {
	c := C.gtk_tree_model_get_column_type(v.native(), C.gint(index))
	return glib.Type(c)
}

// GetIter() is a wrapper around gtk_tree_model_get_iter().
func (v *TreeModel) GetIter(path *TreePath) (*TreeIter, error) {
	var iter C.GtkTreeIter
	c := C.gtk_tree_model_get_iter(v.native(), &iter, path.native())
	if !gobool(c) {
		return nil, errors.New("Unable to set iterator")
	}
	t := &TreeIter{iter}
	return t, nil
}

// GetIterFromString() is a wrapper around
// gtk_tree_model_get_iter_from_string().
func (v *TreeModel) GetIterFromString(path string) (*TreeIter, error) {
	var iter C.GtkTreeIter
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_tree_model_get_iter_from_string(v.native(), &iter,
		(*C.gchar)(cstr))
	if !gobool(c) {
		return nil, errors.New("Unable to set iterator")
	}
	t := &TreeIter{iter}
	return t, nil
}

// GetIterFirst() is a wrapper around gtk_tree_model_get_iter_first().
func (v *TreeModel) GetIterFirst() (*TreeIter, bool) {
	var iter C.GtkTreeIter
	c := C.gtk_tree_model_get_iter_first(v.native(), &iter)
	if !gobool(c) {
		return nil, false
	}
	t := &TreeIter{iter}
	return t, true
}

// GetPath() is a wrapper around gtk_tree_model_get_path().
func (v *TreeModel) GetPath(iter *TreeIter) (*TreePath, error) {
	c := C.gtk_tree_model_get_path(v.native(), iter.native())
	if c == nil {
		return nil, nilPtrErr
	}
	p := &TreePath{c}
	runtime.SetFinalizer(p, (*TreePath).free)
	return p, nil
}

// GetValue() is a wrapper around gtk_tree_model_get_value().
func (v *TreeModel) GetValue(iter *TreeIter, column int) (*glib.Value, error) {
	val, err := glib.ValueAlloc()
	if err != nil {
		return nil, err
	}
	C.gtk_tree_model_get_value(
		(*C.GtkTreeModel)(unsafe.Pointer(v.native())),
		iter.native(),
		C.gint(column),
		(*C.GValue)(unsafe.Pointer(val.Native())))
	return val, nil
}

// IterNext() is a wrapper around gtk_tree_model_iter_next().
func (v *TreeModel) IterNext(iter *TreeIter) bool {
	c := C.gtk_tree_model_iter_next(v.native(), iter.native())
	return gobool(c)
}

// IterPrevious is a wrapper around gtk_tree_model_iter_previous().
func (v *TreeModel) IterPrevious(iter *TreeIter) bool {
	c := C.gtk_tree_model_iter_previous(v.native(), iter.native())
	return gobool(c)
}

// IterNChildren is a wrapper around gtk_tree_model_iter_n_children().
func (v *TreeModel) IterNChildren(iter *TreeIter) int {
	var cIter *C.GtkTreeIter
	if iter != nil {
		cIter = iter.native()
	}
	c := C.gtk_tree_model_iter_n_children(v.native(), cIter)
	return int(c)
}

/*
 * GtkTreePath
 */

// TreePath is a representation of GTK's GtkTreePath.
type TreePath struct {
	GtkTreePath *C.GtkTreePath
}

// Return a TreePath from the GList
func TreePathFromList(list *glib.List) *TreePath {
	if list == nil {
		return nil
	}
	return &TreePath{(*C.GtkTreePath)(unsafe.Pointer(list.Data()))}
}

// native returns a pointer to the underlying GtkTreePath.
func (v *TreePath) native() *C.GtkTreePath {
	if v == nil {
		return nil
	}
	return v.GtkTreePath
}

func marshalTreePath(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &TreePath{(*C.GtkTreePath)(unsafe.Pointer(c))}, nil
}

func (v *TreePath) free() {
	C.gtk_tree_path_free(v.native())
}

// String is a wrapper around gtk_tree_path_to_string().
func (v *TreePath) String() string {
	c := C.gtk_tree_path_to_string(v.native())
	return C.GoString((*C.char)(c))
}

// TreePathNewFromString is a wrapper around gtk_tree_path_new_from_string().
func TreePathNewFromString(path string) (*TreePath, error) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_tree_path_new_from_string((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	t := &TreePath{c}
	runtime.SetFinalizer(t, (*TreePath).free)
	return t, nil
}

/*
 * GtkTreeSelection
 */

// TreeSelection is a representation of GTK's GtkTreeSelection.
type TreeSelection struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkTreeSelection.
func (v *TreeSelection) native() *C.GtkTreeSelection {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTreeSelection(p)
}

func marshalTreeSelection(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapTreeSelection(obj), nil
}

func wrapTreeSelection(obj *glib.Object) *TreeSelection {
	return &TreeSelection{obj}
}

// GetSelected() is a wrapper around gtk_tree_selection_get_selected().
func (v *TreeSelection) GetSelected(model *ITreeModel, iter *TreeIter) bool {
	var pcmodel **C.GtkTreeModel
	if pcmodel != nil {
		cmodel := (*model).toTreeModel()
		pcmodel = &cmodel
	} else {
		pcmodel = nil
	}
	c := C.gtk_tree_selection_get_selected(v.native(),
		pcmodel, iter.native())
	return gobool(c)
}

// SelectPath is a wrapper around gtk_tree_selection_select_path().
func (v *TreeSelection) SelectPath(path *TreePath) {
	C.gtk_tree_selection_select_path(v.native(), path.native())
}

// GetSelectedRows is a wrapper around gtk_tree_selection_get_selected_rows().
//
// Please note that a runtime finalizer is only set on the head of the linked
// list, and must be kept live while accessing any item in the list, or the
// Go garbage collector will free the whole list.
func (v *TreeSelection) GetSelectedRows(model ITreeModel) *glib.List {
	var pcmodel **C.GtkTreeModel
	if model != nil {
		cmodel := model.toTreeModel()
		pcmodel = &cmodel
	}
	clist := C.gtk_tree_selection_get_selected_rows(v.native(), pcmodel)
	glist := (*glib.List)(unsafe.Pointer(clist))
	runtime.SetFinalizer(glist, func(list *glib.List) {
		glist := (*C.GList)(unsafe.Pointer(list))
		C.g_list_free_full((*C.GList)(glist), (C.GDestroyNotify)(C.gtk_tree_path_free))
	})
	return glist
}

// CountSelectedRows() is a wrapper around gtk_tree_selection_count_selected_rows().
func (v *TreeSelection) CountSelectedRows() int {
	return int(C.gtk_tree_selection_count_selected_rows(v.native()))
}

// SelectIter is a wrapper around gtk_tree_selection_select_iter().
func (v *TreeSelection) SelectIter(iter *TreeIter) {
	C.gtk_tree_selection_select_iter(v.native(), iter.native())
}

// SetMode() is a wrapper around gtk_tree_selection_set_mode().
func (v *TreeSelection) SetMode(m SelectionMode) {
	C.gtk_tree_selection_set_mode(v.native(), C.GtkSelectionMode(m))
}

// GetMode() is a wrapper around gtk_tree_selection_get_mode().
func (v *TreeSelection) GetMode() SelectionMode {
	return SelectionMode(C.gtk_tree_selection_get_mode(v.native()))
}

/*
 * GtkTreeStore
 */

// TreeStore is a representation of GTK's GtkTreeStore.
type TreeStore struct {
	*glib.Object

	// Interfaces
	TreeModel
}

// native returns a pointer to the underlying GtkTreeStore.
func (v *TreeStore) native() *C.GtkTreeStore {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTreeStore(p)
}

func marshalTreeStore(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapTreeStore(obj), nil
}

func wrapTreeStore(obj *glib.Object) *TreeStore {
	tm := wrapTreeModel(obj)
	return &TreeStore{obj, *tm}
}

func (v *TreeStore) toTreeModel() *C.GtkTreeModel {
	if v == nil {
		return nil
	}
	return C.toGtkTreeModel(unsafe.Pointer(v.GObject))
}

// TreeStoreNew is a wrapper around gtk_tree_store_newv().
func TreeStoreNew(types ...glib.Type) (*TreeStore, error) {
	gtypes := C.alloc_types(C.int(len(types)))
	for n, val := range types {
		C.set_type(gtypes, C.int(n), C.GType(val))
	}
	defer C.g_free(C.gpointer(gtypes))
	c := C.gtk_tree_store_newv(C.gint(len(types)), gtypes)
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	ts := wrapTreeStore(obj)
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return ts, nil
}

// Append is a wrapper around gtk_tree_store_append().
func (v *TreeStore) Append(parent *TreeIter) *TreeIter {
	var ti C.GtkTreeIter
	var cParent *C.GtkTreeIter
	if parent != nil {
		cParent = parent.native()
	}
	C.gtk_tree_store_append(v.native(), &ti, cParent)
	iter := &TreeIter{ti}
	return iter
}

// SetValue is a wrapper around gtk_tree_store_set_value()
func (v *TreeStore) SetValue(iter *TreeIter, column int, value interface{}) error {
	gv, err := glib.GValue(value)
	if err != nil {
		return err
	}
	C.gtk_tree_store_set_value(v.native(), iter.native(),
		C.gint(column),
		(*C.GValue)(C.gpointer(gv.Native())))
	return nil
}

// Clear is a wrapper around gtk_tree_store_clear().
func (v *TreeStore) Clear() {
	C.gtk_tree_store_clear(v.native())
}

/*
 * GtkTreeViewColumn
 */

// TreeViewColumns is a representation of GTK's GtkTreeViewColumn.
type TreeViewColumn struct {
	glib.InitiallyUnowned
}

// native returns a pointer to the underlying GtkTreeViewColumn.
func (v *TreeViewColumn) native() *C.GtkTreeViewColumn {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTreeViewColumn(p)
}

func marshalTreeViewColumn(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapTreeViewColumn(obj), nil
}

func wrapTreeViewColumn(obj *glib.Object) *TreeViewColumn {
	return &TreeViewColumn{glib.InitiallyUnowned{obj}}
}

// TreeViewColumnNew() is a wrapper around gtk_tree_view_column_new().
func TreeViewColumnNew() (*TreeViewColumn, error) {
	c := C.gtk_tree_view_column_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	t := wrapTreeViewColumn(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return t, nil
}

// TreeViewColumnNewWithAttribute() is a wrapper around
// gtk_tree_view_column_new_with_attributes() that only sets one
// attribute for one column.
func TreeViewColumnNewWithAttribute(title string, renderer ICellRenderer, attribute string, column int) (*TreeViewColumn, error) {
	t_cstr := C.CString(title)
	defer C.free(unsafe.Pointer(t_cstr))
	a_cstr := C.CString(attribute)
	defer C.free(unsafe.Pointer(a_cstr))
	c := C._gtk_tree_view_column_new_with_attributes_one((*C.gchar)(t_cstr),
		renderer.toCellRenderer(), (*C.gchar)(a_cstr), C.gint(column))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	t := wrapTreeViewColumn(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return t, nil
}

// AddAttribute() is a wrapper around gtk_tree_view_column_add_attribute().
func (v *TreeViewColumn) AddAttribute(renderer ICellRenderer, attribute string, column int) {
	cstr := C.CString(attribute)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tree_view_column_add_attribute(v.native(),
		renderer.toCellRenderer(), (*C.gchar)(cstr), C.gint(column))
}

// SetExpand() is a wrapper around gtk_tree_view_column_set_expand().
func (v *TreeViewColumn) SetExpand(expand bool) {
	C.gtk_tree_view_column_set_expand(v.native(), gbool(expand))
}

// GetExpand() is a wrapper around gtk_tree_view_column_get_expand().
func (v *TreeViewColumn) GetExpand() bool {
	c := C.gtk_tree_view_column_get_expand(v.native())
	return gobool(c)
}

// SetMinWidth() is a wrapper around gtk_tree_view_column_set_min_width().
func (v *TreeViewColumn) SetMinWidth(minWidth int) {
	C.gtk_tree_view_column_set_min_width(v.native(), C.gint(minWidth))
}

// GetMinWidth() is a wrapper around gtk_tree_view_column_get_min_width().
func (v *TreeViewColumn) GetMinWidth() int {
	c := C.gtk_tree_view_column_get_min_width(v.native())
	return int(c)
}

/*
 * GtkViewport
 */

// Viewport is a representation of GTK's GtkViewport GInterface.
type Viewport struct {
	Bin

	// Interfaces
	Scrollable
}

// IViewport is an interface type implemented by all structs
// embedding a Viewport.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkViewport.
type IViewport interface {
	toViewport() *C.GtkViewport
}

// native() returns a pointer to the underlying GObject as a GtkViewport.
func (v *Viewport) native() *C.GtkViewport {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkViewport(p)
}

func wrapViewport(obj *glib.Object) *Viewport {
	b := wrapBin(obj)
	s := wrapScrollable(obj)
	return &Viewport{
		Bin:        *b,
		Scrollable: *s,
	}
}

func (v *Viewport) toViewport() *C.GtkViewport {
	if v == nil {
		return nil
	}
	return v.native()
}

// ViewportNew() is a wrapper around gtk_viewport_new().
func ViewportNew(hadjustment, vadjustment *Adjustment) (*Viewport, error) {
	c := C.gtk_viewport_new(hadjustment.native(), vadjustment.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapViewport(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return b, nil
}

func (v *Viewport) SetHAdjustment(adjustment *Adjustment) {
	wrapScrollable(v.Object).SetHAdjustment(adjustment)
}

func (v *Viewport) GetHAdjustment() (*Adjustment, error) {
	return wrapScrollable(v.Object).GetHAdjustment()
}

func (v *Viewport) SetVAdjustment(adjustment *Adjustment) {
	wrapScrollable(v.Object).SetVAdjustment(adjustment)
}

func (v *Viewport) GetVAdjustment() (*Adjustment, error) {
	return wrapScrollable(v.Object).GetVAdjustment()
}

/*
 * GtkVolumeButton
 */

// VolumeButton is a representation of GTK's GtkVolumeButton.
type VolumeButton struct {
	ScaleButton
}

// native() returns a pointer to the underlying GtkVolumeButton.
func (v *VolumeButton) native() *C.GtkVolumeButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkVolumeButton(p)
}

func marshalVolumeButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapVolumeButton(obj), nil
}

func wrapVolumeButton(obj *glib.Object) *VolumeButton {
	return &VolumeButton{ScaleButton{Button{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}}}
}

// VolumeButtonNew() is a wrapper around gtk_button_new().
func VolumeButtonNew() (*VolumeButton, error) {
	c := C.gtk_volume_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapVolumeButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return b, nil
}

/*
 * GtkWidget
 */

// Widget is a representation of GTK's GtkWidget.
type Widget struct {
	glib.InitiallyUnowned
}

// IWidget is an interface type implemented by all structs
// embedding a Widget.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkWidget.
type IWidget interface {
	toWidget() *C.GtkWidget
	Set(string, interface{}) error
}

type IWidgetable interface {
	toWidget() *C.GtkWidget
}

func nullableWidget(v IWidgetable) *C.GtkWidget {
	if v == nil {
		return nil
	}

	return v.toWidget()
}

// native returns a pointer to the underlying GtkWidget.
func (v *Widget) native() *C.GtkWidget {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkWidget(p)
}

func (v *Widget) toWidget() *C.GtkWidget {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalWidget(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapWidget(obj), nil
}

func wrapWidget(obj *glib.Object) *Widget {
	return &Widget{glib.InitiallyUnowned{obj}}
}

// Destroy is a wrapper around gtk_widget_destroy().
func (v *Widget) Destroy() {
	C.gtk_widget_destroy(v.native())
}

func (v *Widget) HideOnDelete() {
	C._gtk_widget_hide_on_delete(v.native())
}

/* TODO
func (v *Widget) DragDestSet(flags DestDefaults, targets []TargetEntry, actions gdk.DragAction) {
	C.gtk_drag_dest_set(v.native(), C.GtkDestDefaults(flags), (*C.GtkTargetEntry)(&targets[0]),
		C.gint(len(targets)), C.GdkDragAction(actions))
}
*/

// ResetStyle is a wrapper around gtk_widget_reset_style().
func (v *Widget) ResetStyle() {
	C.gtk_widget_reset_style(v.native())
}

// InDestruction is a wrapper around gtk_widget_in_destruction().
func (v *Widget) InDestruction() bool {
	return gobool(C.gtk_widget_in_destruction(v.native()))
}

// TODO(jrick) this may require some rethinking
/*
func (v *Widget) Destroyed(widgetPointer **Widget) {
}
*/

// Unparent is a wrapper around gtk_widget_unparent().
func (v *Widget) Unparent() {
	C.gtk_widget_unparent(v.native())
}

// Show is a wrapper around gtk_widget_show().
func (v *Widget) Show() {
	C.gtk_widget_show(v.native())
}

// Hide is a wrapper around gtk_widget_hide().
func (v *Widget) Hide() {
	C.gtk_widget_hide(v.native())
}

// GetCanFocus is a wrapper around gtk_widget_get_can_focus().
func (v *Widget) GetCanFocus() bool {
	c := C.gtk_widget_get_can_focus(v.native())
	return gobool(c)
}

// SetCanFocus is a wrapper around gtk_widget_set_can_focus().
func (v *Widget) SetCanFocus(canFocus bool) {
	C.gtk_widget_set_can_focus(v.native(), gbool(canFocus))
}

// GetMapped is a wrapper around gtk_window_get_mapped().
func (v *Widget) GetMapped() bool {
	c := C.gtk_widget_get_mapped(v.native())
	return gobool(c)
}

// SetMapped is a wrapper around gtk_widget_set_mapped().
func (v *Widget) SetMapped(mapped bool) {
	C.gtk_widget_set_can_focus(v.native(), gbool(mapped))
}

// GetRealized is a wrapper around gtk_window_get_realized().
func (v *Widget) GetRealized() bool {
	c := C.gtk_widget_get_realized(v.native())
	return gobool(c)
}

// SetRealized is a wrapper around gtk_widget_set_realized().
func (v *Widget) SetRealized(realized bool) {
	C.gtk_widget_set_realized(v.native(), gbool(realized))
}

// GetHasWindow is a wrapper around gtk_widget_get_has_window().
func (v *Widget) GetHasWindow() bool {
	c := C.gtk_widget_get_has_window(v.native())
	return gobool(c)
}

// SetHasWindow is a wrapper around gtk_widget_set_has_window().
func (v *Widget) SetHasWindow(hasWindow bool) {
	C.gtk_widget_set_has_window(v.native(), gbool(hasWindow))
}

// ShowNow is a wrapper around gtk_widget_show_now().
func (v *Widget) ShowNow() {
	C.gtk_widget_show_now(v.native())
}

// ShowAll is a wrapper around gtk_widget_show_all().
func (v *Widget) ShowAll() {
	C.gtk_widget_show_all(v.native())
}

// SetNoShowAll is a wrapper around gtk_widget_set_no_show_all().
func (v *Widget) SetNoShowAll(noShowAll bool) {
	C.gtk_widget_set_no_show_all(v.native(), gbool(noShowAll))
}

// GetNoShowAll is a wrapper around gtk_widget_get_no_show_all().
func (v *Widget) GetNoShowAll() bool {
	c := C.gtk_widget_get_no_show_all(v.native())
	return gobool(c)
}

// Map is a wrapper around gtk_widget_map().
func (v *Widget) Map() {
	C.gtk_widget_map(v.native())
}

// Unmap is a wrapper around gtk_widget_unmap().
func (v *Widget) Unmap() {
	C.gtk_widget_unmap(v.native())
}

//void gtk_widget_realize(GtkWidget *widget);
//void gtk_widget_unrealize(GtkWidget *widget);
//void gtk_widget_draw(GtkWidget *widget, cairo_t *cr);
//void gtk_widget_queue_resize(GtkWidget *widget);
//void gtk_widget_queue_resize_no_redraw(GtkWidget *widget);
//GdkFrameClock *gtk_widget_get_frame_clock(GtkWidget *widget);
//guint gtk_widget_add_tick_callback (GtkWidget *widget,
//                                    GtkTickCallback callback,
//                                    gpointer user_data,
//                                    GDestroyNotify notify);
//void gtk_widget_remove_tick_callback(GtkWidget *widget, guint id);

// TODO(jrick) GtkAllocation
/*
func (v *Widget) SizeAllocate() {
}
*/

// Allocation is a representation of GTK's GtkAllocation type.
type Allocation struct {
	gdk.Rectangle
}

// Native returns a pointer to the underlying GtkAllocation.
func (v *Allocation) native() *C.GtkAllocation {
	return (*C.GtkAllocation)(unsafe.Pointer(&v.GdkRectangle))
}

// GetAllocatedWidth() is a wrapper around gtk_widget_get_allocated_width().
func (v *Widget) GetAllocatedWidth() int {
	return int(C.gtk_widget_get_allocated_width(v.native()))
}

// GetAllocatedHeight() is a wrapper around gtk_widget_get_allocated_height().
func (v *Widget) GetAllocatedHeight() int {
	return int(C.gtk_widget_get_allocated_height(v.native()))
}

// Event() is a wrapper around gtk_widget_event().
func (v *Widget) Event(event *gdk.Event) bool {
	c := C.gtk_widget_event(v.native(),
		(*C.GdkEvent)(unsafe.Pointer(event.Native())))
	return gobool(c)
}

// Activate() is a wrapper around gtk_widget_activate().
func (v *Widget) Activate() bool {
	return gobool(C.gtk_widget_activate(v.native()))
}

// TODO(jrick) GdkRectangle
/*
func (v *Widget) Intersect() {
}
*/

// IsFocus() is a wrapper around gtk_widget_is_focus().
func (v *Widget) IsFocus() bool {
	return gobool(C.gtk_widget_is_focus(v.native()))
}

// GrabFocus() is a wrapper around gtk_widget_grab_focus().
func (v *Widget) GrabFocus() {
	C.gtk_widget_grab_focus(v.native())
}

// GrabDefault() is a wrapper around gtk_widget_grab_default().
func (v *Widget) GrabDefault() {
	C.gtk_widget_grab_default(v.native())
}

// SetName() is a wrapper around gtk_widget_set_name().
func (v *Widget) SetName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_widget_set_name(v.native(), (*C.gchar)(cstr))
}

// GetName() is a wrapper around gtk_widget_get_name().  A non-nil
// error is returned in the case that gtk_widget_get_name returns NULL to
// differentiate between NULL and an empty string.
func (v *Widget) GetName() (string, error) {
	c := C.gtk_widget_get_name(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// GetSensitive is a wrapper around gtk_widget_get_sensitive().
func (v *Widget) GetSensitive() bool {
	c := C.gtk_widget_get_sensitive(v.native())
	return gobool(c)
}

// IsSensitive is a wrapper around gtk_widget_is_sensitive().
func (v *Widget) IsSensitive() bool {
	c := C.gtk_widget_is_sensitive(v.native())
	return gobool(c)
}

// SetSensitive is a wrapper around gtk_widget_set_sensitive().
func (v *Widget) SetSensitive(sensitive bool) {
	C.gtk_widget_set_sensitive(v.native(), gbool(sensitive))
}

// GetVisible is a wrapper around gtk_widget_get_visible().
func (v *Widget) GetVisible() bool {
	c := C.gtk_widget_get_visible(v.native())
	return gobool(c)
}

// SetVisible is a wrapper around gtk_widget_set_visible().
func (v *Widget) SetVisible(visible bool) {
	C.gtk_widget_set_visible(v.native(), gbool(visible))
}

// SetParent is a wrapper around gtk_widget_set_parent().
func (v *Widget) SetParent(parent IWidget) {
	C.gtk_widget_set_parent(v.native(), parent.toWidget())
}

// GetParent is a wrapper around gtk_widget_get_parent().
func (v *Widget) GetParent() (*Widget, error) {
	c := C.gtk_widget_get_parent(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// SetSizeRequest is a wrapper around gtk_widget_set_size_request().
func (v *Widget) SetSizeRequest(width, height int) {
	C.gtk_widget_set_size_request(v.native(), C.gint(width), C.gint(height))
}

// GetSizeRequest is a wrapper around gtk_widget_get_size_request().
func (v *Widget) GetSizeRequest() (width, height int) {
	var w, h C.gint
	C.gtk_widget_get_size_request(v.native(), &w, &h)
	return int(w), int(h)
}

// SetParentWindow is a wrapper around gtk_widget_set_parent_window().
func (v *Widget) SetParentWindow(parentWindow *gdk.Window) {
	C.gtk_widget_set_parent_window(v.native(),
		(*C.GdkWindow)(unsafe.Pointer(parentWindow.Native())))
}

// GetParentWindow is a wrapper around gtk_widget_get_parent_window().
func (v *Widget) GetParentWindow() (*gdk.Window, error) {
	c := C.gtk_widget_get_parent_window(v.native())
	if v == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := &gdk.Window{obj}
	w.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// SetEvents is a wrapper around gtk_widget_set_events().
func (v *Widget) SetEvents(events int) {
	C.gtk_widget_set_events(v.native(), C.gint(events))
}

// GetEvents is a wrapper around gtk_widget_get_events().
func (v *Widget) GetEvents() int {
	return int(C.gtk_widget_get_events(v.native()))
}

// AddEvents is a wrapper around gtk_widget_add_events().
func (v *Widget) AddEvents(events int) {
	C.gtk_widget_add_events(v.native(), C.gint(events))
}

// HasDefault is a wrapper around gtk_widget_has_default().
func (v *Widget) HasDefault() bool {
	c := C.gtk_widget_has_default(v.native())
	return gobool(c)
}

// HasFocus is a wrapper around gtk_widget_has_focus().
func (v *Widget) HasFocus() bool {
	c := C.gtk_widget_has_focus(v.native())
	return gobool(c)
}

// HasVisibleFocus is a wrapper around gtk_widget_has_visible_focus().
func (v *Widget) HasVisibleFocus() bool {
	c := C.gtk_widget_has_visible_focus(v.native())
	return gobool(c)
}

// HasGrab is a wrapper around gtk_widget_has_grab().
func (v *Widget) HasGrab() bool {
	c := C.gtk_widget_has_grab(v.native())
	return gobool(c)
}

// IsDrawable is a wrapper around gtk_widget_is_drawable().
func (v *Widget) IsDrawable() bool {
	c := C.gtk_widget_is_drawable(v.native())
	return gobool(c)
}

// IsToplevel is a wrapper around gtk_widget_is_toplevel().
func (v *Widget) IsToplevel() bool {
	c := C.gtk_widget_is_toplevel(v.native())
	return gobool(c)
}

// TODO(jrick) GdkEventMask
/*
func (v *Widget) SetDeviceEvents() {
}
*/

// TODO(jrick) GdkEventMask
/*
func (v *Widget) GetDeviceEvents() {
}
*/

// TODO(jrick) GdkEventMask
/*
func (v *Widget) AddDeviceEvents() {
}
*/

// SetDeviceEnabled is a wrapper around gtk_widget_set_device_enabled().
func (v *Widget) SetDeviceEnabled(device *gdk.Device, enabled bool) {
	C.gtk_widget_set_device_enabled(v.native(),
		(*C.GdkDevice)(unsafe.Pointer(device.Native())), gbool(enabled))
}

// GetDeviceEnabled is a wrapper around gtk_widget_get_device_enabled().
func (v *Widget) GetDeviceEnabled(device *gdk.Device) bool {
	c := C.gtk_widget_get_device_enabled(v.native(),
		(*C.GdkDevice)(unsafe.Pointer(device.Native())))
	return gobool(c)
}

// GetToplevel is a wrapper around gtk_widget_get_toplevel().
func (v *Widget) GetToplevel() (*Widget, error) {
	c := C.gtk_widget_get_toplevel(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// GetTooltipText is a wrapper around gtk_widget_get_tooltip_text().
// A non-nil error is returned in the case that
// gtk_widget_get_tooltip_text returns NULL to differentiate between NULL
// and an empty string.
func (v *Widget) GetTooltipText() (string, error) {
	c := C.gtk_widget_get_tooltip_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// SetTooltipText is a wrapper around gtk_widget_set_tooltip_text().
func (v *Widget) SetTooltipText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_widget_set_tooltip_text(v.native(), (*C.gchar)(cstr))
}

// GetHAlign is a wrapper around gtk_widget_get_halign().
func (v *Widget) GetHAlign() Align {
	c := C.gtk_widget_get_halign(v.native())
	return Align(c)
}

// SetHAlign is a wrapper around gtk_widget_set_halign().
func (v *Widget) SetHAlign(align Align) {
	C.gtk_widget_set_halign(v.native(), C.GtkAlign(align))
}

// GetVAlign is a wrapper around gtk_widget_get_valign().
func (v *Widget) GetVAlign() Align {
	c := C.gtk_widget_get_valign(v.native())
	return Align(c)
}

// SetVAlign is a wrapper around gtk_widget_set_valign().
func (v *Widget) SetVAlign(align Align) {
	C.gtk_widget_set_valign(v.native(), C.GtkAlign(align))
}

// GetMarginTop is a wrapper around gtk_widget_get_margin_top().
func (v *Widget) GetMarginTop() int {
	c := C.gtk_widget_get_margin_top(v.native())
	return int(c)
}

// SetMarginTop is a wrapper around gtk_widget_set_margin_top().
func (v *Widget) SetMarginTop(margin int) {
	C.gtk_widget_set_margin_top(v.native(), C.gint(margin))
}

// GetMarginBottom is a wrapper around gtk_widget_get_margin_bottom().
func (v *Widget) GetMarginBottom() int {
	c := C.gtk_widget_get_margin_bottom(v.native())
	return int(c)
}

// SetMarginBottom is a wrapper around gtk_widget_set_margin_bottom().
func (v *Widget) SetMarginBottom(margin int) {
	C.gtk_widget_set_margin_bottom(v.native(), C.gint(margin))
}

// GetHExpand is a wrapper around gtk_widget_get_hexpand().
func (v *Widget) GetHExpand() bool {
	c := C.gtk_widget_get_hexpand(v.native())
	return gobool(c)
}

// SetHExpand is a wrapper around gtk_widget_set_hexpand().
func (v *Widget) SetHExpand(expand bool) {
	C.gtk_widget_set_hexpand(v.native(), gbool(expand))
}

// GetVExpand is a wrapper around gtk_widget_get_vexpand().
func (v *Widget) GetVExpand() bool {
	c := C.gtk_widget_get_vexpand(v.native())
	return gobool(c)
}

// SetVExpand is a wrapper around gtk_widget_set_vexpand().
func (v *Widget) SetVExpand(expand bool) {
	C.gtk_widget_set_vexpand(v.native(), gbool(expand))
}

// TranslateCoordinates is a wrapper around gtk_widget_translate_coordinates().
func (v *Widget) TranslateCoordinates(dest IWidget, srcX, srcY int) (destX, destY int, e error) {
	cdest := nullableWidget(dest)

	var cdestX, cdestY C.gint
	c := C.gtk_widget_translate_coordinates(v.native(), cdest, C.gint(srcX), C.gint(srcY), &cdestX, &cdestY)
	if !gobool(c) {
		return 0, 0, errors.New("translate coordinates failed")
	}
	return int(cdestX), int(cdestY), nil
}

// SetVisual is a wrapper around gtk_widget_set_visual().
func (v *Widget) SetVisual(visual *gdk.Visual) {
	C.gtk_widget_set_visual(v.native(),
		(*C.GdkVisual)(unsafe.Pointer(visual.Native())))
}

// SetAppPaintable is a wrapper around gtk_widget_set_app_paintable().
func (v *Widget) SetAppPaintable(paintable bool) {
	C.gtk_widget_set_app_paintable(v.native(), gbool(paintable))
}

// GetAppPaintable is a wrapper around gtk_widget_get_app_paintable().
func (v *Widget) GetAppPaintable() bool {
	c := C.gtk_widget_get_app_paintable(v.native())
	return gobool(c)
}

// QueueDraw is a wrapper around gtk_widget_queue_draw().
func (v *Widget) QueueDraw() {
	C.gtk_widget_queue_draw(v.native())
}

// GetAllocation is a wrapper around gtk_widget_get_allocation().
func (v *Widget) GetAllocation() *Allocation {
	var a Allocation
	C.gtk_widget_get_allocation(v.native(), a.native())
	return &a
}

// SetAllocation is a wrapper around gtk_widget_set_allocation().
func (v *Widget) SetAllocation(allocation *Allocation) {
	C.gtk_widget_set_allocation(v.native(), allocation.native())
}

// SizeAllocate is a wrapper around gtk_widget_size_allocate().
func (v *Widget) SizeAllocate(allocation *Allocation) {
	C.gtk_widget_size_allocate(v.native(), allocation.native())
}

// SetStateFlags is a wrapper around gtk_widget_set_state_flags().
func (v *Widget) SetStateFlags(stateFlags StateFlags, clear bool) {
	C.gtk_widget_set_state_flags(v.native(), C.GtkStateFlags(stateFlags), gbool(clear))
}

/*
 * GtkWindow
 */

// Window is a representation of GTK's GtkWindow.
type Window struct {
	Bin
}

// IWindow is an interface type implemented by all structs embedding a
// Window.  It is meant to be used as an argument type for wrapper
// functions that wrap around a C GTK function taking a GtkWindow.
type IWindow interface {
	toWindow() *C.GtkWindow
}

// native returns a pointer to the underlying GtkWindow.
func (v *Window) native() *C.GtkWindow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkWindow(p)
}

func (v *Window) toWindow() *C.GtkWindow {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalWindow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapWindow(obj), nil
}

func wrapWindow(obj *glib.Object) *Window {
	return &Window{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// WindowNew is a wrapper around gtk_window_new().
func WindowNew(t WindowType) (*Window, error) {
	c := C.gtk_window_new(C.GtkWindowType(t))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWindow(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// SetTitle is a wrapper around gtk_window_set_title().
func (v *Window) SetTitle(title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_window_set_title(v.native(), (*C.gchar)(cstr))
}

// SetResizable is a wrapper around gtk_window_set_resizable().
func (v *Window) SetResizable(resizable bool) {
	C.gtk_window_set_resizable(v.native(), gbool(resizable))
}

// GetResizable is a wrapper around gtk_window_get_resizable().
func (v *Window) GetResizable() bool {
	c := C.gtk_window_get_resizable(v.native())
	return gobool(c)
}

// ActivateFocus is a wrapper around gtk_window_activate_focus().
func (v *Window) ActivateFocus() bool {
	c := C.gtk_window_activate_focus(v.native())
	return gobool(c)
}

// ActivateDefault is a wrapper around gtk_window_activate_default().
func (v *Window) ActivateDefault() bool {
	c := C.gtk_window_activate_default(v.native())
	return gobool(c)
}

// SetModal is a wrapper around gtk_window_set_modal().
func (v *Window) SetModal(modal bool) {
	C.gtk_window_set_modal(v.native(), gbool(modal))
}

// SetDefaultSize is a wrapper around gtk_window_set_default_size().
func (v *Window) SetDefaultSize(width, height int) {
	C.gtk_window_set_default_size(v.native(), C.gint(width), C.gint(height))
}

// SetDefaultGeometry is a wrapper around gtk_window_set_default_geometry().
func (v *Window) SetDefaultGeometry(width, height int) {
	C.gtk_window_set_default_geometry(v.native(), C.gint(width),
		C.gint(height))
}

// GetScreen is a wrapper around gtk_window_get_screen().
func (v *Window) GetScreen() (*gdk.Screen, error) {
	c := C.gtk_window_get_screen(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := &gdk.Screen{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

// SetIcon is a wrapper around gtk_window_set_icon().
func (v *Window) SetIcon(icon *gdk.Pixbuf) {
	iconPtr := (*C.GdkPixbuf)(unsafe.Pointer(icon.Native()))
	C.gtk_window_set_icon(v.native(), iconPtr)
}

// TODO(jrick) GdkGeometry GdkWindowHints.
/*
func (v *Window) SetGeometryHints() {
}
*/

// TODO(jrick) GdkGravity.
/*
func (v *Window) SetGravity() {
}
*/

// TODO(jrick) GdkGravity.
/*
func (v *Window) GetGravity() {
}
*/

// SetPosition is a wrapper around gtk_window_set_position().
func (v *Window) SetPosition(position WindowPosition) {
	C.gtk_window_set_position(v.native(), C.GtkWindowPosition(position))
}

// SetTransientFor is a wrapper around gtk_window_set_transient_for().
func (v *Window) SetTransientFor(parent IWindow) {
	var pw *C.GtkWindow = nil
	if parent != nil {
		pw = parent.toWindow()
	}
	C.gtk_window_set_transient_for(v.native(), pw)
}

// TODO gtk_window_set_attached_to().

// SetDestroyWithParent is a wrapper around
// gtk_window_set_destroy_with_parent().
func (v *Window) SetDestroyWithParent(setting bool) {
	C.gtk_window_set_destroy_with_parent(v.native(), gbool(setting))
}

// SetHideTitlebarWhenMaximized is a wrapper around
// gtk_window_set_hide_titlebar_when_maximized().
func (v *Window) SetHideTitlebarWhenMaximized(setting bool) {
	C.gtk_window_set_hide_titlebar_when_maximized(v.native(),
		gbool(setting))
}

// TODO gtk_window_set_screen().

// IsActive is a wrapper around gtk_window_is_active().
func (v *Window) IsActive() bool {
	c := C.gtk_window_is_active(v.native())
	return gobool(c)
}

// HasToplevelFocus is a wrapper around gtk_window_has_toplevel_focus().
func (v *Window) HasToplevelFocus() bool {
	c := C.gtk_window_has_toplevel_focus(v.native())
	return gobool(c)
}

// TODO gtk_window_list_toplevels().

// TODO gtk_window_add_mnemonic().

// TODO gtk_window_remove_mnemonic().

// TODO gtk_window_mnemonic_activate().

// TODO gtk_window_activate_key().

// TODO gtk_window_propogate_key_event().

// TODO gtk_window_get_focus().

// TODO gtk_window_set_focus().

// TODO gtk_window_get_default_widget().

// TODO gtk_window_set_default().

// Present is a wrapper around gtk_window_present().
func (v *Window) Present() {
	C.gtk_window_present(v.native())
}

// PresentWithTime is a wrapper around gtk_window_present_with_time().
func (v *Window) PresentWithTime(ts uint32) {
	C.gtk_window_present_with_time(v.native(), C.guint32(ts))
}

// Iconify is a wrapper around gtk_window_iconify().
func (v *Window) Iconify() {
	C.gtk_window_iconify(v.native())
}

// Deiconify is a wrapper around gtk_window_deiconify().
func (v *Window) Deiconify() {
	C.gtk_window_deiconify(v.native())
}

// Stick is a wrapper around gtk_window_stick().
func (v *Window) Stick() {
	C.gtk_window_stick(v.native())
}

// Unstick is a wrapper around gtk_window_unstick().
func (v *Window) Unstick() {
	C.gtk_window_unstick(v.native())
}

// Maximize is a wrapper around gtk_window_maximize().
func (v *Window) Maximize() {
	C.gtk_window_maximize(v.native())
}

// Unmaximize is a wrapper around gtk_window_unmaximize().
func (v *Window) Unmaximize() {
	C.gtk_window_unmaximize(v.native())
}

// Fullscreen is a wrapper around gtk_window_fullscreen().
func (v *Window) Fullscreen() {
	C.gtk_window_fullscreen(v.native())
}

// Unfullscreen is a wrapper around gtk_window_unfullscreen().
func (v *Window) Unfullscreen() {
	C.gtk_window_unfullscreen(v.native())
}

// SetKeepAbove is a wrapper around gtk_window_set_keep_above().
func (v *Window) SetKeepAbove(setting bool) {
	C.gtk_window_set_keep_above(v.native(), gbool(setting))
}

// SetKeepBelow is a wrapper around gtk_window_set_keep_below().
func (v *Window) SetKeepBelow(setting bool) {
	C.gtk_window_set_keep_below(v.native(), gbool(setting))
}

// TODO gtk_window_begin_resize_drag().

// TODO gtk_window_begin_move_drag().

// SetDecorated is a wrapper around gtk_window_set_decorated().
func (v *Window) SetDecorated(setting bool) {
	C.gtk_window_set_decorated(v.native(), gbool(setting))
}

// SetDeletable is a wrapper around gtk_window_set_deletable().
func (v *Window) SetDeletable(setting bool) {
	C.gtk_window_set_deletable(v.native(), gbool(setting))
}

// TODO gtk_window_set_mnemonic_modifier().

// TODO gtk_window_set_type_hint().

// SetSkipTaskbarHint is a wrapper around gtk_window_set_skip_taskbar_hint().
func (v *Window) SetSkipTaskbarHint(setting bool) {
	C.gtk_window_set_skip_taskbar_hint(v.native(), gbool(setting))
}

// SetSkipPagerHint is a wrapper around gtk_window_set_skip_pager_hint().
func (v *Window) SetSkipPagerHint(setting bool) {
	C.gtk_window_set_skip_pager_hint(v.native(), gbool(setting))
}

// SetUrgencyHint is a wrapper around gtk_window_set_urgency_hint().
func (v *Window) SetUrgencyHint(setting bool) {
	C.gtk_window_set_urgency_hint(v.native(), gbool(setting))
}

// SetAcceptFocus is a wrapper around gtk_window_set_accept_focus().
func (v *Window) SetAcceptFocus(setting bool) {
	C.gtk_window_set_accept_focus(v.native(), gbool(setting))
}

// SetFocusOnMap is a wrapper around gtk_window_set_focus_on_map().
func (v *Window) SetFocusOnMap(setting bool) {
	C.gtk_window_set_focus_on_map(v.native(), gbool(setting))
}

// TODO gtk_window_set_startup_id().

// TODO gtk_window_set_role().

// SetWMClass is a wrapper around gtk_window_set_wmclass().
func (v *Window) SetWMClass(name, class string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cClass := C.CString(class)
	defer C.free(unsafe.Pointer(cClass))
	C.gtk_window_set_wmclass(v.native(), (*C.gchar)(cName), (*C.gchar)(cClass))
}

// GetDecorated is a wrapper around gtk_window_get_decorated().
func (v *Window) GetDecorated() bool {
	c := C.gtk_window_get_decorated(v.native())
	return gobool(c)
}

// GetDeletable is a wrapper around gtk_window_get_deletable().
func (v *Window) GetDeletable() bool {
	c := C.gtk_window_get_deletable(v.native())
	return gobool(c)
}

// TODO get_default_icon_list().

// TODO get_default_icon_name().

// GetDefaultSize is a wrapper around gtk_window_get_default_size().
func (v *Window) GetDefaultSize() (width, height int) {
	var w, h C.gint
	C.gtk_window_get_default_size(v.native(), &w, &h)
	return int(w), int(h)
}

// GetDestroyWithParent is a wrapper around
// gtk_window_get_destroy_with_parent().
func (v *Window) GetDestroyWithParent() bool {
	c := C.gtk_window_get_destroy_with_parent(v.native())
	return gobool(c)
}

// GetHideTitlebarWhenMaximized is a wrapper around
// gtk_window_get_hide_titlebar_when_maximized().
func (v *Window) GetHideTitlebarWhenMaximized() bool {
	c := C.gtk_window_get_hide_titlebar_when_maximized(v.native())
	return gobool(c)
}

// TODO gtk_window_get_icon().

// TODO gtk_window_get_icon_list().

// TODO gtk_window_get_icon_name().

// TODO gtk_window_get_mnemonic_modifier().

// GetModal is a wrapper around gtk_window_get_modal().
func (v *Window) GetModal() bool {
	c := C.gtk_window_get_modal(v.native())
	return gobool(c)
}

// GetPosition is a wrapper around gtk_window_get_position().
func (v *Window) GetPosition() (root_x, root_y int) {
	var x, y C.gint
	C.gtk_window_get_position(v.native(), &x, &y)
	return int(x), int(y)
}

// TODO gtk_window_get_role().

// GetSize is a wrapper around gtk_window_get_size().
func (v *Window) GetSize() (width, height int) {
	var w, h C.gint
	C.gtk_window_get_size(v.native(), &w, &h)
	return int(w), int(h)
}

// TODO gtk_window_get_title().

// TODO gtk_window_get_transient_for().

// TODO gtk_window_get_attached_to().

// TODO gtk_window_get_type_hint().

// GetSkipTaskbarHint is a wrapper around gtk_window_get_skip_taskbar_hint().
func (v *Window) GetSkipTaskbarHint() bool {
	c := C.gtk_window_get_skip_taskbar_hint(v.native())
	return gobool(c)
}

// GetSkipPagerHint is a wrapper around gtk_window_get_skip_pager_hint().
func (v *Window) GetSkipPagerHint() bool {
	c := C.gtk_window_get_skip_taskbar_hint(v.native())
	return gobool(c)
}

// GetUrgencyHint is a wrapper around gtk_window_get_urgency_hint().
func (v *Window) GetUrgencyHint() bool {
	c := C.gtk_window_get_urgency_hint(v.native())
	return gobool(c)
}

// GetAcceptFocus is a wrapper around gtk_window_get_accept_focus().
func (v *Window) GetAcceptFocus() bool {
	c := C.gtk_window_get_accept_focus(v.native())
	return gobool(c)
}

// GetFocusOnMap is a wrapper around gtk_window_get_focus_on_map().
func (v *Window) GetFocusOnMap() bool {
	c := C.gtk_window_get_focus_on_map(v.native())
	return gobool(c)
}

// TODO gtk_window_get_group().

// HasGroup is a wrapper around gtk_window_has_group().
func (v *Window) HasGroup() bool {
	c := C.gtk_window_has_group(v.native())
	return gobool(c)
}

// TODO gtk_window_get_window_type().

// Move is a wrapper around gtk_window_move().
func (v *Window) Move(x, y int) {
	C.gtk_window_move(v.native(), C.gint(x), C.gint(y))
}

// TODO gtk_window_parse_geometry().

// Resize is a wrapper around gtk_window_resize().
func (v *Window) Resize(width, height int) {
	C.gtk_window_resize(v.native(), C.gint(width), C.gint(height))
}

// ResizeToGeometry is a wrapper around gtk_window_resize_to_geometry().
func (v *Window) ResizeToGeometry(width, height int) {
	C.gtk_window_resize_to_geometry(v.native(), C.gint(width), C.gint(height))
}

// TODO gtk_window_set_default_icon_list().

// TODO gtk_window_set_default_icon().

// TODO gtk_window_set_default_icon_from_file().

// TODO gtk_window_set_default_icon_name().

// TODO gtk_window_set_icon().

// TODO gtk_window_set_icon_list().

// SetIconFromFile is a wrapper around gtk_window_set_icon_from_file().
func (v *Window) SetIconFromFile(file string) error {
	cstr := C.CString(file)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	res := C.gtk_window_set_icon_from_file(v.native(), (*C.gchar)(cstr), &err)
	if res == 0 {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

// SetIconName is a wrapper around gtk_window_set_icon_name().
func (v *Window) SetIconName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_window_set_icon_name(v.native(), (*C.gchar)(cstr))
}

// SetAutoStartupNotification is a wrapper around
// gtk_window_set_auto_startup_notification().
// This doesn't seem write.  Might need to rethink?
/*
func (v *Window) SetAutoStartupNotification(setting bool) {
	C.gtk_window_set_auto_startup_notification(gbool(setting))
}
*/

// GetMnemonicsVisible is a wrapper around
// gtk_window_get_mnemonics_visible().
func (v *Window) GetMnemonicsVisible() bool {
	c := C.gtk_window_get_mnemonics_visible(v.native())
	return gobool(c)
}

// SetMnemonicsVisible is a wrapper around
// gtk_window_get_mnemonics_visible().
func (v *Window) SetMnemonicsVisible(setting bool) {
	C.gtk_window_set_mnemonics_visible(v.native(), gbool(setting))
}

// GetFocusVisible is a wrapper around gtk_window_get_focus_visible().
func (v *Window) GetFocusVisible() bool {
	c := C.gtk_window_get_focus_visible(v.native())
	return gobool(c)
}

// SetFocusVisible is a wrapper around gtk_window_set_focus_visible().
func (v *Window) SetFocusVisible(setting bool) {
	C.gtk_window_set_focus_visible(v.native(), gbool(setting))
}

// TODO gtk_window_get_resize_grip_area().

// TODO gtk_window_set_application().

// TODO gtk_window_get_application().

type WrapFn interface{}

var WrapMap = map[string]WrapFn{
	"GtkAboutDialog":         wrapAboutDialog,
	"GtkAccelGroup":          wrapAccelGroup,
	"GtkAccelMao":            wrapAccelMap,
	"GtkAdjustment":          wrapAdjustment,
	"GtkBin":                 wrapBin,
	"GtkBox":                 wrapBox,
	"GtkButton":              wrapButton,
	"GtkCalendar":            wrapCalendar,
	"GtkCellLayout":          wrapCellLayout,
	"GtkCellRenderer":        wrapCellRenderer,
	"GtkCellRendererSpinner": wrapCellRendererSpinner,
	"GtkCellRendererPixbuf":  wrapCellRendererPixbuf,
	"GtkCellRendererText":    wrapCellRendererText,
	"GtkCellRendererToggle":  wrapCellRendererToggle,
	"GtkCheckButton":         wrapCheckButton,
	"GtkCheckMenuItem":       wrapCheckMenuItem,
	"GtkClipboard":           wrapClipboard,
	"GtkComboBox":            wrapComboBox,
	"GtkComboBoxText":        wrapComboBoxText,
	"GtkContainer":           wrapContainer,
	"GtkDialog":              wrapDialog,
	"GtkDrawingArea":         wrapDrawingArea,
	"GtkEditable":            wrapEditable,
	"GtkEntry":               wrapEntry,
	"GtkEntryBuffer":         wrapEntryBuffer,
	"GtkEntryCompletion":     wrapEntryCompletion,
	"GtkEventBox":            wrapEventBox,
	"GtkExpander":            wrapExpander,
	"GtkFrame":               wrapFrame,
	"GtkFileChooser":         wrapFileChooser,
	"GtkFileChooserButton":   wrapFileChooserButton,
	"GtkFileChooserDialog":   wrapFileChooserDialog,
	"GtkFileChooserWidget":   wrapFileChooserWidget,
	"GtkFontButton":          wrapFontButton,
	"GtkGrid":                wrapGrid,
	"GtkIconView":            wrapIconView,
	"GtkImage":               wrapImage,
	"GtkLabel":               wrapLabel,
	"GtkLayout":              wrapLayout,
	"GtkLinkButton":          wrapLinkButton,
	"GtkListStore":           wrapListStore,
	"GtkMenu":                wrapMenu,
	"GtkMenuBar":             wrapMenuBar,
	"GtkMenuButton":          wrapMenuButton,
	"GtkMenuItem":            wrapMenuItem,
	"GtkMenuShell":           wrapMenuShell,
	"GtkMessageDialog":       wrapMessageDialog,
	"GtkNotebook":            wrapNotebook,
	"GtkOffscreenWindow":     wrapOffscreenWindow,
	"GtkOrientable":          wrapOrientable,
	"GtkPaned":               wrapPaned,
	"GtkProgressBar":         wrapProgressBar,
	"GtkRadioButton":         wrapRadioButton,
	"GtkRadioMenuItem":       wrapRadioMenuItem,
	"GtkRange":               wrapRange,
	"GtkRecentChooser":       wrapRecentChooser,
	"GtkRecentChooserMenu":   wrapRecentChooserMenu,
	"GtkRecentFilter":        wrapRecentFilter,
	"GtkRecentManager":       wrapRecentManager,
	"GtkScaleButton":         wrapScaleButton,
	"GtkScale":               wrapScale,
	"GtkScrollable":          wrapScrollable,
	"GtkScrollbar":           wrapScrollbar,
	"GtkScrolledWindow":      wrapScrolledWindow,
	"GtkSearchEntry":         wrapSearchEntry,
	"GtkSeparator":           wrapSeparator,
	"GtkSeparatorMenuItem":   wrapSeparatorMenuItem,
	"GtkSeparatorToolItem":   wrapSeparatorToolItem,
	"GtkSpinButton":          wrapSpinButton,
	"GtkSpinner":             wrapSpinner,
	"GtkStatusbar":           wrapStatusbar,
	"GtkSwitch":              wrapSwitch,
	"GtkTextView":            wrapTextView,
	"GtkTextBuffer":          wrapTextBuffer,
	"GtkTextTag":             wrapTextTag,
	"GtkTextTagTable":        wrapTextTagTable,
	"GtkToggleButton":        wrapToggleButton,
	"GtkToolbar":             wrapToolbar,
	"GtkToolButton":          wrapToolButton,
	"GtkToolItem":            wrapToolItem,
	"GtkTreeModel":           wrapTreeModel,
	"GtkTreeSelection":       wrapTreeSelection,
	"GtkTreeStore":           wrapTreeStore,
	"GtkTreeView":            wrapTreeView,
	"GtkTreeViewColumn":      wrapTreeViewColumn,
	"GtkViewport":            wrapViewport,
	"GtkVolumeButton":        wrapVolumeButton,
	"GtkWidget":              wrapWidget,
	"GtkWindow":              wrapWindow,
}

// cast takes a native GObject and casts it to the appropriate Go struct.
//TODO change all wrapFns to return an IObject
func cast(c *C.GObject) (glib.IObject, error) {
	var (
		className = C.GoString((*C.char)(C.object_get_class_name(c)))
		obj       = &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	)

	runtime.SetFinalizer(obj, (*glib.Object).Unref)

	fn, ok := WrapMap[className]
	if !ok {
		return nil, errors.New("unrecognized class name '" + className + "'")
	}

	rf := reflect.ValueOf(fn)
	if rf.Type().Kind() != reflect.Func {
		return nil, errors.New("wraper is not a function")
	}

	v := reflect.ValueOf(obj)
	rv := rf.Call([]reflect.Value{v})

	if len(rv) != 1 {
		return nil, errors.New("wrapper did not return")
	}

	if k := rv[0].Kind(); k != reflect.Ptr {
		return nil, fmt.Errorf("wrong return type %s", k)
	}

	ret, ok := rv[0].Interface().(glib.IObject)
	if !ok {
		return nil, errors.New("did not return an IObject")
	}

	return ret, nil
}
