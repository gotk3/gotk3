package iface

type GtkSince310 interface {
	Gtk

	RevealerNew() (Revealer, error)
}

func AssertGtkSince310(_ GtkSince310) {}
