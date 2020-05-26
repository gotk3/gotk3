gotk3 [![GoDoc](https://godoc.org/github.com/gotk3/gotk3?status.svg)](https://godoc.org/github.com/gotk3/gotk3)
=====

[![Build Status](https://travis-ci.org/gotk3/gotk3.svg?branch=master)](https://travis-ci.org/gotk3/gotk3)

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
