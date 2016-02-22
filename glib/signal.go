package glib

type Signal interface {
	String() string
} // end of Signal

func AssertSignal(_ Signal) {}
