package gtk

import "github.com/gotk3/gotk3/pango"

type Label interface {
	Widget

	GetAngle() float64
	GetCurrentUri() string
	GetEllipsize() pango.EllipsizeMode
	GetJustify() Justification
	GetLineWrap() bool
	GetMaxWidthChars() int
	GetSelectable() bool
	GetSelectionBounds() (int, int, bool)
	GetSingleLineMode() bool
	GetText() (string, error)
	GetTrackVisitedLinks() bool
	GetUseMarkup() bool
	GetUseUnderline() bool
	GetWidthChars() int
	SelectRegion(int, int)
	SetAngle(float64)
	SetEllipsize(pango.EllipsizeMode)
	SetJustify(Justification)
	SetLabel(string)
	SetLineWrap(bool)
	SetLineWrapMode(pango.WrapMode)
	SetMarkup(string)
	SetMarkupWithMnemonic(string)
	SetMaxWidthChars(int)
	SetPattern(string)
	SetSelectable(bool)
	SetSingleLineMode(bool)
	SetText(string)
	SetTrackVisitedLinks(bool)
	SetUseMarkup(bool)
	SetUseUnderline(bool)
	SetWidthChars(int)
} // end of Label

func AssertLabel(_ Label) {}
