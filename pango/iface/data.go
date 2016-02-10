package iface

type Gravity int

type GravityHint int

type Style int

type Variant int

type Weight int

type Stretch int

type FontMask int

type Scale float64

// WrapMode is a representation of Pango's PangoWrapMode.
type WrapMode int

// EllipsizeMode is a representation of Pango's PangoEllipsizeMode.
type EllipsizeMode int

// Alignment is a representation of Pango's PangoAlignment.
type Alignment int

// AttrType is a representation of Pango's PangoAttrType.
type AttrType int

// Underline is a representation of Pango's PangoUnderline.
type Underline int

// Glyph is a representation of PangoGlyph
type Glyph uint32

var (
	GRAVITY_SOUTH Gravity
	GRAVITY_EAST  Gravity
	GRAVITY_NORTH Gravity
	GRAVITY_WEST  Gravity
	GRAVITY_AUTO  Gravity
)

var (
	GRAVITY_HINT_NATURAL GravityHint
	GRAVITY_HINT_STRONG  GravityHint
	GRAVITY_HINT_LINE    GravityHint
)

var (
	ATTR_INVALID             AttrType
	ATTR_LANGUAGE            AttrType
	ATTR_FAMILY              AttrType
	ATTR_STYLE               AttrType
	ATTR_WEIGHT              AttrType
	ATTR_VARIANT             AttrType
	ATTR_STRETCH             AttrType
	ATTR_SIZE                AttrType
	ATTR_FONT_DESC           AttrType
	ATTR_FOREGROUND          AttrType
	ATTR_BACKGROUND          AttrType
	ATTR_UNDERLINE           AttrType
	ATTR_STRIKETHROUGH       AttrType
	ATTR_RISE                AttrType
	ATTR_SHAPE               AttrType
	ATTR_SCALE               AttrType
	ATTR_FALLBACK            AttrType
	ATTR_LETTER_SPACING      AttrType
	ATTR_UNDERLINE_COLOR     AttrType
	ATTR_STRIKETHROUGH_COLOR AttrType
	ATTR_ABSOLUTE_SIZE       AttrType
	ATTR_GRAVITY             AttrType
	ATTR_GRAVITY_HINT        AttrType
)

var (
	UNDERLINE_NONE   Underline
	UNDERLINE_SINGLE Underline
	UNDERLINE_DOUBLE Underline
	UNDERLINE_LOW    Underline
	UNDERLINE_ERROR  Underline
)

var (
	ATTR_INDEX_FROM_TEXT_BEGINNING uint = 0
	ATTR_INDEX_TO_TEXT_END         uint
)

var (
	PANGO_SCALE int
)

var (
	STYLE_NORMAL  Style
	STYLE_OBLIQUE Style
	STYLE_ITALIC  Style
)

var (
	VARIANT_NORMAL     Variant
	VARIANT_SMALL_CAPS Variant
)

var (
	WEIGHT_THIN       Weight
	WEIGHT_ULTRALIGHT Weight
	WEIGHT_LIGHT      Weight
	WEIGHT_SEMILIGHT  Weight
	WEIGHT_BOOK       Weight
	WEIGHT_NORMAL     Weight
	WEIGHT_MEDIUM     Weight
	WEIGHT_SEMIBOLD   Weight
	WEIGHT_BOLD       Weight
	WEIGHT_ULTRABOLD  Weight
	WEIGHT_HEAVY      Weight
	WEIGHT_ULTRAHEAVY Weight
)

var (
	STRETCH_ULTRA_CONDENSED        Stretch
	STRETCH_EXTRA_CONDENSEDStretch Stretch
	STRETCH_CONDENSEDStretch       Stretch
	STRETCH_SEMI_CONDENSEDStretch  Stretch
	STRETCH_NORMALStretch          Stretch
	STRETCH_SEMI_EXPANDEDStretch   Stretch
	STRETCH_EXPANDEDStretch        Stretch
	STRETCH_EXTRA_EXPANDEDStretch  Stretch
	STRETCH_ULTRA_EXPANDEDStretch  Stretch
)

var (
	FONT_MASK_FAMILY          FontMask
	FONT_MASK_STYLEFontMask   FontMask
	FONT_MASK_VARIANTFontMask FontMask
	FONT_MASK_WEIGHTFontMask  FontMask
	FONT_MASK_STRETCHFontMask FontMask
	FONT_MASK_SIZEFontMask    FontMask
	FONT_MASK_GRAVITYFontMask FontMask
)

var (
	SCALE_XX_SMALL Scale
	SCALE_X_SMALL  Scale
	SCALE_SMALL    Scale
	SCALE_MEDIUM   Scale
	SCALE_LARGE    Scale
	SCALE_X_LARGE  Scale
	SCALE_XX_LARGE Scale
)

var (
	SCALE int
)

var (
	ALIGN_LEFT   Alignment
	ALIGN_CENTER Alignment
	ALIGN_RIGHT  Alignment
)

var (
	WRAP_WORD      WrapMode
	WRAP_CHAR      WrapMode
	WRAP_WORD_CHAR WrapMode
)

var (
	ELLIPSIZE_NONE   EllipsizeMode
	ELLIPSIZE_START  EllipsizeMode
	ELLIPSIZE_MIDDLE EllipsizeMode
	ELLIPSIZE_END    EllipsizeMode
)
