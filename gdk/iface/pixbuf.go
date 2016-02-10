package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type Pixbuf interface {
    glib_iface.Object

    ApplyEmbeddedOrientation() (Pixbuf, error)
    Flip(bool) (Pixbuf, error)
    GetBitsPerSample() int
    GetByteLength() int
    GetColorspace() Colorspace
    GetHasAlpha() bool
    GetHeight() int
    GetNChannels() int
    GetOption(string) (string, bool)
    GetPixels() []byte
    GetRowstride() int
    GetWidth() int
    RotateSimple(PixbufRotation) (Pixbuf, error)
    SaveJPEG(string, int) error
    SavePNG(string, int) error
    ScaleSimple(int, int, InterpType) (Pixbuf, error)
} // end of Pixbuf

func AssertPixbuf(_ Pixbuf) {}
