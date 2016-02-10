package gtk

import "github.com/gotk3/gotk3/gtk/iface"
import gdk_iface "github.com/gotk3/gotk3/gdk/iface"
import cairo_iface "github.com/gotk3/gotk3/cairo/iface"
import glib_iface "github.com/gotk3/gotk3/glib/iface"

type RealGtk struct{}

var Real = &RealGtk{}

func (*RealGtk) AboutDialogNew() (iface.AboutDialog, error) {
	return AboutDialogNew()
}

func (*RealGtk) AccelGroupFromClosure(f interface{}) iface.AccelGroup {
	return AccelGroupFromClosure(f)
}

func (*RealGtk) AccelGroupNew() (iface.AccelGroup, error) {
	return AccelGroupNew()
}

func (*RealGtk) AccelGroupsActivate(obj glib_iface.Object, key uint, mods gdk_iface.ModifierType) bool {
	return AccelGroupsActivate(obj, key, mods)
}

func (*RealGtk) AccelGroupsFromObject(obj glib_iface.Object) glib_iface.SList {
	return AccelGroupsFromObject(obj)
}

func (*RealGtk) AccelMapAddEntry(path string, key uint, mods gdk_iface.ModifierType) {
	AccelMapAddEntry(path, key, mods)
}

func (*RealGtk) AccelMapAddFilter(filter string) {
	AccelMapAddFilter(filter)
}

func (*RealGtk) AccelMapChangeEntry(path string, key uint, mods gdk_iface.ModifierType, replace bool) bool {
	return AccelMapChangeEntry(path, key, mods, replace)
}

func (*RealGtk) AccelMapGet() iface.AccelMap {
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

func (*RealGtk) AccelMapLookupEntry(path string) iface.AccelKey {
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

func (*RealGtk) AcceleratorGetDefaultModMask() gdk_iface.ModifierType {
	return AcceleratorGetDefaultModMask()
}

func (*RealGtk) AcceleratorGetLabel(key uint, mods gdk_iface.ModifierType) string {
	return AcceleratorGetLabel(key, mods)
}

func (*RealGtk) AcceleratorName(key uint, mods gdk_iface.ModifierType) string {
	return AcceleratorName(key, mods)
}

func (*RealGtk) AcceleratorParse(acc string) (uint, gdk_iface.ModifierType) {
	return AcceleratorParse(acc)
}

func (*RealGtk) AcceleratorSetDefaultModMask(mods gdk_iface.ModifierType) {
	AcceleratorSetDefaultModMask(mods)
}

func (*RealGtk) AcceleratorValid(key uint, mods gdk_iface.ModifierType) bool {
	return AcceleratorValid(key, mods)
}

func (*RealGtk) AddProviderForScreen(s gdk_iface.Screen, provider StyleProvider, prio uint) {
	AddProviderForScreen(s, provider, prio)
}

func (*RealGtk) AdjustmentNew(value float64, lower float64, upper float64, stepIncrement float64, pageIncrement float64, pageSize float64) (iface.Adjustment, error) {
	return AdjustmentNew(value, lower, upper, stepIncrement, pageIncrement, pageSize)
}

func (*RealGtk) AppChooserButtonNew(content_type string) (iface.AppChooserButton, error) {
	return AppChooserButtonNew(content_type)
}

func (*RealGtk) AppChooserDialogNewForContentType(parent iface.Window, flags iface.DialogFlags, content_type string) (iface.AppChooserDialog, error) {
	return AppChooserDialogNewForContentType(parent, flags, content_type)
}

func (*RealGtk) AppChooserWidgetNew(content_type string) (iface.AppChooserWidget, error) {
	return AppChooserWidgetNew(content_type)
}

func (*RealGtk) ApplicationNew(appId string, flags glib_iface.ApplicationFlags) (iface.Application, error) {
	return ApplicationNew(appId, flags)
}

func (*RealGtk) ApplicationWindowNew(app iface.Application) (iface.ApplicationWindow, error) {
	return ApplicationWindowNew(app)
}

func (*RealGtk) AssistantNew() (iface.Assistant, error) {
	return AssistantNew()
}

func (*RealGtk) BoxNew(orientation iface.Orientation, spacing int) (iface.Box, error) {
	return BoxNew(orientation, spacing)
}

func (*RealGtk) BuilderNew() (iface.Builder, error) {
	return BuilderNew()
}

func (*RealGtk) ButtonNew() (iface.Button, error) {
	return ButtonNew()
}

func (*RealGtk) ButtonNewWithLabel(label string) (iface.Button, error) {
	return ButtonNewWithLabel(label)
}

func (*RealGtk) ButtonNewWithMnemonic(label string) (iface.Button, error) {
	return ButtonNewWithMnemonic(label)
}

func (*RealGtk) CalendarNew() (iface.Calendar, error) {
	return CalendarNew()
}

func (*RealGtk) CellRendererPixbufNew() (iface.CellRendererPixbuf, error) {
	return CellRendererPixbufNew()
}

func (*RealGtk) CellRendererSpinnerNew() (iface.CellRendererSpinner, error) {
	return CellRendererSpinnerNew()
}

func (*RealGtk) CellRendererTextNew() (iface.CellRendererText, error) {
	return CellRendererTextNew()
}

func (*RealGtk) CellRendererToggleNew() (iface.CellRendererToggle, error) {
	return CellRendererToggleNew()
}

func (*RealGtk) CheckButtonNew() (iface.CheckButton, error) {
	return CheckButtonNew()
}

func (*RealGtk) CheckButtonNewWithLabel(label string) (iface.CheckButton, error) {
	return CheckButtonNewWithLabel(label)
}

func (*RealGtk) CheckButtonNewWithMnemonic(label string) (iface.CheckButton, error) {
	return CheckButtonNewWithMnemonic(label)
}

func (*RealGtk) CheckMenuItemNew() (iface.CheckMenuItem, error) {
	return CheckMenuItemNew()
}

func (*RealGtk) CheckMenuItemNewWithLabel(label string) (iface.CheckMenuItem, error) {
	return CheckMenuItemNewWithLabel(label)
}

func (*RealGtk) CheckMenuItemNewWithMnemonic(label string) (iface.CheckMenuItem, error) {
	return CheckMenuItemNewWithMnemonic(label)
}

func (*RealGtk) CheckVersion(major uint, minor uint, micro uint) error {
	return CheckVersion(major, minor, micro)
}

func (*RealGtk) ClipboardGet(atom gdk_iface.Atom) (iface.Clipboard, error) {
	return ClipboardGet(atom)
}

func (*RealGtk) ClipboardGetForDisplay(display gdk_iface.Display, atom gdk_iface.Atom) (iface.Clipboard, error) {
	return ClipboardGetForDisplay(display, atom)
}

func (*RealGtk) ColorButtonNew() (iface.ColorButton, error) {
	return ColorButtonNew()
}

func (*RealGtk) ColorButtonNewWithRGBA(gdkColor gdk_iface.RGBA) (iface.ColorButton, error) {
	return ColorButtonNewWithRGBA(gdkColor)
}

func (*RealGtk) ColorChooserDialogNew(title string, parent iface.Window) (iface.ColorChooserDialog, error) {
	return ColorChooserDialogNew(title, parent)
}

func (*RealGtk) ComboBoxNew() (iface.ComboBox, error) {
	return ComboBoxNew()
}

func (*RealGtk) ComboBoxNewWithEntry() (iface.ComboBox, error) {
	return ComboBoxNewWithEntry()
}

func (*RealGtk) ComboBoxNewWithModel(model ITreeModel) (iface.ComboBox, error) {
	return ComboBoxNewWithModel(model)
}

func (*RealGtk) ComboBoxTextNew() (iface.ComboBoxText, error) {
	return ComboBoxTextNew()
}

func (*RealGtk) ComboBoxTextNewWithEntry() (iface.ComboBoxText, error) {
	return ComboBoxTextNewWithEntry()
}

func (*RealGtk) CssProviderGetDefault() (iface.CssProvider, error) {
	return CssProviderGetDefault()
}

func (*RealGtk) CssProviderGetNamed(name string, variant string) (iface.CssProvider, error) {
	return CssProviderGetNamed(name, variant)
}

func (*RealGtk) CssProviderNew() (iface.CssProvider, error) {
	return CssProviderNew()
}

func (*RealGtk) DialogNew() (iface.Dialog, error) {
	return DialogNew()
}

func (*RealGtk) DrawingAreaNew() (iface.DrawingArea, error) {
	return DrawingAreaNew()
}

func (*RealGtk) EntryBufferNew(initialChars string, nInitialChars int) (iface.EntryBuffer, error) {
	return EntryBufferNew(initialChars, nInitialChars)
}

func (*RealGtk) EntryNew() (iface.Entry, error) {
	return EntryNew()
}

func (*RealGtk) EntryNewWithBuffer(buffer iface.EntryBuffer) (iface.Entry, error) {
	return EntryNewWithBuffer(buffer)
}

func (*RealGtk) EventBoxNew() (iface.EventBox, error) {
	return EventBoxNew()
}

func (*RealGtk) EventsPending() bool {
	return EventsPending()
}

func (*RealGtk) ExpanderNew(label string) (iface.Expander, error) {
	return ExpanderNew(label)
}

func (*RealGtk) FileChooserButtonNew(title string, action iface.FileChooserAction) (iface.FileChooserButton, error) {
	return FileChooserButtonNew(title, action)
}

func (*RealGtk) FileChooserDialogNewWith1Button(title string, parent iface.Window, action iface.FileChooserAction, first_button_text string, first_button_id iface.ResponseType) (iface.FileChooserDialog, error) {
	return FileChooserDialogNewWith1Button(title, parent, action, first_button_text, first_button_id)
}

func (*RealGtk) FileChooserDialogNewWith2Buttons(title string, parent iface.Window, action iface.FileChooserAction, first_button_text string, first_button_id iface.ResponseType, second_button_text string, second_button_id iface.ResponseType) (iface.FileChooserDialog, error) {
	return FileChooserDialogNewWith2Buttons(title, parent, action, first_button_text, first_button_id, second_button_text, second_button_id)
}

func (*RealGtk) FileChooserWidgetNew(action iface.FileChooserAction) (iface.FileChooserWidget, error) {
	return FileChooserWidgetNew(action)
}

func (*RealGtk) FileFilterNew() (iface.FileFilter, error) {
	return FileFilterNew()
}

func (*RealGtk) FontButtonNew() (iface.FontButton, error) {
	return FontButtonNew()
}

func (*RealGtk) FontButtonNewWithFont(fontname string) (iface.FontButton, error) {
	return FontButtonNewWithFont(fontname)
}

func (*RealGtk) FrameNew(label string) (iface.Frame, error) {
	return FrameNew(label)
}

func (*RealGtk) GdkCairoSetSourcePixBuf(cr cairo_iface.Context, pixbuf gdk_iface.Pixbuf, pixbufX float64, pixbufY float64) {
	GdkCairoSetSourcePixBuf(cr, pixbuf, pixbufX, pixbufY)
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

func (*RealGtk) GridNew() (iface.Grid, error) {
	return GridNew()
}

func (*RealGtk) IconThemeGetDefault() (iface.IconTheme, error) {
	return IconThemeGetDefault()
}

func (*RealGtk) IconThemeGetForScreen(screen gdk_iface.Screen) (iface.IconTheme, error) {
	return IconThemeGetForScreen(screen)
}

func (*RealGtk) IconViewNew() (iface.IconView, error) {
	return IconViewNew()
}

func (*RealGtk) IconViewNewWithModel(model TreeModel) (iface.IconView, error) {
	return IconViewNewWithModel(model)
}

func (*RealGtk) ImageNew() (iface.Image, error) {
	return ImageNew()
}

func (*RealGtk) ImageNewFromFile(filename string) (iface.Image, error) {
	return ImageNewFromFile(filename)
}

func (*RealGtk) ImageNewFromIconName(iconName string, size iface.IconSize) (iface.Image, error) {
	return ImageNewFromIconName(iconName, size)
}

func (*RealGtk) ImageNewFromPixbuf(pixbuf gdk_iface.Pixbuf) (iface.Image, error) {
	return ImageNewFromPixbuf(pixbuf)
}

func (*RealGtk) ImageNewFromResource(resourcePath string) (iface.Image, error) {
	return ImageNewFromResource(resourcePath)
}

func (*RealGtk) InfoBarNew() (iface.InfoBar, error) {
	return InfoBarNew()
}

func (*RealGtk) Init(args *[]string) {
	Init(args)
}

func (*RealGtk) LabelNew(str string) (iface.Label, error) {
	return LabelNew(str)
}

func (*RealGtk) LabelNewWithMnemonic(str string) (iface.Label, error) {
	return LabelNewWithMnemonic(str)
}

func (*RealGtk) LayoutNew(hadjustment iface.Adjustment, vadjustment iface.Adjustment) (iface.Layout, error) {
	return LayoutNew(hadjustment, vadjustment)
}

func (*RealGtk) LevelBarNew() (iface.LevelBar, error) {
	return LevelBarNew()
}

func (*RealGtk) LevelBarNewForInterval(min_value float64, max_value float64) (iface.LevelBar, error) {
	return LevelBarNewForInterval(min_value, max_value)
}

func (*RealGtk) LinkButtonNew(label string) (iface.LinkButton, error) {
	return LinkButtonNew(label)
}

func (*RealGtk) LinkButtonNewWithLabel(uri string, label string) (iface.LinkButton, error) {
	return LinkButtonNewWithLabel(uri, label)
}

func (*RealGtk) ListStoreNew(types ...glib_iface.Type) (iface.ListStore, error) {
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

func (*RealGtk) MenuBarNew() (iface.MenuBar, error) {
	return MenuBarNew()
}

func (*RealGtk) MenuButtonNew() (iface.MenuButton, error) {
	return MenuButtonNew()
}

func (*RealGtk) MenuItemNew() (iface.MenuItem, error) {
	return MenuItemNew()
}

func (*RealGtk) MenuItemNewWithLabel(label string) (iface.MenuItem, error) {
	return MenuItemNewWithLabel(label)
}

func (*RealGtk) MenuItemNewWithMnemonic(label string) (iface.MenuItem, error) {
	return MenuItemNewWithMnemonic(label)
}

func (*RealGtk) MenuNew() (iface.Menu, error) {
	return MenuNew()
}

func (*RealGtk) MessageDialogNew(parent iface.Window, flags iface.DialogFlags, mType iface.MessageType, buttons iface.ButtonsType, format string, a ...interface{}) iface.MessageDialog {
	return MessageDialogNew(parent, flags, mType, buttons, format, a...)
}

func (*RealGtk) MessageDialogNewWithMarkup(parent iface.Window, flags iface.DialogFlags, mType iface.MessageType, buttons iface.ButtonsType, format string, a ...interface{}) iface.MessageDialog {
	return MessageDialogNewWithMarkup(parent, flags, mType, buttons, format, a...)
}

func (*RealGtk) NotebookNew() (iface.Notebook, error) {
	return NotebookNew()
}

func (*RealGtk) OffscreenWindowNew() (iface.OffscreenWindow, error) {
	return OffscreenWindowNew()
}

func (*RealGtk) PanedNew(orientation iface.Orientation) (iface.Paned, error) {
	return PanedNew(orientation)
}

func (*RealGtk) ProgressBarNew() (iface.ProgressBar, error) {
	return ProgressBarNew()
}

func (*RealGtk) RadioButtonNew(group glib_iface.SList) (iface.RadioButton, error) {
	return RadioButtonNew(group)
}

func (*RealGtk) RadioButtonNewFromWidget(radioGroupMember iface.RadioButton) (iface.RadioButton, error) {
	return RadioButtonNewFromWidget(radioGroupMember)
}

func (*RealGtk) RadioButtonNewWithLabel(group glib_iface.SList, label string) (iface.RadioButton, error) {
	return RadioButtonNewWithLabel(group, label)
}

func (*RealGtk) RadioButtonNewWithLabelFromWidget(radioGroupMember iface.RadioButton, label string) (iface.RadioButton, error) {
	return RadioButtonNewWithLabelFromWidget(radioGroupMember, label)
}

func (*RealGtk) RadioButtonNewWithMnemonic(group glib_iface.SList, label string) (iface.RadioButton, error) {
	return RadioButtonNewWithMnemonic(group, label)
}

func (*RealGtk) RadioButtonNewWithMnemonicFromWidget(radioGroupMember iface.RadioButton, label string) (iface.RadioButton, error) {
	return RadioButtonNewWithMnemonicFromWidget(radioGroupMember, label)
}

func (*RealGtk) RadioMenuItemNew(group glib_iface.SList) (iface.RadioMenuItem, error) {
	return RadioMenuItemNew(group)
}

func (*RealGtk) RadioMenuItemNewFromWidget(group iface.RadioMenuItem) (iface.RadioMenuItem, error) {
	return RadioMenuItemNewFromWidget(group)
}

func (*RealGtk) RadioMenuItemNewWithLabel(group glib_iface.SList, label string) (iface.RadioMenuItem, error) {
	return RadioMenuItemNewWithLabel(group, label)
}

func (*RealGtk) RadioMenuItemNewWithLabelFromWidget(group iface.RadioMenuItem, label string) (iface.RadioMenuItem, error) {
	return RadioMenuItemNewWithLabelFromWidget(group, label)
}

func (*RealGtk) RadioMenuItemNewWithMnemonic(group glib_iface.SList, label string) (iface.RadioMenuItem, error) {
	return RadioMenuItemNewWithMnemonic(group, label)
}

func (*RealGtk) RadioMenuItemNewWithMnemonicFromWidget(group iface.RadioMenuItem, label string) (iface.RadioMenuItem, error) {
	return RadioMenuItemNewWithMnemonicFromWidget(group, label)
}

func (*RealGtk) RecentFilterNew() (iface.RecentFilter, error) {
	return RecentFilterNew()
}

func (*RealGtk) RecentManagerGetDefault() (iface.RecentManager, error) {
	return RecentManagerGetDefault()
}

func (*RealGtk) RemoveProviderForScreen(s gdk_iface.Screen, provider IStyleProvider) {
	RemoveProviderForScreen(s, provider)
}

func (*RealGtk) ScaleButtonNew(size iface.IconSize, min float64, max float64, step float64, icons []string) (iface.ScaleButton, error) {
	return ScaleButtonNew(size, min, max, step, icons)
}

func (*RealGtk) ScaleNew(orientation iface.Orientation, adjustment iface.Adjustment) (iface.Scale, error) {
	return ScaleNew(orientation, adjustment)
}

func (*RealGtk) ScaleNewWithRange(orientation iface.Orientation, min float64, max float64, step float64) (iface.Scale, error) {
	return ScaleNewWithRange(orientation, min, max, step)
}

func (*RealGtk) ScrollbarNew(orientation iface.Orientation, adjustment iface.Adjustment) (iface.Scrollbar, error) {
	return ScrollbarNew(orientation, adjustment)
}

func (*RealGtk) ScrolledWindowNew(hadjustment iface.Adjustment, vadjustment iface.Adjustment) (iface.ScrolledWindow, error) {
	return ScrolledWindowNew(hadjustment, vadjustment)
}

func (*RealGtk) SearchEntryNew() (iface.SearchEntry, error) {
	return SearchEntryNew()
}

func (*RealGtk) SeparatorMenuItemNew() (iface.SeparatorMenuItem, error) {
	return SeparatorMenuItemNew()
}

func (*RealGtk) SeparatorNew(orientation iface.Orientation) (iface.Separator, error) {
	return SeparatorNew(orientation)
}

func (*RealGtk) SeparatorToolItemNew() (iface.SeparatorToolItem, error) {
	return SeparatorToolItemNew()
}

func (*RealGtk) SettingsGetDefault() (iface.Settings, error) {
	return SettingsGetDefault()
}

func (*RealGtk) SpinButtonNew(adjustment iface.Adjustment, climbRate float64, digits uint) (iface.SpinButton, error) {
	return SpinButtonNew(adjustment, climbRate, digits)
}

func (*RealGtk) SpinButtonNewWithRange(min float64, max float64, step float64) (iface.SpinButton, error) {
	return SpinButtonNewWithRange(min, max, step)
}

func (*RealGtk) SpinnerNew() (iface.Spinner, error) {
	return SpinnerNew()
}

func (*RealGtk) StatusbarNew() (iface.Statusbar, error) {
	return StatusbarNew()
}

func (*RealGtk) StyleContextResetWidgets(v gdk_iface.Screen) {
	StyleContextResetWidgets(v)
}

func (*RealGtk) SwitchNew() (iface.Switch, error) {
	return SwitchNew()
}

func (*RealGtk) TargetEntryNew(target string, flags iface.TargetFlags, info uint) (iface.TargetEntry, error) {
	return TargetEntryNew(target, flags, info)
}

func (*RealGtk) TextBufferNew(table iface.TextTagTable) (iface.TextBuffer, error) {
	return TextBufferNew(table)
}

func (*RealGtk) TextTagNew(name string) (iface.TextTag, error) {
	return TextTagNew(name)
}

func (*RealGtk) TextTagTableNew() (iface.TextTagTable, error) {
	return TextTagTableNew()
}

func (*RealGtk) TextViewNew() (iface.TextView, error) {
	return TextViewNew()
}

func (*RealGtk) TextViewNewWithBuffer(buf iface.TextBuffer) (iface.TextView, error) {
	return TextViewNewWithBuffer(buf)
}

func (*RealGtk) ToggleButtonNew() (iface.ToggleButton, error) {
	return ToggleButtonNew()
}

func (*RealGtk) ToggleButtonNewWithLabel(label string) (iface.ToggleButton, error) {
	return ToggleButtonNewWithLabel(label)
}

func (*RealGtk) ToggleButtonNewWithMnemonic(label string) (iface.ToggleButton, error) {
	return ToggleButtonNewWithMnemonic(label)
}

func (*RealGtk) ToolButtonNew(iconWidget iface.Widget, label string) (iface.ToolButton, error) {
	return ToolButtonNew(iconWidget, label)
}

func (*RealGtk) ToolItemNew() (iface.ToolItem, error) {
	return ToolItemNew()
}

func (*RealGtk) ToolbarNew() (iface.Toolbar, error) {
	return ToolbarNew()
}

func (*RealGtk) TreePathFromList(list glib_iface.List) iface.TreePath {
	return TreePathFromList(list)
}

func (*RealGtk) TreePathNewFromString(path string) (iface.TreePath, error) {
	return TreePathNewFromString(path)
}

func (*RealGtk) TreeStoreNew(types ...glib_iface.Type) (iface.TreeStore, error) {
	return TreeStoreNew(types...)
}

func (*RealGtk) TreeViewColumnNew() (iface.TreeViewColumn, error) {
	return TreeViewColumnNew()
}

func (*RealGtk) TreeViewColumnNewWithAttribute(title string, renderer iface.CellRenderer, attribute string, column int) (iface.TreeViewColumn, error) {
	return TreeViewColumnNewWithAttribute(title, renderer, attribute, column)
}

func (*RealGtk) TreeViewNew() (iface.TreeView, error) {
	return TreeViewNew()
}

func (*RealGtk) TreeViewNewWithModel(model iface.TreeModel) (iface.TreeView, error) {
	return TreeViewNewWithModel(model)
}

func (*RealGtk) ViewportNew(hadjustment iface.Adjustment, vadjustment iface.Adjustment) (iface.Viewport, error) {
	return ViewportNew(hadjustment, vadjustment)
}

func (*RealGtk) VolumeButtonNew() (iface.VolumeButton, error) {
	return VolumeButtonNew()
}

func (*RealGtk) WindowGetDefaultIconName() (string, error) {
	return WindowGetDefaultIconName()
}

func (*RealGtk) WindowNew(t iface.WindowType) (iface.Window, error) {
	return WindowNew(t)
}

func (*RealGtk) WindowSetDefaultIcon(icon gdk_iface.Pixbuf) {
	WindowSetDefaultIcon(icon)
}

func (*RealGtk) WindowSetDefaultIconFromFile(file string) error {
	return WindowSetDefaultIconFromFile(file)
}

func (*RealGtk) WindowSetDefaultIconName(s string) {
	WindowSetDefaultIconName(s)
}
