package impl

import "github.com/gotk3/gotk3/pango"

func toAttrList(s pango.AttrList) *AttrList {
	if s == nil {
		return nil
	}
	return s.(*AttrList)
}

func toContext(s pango.Context) *Context {
	if s == nil {
		return nil
	}
	return s.(*Context)
}

func toFontDescription(s pango.FontDescription) *FontDescription {
	if s == nil {
		return nil
	}
	return s.(*FontDescription)
}

func toFont(s pango.Font) *Font {
	if s == nil {
		return nil
	}
	return s.(*Font)
}

func toLayout(s pango.Layout) *Layout {
	if s == nil {
		return nil
	}
	return s.(*Layout)
}

func toLayoutLine(s pango.LayoutLine) *LayoutLine {
	if s == nil {
		return nil
	}
	return s.(*LayoutLine)
}

func toGlyphString(s pango.GlyphString) *GlyphString {
	if s == nil {
		return nil
	}
	return s.(*GlyphString)
}

func toGlyphItem(s pango.GlyphItem) *GlyphItem {
	if s == nil {
		return nil
	}
	return s.(*GlyphItem)
}

func toRectangle(s pango.Rectangle) *Rectangle {
	if s == nil {
		return nil
	}
	return s.(*Rectangle)
}
