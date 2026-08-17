gotk3 (fork) [![GoDoc](https://godoc.org/github.com/gotk3/gotk3?status.svg)](https://godoc.org/github.com/gotk3/gotk3)
=====

[![Build Status](https://travis-ci.org/gotk3/gotk3.svg?branch=master)](https://travis-ci.org/gotk3/gotk3)

This fork adds a couple of features and bugfixes, but its main purpose is to
modify the upstream project in such a way that it is possible to compile
programs using this library for Windows XP. This means the following:

- GTK 3.6 is supported (compiles) if the build tag `gtk_3_6` is used.
- GDK 3.6 is supported (compiles) if the build tag `gtk_3_6` is used.
- GLib 2.34 is supported with the build tag `glib_2_34`
- Cairo 1.10.2 is supported with the build tag `cairo_1_10`. The upstream repository already
claims 1.10 support, however there are some 1.12 features (for example CAIRO_ANTIALIAS_GOOD)
that are in untagged files. This fork fixes that, meaning that it is possible to compile this
on a system with Cairo 1.10.2

## Important information

- to compile an application for Windows XP, you need the following build tags:
  `gtk_3_6 glib_2_34 cairo_1_10 pango_1_36 gdk_pixbuf_2_26`

  You will probably need a bundle for this. You can get it
  [here](https://web.archive.org/web/20180228011133/http://win32builder.gnome.org/gtk+-bundle_3.6.4-20130921_win32.zip)

  The SHA256 hash of this bundle is: `1c1a6019746ae58c6a112e59c6b50463b838900d2244c494d94dea8ebea0f715`

- When using the `glib_2_34` build tag, the GLib package must be initialized
  before any initialization code in the other gotk3 packages that uses the
  GObject type system. Prior to GLib 2.36, `g_type_init()` had to be called
  explicitly to initialize the type system.

  Since gotk3 uses private `init()` functions in several packages, this fork
  initializes the GObject type system from a package-level variable in the
  `glib` package. Package-level variables are initialized before `init()`
  functions, and imported packages are initialized before packages that depend
  on them. This ensures that `g_type_init()` is called before the initialization
  code of gotk3 packages that depend on `glib`.

  Since the library already calls this function internally when the package-level
  variable is initialized, no extra code should be required.

- Pango 1.30 is supported. `The pango_1_36` build tag can be used with Pango 1.30
  because the only API exposed between these compatibility levels that is not present
  in Pango 1.30 is `WEIGHT_SEMILIGHT`, which is represented by its numeric value (350)
  and does not reference a newer Pango C symbol.

# The upstream gotk3 README

The gotk3 project provides Go bindings for GTK 3 and dependent
projects.  Each component is given its own subdirectory, which is used
as the import path for the package.  Partial binding support for the
following libraries is currently implemented:

- GTK 3 (3.12 and later)
- GDK 3 (3.12 and later)
- GLib 2 (2.36 and later)
- Cairo (1.10 and later)

Care has been taken for memory management to work seamlessly with Go's
garbage collector without the need to use or understand GObject's
floating references.

for better understanding see
[package reference documation](https://pkg.go.dev/github.com/gotk3/gotk3/gtk?tab=doc)

On Linux, see which version your distribution has [here](https://pkgs.org) with the search terms:
* libgtk-3
* libglib2
* libgdk-pixbuf2

## Sample Use

The following example can be found in [Examples](https://github.com/gotk3/gotk3-examples/).

```Go
package main

import (
    "github.com/gotk3/gotk3/gtk"
    "log"
)

func main() {
    // Initialize GTK without parsing any command line arguments.
    gtk.Init(nil)

    // Create a new toplevel window, set its title, and connect it to the
    // "destroy" signal to exit the GTK main loop when it is destroyed.
    win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
    if err != nil {
        log.Fatal("Unable to create window:", err)
    }
    win.SetTitle("Simple Example")
    win.Connect("destroy", func() {
        gtk.MainQuit()
    })

    // Create a new label widget to show in the window.
    l, err := gtk.LabelNew("Hello, gotk3!")
    if err != nil {
        log.Fatal("Unable to create label:", err)
    }

    // Add the label to the window.
    win.Add(l)

    // Set the default window size.
    win.SetDefaultSize(800, 600)

    // Recursively show all widgets contained in this window.
    win.ShowAll()

    // Begin executing the GTK main loop.  This blocks until
    // gtk.MainQuit() is run.
    gtk.Main()
}
```

To build the example:

```shell
$ go build example.go
```

To build this example with older gtk version you should use gtk_3_10 tag:

```shell
$ go build -tags gtk_3_10 example.go
```

### Example usage

```Go
package main

import (
    "log"
    "os"

    "github.com/gotk3/gotk3/glib"
    "github.com/gotk3/gotk3/gtk"
)

// Simple Gtk3 Application written in go.
// This application creates a window on the application callback activate.
// More GtkApplication info can be found here -> https://wiki.gnome.org/HowDoI/GtkApplication

func main() {
    // Create Gtk Application, change appID to your application domain name reversed.
    const appID = "org.gtk.example"
    application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
    // Check to make sure no errors when creating Gtk Application
    if err != nil {
        log.Fatal("Could not create application.", err)
    }
    // Application signals available
    // startup -> sets up the application when it first starts
    // activate -> shows the default first window of the application (like a new document). This corresponds to the application being launched by the desktop environment.
    // open -> opens files and shows them in a new window. This corresponds to someone trying to open a document (or documents) using the application from the file browser, or similar.
    // shutdown ->  performs shutdown tasks
    // Setup Gtk Application callback signals
    application.Connect("activate", func() { onActivate(application) })
    // Run Gtk application
    os.Exit(application.Run(os.Args))
}

// Callback signal from Gtk Application
func onActivate(application *gtk.Application) {
    // Create ApplicationWindow
    appWindow, err := gtk.ApplicationWindowNew(application)
    if err != nil {
        log.Fatal("Could not create application window.", err)
    }
    // Set ApplicationWindow Properties
    appWindow.SetTitle("Basic Application.")
    appWindow.SetDefaultSize(400, 400)
    appWindow.Show()
}
```

```Go
package main

import (
    "log"
    "os"

    "github.com/gotk3/gotk3/glib"
    "github.com/gotk3/gotk3/gtk"
)

// Simple Gtk3 Application written in go.
// This application creates a window on the application callback activate.
// More GtkApplication info can be found here -> https://wiki.gnome.org/HowDoI/GtkApplication

func main() {
    // Create Gtk Application, change appID to your application domain name reversed.
    const appID = "org.gtk.example"
    application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
    // Check to make sure no errors when creating Gtk Application
    if err != nil {
        log.Fatal("Could not create application.", err)
    }

    // Application signals available
    // startup -> sets up the application when it first starts
    // activate -> shows the default first window of the application (like a new document). This corresponds to the application being launched by the desktop environment.
    // open -> opens files and shows them in a new window. This corresponds to someone trying to open a document (or documents) using the application from the file browser, or similar.
    // shutdown ->  performs shutdown tasks
    // Setup activate signal with a closure function.
    application.Connect("activate", func() {
        // Create ApplicationWindow
        appWindow, err := gtk.ApplicationWindowNew(application)
        if err != nil {
            log.Fatal("Could not create application window.", err)
        }
        // Set ApplicationWindow Properties
        appWindow.SetTitle("Basic Application.")
        appWindow.SetDefaultSize(400, 400)
        appWindow.Show()
    })
    // Run Gtk application
    application.Run(os.Args)
}
```

## Documentation

Each package's internal `go doc` style documentation can be viewed
online without installing this package by using the GoDoc site (links
to [cairo](http://godoc.org/github.com/gotk3/gotk3/cairo),
[glib](http://godoc.org/github.com/gotk3/gotk3/glib),
[gdk](http://godoc.org/github.com/gotk3/gotk3/gdk), and
[gtk](http://godoc.org/github.com/gotk3/gotk3/gtk) documentation).

You can also view the documentation locally once the package is
installed with the `godoc` tool by running `godoc -http=":6060"` and
pointing your browser to
http://localhost:6060/pkg/github.com/gotk3/gotk3

## Installation

gotk3 currently requires GTK 3.6-3.24, GLib 2.36-2.46, and
Cairo 1.10 or 1.12.  A recent Go (1.8 or newer) is also required.

For detailed instructions see the wiki pages: [installation](https://github.com/gotk3/gotk3/wiki#installation)

## Using deprecated features

By default, deprecated GTK features are not included in the build.

By specifying the e.g. build tag `gtk_3_20`, any feature deprecated in GTK 3.20 or earlier will NOT be available.
To enable deprecated features in the build, add the tag `gtk_deprecated`.
Example:
```shell
$ go build -tags "gtk_3_10 gtk_deprecated" example.go
```

The same goes for
* gdk-pixbuf: gdk_pixbuf_deprecated

## TODO

- Add bindings for all of GTK functions
- Add tests for each implemented binding
- See the next steps: [wiki page](https://github.com/gotk3/gotk3/wiki/The-future-and-what-happens-next) and add [your suggestion](https://github.com/gotk3/gotk3/issues/576)


## License

Package gotk3 is licensed under the liberal ISC License.

Actually if you use gotk3, then gotk3 is statically linked into your application (with the ISC licence).
The system libraries (e.g. GTK+, GLib) used via cgo use dynamic linking.
