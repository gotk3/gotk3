package glib

type SList interface {
	Append(uintptr) SList
} // end of SList

func AssertSList(_ SList) {}