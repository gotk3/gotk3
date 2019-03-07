/*
 * Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
 *
 * This file originated from: http://opensource.conformal.com/
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package gtk

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
)

func init() {
	Init(nil)
}

// TestBoolConvs tests the conversion between Go bools and gboolean
// types.
func TestBoolConvs(t *testing.T) {
	if err := testBoolConvs(); err != nil {
		t.Error(err)
	}
}

// TestBox tests creating and adding widgets to a Box
func TestBox(t *testing.T) {
	vbox, err := BoxNew(ORIENTATION_VERTICAL, 0)
	if err != nil {
		t.Error("Unable to create box")
	}

	vbox.Set("homogeneous", true)
	if vbox.GetHomogeneous() != true {
		t.Error("Could not set or get Box homogeneous property")
	}

	vbox.SetHomogeneous(false)
	if vbox.GetHomogeneous() != false {
		t.Error("Could not set or get Box homogeneous property")
	}

	vbox.Set("spacing", 1)
	if vbox.GetSpacing() != 1 {
		t.Error("Could not set or get Box spacing")
	}

	vbox.SetSpacing(2)
	if vbox.GetSpacing() != 2 {
		t.Error("Could not set or get Box spacing")
	}

	// add a child to start and end
	start, err := LabelNew("Start")
	if err != nil {
		t.Error("Unable to create label")
	}

	end, err := LabelNew("End")
	if err != nil {
		t.Error("Unable to create label")
	}

	vbox.PackStart(start, true, true, 3)
	vbox.PackEnd(end, true, true, 3)
}
func TestTextBuffer_WhenSetText_ExpectGetTextReturnsSame(t *testing.T) {
	buffer, err := TextBufferNew(nil)
	if err != nil {
		t.Error("Unable to create text buffer")
	}
	expected := "Hello, World!"
	buffer.SetText(expected)

	start, end := buffer.GetBounds()

	actual, err := buffer.GetText(start, end, true)
	if err != nil {
		t.Error("Unable to get text from buffer")
	}

	if actual != expected {
		t.Errorf("Expected '%s'; Got '%s'", expected, actual)
	}
}

func testTextViewEditable(set bool) error {
	tv, err := TextViewNew()
	if err != nil {
		return err
	}

	exp := set
	tv.SetEditable(exp)
	act := tv.GetEditable()
	if exp != act {
		return fmt.Errorf("Expected GetEditable(): %v; Got: %v", exp, act)
	}
	return nil
}

func TestTextView_WhenSetEditableFalse_ExpectGetEditableReturnsFalse(t *testing.T) {
	if err := testTextViewEditable(false); err != nil {
		t.Error(err)
	}
}

func TestTextView_WhenSetEditableTrue_ExpectGetEditableReturnsTrue(t *testing.T) {
	if err := testTextViewEditable(true); err != nil {
		t.Error(err)
	}
}

func testTextViewWrapMode(set WrapMode) error {
	tv, err := TextViewNew()
	if err != nil {
		return err
	}

	exp := set
	tv.SetWrapMode(set)
	act := tv.GetWrapMode()
	if act != exp {
		return fmt.Errorf("Expected GetWrapMode(): %v; Got: %v", exp, act)
	}
	return nil
}

func TestTextView_WhenSetWrapModeNone_ExpectGetWrapModeReturnsNone(t *testing.T) {
	if err := testTextViewWrapMode(WRAP_NONE); err != nil {
		t.Error(err)
	}
}

func TestTextView_WhenSetWrapModeWord_ExpectGetWrapModeReturnsWord(t *testing.T) {
	if err := testTextViewWrapMode(WRAP_WORD); err != nil {
		t.Error(err)
	}
}

func testTextViewCursorVisible(set bool) error {
	tv, err := TextViewNew()
	if err != nil {
		return err
	}

	exp := set
	tv.SetCursorVisible(set)
	act := tv.GetCursorVisible()
	if act != exp {
		return fmt.Errorf("Expected GetCursorVisible(): %v; Got: %v", exp, act)
	}
	return nil
}

func TestTextView_WhenSetCursorVisibleFalse_ExpectGetCursorVisibleReturnsFalse(t *testing.T) {
	if err := testTextViewCursorVisible(false); err != nil {
		t.Error(err)
	}
}

func TestTextView_WhenSetCursorVisibleTrue_ExpectGetCursorVisibleReturnsTrue(t *testing.T) {
	if err := testTextViewCursorVisible(true); err != nil {
		t.Error(err)
	}
}

func testTextViewOverwrite(set bool) error {
	tv, err := TextViewNew()
	if err != nil {
		return err
	}

	exp := set
	tv.SetOverwrite(set)
	act := tv.GetOverwrite()
	if act != exp {
		return fmt.Errorf("Expected GetOverwrite(): %v; Got: %v", exp, act)
	}
	return nil
}

func TestTextView_WhenSetOverwriteFalse_ExpectGetOverwriteReturnsFalse(t *testing.T) {
	if err := testTextViewOverwrite(false); err != nil {
		t.Error(err)
	}
}

func TestTextView_WhenSetOverwriteTrue_ExpectGetOverwriteReturnsTrue(t *testing.T) {
	if err := testTextViewOverwrite(true); err != nil {
		t.Error(err)
	}
}

func testTextViewJustification(justify Justification) error {
	tv, err := TextViewNew()
	if err != nil {
		return err
	}

	exp := justify
	tv.SetJustification(justify)
	act := tv.GetJustification()
	if act != exp {
		return fmt.Errorf("Expected GetJustification(): %v; Got: %v", exp, act)
	}
	return nil
}

func TestTextView_WhenSetJustificationLeft_ExpectGetJustificationReturnsLeft(t *testing.T) {
	if err := testTextViewJustification(JUSTIFY_LEFT); err != nil {
		t.Error(err)
	}
}

func TestTextView_WhenSetJustificationRight_ExpectGetJustificationReturnsRight(t *testing.T) {
	if err := testTextViewJustification(JUSTIFY_RIGHT); err != nil {
		t.Error(err)
	}
}

func testTextViewAcceptsTab(set bool) error {
	tv, err := TextViewNew()
	if err != nil {
		return err
	}

	exp := set
	tv.SetAcceptsTab(set)
	if act := tv.GetAcceptsTab(); act != exp {
		return fmt.Errorf("Expected GetAcceptsTab(): %v; Got: %v", exp, act)
	}
	return nil
}

func TestTextView_WhenSetAcceptsTabFalse_ExpectGetAcceptsTabReturnsFalse(t *testing.T) {
	if err := testTextViewAcceptsTab(false); err != nil {
		t.Error(err)
	}
}

func TestTextView_WhenSetAcceptsTabTrue_ExpectGetAcceptsTabReturnsTrue(t *testing.T) {
	if err := testTextViewAcceptsTab(true); err != nil {
		t.Error(err)
	}
}

func testIntProperty(val int, set func(int), get func() int) error {
	set(val)
	if exp, act := val, get(); act != exp {
		return fmt.Errorf("Expected: %d; got: %d", exp, act)
	}
	return nil
}

func testTextViewPixelsAboveLines(px int) error {
	tv, err := TextViewNew()
	if err != nil {
		return err
	}
	return testIntProperty(px, (*tv).SetPixelsAboveLines, (*tv).GetPixelsAboveLines)
}

func TestTextView_WhenSetPixelsAboveLines10_ExpectGetPixelsAboveLinesReturns10(t *testing.T) {
	if err := testTextViewPixelsAboveLines(10); err != nil {
		t.Error(err)
	}
}

func TestTextView_WhenSetPixelsAboveLines11_ExpectGetPixelsAboveLinesReturns11(t *testing.T) {
	if err := testTextViewPixelsAboveLines(11); err != nil {
		t.Error(err)
	}
}

func testTextViewPixelsBelowLines(px int) error {
	tv, err := TextViewNew()
	if err != nil {
		return err
	}
	return testIntProperty(px, (*tv).SetPixelsBelowLines, (*tv).GetPixelsBelowLines)
}

func TestTextView_WhenSetPixelsBelowLines10_ExpectGetPixelsAboveLinesReturns10(t *testing.T) {
	if err := testTextViewPixelsBelowLines(10); err != nil {
		t.Error(err)
	}
}

func TestTextView_WhenSetPixelsBelowLines11_ExpectGetPixelsBelowLinesReturns11(t *testing.T) {
	if err := testTextViewPixelsBelowLines(11); err != nil {
		t.Error(err)
	}
}

func testTextViewPixelsInsideWrap(px int) error {
	tv, err := TextViewNew()
	if err != nil {
		return err
	}

	return testIntProperty(px, (*tv).SetPixelsInsideWrap, (*tv).GetPixelsInsideWrap)
}

func TestTextView_WhenSetPixelsInsideWrap10_ExpectGetPixelsInsideWrapReturns11(t *testing.T) {
	if err := testTextViewPixelsInsideWrap(10); err != nil {
		t.Error(err)
	}
}

func TestTextView_WhenSetPixelsInsideWrap11_ExpectGetPixelsInsideWrapReturns11(t *testing.T) {
	if err := testTextViewPixelsInsideWrap(11); err != nil {
		t.Error(err)
	}
}

func testTextViewLeftMargin(margin int) error {
	tv, err := TextViewNew()
	if err != nil {
		return err
	}

	return testIntProperty(margin, (*tv).SetLeftMargin, (*tv).GetLeftMargin)
}

func TestTextView_WhenSetLeftMargin11_ExpectGetLeftMarginReturns11(t *testing.T) {
	if err := testTextViewLeftMargin(11); err != nil {
		t.Error(err)
	}
}

func TestTextView_WhenSetLeftMargin10_ExpectGetLeftMarginReturns10(t *testing.T) {
	if err := testTextViewLeftMargin(10); err != nil {
		t.Error(err)
	}
}

func testTextViewRightMargin(margin int) error {
	tv, err := TextViewNew()
	if err != nil {
		return err
	}

	return testIntProperty(margin, (*tv).SetRightMargin, (*tv).GetRightMargin)
}

func TestTextView_WhenSetRightMargin10_ExpectGetRightMarginReturns10(t *testing.T) {
	if err := testTextViewRightMargin(10); err != nil {
		t.Error(err)
	}
}

func TestTextView_WhenSetRightMargin11_ExpectGetRightMarginReturns11(t *testing.T) {
	if err := testTextViewRightMargin(11); err != nil {
		t.Error(err)
	}
}

func testTextViewIndent(indent int) error {
	tv, err := TextViewNew()
	if err != nil {
		return err
	}

	return testIntProperty(indent, (*tv).SetIndent, (*tv).GetIndent)
}

func TestTextView_WhenSetIndent10_ExpectGetIndentReturns10(t *testing.T) {
	if err := testTextViewIndent(10); err != nil {
		t.Error(err)
	}
}

func TestTextView_WhenSetIndent11_ExpectGetIndentReturns11(t *testing.T) {
	if err := testTextViewIndent(11); err != nil {
		t.Error(err)
	}
}

func testTextViewInputHints(hint InputHints) error {
	tv, err := TextViewNew()
	if err != nil {
		return err
	}

	tv.SetInputHints(hint)
	if exp, act := hint, tv.GetInputHints(); act != exp {
		return fmt.Errorf("Expected %v; Got %v", exp, act)
	}
	return nil
}

func TestTextView_WhenSetInputHintsNone_ExpectGetInputHintsReturnsNone(t *testing.T) {
	if err := testTextViewInputHints(INPUT_HINT_NONE); err != nil {
		t.Error(err)
	}
}

func TestTextView_WhenSetInputHintsSpellCheck_ExpectGetInputHintsReturnsSpellCheck(t *testing.T) {
	if err := testTextViewInputHints(INPUT_HINT_SPELLCHECK); err != nil {
		t.Error(err)
	}
}

func testTextViewInputPurpose(purpose InputPurpose) error {
	tv, err := TextViewNew()
	if err != nil {
		return err
	}

	tv.SetInputPurpose(purpose)
	if exp, act := purpose, tv.GetInputPurpose(); act != exp {
		return fmt.Errorf("Expected %v; Got %v", exp, act)
	}
	return nil
}

func TestTextView_WhenSetInputPurposeURL_ExpectGetInputPurposeReturnsURL(t *testing.T) {
	if err := testTextViewInputPurpose(INPUT_PURPOSE_URL); err != nil {
		t.Error(err)
	}
}

func TestTextView_WhenSetInputPurposeALPHA_ExpectGetInputPurposeReturnsALPHA(t *testing.T) {
	if err := testTextViewInputPurpose(INPUT_PURPOSE_ALPHA); err != nil {
		t.Error(err)
	}
}

func testCellRendererToggleSetRadio(set bool) error {
	renderer, err := CellRendererToggleNew()
	if err != nil {
		return err
	}

	renderer.SetRadio(set)
	if exp, act := set, renderer.GetRadio(); act != exp {
		return fmt.Errorf("Expected GetRadio(): %v; Got: %v", exp, act)
	}
	return nil
}

func TestCellRendererToggle_WhenSetRadioFalse_ExpectGetRadioReturnsFalse(t *testing.T) {
	if err := testCellRendererToggleSetRadio(false); err != nil {
		t.Error(err)
	}
}

func TestCellRendererToggle_WhenSetRadioTrue_ExpectGetRadioReturnsTrue(t *testing.T) {
	if err := testCellRendererToggleSetRadio(true); err != nil {
		t.Error(err)
	}
}

func testCellRendererToggleSetActive(set bool) error {
	renderer, err := CellRendererToggleNew()
	if err != nil {
		return err
	}

	renderer.SetActive(set)
	if exp, act := set, renderer.GetActive(); act != exp {
		return fmt.Errorf("Expected GetActive(): %v; Got: %v", exp, act)
	}
	return nil
}

func TestCellRendererToggle_WhenSetActiveFalse_ExpectGetActiveReturnsFalse(t *testing.T) {
	if err := testCellRendererToggleSetActive(false); err != nil {
		t.Error(err)
	}
}

func TestCellRendererToggle_WhenSetActiveTrue_ExpectGetActiveReturnsTrue(t *testing.T) {
	if err := testCellRendererToggleSetActive(true); err != nil {
		t.Error(err)
	}
}

func testCellRendererToggleSetActivatable(set bool) error {
	renderer, err := CellRendererToggleNew()
	if err != nil {
		return err
	}

	renderer.SetActivatable(set)
	if exp, act := set, renderer.GetActivatable(); act != exp {
		return fmt.Errorf("Expected GetActivatable(): %v; Got: %v", exp, act)
	}
	return nil
}

func TestCellRendererToggle_WhenSetActivatableFalse_ExpectGetActivatableReturnsFalse(t *testing.T) {
	if err := testCellRendererToggleSetActivatable(false); err != nil {
		t.Error(err)
	}
}

func TestCellRendererToggle_WhenSetActivatableTrue_ExpectGetActivatableReturnsTrue(t *testing.T) {
	if err := testCellRendererToggleSetActivatable(true); err != nil {
		t.Error(err)
	}
}

func setupListStore() *ListStore {
	ls, err := ListStoreNew(glib.TYPE_STRING)
	if err != nil {
		log.Fatal("Unexpected err:", err)
	}
	return ls
}

func getLastIter(ls *ListStore) (*TreeIter, bool) {
	iter, listIsntEmpty := ls.GetIterFirst()
	if !listIsntEmpty {
		return iter, listIsntEmpty
	}

	for {
		temp := *iter
		last := &temp
		if !ls.IterNext(iter) {
			return last, true
		}
	}

	panic("Shouldn't get here")
}

// TestListStoreRemoveLastInvalidIterator tests that when a ListStore stores
// one item and it is removed, the iterator becomes invalid.
func TestListStoreRemoveLastInvalidIterator(t *testing.T) {
	ls := setupListStore()

	iter := ls.Append()

	if iterValid := ls.Remove(iter); iterValid {
		t.Fatal("Remove() returned true (iter valid); expected false (iter invalid)")
	}
}

func TestListStoreInsertBefore(t *testing.T) {
	ls := setupListStore()

	// Given 1 iter is already in the liststore
	initialIter := ls.Append()

	// When another iter is inserted before it
	newIter := ls.InsertBefore(initialIter)

	// Expect the newly-inserted iter is first iter in list
	firstIter, listIsntEmpty := ls.GetIterFirst()
	if !listIsntEmpty {
		t.Fatal("Unexpected: liststore is empty")
	}

	if *firstIter != *newIter {
		t.Fatal("Expected the new iter added to front of list")
	}
}

// When 'sibling' parameter is nil, the new iter should be appended to the liststore
func TestListStoreInsertBefore_WhenNilSibling(t *testing.T) {
	ls := setupListStore()

	// Given 2 iters in liststore
	ls.Append()
	ls.Append()

	// When 'sibling' parameter of InsertBefore() is nil...
	newIter := ls.InsertBefore(nil)

	// Then expect newly-inserted iter is the first iter in list
	lastIter, listIsntEmpty := getLastIter(ls)
	if !listIsntEmpty {
		t.Fatal("Unexpected: liststore is empty")
	}

	if *lastIter != *newIter {
		t.Fatal("Expected the new iter added to end of list")
	}
}

func TestListStoreInsertAfter(t *testing.T) {
	ls := setupListStore()

	// Given 1 iter in liststore
	sibling := ls.Append()

	// When InsertAfter(sibling)
	newIter := ls.InsertAfter(sibling)

	// Then expect newly-inserted iter is the last iter in list
	lastIter, listIsntEmpty := getLastIter(ls)
	if !listIsntEmpty {
		t.Fatal("Unexpected: liststore is empty")
	}

	if *lastIter != *newIter {
		t.Fatal("Expected the new iter added to end of list")
	}
}

// When 'sibling' parameter is nil, the new iter should be prepended to the liststore
func TestListStoreInsertAfter_WhenNilSibling(t *testing.T) {
	ls := setupListStore()

	// Given 2 iters in liststore
	ls.Append()
	ls.Append()

	// When InsertAfter(nil)
	newIter := ls.InsertAfter(nil)

	// Then expect newly-inserted iter is the first iter in the list
	first, listIsntEmpty := ls.GetIterFirst()
	if !listIsntEmpty {
		t.Fatal("Unexpected: liststore is empty")
	}

	if *first != *newIter {
		t.Fatal("Expected the new iter was prepended to liststore")
	}
}

func TestBuilder(t *testing.T) {
	builder, err := BuilderNew()
	if err != nil {
		t.Error("Unable to create builder")
	}

	str := `
<interface>
  <object class="GtkDialog" id="dialog1">
    <child internal-child="vbox">
      <object class="GtkBox" id="vbox1">
        <property name="border-width">10</property>
        <child internal-child="action_area">
          <object class="GtkButtonBox" id="hbuttonbox1">
            <property name="border-width">20</property>
            <child>
              <object class="GtkButton" id="ok_button">
                <property name="label">gtk-ok</property>
                <property name="use-stock">TRUE</property>
                <signal name="clicked" handler="ok_button_clicked"/>
              </object>
            </child>
          </object>
        </child>
      </object>
    </child>
  </object>
</interface>
`

	err = builder.AddFromString(str)
	if err != nil {
		t.Error("Unable to add from string")
	}

	widget, err := builder.GetObject("ok_button")
	if err != nil {
		t.Error("Unable to get widget from string")
	}

	button, ok := widget.(*Button)
	if !ok {
		t.Error("Unable to cast to gtk.Button")
	}

	l, err := button.GetLabel()
	if err != nil {
		t.Error("Unable to get button label")
	}

	if l != "gtk-ok" {
		t.Errorf("Label has the wrong value: %q", l)
	}

	done := make(chan bool)

	builder.ConnectSignals(map[string]interface{}{
		"ok_button_clicked": func() {
			done <- true
		},
	})

	go button.Emit("clicked")

	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Error("Failed to call callback")
	}
}

func TestTextTagEvent(t *testing.T) {
	textTag, err := TextTagNew("mytexttag")
	if err != nil {
		t.Error("could not create text tag")
	}

	evk := gdk.EventKeyNew()

	var iter TextIter
	ok := textTag.Event(textTag.Object, evk.Event, &iter)

	if ok {
		t.Error("event should not have been handled")
	}

	textTag.Connect("event", func() bool {
		return true
	})

	ok = textTag.Event(textTag.Object, evk.Event, &iter)

	if !ok {
		t.Error("event should have been handled")
	}

}
