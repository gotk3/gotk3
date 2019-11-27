package gtk

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func TestTestRegisterAllTypes(t *testing.T) {
	TestRegisterAllTypes()
	types := TestListAllTypes()

	if len(types) == 0 {
		t.Error("Expected at least one type to be registered")
	}
}

func TestPointerAtOffset(t *testing.T) {
	// Simulate a C array by using a pointer to the first element
	intArray := []int{4, 8, 2, 5, 9}
	arrayPointer := unsafe.Pointer(&intArray[0])
	elementSize := unsafe.Sizeof(intArray[0])

	for i, val := range intArray {
		intAtOffset := (*int)(pointerAtOffset(arrayPointer, elementSize, uint(i)))
		if val != *intAtOffset {
			t.Errorf("Expected %d at offset %d, got %d", val, i, *intAtOffset)
		}
	}
}

func TestTestFindLabel(t *testing.T) {
	// Build a dummy widget
	box, _ := BoxNew(ORIENTATION_HORIZONTAL, 0)
	label1, _ := LabelNew("First")
	label2, _ := LabelNew("Second")

	box.PackStart(label1, true, true, 0)
	box.PackStart(label2, true, true, 0)

	// Find a label in the box with text matching Fir*
	found, err := TestFindLabel(box, "Fir*")
	if err != nil {
		t.Error("Unexpected error:", err.Error())
	}

	// Should return the label1
	if found == nil {
		t.Error("Return value is nil")
	}
	foundAsLabel, ok := found.(*Label)
	if !ok {
		t.Error("Did not return a label. Received type:", reflect.TypeOf(found))
	}

	text, _ := foundAsLabel.GetText()
	if text != "First" {
		t.Error("Expected: First, Got:", text)
	}

}

func TestTestFindSibling(t *testing.T) {
	// Build a dummy widget
	box, _ := BoxNew(ORIENTATION_HORIZONTAL, 0)
	label1, _ := LabelNew("First")
	label2, _ := LabelNew("Second")

	box.PackStart(label1, true, true, 0)
	box.PackStart(label2, true, true, 0)

	// Finx a sibling to label1, of type label
	found, err := TestFindSibling(label1, glib.TypeFromName("GtkLabel"))
	if err != nil {
		t.Error("Unexpected error:", err.Error())
	}

	// Should return the label2
	if found == nil {
		t.Error("Return value is nil")
	}
	foundAsLabel, ok := found.(*Label)
	if !ok {
		t.Error("Did not return a label. Received type:", reflect.TypeOf(found))
	}

	text, _ := foundAsLabel.GetText()
	if text != "Second" {
		t.Error("Expected: First, Got:", text)
	}

}

func TestTestFindWidget(t *testing.T) {
	// Build a dummy widget
	box, _ := BoxNew(ORIENTATION_HORIZONTAL, 0)
	button1, _ := ButtonNewWithLabel("First")
	button2, _ := ButtonNewWithLabel("Second")

	box.PackStart(button1, true, true, 0)
	box.PackStart(button2, true, true, 0)

	// Find a label in the box with text matching Fir*
	found, err := TestFindWidget(box, "Sec*", glib.TypeFromName("GtkButton"))
	if err != nil {
		t.Error("Unexpected error:", err.Error())
	}

	// Should return the button2
	if found == nil {
		t.Error("Return value is nil")
	}
	foundAsButton, ok := found.(*Button)
	if !ok {
		t.Error("Did not return a button. Received type:", reflect.TypeOf(found))
	}

	text, _ := foundAsButton.GetLabel()
	if text != "Second" {
		t.Error("Expected: Second, Got:", text)
	}

}
