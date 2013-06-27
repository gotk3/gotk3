gotk3
=====

Package gotk3 provides Go bindings for GLib 2, GDK 3, and GTK+3.  Each
component is given its own subdirectory, which is used as the import
path for the package.

Care has been taken for memory management to work seamlessly with Go's
garbage collector without the need to use or understand GObject's
floating references.

## Documentation

Each package's internal `go doc` style documentation can be viewed
online without installing this package by using the GoDoc site (links
to [glib](http://godoc.org/github.com/conformal/gotk3/glib),
[gdk](http://godoc.org/github.com/conformal/gotk3/gdk), and
[gtk](http://godoc.org/github.com/conformal/gotk3/gtk) documentation).

You can also view the documentation locally once the package is
installed with the `godoc` tool by running `godoc -http=":6060"` and
pointing your browser to
http://localhost:6060/pkg/github.com/conformal/gotk3

## Installation

The gtk package requires glib and gdk packages as dependencies, so
only one `go get` is necessary for complete installation.

```bash
$ go get github.com/conformal/gotk3/gtk
```

## TODO
- Add bindings for all of GTK+
- Add tests for each implemented binding
- Add examples for intent

## License

Package gotk3 is licensed under the liberal ISC License.
