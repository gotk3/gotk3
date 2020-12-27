// Same copyright and license as the rest of the files in this project

package gdk

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
// #include "pixbuf.go.h"
import "C"
import (
	"errors"
	"reflect"
	"runtime"
	"strconv"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.gdk_pixbuf_alpha_mode_get_type()), marshalPixbufAlphaMode},

		// Objects/Interfaces
		{glib.Type(C.gdk_pixbuf_get_type()), marshalPixbuf},
		{glib.Type(C.gdk_pixbuf_animation_get_type()), marshalPixbufAnimation},
		{glib.Type(C.gdk_pixbuf_loader_get_type()), marshalPixbufLoader},
	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * Constants
 */

// TODO:
// GdkPixbufError

// PixbufRotation is a representation of GDK's GdkPixbufRotation.
type PixbufRotation int

const (
	PIXBUF_ROTATE_NONE             PixbufRotation = C.GDK_PIXBUF_ROTATE_NONE
	PIXBUF_ROTATE_COUNTERCLOCKWISE PixbufRotation = C.GDK_PIXBUF_ROTATE_COUNTERCLOCKWISE
	PIXBUF_ROTATE_UPSIDEDOWN       PixbufRotation = C.GDK_PIXBUF_ROTATE_UPSIDEDOWN
	PIXBUF_ROTATE_CLOCKWISE        PixbufRotation = C.GDK_PIXBUF_ROTATE_CLOCKWISE
)

// PixbufAlphaMode is a representation of GDK's GdkPixbufAlphaMode.
type PixbufAlphaMode int

const (
	GDK_PIXBUF_ALPHA_BILEVEL PixbufAlphaMode = C.GDK_PIXBUF_ALPHA_BILEVEL
	GDK_PIXBUF_ALPHA_FULL    PixbufAlphaMode = C.GDK_PIXBUF_ALPHA_FULL
)

func marshalPixbufAlphaMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PixbufAlphaMode(c), nil
}

/*
 * GdkPixbuf
 */

// PixbufGetType is a wrapper around gdk_pixbuf_get_type().
func PixbufGetType() glib.Type {
	return glib.Type(C.gdk_pixbuf_get_type())
}

// Pixbuf is a representation of GDK's GdkPixbuf.
type Pixbuf struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkPixbuf.
func (v *Pixbuf) native() *C.GdkPixbuf {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkPixbuf(p)
}

// Native returns a pointer to the underlying GdkPixbuf.
func (v *Pixbuf) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *Pixbuf) NativePrivate() *C.GdkPixbuf {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkPixbuf(p)
}

func marshalPixbuf(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &Pixbuf{obj}, nil
}

// PixbufNew is a wrapper around gdk_pixbuf_new().
func PixbufNew(colorspace Colorspace, hasAlpha bool, bitsPerSample, width, height int) (*Pixbuf, error) {
	c := C.gdk_pixbuf_new(C.GdkColorspace(colorspace), gbool(hasAlpha),
		C.int(bitsPerSample), C.int(width), C.int(height))
	if c == nil {
		return nil, nilPtrErr
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	//obj.Ref()
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })
	return p, nil
}

// File Loading

// PixbufNewFromFile is a wrapper around gdk_pixbuf_new_from_file().
func PixbufNewFromFile(filename string) (*Pixbuf, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError
	c := C.gdk_pixbuf_new_from_file((*C.char)(cstr), &err)
	if c == nil {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	//obj.Ref()
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })
	return p, nil
}

// Image Data in Memory

// PixbufNewFromData is a wrapper around gdk_pixbuf_new_from_data().
func PixbufNewFromData(pixbufData []byte, cs Colorspace, hasAlpha bool, bitsPerSample, width, height, rowStride int) (*Pixbuf, error) {
	arrayPtr := (*C.guchar)(unsafe.Pointer(&pixbufData[0]))

	c := C.gdk_pixbuf_new_from_data(
		arrayPtr,
		C.GdkColorspace(cs),
		gbool(hasAlpha),
		C.int(bitsPerSample),
		C.int(width),
		C.int(height),
		C.int(rowStride),
		nil, // TODO: missing support for GdkPixbufDestroyNotify
		nil,
	)

	if c == nil {
		return nil, nilPtrErr
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	//obj.Ref()
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })

	return p, nil
}

// PixbufNewFromDataOnly is a convenient alternative to PixbufNewFromData() and also a wrapper around gdk_pixbuf_new_from_data().
func PixbufNewFromDataOnly(pixbufData []byte) (*Pixbuf, error) {
	pixbufLoader, err := PixbufLoaderNew()
	if err == nil {
		return nil, err
	}
	return pixbufLoader.WriteAndReturnPixbuf(pixbufData)
}

// TODO:
// gdk_pixbuf_new_from_xpm_data().
// gdk_pixbuf_new_subpixbuf()

// PixbufCopy is a wrapper around gdk_pixbuf_copy().
func PixbufCopy(v *Pixbuf) (*Pixbuf, error) {
	c := C.gdk_pixbuf_copy(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	//obj.Ref()
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })
	return p, nil
}

// The GdkPixbuf Structure

// GetColorspace is a wrapper around gdk_pixbuf_get_colorspace().
func (v *Pixbuf) GetColorspace() Colorspace {
	c := C.gdk_pixbuf_get_colorspace(v.native())
	return Colorspace(c)
}

// GetNChannels is a wrapper around gdk_pixbuf_get_n_channels().
func (v *Pixbuf) GetNChannels() int {
	c := C.gdk_pixbuf_get_n_channels(v.native())
	return int(c)
}

// GetHasAlpha is a wrapper around gdk_pixbuf_get_has_alpha().
func (v *Pixbuf) GetHasAlpha() bool {
	c := C.gdk_pixbuf_get_has_alpha(v.native())
	return gobool(c)
}

// GetBitsPerSample is a wrapper around gdk_pixbuf_get_bits_per_sample().
func (v *Pixbuf) GetBitsPerSample() int {
	c := C.gdk_pixbuf_get_bits_per_sample(v.native())
	return int(c)
}

// GetPixels is a wrapper around gdk_pixbuf_get_pixels_with_length().
// A Go slice is used to represent the underlying Pixbuf data array, one
// byte per channel.
func (v *Pixbuf) GetPixels() (channels []byte) {
	var length C.guint
	c := C.gdk_pixbuf_get_pixels_with_length(v.native(), &length)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&channels))
	sliceHeader.Data = uintptr(unsafe.Pointer(c))
	sliceHeader.Len = int(length)
	sliceHeader.Cap = int(length)

	// To make sure the slice doesn't outlive the Pixbuf, add a reference
	v.Ref()
	runtime.SetFinalizer(&channels, func(_ *[]byte) {
		v.Unref()
	})
	return
}

// GetWidth is a wrapper around gdk_pixbuf_get_width().
func (v *Pixbuf) GetWidth() int {
	c := C.gdk_pixbuf_get_width(v.native())
	return int(c)
}

// GetHeight is a wrapper around gdk_pixbuf_get_height().
func (v *Pixbuf) GetHeight() int {
	c := C.gdk_pixbuf_get_height(v.native())
	return int(c)
}

// GetRowstride is a wrapper around gdk_pixbuf_get_rowstride().
func (v *Pixbuf) GetRowstride() int {
	c := C.gdk_pixbuf_get_rowstride(v.native())
	return int(c)
}

// GetOption is a wrapper around gdk_pixbuf_get_option().  ok is true if
// the key has an associated value.
func (v *Pixbuf) GetOption(key string) (value string, ok bool) {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_pixbuf_get_option(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return "", false
	}
	return C.GoString((*C.char)(c)), true
}

// Scaling

// ScaleSimple is a wrapper around gdk_pixbuf_scale_simple().
func (v *Pixbuf) ScaleSimple(destWidth, destHeight int, interpType InterpType) (*Pixbuf, error) {
	c := C.gdk_pixbuf_scale_simple(v.native(), C.int(destWidth),
		C.int(destHeight), C.GdkInterpType(interpType))
	if c == nil {
		return nil, nilPtrErr
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	//obj.Ref()
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })
	return p, nil
}

// Scale is a wrapper around gdk_pixbuf_scale().
func (v *Pixbuf) Scale(dest *Pixbuf, destX, destY, destWidth, destHeight int, offsetX, offsetY, scaleX, scaleY float64, interpType InterpType) {
	C.gdk_pixbuf_scale(
		v.native(),
		dest.native(),
		C.int(destX),
		C.int(destY),
		C.int(destWidth),
		C.int(destHeight),
		C.double(offsetX),
		C.double(offsetY),
		C.double(scaleX),
		C.double(scaleY),
		C.GdkInterpType(interpType),
	)
}

// Composite is a wrapper around gdk_pixbuf_composite().
func (v *Pixbuf) Composite(dest *Pixbuf, destX, destY, destWidth, destHeight int, offsetX, offsetY, scaleX, scaleY float64, interpType InterpType, overallAlpha int) {
	C.gdk_pixbuf_composite(
		v.native(),
		dest.native(),
		C.int(destX),
		C.int(destY),
		C.int(destWidth),
		C.int(destHeight),
		C.double(offsetX),
		C.double(offsetY),
		C.double(scaleX),
		C.double(scaleY),
		C.GdkInterpType(interpType),
		C.int(overallAlpha),
	)
}

// TODO:
// gdk_pixbuf_composite_color_simple().
// gdk_pixbuf_composite_color().

// Utilities

// TODO:
// gdk_pixbuf_copy_area().
// gdk_pixbuf_saturate_and_pixelate().

// Fill is a wrapper around gdk_pixbuf_fill(). The given pixel is an RGBA value.
func (v *Pixbuf) Fill(pixel uint32) {
	C.gdk_pixbuf_fill(v.native(), C.guint32(pixel))
}

// AddAlpha is a wrapper around gdk_pixbuf_add_alpha(). If substituteColor is
// true, then the color specified by r, g and b will be assigned zero opacity.
func (v *Pixbuf) AddAlpha(substituteColor bool, r, g, b uint8) *Pixbuf {
	c := C.gdk_pixbuf_add_alpha(v.native(), gbool(substituteColor), C.guchar(r), C.guchar(g), C.guchar(b))

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })

	return p
}

// File saving

// TODO:
// gdk_pixbuf_savev().
// gdk_pixbuf_save().

// SaveJPEG is a convenience wrapper around gdk_pixbuf_save() for saving image as jpeg file.
// Quality is a number between 0...100
func (v *Pixbuf) SaveJPEG(path string, quality int) error {
	cpath := C.CString(path)
	cquality := C.CString(strconv.Itoa(quality))
	defer C.free(unsafe.Pointer(cpath))
	defer C.free(unsafe.Pointer(cquality))

	var err *C.GError
	c := C._gdk_pixbuf_save_jpeg(v.native(), cpath, &err, cquality)
	if !gobool(c) {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}

	return nil
}

// SavePNG is a convenience wrapper around gdk_pixbuf_save() for saving image as png file.
// Compression is a number between 0...9
func (v *Pixbuf) SavePNG(path string, compression int) error {
	cpath := C.CString(path)
	ccompression := C.CString(strconv.Itoa(compression))
	defer C.free(unsafe.Pointer(cpath))
	defer C.free(unsafe.Pointer(ccompression))

	var err *C.GError
	c := C._gdk_pixbuf_save_png(v.native(), cpath, &err, ccompression)
	if !gobool(c) {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

/*
 * PixbufFormat
 */

type PixbufFormat struct {
	format *C.GdkPixbufFormat
}

// native returns a pointer to the underlying GdkPixbuf.
func (v *PixbufFormat) native() *C.GdkPixbufFormat {
	if v == nil {
		return nil
	}

	return v.format
}

func wrapPixbufFormat(obj *C.GdkPixbufFormat) *PixbufFormat {
	return &PixbufFormat{obj}
}

// Native returns a pointer to the underlying GdkPixbuf.
func (v *PixbufFormat) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

/*
 * GdkPixbufAnimation
 */

// PixbufAnimation is a representation of GDK's GdkPixbufAnimation.
type PixbufAnimation struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkPixbufAnimation.
func (v *PixbufAnimation) native() *C.GdkPixbufAnimation {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkPixbufAnimation(p)
}

func (v *PixbufAnimation) NativePrivate() *C.GdkPixbufAnimation {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkPixbufAnimation(p)
}

func (v *PixbufAnimation) GetStaticImage() *Pixbuf {
	c := C.gdk_pixbuf_animation_get_static_image(v.native())
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}

	// Add a reference so the pixbuf doesn't outlive the parent pixbuf
	// animation.
	v.Ref()
	runtime.SetFinalizer(p, func(*Pixbuf) { v.Unref() })

	return p
}

func marshalPixbufAnimation(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &PixbufAnimation{obj}, nil
}

// PixbufAnimationNewFromFile is a wrapper around gdk_pixbuf_animation_new_from_file().
func PixbufAnimationNewFromFile(filename string) (*PixbufAnimation, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError
	c := C.gdk_pixbuf_animation_new_from_file((*C.char)(cstr), &err)
	if c == nil {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &PixbufAnimation{obj}
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })
	return p, nil
}

// TODO:
// gdk_pixbuf_animation_new_from_resource().

/*
 * GdkPixbufLoader
 */

// PixbufLoader is a representation of GDK's GdkPixbufLoader.
// Users of PixbufLoader are expected to call Close() when they are finished.
type PixbufLoader struct {
	*glib.Object
}

// native() returns a pointer to the underlying GdkPixbufLoader.
func (v *PixbufLoader) native() *C.GdkPixbufLoader {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkPixbufLoader(p)
}

func marshalPixbufLoader(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &PixbufLoader{obj}, nil
}

// PixbufLoaderNew() is a wrapper around gdk_pixbuf_loader_new().
func PixbufLoaderNew() (*PixbufLoader, error) {
	c := C.gdk_pixbuf_loader_new()
	if c == nil {
		return nil, nilPtrErr
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &PixbufLoader{obj}
	obj.Ref()
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })
	return p, nil
}

// PixbufLoaderNewWithType() is a wrapper around gdk_pixbuf_loader_new_with_type().
func PixbufLoaderNewWithType(t string) (*PixbufLoader, error) {
	var err *C.GError

	cstr := C.CString(t)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gdk_pixbuf_loader_new_with_type((*C.char)(cstr), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}

	if c == nil {
		return nil, nilPtrErr
	}

	return &PixbufLoader{glib.Take(unsafe.Pointer(c))}, nil
}

// Write() is a wrapper around gdk_pixbuf_loader_write().  The
// function signature differs from the C equivalent to satisify the
// io.Writer interface.
func (v *PixbufLoader) Write(data []byte) (int, error) {
	// n is set to 0 on error, and set to len(data) otherwise.
	// This is a tiny hacky to satisfy io.Writer and io.WriteCloser,
	// which would allow access to all io and ioutil goodies,
	// and play along nice with go environment.

	if len(data) == 0 {
		return 0, nil
	}

	var err *C.GError
	c := C.gdk_pixbuf_loader_write(v.native(),
		(*C.guchar)(unsafe.Pointer(&data[0])), C.gsize(len(data)),
		&err)

	if !gobool(c) {
		defer C.g_error_free(err)
		return 0, errors.New(C.GoString((*C.char)(err.message)))
	}

	return len(data), nil
}

func (v *PixbufLoader) WriteAndReturnPixbuf(data []byte) (*Pixbuf, error) {

	if len(data) == 0 {
		return nil, errors.New("no data")
	}

	var err *C.GError
	c := C.gdk_pixbuf_loader_write(v.native(), (*C.guchar)(unsafe.Pointer(&data[0])), C.gsize(len(data)), &err)

	if !gobool(c) {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}

	v.Close()

	c2 := C.gdk_pixbuf_loader_get_pixbuf(v.native())
	if c2 == nil {
		return nil, nilPtrErr
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c2))}
	p := &Pixbuf{obj}
	//obj.Ref() // Don't call Ref here, gdk_pixbuf_loader_get_pixbuf already did that for us.
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })

	return p, nil
}

// Close is a wrapper around gdk_pixbuf_loader_close().  An error is
// returned instead of a bool like the native C function to support the
// io.Closer interface.
func (v *PixbufLoader) Close() error {
	var err *C.GError

	if ok := gobool(C.gdk_pixbuf_loader_close(v.native(), &err)); !ok {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

// GetPixbuf is a wrapper around gdk_pixbuf_loader_get_pixbuf().
func (v *PixbufLoader) GetPixbuf() (*Pixbuf, error) {
	c := C.gdk_pixbuf_loader_get_pixbuf(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	//obj.Ref() // Don't call Ref here, gdk_pixbuf_loader_get_pixbuf already did that for us.
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })
	return p, nil
}

// GetAnimation is a wrapper around gdk_pixbuf_loader_get_animation().
func (v *PixbufLoader) GetAnimation() (*PixbufAnimation, error) {
	c := C.gdk_pixbuf_loader_get_animation(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &PixbufAnimation{obj}
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })
	return p, nil
}

// Convenient function like above for Pixbuf. Write data, close loader and return PixbufAnimation.
func (v *PixbufLoader) WriteAndReturnPixbufAnimation(data []byte) (*PixbufAnimation, error) {

	if len(data) == 0 {
		return nil, errors.New("no data")
	}

	var err *C.GError
	c := C.gdk_pixbuf_loader_write(v.native(), (*C.guchar)(unsafe.Pointer(&data[0])), C.gsize(len(data)), &err)

	if !gobool(c) {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}

	v.Close()

	c2 := C.gdk_pixbuf_loader_get_animation(v.native())
	if c2 == nil {
		return nil, nilPtrErr
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c2))}
	p := &PixbufAnimation{obj}
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })

	return p, nil
}
