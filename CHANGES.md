---

### User visible changes for gotk3 Go bindings for GTK3

---

Changes for Version after 0.6.1:

* **2021-08**: Glib version 2.68 deprecated glib.Binding. **GetSource** and **GetTarget** in favor of **DupSource** and **DupTarget**. Those using glib.Binding should check the [glib changes](https://gitlab.gnome.org/GNOME/glib/-/tags/2.67.1). For those who use **_Glib versions <= 2.66_**, you now need to use the build tag `-tags "glib_2_66"`, see [#828](https://github.com/gotk3/gotk3/pull/828)



Changes for next Version 0.6.0

- Breaking changes in API
- General code cleanup 
- #685 Refactor Gtk callback setters and types enhancement missing binding
- #706 Refactor internal closure handling and several API changes breaking changes
- #746 Add build tag pango_1_42 for Pango
- #743 Solving #741- Add possibility to use GVariant in signal handler
- #740 Add binding for GtkRadioMenuItem
- #738 Adds binding for gtk_cell_layout_clear_attributes()
- #737 Adds bindings for gdk_pixbuf_new_from_resource() and gdk_pixbuf_new_from_resource_at_scale()
- #736 Add bindings/helper methods GdkRectangle GdkPoint
- #735 Add GtkMenuItem bindings
- #734 Add bindings GtkMenuShell
- #732 add as contributor
- #731 add bindings to GtkMenu
- #730 Solve GtkAccelKey issue with golang 1.16
- #728 It is not safe to reference memory returned in a signal callback.
- #687 Don't merge until publication of Golang v1.16: GtkAccelKey v1.16 issue fix next version
- #724 Implemented CellRenderer.SetAlignment
- #723 Added SetOrientation to gkt.SpinButton
- #720 Add Prgname getter and setter
- #716 Add (Get/Set) methods to GdkRGBA & GdkVisual & GdkDisplayManager bind…
- #715 Add some GtkRange bindings
- #712 glib.Take to return nil and gtk.marshal* to allow nil
