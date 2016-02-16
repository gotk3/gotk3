package impl

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

import cairo_impl "github.com/gotk3/gotk3/cairo/impl"
import gdk_impl "github.com/gotk3/gotk3/gdk/impl"
import glib_impl "github.com/gotk3/gotk3/glib/impl"

type RealGtk struct{}

var Real = &RealGtk{}

func (*RealGtk) AboutDialogNew() (gtk.AboutDialog, error) {
	return AboutDialogNew()
}

func (*RealGtk) AccelGroupFromClosure(f interface{}) gtk.AccelGroup {
	return AccelGroupFromClosure(f)
}

func (*RealGtk) AccelGroupNew() (gtk.AccelGroup, error) {
	return AccelGroupNew()
}

func (*RealGtk) AccelGroupsActivate(obj glib.Object, key uint, mods gdk.ModifierType) bool {
	return AccelGroupsActivate(obj.(*glib_impl.Object), key, mods)
}

func (*RealGtk) AccelGroupsFromObject(obj glib.Object) glib.SList {
	return AccelGroupsFromObject(obj.(*glib_impl.Object))
}

func (*RealGtk) AccelMapAddEntry(path string, key uint, mods gdk.ModifierType) {
	AccelMapAddEntry(path, key, mods)
}

func (*RealGtk) AccelMapAddFilter(filter string) {
	AccelMapAddFilter(filter)
}

func (*RealGtk) AccelMapChangeEntry(path string, key uint, mods gdk.ModifierType, replace bool) bool {
	return AccelMapChangeEntry(path, key, mods, replace)
}

func (*RealGtk) AccelMapGet() gtk.AccelMap {
	return AccelMapGet()
}

func (*RealGtk) AccelMapLoad(fileName string) {
	AccelMapLoad(fileName)
}

func (*RealGtk) AccelMapLoadFD(fd int) {
	AccelMapLoadFD(fd)
}

func (*RealGtk) AccelMapLockPath(path string) {
	AccelMapLockPath(path)
}

func (*RealGtk) AccelMapLookupEntry(path string) gtk.AccelKey {
	return AccelMapLookupEntry(path)
}

func (*RealGtk) AccelMapSave(fileName string) {
	AccelMapSave(fileName)
}

func (*RealGtk) AccelMapSaveFD(fd int) {
	AccelMapSaveFD(fd)
}

func (*RealGtk) AccelMapUnlockPath(path string) {
	AccelMapUnlockPath(path)
}

func (*RealGtk) AcceleratorGetDefaultModMask() gdk.ModifierType {
	return AcceleratorGetDefaultModMask()
}

func (*RealGtk) AcceleratorGetLabel(key uint, mods gdk.ModifierType) string {
	return AcceleratorGetLabel(key, mods)
}

func (*RealGtk) AcceleratorName(key uint, mods gdk.ModifierType) string {
	return AcceleratorName(key, mods)
}

func (*RealGtk) AcceleratorParse(acc string) (uint, gdk.ModifierType) {
	return AcceleratorParse(acc)
}

func (*RealGtk) AcceleratorSetDefaultModMask(mods gdk.ModifierType) {
	AcceleratorSetDefaultModMask(mods)
}

func (*RealGtk) AcceleratorValid(key uint, mods gdk.ModifierType) bool {
	return AcceleratorValid(key, mods)
}

func (*RealGtk) AddProviderForScreen(s gdk.Screen, provider gtk.StyleProvider, prio uint) {
	AddProviderForScreen(s.(*gdk_impl.Screen), provider, prio)
}

func (*RealGtk) AdjustmentNew(value float64, lower float64, upper float64, stepIncrement float64, pageIncrement float64, pageSize float64) (gtk.Adjustment, error) {
	return AdjustmentNew(value, lower, upper, stepIncrement, pageIncrement, pageSize)
}

func (*RealGtk) AppChooserButtonNew(content_type string) (gtk.AppChooserButton, error) {
	return AppChooserButtonNew(content_type)
}

func (*RealGtk) AppChooserDialogNewForContentType(parent gtk.Window, flags gtk.DialogFlags, content_type string) (gtk.AppChooserDialog, error) {
	return AppChooserDialogNewForContentType(asWindowImpl(parent), flags, content_type)
}

func (*RealGtk) AppChooserWidgetNew(content_type string) (gtk.AppChooserWidget, error) {
	return AppChooserWidgetNew(content_type)
}

func (*RealGtk) ApplicationNew(appId string, flags glib.ApplicationFlags) (gtk.Application, error) {
	return ApplicationNew(appId, flags)
}

func (*RealGtk) ApplicationWindowNew(app gtk.Application) (gtk.ApplicationWindow, error) {
	return ApplicationWindowNew(app.(*Application))
}

func (*RealGtk) AssistantNew() (gtk.Assistant, error) {
	return AssistantNew()
}

func (*RealGtk) BoxNew(orientation gtk.Orientation, spacing int) (gtk.Box, error) {
	return BoxNew(orientation, spacing)
}

func (*RealGtk) BuilderNew() (gtk.Builder, error) {
	return BuilderNew()
}

func (*RealGtk) ButtonNew() (gtk.Button, error) {
	return ButtonNew()
}

func (*RealGtk) ButtonNewWithLabel(label string) (gtk.Button, error) {
	return ButtonNewWithLabel(label)
}

func (*RealGtk) ButtonNewWithMnemonic(label string) (gtk.Button, error) {
	return ButtonNewWithMnemonic(label)
}

func (*RealGtk) CalendarNew() (gtk.Calendar, error) {
	return CalendarNew()
}

func (*RealGtk) CellRendererPixbufNew() (gtk.CellRendererPixbuf, error) {
	return CellRendererPixbufNew()
}

func (*RealGtk) CellRendererSpinnerNew() (gtk.CellRendererSpinner, error) {
	return CellRendererSpinnerNew()
}

func (*RealGtk) CellRendererTextNew() (gtk.CellRendererText, error) {
	return CellRendererTextNew()
}

func (*RealGtk) CellRendererToggleNew() (gtk.CellRendererToggle, error) {
	return CellRendererToggleNew()
}

func (*RealGtk) CheckButtonNew() (gtk.CheckButton, error) {
	return CheckButtonNew()
}

func (*RealGtk) CheckButtonNewWithLabel(label string) (gtk.CheckButton, error) {
	return CheckButtonNewWithLabel(label)
}

func (*RealGtk) CheckButtonNewWithMnemonic(label string) (gtk.CheckButton, error) {
	return CheckButtonNewWithMnemonic(label)
}

func (*RealGtk) CheckMenuItemNew() (gtk.CheckMenuItem, error) {
	return CheckMenuItemNew()
}

func (*RealGtk) CheckMenuItemNewWithLabel(label string) (gtk.CheckMenuItem, error) {
	return CheckMenuItemNewWithLabel(label)
}

func (*RealGtk) CheckMenuItemNewWithMnemonic(label string) (gtk.CheckMenuItem, error) {
	return CheckMenuItemNewWithMnemonic(label)
}

func (*RealGtk) CheckVersion(major uint, minor uint, micro uint) error {
	return CheckVersion(major, minor, micro)
}

func (*RealGtk) ClipboardGet(atom gdk.Atom) (gtk.Clipboard, error) {
	return ClipboardGet(atom)
}

func (*RealGtk) ClipboardGetForDisplay(display gdk.Display, atom gdk.Atom) (gtk.Clipboard, error) {
	return ClipboardGetForDisplay(display.(*gdk_impl.Display), atom)
}

func (*RealGtk) ColorButtonNew() (gtk.ColorButton, error) {
	return ColorButtonNew()
}

func (*RealGtk) ColorButtonNewWithRGBA(gdkColor gdk.RGBA) (gtk.ColorButton, error) {
	return ColorButtonNewWithRGBA(gdkColor.(*gdk_impl.RGBA))
}

func (*RealGtk) ColorChooserDialogNew(title string, parent gtk.Window) (gtk.ColorChooserDialog, error) {
	return ColorChooserDialogNew(title, asWindowImpl(parent))
}

func (*RealGtk) ComboBoxNew() (gtk.ComboBox, error) {
	return ComboBoxNew()
}

func (*RealGtk) ComboBoxNewWithEntry() (gtk.ComboBox, error) {
	return ComboBoxNewWithEntry()
}

func (*RealGtk) ComboBoxNewWithModel(model gtk.TreeModel) (gtk.ComboBox, error) {
	return ComboBoxNewWithModel(model.(ITreeModel))
}

func (*RealGtk) ComboBoxTextNew() (gtk.ComboBoxText, error) {
	return ComboBoxTextNew()
}

func (*RealGtk) ComboBoxTextNewWithEntry() (gtk.ComboBoxText, error) {
	return ComboBoxTextNewWithEntry()
}

func (*RealGtk) CssProviderGetDefault() (gtk.CssProvider, error) {
	return CssProviderGetDefault()
}

func (*RealGtk) CssProviderGetNamed(name string, variant string) (gtk.CssProvider, error) {
	return CssProviderGetNamed(name, variant)
}

func (*RealGtk) CssProviderNew() (gtk.CssProvider, error) {
	return CssProviderNew()
}

func (*RealGtk) DialogNew() (gtk.Dialog, error) {
	return DialogNew()
}

func (*RealGtk) DrawingAreaNew() (gtk.DrawingArea, error) {
	return DrawingAreaNew()
}

func (*RealGtk) EntryBufferNew(initialChars string, nInitialChars int) (gtk.EntryBuffer, error) {
	return EntryBufferNew(initialChars, nInitialChars)
}

func (*RealGtk) EntryNew() (gtk.Entry, error) {
	return EntryNew()
}

func (*RealGtk) EntryNewWithBuffer(buffer gtk.EntryBuffer) (gtk.Entry, error) {
	return EntryNewWithBuffer(buffer.(*EntryBuffer))
}

func (*RealGtk) EventBoxNew() (gtk.EventBox, error) {
	return EventBoxNew()
}

func (*RealGtk) EventsPending() bool {
	return EventsPending()
}

func (*RealGtk) ExpanderNew(label string) (gtk.Expander, error) {
	return ExpanderNew(label)
}

func (*RealGtk) FileChooserButtonNew(title string, action gtk.FileChooserAction) (gtk.FileChooserButton, error) {
	return FileChooserButtonNew(title, action)
}

func (*RealGtk) FileChooserDialogNewWith1Button(title string, parent gtk.Window, action gtk.FileChooserAction, first_button_text string, first_button_id gtk.ResponseType) (gtk.FileChooserDialog, error) {
	return FileChooserDialogNewWith1Button(title, asWindowImpl(parent), action, first_button_text, first_button_id)
}

func (*RealGtk) FileChooserDialogNewWith2Buttons(title string, parent gtk.Window, action gtk.FileChooserAction, first_button_text string, first_button_id gtk.ResponseType, second_button_text string, second_button_id gtk.ResponseType) (gtk.FileChooserDialog, error) {
	return FileChooserDialogNewWith2Buttons(title, asWindowImpl(parent), action, first_button_text, first_button_id, second_button_text, second_button_id)
}

func (*RealGtk) FileChooserWidgetNew(action gtk.FileChooserAction) (gtk.FileChooserWidget, error) {
	return FileChooserWidgetNew(action)
}

func (*RealGtk) FileFilterNew() (gtk.FileFilter, error) {
	return FileFilterNew()
}

func (*RealGtk) FontButtonNew() (gtk.FontButton, error) {
	return FontButtonNew()
}

func (*RealGtk) FontButtonNewWithFont(fontname string) (gtk.FontButton, error) {
	return FontButtonNewWithFont(fontname)
}

func (*RealGtk) FrameNew(label string) (gtk.Frame, error) {
	return FrameNew(label)
}

func (*RealGtk) GdkCairoSetSourcePixBuf(cr cairo.Context, pixbuf gdk.Pixbuf, pixbufX float64, pixbufY float64) {
	GdkCairoSetSourcePixBuf(cr.(*cairo_impl.Context), pixbuf.(*gdk_impl.Pixbuf), pixbufX, pixbufY)
}

func (*RealGtk) GetMajorVersion() uint {
	return GetMajorVersion()
}

func (*RealGtk) GetMicroVersion() uint {
	return GetMicroVersion()
}

func (*RealGtk) GetMinorVersion() uint {
	return GetMinorVersion()
}

func (*RealGtk) GridNew() (gtk.Grid, error) {
	return GridNew()
}

func (*RealGtk) IconThemeGetDefault() (gtk.IconTheme, error) {
	return IconThemeGetDefault()
}

func (*RealGtk) IconThemeGetForScreen(screen gdk.Screen) (gtk.IconTheme, error) {
	return IconThemeGetForScreen(screen.(*gdk_impl.Screen))
}

func (*RealGtk) IconViewNew() (gtk.IconView, error) {
	return IconViewNew()
}

func (*RealGtk) IconViewNewWithModel(model gtk.TreeModel) (gtk.IconView, error) {
	return IconViewNewWithModel(model.(ITreeModel))
}

func (*RealGtk) ImageNew() (gtk.Image, error) {
	return ImageNew()
}

func (*RealGtk) ImageNewFromFile(filename string) (gtk.Image, error) {
	return ImageNewFromFile(filename)
}

func (*RealGtk) ImageNewFromIconName(iconName string, size gtk.IconSize) (gtk.Image, error) {
	return ImageNewFromIconName(iconName, size)
}

func (*RealGtk) ImageNewFromPixbuf(pixbuf gdk.Pixbuf) (gtk.Image, error) {
	return ImageNewFromPixbuf(pixbuf.(*gdk_impl.Pixbuf))
}

func (*RealGtk) ImageNewFromResource(resourcePath string) (gtk.Image, error) {
	return ImageNewFromResource(resourcePath)
}

func (*RealGtk) InfoBarNew() (gtk.InfoBar, error) {
	return InfoBarNew()
}

func (*RealGtk) Init(args *[]string) {
	Init(args)
}

func (*RealGtk) LabelNew(str string) (gtk.Label, error) {
	return LabelNew(str)
}

func (*RealGtk) LabelNewWithMnemonic(str string) (gtk.Label, error) {
	return LabelNewWithMnemonic(str)
}

func (*RealGtk) LayoutNew(hadjustment gtk.Adjustment, vadjustment gtk.Adjustment) (gtk.Layout, error) {
	return LayoutNew(hadjustment.(*Adjustment), vadjustment.(*Adjustment))
}

func (*RealGtk) LevelBarNew() (gtk.LevelBar, error) {
	return LevelBarNew()
}

func (*RealGtk) LevelBarNewForInterval(min_value float64, max_value float64) (gtk.LevelBar, error) {
	return LevelBarNewForInterval(min_value, max_value)
}

func (*RealGtk) LinkButtonNew(label string) (gtk.LinkButton, error) {
	return LinkButtonNew(label)
}

func (*RealGtk) LinkButtonNewWithLabel(uri string, label string) (gtk.LinkButton, error) {
	return LinkButtonNewWithLabel(uri, label)
}

func (*RealGtk) ListStoreNew(types ...glib.Type) (gtk.ListStore, error) {
	return ListStoreNew(types...)
}

func (*RealGtk) Main() {
	Main()
}

func (*RealGtk) MainIteration() bool {
	return MainIteration()
}

func (*RealGtk) MainIterationDo(blocking bool) bool {
	return MainIterationDo(blocking)
}

func (*RealGtk) MainQuit() {
	MainQuit()
}

func (*RealGtk) MenuBarNew() (gtk.MenuBar, error) {
	return MenuBarNew()
}

func (*RealGtk) MenuButtonNew() (gtk.MenuButton, error) {
	return MenuButtonNew()
}

func (*RealGtk) MenuItemNew() (gtk.MenuItem, error) {
	return MenuItemNew()
}

func (*RealGtk) MenuItemNewWithLabel(label string) (gtk.MenuItem, error) {
	return MenuItemNewWithLabel(label)
}

func (*RealGtk) MenuItemNewWithMnemonic(label string) (gtk.MenuItem, error) {
	return MenuItemNewWithMnemonic(label)
}

func (*RealGtk) MenuNew() (gtk.Menu, error) {
	return MenuNew()
}

func (*RealGtk) MessageDialogNew(parent gtk.Window, flags gtk.DialogFlags, mType gtk.MessageType, buttons gtk.ButtonsType, format string, a ...interface{}) gtk.MessageDialog {
	return MessageDialogNew(parent.(IWindow), flags, mType, buttons, format, a...)
}

func (*RealGtk) MessageDialogNewWithMarkup(parent gtk.Window, flags gtk.DialogFlags, mType gtk.MessageType, buttons gtk.ButtonsType, format string, a ...interface{}) gtk.MessageDialog {
	return MessageDialogNewWithMarkup(parent.(IWindow), flags, mType, buttons, format, a...)
}

func (*RealGtk) NotebookNew() (gtk.Notebook, error) {
	return NotebookNew()
}

func (*RealGtk) OffscreenWindowNew() (gtk.OffscreenWindow, error) {
	return OffscreenWindowNew()
}

func (*RealGtk) PanedNew(orientation gtk.Orientation) (gtk.Paned, error) {
	return PanedNew(orientation)
}

func (*RealGtk) ProgressBarNew() (gtk.ProgressBar, error) {
	return ProgressBarNew()
}

func (*RealGtk) RadioButtonNew(group glib.SList) (gtk.RadioButton, error) {
	return RadioButtonNew(group.(*glib_impl.SList))
}

func (*RealGtk) RadioButtonNewFromWidget(radioGroupMember gtk.RadioButton) (gtk.RadioButton, error) {
	return RadioButtonNewFromWidget(radioGroupMember.(*RadioButton))
}

func (*RealGtk) RadioButtonNewWithLabel(group glib.SList, label string) (gtk.RadioButton, error) {
	return RadioButtonNewWithLabel(group.(*glib_impl.SList), label)
}

func (*RealGtk) RadioButtonNewWithLabelFromWidget(radioGroupMember gtk.RadioButton, label string) (gtk.RadioButton, error) {
	return RadioButtonNewWithLabelFromWidget(radioGroupMember.(*RadioButton), label)
}

func (*RealGtk) RadioButtonNewWithMnemonic(group glib.SList, label string) (gtk.RadioButton, error) {
	return RadioButtonNewWithMnemonic(group.(*glib_impl.SList), label)
}

func (*RealGtk) RadioButtonNewWithMnemonicFromWidget(radioGroupMember gtk.RadioButton, label string) (gtk.RadioButton, error) {
	return RadioButtonNewWithMnemonicFromWidget(radioGroupMember.(*RadioButton), label)
}

func (*RealGtk) RadioMenuItemNew(group glib.SList) (gtk.RadioMenuItem, error) {
	return RadioMenuItemNew(group.(*glib_impl.SList))
}

func (*RealGtk) RadioMenuItemNewFromWidget(group gtk.RadioMenuItem) (gtk.RadioMenuItem, error) {
	return RadioMenuItemNewFromWidget(group.(*RadioMenuItem))
}

func (*RealGtk) RadioMenuItemNewWithLabel(group glib.SList, label string) (gtk.RadioMenuItem, error) {
	return RadioMenuItemNewWithLabel(group.(*glib_impl.SList), label)
}

func (*RealGtk) RadioMenuItemNewWithLabelFromWidget(group gtk.RadioMenuItem, label string) (gtk.RadioMenuItem, error) {
	return RadioMenuItemNewWithLabelFromWidget(group.(*RadioMenuItem), label)
}

func (*RealGtk) RadioMenuItemNewWithMnemonic(group glib.SList, label string) (gtk.RadioMenuItem, error) {
	return RadioMenuItemNewWithMnemonic(group.(*glib_impl.SList), label)
}

func (*RealGtk) RadioMenuItemNewWithMnemonicFromWidget(group gtk.RadioMenuItem, label string) (gtk.RadioMenuItem, error) {
	return RadioMenuItemNewWithMnemonicFromWidget(group.(*RadioMenuItem), label)
}

func (*RealGtk) RecentFilterNew() (gtk.RecentFilter, error) {
	return RecentFilterNew()
}

func (*RealGtk) RecentManagerGetDefault() (gtk.RecentManager, error) {
	return RecentManagerGetDefault()
}

func (*RealGtk) RemoveProviderForScreen(s gdk.Screen, provider gtk.StyleProvider) {
	RemoveProviderForScreen(s.(*gdk_impl.Screen), provider)
}

func (*RealGtk) ScaleButtonNew(size gtk.IconSize, min float64, max float64, step float64, icons []string) (gtk.ScaleButton, error) {
	return ScaleButtonNew(size, min, max, step, icons)
}

func (*RealGtk) ScaleNew(orientation gtk.Orientation, adjustment gtk.Adjustment) (gtk.Scale, error) {
	return ScaleNew(orientation, adjustment.(*Adjustment))
}

func (*RealGtk) ScaleNewWithRange(orientation gtk.Orientation, min float64, max float64, step float64) (gtk.Scale, error) {
	return ScaleNewWithRange(orientation, min, max, step)
}

func (*RealGtk) ScrollbarNew(orientation gtk.Orientation, adjustment gtk.Adjustment) (gtk.Scrollbar, error) {
	return ScrollbarNew(orientation, adjustment.(*Adjustment))
}

func (*RealGtk) ScrolledWindowNew(hadjustment gtk.Adjustment, vadjustment gtk.Adjustment) (gtk.ScrolledWindow, error) {
	return ScrolledWindowNew(hadjustment.(*Adjustment), vadjustment.(*Adjustment))
}

func (*RealGtk) SearchEntryNew() (gtk.SearchEntry, error) {
	return SearchEntryNew()
}

func (*RealGtk) SeparatorMenuItemNew() (gtk.SeparatorMenuItem, error) {
	return SeparatorMenuItemNew()
}

func (*RealGtk) SeparatorNew(orientation gtk.Orientation) (gtk.Separator, error) {
	return SeparatorNew(orientation)
}

func (*RealGtk) SeparatorToolItemNew() (gtk.SeparatorToolItem, error) {
	return SeparatorToolItemNew()
}

func (*RealGtk) SettingsGetDefault() (gtk.Settings, error) {
	return SettingsGetDefault()
}

func (*RealGtk) SpinButtonNew(adjustment gtk.Adjustment, climbRate float64, digits uint) (gtk.SpinButton, error) {
	return SpinButtonNew(adjustment.(*Adjustment), climbRate, digits)
}

func (*RealGtk) SpinButtonNewWithRange(min float64, max float64, step float64) (gtk.SpinButton, error) {
	return SpinButtonNewWithRange(min, max, step)
}

func (*RealGtk) SpinnerNew() (gtk.Spinner, error) {
	return SpinnerNew()
}

func (*RealGtk) StatusbarNew() (gtk.Statusbar, error) {
	return StatusbarNew()
}

func (*RealGtk) StyleContextResetWidgets(v gdk.Screen) {
	StyleContextResetWidgets(v.(*gdk_impl.Screen))
}

func (*RealGtk) SwitchNew() (gtk.Switch, error) {
	return SwitchNew()
}

func (*RealGtk) TargetEntryNew(target string, flags gtk.TargetFlags, info uint) (gtk.TargetEntry, error) {
	return TargetEntryNew(target, flags, info)
}

func (*RealGtk) TextBufferNew(table gtk.TextTagTable) (gtk.TextBuffer, error) {
	return TextBufferNew(table.(*TextTagTable))
}

func (*RealGtk) TextTagNew(name string) (gtk.TextTag, error) {
	return TextTagNew(name)
}

func (*RealGtk) TextTagTableNew() (gtk.TextTagTable, error) {
	return TextTagTableNew()
}

func (*RealGtk) TextViewNew() (gtk.TextView, error) {
	return TextViewNew()
}

func (*RealGtk) TextViewNewWithBuffer(buf gtk.TextBuffer) (gtk.TextView, error) {
	return TextViewNewWithBuffer(buf.(*TextBuffer))
}

func (*RealGtk) ToggleButtonNew() (gtk.ToggleButton, error) {
	return ToggleButtonNew()
}

func (*RealGtk) ToggleButtonNewWithLabel(label string) (gtk.ToggleButton, error) {
	return ToggleButtonNewWithLabel(label)
}

func (*RealGtk) ToggleButtonNewWithMnemonic(label string) (gtk.ToggleButton, error) {
	return ToggleButtonNewWithMnemonic(label)
}

func (*RealGtk) ToolButtonNew(iconWidget gtk.Widget, label string) (gtk.ToolButton, error) {
	return ToolButtonNew(asWidgetImpl(iconWidget), label)
}

func (*RealGtk) ToolItemNew() (gtk.ToolItem, error) {
	return ToolItemNew()
}

func (*RealGtk) ToolbarNew() (gtk.Toolbar, error) {
	return ToolbarNew()
}

func (*RealGtk) TreeIterNew() gtk.TreeIter {
	return TreeIterNew()
}

func (*RealGtk) TreePathFromList(list glib.List) gtk.TreePath {
	return TreePathFromList(list.(*glib_impl.List))
}

func (*RealGtk) TreePathNew() gtk.TreePath {
	return TreePathNew()
}

func (*RealGtk) TreePathNewFromString(path string) (gtk.TreePath, error) {
	return TreePathNewFromString(path)
}

func (*RealGtk) TreeStoreNew(types ...glib.Type) (gtk.TreeStore, error) {
	return TreeStoreNew(types...)
}

func (*RealGtk) TreeViewColumnNew() (gtk.TreeViewColumn, error) {
	return TreeViewColumnNew()
}

func (*RealGtk) TreeViewColumnNewWithAttribute(title string, renderer gtk.CellRenderer, attribute string, column int) (gtk.TreeViewColumn, error) {
	return TreeViewColumnNewWithAttribute(title, renderer.(ICellRenderer), attribute, column)
}

func (*RealGtk) TreeViewNew() (gtk.TreeView, error) {
	return TreeViewNew()
}

func (*RealGtk) TreeViewNewWithModel(model gtk.TreeModel) (gtk.TreeView, error) {
	return TreeViewNewWithModel(model.(ITreeModel))
}

func (*RealGtk) ViewportNew(hadjustment gtk.Adjustment, vadjustment gtk.Adjustment) (gtk.Viewport, error) {
	return ViewportNew(hadjustment.(*Adjustment), vadjustment.(*Adjustment))
}

func (*RealGtk) VolumeButtonNew() (gtk.VolumeButton, error) {
	return VolumeButtonNew()
}

func (*RealGtk) WindowGetDefaultIconName() (string, error) {
	return WindowGetDefaultIconName()
}

func (*RealGtk) WindowNew(t gtk.WindowType) (gtk.Window, error) {
	return WindowNew(t)
}

func (*RealGtk) WindowSetDefaultIcon(icon gdk.Pixbuf) {
	WindowSetDefaultIcon(icon)
}

func (*RealGtk) WindowSetDefaultIconFromFile(file string) error {
	return WindowSetDefaultIconFromFile(file)
}

func (*RealGtk) WindowSetDefaultIconName(s string) {
	WindowSetDefaultIconName(s)
}
