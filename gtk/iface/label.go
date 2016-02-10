package iface

import pango_iface "github.com/gotk3/gotk3/pango/iface"

type Label interface {
    Widget

    GetAngle() float64
    GetCurrentUri() string
    GetEllipsize() pango_iface.EllipsizeMode
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
    SetEllipsize(pango_iface.EllipsizeMode)
    SetJustify(Justification)
    SetLabel(string)
    SetLineWrap(bool)
    SetLineWrapMode(pango_iface.WrapMode)
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
