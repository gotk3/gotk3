package iface

import "unsafe"

type Value interface {
    GetPointer() unsafe.Pointer
    GetString() (string, error)
    GoValue() (interface{}, error)
    SetBool(bool)
    SetDouble(float64)
    SetFloat(float32)
    SetInstance(uintptr)
    SetInt(int)
    SetInt64(int64)
    SetPointer(uintptr)
    SetSChar(int8)
    SetString(string)
    SetUChar(uint8)
    SetUInt(uint)
    SetUInt64(uint64)
    Type() (Type, Type, error)
} // end of Value

func AssertValue(_ Value) {}
