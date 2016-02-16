package impl

import "github.com/gotk3/gotk3/pango"

func init() {
	pango.AssertPango(&RealPango{})
	pango.AssertAttrClass(&AttrClass{})
	pango.AssertAttrColor(&AttrColor{})
	pango.AssertAttrFloat(&AttrFloat{})
	pango.AssertAttrFontDesc(&AttrFontDesc{})
	pango.AssertAttrInt(&AttrInt{})
	pango.AssertAttrLanguage(&AttrLanguage{})
	pango.AssertAttrList(&AttrList{})
	pango.AssertAttrShape(&AttrShape{})
	pango.AssertAttrSize(&AttrSize{})
	pango.AssertAttrString(&AttrString{})
	pango.AssertAttribute(&Attribute{})
	pango.AssertColor(&Color{})
	pango.AssertContext(&Context{})
	pango.AssertEngineLang(&EngineLang{})
	pango.AssertEngineShape(&EngineShape{})
	pango.AssertFont(&Font{})
	pango.AssertFontDescription(&FontDescription{})
	pango.AssertFontMap(&FontMap{})
	pango.AssertFontMetrics(&FontMetrics{})
	pango.AssertGlyphGeometry(&GlyphGeometry{})
	pango.AssertGlyphInfo(&GlyphInfo{})
	pango.AssertGlyphItem(&GlyphItem{})
	pango.AssertGlyphString(&GlyphString{})
	pango.AssertGlyphVisAttr(&GlyphVisAttr{})
	pango.AssertLayout(&Layout{})
	pango.AssertLayoutLine(&LayoutLine{})
	pango.AssertLogAttr(&LogAttr{})
	pango.AssertRealPango(&RealPango{})
	pango.AssertRectangle(&Rectangle{})
}
