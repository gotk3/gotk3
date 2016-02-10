package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type EntryBuffer interface {
    glib_iface.Object

    DeleteText(uint, int) uint
    EmitDeletedText(uint, uint)
    EmitInsertedText(uint, string)
    GetBytes() uint
    GetLength() uint
    GetMaxLength() int
    GetText() (string, error)
    InsertText(uint, string) uint
    SetMaxLength(int)
    SetText(string)
} // end of EntryBuffer

func AssertEntryBuffer(_ EntryBuffer) {}
