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
package impl

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
	cairo_impl "github.com/gotk3/gotk3/cairo/impl"
	"github.com/gotk3/gotk3/gdk"
	gdk_impl "github.com/gotk3/gotk3/gdk/impl"
	"github.com/gotk3/gotk3/glib"
	glib_impl "github.com/gotk3/gotk3/glib/impl"
	"github.com/gotk3/gotk3/gtk"
)

func init() {
	tm := []glib_impl.TypeMarshaler{
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
		{glib.Type(C.gtk_accel_group_get_type()), marshalAccelGroup},
		{glib.Type(C.gtk_accel_map_get_type()), marshalAccelMap},
		{glib.Type(C.gtk_adjustment_get_type()), marshalAdjustment},
		{glib.Type(C.gtk_application_window_get_type()), marshalApplicationWindow},
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
	glib_impl.RegisterGValueMarshalers(tm)

	gtk.ALIGN_FILL = C.GTK_ALIGN_FILL
	gtk.ALIGN_START = C.GTK_ALIGN_START
	gtk.ALIGN_END = C.GTK_ALIGN_END
	gtk.ALIGN_CENTER = C.GTK_ALIGN_CENTER

	gtk.ARROWS_BOTH = C.GTK_ARROWS_BOTH
	gtk.ARROWS_START = C.GTK_ARROWS_START
	gtk.ARROWS_END = C.GTK_ARROWS_END

	gtk.ARROW_UP = C.GTK_ARROW_UP
	gtk.ARROW_DOWN = C.GTK_ARROW_DOWN
	gtk.ARROW_LEFT = C.GTK_ARROW_LEFT
	gtk.ARROW_RIGHT = C.GTK_ARROW_RIGHT
	gtk.ARROW_NONE = C.GTK_ARROW_NONE

	gtk.ASSISTANT_PAGE_CONTENT = C.GTK_ASSISTANT_PAGE_CONTENT
	gtk.ASSISTANT_PAGE_INTRO = C.GTK_ASSISTANT_PAGE_INTRO
	gtk.ASSISTANT_PAGE_CONFIRM = C.GTK_ASSISTANT_PAGE_CONFIRM
	gtk.ASSISTANT_PAGE_SUMMARY = C.GTK_ASSISTANT_PAGE_SUMMARY
	gtk.ASSISTANT_PAGE_PROGRESS = C.GTK_ASSISTANT_PAGE_PROGRESS
	gtk.ASSISTANT_PAGE_CUSTOM = C.GTK_ASSISTANT_PAGE_CUSTOM

	gtk.BUTTONS_NONE = C.GTK_BUTTONS_NONE
	gtk.BUTTONS_OK = C.GTK_BUTTONS_OK
	gtk.BUTTONS_CLOSE = C.GTK_BUTTONS_CLOSE
	gtk.BUTTONS_CANCEL = C.GTK_BUTTONS_CANCEL
	gtk.BUTTONS_YES_NO = C.GTK_BUTTONS_YES_NO
	gtk.BUTTONS_OK_CANCEL = C.GTK_BUTTONS_OK_CANCEL

	gtk.CALENDAR_SHOW_HEADING = C.GTK_CALENDAR_SHOW_HEADING
	gtk.CALENDAR_SHOW_DAY_NAMES = C.GTK_CALENDAR_SHOW_DAY_NAMES
	gtk.CALENDAR_NO_MONTH_CHANGE = C.GTK_CALENDAR_NO_MONTH_CHANGE
	gtk.CALENDAR_SHOW_WEEK_NUMBERS = C.GTK_CALENDAR_SHOW_WEEK_NUMBERS
	gtk.CALENDAR_SHOW_DETAILS = C.GTK_CALENDAR_SHOW_DETAILS

	gtk.DEST_DEFAULT_MOTION = C.GTK_DEST_DEFAULT_MOTION
	gtk.DEST_DEFAULT_HIGHLIGHT = C.GTK_DEST_DEFAULT_HIGHLIGHT
	gtk.DEST_DEFAULT_DROP = C.GTK_DEST_DEFAULT_DROP
	gtk.DEST_DEFAULT_ALL = C.GTK_DEST_DEFAULT_ALL

	gtk.DIALOG_MODAL = C.GTK_DIALOG_MODAL
	gtk.DIALOG_DESTROY_WITH_PARENT = C.GTK_DIALOG_DESTROY_WITH_PARENT

	gtk.ENTRY_ICON_PRIMARY = C.GTK_ENTRY_ICON_PRIMARY
	gtk.ENTRY_ICON_SECONDARY = C.GTK_ENTRY_ICON_SECONDARY

	gtk.FILE_CHOOSER_ACTION_OPEN = C.GTK_FILE_CHOOSER_ACTION_OPEN
	gtk.FILE_CHOOSER_ACTION_SAVE = C.GTK_FILE_CHOOSER_ACTION_SAVE
	gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER = C.GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER
	gtk.FILE_CHOOSER_ACTION_CREATE_FOLDER = C.GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER

	gtk.ICON_LOOKUP_NO_SVG = C.GTK_ICON_LOOKUP_NO_SVG
	gtk.ICON_LOOKUP_FORCE_SVG = C.GTK_ICON_LOOKUP_FORCE_SVG
	gtk.ICON_LOOKUP_USE_BUILTIN = C.GTK_ICON_LOOKUP_USE_BUILTIN
	gtk.ICON_LOOKUP_GENERIC_FALLBACK = C.GTK_ICON_LOOKUP_GENERIC_FALLBACK
	gtk.ICON_LOOKUP_FORCE_SIZE = C.GTK_ICON_LOOKUP_FORCE_SIZE

	gtk.ICON_SIZE_INVALID = C.GTK_ICON_SIZE_INVALID
	gtk.ICON_SIZE_MENU = C.GTK_ICON_SIZE_MENU
	gtk.ICON_SIZE_SMALL_TOOLBAR = C.GTK_ICON_SIZE_SMALL_TOOLBAR
	gtk.ICON_SIZE_LARGE_TOOLBAR = C.GTK_ICON_SIZE_LARGE_TOOLBAR
	gtk.ICON_SIZE_BUTTON = C.GTK_ICON_SIZE_BUTTON
	gtk.ICON_SIZE_DND = C.GTK_ICON_SIZE_DND
	gtk.ICON_SIZE_DIALOG = C.GTK_ICON_SIZE_DIALOG

	gtk.IMAGE_EMPTY = C.GTK_IMAGE_EMPTY
	gtk.IMAGE_PIXBUF = C.GTK_IMAGE_PIXBUF
	gtk.IMAGE_STOCK = C.GTK_IMAGE_STOCK
	gtk.IMAGE_ICON_SET = C.GTK_IMAGE_ICON_SET
	gtk.IMAGE_ANIMATION = C.GTK_IMAGE_ANIMATION
	gtk.IMAGE_ICON_NAME = C.GTK_IMAGE_ICON_NAME
	gtk.IMAGE_GICON = C.GTK_IMAGE_GICON

	gtk.INPUT_HINT_NONE = C.GTK_INPUT_HINT_NONE
	gtk.INPUT_HINT_SPELLCHECK = C.GTK_INPUT_HINT_SPELLCHECK
	gtk.INPUT_HINT_NO_SPELLCHECK = C.GTK_INPUT_HINT_NO_SPELLCHECK
	gtk.INPUT_HINT_WORD_COMPLETION = C.GTK_INPUT_HINT_WORD_COMPLETION
	gtk.INPUT_HINT_LOWERCASE = C.GTK_INPUT_HINT_LOWERCASE
	gtk.INPUT_HINT_UPPERCASE_CHARS = C.GTK_INPUT_HINT_UPPERCASE_CHARS
	gtk.INPUT_HINT_UPPERCASE_WORDS = C.GTK_INPUT_HINT_UPPERCASE_WORDS
	gtk.INPUT_HINT_UPPERCASE_SENTENCES = C.GTK_INPUT_HINT_UPPERCASE_SENTENCES
	gtk.INPUT_HINT_INHIBIT_OSK = C.GTK_INPUT_HINT_INHIBIT_OSK

	gtk.INPUT_PURPOSE_FREE_FORM = C.GTK_INPUT_PURPOSE_FREE_FORM
	gtk.INPUT_PURPOSE_ALPHA = C.GTK_INPUT_PURPOSE_ALPHA
	gtk.INPUT_PURPOSE_DIGITS = C.GTK_INPUT_PURPOSE_DIGITS
	gtk.INPUT_PURPOSE_NUMBER = C.GTK_INPUT_PURPOSE_NUMBER
	gtk.INPUT_PURPOSE_PHONE = C.GTK_INPUT_PURPOSE_PHONE
	gtk.INPUT_PURPOSE_URL = C.GTK_INPUT_PURPOSE_URL
	gtk.INPUT_PURPOSE_EMAIL = C.GTK_INPUT_PURPOSE_EMAIL
	gtk.INPUT_PURPOSE_NAME = C.GTK_INPUT_PURPOSE_NAME
	gtk.INPUT_PURPOSE_PASSWORD = C.GTK_INPUT_PURPOSE_PASSWORD
	gtk.INPUT_PURPOSE_PIN = C.GTK_INPUT_PURPOSE_PIN

	gtk.JUSTIFY_LEFT = C.GTK_JUSTIFY_LEFT
	gtk.JUSTIFY_RIGHT = C.GTK_JUSTIFY_RIGHT
	gtk.JUSTIFY_CENTER = C.GTK_JUSTIFY_CENTER
	gtk.JUSTIFY_FILL = C.GTK_JUSTIFY_FILL

	gtk.LICENSE_UNKNOWN = C.GTK_LICENSE_UNKNOWN
	gtk.LICENSE_CUSTOM = C.GTK_LICENSE_CUSTOM
	gtk.LICENSE_GPL_2_0 = C.GTK_LICENSE_GPL_2_0
	gtk.LICENSE_GPL_3_0 = C.GTK_LICENSE_GPL_3_0
	gtk.LICENSE_LGPL_2_1 = C.GTK_LICENSE_LGPL_2_1
	gtk.LICENSE_LGPL_3_0 = C.GTK_LICENSE_LGPL_3_0
	gtk.LICENSE_BSD = C.GTK_LICENSE_BSD
	gtk.LICENSE_MIT_X11 = C.GTK_LICENSE_MIT_X11
	gtk.LICENSE_GTK_ARTISTIC = C.GTK_LICENSE_ARTISTIC

	gtk.MESSAGE_INFO = C.GTK_MESSAGE_INFO
	gtk.MESSAGE_WARNING = C.GTK_MESSAGE_WARNING
	gtk.MESSAGE_QUESTION = C.GTK_MESSAGE_QUESTION
	gtk.MESSAGE_ERROR = C.GTK_MESSAGE_ERROR
	gtk.MESSAGE_OTHER = C.GTK_MESSAGE_OTHER

	gtk.ORIENTATION_HORIZONTAL = C.GTK_ORIENTATION_HORIZONTAL
	gtk.ORIENTATION_VERTICAL = C.GTK_ORIENTATION_VERTICAL

	gtk.PACK_START = C.GTK_PACK_START
	gtk.PACK_END = C.GTK_PACK_END

	gtk.PATH_WIDGET = C.GTK_PATH_WIDGET
	gtk.PATH_WIDGET_CLASS = C.GTK_PATH_WIDGET_CLASS
	gtk.PATH_CLASS = C.GTK_PATH_CLASS

	gtk.POLICY_ALWAYS = C.GTK_POLICY_ALWAYS
	gtk.POLICY_AUTOMATIC = C.GTK_POLICY_AUTOMATIC
	gtk.POLICY_NEVER = C.GTK_POLICY_NEVER

	gtk.POS_LEFT = C.GTK_POS_LEFT
	gtk.POS_RIGHT = C.GTK_POS_RIGHT
	gtk.POS_TOP = C.GTK_POS_TOP
	gtk.POS_BOTTOM = C.GTK_POS_BOTTOM

	gtk.RELIEF_NORMAL = C.GTK_RELIEF_NORMAL
	gtk.RELIEF_HALF = C.GTK_RELIEF_HALF
	gtk.RELIEF_NONE = C.GTK_RELIEF_NONE

	gtk.RESPONSE_NONE = C.GTK_RESPONSE_NONE
	gtk.RESPONSE_REJECT = C.GTK_RESPONSE_REJECT
	gtk.RESPONSE_ACCEPT = C.GTK_RESPONSE_ACCEPT
	gtk.RESPONSE_DELETE_EVENT = C.GTK_RESPONSE_DELETE_EVENT
	gtk.RESPONSE_OK = C.GTK_RESPONSE_OK
	gtk.RESPONSE_CANCEL = C.GTK_RESPONSE_CANCEL
	gtk.RESPONSE_CLOSE = C.GTK_RESPONSE_CLOSE
	gtk.RESPONSE_YES = C.GTK_RESPONSE_YES
	gtk.RESPONSE_NO = C.GTK_RESPONSE_NO
	gtk.RESPONSE_APPLY = C.GTK_RESPONSE_APPLY
	gtk.RESPONSE_HELP = C.GTK_RESPONSE_HELP

	gtk.SELECTION_NONE = C.GTK_SELECTION_NONE
	gtk.SELECTION_SINGLE = C.GTK_SELECTION_SINGLE
	gtk.SELECTION_BROWSE = C.GTK_SELECTION_BROWSE
	gtk.SELECTION_MULTIPLE = C.GTK_SELECTION_MULTIPLE

	gtk.SHADOW_NONE = C.GTK_SHADOW_NONE
	gtk.SHADOW_IN = C.GTK_SHADOW_IN
	gtk.SHADOW_OUT = C.GTK_SHADOW_OUT
	gtk.SHADOW_ETCHED_IN = C.GTK_SHADOW_ETCHED_IN
	gtk.SHADOW_ETCHED_OUT = C.GTK_SHADOW_ETCHED_OUT

	gtk.SORT_ASCENDING = C.GTK_SORT_ASCENDING
	gtk.SORT_DESCENDING = C.GTK_SORT_DESCENDING

	gtk.STATE_FLAG_NORMAL = C.GTK_STATE_FLAG_NORMAL
	gtk.STATE_FLAG_ACTIVE = C.GTK_STATE_FLAG_ACTIVE
	gtk.STATE_FLAG_PRELIGHT = C.GTK_STATE_FLAG_PRELIGHT
	gtk.STATE_FLAG_SELECTED = C.GTK_STATE_FLAG_SELECTED
	gtk.STATE_FLAG_INSENSITIVE = C.GTK_STATE_FLAG_INSENSITIVE
	gtk.STATE_FLAG_INCONSISTENT = C.GTK_STATE_FLAG_INCONSISTENT
	gtk.STATE_FLAG_FOCUSED = C.GTK_STATE_FLAG_FOCUSED
	gtk.STATE_FLAG_BACKDROP = C.GTK_STATE_FLAG_BACKDROP

	gtk.TARGET_SAME_APP = C.GTK_TARGET_SAME_APP
	gtk.TARGET_SAME_WIDGET = C.GTK_TARGET_SAME_WIDGET
	gtk.TARGET_OTHER_APP = C.GTK_TARGET_OTHER_APP
	gtk.TARGET_OTHER_WIDGET = C.GTK_TARGET_OTHER_WIDGET

	gtk.TOOLBAR_ICONS = C.GTK_TOOLBAR_ICONS
	gtk.TOOLBAR_TEXT = C.GTK_TOOLBAR_TEXT
	gtk.TOOLBAR_BOTH = C.GTK_TOOLBAR_BOTH
	gtk.TOOLBAR_BOTH_HORIZ = C.GTK_TOOLBAR_BOTH_HORIZ

	gtk.TREE_MODEL_ITERS_PERSIST = C.GTK_TREE_MODEL_ITERS_PERSIST
	gtk.TREE_MODEL_LIST_ONLY = C.GTK_TREE_MODEL_LIST_ONLY

	gtk.WIN_POS_NONE = C.GTK_WIN_POS_NONE
	gtk.WIN_POS_CENTER = C.GTK_WIN_POS_CENTER
	gtk.WIN_POS_MOUSE = C.GTK_WIN_POS_MOUSE
	gtk.WIN_POS_CENTER_ALWAYS = C.GTK_WIN_POS_CENTER_ALWAYS
	gtk.WIN_POS_CENTER_ON_PARENT = C.GTK_WIN_POS_CENTER_ON_PARENT

	gtk.WINDOW_TOPLEVEL = C.GTK_WINDOW_TOPLEVEL
	gtk.WINDOW_POPUP = C.GTK_WINDOW_POPUP

	gtk.WRAP_NONE = C.GTK_WRAP_NONE
	gtk.WRAP_CHAR = C.GTK_WRAP_CHAR
	gtk.WRAP_WORD = C.GTK_WRAP_WORD
	gtk.WRAP_WORD_CHAR = C.GTK_WRAP_WORD_CHAR
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
func wrapObject(ptr unsafe.Pointer) *glib_impl.Object {
	obj := &glib_impl.Object{glib_impl.ToGObject(ptr)}

	if obj.IsFloating() {
		obj.RefSink()
	} else {
		obj.Ref()
	}

	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
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

func marshalAlign(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.Align(c), nil
}

func marshalArrowPlacement(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.ArrowPlacement(c), nil
}

func marshalArrowType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.ArrowType(c), nil
}

func marshalAssistantPageType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.AssistantPageType(c), nil
}

func marshalButtonsType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.ButtonsType(c), nil
}

func marshalCalendarDisplayOptions(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.CalendarDisplayOptions(c), nil
}

func marshalDestDefaults(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.DestDefaults(c), nil
}

func marshalDialogFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.DialogFlags(c), nil
}

func marshalEntryIconPosition(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.EntryIconPosition(c), nil
}

func marshalFileChooserAction(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.FileChooserAction(c), nil
}

func marshalIconLookupFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.IconLookupFlags(c), nil
}

func marshalIconSize(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.IconSize(c), nil
}

func marshalImageType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.ImageType(c), nil
}

func marshalInputHints(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.InputHints(c), nil
}

func marshalInputPurpose(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.InputPurpose(c), nil
}

func marshalJustification(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.Justification(c), nil
}

func marshalLicense(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.License(c), nil
}

func marshalMessageType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.MessageType(c), nil
}

func marshalOrientation(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.Orientation(c), nil
}

func marshalPackType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.PackType(c), nil
}

func marshalPathType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.PathType(c), nil
}

func marshalPolicyType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.PolicyType(c), nil
}

func marshalPositionType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.PositionType(c), nil
}

func marshalReliefStyle(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.ReliefStyle(c), nil
}

func marshalResponseType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.ResponseType(c), nil
}

func marshalSelectionMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.SelectionMode(c), nil
}

func marshalShadowType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.ShadowType(c), nil
}

func marshalSortType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.SortType(c), nil
}

func marshalStateFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.StateFlags(c), nil
}

func marshalTargetFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.TargetFlags(c), nil
}

func marshalToolbarStyle(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.ToolbarStyle(c), nil
}

func marshalTreeModelFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.TreeModelFlags(c), nil
}

func marshalWindowPosition(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.WindowPosition(c), nil
}

func marshalWindowType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.WindowType(c), nil
}

func marshalWrapMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.WrapMode(c), nil
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
		argv := C.make_strings(argc)
		defer C.destroy_strings(argv)

		for i, arg := range *args {
			cstr := C.CString(arg)
			C.set_string(argv, C.int(i), (*C.gchar)(cstr))
		}

		C.gtk_init((*C.int)(unsafe.Pointer(&argc)),
			(***C.char)(unsafe.Pointer(&argv)))

		unhandled := make([]string, argc)
		for i := 0; i < int(argc); i++ {
			cstr := C.get_string(argv, C.int(i))
			unhandled[i] = C.GoString((*C.char)(cstr))
			C.free(unsafe.Pointer(cstr))
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

// MainIteration is a wrapper around gtk_main_iteration.
func MainIteration() bool {
	return gobool(C.gtk_main_iteration())
}

// MainIterationDo is a wrapper around gtk_main_iteration_do.
func MainIterationDo(blocking bool) bool {
	return gobool(C.gtk_main_iteration_do(gbool(blocking)))
}

// EventsPending is a wrapper around gtk_events_pending.
func EventsPending() bool {
	return gobool(C.gtk_events_pending())
}

// MainQuit() is a wrapper around gtk_main_quit() is used to terminate
// the GTK main loop (started by Main()).
func MainQuit() {
	C.gtk_main_quit()
}

/*
 * GtkAdjustment
 */

// Adjustment is a representation of GTK's GtkAdjustment.
type adjustment struct {
	glib_impl.InitiallyUnowned
}

// native returns a pointer to the underlying GtkAdjustment.
func (v *adjustment) native() *C.GtkAdjustment {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkAdjustment(p)
}

func marshalAdjustment(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAdjustment(obj), nil
}

func wrapAdjustment(obj *glib_impl.Object) *adjustment {
	return &adjustment{glib_impl.InitiallyUnowned{obj}}
}

// AdjustmentNew is a wrapper around gtk_adjustment_new().
func AdjustmentNew(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) (*adjustment, error) {
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
func (v *adjustment) GetValue() float64 {
	c := C.gtk_adjustment_get_value(v.native())
	return float64(c)
}

// SetValue is a wrapper around gtk_adjustment_set_value().
func (v *adjustment) SetValue(value float64) {
	C.gtk_adjustment_set_value(v.native(), C.gdouble(value))
}

// GetLower is a wrapper around gtk_adjustment_get_lower().
func (v *adjustment) GetLower() float64 {
	c := C.gtk_adjustment_get_lower(v.native())
	return float64(c)
}

// GetPageSize is a wrapper around gtk_adjustment_get_page_size().
func (v *adjustment) GetPageSize() float64 {
	return float64(C.gtk_adjustment_get_page_size(v.native()))
}

// SetPageSize is a wrapper around gtk_adjustment_set_page_size().
func (v *adjustment) SetPageSize(value float64) {
	C.gtk_adjustment_set_page_size(v.native(), C.gdouble(value))
}

// Configure is a wrapper around gtk_adjustment_configure().
func (v *adjustment) Configure(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) {
	C.gtk_adjustment_configure(v.native(), C.gdouble(value),
		C.gdouble(lower), C.gdouble(upper), C.gdouble(stepIncrement),
		C.gdouble(pageIncrement), C.gdouble(pageSize))
}

// SetLower is a wrapper around gtk_adjustment_set_lower().
func (v *adjustment) SetLower(value float64) {
	C.gtk_adjustment_set_lower(v.native(), C.gdouble(value))
}

// GetUpper is a wrapper around gtk_adjustment_get_upper().
func (v *adjustment) GetUpper() float64 {
	c := C.gtk_adjustment_get_upper(v.native())
	return float64(c)
}

// SetUpper is a wrapper around gtk_adjustment_set_upper().
func (v *adjustment) SetUpper(value float64) {
	C.gtk_adjustment_set_upper(v.native(), C.gdouble(value))
}

// GetPageIncrement is a wrapper around gtk_adjustment_get_page_increment().
func (v *adjustment) GetPageIncrement() float64 {
	c := C.gtk_adjustment_get_page_increment(v.native())
	return float64(c)
}

// SetPageIncrement is a wrapper around gtk_adjustment_set_page_increment().
func (v *adjustment) SetPageIncrement(value float64) {
	C.gtk_adjustment_set_page_increment(v.native(), C.gdouble(value))
}

// GetStepIncrement is a wrapper around gtk_adjustment_get_step_increment().
func (v *adjustment) GetStepIncrement() float64 {
	c := C.gtk_adjustment_get_step_increment(v.native())
	return float64(c)
}

// SetStepIncrement is a wrapper around gtk_adjustment_set_step_increment().
func (v *adjustment) SetStepIncrement(value float64) {
	C.gtk_adjustment_set_step_increment(v.native(), C.gdouble(value))
}

// GetMinimumIncrement is a wrapper around gtk_adjustment_get_minimum_increment().
func (v *adjustment) GetMinimumIncrement() float64 {
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
type assistant struct {
	window
}

// native returns a pointer to the underlying GtkAssistant.
func (v *assistant) native() *C.GtkAssistant {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkAssistant(p)
}

func marshalAssistant(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAssistant(obj), nil
}

func wrapAssistant(obj *glib_impl.Object) *assistant {
	return &assistant{window{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}}
}

// AssistantNew is a wrapper around gtk_assistant_new().
func AssistantNew() (*assistant, error) {
	c := C.gtk_assistant_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAssistant(obj), nil
}

// GetCurrentPage is a wrapper around gtk_assistant_get_current_page().
func (v *assistant) GetCurrentPage() int {
	c := C.gtk_assistant_get_current_page(v.native())
	return int(c)
}

// SetCurrentPage is a wrapper around gtk_assistant_set_current_page().
func (v *assistant) SetCurrentPage(pageNum int) {
	C.gtk_assistant_set_current_page(v.native(), C.gint(pageNum))
}

// GetNPages is a wrapper around gtk_assistant_get_n_pages().
func (v *assistant) GetNPages() int {
	c := C.gtk_assistant_get_n_pages(v.native())
	return int(c)
}

// GetNthPage is a wrapper around gtk_assistant_get_nth_page().
func (v *assistant) GetNthPage(pageNum int) (gtk.Widget, error) {
	c := C.gtk_assistant_get_nth_page(v.native(), C.gint(pageNum))
	if c == nil {
		return nil, fmt.Errorf("page %d is out of bounds", pageNum)
	}

	obj := wrapObject(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// PrependPage is a wrapper around gtk_assistant_prepend_page().
func (v *assistant) PrependPage(page gtk.Widget) int {
	c := C.gtk_assistant_prepend_page(v.native(), page.(IWidget).toWidget())
	return int(c)
}

// AppendPage is a wrapper around gtk_assistant_append_page().
func (v *assistant) AppendPage(page gtk.Widget) int {
	c := C.gtk_assistant_append_page(v.native(), page.(IWidget).toWidget())
	return int(c)
}

// InsertPage is a wrapper around gtk_assistant_insert_page().
func (v *assistant) InsertPage(page gtk.Widget, position int) int {
	c := C.gtk_assistant_insert_page(v.native(), page.(IWidget).toWidget(),
		C.gint(position))
	return int(c)
}

// RemovePage is a wrapper around gtk_assistant_remove_page().
func (v *assistant) RemovePage(pageNum int) {
	C.gtk_assistant_remove_page(v.native(), C.gint(pageNum))
}

// TODO: gtk_assistant_set_forward_page_func

// SetPageType is a wrapper around gtk_assistant_set_page_type().
func (v *assistant) SetPageType(page gtk.Widget, ptype gtk.AssistantPageType) {
	C.gtk_assistant_set_page_type(v.native(), page.(IWidget).toWidget(),
		C.GtkAssistantPageType(ptype))
}

// GetPageType is a wrapper around gtk_assistant_get_page_type().
func (v *assistant) GetPageType(page gtk.Widget) gtk.AssistantPageType {
	c := C.gtk_assistant_get_page_type(v.native(), page.(IWidget).toWidget())
	return gtk.AssistantPageType(c)
}

// SetPageTitle is a wrapper around gtk_assistant_set_page_title().
func (v *assistant) SetPageTitle(page gtk.Widget, title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_assistant_set_page_title(v.native(), page.(IWidget).toWidget(),
		(*C.gchar)(cstr))
}

// GetPageTitle is a wrapper around gtk_assistant_get_page_title().
func (v *assistant) GetPageTitle(page gtk.Widget) string {
	c := C.gtk_assistant_get_page_title(v.native(), page.(IWidget).toWidget())
	return C.GoString((*C.char)(c))
}

// SetPageComplete is a wrapper around gtk_assistant_set_page_complete().
func (v *assistant) SetPageComplete(page gtk.Widget, complete bool) {
	C.gtk_assistant_set_page_complete(v.native(), page.(IWidget).toWidget(),
		gbool(complete))
}

// GetPageComplete is a wrapper around gtk_assistant_get_page_complete().
func (v *assistant) GetPageComplete(page gtk.Widget) bool {
	c := C.gtk_assistant_get_page_complete(v.native(), page.(IWidget).toWidget())
	return gobool(c)
}

// AddActionWidget is a wrapper around gtk_assistant_add_action_widget().
func (v *assistant) AddActionWidget(child gtk.Widget) {
	C.gtk_assistant_add_action_widget(v.native(), child.(IWidget).(IWidget).toWidget())
}

// RemoveActionWidget is a wrapper around gtk_assistant_remove_action_widget().
func (v *assistant) RemoveActionWidget(child gtk.Widget) {
	C.gtk_assistant_remove_action_widget(v.native(), child.(IWidget).toWidget())
}

// UpdateButtonsState is a wrapper around gtk_assistant_update_buttons_state().
func (v *assistant) UpdateButtonsState() {
	C.gtk_assistant_update_buttons_state(v.native())
}

// Commit is a wrapper around gtk_assistant_commit().
func (v *assistant) Commit() {
	C.gtk_assistant_commit(v.native())
}

// NextPage is a wrapper around gtk_assistant_next_page().
func (v *assistant) NextPage() {
	C.gtk_assistant_next_page(v.native())
}

// PreviousPage is a wrapper around gtk_assistant_previous_page().
func (v *assistant) PreviousPage() {
	C.gtk_assistant_previous_page(v.native())
}

/*
 * GtkBin
 */

// Bin is a representation of GTK's GtkBin.
type bin struct {
	container
}

// native returns a pointer to the underlying GtkBin.
func (v *bin) native() *C.GtkBin {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkBin(p)
}

func marshalBin(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapBin(obj), nil
}

func wrapBin(obj *glib_impl.Object) *bin {
	return &bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}
}

// GetChild is a wrapper around gtk_bin_get_child().
func (v *bin) GetChild() (gtk.Widget, error) {
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
type builder struct {
	*glib_impl.Object
}

// native() returns a pointer to the underlying GtkBuilder.
func (b *builder) native() *C.GtkBuilder {
	if b == nil || b.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(b.GObject)
	return C.toGtkBuilder(p)
}

func marshalBuilder(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return &builder{obj}, nil
}

// BuilderNew is a wrapper around gtk_builder_new().
func BuilderNew() (*builder, error) {
	c := C.gtk_builder_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return &builder{obj}, nil
}

// AddFromFile is a wrapper around gtk_builder_add_from_file().
func (b *builder) AddFromFile(filename string) error {
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
func (b *builder) AddFromResource(path string) error {
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
func (b *builder) AddFromString(str string) error {
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
//   if w, ok := castToWindow(obj); ok {
//       // do stuff with w here
//   } else {
//       // not a *gtk.Window
//   }
//
func (b *builder) GetObject(name string) (glib.Object, error) {
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
	return obj.(glib.Object), nil
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
func (b *builder) ConnectSignals(signals map[string]interface{}) {
	builderSignals.Lock()
	builderSignals.m[b.native()] = signals
	builderSignals.Unlock()

	C._gtk_builder_connect_signals_full(b.native())
}

/*
 * GtkButton
 */

// Button is a representation of GTK's GtkButton.
type button struct {
	bin
}

// native() returns a pointer to the underlying GtkButton.
func (v *button) native() *C.GtkButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkButton(p)
}

func marshalButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapButton(obj), nil
}

func wrapButton(obj *glib_impl.Object) *button {
	return &button{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}
}

// ButtonNew() is a wrapper around gtk_button_new().
func ButtonNew() (*button, error) {
	c := C.gtk_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapButton(obj), nil
}

// ButtonNewWithLabel() is a wrapper around gtk_button_new_with_label().
func ButtonNewWithLabel(label string) (*button, error) {
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
func ButtonNewWithMnemonic(label string) (*button, error) {
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
func (v *button) Clicked() {
	C.gtk_button_clicked(v.native())
}

// SetRelief() is a wrapper around gtk_button_set_relief().
func (v *button) SetRelief(newStyle gtk.ReliefStyle) {
	C.gtk_button_set_relief(v.native(), C.GtkReliefStyle(newStyle))
}

// GetRelief() is a wrapper around gtk_button_get_relief().
func (v *button) GetRelief() gtk.ReliefStyle {
	c := C.gtk_button_get_relief(v.native())
	return gtk.ReliefStyle(c)
}

// SetLabel() is a wrapper around gtk_button_set_label().
func (v *button) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_button_set_label(v.native(), (*C.gchar)(cstr))
}

// GetLabel() is a wrapper around gtk_button_get_label().
func (v *button) GetLabel() (string, error) {
	c := C.gtk_button_get_label(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// SetUseUnderline() is a wrapper around gtk_button_set_use_underline().
func (v *button) SetUseUnderline(useUnderline bool) {
	C.gtk_button_set_use_underline(v.native(), gbool(useUnderline))
}

// GetUseUnderline() is a wrapper around gtk_button_get_use_underline().
func (v *button) GetUseUnderline() bool {
	c := C.gtk_button_get_use_underline(v.native())
	return gobool(c)
}

// SetFocusOnClick() is a wrapper around gtk_button_set_focus_on_click().
func (v *button) SetFocusOnClick(focusOnClick bool) {
	C.gtk_button_set_focus_on_click(v.native(), gbool(focusOnClick))
}

// GetFocusOnClick() is a wrapper around gtk_button_get_focus_on_click().
func (v *button) GetFocusOnClick() bool {
	c := C.gtk_button_get_focus_on_click(v.native())
	return gobool(c)
}

// SetImage() is a wrapper around gtk_button_set_image().
func (v *button) SetImage(image gtk.Widget) {
	C.gtk_button_set_image(v.native(), image.(IWidget).toWidget())
}

// GetImage() is a wrapper around gtk_button_get_image().
func (v *button) GetImage() (gtk.Widget, error) {
	c := C.gtk_button_get_image(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// SetImagePosition() is a wrapper around gtk_button_set_image_position().
func (v *button) SetImagePosition(position gtk.PositionType) {
	C.gtk_button_set_image_position(v.native(), C.GtkPositionType(position))
}

// GetImagePosition() is a wrapper around gtk_button_get_image_position().
func (v *button) GetImagePosition() gtk.PositionType {
	c := C.gtk_button_get_image_position(v.native())
	return gtk.PositionType(c)
}

// SetAlwaysShowImage() is a wrapper around gtk_button_set_always_show_image().
func (v *button) SetAlwaysShowImage(alwaysShow bool) {
	C.gtk_button_set_always_show_image(v.native(), gbool(alwaysShow))
}

// GetAlwaysShowImage() is a wrapper around gtk_button_get_always_show_image().
func (v *button) GetAlwaysShowImage() bool {
	c := C.gtk_button_get_always_show_image(v.native())
	return gobool(c)
}

// GetEventWindow() is a wrapper around gtk_button_get_event_window().
func (v *button) GetEventWindow() (gdk.Window, error) {
	c := C.gtk_button_get_event_window(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	w := &gdk_impl.Window{wrapObject(unsafe.Pointer(c))}
	return w, nil
}

/*
 * GtkColorButton
 */

// ColorButton is a representation of GTK's GtkColorButton.
type colorButton struct {
	button

	// Interfaces
	colorChooser
}

// Native returns a pointer to the underlying GtkColorButton.
func (v *colorButton) native() *C.GtkColorButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkColorButton(p)
}

func wrapColorButton(obj *glib_impl.Object) *colorButton {
	cc := wrapColorChooser(obj)
	return &colorButton{button{bin{container{widget{
		glib_impl.InitiallyUnowned{obj}}}}}, *cc}
}

// ColorButtonNew is a wrapper around gtk_color_button_new().
func ColorButtonNew() (*colorButton, error) {
	c := C.gtk_color_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapColorButton(wrapObject(unsafe.Pointer(c))), nil
}

// ColorButtonNewWithRGBA is a wrapper around gtk_color_button_new_with_rgba().
func ColorButtonNewWithRGBA(gdkColor *gdk_impl.RGBA) (*colorButton, error) {
	c := C.gtk_color_button_new_with_rgba((*C.GdkRGBA)(unsafe.Pointer(gdkColor.Native())))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapColorButton(wrapObject(unsafe.Pointer(c))), nil
}

/*
 * GtkBox
 */

// Box is a representation of GTK's GtkBox.
type box struct {
	container
}

// native() returns a pointer to the underlying GtkBox.
func (v *box) native() *C.GtkBox {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkBox(p)
}

func marshalBox(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapBox(obj), nil
}

func wrapBox(obj *glib_impl.Object) *box {
	return &box{container{widget{glib_impl.InitiallyUnowned{obj}}}}
}

// BoxNew() is a wrapper around gtk_box_new().
func BoxNew(orientation gtk.Orientation, spacing int) (*box, error) {
	c := C.gtk_box_new(C.GtkOrientation(orientation), C.gint(spacing))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapBox(obj), nil
}

// PackStart() is a wrapper around gtk_box_pack_start().
func (v *box) PackStart(child gtk.Widget, expand, fill bool, padding uint) {
	C.gtk_box_pack_start(v.native(), child.(IWidget).toWidget(), gbool(expand),
		gbool(fill), C.guint(padding))
}

// PackEnd() is a wrapper around gtk_box_pack_end().
func (v *box) PackEnd(child gtk.Widget, expand, fill bool, padding uint) {
	C.gtk_box_pack_end(v.native(), child.(IWidget).toWidget(), gbool(expand),
		gbool(fill), C.guint(padding))
}

// GetHomogeneous() is a wrapper around gtk_box_get_homogeneous().
func (v *box) GetHomogeneous() bool {
	c := C.gtk_box_get_homogeneous(v.native())
	return gobool(c)
}

// SetHomogeneous() is a wrapper around gtk_box_set_homogeneous().
func (v *box) SetHomogeneous(homogeneous bool) {
	C.gtk_box_set_homogeneous(v.native(), gbool(homogeneous))
}

// GetSpacing() is a wrapper around gtk_box_get_spacing().
func (v *box) GetSpacing() int {
	c := C.gtk_box_get_spacing(v.native())
	return int(c)
}

// SetSpacing() is a wrapper around gtk_box_set_spacing()
func (v *box) SetSpacing(spacing int) {
	C.gtk_box_set_spacing(v.native(), C.gint(spacing))
}

// ReorderChild() is a wrapper around gtk_box_reorder_child().
func (v *box) ReorderChild(child gtk.Widget, position int) {
	C.gtk_box_reorder_child(v.native(), child.(IWidget).toWidget(), C.gint(position))
}

// QueryChildPacking() is a wrapper around gtk_box_query_child_packing().
func (v *box) QueryChildPacking(child gtk.Widget) (expand, fill bool, padding uint, packType gtk.PackType) {
	var cexpand, cfill C.gboolean
	var cpadding C.guint
	var cpackType C.GtkPackType

	C.gtk_box_query_child_packing(v.native(), child.(IWidget).toWidget(), &cexpand,
		&cfill, &cpadding, &cpackType)
	return gobool(cexpand), gobool(cfill), uint(cpadding), gtk.PackType(cpackType)
}

// SetChildPacking() is a wrapper around gtk_box_set_child_packing().
func (v *box) SetChildPacking(child gtk.Widget, expand, fill bool, padding uint, packType gtk.PackType) {
	C.gtk_box_set_child_packing(v.native(), child.(IWidget).toWidget(), gbool(expand),
		gbool(fill), C.guint(padding), C.GtkPackType(packType))
}

/*
 * GtkCalendar
 */

// Calendar is a representation of GTK's GtkCalendar.
type calendar struct {
	widget
}

// native() returns a pointer to the underlying GtkCalendar.
func (v *calendar) native() *C.GtkCalendar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCalendar(p)
}

func marshalCalendar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCalendar(obj), nil
}

func wrapCalendar(obj *glib_impl.Object) *calendar {
	return &calendar{widget{glib_impl.InitiallyUnowned{obj}}}
}

// CalendarNew is a wrapper around gtk_calendar_new().
func CalendarNew() (*calendar, error) {
	c := C.gtk_calendar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCalendar(obj), nil
}

// SelectMonth is a wrapper around gtk_calendar_select_month().
func (v *calendar) SelectMonth(month, year uint) {
	C.gtk_calendar_select_month(v.native(), C.guint(month), C.guint(year))
}

// SelectDay is a wrapper around gtk_calendar_select_day().
func (v *calendar) SelectDay(day uint) {
	C.gtk_calendar_select_day(v.native(), C.guint(day))
}

// MarkDay is a wrapper around gtk_calendar_mark_day().
func (v *calendar) MarkDay(day uint) {
	C.gtk_calendar_mark_day(v.native(), C.guint(day))
}

// UnmarkDay is a wrapper around gtk_calendar_unmark_day().
func (v *calendar) UnmarkDay(day uint) {
	C.gtk_calendar_unmark_day(v.native(), C.guint(day))
}

// GetDayIsMarked is a wrapper around gtk_calendar_get_day_is_marked().
func (v *calendar) GetDayIsMarked(day uint) bool {
	c := C.gtk_calendar_get_day_is_marked(v.native(), C.guint(day))
	return gobool(c)
}

// ClearMarks is a wrapper around gtk_calendar_clear_marks().
func (v *calendar) ClearMarks() {
	C.gtk_calendar_clear_marks(v.native())
}

// GetDisplayOptions is a wrapper around gtk_calendar_get_display_options().
func (v *calendar) GetDisplayOptions() gtk.CalendarDisplayOptions {
	c := C.gtk_calendar_get_display_options(v.native())
	return gtk.CalendarDisplayOptions(c)
}

// SetDisplayOptions is a wrapper around gtk_calendar_set_display_options().
func (v *calendar) SetDisplayOptions(flags gtk.CalendarDisplayOptions) {
	C.gtk_calendar_set_display_options(v.native(),
		C.GtkCalendarDisplayOptions(flags))
}

// GetDate is a wrapper around gtk_calendar_get_date().
func (v *calendar) GetDate() (year, month, day uint) {
	var cyear, cmonth, cday C.guint
	C.gtk_calendar_get_date(v.native(), &cyear, &cmonth, &cday)
	return uint(cyear), uint(cmonth), uint(cday)
}

// TODO gtk_calendar_set_detail_func

// GetDetailWidthChars is a wrapper around gtk_calendar_get_detail_width_chars().
func (v *calendar) GetDetailWidthChars() int {
	c := C.gtk_calendar_get_detail_width_chars(v.native())
	return int(c)
}

// SetDetailWidthChars is a wrapper around gtk_calendar_set_detail_width_chars().
func (v *calendar) SetDetailWidthChars(chars int) {
	C.gtk_calendar_set_detail_width_chars(v.native(), C.gint(chars))
}

// GetDetailHeightRows is a wrapper around gtk_calendar_get_detail_height_rows().
func (v *calendar) GetDetailHeightRows() int {
	c := C.gtk_calendar_get_detail_height_rows(v.native())
	return int(c)
}

// SetDetailHeightRows is a wrapper around gtk_calendar_set_detail_height_rows().
func (v *calendar) SetDetailHeightRows(rows int) {
	C.gtk_calendar_set_detail_height_rows(v.native(), C.gint(rows))
}

/*
 * GtkCellLayout
 */

// CellLayout is a representation of GTK's GtkCellLayout GInterface.
type cellLayout struct {
	*glib_impl.Object
}

// ICellLayout is an interface type implemented by all structs
// embedding a CellLayout.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkCellLayout.
type ICellLayout interface {
	toCellLayout() *C.GtkCellLayout
}

// native() returns a pointer to the underlying GObject as a GtkCellLayout.
func (v *cellLayout) native() *C.GtkCellLayout {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCellLayout(p)
}

func marshalCellLayout(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCellLayout(obj), nil
}

func wrapCellLayout(obj *glib_impl.Object) *cellLayout {
	return &cellLayout{obj}
}

func (v *cellLayout) toCellLayout() *C.GtkCellLayout {
	if v == nil {
		return nil
	}
	return v.native()
}

// PackStart() is a wrapper around gtk_cell_layout_pack_start().
func (v *cellLayout) PackStart(cell gtk.CellRenderer, expand bool) {
	C.gtk_cell_layout_pack_start(v.native(), cell.(ICellRenderer).toCellRenderer(),
		gbool(expand))
}

// AddAttribute() is a wrapper around gtk_cell_layout_add_attribute().
func (v *cellLayout) AddAttribute(cell gtk.CellRenderer, attribute string, column int) {
	cstr := C.CString(attribute)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_cell_layout_add_attribute(v.native(), cell.(ICellRenderer).toCellRenderer(),
		(*C.gchar)(cstr), C.gint(column))
}

/*
 * GtkCellRenderer
 */

// CellRenderer is a representation of GTK's GtkCellRenderer.
type cellRenderer struct {
	glib_impl.InitiallyUnowned
}

// ICellRenderer is an interface type implemented by all structs
// embedding a CellRenderer.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkCellRenderer.
type ICellRenderer interface {
	toCellRenderer() *C.GtkCellRenderer
}

// native returns a pointer to the underlying GtkCellRenderer.
func (v *cellRenderer) native() *C.GtkCellRenderer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCellRenderer(p)
}

func (v *cellRenderer) toCellRenderer() *C.GtkCellRenderer {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalCellRenderer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCellRenderer(obj), nil
}

func wrapCellRenderer(obj *glib_impl.Object) *cellRenderer {
	return &cellRenderer{glib_impl.InitiallyUnowned{obj}}
}

/*
 * GtkCellRendererSpinner
 */

// CellRendererSpinner is a representation of GTK's GtkCellRendererSpinner.
type cellRendererSpinner struct {
	cellRenderer
}

// native returns a pointer to the underlying GtkCellRendererSpinner.
func (v *cellRendererSpinner) native() *C.GtkCellRendererSpinner {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCellRendererSpinner(p)
}

func marshalCellRendererSpinner(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCellRendererSpinner(obj), nil
}

func wrapCellRendererSpinner(obj *glib_impl.Object) *cellRendererSpinner {
	return &cellRendererSpinner{cellRenderer{glib_impl.InitiallyUnowned{obj}}}
}

// CellRendererSpinnerNew is a wrapper around gtk_cell_renderer_text_new().
func CellRendererSpinnerNew() (*cellRendererSpinner, error) {
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
type cellRendererPixbuf struct {
	cellRenderer
}

// native returns a pointer to the underlying GtkCellRendererPixbuf.
func (v *cellRendererPixbuf) native() *C.GtkCellRendererPixbuf {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCellRendererPixbuf(p)
}

func marshalCellRendererPixbuf(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCellRendererPixbuf(obj), nil
}

func wrapCellRendererPixbuf(obj *glib_impl.Object) *cellRendererPixbuf {
	return &cellRendererPixbuf{cellRenderer{glib_impl.InitiallyUnowned{obj}}}
}

// CellRendererPixbufNew is a wrapper around gtk_cell_renderer_pixbuf_new().
func CellRendererPixbufNew() (*cellRendererPixbuf, error) {
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
type cellRendererText struct {
	cellRenderer
}

// native returns a pointer to the underlying GtkCellRendererText.
func (v *cellRendererText) native() *C.GtkCellRendererText {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCellRendererText(p)
}

func marshalCellRendererText(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCellRendererText(obj), nil
}

func wrapCellRendererText(obj *glib_impl.Object) *cellRendererText {
	return &cellRendererText{cellRenderer{glib_impl.InitiallyUnowned{obj}}}
}

// CellRendererTextNew is a wrapper around gtk_cell_renderer_text_new().
func CellRendererTextNew() (*cellRendererText, error) {
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
type cellRendererToggle struct {
	cellRenderer
}

// native returns a pointer to the underlying GtkCellRendererToggle.
func (v *cellRendererToggle) native() *C.GtkCellRendererToggle {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCellRendererToggle(p)
}

func (v *cellRendererToggle) toCellRenderer() *C.GtkCellRenderer {
	if v == nil {
		return nil
	}
	return v.cellRenderer.native()
}

func marshalCellRendererToggle(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCellRendererToggle(obj), nil
}

func wrapCellRendererToggle(obj *glib_impl.Object) *cellRendererToggle {
	return &cellRendererToggle{cellRenderer{glib_impl.InitiallyUnowned{obj}}}
}

// CellRendererToggleNew is a wrapper around gtk_cell_renderer_toggle_new().
func CellRendererToggleNew() (*cellRendererToggle, error) {
	c := C.gtk_cell_renderer_toggle_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCellRendererToggle(obj), nil
}

// SetRadio is a wrapper around gtk_cell_renderer_toggle_set_radio().
func (v *cellRendererToggle) SetRadio(set bool) {
	C.gtk_cell_renderer_toggle_set_radio(v.native(), gbool(set))
}

// GetRadio is a wrapper around gtk_cell_renderer_toggle_get_radio().
func (v *cellRendererToggle) GetRadio() bool {
	c := C.gtk_cell_renderer_toggle_get_radio(v.native())
	return gobool(c)
}

// SetActive is a wrapper arround gtk_cell_renderer_toggle_set_active().
func (v *cellRendererToggle) SetActive(active bool) {
	C.gtk_cell_renderer_toggle_set_active(v.native(), gbool(active))
}

// GetActive is a wrapper around gtk_cell_renderer_toggle_get_active().
func (v *cellRendererToggle) GetActive() bool {
	c := C.gtk_cell_renderer_toggle_get_active(v.native())
	return gobool(c)
}

// SetActivatable is a wrapper around gtk_cell_renderer_toggle_set_activatable().
func (v *cellRendererToggle) SetActivatable(activatable bool) {
	C.gtk_cell_renderer_toggle_set_activatable(v.native(),
		gbool(activatable))
}

// GetActivatable is a wrapper around gtk_cell_renderer_toggle_get_activatable().
func (v *cellRendererToggle) GetActivatable() bool {
	c := C.gtk_cell_renderer_toggle_get_activatable(v.native())
	return gobool(c)
}

/*
 * GtkCheckButton
 */

// CheckButton is a wrapper around GTK's GtkCheckButton.
type checkButton struct {
	toggleButton
}

// native returns a pointer to the underlying GtkCheckButton.
func (v *checkButton) native() *C.GtkCheckButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCheckButton(p)
}

func marshalCheckButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCheckButton(obj), nil
}

func wrapCheckButton(obj *glib_impl.Object) *checkButton {
	return &checkButton{toggleButton{button{bin{container{widget{
		glib_impl.InitiallyUnowned{obj}}}}}}}
}

// CheckButtonNew is a wrapper around gtk_check_button_new().
func CheckButtonNew() (*checkButton, error) {
	c := C.gtk_check_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCheckButton(obj), nil
}

// CheckButtonNewWithLabel is a wrapper around
// gtk_check_button_new_with_label().
func CheckButtonNewWithLabel(label string) (*checkButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_check_button_new_with_label((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapCheckButton(wrapObject(unsafe.Pointer(c))), nil
}

// CheckButtonNewWithMnemonic is a wrapper around
// gtk_check_button_new_with_mnemonic().
func CheckButtonNewWithMnemonic(label string) (*checkButton, error) {
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

type checkMenuItem struct {
	menuItem
}

// native returns a pointer to the underlying GtkCheckMenuItem.
func (v *checkMenuItem) native() *C.GtkCheckMenuItem {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCheckMenuItem(p)
}

func marshalCheckMenuItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCheckMenuItem(obj), nil
}

func wrapCheckMenuItem(obj *glib_impl.Object) *checkMenuItem {
	return &checkMenuItem{menuItem{bin{container{widget{
		glib_impl.InitiallyUnowned{obj}}}}}}
}

// CheckMenuItemNew is a wrapper around gtk_check_menu_item_new().
func CheckMenuItemNew() (*checkMenuItem, error) {
	c := C.gtk_check_menu_item_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapCheckMenuItem(obj), nil
}

// CheckMenuItemNewWithLabel is a wrapper around
// gtk_check_menu_item_new_with_label().
func CheckMenuItemNewWithLabel(label string) (*checkMenuItem, error) {
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
func CheckMenuItemNewWithMnemonic(label string) (*checkMenuItem, error) {
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
func (v *checkMenuItem) GetActive() bool {
	c := C.gtk_check_menu_item_get_active(v.native())
	return gobool(c)
}

// SetActive is a wrapper around gtk_check_menu_item_set_active().
func (v *checkMenuItem) SetActive(isActive bool) {
	C.gtk_check_menu_item_set_active(v.native(), gbool(isActive))
}

// Toggled is a wrapper around gtk_check_menu_item_toggled().
func (v *checkMenuItem) Toggled() {
	C.gtk_check_menu_item_toggled(v.native())
}

// GetInconsistent is a wrapper around gtk_check_menu_item_get_inconsistent().
func (v *checkMenuItem) GetInconsistent() bool {
	c := C.gtk_check_menu_item_get_inconsistent(v.native())
	return gobool(c)
}

// SetInconsistent is a wrapper around gtk_check_menu_item_set_inconsistent().
func (v *checkMenuItem) SetInconsistent(setting bool) {
	C.gtk_check_menu_item_set_inconsistent(v.native(), gbool(setting))
}

// SetDrawAsRadio is a wrapper around gtk_check_menu_item_set_draw_as_radio().
func (v *checkMenuItem) SetDrawAsRadio(drawAsRadio bool) {
	C.gtk_check_menu_item_set_draw_as_radio(v.native(), gbool(drawAsRadio))
}

// GetDrawAsRadio is a wrapper around gtk_check_menu_item_get_draw_as_radio().
func (v *checkMenuItem) GetDrawAsRadio() bool {
	c := C.gtk_check_menu_item_get_draw_as_radio(v.native())
	return gobool(c)
}

/*
 * GtkClipboard
 */

// Clipboard is a wrapper around GTK's GtkClipboard.
type clipboard struct {
	*glib_impl.Object
}

// native returns a pointer to the underlying GtkClipboard.
func (v *clipboard) native() *C.GtkClipboard {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkClipboard(p)
}

func marshalClipboard(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapClipboard(obj), nil
}

func wrapClipboard(obj *glib_impl.Object) *clipboard {
	return &clipboard{obj}
}

// Store is a wrapper around gtk_clipboard_store
func (v *clipboard) Store() {
	C.gtk_clipboard_store(v.native())
}

// ClipboardGet() is a wrapper around gtk_clipboard_get().
func ClipboardGet(atom gdk.Atom) (*clipboard, error) {
	c := C.gtk_clipboard_get(C.GdkAtom(unsafe.Pointer(atom)))
	if c == nil {
		return nil, nilPtrErr
	}

	cb := &clipboard{wrapObject(unsafe.Pointer(c))}
	return cb, nil
}

// ClipboardGetForDisplay() is a wrapper around gtk_clipboard_get_for_display().
func ClipboardGetForDisplay(display *gdk_impl.Display, atom gdk.Atom) (*clipboard, error) {
	displayPtr := (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	c := C.gtk_clipboard_get_for_display(displayPtr,
		C.GdkAtom(unsafe.Pointer(atom)))
	if c == nil {
		return nil, nilPtrErr
	}

	cb := &clipboard{wrapObject(unsafe.Pointer(c))}
	return cb, nil
}

// WaitIsTextAvailable is a wrapper around gtk_clipboard_wait_is_text_available
func (v *clipboard) WaitIsTextAvailable() bool {
	c := C.gtk_clipboard_wait_is_text_available(v.native())
	return gobool(c)
}

// WaitForText is a wrapper around gtk_clipboard_wait_for_text
func (v *clipboard) WaitForText() (string, error) {
	c := C.gtk_clipboard_wait_for_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	defer C.g_free(C.gpointer(c))
	return C.GoString((*C.char)(c)), nil
}

// SetText() is a wrapper around gtk_clipboard_set_text().
func (v *clipboard) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_clipboard_set_text(v.native(), (*C.gchar)(cstr),
		C.gint(len(text)))
}

// WaitIsRichTextAvailable is a wrapper around gtk_clipboard_wait_is_rich_text_available
func (v *clipboard) WaitIsRichTextAvailable(buf gtk.TextBuffer) bool {
	c := C.gtk_clipboard_wait_is_rich_text_available(v.native(), castToTextBuffer(buf).native())
	return gobool(c)
}

// WaitIsUrisAvailable is a wrapper around gtk_clipboard_wait_is_uris_available
func (v *clipboard) WaitIsUrisAvailable() bool {
	c := C.gtk_clipboard_wait_is_uris_available(v.native())
	return gobool(c)
}

// WaitIsImageAvailable is a wrapper around gtk_clipboard_wait_is_image_available
func (v *clipboard) WaitIsImageAvailable() bool {
	c := C.gtk_clipboard_wait_is_image_available(v.native())
	return gobool(c)
}

// SetImage is a wrapper around gtk_clipboard_set_image
func (v *clipboard) SetImage(pixbuf gdk.Pixbuf) {
	C.gtk_clipboard_set_image(v.native(), (*C.GdkPixbuf)(unsafe.Pointer(gdk_impl.CastToPixbuf(pixbuf).Native())))
}

// WaitForImage is a wrapper around gtk_clipboard_wait_for_image
func (v *clipboard) WaitForImage() (gdk.Pixbuf, error) {
	c := C.gtk_clipboard_wait_for_image(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	p := &gdk_impl.Pixbuf{wrapObject(unsafe.Pointer(c))}
	return p, nil
}

// WaitIsTargetAvailable is a wrapper around gtk_clipboard_wait_is_target_available
func (v *clipboard) WaitIsTargetAvailable(target gdk.Atom) bool {
	c := C.gtk_clipboard_wait_is_target_available(v.native(), C.GdkAtom(unsafe.Pointer(target)))
	return gobool(c)
}

// WaitForContents is a wrapper around gtk_clipboard_wait_for_contents
func (v *clipboard) WaitForContents(target gdk.Atom) (gtk.SelectionData, error) {
	c := C.gtk_clipboard_wait_for_contents(v.native(), C.GdkAtom(unsafe.Pointer(target)))
	if c == nil {
		return nil, nilPtrErr
	}
	p := &selectionData{c}
	runtime.SetFinalizer(p, (*selectionData).free)
	return p, nil
}

/*
 * GtkContainer
 */

// Container is a representation of GTK's GtkContainer.
type container struct {
	widget
}

// native returns a pointer to the underlying GtkContainer.
func (v *container) native() *C.GtkContainer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkContainer(p)
}

func marshalContainer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapContainer(obj), nil
}

func wrapContainer(obj *glib_impl.Object) *container {
	return &container{widget{glib_impl.InitiallyUnowned{obj}}}
}

// Add is a wrapper around gtk_container_add().
func (v *container) Add(w gtk.Widget) {
	C.gtk_container_add(v.native(), w.(IWidget).(IWidget).toWidget())
}

// Remove is a wrapper around gtk_container_remove().
func (v *container) Remove(w gtk.Widget) {
	C.gtk_container_remove(v.native(), w.(IWidget).(IWidget).toWidget())
}

// TODO: gtk_container_add_with_properties

// CheckResize is a wrapper around gtk_container_check_resize().
func (v *container) CheckResize() {
	C.gtk_container_check_resize(v.native())
}

// TODO: gtk_container_foreach
// TODO: gtk_container_get_children
// TODO: gtk_container_get_path_for_child

// GetFocusChild is a wrapper around gtk_container_get_focus_child().
func (v *container) GetFocusChild() gtk.Widget {
	c := C.gtk_container_get_focus_child(v.native())
	if c == nil {
		return nil
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapWidget(obj)
}

// SetFocusChild is a wrapper around gtk_container_set_focus_child().
func (v *container) SetFocusChild(child gtk.Widget) {
	C.gtk_container_set_focus_child(v.native(), child.(IWidget).toWidget())
}

// GetFocusVAdjustment is a wrapper around
// gtk_container_get_focus_vadjustment().
func (v *container) GetFocusVAdjustment() gtk.Adjustment {
	c := C.gtk_container_get_focus_vadjustment(v.native())
	if c == nil {
		return nil
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAdjustment(obj)
}

// SetFocusVAdjustment is a wrapper around
// gtk_container_set_focus_vadjustment().
func (v *container) SetFocusVAdjustment(adjustment gtk.Adjustment) {
	C.gtk_container_set_focus_vadjustment(v.native(), castToAdjustment(adjustment).native())
}

// GetFocusHAdjustment is a wrapper around
// gtk_container_get_focus_hadjustment().
func (v *container) GetFocusHAdjustment() gtk.Adjustment {
	c := C.gtk_container_get_focus_hadjustment(v.native())
	if c == nil {
		return nil
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAdjustment(obj)
}

// SetFocusHAdjustment is a wrapper around
// gtk_container_set_focus_hadjustment().
func (v *container) SetFocusHAdjustment(adjustment gtk.Adjustment) {
	C.gtk_container_set_focus_hadjustment(v.native(), castToAdjustment(adjustment).native())
}

// ChildType is a wrapper around gtk_container_child_type().
func (v *container) ChildType() glib.Type {
	c := C.gtk_container_child_type(v.native())
	return glib.Type(c)
}

// TODO: gtk_container_child_get_valist
// TODO: gtk_container_child_set_valist

// ChildNotify is a wrapper around gtk_container_child_notify().
func (v *container) ChildNotify(child gtk.Widget, childProperty string) {
	cstr := C.CString(childProperty)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_container_child_notify(v.native(), child.(IWidget).toWidget(),
		(*C.gchar)(cstr))
}

// ChildSetProperty is a wrapper around gtk_container_child_set_property().
func (v *container) ChildSetProperty(child gtk.Widget, name string, value interface{}) error {
	gv, e := glib_impl.GValue(value)
	if e != nil {
		return e
	}
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_container_child_set_property(v.native(), child.(IWidget).toWidget(), (*C.gchar)(cstr), (*C.GValue)(unsafe.Pointer(gv)))
	return nil
}

// TODO: gtk_container_forall

// GetBorderWidth is a wrapper around gtk_container_get_border_width().
func (v *container) GetBorderWidth() uint {
	c := C.gtk_container_get_border_width(v.native())
	return uint(c)
}

// SetBorderWidth is a wrapper around gtk_container_set_border_width().
func (v *container) SetBorderWidth(borderWidth uint) {
	C.gtk_container_set_border_width(v.native(), C.guint(borderWidth))
}

// PropagateDraw is a wrapper around gtk_container_propagate_draw().
func (v *container) PropagateDraw(child gtk.Widget, cr cairo.Context) {
	context := (*C.cairo_t)(unsafe.Pointer(cairo_impl.CastToContext(cr).Native()))
	C.gtk_container_propagate_draw(v.native(), child.(IWidget).toWidget(), context)
}

// GdkCairoSetSourcePixBuf() is a wrapper around gdk_cairo_set_source_pixbuf().
func GdkCairoSetSourcePixBuf(cr *cairo_impl.Context, pixbuf *gdk_impl.Pixbuf, pixbufX, pixbufY float64) {
	context := (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	ptr := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	C.gdk_cairo_set_source_pixbuf(context, ptr, C.gdouble(pixbufX), C.gdouble(pixbufY))
}

// GetFocusChain is a wrapper around gtk_container_get_focus_chain().
func (v *container) GetFocusChain() ([]gtk.Widget, bool) {
	var cwlist *C.GList
	c := C.gtk_container_get_focus_chain(v.native(), &cwlist)

	var widgets []gtk.Widget
	var wlist glib.List = glib_impl.WrapList(uintptr(unsafe.Pointer(cwlist)))
	for ; wlist.Data() != nil; wlist = wlist.Next() {
		widgets = append(widgets, wrapWidget(wrapObject(wlist.Data().(unsafe.Pointer))))
	}
	return widgets, gobool(c)
}

// SetFocusChain is a wrapper around gtk_container_set_focus_chain().
func (v *container) SetFocusChain(focusableWidgets []gtk.Widget) {
	var list glib.List
	for _, w := range focusableWidgets {
		data := uintptr(unsafe.Pointer(w.(IWidget).toWidget()))
		list = list.Append(data)
	}
	glist := (*C.GList)(unsafe.Pointer(glib_impl.CastToList(list)))
	C.gtk_container_set_focus_chain(v.native(), glist)
}

/*
 * GtkCssProvider
 */

// CssProvider is a representation of GTK's GtkCssProvider.
type cssProvider struct {
	*glib_impl.Object
}

func (v *cssProvider) toStyleProvider() *C.GtkStyleProvider {
	if v == nil {
		return nil
	}
	return C.toGtkStyleProvider(unsafe.Pointer(v.native()))
}

// native returns a pointer to the underlying GtkCssProvider.
func (v *cssProvider) native() *C.GtkCssProvider {
	if v == nil || v.Object == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCssProvider(p)
}

func wrapCssProvider(obj *glib_impl.Object) *cssProvider {
	return &cssProvider{obj}
}

// CssProviderNew is a wrapper around gtk_css_provider_new().
func CssProviderNew() (*cssProvider, error) {
	c := C.gtk_css_provider_new()
	if c == nil {
		return nil, nilPtrErr
	}

	return wrapCssProvider(wrapObject(unsafe.Pointer(c))), nil
}

// LoadFromPath is a wrapper around gtk_css_provider_load_from_path().
func (v *cssProvider) LoadFromPath(path string) error {
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
func (v *cssProvider) LoadFromData(data string) error {
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
func (v *cssProvider) ToString() (string, error) {
	c := C.gtk_css_provider_to_string(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString(c), nil
}

// CssProviderGetDefault is a wrapper around gtk_css_provider_get_default().
func CssProviderGetDefault() (*cssProvider, error) {
	c := C.gtk_css_provider_get_default()
	if c == nil {
		return nil, nilPtrErr
	}

	obj := wrapObject(unsafe.Pointer(c))
	return wrapCssProvider(obj), nil
}

// GetNamed is a wrapper around gtk_css_provider_get_named().
func CssProviderGetNamed(name string, variant string) (*cssProvider, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cvariant := C.CString(variant)
	defer C.free(unsafe.Pointer(cvariant))

	c := C.gtk_css_provider_get_named((*C.gchar)(cname), (*C.gchar)(cvariant))
	if c == nil {
		return nil, nilPtrErr
	}

	obj := wrapObject(unsafe.Pointer(c))
	return wrapCssProvider(obj), nil
}

/*
 * GtkDialog
 */

// Dialog is a representation of GTK's GtkDialog.
type dialog struct {
	window
}

// native returns a pointer to the underlying GtkDialog.
func (v *dialog) native() *C.GtkDialog {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkDialog(p)
}

func marshalDialog(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapDialog(obj), nil
}

func wrapDialog(obj *glib_impl.Object) *dialog {
	return &dialog{window{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}}
}

// DialogNew() is a wrapper around gtk_dialog_new().
func DialogNew() (*dialog, error) {
	c := C.gtk_dialog_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapDialog(obj), nil
}

// Run() is a wrapper around gtk_dialog_run().
func (v *dialog) Run() int {
	c := C.gtk_dialog_run(v.native())
	return int(c)
}

// Response() is a wrapper around gtk_dialog_response().
func (v *dialog) Response(response gtk.ResponseType) {
	C.gtk_dialog_response(v.native(), C.gint(response))
}

// AddButton() is a wrapper around gtk_dialog_add_button().  text may
// be either the literal button text, or if using GTK 3.8 or earlier, a
// Stock type converted to a string.
func (v *dialog) AddButton(text string, id gtk.ResponseType) (gtk.Button, error) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_dialog_add_button(v.native(), (*C.gchar)(cstr), C.gint(id))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return &button{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}, nil
}

// AddActionWidget() is a wrapper around gtk_dialog_add_action_widget().
func (v *dialog) AddActionWidget(child gtk.Widget, id gtk.ResponseType) {
	C.gtk_dialog_add_action_widget(v.native(), child.(IWidget).(IWidget).toWidget(), C.gint(id))
}

// SetDefaultResponse() is a wrapper around gtk_dialog_set_default_response().
func (v *dialog) SetDefaultResponse(id gtk.ResponseType) {
	C.gtk_dialog_set_default_response(v.native(), C.gint(id))
}

// SetResponseSensitive() is a wrapper around
// gtk_dialog_set_response_sensitive().
func (v *dialog) SetResponseSensitive(id gtk.ResponseType, setting bool) {
	C.gtk_dialog_set_response_sensitive(v.native(), C.gint(id),
		gbool(setting))
}

// GetResponseForWidget() is a wrapper around
// gtk_dialog_get_response_for_widget().
func (v *dialog) GetResponseForWidget(widget gtk.Widget) gtk.ResponseType {
	c := C.gtk_dialog_get_response_for_widget(v.native(), widget.(IWidget).toWidget())
	return gtk.ResponseType(c)
}

// GetWidgetForResponse() is a wrapper around
// gtk_dialog_get_widget_for_response().
func (v *dialog) GetWidgetForResponse(id gtk.ResponseType) (gtk.Widget, error) {
	c := C.gtk_dialog_get_widget_for_response(v.native(), C.gint(id))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// GetContentArea() is a wrapper around gtk_dialog_get_content_area().
func (v *dialog) GetContentArea() (gtk.Box, error) {
	c := C.gtk_dialog_get_content_area(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	b := &box{container{widget{glib_impl.InitiallyUnowned{obj}}}}
	return b, nil
}

// TODO(jrick)
/*
func (v *gdk_impl.Screen) AlternativeDialogButtonOrder() bool {
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
type drawingArea struct {
	widget
}

// native returns a pointer to the underlying GtkDrawingArea.
func (v *drawingArea) native() *C.GtkDrawingArea {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkDrawingArea(p)
}

func marshalDrawingArea(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapDrawingArea(obj), nil
}

func wrapDrawingArea(obj *glib_impl.Object) *drawingArea {
	return &drawingArea{widget{glib_impl.InitiallyUnowned{obj}}}
}

// DrawingAreaNew is a wrapper around gtk_drawing_area_new().
func DrawingAreaNew() (*drawingArea, error) {
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
type editable struct {
	*glib_impl.Object
}

// IEditable is an interface type implemented by all structs
// embedding an Editable.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkEditable.
type IEditable interface {
	toEditable() *C.GtkEditable
}

// native() returns a pointer to the underlying GObject as a GtkEditable.
func (v *editable) native() *C.GtkEditable {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkEditable(p)
}

func marshalEditable(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapEditable(obj), nil
}

func wrapEditable(obj *glib_impl.Object) *editable {
	return &editable{obj}
}

func (v *editable) toEditable() *C.GtkEditable {
	if v == nil {
		return nil
	}
	return v.native()
}

// SelectRegion is a wrapper around gtk_editable_select_region().
func (v *editable) SelectRegion(startPos, endPos int) {
	C.gtk_editable_select_region(v.native(), C.gint(startPos),
		C.gint(endPos))
}

// GetSelectionBounds is a wrapper around gtk_editable_get_selection_bounds().
func (v *editable) GetSelectionBounds() (start, end int, nonEmpty bool) {
	var cstart, cend C.gint
	c := C.gtk_editable_get_selection_bounds(v.native(), &cstart, &cend)
	return int(cstart), int(cend), gobool(c)
}

// InsertText is a wrapper around gtk_editable_insert_text(). The returned
// int is the position after the inserted text.
func (v *editable) InsertText(newText string, position int) int {
	cstr := C.CString(newText)
	defer C.free(unsafe.Pointer(cstr))
	pos := new(C.gint)
	*pos = C.gint(position)
	C.gtk_editable_insert_text(v.native(), (*C.gchar)(cstr),
		C.gint(len(newText)), pos)
	return int(*pos)
}

// DeleteText is a wrapper around gtk_editable_delete_text().
func (v *editable) DeleteText(startPos, endPos int) {
	C.gtk_editable_delete_text(v.native(), C.gint(startPos), C.gint(endPos))
}

// GetChars is a wrapper around gtk_editable_get_chars().
func (v *editable) GetChars(startPos, endPos int) string {
	c := C.gtk_editable_get_chars(v.native(), C.gint(startPos),
		C.gint(endPos))
	defer C.free(unsafe.Pointer(c))
	return C.GoString((*C.char)(c))
}

// CutClipboard is a wrapper around gtk_editable_cut_clipboard().
func (v *editable) CutClipboard() {
	C.gtk_editable_cut_clipboard(v.native())
}

// CopyClipboard is a wrapper around gtk_editable_copy_clipboard().
func (v *editable) CopyClipboard() {
	C.gtk_editable_copy_clipboard(v.native())
}

// PasteClipboard is a wrapper around gtk_editable_paste_clipboard().
func (v *editable) PasteClipboard() {
	C.gtk_editable_paste_clipboard(v.native())
}

// DeleteSelection is a wrapper around gtk_editable_delete_selection().
func (v *editable) DeleteSelection() {
	C.gtk_editable_delete_selection(v.native())
}

// SetPosition is a wrapper around gtk_editable_set_position().
func (v *editable) SetPosition(position int) {
	C.gtk_editable_set_position(v.native(), C.gint(position))
}

// GetPosition is a wrapper around gtk_editable_get_position().
func (v *editable) GetPosition() int {
	c := C.gtk_editable_get_position(v.native())
	return int(c)
}

// SetEditable is a wrapper around gtk_editable_set_editable().
func (v *editable) SetEditable(isEditable bool) {
	C.gtk_editable_set_editable(v.native(), gbool(isEditable))
}

// GetEditable is a wrapper around gtk_editable_get_editable().
func (v *editable) GetEditable() bool {
	c := C.gtk_editable_get_editable(v.native())
	return gobool(c)
}

/*
 * GtkEntry
 */

// Entry is a representation of GTK's GtkEntry.
type entry struct {
	widget

	// Interfaces
	editable
}

type IEntry interface {
	toEntry() *C.GtkEntry
}

func (v *entry) toEntry() *C.GtkEntry {
	return v.native()
}

// native returns a pointer to the underlying GtkEntry.
func (v *entry) native() *C.GtkEntry {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkEntry(p)
}

func marshalEntry(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapEntry(obj), nil
}

func wrapEntry(obj *glib_impl.Object) *entry {
	e := wrapEditable(obj)
	return &entry{widget{glib_impl.InitiallyUnowned{obj}}, *e}
}

// EntryNew() is a wrapper around gtk_entry_new().
func EntryNew() (*entry, error) {
	c := C.gtk_entry_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapEntry(obj), nil
}

// EntryNewWithBuffer() is a wrapper around gtk_entry_new_with_buffer().
func EntryNewWithBuffer(buffer *entryBuffer) (*entry, error) {
	c := C.gtk_entry_new_with_buffer(buffer.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapEntry(obj), nil
}

// GetBuffer() is a wrapper around gtk_entry_get_buffer().
func (v *entry) GetBuffer() (gtk.EntryBuffer, error) {
	c := C.gtk_entry_get_buffer(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return &entryBuffer{obj}, nil
}

// SetBuffer() is a wrapper around gtk_entry_set_buffer().
func (v *entry) SetBuffer(buffer gtk.EntryBuffer) {
	C.gtk_entry_set_buffer(v.native(), castToEntryBuffer(buffer).native())
}

// SetText() is a wrapper around gtk_entry_set_text().
func (v *entry) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_set_text(v.native(), (*C.gchar)(cstr))
}

// GetText() is a wrapper around gtk_entry_get_text().
func (v *entry) GetText() (string, error) {
	c := C.gtk_entry_get_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// GetTextLength() is a wrapper around gtk_entry_get_text_length().
func (v *entry) GetTextLength() uint16 {
	c := C.gtk_entry_get_text_length(v.native())
	return uint16(c)
}

// TODO(jrick) GdkRectangle
/*
func (v *entry) GetTextArea() {
}
*/

// SetVisibility() is a wrapper around gtk_entry_set_visibility().
func (v *entry) SetVisibility(visible bool) {
	C.gtk_entry_set_visibility(v.native(), gbool(visible))
}

// SetInvisibleChar() is a wrapper around gtk_entry_set_invisible_char().
func (v *entry) SetInvisibleChar(ch rune) {
	C.gtk_entry_set_invisible_char(v.native(), C.gunichar(ch))
}

// UnsetInvisibleChar() is a wrapper around gtk_entry_unset_invisible_char().
func (v *entry) UnsetInvisibleChar() {
	C.gtk_entry_unset_invisible_char(v.native())
}

// SetMaxLength() is a wrapper around gtk_entry_set_max_length().
func (v *entry) SetMaxLength(len int) {
	C.gtk_entry_set_max_length(v.native(), C.gint(len))
}

// GetActivatesDefault() is a wrapper around gtk_entry_get_activates_default().
func (v *entry) GetActivatesDefault() bool {
	c := C.gtk_entry_get_activates_default(v.native())
	return gobool(c)
}

// GetHasFrame() is a wrapper around gtk_entry_get_has_frame().
func (v *entry) GetHasFrame() bool {
	c := C.gtk_entry_get_has_frame(v.native())
	return gobool(c)
}

// GetWidthChars() is a wrapper around gtk_entry_get_width_chars().
func (v *entry) GetWidthChars() int {
	c := C.gtk_entry_get_width_chars(v.native())
	return int(c)
}

// SetActivatesDefault() is a wrapper around gtk_entry_set_activates_default().
func (v *entry) SetActivatesDefault(setting bool) {
	C.gtk_entry_set_activates_default(v.native(), gbool(setting))
}

// SetHasFrame() is a wrapper around gtk_entry_set_has_frame().
func (v *entry) SetHasFrame(setting bool) {
	C.gtk_entry_set_has_frame(v.native(), gbool(setting))
}

// SetWidthChars() is a wrapper around gtk_entry_set_width_chars().
func (v *entry) SetWidthChars(nChars int) {
	C.gtk_entry_set_width_chars(v.native(), C.gint(nChars))
}

// GetInvisibleChar() is a wrapper around gtk_entry_get_invisible_char().
func (v *entry) GetInvisibleChar() rune {
	c := C.gtk_entry_get_invisible_char(v.native())
	return rune(c)
}

// SetAlignment() is a wrapper around gtk_entry_set_alignment().
func (v *entry) SetAlignment(xalign float32) {
	C.gtk_entry_set_alignment(v.native(), C.gfloat(xalign))
}

// GetAlignment() is a wrapper around gtk_entry_get_alignment().
func (v *entry) GetAlignment() float32 {
	c := C.gtk_entry_get_alignment(v.native())
	return float32(c)
}

// SetPlaceholderText() is a wrapper around gtk_entry_set_placeholder_text().
func (v *entry) SetPlaceholderText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_set_placeholder_text(v.native(), (*C.gchar)(cstr))
}

// GetPlaceholderText() is a wrapper around gtk_entry_get_placeholder_text().
func (v *entry) GetPlaceholderText() (string, error) {
	c := C.gtk_entry_get_placeholder_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// SetOverwriteMode() is a wrapper around gtk_entry_set_overwrite_mode().
func (v *entry) SetOverwriteMode(overwrite bool) {
	C.gtk_entry_set_overwrite_mode(v.native(), gbool(overwrite))
}

// GetOverwriteMode() is a wrapper around gtk_entry_get_overwrite_mode().
func (v *entry) GetOverwriteMode() bool {
	c := C.gtk_entry_get_overwrite_mode(v.native())
	return gobool(c)
}

// TODO(jrick) Pangolayout
/*
func (v *entry) GetLayout() {
}
*/

// GetLayoutOffsets() is a wrapper around gtk_entry_get_layout_offsets().
func (v *entry) GetLayoutOffsets() (x, y int) {
	var gx, gy C.gint
	C.gtk_entry_get_layout_offsets(v.native(), &gx, &gy)
	return int(gx), int(gy)
}

// LayoutIndexToTextIndex() is a wrapper around
// gtk_entry_layout_index_to_text_index().
func (v *entry) LayoutIndexToTextIndex(layoutIndex int) int {
	c := C.gtk_entry_layout_index_to_text_index(v.native(),
		C.gint(layoutIndex))
	return int(c)
}

// TextIndexToLayoutIndex() is a wrapper around
// gtk_entry_text_index_to_layout_index().
func (v *entry) TextIndexToLayoutIndex(textIndex int) int {
	c := C.gtk_entry_text_index_to_layout_index(v.native(),
		C.gint(textIndex))
	return int(c)
}

// TODO(jrick) PandoAttrList
/*
func (v *entry) SetAttributes() {
}
*/

// TODO(jrick) PandoAttrList
/*
func (v *entry) GetAttributes() {
}
*/

// GetMaxLength() is a wrapper around gtk_entry_get_max_length().
func (v *entry) GetMaxLength() int {
	c := C.gtk_entry_get_max_length(v.native())
	return int(c)
}

// GetVisibility() is a wrapper around gtk_entry_get_visibility().
func (v *entry) GetVisibility() bool {
	c := C.gtk_entry_get_visibility(v.native())
	return gobool(c)
}

// SetCompletion() is a wrapper around gtk_entry_set_completion().
func (v *entry) SetCompletion(completion gtk.EntryCompletion) {
	C.gtk_entry_set_completion(v.native(), castToEntryCompletion(completion).native())
}

// GetCompletion() is a wrapper around gtk_entry_get_completion().
func (v *entry) GetCompletion() (gtk.EntryCompletion, error) {
	c := C.gtk_entry_get_completion(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	e := &entryCompletion{wrapObject(unsafe.Pointer(c))}
	return e, nil
}

// SetCursorHAdjustment() is a wrapper around
// gtk_entry_set_cursor_hadjustment().
func (v *entry) SetCursorHAdjustment(adjustment gtk.Adjustment) {
	C.gtk_entry_set_cursor_hadjustment(v.native(), castToAdjustment(adjustment).native())
}

// GetCursorHAdjustment() is a wrapper around
// gtk_entry_get_cursor_hadjustment().
func (v *entry) GetCursorHAdjustment() (gtk.Adjustment, error) {
	c := C.gtk_entry_get_cursor_hadjustment(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return &adjustment{glib_impl.InitiallyUnowned{obj}}, nil
}

// SetProgressFraction() is a wrapper around gtk_entry_set_progress_fraction().
func (v *entry) SetProgressFraction(fraction float64) {
	C.gtk_entry_set_progress_fraction(v.native(), C.gdouble(fraction))
}

// GetProgressFraction() is a wrapper around gtk_entry_get_progress_fraction().
func (v *entry) GetProgressFraction() float64 {
	c := C.gtk_entry_get_progress_fraction(v.native())
	return float64(c)
}

// SetProgressPulseStep() is a wrapper around
// gtk_entry_set_progress_pulse_step().
func (v *entry) SetProgressPulseStep(fraction float64) {
	C.gtk_entry_set_progress_pulse_step(v.native(), C.gdouble(fraction))
}

// GetProgressPulseStep() is a wrapper around
// gtk_entry_get_progress_pulse_step().
func (v *entry) GetProgressPulseStep() float64 {
	c := C.gtk_entry_get_progress_pulse_step(v.native())
	return float64(c)
}

// ProgressPulse() is a wrapper around gtk_entry_progress_pulse().
func (v *entry) ProgressPulse() {
	C.gtk_entry_progress_pulse(v.native())
}

// TODO(jrick) GdkEventKey
/*
func (v *entry) IMContextFilterKeypress() {
}
*/

// ResetIMContext() is a wrapper around gtk_entry_reset_im_context().
func (v *entry) ResetIMContext() {
	C.gtk_entry_reset_im_context(v.native())
}

// TODO(jrick) GdkPixbuf
/*
func (v *entry) SetIconFromPixbuf() {
}
*/

// SetIconFromIconName() is a wrapper around
// gtk_entry_set_icon_from_icon_name().
func (v *entry) SetIconFromIconName(iconPos gtk.EntryIconPosition, name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_set_icon_from_icon_name(v.native(),
		C.GtkEntryIconPosition(iconPos), (*C.gchar)(cstr))
}

// TODO(jrick) GIcon
/*
func (v *entry) SetIconFromGIcon() {
}
*/

// GetIconStorageType() is a wrapper around gtk_entry_get_icon_storage_type().
func (v *entry) GetIconStorageType(iconPos gtk.EntryIconPosition) gtk.ImageType {
	c := C.gtk_entry_get_icon_storage_type(v.native(),
		C.GtkEntryIconPosition(iconPos))
	return gtk.ImageType(c)
}

// TODO(jrick) GdkPixbuf
/*
func (v *entry) GetIconPixbuf() {
}
*/

// GetIconName() is a wrapper around gtk_entry_get_icon_name().
func (v *entry) GetIconName(iconPos gtk.EntryIconPosition) (string, error) {
	c := C.gtk_entry_get_icon_name(v.native(),
		C.GtkEntryIconPosition(iconPos))
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// TODO(jrick) GIcon
/*
func (v *entry) GetIconGIcon() {
}
*/

// SetIconActivatable() is a wrapper around gtk_entry_set_icon_activatable().
func (v *entry) SetIconActivatable(iconPos gtk.EntryIconPosition, activatable bool) {
	C.gtk_entry_set_icon_activatable(v.native(),
		C.GtkEntryIconPosition(iconPos), gbool(activatable))
}

// GetIconActivatable() is a wrapper around gtk_entry_get_icon_activatable().
func (v *entry) GetIconActivatable(iconPos gtk.EntryIconPosition) bool {
	c := C.gtk_entry_get_icon_activatable(v.native(),
		C.GtkEntryIconPosition(iconPos))
	return gobool(c)
}

// SetIconSensitive() is a wrapper around gtk_entry_set_icon_sensitive().
func (v *entry) SetIconSensitive(iconPos gtk.EntryIconPosition, sensitive bool) {
	C.gtk_entry_set_icon_sensitive(v.native(),
		C.GtkEntryIconPosition(iconPos), gbool(sensitive))
}

// GetIconSensitive() is a wrapper around gtk_entry_get_icon_sensitive().
func (v *entry) GetIconSensitive(iconPos gtk.EntryIconPosition) bool {
	c := C.gtk_entry_get_icon_sensitive(v.native(),
		C.GtkEntryIconPosition(iconPos))
	return gobool(c)
}

// GetIconAtPos() is a wrapper around gtk_entry_get_icon_at_pos().
func (v *entry) GetIconAtPos(x, y int) int {
	c := C.gtk_entry_get_icon_at_pos(v.native(), C.gint(x), C.gint(y))
	return int(c)
}

// SetIconTooltipText() is a wrapper around gtk_entry_set_icon_tooltip_text().
func (v *entry) SetIconTooltipText(iconPos gtk.EntryIconPosition, tooltip string) {
	cstr := C.CString(tooltip)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_set_icon_tooltip_text(v.native(),
		C.GtkEntryIconPosition(iconPos), (*C.gchar)(cstr))
}

// GetIconTooltipText() is a wrapper around gtk_entry_get_icon_tooltip_text().
func (v *entry) GetIconTooltipText(iconPos gtk.EntryIconPosition) (string, error) {
	c := C.gtk_entry_get_icon_tooltip_text(v.native(),
		C.GtkEntryIconPosition(iconPos))
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// SetIconTooltipMarkup() is a wrapper around
// gtk_entry_set_icon_tooltip_markup().
func (v *entry) SetIconTooltipMarkup(iconPos gtk.EntryIconPosition, tooltip string) {
	cstr := C.CString(tooltip)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_set_icon_tooltip_markup(v.native(),
		C.GtkEntryIconPosition(iconPos), (*C.gchar)(cstr))
}

// GetIconTooltipMarkup() is a wrapper around
// gtk_entry_get_icon_tooltip_markup().
func (v *entry) GetIconTooltipMarkup(iconPos gtk.EntryIconPosition) (string, error) {
	c := C.gtk_entry_get_icon_tooltip_markup(v.native(),
		C.GtkEntryIconPosition(iconPos))
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// TODO(jrick) GdkDragAction
/*
func (v *entry) SetIconDragSource() {
}
*/

// GetCurrentIconDragSource() is a wrapper around
// gtk_entry_get_current_icon_drag_source().
func (v *entry) GetCurrentIconDragSource() int {
	c := C.gtk_entry_get_current_icon_drag_source(v.native())
	return int(c)
}

// TODO(jrick) GdkRectangle
/*
func (v *entry) GetIconArea() {
}
*/

// SetInputPurpose() is a wrapper around gtk_entry_set_input_purpose().
func (v *entry) SetInputPurpose(purpose gtk.InputPurpose) {
	C.gtk_entry_set_input_purpose(v.native(), C.GtkInputPurpose(purpose))
}

// GetInputPurpose() is a wrapper around gtk_entry_get_input_purpose().
func (v *entry) GetInputPurpose() gtk.InputPurpose {
	c := C.gtk_entry_get_input_purpose(v.native())
	return gtk.InputPurpose(c)
}

// SetInputHints() is a wrapper around gtk_entry_set_input_hints().
func (v *entry) SetInputHints(hints gtk.InputHints) {
	C.gtk_entry_set_input_hints(v.native(), C.GtkInputHints(hints))
}

// GetInputHints() is a wrapper around gtk_entry_get_input_hints().
func (v *entry) GetInputHints() gtk.InputHints {
	c := C.gtk_entry_get_input_hints(v.native())
	return gtk.InputHints(c)
}

/*
 * GtkEntryBuffer
 */

// EntryBuffer is a representation of GTK's GtkEntryBuffer.
type entryBuffer struct {
	*glib_impl.Object
}

// native returns a pointer to the underlying GtkEntryBuffer.
func (v *entryBuffer) native() *C.GtkEntryBuffer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkEntryBuffer(p)
}

func marshalEntryBuffer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapEntryBuffer(obj), nil
}

func wrapEntryBuffer(obj *glib_impl.Object) *entryBuffer {
	return &entryBuffer{obj}
}

// EntryBufferNew() is a wrapper around gtk_entry_buffer_new().
func EntryBufferNew(initialChars string, nInitialChars int) (*entryBuffer, error) {
	cstr := C.CString(initialChars)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_entry_buffer_new((*C.gchar)(cstr), C.gint(nInitialChars))
	if c == nil {
		return nil, nilPtrErr
	}

	e := wrapEntryBuffer(wrapObject(unsafe.Pointer(c)))
	return e, nil
}

// GetText() is a wrapper around gtk_entry_buffer_get_text().  A
// non-nil error is returned in the case that gtk_entry_buffer_get_text
// returns NULL to differentiate between NULL and an empty string.
func (v *entryBuffer) GetText() (string, error) {
	c := C.gtk_entry_buffer_get_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// SetText() is a wrapper around gtk_entry_buffer_set_text().
func (v *entryBuffer) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_buffer_set_text(v.native(), (*C.gchar)(cstr),
		C.gint(len(text)))
}

// GetBytes() is a wrapper around gtk_entry_buffer_get_bytes().
func (v *entryBuffer) GetBytes() uint {
	c := C.gtk_entry_buffer_get_bytes(v.native())
	return uint(c)
}

// GetLength() is a wrapper around gtk_entry_buffer_get_length().
func (v *entryBuffer) GetLength() uint {
	c := C.gtk_entry_buffer_get_length(v.native())
	return uint(c)
}

// GetMaxLength() is a wrapper around gtk_entry_buffer_get_max_length().
func (v *entryBuffer) GetMaxLength() int {
	c := C.gtk_entry_buffer_get_max_length(v.native())
	return int(c)
}

// SetMaxLength() is a wrapper around gtk_entry_buffer_set_max_length().
func (v *entryBuffer) SetMaxLength(maxLength int) {
	C.gtk_entry_buffer_set_max_length(v.native(), C.gint(maxLength))
}

// InsertText() is a wrapper around gtk_entry_buffer_insert_text().
func (v *entryBuffer) InsertText(position uint, text string) uint {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_entry_buffer_insert_text(v.native(), C.guint(position),
		(*C.gchar)(cstr), C.gint(len(text)))
	return uint(c)
}

// DeleteText() is a wrapper around gtk_entry_buffer_delete_text().
func (v *entryBuffer) DeleteText(position uint, nChars int) uint {
	c := C.gtk_entry_buffer_delete_text(v.native(), C.guint(position),
		C.gint(nChars))
	return uint(c)
}

// EmitDeletedText() is a wrapper around gtk_entry_buffer_emit_deleted_text().
func (v *entryBuffer) EmitDeletedText(pos, nChars uint) {
	C.gtk_entry_buffer_emit_deleted_text(v.native(), C.guint(pos),
		C.guint(nChars))
}

// EmitInsertedText() is a wrapper around gtk_entry_buffer_emit_inserted_text().
func (v *entryBuffer) EmitInsertedText(pos uint, text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_buffer_emit_inserted_text(v.native(), C.guint(pos),
		(*C.gchar)(cstr), C.guint(len(text)))
}

/*
 * GtkEntryCompletion
 */

// EntryCompletion is a representation of GTK's GtkEntryCompletion.
type entryCompletion struct {
	*glib_impl.Object
}

// native returns a pointer to the underlying GtkEntryCompletion.
func (v *entryCompletion) native() *C.GtkEntryCompletion {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkEntryCompletion(p)
}

func marshalEntryCompletion(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapEntryCompletion(obj), nil
}

func wrapEntryCompletion(obj *glib_impl.Object) *entryCompletion {
	return &entryCompletion{obj}
}

/*
 * GtkEventBox
 */

// EventBox is a representation of GTK's GtkEventBox.
type eventBox struct {
	bin
}

// native returns a pointer to the underlying GtkEventBox.
func (v *eventBox) native() *C.GtkEventBox {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkEventBox(p)
}

func marshalEventBox(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapEventBox(obj), nil
}

func wrapEventBox(obj *glib_impl.Object) *eventBox {
	return &eventBox{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}
}

// EventBoxNew is a wrapper around gtk_event_box_new().
func EventBoxNew() (*eventBox, error) {
	c := C.gtk_event_box_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapEventBox(obj), nil
}

// SetAboveChild is a wrapper around gtk_event_box_set_above_child().
func (v *eventBox) SetAboveChild(aboveChild bool) {
	C.gtk_event_box_set_above_child(v.native(), gbool(aboveChild))
}

// GetAboveChild is a wrapper around gtk_event_box_get_above_child().
func (v *eventBox) GetAboveChild() bool {
	c := C.gtk_event_box_get_above_child(v.native())
	return gobool(c)
}

// SetVisibleWindow is a wrapper around gtk_event_box_set_visible_window().
func (v *eventBox) SetVisibleWindow(visibleWindow bool) {
	C.gtk_event_box_set_visible_window(v.native(), gbool(visibleWindow))
}

// GetVisibleWindow is a wrapper around gtk_event_box_get_visible_window().
func (v *eventBox) GetVisibleWindow() bool {
	c := C.gtk_event_box_get_visible_window(v.native())
	return gobool(c)
}

/*
 * GtkExpander
 */

// Expander is a representation of GTK's GtkExpander.
type expander struct {
	bin
}

// native returns a pointer to the underlying GtkExpander.
func (v *expander) native() *C.GtkExpander {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkExpander(p)
}

func marshalExpander(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapExpander(obj), nil
}

func wrapExpander(obj *glib_impl.Object) *expander {
	return &expander{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}
}

// ExpanderNew is a wrapper around gtk_expander_new().
func ExpanderNew(label string) (*expander, error) {
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
func (v *expander) SetExpanded(expanded bool) {
	C.gtk_expander_set_expanded(v.native(), gbool(expanded))
}

// GetExpanded is a wrapper around gtk_expander_get_expanded().
func (v *expander) GetExpanded() bool {
	c := C.gtk_expander_get_expanded(v.native())
	return gobool(c)
}

// SetLabel is a wrapper around gtk_expander_set_label().
func (v *expander) SetLabel(label string) {
	var cstr *C.char
	if label != "" {
		cstr = C.CString(label)
		defer C.free(unsafe.Pointer(cstr))
	}
	C.gtk_expander_set_label(v.native(), (*C.gchar)(cstr))
}

// GetLabel is a wrapper around gtk_expander_get_label().
func (v *expander) GetLabel() string {
	c := C.gtk_expander_get_label(v.native())
	return C.GoString((*C.char)(c))
}

// SetLabelWidget is a wrapper around gtk_expander_set_label_widget().
func (v *expander) SetLabelWidget(widget gtk.Widget) {
	C.gtk_expander_set_label_widget(v.native(), widget.(IWidget).toWidget())
}

/*
 * GtkFileChooser
 */

// FileChoser is a representation of GTK's GtkFileChooser GInterface.
type fileChooser struct {
	*glib_impl.Object
}

// native returns a pointer to the underlying GObject as a GtkFileChooser.
func (v *fileChooser) native() *C.GtkFileChooser {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFileChooser(p)
}

func marshalFileChooser(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapFileChooser(obj), nil
}

func wrapFileChooser(obj *glib_impl.Object) *fileChooser {
	return &fileChooser{obj}
}

// GetFilename is a wrapper around gtk_file_chooser_get_filename().
func (v *fileChooser) GetFilename() string {
	c := C.gtk_file_chooser_get_filename(v.native())
	s := C.GoString((*C.char)(c))
	defer C.g_free((C.gpointer)(c))
	return s
}

// SetCurrentName is a wrapper around gtk_file_chooser_set_current_name().
func (v *fileChooser) SetCurrentName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_file_chooser_set_current_name(v.native(), (*C.gchar)(cstr))
	return
}

// SetCurrentFolder is a wrapper around gtk_file_chooser_set_current_folder().
func (v *fileChooser) SetCurrentFolder(folder string) bool {
	cstr := C.CString(folder)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_file_chooser_set_current_folder(v.native(), (*C.gchar)(cstr))
	return gobool(c)
}

// GetCurrentFolder is a wrapper around gtk_file_chooser_get_current_folder().
func (v *fileChooser) GetCurrentFolder() (string, error) {
	c := C.gtk_file_chooser_get_current_folder(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	defer C.free(unsafe.Pointer(c))
	return C.GoString((*C.char)(c)), nil
}

// SetPreviewWidget is a wrapper around gtk_file_chooser_set_preview_widget().
func (v *fileChooser) SetPreviewWidget(widget gtk.Widget) {
	C.gtk_file_chooser_set_preview_widget(v.native(), widget.(IWidget).toWidget())
}

// SetPreviewWidgetActive is a wrapper around gtk_file_chooser_set_preview_widget_active().
func (v *fileChooser) SetPreviewWidgetActive(active bool) {
	C.gtk_file_chooser_set_preview_widget_active(v.native(), gbool(active))
}

// GetPreviewFilename is a wrapper around gtk_file_chooser_get_preview_filename().
func (v *fileChooser) GetPreviewFilename() string {
	c := C.gtk_file_chooser_get_preview_filename(v.native())
	defer C.free(unsafe.Pointer(c))
	return C.GoString(c)
}

// AddFilter is a wrapper around gtk_file_chooser_add_filter().
func (v *fileChooser) AddFilter(filter gtk.FileFilter) {
	C.gtk_file_chooser_add_filter(v.native(), castToFileFilter(filter).native())
}

// GetURI is a wrapper around gtk_file_chooser_get_uri().
func (v *fileChooser) GetURI() string {
	c := C.gtk_file_chooser_get_uri(v.native())
	s := C.GoString((*C.char)(c))
	defer C.g_free((C.gpointer)(c))
	return s
}

// AddShortcutFolder is a wrapper around gtk_file_chooser_add_shortcut_folder().
func (v *fileChooser) AddShortcutFolder(folder string) bool {
	cstr := C.CString(folder)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_file_chooser_add_shortcut_folder(v.native(), cstr, nil)
	return gobool(c)
}

/*
 * GtkFileChooserButton
 */

// FileChooserButton is a representation of GTK's GtkFileChooserButton.
type fileChooserButton struct {
	box

	// Interfaces
	fileChooser
}

// native returns a pointer to the underlying GtkFileChooserButton.
func (v *fileChooserButton) native() *C.GtkFileChooserButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFileChooserButton(p)
}

func marshalFileChooserButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapFileChooserButton(obj), nil
}

func wrapFileChooserButton(obj *glib_impl.Object) *fileChooserButton {
	fc := wrapFileChooser(obj)
	return &fileChooserButton{box{container{widget{glib_impl.InitiallyUnowned{obj}}}}, *fc}
}

// FileChooserButtonNew is a wrapper around gtk_file_chooser_button_new().
func FileChooserButtonNew(title string, action gtk.FileChooserAction) (*fileChooserButton, error) {
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
type fileChooserDialog struct {
	dialog

	// Interfaces
	fileChooser
}

// native returns a pointer to the underlying GtkFileChooserDialog.
func (v *fileChooserDialog) native() *C.GtkFileChooserDialog {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFileChooserDialog(p)
}

func marshalFileChooserDialog(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapFileChooserDialog(obj), nil
}

func wrapFileChooserDialog(obj *glib_impl.Object) *fileChooserDialog {
	fc := wrapFileChooser(obj)
	return &fileChooserDialog{dialog{window{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}}, *fc}
}

// FileChooserDialogNewWith1Button is a wrapper around gtk_file_chooser_dialog_new() with one button.
func FileChooserDialogNewWith1Button(
	title string,
	parent *window,
	action gtk.FileChooserAction,
	first_button_text string,
	first_button_id gtk.ResponseType) (*fileChooserDialog, error) {
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
	parent *window,
	action gtk.FileChooserAction,
	first_button_text string,
	first_button_id gtk.ResponseType,
	second_button_text string,
	second_button_id gtk.ResponseType) (*fileChooserDialog, error) {
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
type fileChooserWidget struct {
	box

	// Interfaces
	fileChooser
}

// native returns a pointer to the underlying GtkFileChooserWidget.
func (v *fileChooserWidget) native() *C.GtkFileChooserWidget {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFileChooserWidget(p)
}

func marshalFileChooserWidget(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapFileChooserWidget(obj), nil
}

func wrapFileChooserWidget(obj *glib_impl.Object) *fileChooserWidget {
	fc := wrapFileChooser(obj)
	return &fileChooserWidget{box{container{widget{glib_impl.InitiallyUnowned{obj}}}}, *fc}
}

// FileChooserWidgetNew is a wrapper around gtk_file_chooser_widget_new().
func FileChooserWidgetNew(action gtk.FileChooserAction) (*fileChooserWidget, error) {
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
type fileFilter struct {
	*glib_impl.Object
}

// native returns a pointer to the underlying GObject as a GtkFileFilter.
func (v *fileFilter) native() *C.GtkFileFilter {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFileFilter(p)
}

func marshalFileFilter(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapFileFilter(obj), nil
}

func wrapFileFilter(obj *glib_impl.Object) *fileFilter {
	return &fileFilter{obj}
}

// FileFilterNew is a wrapper around gtk_file_filter_new().
func FileFilterNew() (*fileFilter, error) {
	c := C.gtk_file_filter_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapFileFilter(obj), nil
}

// SetName is a wrapper around gtk_file_filter_set_name().
func (v *fileFilter) SetName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_file_filter_set_name(v.native(), (*C.gchar)(cstr))
}

// AddPattern is a wrapper around gtk_file_filter_add_pattern().
func (v *fileFilter) AddPattern(pattern string) {
	cstr := C.CString(pattern)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_file_filter_add_pattern(v.native(), (*C.gchar)(cstr))
}

// AddPixbufFormats is a wrapper around gtk_file_filter_add_pixbuf_formats().
func (v *fileFilter) AddPixbufFormats() {
	C.gtk_file_filter_add_pixbuf_formats(v.native())
}

/*
 * GtkFontButton
 */

// FontButton is a representation of GTK's GtkFontButton.
type fontButton struct {
	button
}

// native returns a pointer to the underlying GtkFontButton.
func (v *fontButton) native() *C.GtkFontButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFontButton(p)
}

func marshalFontButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapFontButton(obj), nil
}

func wrapFontButton(obj *glib_impl.Object) *fontButton {
	return &fontButton{button{bin{container{widget{
		glib_impl.InitiallyUnowned{obj}}}}}}
}

// FontButtonNew is a wrapper around gtk_font_button_new().
func FontButtonNew() (*fontButton, error) {
	c := C.gtk_font_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapFontButton(obj), nil
}

// FontButtonNewWithFont is a wrapper around gtk_font_button_new_with_font().
func FontButtonNewWithFont(fontname string) (*fontButton, error) {
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
func (v *fontButton) GetFontName() string {
	c := C.gtk_font_button_get_font_name(v.native())
	return C.GoString((*C.char)(c))
}

// SetFontName is a wrapper around gtk_font_button_set_font_name().
func (v *fontButton) SetFontName(fontname string) bool {
	cstr := C.CString(fontname)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_font_button_set_font_name(v.native(), (*C.gchar)(cstr))
	return gobool(c)
}

/*
 * GtkFrame
 */

// Frame is a representation of GTK's GtkFrame.
type frame struct {
	bin
}

// native returns a pointer to the underlying GtkFrame.
func (v *frame) native() *C.GtkFrame {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFrame(p)
}

func marshalFrame(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapFrame(obj), nil
}

func wrapFrame(obj *glib_impl.Object) *frame {
	return &frame{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}
}

// FrameNew is a wrapper around gtk_frame_new().
func FrameNew(label string) (*frame, error) {
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
func (v *frame) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_frame_set_label(v.native(), (*C.gchar)(cstr))
}

// SetLabelWidget is a wrapper around gtk_frame_set_label_widget().
func (v *frame) SetLabelWidget(labelWidget gtk.Widget) {
	C.gtk_frame_set_label_widget(v.native(), labelWidget.(IWidget).toWidget())
}

// SetLabelAlign is a wrapper around gtk_frame_set_label_align().
func (v *frame) SetLabelAlign(xAlign, yAlign float32) {
	C.gtk_frame_set_label_align(v.native(), C.gfloat(xAlign),
		C.gfloat(yAlign))
}

// SetShadowType is a wrapper around gtk_frame_set_shadow_type().
func (v *frame) SetShadowType(t gtk.ShadowType) {
	C.gtk_frame_set_shadow_type(v.native(), C.GtkShadowType(t))
}

// GetLabel is a wrapper around gtk_frame_get_label().
func (v *frame) GetLabel() string {
	c := C.gtk_frame_get_label(v.native())
	return C.GoString((*C.char)(c))
}

// GetLabelAlign is a wrapper around gtk_frame_get_label_align().
func (v *frame) GetLabelAlign() (xAlign, yAlign float32) {
	var x, y C.gfloat
	C.gtk_frame_get_label_align(v.native(), &x, &y)
	return float32(x), float32(y)
}

// GetLabelWidget is a wrapper around gtk_frame_get_label_widget().
func (v *frame) GetLabelWidget() (gtk.Widget, error) {
	c := C.gtk_frame_get_label_widget(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// GetShadowType is a wrapper around gtk_frame_get_shadow_type().
func (v *frame) GetShadowType() gtk.ShadowType {
	c := C.gtk_frame_get_shadow_type(v.native())
	return gtk.ShadowType(c)
}

/*
 * GtkGrid
 */

// Grid is a representation of GTK's GtkGrid.
type grid struct {
	container

	// Interfaces
	orientable
}

// native returns a pointer to the underlying GtkGrid.
func (v *grid) native() *C.GtkGrid {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkGrid(p)
}

func (v *grid) toOrientable() *C.GtkOrientable {
	if v == nil {
		return nil
	}
	return C.toGtkOrientable(unsafe.Pointer(v.GObject))
}

func marshalGrid(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapGrid(obj), nil
}

func wrapGrid(obj *glib_impl.Object) *grid {
	o := wrapOrientable(obj)
	return &grid{container{widget{glib_impl.InitiallyUnowned{obj}}}, *o}
}

// GridNew() is a wrapper around gtk_grid_new().
func GridNew() (*grid, error) {
	c := C.gtk_grid_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapGrid(obj), nil
}

// Attach() is a wrapper around gtk_grid_attach().
func (v *grid) Attach(child gtk.Widget, left, top, width, height int) {
	C.gtk_grid_attach(v.native(), child.(IWidget).toWidget(), C.gint(left),
		C.gint(top), C.gint(width), C.gint(height))
}

// AttachNextTo() is a wrapper around gtk_grid_attach_next_to().
func (v *grid) AttachNextTo(child, sibling gtk.Widget, side gtk.PositionType, width, height int) {
	C.gtk_grid_attach_next_to(v.native(), child.(IWidget).toWidget(),
		sibling.(IWidget).toWidget(), C.GtkPositionType(side), C.gint(width),
		C.gint(height))
}

// GetChildAt() is a wrapper around gtk_grid_get_child_at().
func (v *grid) GetChildAt(left, top int) (gtk.Widget, error) {
	c := C.gtk_grid_get_child_at(v.native(), C.gint(left), C.gint(top))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// InsertRow() is a wrapper around gtk_grid_insert_row().
func (v *grid) InsertRow(position int) {
	C.gtk_grid_insert_row(v.native(), C.gint(position))
}

// InsertColumn() is a wrapper around gtk_grid_insert_column().
func (v *grid) InsertColumn(position int) {
	C.gtk_grid_insert_column(v.native(), C.gint(position))
}

// InsertNextTo() is a wrapper around gtk_grid_insert_next_to()
func (v *grid) InsertNextTo(sibling gtk.Widget, side gtk.PositionType) {
	C.gtk_grid_insert_next_to(v.native(), sibling.(IWidget).toWidget(),
		C.GtkPositionType(side))
}

// SetRowHomogeneous() is a wrapper around gtk_grid_set_row_homogeneous().
func (v *grid) SetRowHomogeneous(homogeneous bool) {
	C.gtk_grid_set_row_homogeneous(v.native(), gbool(homogeneous))
}

// GetRowHomogeneous() is a wrapper around gtk_grid_get_row_homogeneous().
func (v *grid) GetRowHomogeneous() bool {
	c := C.gtk_grid_get_row_homogeneous(v.native())
	return gobool(c)
}

// SetRowSpacing() is a wrapper around gtk_grid_set_row_spacing().
func (v *grid) SetRowSpacing(spacing uint) {
	C.gtk_grid_set_row_spacing(v.native(), C.guint(spacing))
}

// GetRowSpacing() is a wrapper around gtk_grid_get_row_spacing().
func (v *grid) GetRowSpacing() uint {
	c := C.gtk_grid_get_row_spacing(v.native())
	return uint(c)
}

// SetColumnHomogeneous() is a wrapper around gtk_grid_set_column_homogeneous().
func (v *grid) SetColumnHomogeneous(homogeneous bool) {
	C.gtk_grid_set_column_homogeneous(v.native(), gbool(homogeneous))
}

// GetColumnHomogeneous() is a wrapper around gtk_grid_get_column_homogeneous().
func (v *grid) GetColumnHomogeneous() bool {
	c := C.gtk_grid_get_column_homogeneous(v.native())
	return gobool(c)
}

// SetColumnSpacing() is a wrapper around gtk_grid_set_column_spacing().
func (v *grid) SetColumnSpacing(spacing uint) {
	C.gtk_grid_set_column_spacing(v.native(), C.guint(spacing))
}

// GetColumnSpacing() is a wrapper around gtk_grid_get_column_spacing().
func (v *grid) GetColumnSpacing() uint {
	c := C.gtk_grid_get_column_spacing(v.native())
	return uint(c)
}

/*
 * GtkIconTheme
 */

// IconTheme is a representation of GTK's GtkIconTheme
type iconTheme struct {
	Theme *C.GtkIconTheme
}

// IconThemeGetDefault is a wrapper around gtk_icon_theme_get_default().
func IconThemeGetDefault() (*iconTheme, error) {
	c := C.gtk_icon_theme_get_default()
	if c == nil {
		return nil, nilPtrErr
	}
	return &iconTheme{c}, nil
}

// IconThemeGetForScreen is a wrapper around gtk_icon_theme_get_for_screen().
func IconThemeGetForScreen(screen *gdk_impl.Screen) (*iconTheme, error) {
	cScreen := (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	c := C.gtk_icon_theme_get_for_screen(cScreen)
	if c == nil {
		return nil, nilPtrErr
	}
	return &iconTheme{c}, nil
}

// LoadIcon is a wrapper around gtk_icon_theme_load_icon().
func (v *iconTheme) LoadIcon(iconName string, size int, flags gtk.IconLookupFlags) (gdk.Pixbuf, error) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	c := C.gtk_icon_theme_load_icon(v.Theme, (*C.gchar)(cstr), C.gint(size), C.GtkIconLookupFlags(flags), &err)
	if c == nil {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}
	return &gdk_impl.Pixbuf{wrapObject(unsafe.Pointer(c))}, nil
}

/*
 * GtkIconView
 */

// IconView is a representation of GTK's GtkIconView.
type iconView struct {
	container
}

// native returns a pointer to the underlying GtkIconView.
func (v *iconView) native() *C.GtkIconView {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkIconView(p)
}

func marshalIconView(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapIconView(obj), nil
}

func wrapIconView(obj *glib_impl.Object) *iconView {
	return &iconView{container{widget{glib_impl.InitiallyUnowned{obj}}}}
}

// IconViewNew is a wrapper around gtk_icon_view_new().
func IconViewNew() (*iconView, error) {
	c := C.gtk_icon_view_new()
	if c == nil {
		return nil, nilPtrErr
	}

	return wrapIconView(wrapObject(unsafe.Pointer(c))), nil
}

// IconViewNewWithModel is a wrapper around gtk_icon_view_new_with_model().
func IconViewNewWithModel(model ITreeModel) (*iconView, error) {
	c := C.gtk_icon_view_new_with_model(model.toTreeModel())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapIconView(obj), nil
}

// GetModel is a wrapper around gtk_icon_view_get_model().
func (v *iconView) GetModel() (gtk.TreeModel, error) {
	c := C.gtk_icon_view_get_model(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapTreeModel(obj), nil
}

// SetModel is a wrapper around gtk_icon_view_set_model().
func (v *iconView) SetModel(model gtk.TreeModel) {
	C.gtk_icon_view_set_model(v.native(), model.(ITreeModel).toTreeModel())
}

// SelectPath is a wrapper around gtk_icon_view_select_path().
func (v *iconView) SelectPath(path gtk.TreePath) {
	C.gtk_icon_view_select_path(v.native(), castToTreePath(path).native())
}

// ScrollToPath is a wrapper around gtk_icon_view_scroll_to_path().
func (v *iconView) ScrollToPath(path gtk.TreePath, useAlign bool, rowAlign, colAlign float64) {
	C.gtk_icon_view_scroll_to_path(v.native(), castToTreePath(path).native(), gbool(useAlign),
		C.gfloat(rowAlign), C.gfloat(colAlign))
}

/*
 * GtkImage
 */

// Image is a representation of GTK's GtkImage.
type image struct {
	widget
}

// native returns a pointer to the underlying GtkImage.
func (v *image) native() *C.GtkImage {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkImage(p)
}

func marshalImage(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapImage(obj), nil
}

func wrapImage(obj *glib_impl.Object) *image {
	return &image{widget{glib_impl.InitiallyUnowned{obj}}}
}

// ImageNew() is a wrapper around gtk_image_new().
func ImageNew() (*image, error) {
	c := C.gtk_image_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapImage(obj), nil
}

// ImageNewFromFile() is a wrapper around gtk_image_new_from_file().
func ImageNewFromFile(filename string) (*image, error) {
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
func ImageNewFromResource(resourcePath string) (*image, error) {
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
func ImageNewFromPixbuf(pixbuf *gdk_impl.Pixbuf) (*image, error) {
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
func ImageNewFromIconName(iconName string, size gtk.IconSize) (*image, error) {
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
func (v *image) Clear() {
	C.gtk_image_clear(v.native())
}

// SetFromFile() is a wrapper around gtk_image_set_from_file().
func (v *image) SetFromFile(filename string) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_image_set_from_file(v.native(), (*C.gchar)(cstr))
}

// SetFromResource() is a wrapper around gtk_image_set_from_resource().
func (v *image) SetFromResource(resourcePath string) {
	cstr := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_image_set_from_resource(v.native(), (*C.gchar)(cstr))
}

// SetFromFixbuf is a wrapper around gtk_image_set_from_pixbuf().
func (v *image) SetFromPixbuf(pixbuf gdk.Pixbuf) {
	pbptr := (*C.GdkPixbuf)(unsafe.Pointer(gdk_impl.CastToPixbuf(pixbuf).Native()))
	C.gtk_image_set_from_pixbuf(v.native(), pbptr)
}

// TODO(jrick) GtkIconSet
/*
func (v *image) SetFromIconSet() {
}
*/

// TODO(jrick) GdkPixbufAnimation
/*
func (v *image) SetFromAnimation() {
}
*/

// SetFromIconName() is a wrapper around gtk_image_set_from_icon_name().
func (v *image) SetFromIconName(iconName string, size gtk.IconSize) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_image_set_from_icon_name(v.native(), (*C.gchar)(cstr),
		C.GtkIconSize(size))
}

// TODO(jrick) GIcon
/*
func (v *image) SetFromGIcon() {
}
*/

// SetPixelSize() is a wrapper around gtk_image_set_pixel_size().
func (v *image) SetPixelSize(pixelSize int) {
	C.gtk_image_set_pixel_size(v.native(), C.gint(pixelSize))
}

// GetStorageType() is a wrapper around gtk_image_get_storage_type().
func (v *image) GetStorageType() gtk.ImageType {
	c := C.gtk_image_get_storage_type(v.native())
	return gtk.ImageType(c)
}

// GetPixbuf() is a wrapper around gtk_image_get_pixbuf().
func (v *image) GetPixbuf() gdk.Pixbuf {
	c := C.gtk_image_get_pixbuf(v.native())
	if c == nil {
		return nil
	}

	pb := &gdk_impl.Pixbuf{wrapObject(unsafe.Pointer(c))}
	return pb
}

// TODO(jrick) GtkIconSet
/*
func (v *image) GetIconSet() {
}
*/

// TODO(jrick) GdkPixbufAnimation
/*
func (v *image) GetAnimation() {
}
*/

// GetIconName() is a wrapper around gtk_image_get_icon_name().
func (v *image) GetIconName() (string, gtk.IconSize) {
	var iconName *C.gchar
	var size C.GtkIconSize
	C.gtk_image_get_icon_name(v.native(), &iconName, &size)
	return C.GoString((*C.char)(iconName)), gtk.IconSize(size)
}

// TODO(jrick) GIcon
/*
func (v *image) GetGIcon() {
}
*/

// GetPixelSize() is a wrapper around gtk_image_get_pixel_size().
func (v *image) GetPixelSize() int {
	c := C.gtk_image_get_pixel_size(v.native())
	return int(c)
}

// added by terrak
/*
 * GtkLayout
 */

// Layout is a representation of GTK's GtkLayout.
type layout struct {
	container
}

// native returns a pointer to the underlying GtkDrawingArea.
func (v *layout) native() *C.GtkLayout {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkLayout(p)
}

func marshalLayout(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapLayout(obj), nil
}

func wrapLayout(obj *glib_impl.Object) *layout {
	return &layout{container{widget{glib_impl.InitiallyUnowned{obj}}}}
}

// LayoutNew is a wrapper around gtk_layout_new().
func LayoutNew(hadjustment, vadjustment *adjustment) (*layout, error) {
	c := C.gtk_layout_new(hadjustment.native(), vadjustment.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapLayout(obj), nil
}

// Layout.Put is a wrapper around gtk_layout_put().
func (v *layout) Put(w gtk.Widget, x, y int) {
	C.gtk_layout_put(v.native(), w.(IWidget).toWidget(), C.gint(x), C.gint(y))
}

// Layout.Move is a wrapper around gtk_layout_move().
func (v *layout) Move(w gtk.Widget, x, y int) {
	C.gtk_layout_move(v.native(), w.(IWidget).toWidget(), C.gint(x), C.gint(y))
}

// Layout.SetSize is a wrapper around gtk_layout_set_size
func (v *layout) SetSize(width, height uint) {
	C.gtk_layout_set_size(v.native(), C.guint(width), C.guint(height))
}

// Layout.GetSize is a wrapper around gtk_layout_get_size
func (v *layout) GetSize() (width, height uint) {
	var w, h C.guint
	C.gtk_layout_get_size(v.native(), &w, &h)
	return uint(w), uint(h)
}

/*
 * GtkLinkButton
 */

// LinkButton is a representation of GTK's GtkLinkButton.
type linkButton struct {
	button
}

// native returns a pointer to the underlying GtkLinkButton.
func (v *linkButton) native() *C.GtkLinkButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkLinkButton(p)
}

func marshalLinkButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapLinkButton(obj), nil
}

func wrapLinkButton(obj *glib_impl.Object) *linkButton {
	return &linkButton{button{bin{container{widget{
		glib_impl.InitiallyUnowned{obj}}}}}}
}

// LinkButtonNew is a wrapper around gtk_link_button_new().
func LinkButtonNew(label string) (*linkButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_link_button_new((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapLinkButton(wrapObject(unsafe.Pointer(c))), nil
}

// LinkButtonNewWithLabel is a wrapper around gtk_link_button_new_with_label().
func LinkButtonNewWithLabel(uri, label string) (*linkButton, error) {
	curi := C.CString(uri)
	defer C.free(unsafe.Pointer(curi))
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	c := C.gtk_link_button_new_with_label((*C.gchar)(curi), (*C.gchar)(clabel))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapLinkButton(wrapObject(unsafe.Pointer(c))), nil
}

// GetUri is a wrapper around gtk_link_button_get_uri().
func (v *linkButton) GetUri() string {
	c := C.gtk_link_button_get_uri(v.native())
	return C.GoString((*C.char)(c))
}

// SetUri is a wrapper around gtk_link_button_set_uri().
func (v *linkButton) SetUri(uri string) {
	cstr := C.CString(uri)
	C.gtk_link_button_set_uri(v.native(), (*C.gchar)(cstr))
}

/*
 * GtkListStore
 */

// ListStore is a representation of GTK's GtkListStore.
type listStore struct {
	*glib_impl.Object

	// Interfaces
	treeModel
}

// native returns a pointer to the underlying GtkListStore.
func (v *listStore) native() *C.GtkListStore {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkListStore(p)
}

func marshalListStore(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapListStore(obj), nil
}

func wrapListStore(obj *glib_impl.Object) *listStore {
	tm := wrapTreeModel(obj)
	return &listStore{obj, *tm}
}

func (v *listStore) toTreeModel() *C.GtkTreeModel {
	if v == nil {
		return nil
	}
	return C.toGtkTreeModel(unsafe.Pointer(v.GObject))
}

// ListStoreNew is a wrapper around gtk_list_store_newv().
func ListStoreNew(types ...glib.Type) (*listStore, error) {
	gtypes := C.alloc_types(C.int(len(types)))
	for n, val := range types {
		C.set_type(gtypes, C.int(n), C.GType(val))
	}
	defer C.g_free(C.gpointer(gtypes))
	c := C.gtk_list_store_newv(C.gint(len(types)), gtypes)
	if c == nil {
		return nil, nilPtrErr
	}

	ls := wrapListStore(wrapObject(unsafe.Pointer(c)))
	return ls, nil
}

// Remove is a wrapper around gtk_list_store_remove().
func (v *listStore) Remove(iter gtk.TreeIter) bool {
	c := C.gtk_list_store_remove(v.native(), castToTreeIter(iter).native())
	return gobool(c)
}

// TODO(jrick)
/*
func (v *listStore) SetColumnTypes(types ...glib.Type) {
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
func (v *listStore) Set2(iter gtk.TreeIter, columns []int, values []interface{}) error {
	if len(columns) != len(values) {
		return errors.New("columns and values lengths do not match")
	}
	for i, val := range values {
		v.SetValue(castToTreeIter(iter), columns[i], val)
	}
	return nil
}

// SetValue is a wrapper around gtk_list_store_set_value().
func (v *listStore) SetValue(iter gtk.TreeIter, column int, value interface{}) error {
	switch value.(type) {
	case *gdk_impl.Pixbuf:
		pix := value.(*gdk_impl.Pixbuf)
		C._gtk_list_store_set(v.native(), castToTreeIter(iter).native(), C.gint(column), unsafe.Pointer(pix.Native()))

	default:
		gv, err := glib_impl.GValue(value)
		if err != nil {
			return err
		}

		C.gtk_list_store_set_value(v.native(), castToTreeIter(iter).native(),
			C.gint(column),
			(*C.GValue)(unsafe.Pointer(gv.Native())))
	}

	return nil
}

// func (v *listStore) Model(model ITreeModel) {
// 	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(model.toTreeModel()))}
//	v.TreeModel = *wrapTreeModel(obj)
//}

// SetSortColumnId() is a wrapper around gtk_tree_sortable_set_sort_column_id().
func (v *listStore) SetSortColumnId(column int, order gtk.SortType) {
	sort := C.toGtkTreeSortable(unsafe.Pointer(v.Native()))
	C.gtk_tree_sortable_set_sort_column_id(sort, C.gint(column), C.GtkSortType(order))
}

func (v *listStore) SetCols(iter gtk.TreeIter, cols gtk.Cols) error {
	for key, value := range cols {
		err := v.SetValue(iter, key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

// TODO(jrick)
/*
func (v *listStore) InsertWithValues(iter *TreeIter, position int, columns []int, values []glib_impl.Value) {
		var ccolumns *C.gint
		var cvalues *C.GValue

		C.gtk_list_store_insert_with_values(v.native(), iter.native(),
			C.gint(position), columns, values, C.gint(len(values)))
}
*/

// InsertBefore() is a wrapper around gtk_list_store_insert_before().
func (v *listStore) InsertBefore(sibling gtk.TreeIter) gtk.TreeIter {
	var ti C.GtkTreeIter
	C.gtk_list_store_insert_before(v.native(), &ti, castToTreeIter(sibling).native())
	iter := &treeIter{ti}
	return iter
}

// InsertAfter() is a wrapper around gtk_list_store_insert_after().
func (v *listStore) InsertAfter(sibling gtk.TreeIter) gtk.TreeIter {
	var ti C.GtkTreeIter
	C.gtk_list_store_insert_after(v.native(), &ti, castToTreeIter(sibling).native())
	iter := &treeIter{ti}
	return iter
}

// Prepend() is a wrapper around gtk_list_store_prepend().
func (v *listStore) Prepend() gtk.TreeIter {
	var ti C.GtkTreeIter
	C.gtk_list_store_prepend(v.native(), &ti)
	iter := &treeIter{ti}
	return iter
}

// Append() is a wrapper around gtk_list_store_append().
func (v *listStore) Append() gtk.TreeIter {
	var ti C.GtkTreeIter
	C.gtk_list_store_append(v.native(), &ti)
	iter := &treeIter{ti}
	return iter
}

// Clear() is a wrapper around gtk_list_store_clear().
func (v *listStore) Clear() {
	C.gtk_list_store_clear(v.native())
}

// IterIsValid() is a wrapper around gtk_list_store_iter_is_valid().
func (v *listStore) IterIsValid(iter gtk.TreeIter) bool {
	c := C.gtk_list_store_iter_is_valid(v.native(), castToTreeIter(iter).native())
	return gobool(c)
}

// TODO(jrick)
/*
func (v *listStore) Reorder(newOrder []int) {
}
*/

// Swap() is a wrapper around gtk_list_store_swap().
func (v *listStore) Swap(a, b gtk.TreeIter) {
	C.gtk_list_store_swap(v.native(), castToTreeIter(a).native(), castToTreeIter(b).native())
}

// MoveBefore() is a wrapper around gtk_list_store_move_before().
func (v *listStore) MoveBefore(iter, position gtk.TreeIter) {
	C.gtk_list_store_move_before(v.native(), castToTreeIter(position).native(),
		castToTreeIter(position).native())
}

// MoveAfter() is a wrapper around gtk_list_store_move_after().
func (v *listStore) MoveAfter(iter, position gtk.TreeIter) {
	C.gtk_list_store_move_after(v.native(), castToTreeIter(position).native(),
		castToTreeIter(position).native())
}

/*
 * GtkMenu
 */

// Menu is a representation of GTK's GtkMenu.
type menu struct {
	menuShell
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
func (v *menu) native() *C.GtkMenu {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenu(p)
}

func (v *menu) toMenu() *C.GtkMenu {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalMenu(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapMenu(obj), nil
}

func wrapMenu(obj *glib_impl.Object) *menu {
	return &menu{menuShell{container{widget{glib_impl.InitiallyUnowned{obj}}}}}
}

// MenuNew() is a wrapper around gtk_menu_new().
func MenuNew() (*menu, error) {
	c := C.gtk_menu_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapMenu(wrapObject(unsafe.Pointer(c))), nil
}

// PopupAtMouse() is a wrapper for gtk_menu_popup(), without the option for a custom positioning function.
func (v *menu) PopupAtMouseCursor(parentMenuShell gtk.Menu, parentMenuItem gtk.MenuItem, button int, activateTime uint32) {
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
func (v *menu) Popdown() {
	C.gtk_menu_popdown(v.native())
}

// ReorderChild() is a wrapper around gtk_menu_reorder_child().
func (v *menu) ReorderChild(child gtk.Widget, position int) {
	C.gtk_menu_reorder_child(v.native(), child.(IWidget).toWidget(), C.gint(position))
}

/*
 * GtkMenuBar
 */

// MenuBar is a representation of GTK's GtkMenuBar.
type menuBar struct {
	menuShell
}

// native() returns a pointer to the underlying GtkMenuBar.
func (v *menuBar) native() *C.GtkMenuBar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenuBar(p)
}

func marshalMenuBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapMenuBar(obj), nil
}

func wrapMenuBar(obj *glib_impl.Object) *menuBar {
	return &menuBar{menuShell{container{widget{glib_impl.InitiallyUnowned{obj}}}}}
}

// MenuBarNew() is a wrapper around gtk_menu_bar_new().
func MenuBarNew() (*menuBar, error) {
	c := C.gtk_menu_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapMenuBar(wrapObject(unsafe.Pointer(c))), nil
}

/*
 * GtkMenuButton
 */

// MenuButton is a representation of GTK's GtkMenuButton.
type menuButton struct {
	toggleButton
}

// native returns a pointer to the underlying GtkMenuButton.
func (v *menuButton) native() *C.GtkMenuButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenuButton(p)
}

func marshalMenuButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapMenuButton(obj), nil
}

func wrapMenuButton(obj *glib_impl.Object) *menuButton {
	return &menuButton{toggleButton{button{bin{container{widget{
		glib_impl.InitiallyUnowned{obj}}}}}}}
}

// MenuButtonNew is a wrapper around gtk_menu_button_new().
func MenuButtonNew() (*menuButton, error) {
	c := C.gtk_menu_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapMenuButton(wrapObject(unsafe.Pointer(c))), nil
}

// SetPopup is a wrapper around gtk_menu_button_set_popup().
func (v *menuButton) SetPopup(menu gtk.Menu) {
	C.gtk_menu_button_set_popup(v.native(), menu.(IMenu).toWidget())
}

// GetPopup is a wrapper around gtk_menu_button_get_popup().
func (v *menuButton) GetPopup() gtk.Menu {
	c := C.gtk_menu_button_get_popup(v.native())
	if c == nil {
		return nil
	}
	return wrapMenu(wrapObject(unsafe.Pointer(c)))
}

// TODO: gtk_menu_button_set_menu_model
// TODO: gtk_menu_button_get_menu_model

// SetDirection is a wrapper around gtk_menu_button_set_direction().
func (v *menuButton) SetDirection(direction gtk.ArrowType) {
	C.gtk_menu_button_set_direction(v.native(), C.GtkArrowType(direction))
}

// GetDirection is a wrapper around gtk_menu_button_get_direction().
func (v *menuButton) GetDirection() gtk.ArrowType {
	c := C.gtk_menu_button_get_direction(v.native())
	return gtk.ArrowType(c)
}

// SetAlignWidget is a wrapper around gtk_menu_button_set_align_widget().
func (v *menuButton) SetAlignWidget(alignWidget gtk.Widget) {
	C.gtk_menu_button_set_align_widget(v.native(), alignWidget.(IWidget).toWidget())
}

// GetAlignWidget is a wrapper around gtk_menu_button_get_align_widget().
func (v *menuButton) GetAlignWidget() gtk.Widget {
	c := C.gtk_menu_button_get_align_widget(v.native())
	if c == nil {
		return nil
	}
	return wrapWidget(wrapObject(unsafe.Pointer(c)))
}

/*
 * GtkMenuItem
 */

// MenuItem is a representation of GTK's GtkMenuItem.
type menuItem struct {
	bin
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
func (v *menuItem) native() *C.GtkMenuItem {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenuItem(p)
}

func (v *menuItem) toMenuItem() *C.GtkMenuItem {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalMenuItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapMenuItem(obj), nil
}

func wrapMenuItem(obj *glib_impl.Object) *menuItem {
	return &menuItem{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}
}

// MenuItemNew() is a wrapper around gtk_menu_item_new().
func MenuItemNew() (*menuItem, error) {
	c := C.gtk_menu_item_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapMenuItem(wrapObject(unsafe.Pointer(c))), nil
}

// MenuItemNewWithLabel() is a wrapper around gtk_menu_item_new_with_label().
func MenuItemNewWithLabel(label string) (*menuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_menu_item_new_with_label((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapMenuItem(wrapObject(unsafe.Pointer(c))), nil
}

// MenuItemNewWithMnemonic() is a wrapper around
// gtk_menu_item_new_with_mnemonic().
func MenuItemNewWithMnemonic(label string) (*menuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_menu_item_new_with_mnemonic((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapMenuItem(wrapObject(unsafe.Pointer(c))), nil
}

// SetSubmenu() is a wrapper around gtk_menu_item_set_submenu().
func (v *menuItem) SetSubmenu(submenu gtk.Widget) {
	C.gtk_menu_item_set_submenu(v.native(), submenu.(IWidget).toWidget())
}

// Sets text on the menu_item label
func (v *menuItem) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_menu_item_set_label(v.native(), (*C.gchar)(cstr))
}

/*
 * GtkMessageDialog
 */

// MessageDialog is a representation of GTK's GtkMessageDialog.
type messageDialog struct {
	dialog
}

// native returns a pointer to the underlying GtkMessageDialog.
func (v *messageDialog) native() *C.GtkMessageDialog {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMessageDialog(p)
}

func marshalMessageDialog(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapMessageDialog(obj), nil
}

func wrapMessageDialog(obj *glib_impl.Object) *messageDialog {
	return &messageDialog{dialog{window{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}}}
}

// MessageDialogNew() is a wrapper around gtk_message_dialog_new().
// The text is created and formatted by the format specifier and any
// additional arguments.
func MessageDialogNew(parent IWindow, flags gtk.DialogFlags, mType gtk.MessageType, buttons gtk.ButtonsType, format string, a ...interface{}) *messageDialog {
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
	return wrapMessageDialog(wrapObject(unsafe.Pointer(c)))
}

// MessageDialogNewWithMarkup is a wrapper around
// gtk_message_dialog_new_with_markup().
func MessageDialogNewWithMarkup(parent IWindow, flags gtk.DialogFlags, mType gtk.MessageType, buttons gtk.ButtonsType, format string, a ...interface{}) *messageDialog {
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
	return wrapMessageDialog(wrapObject(unsafe.Pointer(c)))
}

// SetMarkup is a wrapper around gtk_message_dialog_set_markup().
func (v *messageDialog) SetMarkup(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_message_dialog_set_markup(v.native(), (*C.gchar)(cstr))
}

// FormatSecondaryText is a wrapper around
// gtk_message_dialog_format_secondary_text().
func (v *messageDialog) FormatSecondaryText(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C._gtk_message_dialog_format_secondary_text(v.native(),
		(*C.gchar)(cstr))
}

// FormatSecondaryMarkup is a wrapper around
// gtk_message_dialog_format_secondary_text().
func (v *messageDialog) FormatSecondaryMarkup(format string, a ...interface{}) {
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
type notebook struct {
	container
}

// native returns a pointer to the underlying GtkNotebook.
func (v *notebook) native() *C.GtkNotebook {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkNotebook(p)
}

func marshalNotebook(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapNotebook(obj), nil
}

func wrapNotebook(obj *glib_impl.Object) *notebook {
	return &notebook{container{widget{glib_impl.InitiallyUnowned{obj}}}}
}

// NotebookNew() is a wrapper around gtk_notebook_new().
func NotebookNew() (*notebook, error) {
	c := C.gtk_notebook_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapNotebook(wrapObject(unsafe.Pointer(c))), nil
}

// AppendPage() is a wrapper around gtk_notebook_append_page().
func (v *notebook) AppendPage(child gtk.Widget, tabLabel gtk.Widget) int {
	cTabLabel := nullableWidget(tabLabel)
	c := C.gtk_notebook_append_page(v.native(), child.(IWidget).toWidget(), cTabLabel)
	return int(c)
}

// AppendPageMenu() is a wrapper around gtk_notebook_append_page_menu().
func (v *notebook) AppendPageMenu(child gtk.Widget, tabLabel gtk.Widget, menuLabel gtk.Widget) int {
	c := C.gtk_notebook_append_page_menu(v.native(), child.(IWidget).toWidget(),
		tabLabel.(IWidget).toWidget(), menuLabel.(IWidget).toWidget())
	return int(c)
}

// PrependPage() is a wrapper around gtk_notebook_prepend_page().
func (v *notebook) PrependPage(child gtk.Widget, tabLabel gtk.Widget) int {
	cTabLabel := nullableWidget(tabLabel)
	c := C.gtk_notebook_prepend_page(v.native(), child.(IWidget).toWidget(), cTabLabel)
	return int(c)
}

// PrependPageMenu() is a wrapper around gtk_notebook_prepend_page_menu().
func (v *notebook) PrependPageMenu(child gtk.Widget, tabLabel gtk.Widget, menuLabel gtk.Widget) int {
	c := C.gtk_notebook_prepend_page_menu(v.native(), child.(IWidget).toWidget(),
		tabLabel.(IWidget).toWidget(), menuLabel.(IWidget).toWidget())
	return int(c)
}

// InsertPage() is a wrapper around gtk_notebook_insert_page().
func (v *notebook) InsertPage(child gtk.Widget, tabLabel gtk.Widget, position int) int {
	label := nullableWidget(tabLabel)
	c := C.gtk_notebook_insert_page(v.native(), child.(IWidget).toWidget(), label, C.gint(position))

	return int(c)
}

// InsertPageMenu() is a wrapper around gtk_notebook_insert_page_menu().
func (v *notebook) InsertPageMenu(child gtk.Widget, tabLabel gtk.Widget, menuLabel gtk.Widget, position int) int {
	c := C.gtk_notebook_insert_page_menu(v.native(), child.(IWidget).toWidget(),
		tabLabel.(IWidget).toWidget(), menuLabel.(IWidget).toWidget(), C.gint(position))
	return int(c)
}

// RemovePage() is a wrapper around gtk_notebook_remove_page().
func (v *notebook) RemovePage(pageNum int) {
	C.gtk_notebook_remove_page(v.native(), C.gint(pageNum))
}

// PageNum() is a wrapper around gtk_notebook_page_num().
func (v *notebook) PageNum(child gtk.Widget) int {
	c := C.gtk_notebook_page_num(v.native(), child.(IWidget).toWidget())
	return int(c)
}

// NextPage() is a wrapper around gtk_notebook_next_page().
func (v *notebook) NextPage() {
	C.gtk_notebook_next_page(v.native())
}

// PrevPage() is a wrapper around gtk_notebook_prev_page().
func (v *notebook) PrevPage() {
	C.gtk_notebook_prev_page(v.native())
}

// ReorderChild() is a wrapper around gtk_notebook_reorder_child().
func (v *notebook) ReorderChild(child gtk.Widget, position int) {
	C.gtk_notebook_reorder_child(v.native(), child.(IWidget).toWidget(),
		C.gint(position))
}

// SetTabPos() is a wrapper around gtk_notebook_set_tab_pos().
func (v *notebook) SetTabPos(pos gtk.PositionType) {
	C.gtk_notebook_set_tab_pos(v.native(), C.GtkPositionType(pos))
}

// SetShowTabs() is a wrapper around gtk_notebook_set_show_tabs().
func (v *notebook) SetShowTabs(showTabs bool) {
	C.gtk_notebook_set_show_tabs(v.native(), gbool(showTabs))
}

// SetShowBorder() is a wrapper around gtk_notebook_set_show_border().
func (v *notebook) SetShowBorder(showBorder bool) {
	C.gtk_notebook_set_show_border(v.native(), gbool(showBorder))
}

// SetScrollable() is a wrapper around gtk_notebook_set_scrollable().
func (v *notebook) SetScrollable(scrollable bool) {
	C.gtk_notebook_set_scrollable(v.native(), gbool(scrollable))
}

// PopupEnable() is a wrapper around gtk_notebook_popup_enable().
func (v *notebook) PopupEnable() {
	C.gtk_notebook_popup_enable(v.native())
}

// PopupDisable() is a wrapper around gtk_notebook_popup_disable().
func (v *notebook) PopupDisable() {
	C.gtk_notebook_popup_disable(v.native())
}

// GetCurrentPage() is a wrapper around gtk_notebook_get_current_page().
func (v *notebook) GetCurrentPage() int {
	c := C.gtk_notebook_get_current_page(v.native())
	return int(c)
}

// GetMenuLabel() is a wrapper around gtk_notebook_get_menu_label().
func (v *notebook) GetMenuLabel(child gtk.Widget) (gtk.Widget, error) {
	c := C.gtk_notebook_get_menu_label(v.native(), child.(IWidget).toWidget())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWidget(wrapObject(unsafe.Pointer(c))), nil
}

// GetNthPage() is a wrapper around gtk_notebook_get_nth_page().
func (v *notebook) GetNthPage(pageNum int) (gtk.Widget, error) {
	c := C.gtk_notebook_get_nth_page(v.native(), C.gint(pageNum))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWidget(wrapObject(unsafe.Pointer(c))), nil
}

// GetNPages() is a wrapper around gtk_notebook_get_n_pages().
func (v *notebook) GetNPages() int {
	c := C.gtk_notebook_get_n_pages(v.native())
	return int(c)
}

// GetTabLabel() is a wrapper around gtk_notebook_get_tab_label().
func (v *notebook) GetTabLabel(child gtk.Widget) (gtk.Widget, error) {
	c := C.gtk_notebook_get_tab_label(v.native(), child.(IWidget).toWidget())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWidget(wrapObject(unsafe.Pointer(c))), nil
}

// SetMenuLabel() is a wrapper around gtk_notebook_set_menu_label().
func (v *notebook) SetMenuLabel(child, menuLabel gtk.Widget) {
	C.gtk_notebook_set_menu_label(v.native(), child.(IWidget).toWidget(),
		menuLabel.(IWidget).toWidget())
}

// SetMenuLabelText() is a wrapper around gtk_notebook_set_menu_label_text().
func (v *notebook) SetMenuLabelText(child gtk.Widget, menuText string) {
	cstr := C.CString(menuText)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_notebook_set_menu_label_text(v.native(), child.(IWidget).toWidget(),
		(*C.gchar)(cstr))
}

// SetTabLabel() is a wrapper around gtk_notebook_set_tab_label().
func (v *notebook) SetTabLabel(child, tabLabel gtk.Widget) {
	C.gtk_notebook_set_tab_label(v.native(), child.(IWidget).toWidget(),
		tabLabel.(IWidget).toWidget())
}

// SetTabLabelText() is a wrapper around gtk_notebook_set_tab_label_text().
func (v *notebook) SetTabLabelText(child gtk.Widget, tabText string) {
	cstr := C.CString(tabText)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_notebook_set_tab_label_text(v.native(), child.(IWidget).toWidget(),
		(*C.gchar)(cstr))
}

// SetTabReorderable() is a wrapper around gtk_notebook_set_tab_reorderable().
func (v *notebook) SetTabReorderable(child gtk.Widget, reorderable bool) {
	C.gtk_notebook_set_tab_reorderable(v.native(), child.(IWidget).toWidget(),
		gbool(reorderable))
}

// SetTabDetachable() is a wrapper around gtk_notebook_set_tab_detachable().
func (v *notebook) SetTabDetachable(child gtk.Widget, detachable bool) {
	C.gtk_notebook_set_tab_detachable(v.native(), child.(IWidget).toWidget(),
		gbool(detachable))
}

// GetMenuLabelText() is a wrapper around gtk_notebook_get_menu_label_text().
func (v *notebook) GetMenuLabelText(child gtk.Widget) (string, error) {
	c := C.gtk_notebook_get_menu_label_text(v.native(), child.(IWidget).toWidget())
	if c == nil {
		return "", errors.New("No menu label for widget")
	}
	return C.GoString((*C.char)(c)), nil
}

// GetScrollable() is a wrapper around gtk_notebook_get_scrollable().
func (v *notebook) GetScrollable() bool {
	c := C.gtk_notebook_get_scrollable(v.native())
	return gobool(c)
}

// GetShowBorder() is a wrapper around gtk_notebook_get_show_border().
func (v *notebook) GetShowBorder() bool {
	c := C.gtk_notebook_get_show_border(v.native())
	return gobool(c)
}

// GetShowTabs() is a wrapper around gtk_notebook_get_show_tabs().
func (v *notebook) GetShowTabs() bool {
	c := C.gtk_notebook_get_show_tabs(v.native())
	return gobool(c)
}

// GetTabLabelText() is a wrapper around gtk_notebook_get_tab_label_text().
func (v *notebook) GetTabLabelText(child gtk.Widget) (string, error) {
	c := C.gtk_notebook_get_tab_label_text(v.native(), child.(IWidget).toWidget())
	if c == nil {
		return "", errors.New("No tab label for widget")
	}
	return C.GoString((*C.char)(c)), nil
}

// GetTabPos() is a wrapper around gtk_notebook_get_tab_pos().
func (v *notebook) GetTabPos() gtk.PositionType {
	c := C.gtk_notebook_get_tab_pos(v.native())
	return gtk.PositionType(c)
}

// GetTabReorderable() is a wrapper around gtk_notebook_get_tab_reorderable().
func (v *notebook) GetTabReorderable(child gtk.Widget) bool {
	c := C.gtk_notebook_get_tab_reorderable(v.native(), child.(IWidget).toWidget())
	return gobool(c)
}

// GetTabDetachable() is a wrapper around gtk_notebook_get_tab_detachable().
func (v *notebook) GetTabDetachable(child gtk.Widget) bool {
	c := C.gtk_notebook_get_tab_detachable(v.native(), child.(IWidget).toWidget())
	return gobool(c)
}

// SetCurrentPage() is a wrapper around gtk_notebook_set_current_page().
func (v *notebook) SetCurrentPage(pageNum int) {
	C.gtk_notebook_set_current_page(v.native(), C.gint(pageNum))
}

// SetGroupName() is a wrapper around gtk_notebook_set_group_name().
func (v *notebook) SetGroupName(groupName string) {
	cstr := C.CString(groupName)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_notebook_set_group_name(v.native(), (*C.gchar)(cstr))
}

// GetGroupName() is a wrapper around gtk_notebook_get_group_name().
func (v *notebook) GetGroupName() (string, error) {
	c := C.gtk_notebook_get_group_name(v.native())
	if c == nil {
		return "", errors.New("No group name")
	}
	return C.GoString((*C.char)(c)), nil
}

// SetActionWidget() is a wrapper around gtk_notebook_set_action_widget().
func (v *notebook) SetActionWidget(widget gtk.Widget, packType gtk.PackType) {
	C.gtk_notebook_set_action_widget(v.native(), widget.(IWidget).toWidget(),
		C.GtkPackType(packType))
}

// GetActionWidget() is a wrapper around gtk_notebook_get_action_widget().
func (v *notebook) GetActionWidget(packType gtk.PackType) (gtk.Widget, error) {
	c := C.gtk_notebook_get_action_widget(v.native(),
		C.GtkPackType(packType))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWidget(wrapObject(unsafe.Pointer(c))), nil
}

/*
 * GtkOffscreenWindow
 */

// OffscreenWindow is a representation of GTK's GtkOffscreenWindow.
type offscreenWindow struct {
	window
}

// native returns a pointer to the underlying GtkOffscreenWindow.
func (v *offscreenWindow) native() *C.GtkOffscreenWindow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkOffscreenWindow(p)
}

func marshalOffscreenWindow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapOffscreenWindow(obj), nil
}

func wrapOffscreenWindow(obj *glib_impl.Object) *offscreenWindow {
	return &offscreenWindow{window{bin{container{widget{
		glib_impl.InitiallyUnowned{obj}}}}}}
}

// OffscreenWindowNew is a wrapper around gtk_offscreen_window_new().
func OffscreenWindowNew() (*offscreenWindow, error) {
	c := C.gtk_offscreen_window_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapOffscreenWindow(wrapObject(unsafe.Pointer(c))), nil
}

// GetSurface is a wrapper around gtk_offscreen_window_get_surface().
// The returned surface is safe to use over window resizes.
func (v *offscreenWindow) GetSurface() (cairo.Surface, error) {
	c := C.gtk_offscreen_window_get_surface(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	cairoPtr := (uintptr)(unsafe.Pointer(c))
	s := cairo_impl.NewSurface(cairoPtr, true)
	return s, nil
}

// GetPixbuf is a wrapper around gtk_offscreen_window_get_pixbuf().
func (v *offscreenWindow) GetPixbuf() (gdk.Pixbuf, error) {
	c := C.gtk_offscreen_window_get_pixbuf(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	// Pixbuf is returned with ref count of 1, so don't increment.
	// Is it a floating reference?
	pb := &gdk_impl.Pixbuf{wrapObject(unsafe.Pointer(c))}
	return pb, nil
}

/*
 * GtkOrientable
 */

// Orientable is a representation of GTK's GtkOrientable GInterface.
type orientable struct {
	*glib_impl.Object
}

// IOrientable is an interface type implemented by all structs
// embedding an Orientable.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkOrientable.
type IOrientable interface {
	toOrientable() *C.GtkOrientable
}

// native returns a pointer to the underlying GObject as a GtkOrientable.
func (v *orientable) native() *C.GtkOrientable {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkOrientable(p)
}

func marshalOrientable(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapOrientable(obj), nil
}

func wrapOrientable(obj *glib_impl.Object) *orientable {
	return &orientable{obj}
}

// GetOrientation() is a wrapper around gtk_orientable_get_orientation().
func (v *orientable) GetOrientation() gtk.Orientation {
	c := C.gtk_orientable_get_orientation(v.native())
	return gtk.Orientation(c)
}

// SetOrientation() is a wrapper around gtk_orientable_set_orientation().
func (v *orientable) SetOrientation(orientation gtk.Orientation) {
	C.gtk_orientable_set_orientation(v.native(),
		C.GtkOrientation(orientation))
}

/*
 * GtkPaned
 */

// Paned is a representation of GTK's GtkPaned.
type paned struct {
	bin
}

// native returns a pointer to the underlying GtkPaned.
func (v *paned) native() *C.GtkPaned {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkPaned(p)
}

func marshalPaned(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapPaned(obj), nil
}

func wrapPaned(obj *glib_impl.Object) *paned {
	return &paned{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}
}

// PanedNew() is a wrapper around gtk_scrolled_window_new().
func PanedNew(orientation gtk.Orientation) (*paned, error) {
	c := C.gtk_paned_new(C.GtkOrientation(orientation))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapPaned(wrapObject(unsafe.Pointer(c))), nil
}

// Add1() is a wrapper around gtk_paned_add1().
func (v *paned) Add1(child gtk.Widget) {
	C.gtk_paned_add1(v.native(), child.(IWidget).toWidget())
}

// Add2() is a wrapper around gtk_paned_add2().
func (v *paned) Add2(child gtk.Widget) {
	C.gtk_paned_add2(v.native(), child.(IWidget).toWidget())
}

// Pack1() is a wrapper around gtk_paned_pack1().
func (v *paned) Pack1(child gtk.Widget, resize, shrink bool) {
	C.gtk_paned_pack1(v.native(), child.(IWidget).toWidget(), gbool(resize), gbool(shrink))
}

// Pack2() is a wrapper around gtk_paned_pack2().
func (v *paned) Pack2(child gtk.Widget, resize, shrink bool) {
	C.gtk_paned_pack2(v.native(), child.(IWidget).toWidget(), gbool(resize), gbool(shrink))
}

// SetPosition() is a wrapper around gtk_paned_set_position().
func (v *paned) SetPosition(position int) {
	C.gtk_paned_set_position(v.native(), C.gint(position))
}

// GetChild1() is a wrapper around gtk_paned_get_child1().
func (v *paned) GetChild1() (gtk.Widget, error) {
	c := C.gtk_paned_get_child1(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWidget(wrapObject(unsafe.Pointer(c))), nil
}

// GetChild2() is a wrapper around gtk_paned_get_child2().
func (v *paned) GetChild2() (gtk.Widget, error) {
	c := C.gtk_paned_get_child2(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWidget(wrapObject(unsafe.Pointer(c))), nil
}

// GetHandleWindow() is a wrapper around gtk_paned_get_handle_window().
func (v *paned) GetHandleWindow() (gtk.Window, error) {
	c := C.gtk_paned_get_handle_window(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWindow(wrapObject(unsafe.Pointer(c))), nil
}

// GetPosition() is a wrapper around gtk_paned_get_position().
func (v *paned) GetPosition() int {
	return int(C.gtk_paned_get_position(v.native()))
}

/*
 * GtkProgressBar
 */

// ProgressBar is a representation of GTK's GtkProgressBar.
type progressBar struct {
	widget
}

// native returns a pointer to the underlying GtkProgressBar.
func (v *progressBar) native() *C.GtkProgressBar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkProgressBar(p)
}

func marshalProgressBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapProgressBar(obj), nil
}

func wrapProgressBar(obj *glib_impl.Object) *progressBar {
	return &progressBar{widget{glib_impl.InitiallyUnowned{obj}}}
}

// ProgressBarNew() is a wrapper around gtk_progress_bar_new().
func ProgressBarNew() (*progressBar, error) {
	c := C.gtk_progress_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapProgressBar(wrapObject(unsafe.Pointer(c))), nil
}

// SetFraction() is a wrapper around gtk_progress_bar_set_fraction().
func (v *progressBar) SetFraction(fraction float64) {
	C.gtk_progress_bar_set_fraction(v.native(), C.gdouble(fraction))
}

// GetFraction() is a wrapper around gtk_progress_bar_get_fraction().
func (v *progressBar) GetFraction() float64 {
	c := C.gtk_progress_bar_get_fraction(v.native())
	return float64(c)
}

// SetShowText is a wrapper around gtk_progress_bar_set_show_text().
func (v *progressBar) SetShowText(showText bool) {
	C.gtk_progress_bar_set_show_text(v.native(), gbool(showText))
}

// GetShowText is a wrapper around gtk_progress_bar_get_show_text().
func (v *progressBar) GetShowText() bool {
	c := C.gtk_progress_bar_get_show_text(v.native())
	return gobool(c)
}

// SetText() is a wrapper around gtk_progress_bar_set_text().
func (v *progressBar) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_progress_bar_set_text(v.native(), (*C.gchar)(cstr))
}

/*
 * GtkRadioButton
 */

// RadioButton is a representation of GTK's GtkRadioButton.
type radioButton struct {
	checkButton
}

// native returns a pointer to the underlying GtkRadioButton.
func (v *radioButton) native() *C.GtkRadioButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkRadioButton(p)
}

func marshalRadioButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapRadioButton(obj), nil
}

func wrapRadioButton(obj *glib_impl.Object) *radioButton {
	return &radioButton{checkButton{toggleButton{button{bin{container{
		widget{glib_impl.InitiallyUnowned{obj}}}}}}}}
}

// RadioButtonNew is a wrapper around gtk_radio_button_new().
func RadioButtonNew(group *glib_impl.SList) (*radioButton, error) {
	gslist := (*C.GSList)(unsafe.Pointer(group.Native()))
	c := C.gtk_radio_button_new(gslist)
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapRadioButton(wrapObject(unsafe.Pointer(c))), nil
}

// RadioButtonNewFromWidget is a wrapper around
// gtk_radio_button_new_from_widget().
func RadioButtonNewFromWidget(radioGroupMember *radioButton) (*radioButton, error) {
	c := C.gtk_radio_button_new_from_widget(radioGroupMember.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapRadioButton(wrapObject(unsafe.Pointer(c))), nil
}

// RadioButtonNewWithLabel is a wrapper around
// gtk_radio_button_new_with_label().
func RadioButtonNewWithLabel(group *glib_impl.SList, label string) (*radioButton, error) {
	gslist := (*C.GSList)(unsafe.Pointer(group.Native()))
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_button_new_with_label(gslist, (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapRadioButton(wrapObject(unsafe.Pointer(c))), nil
}

// RadioButtonNewWithLabelFromWidget is a wrapper around
// gtk_radio_button_new_with_label_from_widget().
func RadioButtonNewWithLabelFromWidget(radioGroupMember *radioButton, label string) (*radioButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_button_new_with_label_from_widget(radioGroupMember.native(),
		(*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapRadioButton(wrapObject(unsafe.Pointer(c))), nil
}

// RadioButtonNewWithMnemonic is a wrapper around
// gtk_radio_button_new_with_mnemonic()
func RadioButtonNewWithMnemonic(group *glib_impl.SList, label string) (*radioButton, error) {
	gslist := (*C.GSList)(unsafe.Pointer(group.Native()))
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_button_new_with_mnemonic(gslist, (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapRadioButton(wrapObject(unsafe.Pointer(c))), nil
}

// RadioButtonNewWithMnemonicFromWidget is a wrapper around
// gtk_radio_button_new_with_mnemonic_from_widget().
func RadioButtonNewWithMnemonicFromWidget(radioGroupMember *radioButton, label string) (*radioButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_button_new_with_mnemonic_from_widget(radioGroupMember.native(),
		(*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapRadioButton(wrapObject(unsafe.Pointer(c))), nil
}

// SetGroup is a wrapper around gtk_radio_button_set_group().
func (v *radioButton) SetGroup(group glib.SList) {
	gslist := (*C.GSList)(unsafe.Pointer(glib_impl.CastToSList(group).Native()))
	C.gtk_radio_button_set_group(v.native(), gslist)
}

// GetGroup is a wrapper around gtk_radio_button_get_group().
func (v *radioButton) GetGroup() (glib.SList, error) {
	c := C.gtk_radio_button_get_group(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return glib_impl.WrapSList(uintptr(unsafe.Pointer(c))), nil
}

// JoinGroup is a wrapper around gtk_radio_button_join_group().
func (v *radioButton) JoinGroup(groupSource gtk.RadioButton) {
	C.gtk_radio_button_join_group(v.native(), castToRadioButton(groupSource).native())
}

/*
 * GtkRadioMenuItem
 */

// RadioMenuItem is a representation of GTK's GtkRadioMenuItem.
type radioMenuItem struct {
	checkMenuItem
}

// native returns a pointer to the underlying GtkRadioMenuItem.
func (v *radioMenuItem) native() *C.GtkRadioMenuItem {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkRadioMenuItem(p)
}

func marshalRadioMenuItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapRadioMenuItem(obj), nil
}

func wrapRadioMenuItem(obj *glib_impl.Object) *radioMenuItem {
	return &radioMenuItem{checkMenuItem{menuItem{bin{container{
		widget{glib_impl.InitiallyUnowned{obj}}}}}}}
}

// RadioMenuItemNew is a wrapper around gtk_radio_menu_item_new().
func RadioMenuItemNew(group *glib_impl.SList) (*radioMenuItem, error) {
	gslist := (*C.GSList)(unsafe.Pointer(group.Native()))
	c := C.gtk_radio_menu_item_new(gslist)
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapRadioMenuItem(wrapObject(unsafe.Pointer(c))), nil
}

// RadioMenuItemNewWithLabel is a wrapper around
// gtk_radio_menu_item_new_with_label().
func RadioMenuItemNewWithLabel(group *glib_impl.SList, label string) (*radioMenuItem, error) {
	gslist := (*C.GSList)(unsafe.Pointer(group.Native()))
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_menu_item_new_with_label(gslist, (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapRadioMenuItem(wrapObject(unsafe.Pointer(c))), nil
}

// RadioMenuItemNewWithMnemonic is a wrapper around
// gtk_radio_menu_item_new_with_mnemonic().
func RadioMenuItemNewWithMnemonic(group *glib_impl.SList, label string) (*radioMenuItem, error) {
	gslist := (*C.GSList)(unsafe.Pointer(group.Native()))
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_menu_item_new_with_mnemonic(gslist, (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapRadioMenuItem(wrapObject(unsafe.Pointer(c))), nil
}

// RadioMenuItemNewFromWidget is a wrapper around
// gtk_radio_menu_item_new_from_widget().
func RadioMenuItemNewFromWidget(group *radioMenuItem) (*radioMenuItem, error) {
	c := C.gtk_radio_menu_item_new_from_widget(group.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapRadioMenuItem(wrapObject(unsafe.Pointer(c))), nil
}

// RadioMenuItemNewWithLabelFromWidget is a wrapper around
// gtk_radio_menu_item_new_with_label_from_widget().
func RadioMenuItemNewWithLabelFromWidget(group *radioMenuItem, label string) (*radioMenuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_menu_item_new_with_label_from_widget(group.native(),
		(*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapRadioMenuItem(wrapObject(unsafe.Pointer(c))), nil
}

// RadioMenuItemNewWithMnemonicFromWidget is a wrapper around
// gtk_radio_menu_item_new_with_mnemonic_from_widget().
func RadioMenuItemNewWithMnemonicFromWidget(group *radioMenuItem, label string) (*radioMenuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_menu_item_new_with_mnemonic_from_widget(group.native(),
		(*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapRadioMenuItem(wrapObject(unsafe.Pointer(c))), nil
}

// SetGroup is a wrapper around gtk_radio_menu_item_set_group().
func (v *radioMenuItem) SetGroup(group glib.SList) {
	gslist := (*C.GSList)(unsafe.Pointer(glib_impl.CastToSList(group).Native()))
	C.gtk_radio_menu_item_set_group(v.native(), gslist)
}

// GetGroup is a wrapper around gtk_radio_menu_item_get_group().
func (v *radioMenuItem) GetGroup() (glib.SList, error) {
	c := C.gtk_radio_menu_item_get_group(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return glib_impl.WrapSList(uintptr(unsafe.Pointer(c))), nil
}

/*
 * GtkRange
 */

// Range is a representation of GTK's GtkRange.
type _range struct {
	widget
}

// native returns a pointer to the underlying GtkRange.
func (v *_range) native() *C.GtkRange {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkRange(p)
}

func marshalRange(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapRange(obj), nil
}

func wrapRange(obj *glib_impl.Object) *_range {
	return &_range{widget{glib_impl.InitiallyUnowned{obj}}}
}

// GetValue is a wrapper around gtk_range_get_value().
func (v *_range) GetValue() float64 {
	c := C.gtk_range_get_value(v.native())
	return float64(c)
}

// SetValue is a wrapper around gtk_range_set_value().
func (v *_range) SetValue(value float64) {
	C.gtk_range_set_value(v.native(), C.gdouble(value))
}

// SetIncrements() is a wrapper around gtk_range_set_increments().
func (v *_range) SetIncrements(step, page float64) {
	C.gtk_range_set_increments(v.native(), C.gdouble(step), C.gdouble(page))
}

// SetRange() is a wrapper around gtk_range_set_range().
func (v *_range) SetRange(min, max float64) {
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
type recentChooser struct {
	*glib_impl.Object
}

// native returns a pointer to the underlying GtkRecentChooser.
func (v *recentChooser) native() *C.GtkRecentChooser {
	if v == nil || v.Object == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkRecentChooser(p)
}

func wrapRecentChooser(obj *glib_impl.Object) *recentChooser {
	return &recentChooser{obj}
}

func (v *recentChooser) toRecentChooser() *C.GtkRecentChooser {
	return v.native()
}

func (v *recentChooser) GetCurrentUri() string {
	curi := C.gtk_recent_chooser_get_current_uri(v.native())
	uri := C.GoString((*C.char)(curi))
	return uri
}

func (v *recentChooser) AddFilter(filter gtk.RecentFilter) {
	C.gtk_recent_chooser_add_filter(v.native(), castToRecentFilter(filter).native())
}

func (v *recentChooser) RemoveFilter(filter gtk.RecentFilter) {
	C.gtk_recent_chooser_remove_filter(v.native(), castToRecentFilter(filter).native())
}

/*
 * GtkRecentChooserMenu
 */

// RecentChooserMenu is a representation of GTK's GtkRecentChooserMenu.
type recentChooserMenu struct {
	menu
	recentChooser
}

// native returns a pointer to the underlying GtkRecentManager.
func (v *recentChooserMenu) native() *C.GtkRecentChooserMenu {
	if v == nil || v.Object == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkRecentChooserMenu(p)
}

func wrapRecentChooserMenu(obj *glib_impl.Object) *recentChooserMenu {
	return &recentChooserMenu{
		menu{menuShell{container{widget{glib_impl.InitiallyUnowned{obj}}}}},
		recentChooser{obj},
	}
}

/*
 * GtkRecentFilter
 */

// RecentFilter is a representation of GTK's GtkRecentFilter.
type recentFilter struct {
	glib_impl.InitiallyUnowned
}

// native returns a pointer to the underlying GtkRecentFilter.
func (v *recentFilter) native() *C.GtkRecentFilter {
	if v == nil || v.Object == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkRecentFilter(p)
}

func wrapRecentFilter(obj *glib_impl.Object) *recentFilter {
	return &recentFilter{glib_impl.InitiallyUnowned{obj}}
}

// RecentFilterNew is a wrapper around gtk_recent_filter_new().
func RecentFilterNew() (*recentFilter, error) {
	c := C.gtk_recent_filter_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapRecentFilter(wrapObject(unsafe.Pointer(c))), nil
}

/*
 * GtkRecentManager
 */

// RecentManager is a representation of GTK's GtkRecentManager.
type recentManager struct {
	*glib_impl.Object
}

// native returns a pointer to the underlying GtkRecentManager.
func (v *recentManager) native() *C.GtkRecentManager {
	if v == nil || v.Object == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkRecentManager(p)
}

func marshalRecentManager(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapRecentManager(obj), nil
}

func wrapRecentManager(obj *glib_impl.Object) *recentManager {
	return &recentManager{obj}
}

// RecentManagerGetDefault is a wrapper around gtk_recent_manager_get_default().
func RecentManagerGetDefault() (*recentManager, error) {
	c := C.gtk_recent_manager_get_default()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	v := wrapRecentManager(obj)
	return v, nil
}

// AddItem is a wrapper around gtk_recent_manager_add_item().
func (v *recentManager) AddItem(fileURI string) bool {
	cstr := C.CString(fileURI)
	defer C.free(unsafe.Pointer(cstr))
	cok := C.gtk_recent_manager_add_item(v.native(), (*C.gchar)(cstr))
	return gobool(cok)
}

/*
 * GtkScale
 */

// Scale is a representation of GTK's GtkScale.
type scale struct {
	_range
}

// native returns a pointer to the underlying GtkScale.
func (v *scale) native() *C.GtkScale {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkScale(p)
}

func marshalScale(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapScale(obj), nil
}

func wrapScale(obj *glib_impl.Object) *scale {
	return &scale{_range{widget{glib_impl.InitiallyUnowned{obj}}}}
}

// ScaleNew is a wrapper around gtk_scale_new().
func ScaleNew(orientation gtk.Orientation, adjustment *adjustment) (*scale, error) {
	c := C.gtk_scale_new(C.GtkOrientation(orientation), adjustment.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapScale(wrapObject(unsafe.Pointer(c))), nil
}

// ScaleNewWithRange is a wrapper around gtk_scale_new_with_range().
func ScaleNewWithRange(orientation gtk.Orientation, min, max, step float64) (*scale, error) {
	c := C.gtk_scale_new_with_range(C.GtkOrientation(orientation),
		C.gdouble(min), C.gdouble(max), C.gdouble(step))

	if c == nil {
		return nil, nilPtrErr
	}
	return wrapScale(wrapObject(unsafe.Pointer(c))), nil
}

/*
 * GtkScaleButton
 */

// ScaleButton is a representation of GTK's GtkScaleButton.
type scaleButton struct {
	button
}

// native() returns a pointer to the underlying GtkScaleButton.
func (v *scaleButton) native() *C.GtkScaleButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkScaleButton(p)
}

func marshalScaleButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapScaleButton(obj), nil
}

func wrapScaleButton(obj *glib_impl.Object) *scaleButton {
	return &scaleButton{button{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}}
}

// ScaleButtonNew() is a wrapper around gtk_scale_button_new().
func ScaleButtonNew(size gtk.IconSize, min, max, step float64, icons []string) (*scaleButton, error) {
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
	return wrapScaleButton(wrapObject(unsafe.Pointer(c))), nil
}

// GetAdjustment() is a wrapper around gtk_scale_button_get_adjustment().
func (v *scaleButton) GetAdjustment() gtk.Adjustment {
	c := C.gtk_scale_button_get_adjustment(v.native())
	obj := wrapObject(unsafe.Pointer(c))
	return &adjustment{glib_impl.InitiallyUnowned{obj}}
}

// GetPopup() is a wrapper around gtk_scale_button_get_popup().
func (v *scaleButton) GetPopup() (gtk.Widget, error) {
	c := C.gtk_scale_button_get_popup(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWidget(wrapObject(unsafe.Pointer(c))), nil
}

// GetValue() is a wrapper around gtk_scale_button_get_value().
func (v *scaleButton) GetValue() float64 {
	return float64(C.gtk_scale_button_get_value(v.native()))
}

// SetAdjustment() is a wrapper around gtk_scale_button_set_adjustment().
func (v *scaleButton) SetAdjustment(adjustment gtk.Adjustment) {
	C.gtk_scale_button_set_adjustment(v.native(), castToAdjustment(adjustment).native())
}

// SetValue() is a wrapper around gtk_scale_button_set_value().
func (v *scaleButton) SetValue(value float64) {
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
type scrollable struct {
	*glib_impl.Object
}

// native() returns a pointer to the underlying GObject as a GtkScrollable.
func (v *scrollable) native() *C.GtkScrollable {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkScrollable(p)
}

func wrapScrollable(obj *glib_impl.Object) *scrollable {
	return &scrollable{obj}
}

func (v *scrollable) toScrollable() *C.GtkScrollable {
	if v == nil {
		return nil
	}
	return v.native()
}

// SetHAdjustment is a wrapper around gtk_scrollable_set_hadjustment().
func (v *scrollable) SetHAdjustment(adjustment gtk.Adjustment) {
	C.gtk_scrollable_set_hadjustment(v.native(), castToAdjustment(adjustment).native())
}

// GetHAdjustment is a wrapper around gtk_scrollable_get_hadjustment().
func (v *scrollable) GetHAdjustment() (gtk.Adjustment, error) {
	c := C.gtk_scrollable_get_hadjustment(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapAdjustment(wrapObject(unsafe.Pointer(c))), nil
}

// SetVAdjustment is a wrapper around gtk_scrollable_set_vadjustment().
func (v *scrollable) SetVAdjustment(adjustment gtk.Adjustment) {
	C.gtk_scrollable_set_vadjustment(v.native(), castToAdjustment(adjustment).native())
}

// GetVAdjustment is a wrapper around gtk_scrollable_get_vadjustment().
func (v *scrollable) GetVAdjustment() (gtk.Adjustment, error) {
	c := C.gtk_scrollable_get_vadjustment(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapAdjustment(wrapObject(unsafe.Pointer(c))), nil
}

/*
 * GtkScrollbar
 */

// Scrollbar is a representation of GTK's GtkScrollbar.
type scrollbar struct {
	_range
}

// native returns a pointer to the underlying GtkScrollbar.
func (v *scrollbar) native() *C.GtkScrollbar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkScrollbar(p)
}

func marshalScrollbar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapScrollbar(obj), nil
}

func wrapScrollbar(obj *glib_impl.Object) *scrollbar {
	return &scrollbar{_range{widget{glib_impl.InitiallyUnowned{obj}}}}
}

// ScrollbarNew is a wrapper around gtk_scrollbar_new().
func ScrollbarNew(orientation gtk.Orientation, adjustment *adjustment) (*scrollbar, error) {
	c := C.gtk_scrollbar_new(C.GtkOrientation(orientation), adjustment.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapScrollbar(wrapObject(unsafe.Pointer(c))), nil
}

/*
 * GtkScrolledWindow
 */

// ScrolledWindow is a representation of GTK's GtkScrolledWindow.
type scrolledWindow struct {
	bin
}

// native returns a pointer to the underlying GtkScrolledWindow.
func (v *scrolledWindow) native() *C.GtkScrolledWindow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkScrolledWindow(p)
}

func marshalScrolledWindow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapScrolledWindow(obj), nil
}

func wrapScrolledWindow(obj *glib_impl.Object) *scrolledWindow {
	return &scrolledWindow{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}
}

// ScrolledWindowNew() is a wrapper around gtk_scrolled_window_new().
func ScrolledWindowNew(hadjustment, vadjustment *adjustment) (*scrolledWindow, error) {
	c := C.gtk_scrolled_window_new(hadjustment.native(),
		vadjustment.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapScrolledWindow(wrapObject(unsafe.Pointer(c))), nil
}

// SetPolicy() is a wrapper around gtk_scrolled_window_set_policy().
func (v *scrolledWindow) SetPolicy(hScrollbarPolicy, vScrollbarPolicy gtk.PolicyType) {
	C.gtk_scrolled_window_set_policy(v.native(),
		C.GtkPolicyType(hScrollbarPolicy),
		C.GtkPolicyType(vScrollbarPolicy))
}

// GetHAdjustment() is a wrapper around gtk_scrolled_window_get_hadjustment().
func (v *scrolledWindow) GetHAdjustment() gtk.Adjustment {
	c := C.gtk_scrolled_window_get_hadjustment(v.native())
	if c == nil {
		return nil
	}
	return wrapAdjustment(wrapObject(unsafe.Pointer(c)))
}

// SetHAdjustment is a wrapper around gtk_scrolled_window_set_hadjustment().
func (v *scrolledWindow) SetHAdjustment(adjustment gtk.Adjustment) {
	C.gtk_scrolled_window_set_hadjustment(v.native(), castToAdjustment(adjustment).native())
}

// GetVAdjustment() is a wrapper around gtk_scrolled_window_get_vadjustment().
func (v *scrolledWindow) GetVAdjustment() gtk.Adjustment {
	c := C.gtk_scrolled_window_get_vadjustment(v.native())
	if c == nil {
		return nil
	}
	return wrapAdjustment(wrapObject(unsafe.Pointer(c)))
}

// SetVAdjustment is a wrapper around gtk_scrolled_window_set_vadjustment().
func (v *scrolledWindow) SetVAdjustment(adjustment gtk.Adjustment) {
	C.gtk_scrolled_window_set_vadjustment(v.native(), castToAdjustment(adjustment).native())
}

/*
 * GtkSearchEntry
 */

// SearchEntry is a reprensentation of GTK's GtkSearchEntry.
type searchEntry struct {
	entry
}

// native returns a pointer to the underlying GtkSearchEntry.
func (v *searchEntry) native() *C.GtkSearchEntry {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSearchEntry(p)
}

func marshalSearchEntry(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapSearchEntry(obj), nil
}

func wrapSearchEntry(obj *glib_impl.Object) *searchEntry {
	e := wrapEditable(obj)
	return &searchEntry{entry{widget{glib_impl.InitiallyUnowned{obj}}, *e}}
}

// SearchEntryNew is a wrapper around gtk_search_entry_new().
func SearchEntryNew() (*searchEntry, error) {
	c := C.gtk_search_entry_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSearchEntry(wrapObject(unsafe.Pointer(c))), nil
}

/*
* GtkSelectionData
 */
type selectionData struct {
	GtkSelectionData *C.GtkSelectionData
}

func marshalSelectionData(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*selectionData)(unsafe.Pointer(c)), nil
}

// native returns a pointer to the underlying GtkSelectionData.
func (v *selectionData) native() *C.GtkSelectionData {
	if v == nil {
		return nil
	}
	return v.GtkSelectionData
}

// GetLength is a wrapper around gtk_selection_data_get_length
func (v *selectionData) GetLength() int {
	return int(C.gtk_selection_data_get_length(v.native()))
}

// GetData is a wrapper around gtk_selection_data_get_data_with_length.
// It returns a slice of the correct size with the selection's data.
func (v *selectionData) GetData() (data []byte) {
	var length C.gint
	c := C.gtk_selection_data_get_data_with_length(v.native(), &length)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sliceHeader.Data = uintptr(unsafe.Pointer(c))
	sliceHeader.Len = int(length)
	sliceHeader.Cap = int(length)
	return
}

func (v *selectionData) free() {
	C.gtk_selection_data_free(v.native())
}

/*
 * GtkSeparator
 */

// Separator is a representation of GTK's GtkSeparator.
type separator struct {
	widget
}

// native returns a pointer to the underlying GtkSeperator.
func (v *separator) native() *C.GtkSeparator {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSeparator(p)
}

func marshalSeparator(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapSeparator(obj), nil
}

func wrapSeparator(obj *glib_impl.Object) *separator {
	return &separator{widget{glib_impl.InitiallyUnowned{obj}}}
}

// SeparatorNew is a wrapper around gtk_separator_new().
func SeparatorNew(orientation gtk.Orientation) (*separator, error) {
	c := C.gtk_separator_new(C.GtkOrientation(orientation))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSeparator(wrapObject(unsafe.Pointer(c))), nil
}

/*
 * GtkSeparatorMenuItem
 */

// SeparatorMenuItem is a representation of GTK's GtkSeparatorMenuItem.
type separatorMenuItem struct {
	menuItem
}

// native returns a pointer to the underlying GtkSeparatorMenuItem.
func (v *separatorMenuItem) native() *C.GtkSeparatorMenuItem {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSeparatorMenuItem(p)
}

func marshalSeparatorMenuItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapSeparatorMenuItem(obj), nil
}

func wrapSeparatorMenuItem(obj *glib_impl.Object) *separatorMenuItem {
	return &separatorMenuItem{menuItem{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}}
}

// SeparatorMenuItemNew is a wrapper around gtk_separator_menu_item_new().
func SeparatorMenuItemNew() (*separatorMenuItem, error) {
	c := C.gtk_separator_menu_item_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSeparatorMenuItem(wrapObject(unsafe.Pointer(c))), nil
}

/*
 * GtkSeparatorToolItem
 */

// SeparatorToolItem is a representation of GTK's GtkSeparatorToolItem.
type separatorToolItem struct {
	toolItem
}

// native returns a pointer to the underlying GtkSeparatorToolItem.
func (v *separatorToolItem) native() *C.GtkSeparatorToolItem {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSeparatorToolItem(p)
}

func marshalSeparatorToolItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapSeparatorToolItem(obj), nil
}

func wrapSeparatorToolItem(obj *glib_impl.Object) *separatorToolItem {
	return &separatorToolItem{toolItem{bin{container{widget{
		glib_impl.InitiallyUnowned{obj}}}}}}
}

// SeparatorToolItemNew is a wrapper around gtk_separator_tool_item_new().
func SeparatorToolItemNew() (*separatorToolItem, error) {
	c := C.gtk_separator_tool_item_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSeparatorToolItem(wrapObject(unsafe.Pointer(c))), nil
}

// SetDraw is a wrapper around gtk_separator_tool_item_set_draw().
func (v *separatorToolItem) SetDraw(draw bool) {
	C.gtk_separator_tool_item_set_draw(v.native(), gbool(draw))
}

// GetDraw is a wrapper around gtk_separator_tool_item_get_draw().
func (v *separatorToolItem) GetDraw() bool {
	c := C.gtk_separator_tool_item_get_draw(v.native())
	return gobool(c)
}

/*
 * GtkSpinButton
 */

// SpinButton is a representation of GTK's GtkSpinButton.
type spinButton struct {
	entry
}

// native returns a pointer to the underlying GtkSpinButton.
func (v *spinButton) native() *C.GtkSpinButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSpinButton(p)
}

func marshalSpinButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapSpinButton(obj), nil
}

func wrapSpinButton(obj *glib_impl.Object) *spinButton {
	e := wrapEditable(obj)
	return &spinButton{entry{widget{glib_impl.InitiallyUnowned{obj}}, *e}}
}

// Configure() is a wrapper around gtk_spin_button_configure().
func (v *spinButton) Configure(adjustment gtk.Adjustment, climbRate float64, digits uint) {
	C.gtk_spin_button_configure(v.native(), castToAdjustment(adjustment).native(),
		C.gdouble(climbRate), C.guint(digits))
}

// SpinButtonNew() is a wrapper around gtk_spin_button_new().
func SpinButtonNew(adjustment gtk.Adjustment, climbRate float64, digits uint) (gtk.SpinButton, error) {
	c := C.gtk_spin_button_new(castToAdjustment(adjustment).native(),
		C.gdouble(climbRate), C.guint(digits))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSpinButton(wrapObject(unsafe.Pointer(c))), nil
}

// SpinButtonNewWithRange() is a wrapper around
// gtk_spin_button_new_with_range().
func SpinButtonNewWithRange(min, max, step float64) (*spinButton, error) {
	c := C.gtk_spin_button_new_with_range(C.gdouble(min), C.gdouble(max),
		C.gdouble(step))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSpinButton(wrapObject(unsafe.Pointer(c))), nil
}

// GetValueAsInt() is a wrapper around gtk_spin_button_get_value_as_int().
func (v *spinButton) GetValueAsInt() int {
	c := C.gtk_spin_button_get_value_as_int(v.native())
	return int(c)
}

// SetValue() is a wrapper around gtk_spin_button_set_value().
func (v *spinButton) SetValue(value float64) {
	C.gtk_spin_button_set_value(v.native(), C.gdouble(value))
}

// GetValue() is a wrapper around gtk_spin_button_get_value().
func (v *spinButton) GetValue() float64 {
	c := C.gtk_spin_button_get_value(v.native())
	return float64(c)
}

// GetAdjustment() is a wrapper around gtk_spin_button_get_adjustment
func (v *spinButton) GetAdjustment() gtk.Adjustment {
	c := C.gtk_spin_button_get_adjustment(v.native())
	if c == nil {
		return nil
	}
	return wrapAdjustment(wrapObject(unsafe.Pointer(c)))
}

// SetRange is a wrapper around gtk_spin_button_set_range().
func (v *spinButton) SetRange(min, max float64) {
	C.gtk_spin_button_set_range(v.native(), C.gdouble(min), C.gdouble(max))
}

// SetIncrements() is a wrapper around gtk_spin_button_set_increments().
func (v *spinButton) SetIncrements(step, page float64) {
	C.gtk_spin_button_set_increments(v.native(), C.gdouble(step), C.gdouble(page))
}

/*
 * GtkSpinner
 */

// Spinner is a representation of GTK's GtkSpinner.
type spinner struct {
	widget
}

// native returns a pointer to the underlying GtkSpinner.
func (v *spinner) native() *C.GtkSpinner {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSpinner(p)
}

func marshalSpinner(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapSpinner(obj), nil
}

func wrapSpinner(obj *glib_impl.Object) *spinner {
	return &spinner{widget{glib_impl.InitiallyUnowned{obj}}}
}

// SpinnerNew is a wrapper around gtk_spinner_new().
func SpinnerNew() (*spinner, error) {
	c := C.gtk_spinner_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSpinner(wrapObject(unsafe.Pointer(c))), nil
}

// Start is a wrapper around gtk_spinner_start().
func (v *spinner) Start() {
	C.gtk_spinner_start(v.native())
}

// Stop is a wrapper around gtk_spinner_stop().
func (v *spinner) Stop() {
	C.gtk_spinner_stop(v.native())
}

/*
 * GtkStatusbar
 */

// Statusbar is a representation of GTK's GtkStatusbar
type statusbar struct {
	box
}

// native returns a pointer to the underlying GtkStatusbar
func (v *statusbar) native() *C.GtkStatusbar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkStatusbar(p)
}

func marshalStatusbar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapStatusbar(obj), nil
}

func wrapStatusbar(obj *glib_impl.Object) *statusbar {
	return &statusbar{box{container{widget{glib_impl.InitiallyUnowned{obj}}}}}
}

// StatusbarNew() is a wrapper around gtk_statusbar_new().
func StatusbarNew() (*statusbar, error) {
	c := C.gtk_statusbar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapStatusbar(wrapObject(unsafe.Pointer(c))), nil
}

// GetContextId() is a wrapper around gtk_statusbar_get_context_id().
func (v *statusbar) GetContextId(contextDescription string) uint {
	cstr := C.CString(contextDescription)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_statusbar_get_context_id(v.native(), (*C.gchar)(cstr))
	return uint(c)
}

// Push() is a wrapper around gtk_statusbar_push().
func (v *statusbar) Push(contextID uint, text string) uint {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_statusbar_push(v.native(), C.guint(contextID),
		(*C.gchar)(cstr))
	return uint(c)
}

// Pop() is a wrapper around gtk_statusbar_pop().
func (v *statusbar) Pop(contextID uint) {
	C.gtk_statusbar_pop(v.native(), C.guint(contextID))
}

// GetMessageArea() is a wrapper around gtk_statusbar_get_message_area().
func (v *statusbar) GetMessageArea() (gtk.Box, error) {
	c := C.gtk_statusbar_get_message_area(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return &box{container{widget{glib_impl.InitiallyUnowned{obj}}}}, nil
}

/*
 * GtkSwitch
 */

// Switch is a representation of GTK's GtkSwitch.
type _switch struct {
	widget
}

// native returns a pointer to the underlying GtkSwitch.
func (v *_switch) native() *C.GtkSwitch {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSwitch(p)
}

func marshalSwitch(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapSwitch(obj), nil
}

func wrapSwitch(obj *glib_impl.Object) *_switch {
	return &_switch{widget{glib_impl.InitiallyUnowned{obj}}}
}

// SwitchNew is a wrapper around gtk_switch_new().
func SwitchNew() (*_switch, error) {
	c := C.gtk_switch_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSwitch(wrapObject(unsafe.Pointer(c))), nil
}

// GetActive is a wrapper around gtk_switch_get_active().
func (v *_switch) GetActive() bool {
	c := C.gtk_switch_get_active(v.native())
	return gobool(c)
}

// SetActive is a wrapper around gtk_switch_set_active().
func (v *_switch) SetActive(isActive bool) {
	C.gtk_switch_set_active(v.native(), gbool(isActive))
}

/*
 * GtkTargetEntry
 */

// TargetEntry is a representation of GTK's GtkTargetEntry
type targetEntry C.GtkTargetEntry

func marshalTargetEntry(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*targetEntry)(unsafe.Pointer(c)), nil
}

func (v *targetEntry) native() *C.GtkTargetEntry {
	return (*C.GtkTargetEntry)(unsafe.Pointer(v))
}

// TargetEntryNew is a wrapper aroud gtk_target_entry_new().
func TargetEntryNew(target string, flags gtk.TargetFlags, info uint) (gtk.TargetEntry, error) {
	cstr := C.CString(target)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_target_entry_new((*C.gchar)(cstr), C.guint(flags), C.guint(info))
	if c == nil {
		return nil, nilPtrErr
	}
	t := (*targetEntry)(unsafe.Pointer(c))
	runtime.SetFinalizer(t, (*targetEntry).free)
	return t, nil
}

func (v *targetEntry) free() {
	C.gtk_target_entry_free(v.native())
}

/*
 * GtkTextView
 */

// TextView is a representation of GTK's GtkTextView
type textView struct {
	container
}

// native returns a pointer to the underlying GtkTextView.
func (v *textView) native() *C.GtkTextView {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTextView(p)
}

func marshalTextView(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapTextView(obj), nil
}

func wrapTextView(obj *glib_impl.Object) *textView {
	return &textView{container{widget{glib_impl.InitiallyUnowned{obj}}}}
}

// TextViewNew is a wrapper around gtk_text_view_new().
func TextViewNew() (*textView, error) {
	c := C.gtk_text_view_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapTextView(wrapObject(unsafe.Pointer(c))), nil
}

// TextViewNewWithBuffer is a wrapper around gtk_text_view_new_with_buffer().
func TextViewNewWithBuffer(buf *textBuffer) (*textView, error) {
	cbuf := buf.native()
	c := C.gtk_text_view_new_with_buffer(cbuf)
	return wrapTextView(wrapObject(unsafe.Pointer(c))), nil
}

// GetBuffer is a wrapper around gtk_text_view_get_buffer().
func (v *textView) GetBuffer() (gtk.TextBuffer, error) {
	c := C.gtk_text_view_get_buffer(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapTextBuffer(wrapObject(unsafe.Pointer(c))), nil
}

// SetBuffer is a wrapper around gtk_text_view_set_buffer().
func (v *textView) SetBuffer(buffer gtk.TextBuffer) {
	C.gtk_text_view_set_buffer(v.native(), castToTextBuffer(buffer).native())
}

// SetEditable is a wrapper around gtk_text_view_set_editable().
func (v *textView) SetEditable(editable bool) {
	C.gtk_text_view_set_editable(v.native(), gbool(editable))
}

// GetEditable is a wrapper around gtk_text_view_get_editable().
func (v *textView) GetEditable() bool {
	c := C.gtk_text_view_get_editable(v.native())
	return gobool(c)
}

// SetWrapMode is a wrapper around gtk_text_view_set_wrap_mode().
func (v *textView) SetWrapMode(wrapMode gtk.WrapMode) {
	C.gtk_text_view_set_wrap_mode(v.native(), C.GtkWrapMode(wrapMode))
}

// GetWrapMode is a wrapper around gtk_text_view_get_wrap_mode().
func (v *textView) GetWrapMode() gtk.WrapMode {
	return gtk.WrapMode(C.gtk_text_view_get_wrap_mode(v.native()))
}

// SetCursorVisible is a wrapper around gtk_text_view_set_cursor_visible().
func (v *textView) SetCursorVisible(visible bool) {
	C.gtk_text_view_set_cursor_visible(v.native(), gbool(visible))
}

// GetCursorVisible is a wrapper around gtk_text_view_get_cursor_visible().
func (v *textView) GetCursorVisible() bool {
	c := C.gtk_text_view_get_cursor_visible(v.native())
	return gobool(c)
}

// SetOverwrite is a wrapper around gtk_text_view_set_overwrite().
func (v *textView) SetOverwrite(overwrite bool) {
	C.gtk_text_view_set_overwrite(v.native(), gbool(overwrite))
}

// GetOverwrite is a wrapper around gtk_text_view_get_overwrite().
func (v *textView) GetOverwrite() bool {
	c := C.gtk_text_view_get_overwrite(v.native())
	return gobool(c)
}

// SetJustification is a wrapper around gtk_text_view_set_justification().
func (v *textView) SetJustification(justify gtk.Justification) {
	C.gtk_text_view_set_justification(v.native(), C.GtkJustification(justify))
}

// GetJustification is a wrapper around gtk_text_view_get_justification().
func (v *textView) GetJustification() gtk.Justification {
	c := C.gtk_text_view_get_justification(v.native())
	return gtk.Justification(c)
}

// SetAcceptsTab is a wrapper around gtk_text_view_set_accepts_tab().
func (v *textView) SetAcceptsTab(acceptsTab bool) {
	C.gtk_text_view_set_accepts_tab(v.native(), gbool(acceptsTab))
}

// GetAcceptsTab is a wrapper around gtk_text_view_get_accepts_tab().
func (v *textView) GetAcceptsTab() bool {
	c := C.gtk_text_view_get_accepts_tab(v.native())
	return gobool(c)
}

// SetPixelsAboveLines is a wrapper around gtk_text_view_set_pixels_above_lines().
func (v *textView) SetPixelsAboveLines(px int) {
	C.gtk_text_view_set_pixels_above_lines(v.native(), C.gint(px))
}

// GetPixelsAboveLines is a wrapper around gtk_text_view_get_pixels_above_lines().
func (v *textView) GetPixelsAboveLines() int {
	c := C.gtk_text_view_get_pixels_above_lines(v.native())
	return int(c)
}

// SetPixelsBelowLines is a wrapper around gtk_text_view_set_pixels_below_lines().
func (v *textView) SetPixelsBelowLines(px int) {
	C.gtk_text_view_set_pixels_below_lines(v.native(), C.gint(px))
}

// GetPixelsBelowLines is a wrapper around gtk_text_view_get_pixels_below_lines().
func (v *textView) GetPixelsBelowLines() int {
	c := C.gtk_text_view_get_pixels_below_lines(v.native())
	return int(c)
}

// SetPixelsInsideWrap is a wrapper around gtk_text_view_set_pixels_inside_wrap().
func (v *textView) SetPixelsInsideWrap(px int) {
	C.gtk_text_view_set_pixels_inside_wrap(v.native(), C.gint(px))
}

// GetPixelsInsideWrap is a wrapper around gtk_text_view_get_pixels_inside_wrap().
func (v *textView) GetPixelsInsideWrap() int {
	c := C.gtk_text_view_get_pixels_inside_wrap(v.native())
	return int(c)
}

// SetLeftMargin is a wrapper around gtk_text_view_set_left_margin().
func (v *textView) SetLeftMargin(margin int) {
	C.gtk_text_view_set_left_margin(v.native(), C.gint(margin))
}

// GetLeftMargin is a wrapper around gtk_text_view_get_left_margin().
func (v *textView) GetLeftMargin() int {
	c := C.gtk_text_view_get_left_margin(v.native())
	return int(c)
}

// SetRightMargin is a wrapper around gtk_text_view_set_right_margin().
func (v *textView) SetRightMargin(margin int) {
	C.gtk_text_view_set_right_margin(v.native(), C.gint(margin))
}

// GetRightMargin is a wrapper around gtk_text_view_get_right_margin().
func (v *textView) GetRightMargin() int {
	c := C.gtk_text_view_get_right_margin(v.native())
	return int(c)
}

// SetIndent is a wrapper around gtk_text_view_set_indent().
func (v *textView) SetIndent(indent int) {
	C.gtk_text_view_set_indent(v.native(), C.gint(indent))
}

// GetIndent is a wrapper around gtk_text_view_get_indent().
func (v *textView) GetIndent() int {
	c := C.gtk_text_view_get_indent(v.native())
	return int(c)
}

// SetInputHints is a wrapper around gtk_text_view_set_input_hints().
func (v *textView) SetInputHints(hints gtk.InputHints) {
	C.gtk_text_view_set_input_hints(v.native(), C.GtkInputHints(hints))
}

// GetInputHints is a wrapper around gtk_text_view_get_input_hints().
func (v *textView) GetInputHints() gtk.InputHints {
	c := C.gtk_text_view_get_input_hints(v.native())
	return gtk.InputHints(c)
}

// SetInputPurpose is a wrapper around gtk_text_view_set_input_purpose().
func (v *textView) SetInputPurpose(purpose gtk.InputPurpose) {
	C.gtk_text_view_set_input_purpose(v.native(),
		C.GtkInputPurpose(purpose))
}

// GetInputPurpose is a wrapper around gtk_text_view_get_input_purpose().
func (v *textView) GetInputPurpose() gtk.InputPurpose {
	c := C.gtk_text_view_get_input_purpose(v.native())
	return gtk.InputPurpose(c)
}

/*
 * GtkTextTag
 */

type textTag struct {
	*glib_impl.Object
}

// native returns a pointer to the underlying GObject as a GtkTextTag.
func (v *textTag) native() *C.GtkTextTag {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTextTag(p)
}

func marshalTextTag(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapTextTag(obj), nil
}

func wrapTextTag(obj *glib_impl.Object) *textTag {
	return &textTag{obj}
}

func TextTagNew(name string) (*textTag, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	c := C.gtk_text_tag_new((*C.gchar)(cname))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapTextTag(wrapObject(unsafe.Pointer(c))), nil
}

// GetPriority() is a wrapper around gtk_text_tag_get_priority().
func (v *textTag) GetPriority() int {
	return int(C.gtk_text_tag_get_priority(v.native()))
}

// SetPriority() is a wrapper around gtk_text_tag_set_priority().
func (v *textTag) SetPriority(priority int) {
	C.gtk_text_tag_set_priority(v.native(), C.gint(priority))
}

// Event() is a wrapper around gtk_text_tag_event().
func (v *textTag) Event(eventObject glib.Object, event gdk.Event, iter gtk.TextIter) bool {
	ok := C.gtk_text_tag_event(v.native(),
		(*C.GObject)(unsafe.Pointer(glib_impl.CastToObject(eventObject).Native())),
		(*C.GdkEvent)(unsafe.Pointer(gdk_impl.CastToEvent(event).Native())),
		(*C.GtkTextIter)(castToTextIter(iter)),
	)
	return gobool(ok)
}

/*
 * GtkTextTagTable
 */

type textTagTable struct {
	*glib_impl.Object
}

// native returns a pointer to the underlying GObject as a GtkTextTagTable.
func (v *textTagTable) native() *C.GtkTextTagTable {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTextTagTable(p)
}

func marshalTextTagTable(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapTextTagTable(obj), nil
}

func wrapTextTagTable(obj *glib_impl.Object) *textTagTable {
	return &textTagTable{obj}
}

func TextTagTableNew() (*textTagTable, error) {
	c := C.gtk_text_tag_table_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapTextTagTable(wrapObject(unsafe.Pointer(c))), nil
}

// Add() is a wrapper around gtk_text_tag_table_add().
func (v *textTagTable) Add(tag gtk.TextTag) {
	C.gtk_text_tag_table_add(v.native(), castToTextTag(tag).native())
	//return gobool(c) // TODO version-separate
}

// Lookup() is a wrapper around gtk_text_tag_table_lookup().
func (v *textTagTable) Lookup(name string) (gtk.TextTag, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	c := C.gtk_text_tag_table_lookup(v.native(), (*C.gchar)(cname))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapTextTag(wrapObject(unsafe.Pointer(c))), nil
}

// Remove() is a wrapper around gtk_text_tag_table_remove().
func (v *textTagTable) Remove(tag gtk.TextTag) {
	C.gtk_text_tag_table_remove(v.native(), castToTextTag(tag).native())
}

/*
 * GtkTextBuffer
 */

// TextBuffer is a representation of GTK's GtkTextBuffer.
type textBuffer struct {
	*glib_impl.Object
}

// native returns a pointer to the underlying GtkTextBuffer.
func (v *textBuffer) native() *C.GtkTextBuffer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTextBuffer(p)
}

func marshalTextBuffer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapTextBuffer(obj), nil
}

func wrapTextBuffer(obj *glib_impl.Object) *textBuffer {
	return &textBuffer{obj}
}

// TextBufferNew() is a wrapper around gtk_text_buffer_new().
func TextBufferNew(table *textTagTable) (*textBuffer, error) {
	c := C.gtk_text_buffer_new(table.native())
	if c == nil {
		return nil, nilPtrErr
	}

	e := wrapTextBuffer(wrapObject(unsafe.Pointer(c)))
	return e, nil
}

// ApplyTag() is a wrapper around gtk_text_buffer_apply_tag().
func (v *textBuffer) ApplyTag(tag gtk.TextTag, start, end gtk.TextIter) {
	C.gtk_text_buffer_apply_tag(v.native(), castToTextTag(tag).native(), (*C.GtkTextIter)(castToTextIter(start)), (*C.GtkTextIter)(castToTextIter(end)))
}

// ApplyTagByName() is a wrapper around gtk_text_buffer_apply_tag_by_name().
func (v *textBuffer) ApplyTagByName(name string, start, end gtk.TextIter) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_apply_tag_by_name(v.native(), (*C.gchar)(cstr),
		(*C.GtkTextIter)(castToTextIter(start)), (*C.GtkTextIter)(castToTextIter(end)))
}

// Delete() is a wrapper around gtk_text_buffer_delete().
func (v *textBuffer) Delete(start, end gtk.TextIter) {
	C.gtk_text_buffer_delete(v.native(), (*C.GtkTextIter)(castToTextIter(start)), (*C.GtkTextIter)(castToTextIter(end)))
}

func (v *textBuffer) GetBounds() (start, end gtk.TextIter) {
	start, end = new(textIter), new(textIter)
	C.gtk_text_buffer_get_bounds(v.native(), (*C.GtkTextIter)(castToTextIter(start)), (*C.GtkTextIter)(castToTextIter(end)))
	return
}

// GetCharCount() is a wrapper around gtk_text_buffer_get_char_count().
func (v *textBuffer) GetCharCount() int {
	return int(C.gtk_text_buffer_get_char_count(v.native()))
}

// GetIterAtOffset() is a wrapper around gtk_text_buffer_get_iter_at_offset().
func (v *textBuffer) GetIterAtOffset(charOffset int) gtk.TextIter {
	var iter C.GtkTextIter
	C.gtk_text_buffer_get_iter_at_offset(v.native(), &iter, C.gint(charOffset))
	return (*textIter)(&iter)
}

// GetStartIter() is a wrapper around gtk_text_buffer_get_start_iter().
func (v *textBuffer) GetStartIter() gtk.TextIter {
	var iter C.GtkTextIter
	C.gtk_text_buffer_get_start_iter(v.native(), &iter)
	return (*textIter)(&iter)
}

// GetEndIter() is a wrapper around gtk_text_buffer_get_end_iter().
func (v *textBuffer) GetEndIter() gtk.TextIter {
	var iter C.GtkTextIter
	C.gtk_text_buffer_get_end_iter(v.native(), &iter)
	return (*textIter)(&iter)
}

// GetLineCount() is a wrapper around gtk_text_buffer_get_line_count().
func (v *textBuffer) GetLineCount() int {
	return int(C.gtk_text_buffer_get_line_count(v.native()))
}

// GetModified() is a wrapper around gtk_text_buffer_get_modified().
func (v *textBuffer) GetModified() bool {
	return gobool(C.gtk_text_buffer_get_modified(v.native()))
}

// GetTagTable() is a wrapper around gtk_text_buffer_get_tag_table().
func (v *textBuffer) GetTagTable() (gtk.TextTagTable, error) {
	c := C.gtk_text_buffer_get_tag_table(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapTextTagTable(obj), nil
}

func (v *textBuffer) GetText(start, end gtk.TextIter, includeHiddenChars bool) (string, error) {
	c := C.gtk_text_buffer_get_text(
		v.native(), (*C.GtkTextIter)(castToTextIter(start)), (*C.GtkTextIter)(castToTextIter(end)), gbool(includeHiddenChars),
	)
	if c == nil {
		return "", nilPtrErr
	}
	gostr := C.GoString((*C.char)(c))
	C.g_free(C.gpointer(c))
	return gostr, nil
}

// Insert() is a wrapper around gtk_text_buffer_insert().
func (v *textBuffer) Insert(iter gtk.TextIter, text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_insert(v.native(), (*C.GtkTextIter)(castToTextIter(iter)), (*C.gchar)(cstr), C.gint(len(text)))
}

// InsertAtCursor() is a wrapper around gtk_text_buffer_insert_at_cursor().
func (v *textBuffer) InsertAtCursor(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_insert_at_cursor(v.native(), (*C.gchar)(cstr), C.gint(len(text)))
}

// RemoveTag() is a wrapper around gtk_text_buffer_remove_tag().
func (v *textBuffer) RemoveTag(tag gtk.TextTag, start, end gtk.TextIter) {
	C.gtk_text_buffer_remove_tag(v.native(), castToTextTag(tag).native(), (*C.GtkTextIter)(castToTextIter(start)), (*C.GtkTextIter)(castToTextIter(end)))
}

// SetModified() is a wrapper around gtk_text_buffer_set_modified().
func (v *textBuffer) SetModified(setting bool) {
	C.gtk_text_buffer_set_modified(v.native(), gbool(setting))
}

func (v *textBuffer) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_set_text(v.native(), (*C.gchar)(cstr),
		C.gint(len(text)))
}

/*
 * GtkTextIter
 */

// TextIter is a representation of GTK's GtkTextIter
type textIter C.GtkTextIter

func marshalTextIter(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*textIter)(unsafe.Pointer(c)), nil
}

/*
 * GtkToggleButton
 */

// ToggleButton is a representation of GTK's GtkToggleButton.
type toggleButton struct {
	button
}

// native returns a pointer to the underlying GtkToggleButton.
func (v *toggleButton) native() *C.GtkToggleButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkToggleButton(p)
}

func marshalToggleButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapToggleButton(obj), nil
}

func wrapToggleButton(obj *glib_impl.Object) *toggleButton {
	return &toggleButton{button{bin{container{widget{
		glib_impl.InitiallyUnowned{obj}}}}}}
}

// ToggleButtonNew is a wrapper around gtk_toggle_button_new().
func ToggleButtonNew() (*toggleButton, error) {
	c := C.gtk_toggle_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapToggleButton(wrapObject(unsafe.Pointer(c))), nil
}

// ToggleButtonNewWithLabel is a wrapper around
// gtk_toggle_button_new_with_label().
func ToggleButtonNewWithLabel(label string) (*toggleButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_toggle_button_new_with_label((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapToggleButton(wrapObject(unsafe.Pointer(c))), nil
}

// ToggleButtonNewWithMnemonic is a wrapper around
// gtk_toggle_button_new_with_mnemonic().
func ToggleButtonNewWithMnemonic(label string) (*toggleButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_toggle_button_new_with_mnemonic((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapToggleButton(wrapObject(unsafe.Pointer(c))), nil
}

// GetActive is a wrapper around gtk_toggle_button_get_active().
func (v *toggleButton) GetActive() bool {
	c := C.gtk_toggle_button_get_active(v.native())
	return gobool(c)
}

// SetActive is a wrapper around gtk_toggle_button_set_active().
func (v *toggleButton) SetActive(isActive bool) {
	C.gtk_toggle_button_set_active(v.native(), gbool(isActive))
}

/*
 * GtkToolbar
 */

// Toolbar is a representation of GTK's GtkToolbar.
type toolbar struct {
	container
}

// native returns a pointer to the underlying GtkToolbar.
func (v *toolbar) native() *C.GtkToolbar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkToolbar(p)
}

func marshalToolbar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapToolbar(obj), nil
}

func wrapToolbar(obj *glib_impl.Object) *toolbar {
	return &toolbar{container{widget{glib_impl.InitiallyUnowned{obj}}}}
}

// ToolbarNew is a wrapper around gtk_toolbar_new().
func ToolbarNew() (*toolbar, error) {
	c := C.gtk_toolbar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapToolbar(wrapObject(unsafe.Pointer(c))), nil
}

// Insert is a wrapper around gtk_toolbar_insert().
func (v *toolbar) Insert(item gtk.ToolItem, pos int) {
	C.gtk_toolbar_insert(v.native(), item.(IToolItem).toToolItem(), C.gint(pos))
}

// GetItemIndex is a wrapper around gtk_toolbar_get_item_index().
func (v *toolbar) GetItemIndex(item gtk.ToolItem) int {
	c := C.gtk_toolbar_get_item_index(v.native(), item.(IToolItem).toToolItem())
	return int(c)
}

// GetNItems is a wrapper around gtk_toolbar_get_n_items().
func (v *toolbar) GetNItems() int {
	c := C.gtk_toolbar_get_n_items(v.native())
	return int(c)
}

// GetNthItem is a wrapper around gtk_toolbar_get_nth_item().
func (v *toolbar) GetNthItem(n int) gtk.ToolItem {
	c := C.gtk_toolbar_get_nth_item(v.native(), C.gint(n))
	if c == nil {
		return nil
	}
	return wrapToolItem(wrapObject(unsafe.Pointer(c)))
}

// GetDropIndex is a wrapper around gtk_toolbar_get_drop_index().
func (v *toolbar) GetDropIndex(x, y int) int {
	c := C.gtk_toolbar_get_drop_index(v.native(), C.gint(x), C.gint(y))
	return int(c)
}

// SetDropHighlightItem is a wrapper around
// gtk_toolbar_set_drop_highlight_item().
func (v *toolbar) SetDropHighlightItem(toolItem gtk.ToolItem, index int) {
	C.gtk_toolbar_set_drop_highlight_item(v.native(),
		toolItem.(IToolItem).toToolItem(), C.gint(index))
}

// SetShowArrow is a wrapper around gtk_toolbar_set_show_arrow().
func (v *toolbar) SetShowArrow(showArrow bool) {
	C.gtk_toolbar_set_show_arrow(v.native(), gbool(showArrow))
}

// UnsetIconSize is a wrapper around gtk_toolbar_unset_icon_size().
func (v *toolbar) UnsetIconSize() {
	C.gtk_toolbar_unset_icon_size(v.native())
}

// GetShowArrow is a wrapper around gtk_toolbar_get_show_arrow().
func (v *toolbar) GetShowArrow() bool {
	c := C.gtk_toolbar_get_show_arrow(v.native())
	return gobool(c)
}

// GetStyle is a wrapper around gtk_toolbar_get_style().
func (v *toolbar) GetStyle() gtk.ToolbarStyle {
	c := C.gtk_toolbar_get_style(v.native())
	return gtk.ToolbarStyle(c)
}

// GetIconSize is a wrapper around gtk_toolbar_get_icon_size().
func (v *toolbar) GetIconSize() gtk.IconSize {
	c := C.gtk_toolbar_get_icon_size(v.native())
	return gtk.IconSize(c)
}

// GetReliefStyle is a wrapper around gtk_toolbar_get_relief_style().
func (v *toolbar) GetReliefStyle() gtk.ReliefStyle {
	c := C.gtk_toolbar_get_relief_style(v.native())
	return gtk.ReliefStyle(c)
}

// SetStyle is a wrapper around gtk_toolbar_set_style().
func (v *toolbar) SetStyle(style gtk.ToolbarStyle) {
	C.gtk_toolbar_set_style(v.native(), C.GtkToolbarStyle(style))
}

// SetIconSize is a wrapper around gtk_toolbar_set_icon_size().
func (v *toolbar) SetIconSize(iconSize gtk.IconSize) {
	C.gtk_toolbar_set_icon_size(v.native(), C.GtkIconSize(iconSize))
}

// UnsetStyle is a wrapper around gtk_toolbar_unset_style().
func (v *toolbar) UnsetStyle() {
	C.gtk_toolbar_unset_style(v.native())
}

/*
 * GtkToolButton
 */

// ToolButton is a representation of GTK's GtkToolButton.
type toolButton struct {
	toolItem
}

// native returns a pointer to the underlying GtkToolButton.
func (v *toolButton) native() *C.GtkToolButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkToolButton(p)
}

func marshalToolButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapToolButton(obj), nil
}

func wrapToolButton(obj *glib_impl.Object) *toolButton {
	return &toolButton{toolItem{bin{container{widget{
		glib_impl.InitiallyUnowned{obj}}}}}}
}

// ToolButtonNew is a wrapper around gtk_tool_button_new().
func ToolButtonNew(iconWidget gtk.Widget, label string) (*toolButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	w := nullableWidget(iconWidget)
	c := C.gtk_tool_button_new(w, (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapToolButton(wrapObject(unsafe.Pointer(c))), nil
}

// SetLabel is a wrapper around gtk_tool_button_set_label().
func (v *toolButton) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tool_button_set_label(v.native(), (*C.gchar)(cstr))
}

// GetLabel is a wrapper aroud gtk_tool_button_get_label().
func (v *toolButton) GetLabel() string {
	c := C.gtk_tool_button_get_label(v.native())
	return C.GoString((*C.char)(c))
}

// SetUseUnderline is a wrapper around gtk_tool_button_set_use_underline().
func (v *toolButton) SetGetUnderline(useUnderline bool) {
	C.gtk_tool_button_set_use_underline(v.native(), gbool(useUnderline))
}

// GetUseUnderline is a wrapper around gtk_tool_button_get_use_underline().
func (v *toolButton) GetuseUnderline() bool {
	c := C.gtk_tool_button_get_use_underline(v.native())
	return gobool(c)
}

// SetIconName is a wrapper around gtk_tool_button_set_icon_name().
func (v *toolButton) SetIconName(iconName string) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tool_button_set_icon_name(v.native(), (*C.gchar)(cstr))
}

// GetIconName is a wrapper around gtk_tool_button_get_icon_name().
func (v *toolButton) GetIconName() string {
	c := C.gtk_tool_button_get_icon_name(v.native())
	return C.GoString((*C.char)(c))
}

// SetIconWidget is a wrapper around gtk_tool_button_set_icon_widget().
func (v *toolButton) SetIconWidget(iconWidget gtk.Widget) {
	C.gtk_tool_button_set_icon_widget(v.native(), iconWidget.(IWidget).toWidget())
}

// GetIconWidget is a wrapper around gtk_tool_button_get_icon_widget().
func (v *toolButton) GetIconWidget() gtk.Widget {
	c := C.gtk_tool_button_get_icon_widget(v.native())
	if c == nil {
		return nil
	}
	return wrapWidget(wrapObject(unsafe.Pointer(c)))
}

// SetLabelWidget is a wrapper around gtk_tool_button_set_label_widget().
func (v *toolButton) SetLabelWidget(labelWidget gtk.Widget) {
	C.gtk_tool_button_set_label_widget(v.native(), labelWidget.(IWidget).toWidget())
}

// GetLabelWidget is a wrapper around gtk_tool_button_get_label_widget().
func (v *toolButton) GetLabelWidget() gtk.Widget {
	c := C.gtk_tool_button_get_label_widget(v.native())
	if c == nil {
		return nil
	}
	return wrapWidget(wrapObject(unsafe.Pointer(c)))
}

/*
 * GtkToolItem
 */

// ToolItem is a representation of GTK's GtkToolItem.
type toolItem struct {
	bin
}

// IToolItem is an interface type implemented by all structs embedding
// a ToolItem.  It is meant to be used as an argument type for wrapper
// functions that wrap around a C GTK function taking a GtkToolItem.
type IToolItem interface {
	toToolItem() *C.GtkToolItem
}

// native returns a pointer to the underlying GtkToolItem.
func (v *toolItem) native() *C.GtkToolItem {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkToolItem(p)
}

func (v *toolItem) toToolItem() *C.GtkToolItem {
	return v.native()
}

func marshalToolItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapToolItem(obj), nil
}

func wrapToolItem(obj *glib_impl.Object) *toolItem {
	return &toolItem{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}
}

// ToolItemNew is a wrapper around gtk_tool_item_new().
func ToolItemNew() (*toolItem, error) {
	c := C.gtk_tool_item_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapToolItem(wrapObject(unsafe.Pointer(c))), nil
}

// SetHomogeneous is a wrapper around gtk_tool_item_set_homogeneous().
func (v *toolItem) SetHomogeneous(homogeneous bool) {
	C.gtk_tool_item_set_homogeneous(v.native(), gbool(homogeneous))
}

// GetHomogeneous is a wrapper around gtk_tool_item_get_homogeneous().
func (v *toolItem) GetHomogeneous() bool {
	c := C.gtk_tool_item_get_homogeneous(v.native())
	return gobool(c)
}

// SetExpand is a wrapper around gtk_tool_item_set_expand().
func (v *toolItem) SetExpand(expand bool) {
	C.gtk_tool_item_set_expand(v.native(), gbool(expand))
}

// GetExpand is a wrapper around gtk_tool_item_get_expand().
func (v *toolItem) GetExpand() bool {
	c := C.gtk_tool_item_get_expand(v.native())
	return gobool(c)
}

// SetTooltipText is a wrapper around gtk_tool_item_set_tooltip_text().
func (v *toolItem) SetTooltipText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tool_item_set_tooltip_text(v.native(), (*C.gchar)(cstr))
}

// SetTooltipMarkup is a wrapper around gtk_tool_item_set_tooltip_markup().
func (v *toolItem) SetTooltipMarkup(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tool_item_set_tooltip_markup(v.native(), (*C.gchar)(cstr))
}

// SetUseDragWindow is a wrapper around gtk_tool_item_set_use_drag_window().
func (v *toolItem) SetUseDragWindow(useDragWindow bool) {
	C.gtk_tool_item_set_use_drag_window(v.native(), gbool(useDragWindow))
}

// GetUseDragWindow is a wrapper around gtk_tool_item_get_use_drag_window().
func (v *toolItem) GetUseDragWindow() bool {
	c := C.gtk_tool_item_get_use_drag_window(v.native())
	return gobool(c)
}

// SetVisibleHorizontal is a wrapper around
// gtk_tool_item_set_visible_horizontal().
func (v *toolItem) SetVisibleHorizontal(visibleHorizontal bool) {
	C.gtk_tool_item_set_visible_horizontal(v.native(),
		gbool(visibleHorizontal))
}

// GetVisibleHorizontal is a wrapper around
// gtk_tool_item_get_visible_horizontal().
func (v *toolItem) GetVisibleHorizontal() bool {
	c := C.gtk_tool_item_get_visible_horizontal(v.native())
	return gobool(c)
}

// SetVisibleVertical is a wrapper around gtk_tool_item_set_visible_vertical().
func (v *toolItem) SetVisibleVertical(visibleVertical bool) {
	C.gtk_tool_item_set_visible_vertical(v.native(), gbool(visibleVertical))
}

// GetVisibleVertical is a wrapper around gtk_tool_item_get_visible_vertical().
func (v *toolItem) GetVisibleVertical() bool {
	c := C.gtk_tool_item_get_visible_vertical(v.native())
	return gobool(c)
}

// SetIsImportant is a wrapper around gtk_tool_item_set_is_important().
func (v *toolItem) SetIsImportant(isImportant bool) {
	C.gtk_tool_item_set_is_important(v.native(), gbool(isImportant))
}

// GetIsImportant is a wrapper around gtk_tool_item_get_is_important().
func (v *toolItem) GetIsImportant() bool {
	c := C.gtk_tool_item_get_is_important(v.native())
	return gobool(c)
}

// TODO: gtk_tool_item_get_ellipsize_mode

// GetIconSize is a wrapper around gtk_tool_item_get_icon_size().
func (v *toolItem) GetIconSize() gtk.IconSize {
	c := C.gtk_tool_item_get_icon_size(v.native())
	return gtk.IconSize(c)
}

// GetOrientation is a wrapper around gtk_tool_item_get_orientation().
func (v *toolItem) GetOrientation() gtk.Orientation {
	c := C.gtk_tool_item_get_orientation(v.native())
	return gtk.Orientation(c)
}

// GetToolbarStyle is a wrapper around gtk_tool_item_get_toolbar_style().
func (v *toolItem) gtk_tool_item_get_toolbar_style() gtk.ToolbarStyle {
	c := C.gtk_tool_item_get_toolbar_style(v.native())
	return gtk.ToolbarStyle(c)
}

// GetReliefStyle is a wrapper around gtk_tool_item_get_relief_style().
func (v *toolItem) GetReliefStyle() gtk.ReliefStyle {
	c := C.gtk_tool_item_get_relief_style(v.native())
	return gtk.ReliefStyle(c)
}

// GetTextAlignment is a wrapper around gtk_tool_item_get_text_alignment().
func (v *toolItem) GetTextAlignment() float32 {
	c := C.gtk_tool_item_get_text_alignment(v.native())
	return float32(c)
}

// GetTextOrientation is a wrapper around gtk_tool_item_get_text_orientation().
func (v *toolItem) GetTextOrientation() gtk.Orientation {
	c := C.gtk_tool_item_get_text_orientation(v.native())
	return gtk.Orientation(c)
}

// RetrieveProxyMenuItem is a wrapper around
// gtk_tool_item_retrieve_proxy_menu_item()
func (v *toolItem) RetrieveProxyMenuItem() gtk.MenuItem {
	c := C.gtk_tool_item_retrieve_proxy_menu_item(v.native())
	if c == nil {
		return nil
	}
	return wrapMenuItem(wrapObject(unsafe.Pointer(c)))
}

// SetProxyMenuItem is a wrapper around gtk_tool_item_set_proxy_menu_item().
func (v *toolItem) SetProxyMenuItem(menuItemId string, menuItem gtk.MenuItem) {
	cstr := C.CString(menuItemId)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tool_item_set_proxy_menu_item(v.native(), (*C.gchar)(cstr),
		C.toGtkWidget(unsafe.Pointer(menuItem.(IMenuItem).toMenuItem())))
}

// RebuildMenu is a wrapper around gtk_tool_item_rebuild_menu().
func (v *toolItem) RebuildMenu() {
	C.gtk_tool_item_rebuild_menu(v.native())
}

// ToolbarReconfigured is a wrapper around gtk_tool_item_toolbar_reconfigured().
func (v *toolItem) ToolbarReconfigured() {
	C.gtk_tool_item_toolbar_reconfigured(v.native())
}

// TODO: gtk_tool_item_get_text_size_group

/*
 * GtkTreeIter
 */

// TreeIter is a representation of GTK's GtkTreeIter.
type treeIter struct {
	GtkTreeIter C.GtkTreeIter
}

// native returns a pointer to the underlying GtkTreeIter.
func (v *treeIter) native() *C.GtkTreeIter {
	if v == nil {
		return nil
	}
	return &v.GtkTreeIter
}

func marshalTreeIter(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*treeIter)(unsafe.Pointer(c)), nil
}

func (v *treeIter) free() {
	C.gtk_tree_iter_free(v.native())
}

// TreeIterNew creates a new TreeIter
func TreeIterNew() gtk.TreeIter {
	var iter treeIter
	return &iter
}

// Copy() is a wrapper around gtk_tree_iter_copy().
func (v *treeIter) Copy() (gtk.TreeIter, error) {
	c := C.gtk_tree_iter_copy(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	t := &treeIter{*c}
	runtime.SetFinalizer(t, (*treeIter).free)
	return t, nil
}

/*
 * GtkTreeModel
 */

// TreeModel is a representation of GTK's GtkTreeModel GInterface.
type treeModel struct {
	*glib_impl.Object
}

// ITreeModel is an interface type implemented by all structs
// embedding a TreeModel.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkTreeModel.
type ITreeModel interface {
	toTreeModel() *C.GtkTreeModel
}

// native returns a pointer to the underlying GObject as a GtkTreeModel.
func (v *treeModel) native() *C.GtkTreeModel {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTreeModel(p)
}

func (v *treeModel) toTreeModel() *C.GtkTreeModel {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalTreeModel(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapTreeModel(obj), nil
}

func wrapTreeModel(obj *glib_impl.Object) *treeModel {
	return &treeModel{obj}
}

// GetFlags() is a wrapper around gtk_tree_model_get_flags().
func (v *treeModel) GetFlags() gtk.TreeModelFlags {
	c := C.gtk_tree_model_get_flags(v.native())
	return gtk.TreeModelFlags(c)
}

// GetNColumns() is a wrapper around gtk_tree_model_get_n_columns().
func (v *treeModel) GetNColumns() int {
	c := C.gtk_tree_model_get_n_columns(v.native())
	return int(c)
}

// GetColumnType() is a wrapper around gtk_tree_model_get_column_type().
func (v *treeModel) GetColumnType(index int) glib.Type {
	c := C.gtk_tree_model_get_column_type(v.native(), C.gint(index))
	return glib.Type(c)
}

// GetIter() is a wrapper around gtk_tree_model_get_iter().
func (v *treeModel) GetIter(path gtk.TreePath) (gtk.TreeIter, error) {
	var iter C.GtkTreeIter
	c := C.gtk_tree_model_get_iter(v.native(), &iter, castToTreePath(path).native())
	if !gobool(c) {
		return nil, errors.New("Unable to set iterator")
	}
	t := &treeIter{iter}
	return t, nil
}

// GetIterFromString() is a wrapper around
// gtk_tree_model_get_iter_from_string().
func (v *treeModel) GetIterFromString(path string) (gtk.TreeIter, error) {
	var iter C.GtkTreeIter
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_tree_model_get_iter_from_string(v.native(), &iter,
		(*C.gchar)(cstr))
	if !gobool(c) {
		return nil, errors.New("Unable to set iterator")
	}
	t := &treeIter{iter}
	return t, nil
}

// GetIterFirst() is a wrapper around gtk_tree_model_get_iter_first().
func (v *treeModel) GetIterFirst() (gtk.TreeIter, bool) {
	var iter C.GtkTreeIter
	c := C.gtk_tree_model_get_iter_first(v.native(), &iter)
	if !gobool(c) {
		return nil, false
	}
	t := &treeIter{iter}
	return t, true
}

// GetPath() is a wrapper around gtk_tree_model_get_path().
func (v *treeModel) GetPath(iter gtk.TreeIter) (gtk.TreePath, error) {
	c := C.gtk_tree_model_get_path(v.native(), castToTreeIter(iter).native())
	if c == nil {
		return nil, nilPtrErr
	}
	p := &treePath{c}
	runtime.SetFinalizer(p, (*treePath).free)
	return p, nil
}

// GetValue() is a wrapper around gtk_tree_model_get_value().
func (v *treeModel) GetValue(iter gtk.TreeIter, column int) (glib.Value, error) {
	val, err := glib_impl.ValueAlloc()
	if err != nil {
		return nil, err
	}
	C.gtk_tree_model_get_value(
		(*C.GtkTreeModel)(unsafe.Pointer(v.native())),
		castToTreeIter(iter).native(),
		C.gint(column),
		(*C.GValue)(unsafe.Pointer(val.Native())))
	return val, nil
}

// IterNext() is a wrapper around gtk_tree_model_iter_next().
func (v *treeModel) IterNext(iter gtk.TreeIter) bool {
	c := C.gtk_tree_model_iter_next(v.native(), castToTreeIter(iter).native())
	return gobool(c)
}

// IterPrevious is a wrapper around gtk_tree_model_iter_previous().
func (v *treeModel) IterPrevious(iter gtk.TreeIter) bool {
	c := C.gtk_tree_model_iter_previous(v.native(), castToTreeIter(iter).native())
	return gobool(c)
}

// IterChildren is a wrapper around gtk_tree_model_iter_children().
func (v *treeModel) IterChildren(iter, child gtk.TreeIter) bool {
	var cIter, cChild *C.GtkTreeIter
	if iter != nil {
		cIter = castToTreeIter(iter).native()
	}
	cChild = castToTreeIter(child).native()
	c := C.gtk_tree_model_iter_children(v.native(), cChild, cIter)
	return gobool(c)
}

// IterNChildren is a wrapper around gtk_tree_model_iter_n_children().
func (v *treeModel) IterNChildren(iter gtk.TreeIter) int {
	var cIter *C.GtkTreeIter
	if iter != nil {
		cIter = castToTreeIter(iter).native()
	}
	c := C.gtk_tree_model_iter_n_children(v.native(), cIter)
	return int(c)
}

/*
 * GtkTreePath
 */

// TreePath is a representation of GTK's GtkTreePath.
type treePath struct {
	GtkTreePath *C.GtkTreePath
}

// Return a TreePath from the GList
func TreePathFromList(list *glib_impl.List) *treePath {
	if list == nil {
		return nil
	}
	return &treePath{(*C.GtkTreePath)(list.Data().(unsafe.Pointer))}
}

// native returns a pointer to the underlying GtkTreePath.
func (v *treePath) native() *C.GtkTreePath {
	if v == nil {
		return nil
	}
	return v.GtkTreePath
}

func marshalTreePath(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &treePath{(*C.GtkTreePath)(unsafe.Pointer(c))}, nil
}

func (v *treePath) free() {
	C.gtk_tree_path_free(v.native())
}

// TreePathNew creates a new TreePath
func TreePathNew() gtk.TreePath {
	return new(treePath)
}

// String is a wrapper around gtk_tree_path_to_string().
func (v *treePath) String() string {
	c := C.gtk_tree_path_to_string(v.native())
	return C.GoString((*C.char)(c))
}

// TreePathNewFromString is a wrapper around gtk_tree_path_new_from_string().
func TreePathNewFromString(path string) (*treePath, error) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_tree_path_new_from_string((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	t := &treePath{c}
	runtime.SetFinalizer(t, (*treePath).free)
	return t, nil
}

/*
 * GtkTreeSelection
 */

// TreeSelection is a representation of GTK's GtkTreeSelection.
type treeSelection struct {
	*glib_impl.Object
}

// native returns a pointer to the underlying GtkTreeSelection.
func (v *treeSelection) native() *C.GtkTreeSelection {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTreeSelection(p)
}

func marshalTreeSelection(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapTreeSelection(obj), nil
}

func wrapTreeSelection(obj *glib_impl.Object) *treeSelection {
	return &treeSelection{obj}
}

// GetSelected() is a wrapper around gtk_tree_selection_get_selected().
func (v *treeSelection) GetSelected() (model gtk.TreeModel, iter gtk.TreeIter, ok bool) {
	var cmodel *C.GtkTreeModel
	var citer C.GtkTreeIter
	c := C.gtk_tree_selection_get_selected(v.native(),
		&cmodel, &citer)
	model = wrapTreeModel(wrapObject(unsafe.Pointer(cmodel)))
	iter = &treeIter{citer}
	ok = gobool(c)
	return
}

// SelectPath is a wrapper around gtk_tree_selection_select_path().
func (v *treeSelection) SelectPath(path gtk.TreePath) {
	C.gtk_tree_selection_select_path(v.native(), castToTreePath(path).native())
}

// UnselectPath is a wrapper around gtk_tree_selection_unselect_path().
func (v *treeSelection) UnselectPath(path gtk.TreePath) {
	C.gtk_tree_selection_unselect_path(v.native(), castToTreePath(path).native())
}

// GetSelectedRows is a wrapper around gtk_tree_selection_get_selected_rows().
// All the elements of returned list are wrapped into (*gtk.TreePath) values.
//
// Please note that a runtime finalizer is only set on the head of the linked
// list, and must be kept live while accessing any item in the list, or the
// Go garbage collector will free the whole list.
func (v *treeSelection) GetSelectedRows(model gtk.TreeModel) glib.List {
	var pcmodel **C.GtkTreeModel
	if model != nil {
		cmodel := model.(ITreeModel).toTreeModel()
		pcmodel = &cmodel
	}

	clist := C.gtk_tree_selection_get_selected_rows(v.native(), pcmodel)
	if clist == nil {
		return nil
	}

	glist := glib_impl.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		return &treePath{(*C.GtkTreePath)(ptr)}
	})
	runtime.SetFinalizer(glist, func(glist *glib_impl.List) {
		glist.FreeFull(func(item interface{}) {
			path := item.(treePath)
			C.gtk_tree_path_free(path.GtkTreePath)
		})
	})

	return glist
}

// CountSelectedRows() is a wrapper around gtk_tree_selection_count_selected_rows().
func (v *treeSelection) CountSelectedRows() int {
	return int(C.gtk_tree_selection_count_selected_rows(v.native()))
}

// SelectIter is a wrapper around gtk_tree_selection_select_iter().
func (v *treeSelection) SelectIter(iter gtk.TreeIter) {
	C.gtk_tree_selection_select_iter(v.native(), castToTreeIter(iter).native())
}

// SetMode() is a wrapper around gtk_tree_selection_set_mode().
func (v *treeSelection) SetMode(m gtk.SelectionMode) {
	C.gtk_tree_selection_set_mode(v.native(), C.GtkSelectionMode(m))
}

// GetMode() is a wrapper around gtk_tree_selection_get_mode().
func (v *treeSelection) GetMode() gtk.SelectionMode {
	return gtk.SelectionMode(C.gtk_tree_selection_get_mode(v.native()))
}

/*
 * GtkTreeStore
 */

// TreeStore is a representation of GTK's GtkTreeStore.
type treeStore struct {
	*glib_impl.Object

	// Interfaces
	treeModel
}

// native returns a pointer to the underlying GtkTreeStore.
func (v *treeStore) native() *C.GtkTreeStore {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTreeStore(p)
}

func marshalTreeStore(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapTreeStore(obj), nil
}

func wrapTreeStore(obj *glib_impl.Object) *treeStore {
	tm := wrapTreeModel(obj)
	return &treeStore{obj, *tm}
}

func (v *treeStore) toTreeModel() *C.GtkTreeModel {
	if v == nil {
		return nil
	}
	return C.toGtkTreeModel(unsafe.Pointer(v.GObject))
}

// TreeStoreNew is a wrapper around gtk_tree_store_newv().
func TreeStoreNew(types ...glib.Type) (*treeStore, error) {
	gtypes := C.alloc_types(C.int(len(types)))
	for n, val := range types {
		C.set_type(gtypes, C.int(n), C.GType(val))
	}
	defer C.g_free(C.gpointer(gtypes))
	c := C.gtk_tree_store_newv(C.gint(len(types)), gtypes)
	if c == nil {
		return nil, nilPtrErr
	}

	ts := wrapTreeStore(wrapObject(unsafe.Pointer(c)))
	return ts, nil
}

// Append is a wrapper around gtk_tree_store_append().
func (v *treeStore) Append(parent gtk.TreeIter) gtk.TreeIter {
	var ti C.GtkTreeIter
	var cParent *C.GtkTreeIter
	if parent != nil {
		cParent = castToTreeIter(parent).native()
	}
	C.gtk_tree_store_append(v.native(), &ti, cParent)
	iter := &treeIter{ti}
	return iter
}

// Insert is a wrapper around gtk_tree_store_insert
func (v *treeStore) Insert(parent gtk.TreeIter, position int) gtk.TreeIter {
	var ti C.GtkTreeIter
	var cParent *C.GtkTreeIter
	if parent != nil {
		cParent = castToTreeIter(parent).native()
	}
	C.gtk_tree_store_insert(v.native(), &ti, cParent, C.gint(position))
	iter := &treeIter{ti}
	return iter
}

// SetValue is a wrapper around gtk_tree_store_set_value()
func (v *treeStore) SetValue(iter gtk.TreeIter, column int, value interface{}) error {
	switch value.(type) {
	case *gdk_impl.Pixbuf:
		pix := value.(*gdk_impl.Pixbuf)
		C._gtk_tree_store_set(v.native(), castToTreeIter(iter).native(), C.gint(column), unsafe.Pointer(pix.Native()))

	default:
		gv, err := glib_impl.GValue(value)
		if err != nil {
			return err
		}
		C.gtk_tree_store_set_value(v.native(), castToTreeIter(iter).native(),
			C.gint(column),
			(*C.GValue)(C.gpointer(gv.Native())))
	}
	return nil
}

// Remove is a wrapper around gtk_tree_store_remove().
func (v *treeStore) Remove(iter gtk.TreeIter) bool {
	var ti *C.GtkTreeIter
	if iter != nil {
		ti = castToTreeIter(iter).native()
	}
	return 0 != C.gtk_tree_store_remove(v.native(), ti)
}

// Clear is a wrapper around gtk_tree_store_clear().
func (v *treeStore) Clear() {
	C.gtk_tree_store_clear(v.native())
}

/*
 * GtkViewport
 */

// Viewport is a representation of GTK's GtkViewport GInterface.
type viewport struct {
	bin

	// Interfaces
	scrollable
}

// IViewport is an interface type implemented by all structs
// embedding a Viewport.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkViewport.
type IViewport interface {
	toViewport() *C.GtkViewport
}

// native() returns a pointer to the underlying GObject as a GtkViewport.
func (v *viewport) native() *C.GtkViewport {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkViewport(p)
}

func wrapViewport(obj *glib_impl.Object) *viewport {
	b := wrapBin(obj)
	s := wrapScrollable(obj)
	return &viewport{
		bin:        *b,
		scrollable: *s,
	}
}

func (v *viewport) toViewport() *C.GtkViewport {
	if v == nil {
		return nil
	}
	return v.native()
}

// ViewportNew() is a wrapper around gtk_viewport_new().
func ViewportNew(hadjustment, vadjustment *adjustment) (*viewport, error) {
	c := C.gtk_viewport_new(hadjustment.native(), vadjustment.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapViewport(wrapObject(unsafe.Pointer(c))), nil
}

func (v *viewport) SetHAdjustment(adjustment gtk.Adjustment) {
	wrapScrollable(v.Object).SetHAdjustment(castToAdjustment(adjustment))
}

func (v *viewport) GetHAdjustment() (gtk.Adjustment, error) {
	return wrapScrollable(v.Object).GetHAdjustment()
}

func (v *viewport) SetVAdjustment(adjustment gtk.Adjustment) {
	wrapScrollable(v.Object).SetVAdjustment(castToAdjustment(adjustment))
}

func (v *viewport) GetVAdjustment() (gtk.Adjustment, error) {
	return wrapScrollable(v.Object).GetVAdjustment()
}

/*
 * GtkVolumeButton
 */

// VolumeButton is a representation of GTK's GtkVolumeButton.
type volumeButton struct {
	scaleButton
}

// native() returns a pointer to the underlying GtkVolumeButton.
func (v *volumeButton) native() *C.GtkVolumeButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkVolumeButton(p)
}

func marshalVolumeButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapVolumeButton(obj), nil
}

func wrapVolumeButton(obj *glib_impl.Object) *volumeButton {
	return &volumeButton{scaleButton{button{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}}}
}

// VolumeButtonNew() is a wrapper around gtk_button_new().
func VolumeButtonNew() (*volumeButton, error) {
	c := C.gtk_volume_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapVolumeButton(wrapObject(unsafe.Pointer(c))), nil
}

type WrapFn interface{}

var WrapMap = map[string]WrapFn{
	"GtkAccelGroup":          wrapAccelGroup,
	"GtkAccelMao":            wrapAccelMap,
	"GtkAdjustment":          wrapAdjustment,
	"GtkApplicationWindow":   wrapApplicationWindow,
	"GtkAssistant":           wrapAssistant,
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
func cast(c *C.GObject) (glib_impl.IObject, error) {
	var (
		className = C.GoString((*C.char)(C.object_get_class_name(c)))
		obj       = wrapObject(unsafe.Pointer(c))
	)

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

	ret, ok := rv[0].Interface().(glib_impl.IObject)
	if !ok {
		return nil, errors.New("did not return an IObject")
	}

	return ret, nil
}
