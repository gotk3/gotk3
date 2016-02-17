package impl

import "github.com/gotk3/gotk3/pango"

func toAttrList(s pango.AttrList) *attrList {
	if s == nil {
		return nil
	}
	return s.(*attrList)
}

func toContext(s pango.Context) *context {
	if s == nil {
		return nil
	}
	return s.(*context)
}

func toFontDescription(s pango.FontDescription) *fontDescription {
	if s == nil {
		return nil
	}
	return s.(*fontDescription)
}

func toFont(s pango.Font) *font {
	if s == nil {
		return nil
	}
	return s.(*font)
}

func toLayout(s pango.Layout) *layout {
	if s == nil {
		return nil
	}
	return s.(*layout)
}

func toLayoutLine(s pango.LayoutLine) *layoutLine {
	if s == nil {
		return nil
	}
	return s.(*layoutLine)
}

func toGlyphString(s pango.GlyphString) *glyphString {
	if s == nil {
		return nil
	}
	return s.(*glyphString)
}

func toGlyphItem(s pango.GlyphItem) *glyphItem {
	if s == nil {
		return nil
	}
	return s.(*glyphItem)
}

func toRectangle(s pango.Rectangle) *rectangle {
	if s == nil {
		return nil
	}
	return s.(*rectangle)
}
