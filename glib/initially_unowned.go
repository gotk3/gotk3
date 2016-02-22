package glib

type InitiallyUnowned interface {
	Object
} // end of InitiallyUnowned

func AssertInitiallyUnowned(_ InitiallyUnowned) {}
