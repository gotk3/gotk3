package cairo

// Antialias is a representation of Cairo's cairo_antialias_t.
type Antialias int

// Content is a representation of Cairo's cairo_content_t.
type Content int

// FillRule is a representation of Cairo's cairo_fill_rule_t.
type FillRule int

// LineCap is a representation of Cairo's cairo_line_cap_t.
type LineCap int

// LineJoin is a representation of Cairo's cairo_line_join_t.
type LineJoin int

// MimeType is a representation of Cairo's CAIRO_MIME_TYPE_*
// preprocessor constants.
type MimeType string

// Operator is a representation of Cairo's cairo_operator_t.
type Operator int

// Status is a representation of Cairo's cairo_status_t.
type Status int

// SurfaceType is a representation of Cairo's cairo_surface_type_t.
type SurfaceType int

// FontSlant is a representation of Cairo's cairo_font_slant_t
type FontSlant int

// FontWeight is a representation of Cairo's cairo_font_weight_t
type FontWeight int

var (
	FONT_SLANT_NORMAL  FontSlant
	FONT_SLANT_ITALIC  FontSlant
	FONT_SLANT_OBLIQUE FontSlant
)

var (
	FONT_WEIGHT_NORMAL FontWeight
	FONT_WEIGHT_BOLD   FontWeight
)

type FontExtents struct {
	Ascent      float64
	Descent     float64
	Height      float64
	MaxXAdvance float64
	MaxYAdvance float64
}

type TextExtents struct {
	XBearing float64
	YBearing float64
	Width    float64
	Height   float64
	XAdvance float64
	YAdvance float64
}

var (
	STATUS_SUCCESS                   Status
	STATUS_NO_MEMORY                 Status
	STATUS_INVALID_RESTORE           Status
	STATUS_INVALID_POP_GROUP         Status
	STATUS_NO_CURRENT_POINT          Status
	STATUS_INVALID_MATRIX            Status
	STATUS_INVALID_STATUS            Status
	STATUS_NULL_POINTER              Status
	STATUS_INVALID_STRING            Status
	STATUS_INVALID_PATH_DATA         Status
	STATUS_READ_ERROR                Status
	STATUS_WRITE_ERROR               Status
	STATUS_SURFACE_FINISHED          Status
	STATUS_SURFACE_TYPE_MISMATCH     Status
	STATUS_PATTERN_TYPE_MISMATCH     Status
	STATUS_INVALID_CONTENT           Status
	STATUS_INVALID_FORMAT            Status
	STATUS_INVALID_VISUAL            Status
	STATUS_FILE_NOT_FOUND            Status
	STATUS_INVALID_DASH              Status
	STATUS_INVALID_DSC_COMMENT       Status
	STATUS_INVALID_INDEX             Status
	STATUS_CLIP_NOT_REPRESENTABLE    Status
	STATUS_TEMP_FILE_ERROR           Status
	STATUS_INVALID_STRIDE            Status
	STATUS_FONT_TYPE_MISMATCH        Status
	STATUS_USER_FONT_IMMUTABLE       Status
	STATUS_USER_FONT_ERROR           Status
	STATUS_NEGATIVE_COUNT            Status
	STATUS_INVALID_CLUSTERS          Status
	STATUS_INVALID_SLANT             Status
	STATUS_INVALID_WEIGHT            Status
	STATUS_INVALID_SIZE              Status
	STATUS_USER_FONT_NOT_IMPLEMENTED Status
	STATUS_DEVICE_TYPE_MISMATCH      Status
	STATUS_DEVICE_ERROR              Status
)

var (
	SURFACE_TYPE_IMAGE          SurfaceType
	SURFACE_TYPE_PDF            SurfaceType
	SURFACE_TYPE_PS             SurfaceType
	SURFACE_TYPE_XLIB           SurfaceType
	SURFACE_TYPE_XCB            SurfaceType
	SURFACE_TYPE_GLITZ          SurfaceType
	SURFACE_TYPE_QUARTZ         SurfaceType
	SURFACE_TYPE_WIN32          SurfaceType
	SURFACE_TYPE_BEOS           SurfaceType
	SURFACE_TYPE_DIRECTFB       SurfaceType
	SURFACE_TYPE_SVG            SurfaceType
	SURFACE_TYPE_OS2            SurfaceType
	SURFACE_TYPE_WIN32_PRINTING SurfaceType
	SURFACE_TYPE_QUARTZ_IMAGE   SurfaceType
	SURFACE_TYPE_SCRIPT         SurfaceType
	SURFACE_TYPE_QT             SurfaceType
	SURFACE_TYPE_RECORDING      SurfaceType
	SURFACE_TYPE_VG             SurfaceType
	SURFACE_TYPE_GL             SurfaceType
	SURFACE_TYPE_DRM            SurfaceType
	SURFACE_TYPE_TEE            SurfaceType
	SURFACE_TYPE_XML            SurfaceType
	SURFACE_TYPE_SKIA           SurfaceType
	SURFACE_TYPE_SUBSURFACE     SurfaceType
)

var (
	MIME_TYPE_JP2       MimeType
	MIME_TYPE_JPEG      MimeType
	MIME_TYPE_PNG       MimeType
	MIME_TYPE_URI       MimeType
	MIME_TYPE_UNIQUE_ID MimeType
)

var (
	LINE_JOIN_MITER LineJoin
	LINE_JOIN_ROUND LineJoin
	LINE_JOIN_BEVEL LineJoin
)

var (
	CONTENT_COLOR       Content
	CONTENT_ALPHA       Content
	CONTENT_COLOR_ALPHA Content
)

var (
	OPERATOR_CLEAR          Operator
	OPERATOR_SOURCE         Operator
	OPERATOR_OVER           Operator
	OPERATOR_IN             Operator
	OPERATOR_OUT            Operator
	OPERATOR_ATOP           Operator
	OPERATOR_DEST           Operator
	OPERATOR_DEST_OVER      Operator
	OPERATOR_DEST_IN        Operator
	OPERATOR_DEST_OUT       Operator
	OPERATOR_DEST_ATOP      Operator
	OPERATOR_XOR            Operator
	OPERATOR_ADD            Operator
	OPERATOR_SATURATE       Operator
	OPERATOR_MULTIPLY       Operator
	OPERATOR_SCREEN         Operator
	OPERATOR_OVERLAY        Operator
	OPERATOR_DARKEN         Operator
	OPERATOR_LIGHTEN        Operator
	OPERATOR_COLOR_DODGE    Operator
	OPERATOR_COLOR_BURN     Operator
	OPERATOR_HARD_LIGHT     Operator
	OPERATOR_SOFT_LIGHT     Operator
	OPERATOR_DIFFERENCE     Operator
	OPERATOR_EXCLUSION      Operator
	OPERATOR_HSL_HUE        Operator
	OPERATOR_HSL_SATURATION Operator
	OPERATOR_HSL_COLOR      Operator
	OPERATOR_HSL_LUMINOSITY Operator
)

var (
	FILL_RULE_WINDING  FillRule
	FILL_RULE_EVEN_ODD FillRule
)

var (
	ANTIALIAS_DEFAULT  Antialias
	ANTIALIAS_NONE     Antialias
	ANTIALIAS_GRAY     Antialias
	ANTIALIAS_SUBPIXEL Antialias
	// ANTIALIAS_FAST      = C.CAIRO_ANTIALIAS_FAST (since 1.12)
	// ANTIALIAS_GOOD      = C.CAIRO_ANTIALIAS_GOOD (since 1.12)
	// ANTIALIAS_BEST     iface.Antialias = C.CAIRO_ANTIALIAS_BEST (since 1.12)
)

var (
	LINE_CAP_BUTT   LineCap
	LINE_CAP_ROUND  LineCap
	LINE_CAP_SQUARE LineCap
)
