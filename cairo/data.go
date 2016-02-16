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
