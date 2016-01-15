package glib

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestList_Basics(t *testing.T) {
	list := (&List{}).Append(0).Append(1).Append(2)
	if list.Length() != 3 {
		t.Errorf("Length of list with 3 appended elements must be 3. (Got %v).", list.Length())
	}

	list = (&List{}).Prepend(0).Prepend(1).Prepend(2)
	if list.Length() != 3 {
		t.Errorf("Length of list with 3 prepended elements must be 3. (Got %v).", list.Length())
	}

	list = (&List{}).Insert(0, 0).Insert(1, 0).Insert(2, 0)
	if list.Length() != 3 {
		t.Errorf("Length of list with 3 inserted elements must be 3. (Got %v).", list.Length())
	}
}

func TestList_DataWrapper(t *testing.T) {
	list := (&List{}).Append(0).Append(1).Append(2)
	list.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		return fmt.Sprintf("Value %v", uintptr(ptr))
	})

	for l := list; l != nil; l = l.Next() {
		expect := fmt.Sprintf("Value %v", uintptr(l.Data()))
		actual, ok := l.DataWrapped().(string)
		if !ok {
			t.Error("DataWrapper must have returned a string!")
		}
		if actual != expect {
			t.Errorf("DataWrapper returned unexpected result. Expected '%v', got '%v'.", expect, actual)
		}
	}
}
