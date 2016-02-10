package iface

import "unsafe"

type List interface {
    Append(uintptr) List
    Data() interface{}
    DataWrapper(func(unsafe.Pointer) interface{})
    Foreach(func(interface{}))
    Free()
    FreeFull(func(interface{}))
    Insert(uintptr, int) List
    Length() uint
    Next() List
    Nth(uint) List
    NthData(uint) interface{}
    Prepend(uintptr) List
    Previous() List
} // end of List

func AssertList(_ List) {}
