package gtkf

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	gdk_impl "github.com/gotk3/gotk3/gdkf"
	"github.com/gotk3/gotk3/glib"
	glib_impl "github.com/gotk3/gotk3/glibf"
	"github.com/gotk3/gotk3/gtk"
)

func init() {
	tm := []glib_impl.TypeMarshaler{
		{glib.Type(C.gtk_about_dialog_get_type()), marshalAboutDialog},
	}

	glib_impl.RegisterGValueMarshalers(tm)

	WrapMap["GtkAboutDialog"] = wrapAboutDialog
}

/*
 * GtkAboutDialog
 */

// AboutDialog is a representation of GTK's GtkAboutDialog.
type aboutDialog struct {
	dialog
}

// native returns a pointer to the underlying GtkAboutDialog.
func (v *aboutDialog) native() *C.GtkAboutDialog {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkAboutDialog(p)
}

func marshalAboutDialog(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAboutDialog(obj), nil
}

func wrapAboutDialog(obj *glib_impl.Object) *aboutDialog {
	return &aboutDialog{dialog{window{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}}}
}

// AboutDialogNew is a wrapper around gtk_about_dialog_new().
func AboutDialogNew() (*aboutDialog, error) {
	c := C.gtk_about_dialog_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAboutDialog(obj), nil
}

// GetComments is a wrapper around gtk_about_dialog_get_comments().
func (v *aboutDialog) GetComments() string {
	c := C.gtk_about_dialog_get_comments(v.native())
	return C.GoString((*C.char)(c))
}

// SetComments is a wrapper around gtk_about_dialog_set_comments().
func (v *aboutDialog) SetComments(comments string) {
	cstr := C.CString(comments)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_comments(v.native(), (*C.gchar)(cstr))
}

// GetCopyright is a wrapper around gtk_about_dialog_get_copyright().
func (v *aboutDialog) GetCopyright() string {
	c := C.gtk_about_dialog_get_copyright(v.native())
	return C.GoString((*C.char)(c))
}

// SetCopyright is a wrapper around gtk_about_dialog_set_copyright().
func (v *aboutDialog) SetCopyright(copyright string) {
	cstr := C.CString(copyright)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_copyright(v.native(), (*C.gchar)(cstr))
}

// GetLicense is a wrapper around gtk_about_dialog_get_license().
func (v *aboutDialog) GetLicense() string {
	c := C.gtk_about_dialog_get_license(v.native())
	return C.GoString((*C.char)(c))
}

// SetLicense is a wrapper around gtk_about_dialog_set_license().
func (v *aboutDialog) SetLicense(license string) {
	cstr := C.CString(license)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_license(v.native(), (*C.gchar)(cstr))
}

// GetLicenseType is a wrapper around gtk_about_dialog_get_license_type().
func (v *aboutDialog) GetLicenseType() gtk.License {
	c := C.gtk_about_dialog_get_license_type(v.native())
	return gtk.License(c)
}

// SetLicenseType is a wrapper around gtk_about_dialog_set_license_type().
func (v *aboutDialog) SetLicenseType(license gtk.License) {
	C.gtk_about_dialog_set_license_type(v.native(), C.GtkLicense(license))
}

// SetLogo is a wrapper around gtk_about_dialog_set_logo().
func (v *aboutDialog) SetLogo(logo gdk.Pixbuf) {
	logoPtr := (*C.GdkPixbuf)(unsafe.Pointer(gdk_impl.CastToPixbuf(logo).Native()))
	C.gtk_about_dialog_set_logo(v.native(), logoPtr)
}

// GetLogoIconName is a wrapper around gtk_about_dialog_get_logo_icon_name().
func (v *aboutDialog) GetLogoIconName() string {
	c := C.gtk_about_dialog_get_logo_icon_name(v.native())
	return C.GoString((*C.char)(c))
}

// SetLogoIconName is a wrapper around gtk_about_dialog_set_logo_icon_name().
func (v *aboutDialog) SetLogoIconName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_logo_icon_name(v.native(), (*C.gchar)(cstr))
}

// GetProgramName is a wrapper around gtk_about_dialog_get_program_name().
func (v *aboutDialog) GetProgramName() string {
	c := C.gtk_about_dialog_get_program_name(v.native())
	return C.GoString((*C.char)(c))
}

// SetProgramName is a wrapper around gtk_about_dialog_set_program_name().
func (v *aboutDialog) SetProgramName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_program_name(v.native(), (*C.gchar)(cstr))
}

// GetAuthors is a wrapper around gtk_about_dialog_get_authors().
func (v *aboutDialog) GetAuthors() []string {
	var authors []string
	cauthors := C.gtk_about_dialog_get_authors(v.native())
	if cauthors == nil {
		return nil
	}
	for {
		if *cauthors == nil {
			break
		}
		authors = append(authors, C.GoString((*C.char)(*cauthors)))
		cauthors = C.next_gcharptr(cauthors)
	}
	return authors
}

// SetAuthors is a wrapper around gtk_about_dialog_set_authors().
func (v *aboutDialog) SetAuthors(authors []string) {
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

// GetArtists is a wrapper around gtk_about_dialog_get_artists().
func (v *aboutDialog) GetArtists() []string {
	var artists []string
	cartists := C.gtk_about_dialog_get_artists(v.native())
	if cartists == nil {
		return nil
	}
	for {
		if *cartists == nil {
			break
		}
		artists = append(artists, C.GoString((*C.char)(*cartists)))
		cartists = C.next_gcharptr(cartists)
	}
	return artists
}

// SetArtists is a wrapper around gtk_about_dialog_set_artists().
func (v *aboutDialog) SetArtists(artists []string) {
	cartists := C.make_strings(C.int(len(artists) + 1))
	for i, artist := range artists {
		cstr := C.CString(artist)
		defer C.free(unsafe.Pointer(cstr))
		C.set_string(cartists, C.int(i), (*C.gchar)(cstr))
	}

	C.set_string(cartists, C.int(len(artists)), nil)
	C.gtk_about_dialog_set_artists(v.native(), cartists)
	C.destroy_strings(cartists)
}

// GetDocumenters is a wrapper around gtk_about_dialog_get_documenters().
func (v *aboutDialog) GetDocumenters() []string {
	var documenters []string
	cdocumenters := C.gtk_about_dialog_get_documenters(v.native())
	if cdocumenters == nil {
		return nil
	}
	for {
		if *cdocumenters == nil {
			break
		}
		documenters = append(documenters, C.GoString((*C.char)(*cdocumenters)))
		cdocumenters = C.next_gcharptr(cdocumenters)
	}
	return documenters
}

// SetDocumenters is a wrapper around gtk_about_dialog_set_documenters().
func (v *aboutDialog) SetDocumenters(documenters []string) {
	cdocumenters := C.make_strings(C.int(len(documenters) + 1))
	for i, doc := range documenters {
		cstr := C.CString(doc)
		defer C.free(unsafe.Pointer(cstr))
		C.set_string(cdocumenters, C.int(i), (*C.gchar)(cstr))
	}

	C.set_string(cdocumenters, C.int(len(documenters)), nil)
	C.gtk_about_dialog_set_documenters(v.native(), cdocumenters)
	C.destroy_strings(cdocumenters)
}

// GetTranslatorCredits is a wrapper around gtk_about_dialog_get_translator_credits().
func (v *aboutDialog) GetTranslatorCredits() string {
	c := C.gtk_about_dialog_get_translator_credits(v.native())
	return C.GoString((*C.char)(c))
}

// SetTranslatorCredits is a wrapper around gtk_about_dialog_set_translator_credits().
func (v *aboutDialog) SetTranslatorCredits(translatorCredits string) {
	cstr := C.CString(translatorCredits)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_translator_credits(v.native(), (*C.gchar)(cstr))
}

// GetVersion is a wrapper around gtk_about_dialog_get_version().
func (v *aboutDialog) GetVersion() string {
	c := C.gtk_about_dialog_get_version(v.native())
	return C.GoString((*C.char)(c))
}

// SetVersion is a wrapper around gtk_about_dialog_set_version().
func (v *aboutDialog) SetVersion(version string) {
	cstr := C.CString(version)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_version(v.native(), (*C.gchar)(cstr))
}

// GetWebsite is a wrapper around gtk_about_dialog_get_website().
func (v *aboutDialog) GetWebsite() string {
	c := C.gtk_about_dialog_get_website(v.native())
	return C.GoString((*C.char)(c))
}

// SetWebsite is a wrapper around gtk_about_dialog_set_website().
func (v *aboutDialog) SetWebsite(website string) {
	cstr := C.CString(website)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_website(v.native(), (*C.gchar)(cstr))
}

// GetWebsiteLabel is a wrapper around gtk_about_dialog_get_website_label().
func (v *aboutDialog) GetWebsiteLabel() string {
	c := C.gtk_about_dialog_get_website_label(v.native())
	return C.GoString((*C.char)(c))
}

// SetWebsiteLabel is a wrapper around gtk_about_dialog_set_website_label().
func (v *aboutDialog) SetWebsiteLabel(websiteLabel string) {
	cstr := C.CString(websiteLabel)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_website_label(v.native(), (*C.gchar)(cstr))
}

// GetWrapLicense is a wrapper around gtk_about_dialog_get_wrap_license().
func (v *aboutDialog) GetWrapLicense() bool {
	return gobool(C.gtk_about_dialog_get_wrap_license(v.native()))
}

// SetWrapLicense is a wrapper around gtk_about_dialog_set_wrap_license().
func (v *aboutDialog) SetWrapLicense(wrapLicense bool) {
	C.gtk_about_dialog_set_wrap_license(v.native(), gbool(wrapLicense))
}

// AddCreditSection is a wrapper around gtk_about_dialog_add_credit_section().
func (v *aboutDialog) AddCreditSection(sectionName string, people []string) {
	cname := (*C.gchar)(C.CString(sectionName))
	defer C.free(unsafe.Pointer(cname))

	cpeople := C.make_strings(C.int(len(people)) + 1)
	defer C.destroy_strings(cpeople)
	for i, p := range people {
		cp := (*C.gchar)(C.CString(p))
		defer C.free(unsafe.Pointer(cp))
		C.set_string(cpeople, C.int(i), cp)
	}
	C.set_string(cpeople, C.int(len(people)), nil)

	C.gtk_about_dialog_add_credit_section(v.native(), cname, cpeople)
}
