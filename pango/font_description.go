package pango

type FontDescription interface {
	BetterMatch(FontDescription, FontDescription) bool
	Copy() FontDescription
	CopyStatic() FontDescription
	Equal(FontDescription) bool
	Free()
	GetFamily() string
	GetGravity() Gravity
	GetSetFields() FontMask
	GetSize() int
	GetSizeIsAbsolute() bool
	GetStretch() Stretch
	GetStyle() Style
	GetUnsetFields(FontMask)
	GetWeight() Weight
	Hash() uint
	Merge(FontDescription, bool)
	MergeStatic(FontDescription, bool)
	SetAbsoluteSize(float64)
	SetFamily(string)
	SetFamilyStatic(string)
	SetGravity(Gravity)
	SetSize(int)
	SetStretch(Stretch)
	SetStyle(Style)
	SetWeight(Weight)
	ToFilename() string
	ToString() string
} // end of FontDescription

func AssertFontDescription(_ FontDescription) {}
