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

// This file includes wrapers for symbols deprecated beginning with GTK 3.10,
// and should only be included in a build targeted intended to target GTK
// 3.8 or earlier.  To target an earlier build build, use the build tag
// gtk_MAJOR_MINOR.  For example, to target GTK 3.8, run
// 'go build -tags gtk_3_8'.
// +build gtk_3_6 gtk_3_8

package gtkf

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/gtk"
)

// ButtonNewFromStock is a wrapper around gtk_button_new_from_stock().
func ButtonNewFromStock(stock gtk.Stock) (*button, error) {
	cstr := C.CString(string(stock))
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_button_new_from_stock((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapButton(wrapObject(unsafe.Pointer(c))), nil
}

// SetUseStock is a wrapper around gtk_button_set_use_stock().
func (v *button) SetUseStock(useStock bool) {
	C.gtk_button_set_use_stock(v.native(), gbool(useStock))
}

// GetUseStock is a wrapper around gtk_button_get_use_stock().
func (v *button) GetUseStock() bool {
	c := C.gtk_button_get_use_stock(v.native())
	return gobool(c)
}

// GetIconStock is a wrapper around gtk_entry_get_icon_stock().
func (v *entry) GetIconStock(iconPos entryIconPosition) (string, error) {
	c := C.gtk_entry_get_icon_stock(v.native(),
		C.GtkEntryIconPosition(iconPos))
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// SetIconFromStock is a wrapper around gtk_entry_set_icon_from_stock().
func (v *entry) SetIconFromStock(iconPos entryIconPosition, stockID string) {
	cstr := C.CString(stockID)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_set_icon_from_stock(v.native(),
		C.GtkEntryIconPosition(iconPos), (*C.gchar)(cstr))
}

// ImageNewFromStock is a wrapper around gtk_image_new_from_stock().
func ImageNewFromStock(stock gtk.Stock, size iconSize) (*image, error) {
	cstr := C.CString(string(stock))
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_image_new_from_stock((*C.gchar)(cstr), C.GtkIconSize(size))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapImage(wrapObject(unsafe.Pointer(c))), nil
}

// SetFromStock is a wrapper around gtk_image_set_from_stock().
func (v *image) SetFromStock(stock gtk.Stock, size IconSize) {
	cstr := C.CString(string(stock))
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_image_set_from_stock(v.native(), (*C.gchar)(cstr),
		C.GtkIconSize(size))
}

const (
	STOCK_ABOUT                         gtk.Stock = C.GTK_STOCK_ABOUT
	STOCK_ADD                           gtk.Stock = C.GTK_STOCK_ADD
	STOCK_APPLY                         gtk.Stock = C.GTK_STOCK_APPLY
	STOCK_BOLD                          gtk.Stock = C.GTK_STOCK_BOLD
	STOCK_CANCEL                        gtk.Stock = C.GTK_STOCK_CANCEL
	STOCK_CAPS_LOCK_WARNING             gtk.Stock = C.GTK_STOCK_CAPS_LOCK_WARNING
	STOCK_CDROM                         gtk.Stock = C.GTK_STOCK_CDROM
	STOCK_CLEAR                         gtk.Stock = C.GTK_STOCK_CLEAR
	STOCK_CLOSE                         gtk.Stock = C.GTK_STOCK_CLOSE
	STOCK_COLOR_PICKER                  gtk.Stock = C.GTK_STOCK_COLOR_PICKER
	STOCK_CONNECT                       gtk.Stock = C.GTK_STOCK_CONNECT
	STOCK_CONVERT                       gtk.Stock = C.GTK_STOCK_CONVERT
	STOCK_COPY                          gtk.Stock = C.GTK_STOCK_COPY
	STOCK_CUT                           gtk.Stock = C.GTK_STOCK_CUT
	STOCK_DELETE                        gtk.Stock = C.GTK_STOCK_DELETE
	STOCK_DIALOG_AUTHENTICATION         gtk.Stock = C.GTK_STOCK_DIALOG_AUTHENTICATION
	STOCK_DIALOG_INFO                   gtk.Stock = C.GTK_STOCK_DIALOG_INFO
	STOCK_DIALOG_WARNING                gtk.Stock = C.GTK_STOCK_DIALOG_WARNING
	STOCK_DIALOG_ERROR                  gtk.Stock = C.GTK_STOCK_DIALOG_ERROR
	STOCK_DIALOG_QUESTION               gtk.Stock = C.GTK_STOCK_DIALOG_QUESTION
	STOCK_DIRECTORY                     gtk.Stock = C.GTK_STOCK_DIRECTORY
	STOCK_DISCARD                       gtk.Stock = C.GTK_STOCK_DISCARD
	STOCK_DISCONNECT                    gtk.Stock = C.GTK_STOCK_DISCONNECT
	STOCK_DND                           gtk.Stock = C.GTK_STOCK_DND
	STOCK_DND_MULTIPLE                  gtk.Stock = C.GTK_STOCK_DND_MULTIPLE
	STOCK_EDIT                          gtk.Stock = C.GTK_STOCK_EDIT
	STOCK_EXECUTE                       gtk.Stock = C.GTK_STOCK_EXECUTE
	STOCK_FILE                          gtk.Stock = C.GTK_STOCK_FILE
	STOCK_FIND                          gtk.Stock = C.GTK_STOCK_FIND
	STOCK_FIND_AND_REPLACE              gtk.Stock = C.GTK_STOCK_FIND_AND_REPLACE
	STOCK_FLOPPY                        gtk.Stock = C.GTK_STOCK_FLOPPY
	STOCK_FULLSCREEN                    gtk.Stock = C.GTK_STOCK_FULLSCREEN
	STOCK_GOTO_BOTTOM                   gtk.Stock = C.GTK_STOCK_GOTO_BOTTOM
	STOCK_GOTO_FIRST                    gtk.Stock = C.GTK_STOCK_GOTO_FIRST
	STOCK_GOTO_LAST                     gtk.Stock = C.GTK_STOCK_GOTO_LAST
	STOCK_GOTO_TOP                      gtk.Stock = C.GTK_STOCK_GOTO_TOP
	STOCK_GO_BACK                       gtk.Stock = C.GTK_STOCK_GO_BACK
	STOCK_GO_DOWN                       gtk.Stock = C.GTK_STOCK_GO_DOWN
	STOCK_GO_FORWARD                    gtk.Stock = C.GTK_STOCK_GO_FORWARD
	STOCK_GO_UP                         gtk.Stock = C.GTK_STOCK_GO_UP
	STOCK_HARDDISK                      gtk.Stock = C.GTK_STOCK_HARDDISK
	STOCK_HELP                          gtk.Stock = C.GTK_STOCK_HELP
	STOCK_HOME                          gtk.Stock = C.GTK_STOCK_HOME
	STOCK_INDEX                         gtk.Stock = C.GTK_STOCK_INDEX
	STOCK_INDENT                        gtk.Stock = C.GTK_STOCK_INDENT
	STOCK_INFO                          gtk.Stock = C.GTK_STOCK_INFO
	STOCK_ITALIC                        gtk.Stock = C.GTK_STOCK_ITALIC
	STOCK_JUMP_TO                       gtk.Stock = C.GTK_STOCK_JUMP_TO
	STOCK_JUSTIFY_CENTER                gtk.Stock = C.GTK_STOCK_JUSTIFY_CENTER
	STOCK_JUSTIFY_FILL                  gtk.Stock = C.GTK_STOCK_JUSTIFY_FILL
	STOCK_JUSTIFY_LEFT                  gtk.Stock = C.GTK_STOCK_JUSTIFY_LEFT
	STOCK_JUSTIFY_RIGHT                 gtk.Stock = C.GTK_STOCK_JUSTIFY_RIGHT
	STOCK_LEAVE_FULLSCREEN              gtk.Stock = C.GTK_STOCK_LEAVE_FULLSCREEN
	STOCK_MISSING_IMAGE                 gtk.Stock = C.GTK_STOCK_MISSING_IMAGE
	STOCK_MEDIA_FORWARD                 gtk.Stock = C.GTK_STOCK_MEDIA_FORWARD
	STOCK_MEDIA_NEXT                    gtk.Stock = C.GTK_STOCK_MEDIA_NEXT
	STOCK_MEDIA_PAUSE                   gtk.Stock = C.GTK_STOCK_MEDIA_PAUSE
	STOCK_MEDIA_PLAY                    gtk.Stock = C.GTK_STOCK_MEDIA_PLAY
	STOCK_MEDIA_PREVIOUS                gtk.Stock = C.GTK_STOCK_MEDIA_PREVIOUS
	STOCK_MEDIA_RECORD                  gtk.Stock = C.GTK_STOCK_MEDIA_RECORD
	STOCK_MEDIA_REWIND                  gtk.Stock = C.GTK_STOCK_MEDIA_REWIND
	STOCK_MEDIA_STOP                    gtk.Stock = C.GTK_STOCK_MEDIA_STOP
	STOCK_NETWORK                       gtk.Stock = C.GTK_STOCK_NETWORK
	STOCK_NEW                           gtk.Stock = C.GTK_STOCK_NEW
	STOCK_NO                            gtk.Stock = C.GTK_STOCK_NO
	STOCK_OK                            gtk.Stock = C.GTK_STOCK_OK
	STOCK_OPEN                          gtk.Stock = C.GTK_STOCK_OPEN
	STOCK_ORIENTATION_PORTRAIT          gtk.Stock = C.GTK_STOCK_ORIENTATION_PORTRAIT
	STOCK_ORIENTATION_LANDSCAPE         gtk.Stock = C.GTK_STOCK_ORIENTATION_LANDSCAPE
	STOCK_ORIENTATION_REVERSE_LANDSCAPE gtk.Stock = C.GTK_STOCK_ORIENTATION_REVERSE_LANDSCAPE
	STOCK_ORIENTATION_REVERSE_PORTRAIT  gtk.Stock = C.GTK_STOCK_ORIENTATION_REVERSE_PORTRAIT
	STOCK_PAGE_SETUP                    gtk.Stock = C.GTK_STOCK_PAGE_SETUP
	STOCK_PASTE                         gtk.Stock = C.GTK_STOCK_PASTE
	STOCK_PREFERENCES                   gtk.Stock = C.GTK_STOCK_PREFERENCES
	STOCK_PRINT                         gtk.Stock = C.GTK_STOCK_PRINT
	STOCK_PRINT_ERROR                   gtk.Stock = C.GTK_STOCK_PRINT_ERROR
	STOCK_PRINT_PAUSED                  gtk.Stock = C.GTK_STOCK_PRINT_PAUSED
	STOCK_PRINT_PREVIEW                 gtk.Stock = C.GTK_STOCK_PRINT_PREVIEW
	STOCK_PRINT_REPORT                  gtk.Stock = C.GTK_STOCK_PRINT_REPORT
	STOCK_PRINT_WARNING                 gtk.Stock = C.GTK_STOCK_PRINT_WARNING
	STOCK_PROPERTIES                    gtk.Stock = C.GTK_STOCK_PROPERTIES
	STOCK_QUIT                          gtk.Stock = C.GTK_STOCK_QUIT
	STOCK_REDO                          gtk.Stock = C.GTK_STOCK_REDO
	STOCK_REFRESH                       gtk.Stock = C.GTK_STOCK_REFRESH
	STOCK_REMOVE                        gtk.Stock = C.GTK_STOCK_REMOVE
	STOCK_REVERT_TO_SAVED               gtk.Stock = C.GTK_STOCK_REVERT_TO_SAVED
	STOCK_SAVE                          gtk.Stock = C.GTK_STOCK_SAVE
	STOCK_SAVE_AS                       gtk.Stock = C.GTK_STOCK_SAVE_AS
	STOCK_SELECT_ALL                    gtk.Stock = C.GTK_STOCK_SELECT_ALL
	STOCK_SELECT_COLOR                  gtk.Stock = C.GTK_STOCK_SELECT_COLOR
	STOCK_SELECT_FONT                   gtk.Stock = C.GTK_STOCK_SELECT_FONT
	STOCK_SORT_ASCENDING                gtk.Stock = C.GTK_STOCK_SORT_ASCENDING
	STOCK_SORT_DESCENDING               gtk.Stock = C.GTK_STOCK_SORT_DESCENDING
	STOCK_SPELL_CHECK                   gtk.Stock = C.GTK_STOCK_SPELL_CHECK
	STOCK_STOP                          gtk.Stock = C.GTK_STOCK_STOP
	STOCK_STRIKETHROUGH                 gtk.Stock = C.GTK_STOCK_STRIKETHROUGH
	STOCK_UNDELETE                      gtk.Stock = C.GTK_STOCK_UNDELETE
	STOCK_UNDERLINE                     gtk.Stock = C.GTK_STOCK_UNDERLINE
	STOCK_UNDO                          gtk.Stock = C.GTK_STOCK_UNDO
	STOCK_UNINDENT                      gtk.Stock = C.GTK_STOCK_UNINDENT
	STOCK_YES                           gtk.Stock = C.GTK_STOCK_YES
	STOCK_ZOOM_100                      gtk.Stock = C.GTK_STOCK_ZOOM_100
	STOCK_ZOOM_FIT                      gtk.Stock = C.GTK_STOCK_ZOOM_FIT
	STOCK_ZOOM_IN                       gtk.Stock = C.GTK_STOCK_ZOOM_IN
	STOCK_ZOOM_OUT                      gtk.Stock = C.GTK_STOCK_ZOOM_OUT
)

// ReshowWithInitialSize is a wrapper around
// gtk_window_reshow_with_initial_size().
func (v *window) ReshowWithInitialSize() {
	C.gtk_window_reshow_with_initial_size(v.native())
}
