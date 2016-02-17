package pango

type Layout interface {
	Copy() Layout
	GetAttributes() AttrList
	GetCharacterCount() int
	GetContext() Context
	GetFontDescription() FontDescription
	GetHeight() int
	GetIndent() int
	GetSize() (int, int)
	GetText() string
	GetWidth() int
	GetWrap() WrapMode
	IsWrapped() bool
	SetAttributes(AttrList)
	SetFontDescription(FontDescription)
	SetHeight(int)
	SetIndent(int)
	SetMarkup(string, int)
	SetText(string, int)
	SetWidth(int)
	SetWrap(WrapMode)
} // end of Layout

func AssertLayout(_ Layout) {}
