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

	i := 0
	for l := list; l != nil; l = l.Next() {
		expect := fmt.Sprintf("Value %v", i)
		i++
		actual, ok := l.Data().(string)
		if !ok {
			t.Error("DataWrapper must have returned a string!")
		}
		if actual != expect {
			t.Errorf("DataWrapper returned unexpected result. Expected '%v', got '%v'.", expect, actual)
		}
	}
}

func TestList_Foreach(t *testing.T) {
	list := (&List{}).Append(0).Append(1).Append(2)
	list.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		return int(uintptr(ptr) + 1)
	})

	sum := 0
	list.Foreach(func(item interface{}) {
		sum += item.(int)
	})

	if sum != 6 {
		t.Errorf("Foreach resulted into wrong sum. Got %v, expected %v.", sum, 6)
	}
}

func TestList_Nth(t *testing.T) {
	list := (&List{}).Append(0).Append(1).Append(2)
	list.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		return int(uintptr(ptr) + 1)
	})

	for i := uint(0); i < 3; i++ {
		nth := list.Nth(i).Data().(int)
		nthData := list.NthData(i).(int)

		if nth != nthData {
			t.Errorf("%v's element didn't match. Nth->Data returned %v; NthData returned %v.", i, nth, nthData)
		}
	}
}
