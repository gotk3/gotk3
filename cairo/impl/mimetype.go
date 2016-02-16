package impl

import "github.com/gotk3/gotk3/cairo"

func init() {
	cairo.MIME_TYPE_JP2 = "image/jp2"
	cairo.MIME_TYPE_JPEG = "image/jpeg"
	cairo.MIME_TYPE_PNG = "image/png"
	cairo.MIME_TYPE_URI = "image/x-uri"
	cairo.MIME_TYPE_UNIQUE_ID = "application/x-cairo.uuid"
}
