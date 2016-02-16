package gtk

// AccelFlags is a representation of GTK's GtkAccelFlags
type AccelFlags int

var (
	ACCEL_VISIBLE AccelFlags
	ACCEL_LOCKED  AccelFlags
	ACCEL_MASK    AccelFlags
)

// ApplicationInhibitFlags is a representation of GTK's GtkApplicationInhibitFlags.
type ApplicationInhibitFlags int

var (
	APPLICATION_INHIBIT_LOGOUT  ApplicationInhibitFlags
	APPLICATION_INHIBIT_SWITCH  ApplicationInhibitFlags
	APPLICATION_INHIBIT_SUSPEND ApplicationInhibitFlags
	APPLICATION_INHIBIT_IDLE    ApplicationInhibitFlags
)

// Stock is a special type that does not have an equivalent type in
// GTK.  It is the type used as a parameter anytime an identifier for
// stock icons are needed.  A Stock must be type converted to string when
// function parameters may take a Stock, but when other string values are
// valid as well.
type Stock string

var (
	STOCK_ABOUT                         Stock
	STOCK_ADD                           Stock
	STOCK_APPLY                         Stock
	STOCK_BOLD                          Stock
	STOCK_CANCEL                        Stock
	STOCK_CAPS_LOCK_WARNING             Stock
	STOCK_CDROM                         Stock
	STOCK_CLEAR                         Stock
	STOCK_CLOSE                         Stock
	STOCK_COLOR_PICKER                  Stock
	STOCK_CONNECT                       Stock
	STOCK_CONVERT                       Stock
	STOCK_COPY                          Stock
	STOCK_CUT                           Stock
	STOCK_DELETE                        Stock
	STOCK_DIALOG_AUTHENTICATION         Stock
	STOCK_DIALOG_INFO                   Stock
	STOCK_DIALOG_WARNING                Stock
	STOCK_DIALOG_ERROR                  Stock
	STOCK_DIALOG_QUESTION               Stock
	STOCK_DIRECTORY                     Stock
	STOCK_DISCARD                       Stock
	STOCK_DISCONNECT                    Stock
	STOCK_DND                           Stock
	STOCK_DND_MULTIPLE                  Stock
	STOCK_EDIT                          Stock
	STOCK_EXECUTE                       Stock
	STOCK_FILE                          Stock
	STOCK_FIND                          Stock
	STOCK_FIND_AND_REPLACE              Stock
	STOCK_FLOPPY                        Stock
	STOCK_FULLSCREEN                    Stock
	STOCK_GOTO_BOTTOM                   Stock
	STOCK_GOTO_FIRST                    Stock
	STOCK_GOTO_LAST                     Stock
	STOCK_GOTO_TOP                      Stock
	STOCK_GO_BACK                       Stock
	STOCK_GO_DOWN                       Stock
	STOCK_GO_FORWARD                    Stock
	STOCK_GO_UP                         Stock
	STOCK_HARDDISK                      Stock
	STOCK_HELP                          Stock
	STOCK_HOME                          Stock
	STOCK_INDEX                         Stock
	STOCK_INDENT                        Stock
	STOCK_INFO                          Stock
	STOCK_ITALIC                        Stock
	STOCK_JUMP_TO                       Stock
	STOCK_JUSTIFY_CENTER                Stock
	STOCK_JUSTIFY_FILL                  Stock
	STOCK_JUSTIFY_LEFT                  Stock
	STOCK_JUSTIFY_RIGHT                 Stock
	STOCK_LEAVE_FULLSCREEN              Stock
	STOCK_MISSING_IMAGE                 Stock
	STOCK_MEDIA_FORWARD                 Stock
	STOCK_MEDIA_NEXT                    Stock
	STOCK_MEDIA_PAUSE                   Stock
	STOCK_MEDIA_PLAY                    Stock
	STOCK_MEDIA_PREVIOUS                Stock
	STOCK_MEDIA_RECORD                  Stock
	STOCK_MEDIA_REWIND                  Stock
	STOCK_MEDIA_STOP                    Stock
	STOCK_NETWORK                       Stock
	STOCK_NEW                           Stock
	STOCK_NO                            Stock
	STOCK_OK                            Stock
	STOCK_OPEN                          Stock
	STOCK_ORIENTATION_PORTRAIT          Stock
	STOCK_ORIENTATION_LANDSCAPE         Stock
	STOCK_ORIENTATION_REVERSE_LANDSCAPE Stock
	STOCK_ORIENTATION_REVERSE_PORTRAIT  Stock
	STOCK_PAGE_SETUP                    Stock
	STOCK_PASTE                         Stock
	STOCK_PREFERENCES                   Stock
	STOCK_PRINT                         Stock
	STOCK_PRINT_ERROR                   Stock
	STOCK_PRINT_PAUSED                  Stock
	STOCK_PRINT_PREVIEW                 Stock
	STOCK_PRINT_REPORT                  Stock
	STOCK_PRINT_WARNING                 Stock
	STOCK_PROPERTIES                    Stock
	STOCK_QUIT                          Stock
	STOCK_REDO                          Stock
	STOCK_REFRESH                       Stock
	STOCK_REMOVE                        Stock
	STOCK_REVERT_TO_SAVED               Stock
	STOCK_SAVE                          Stock
	STOCK_SAVE_AS                       Stock
	STOCK_SELECT_ALL                    Stock
	STOCK_SELECT_COLOR                  Stock
	STOCK_SELECT_FONT                   Stock
	STOCK_SORT_ASCENDING                Stock
	STOCK_SORT_DESCENDING               Stock
	STOCK_SPELL_CHECK                   Stock
	STOCK_STOP                          Stock
	STOCK_STRIKETHROUGH                 Stock
	STOCK_UNDELETE                      Stock
	STOCK_UNDERLINE                     Stock
	STOCK_UNDO                          Stock
	STOCK_UNINDENT                      Stock
	STOCK_YES                           Stock
	STOCK_ZOOM_100                      Stock
	STOCK_ZOOM_FIT                      Stock
	STOCK_ZOOM_IN                       Stock
	STOCK_ZOOM_OUT                      Stock
)

// Align is a representation of GTK's GtkAlign.
type Align int

var (
	ALIGN_FILL   Align
	ALIGN_START  Align
	ALIGN_END    Align
	ALIGN_CENTER Align
)

// ArrowPlacement is a representation of GTK's GtkArrowPlacement.
type ArrowPlacement int

var (
	ARROWS_BOTH  ArrowPlacement
	ARROWS_START ArrowPlacement
	ARROWS_END   ArrowPlacement
)

// ArrowType is a representation of GTK's GtkArrowType.
type ArrowType int

var (
	ARROW_UP    ArrowType
	ARROW_DOWN  ArrowType
	ARROW_LEFT  ArrowType
	ARROW_RIGHT ArrowType
	ARROW_NONE  ArrowType
)

// AssistantPageType is a representation of GTK's GtkAssistantPageType.
type AssistantPageType int

var (
	ASSISTANT_PAGE_CONTENT  AssistantPageType
	ASSISTANT_PAGE_INTRO    AssistantPageType
	ASSISTANT_PAGE_CONFIRM  AssistantPageType
	ASSISTANT_PAGE_SUMMARY  AssistantPageType
	ASSISTANT_PAGE_PROGRESS AssistantPageType
	ASSISTANT_PAGE_CUSTOM   AssistantPageType
)

// ButtonsType is a representation of GTK's GtkButtonsType.
type ButtonsType int

var (
	BUTTONS_NONE      ButtonsType
	BUTTONS_OK        ButtonsType
	BUTTONS_CLOSE     ButtonsType
	BUTTONS_CANCEL    ButtonsType
	BUTTONS_YES_NO    ButtonsType
	BUTTONS_OK_CANCEL ButtonsType
)

// CalendarDisplayOptions is a representation of GTK's GtkCalendarDisplayOptions
type CalendarDisplayOptions int

var (
	CALENDAR_SHOW_HEADING      CalendarDisplayOptions
	CALENDAR_SHOW_DAY_NAMES    CalendarDisplayOptions
	CALENDAR_NO_MONTH_CHANGE   CalendarDisplayOptions
	CALENDAR_SHOW_WEEK_NUMBERS CalendarDisplayOptions
	CALENDAR_SHOW_DETAILS      CalendarDisplayOptions
)

// DestDefaults is a representation of GTK's GtkDestDefaults.
type DestDefaults int

var (
	DEST_DEFAULT_MOTION    DestDefaults
	DEST_DEFAULT_HIGHLIGHT DestDefaults
	DEST_DEFAULT_DROP      DestDefaults
	DEST_DEFAULT_ALL       DestDefaults
)

// DialogFlags is a representation of GTK's GtkDialogFlags.
type DialogFlags int

var (
	DIALOG_MODAL               DialogFlags
	DIALOG_DESTROY_WITH_PARENT DialogFlags
)

// EntryIconPosition is a representation of GTK's GtkEntryIconPosition.
type EntryIconPosition int

var (
	ENTRY_ICON_PRIMARY   EntryIconPosition
	ENTRY_ICON_SECONDARY EntryIconPosition
)

// FileChooserAction is a representation of GTK's GtkFileChooserAction.
type FileChooserAction int

var (
	FILE_CHOOSER_ACTION_OPEN          FileChooserAction
	FILE_CHOOSER_ACTION_SAVE          FileChooserAction
	FILE_CHOOSER_ACTION_SELECT_FOLDER FileChooserAction
	FILE_CHOOSER_ACTION_CREATE_FOLDER FileChooserAction
)

// IconLookupFlags is a representation of GTK's GtkIconLookupFlags.
type IconLookupFlags int

var (
	ICON_LOOKUP_NO_SVG           IconLookupFlags
	ICON_LOOKUP_FORCE_SVG        IconLookupFlags
	ICON_LOOKUP_USE_BUILTIN      IconLookupFlags
	ICON_LOOKUP_GENERIC_FALLBACK IconLookupFlags
	ICON_LOOKUP_FORCE_SIZE       IconLookupFlags
)

// IconSize is a representation of GTK's GtkIconSize.
type IconSize int

var (
	ICON_SIZE_INVALID       IconSize
	ICON_SIZE_MENU          IconSize
	ICON_SIZE_SMALL_TOOLBAR IconSize
	ICON_SIZE_LARGE_TOOLBAR IconSize
	ICON_SIZE_BUTTON        IconSize
	ICON_SIZE_DND           IconSize
	ICON_SIZE_DIALOG        IconSize
)

// ImageType is a representation of GTK's GtkImageType.
type ImageType int

var (
	IMAGE_EMPTY     ImageType
	IMAGE_PIXBUF    ImageType
	IMAGE_STOCK     ImageType
	IMAGE_ICON_SET  ImageType
	IMAGE_ANIMATION ImageType
	IMAGE_ICON_NAME ImageType
	IMAGE_GICON     ImageType
)

// InputHints is a representation of GTK's GtkInputHints.
type InputHints int

var (
	INPUT_HINT_NONE                InputHints
	INPUT_HINT_SPELLCHECK          InputHints
	INPUT_HINT_NO_SPELLCHECK       InputHints
	INPUT_HINT_WORD_COMPLETION     InputHints
	INPUT_HINT_LOWERCASE           InputHints
	INPUT_HINT_UPPERCASE_CHARS     InputHints
	INPUT_HINT_UPPERCASE_WORDS     InputHints
	INPUT_HINT_UPPERCASE_SENTENCES InputHints
	INPUT_HINT_INHIBIT_OSK         InputHints
)

// InputPurpose is a representation of GTK's GtkInputPurpose.
type InputPurpose int

var (
	INPUT_PURPOSE_FREE_FORM InputPurpose
	INPUT_PURPOSE_ALPHA     InputPurpose
	INPUT_PURPOSE_DIGITS    InputPurpose
	INPUT_PURPOSE_NUMBER    InputPurpose
	INPUT_PURPOSE_PHONE     InputPurpose
	INPUT_PURPOSE_URL       InputPurpose
	INPUT_PURPOSE_EMAIL     InputPurpose
	INPUT_PURPOSE_NAME      InputPurpose
	INPUT_PURPOSE_PASSWORD  InputPurpose
	INPUT_PURPOSE_PIN       InputPurpose
)

// Justify is a representation of GTK's GtkJustification.
type Justification int

var (
	JUSTIFY_LEFT   Justification
	JUSTIFY_RIGHT  Justification
	JUSTIFY_CENTER Justification
	JUSTIFY_FILL   Justification
)

// License is a representation of GTK's GtkLicense.
type License int

var (
	LICENSE_UNKNOWN      License
	LICENSE_CUSTOM       License
	LICENSE_GPL_2_0      License
	LICENSE_GPL_3_0      License
	LICENSE_LGPL_2_1     License
	LICENSE_LGPL_3_0     License
	LICENSE_BSD          License
	LICENSE_MIT_X11      License
	LICENSE_GTK_ARTISTIC License
)

// MessageType is a representation of GTK's GtkMessageType.
type MessageType int

var (
	MESSAGE_INFO     MessageType
	MESSAGE_WARNING  MessageType
	MESSAGE_QUESTION MessageType
	MESSAGE_ERROR    MessageType
	MESSAGE_OTHER    MessageType
)

// Orientation is a representation of GTK's GtkOrientation.
type Orientation int

var (
	ORIENTATION_HORIZONTAL Orientation
	ORIENTATION_VERTICAL   Orientation
)

// PackType is a representation of GTK's GtkPackType.
type PackType int

var (
	PACK_START PackType
	PACK_END   PackType
)

// PathType is a representation of GTK's GtkPathType.
type PathType int

var (
	PATH_WIDGET       PathType
	PATH_WIDGET_CLASS PathType
	PATH_CLASS        PathType
)

// PolicyType is a representation of GTK's GtkPolicyType.
type PolicyType int

var (
	POLICY_ALWAYS    PolicyType
	POLICY_AUTOMATIC PolicyType
	POLICY_NEVER     PolicyType
)

// PositionType is a representation of GTK's GtkPositionType.
type PositionType int

var (
	POS_LEFT   PositionType
	POS_RIGHT  PositionType
	POS_TOP    PositionType
	POS_BOTTOM PositionType
)

// ReliefStyle is a representation of GTK's GtkReliefStyle.
type ReliefStyle int

var (
	RELIEF_NORMAL ReliefStyle
	RELIEF_HALF   ReliefStyle
	RELIEF_NONE   ReliefStyle
)

// ResponseType is a representation of GTK's GtkResponseType.
type ResponseType int

var (
	RESPONSE_NONE         ResponseType
	RESPONSE_REJECT       ResponseType
	RESPONSE_ACCEPT       ResponseType
	RESPONSE_DELETE_EVENT ResponseType
	RESPONSE_OK           ResponseType
	RESPONSE_CANCEL       ResponseType
	RESPONSE_CLOSE        ResponseType
	RESPONSE_YES          ResponseType
	RESPONSE_NO           ResponseType
	RESPONSE_APPLY        ResponseType
	RESPONSE_HELP         ResponseType
)

// SelectionMode is a representation of GTK's GtkSelectionMode.
type SelectionMode int

var (
	SELECTION_NONE     SelectionMode
	SELECTION_SINGLE   SelectionMode
	SELECTION_BROWSE   SelectionMode
	SELECTION_MULTIPLE SelectionMode
)

// ShadowType is a representation of GTK's GtkShadowType.
type ShadowType int

var (
	SHADOW_NONE       ShadowType
	SHADOW_IN         ShadowType
	SHADOW_OUT        ShadowType
	SHADOW_ETCHED_IN  ShadowType
	SHADOW_ETCHED_OUT ShadowType
)

// SortType is a representation of GTK's GtkSortType.
type SortType int

var (
	SORT_ASCENDING  SortType
	SORT_DESCENDING SortType
)

// StateFlags is a representation of GTK's GtkStateFlags.
type StateFlags int

var (
	STATE_FLAG_NORMAL       StateFlags
	STATE_FLAG_ACTIVE       StateFlags
	STATE_FLAG_PRELIGHT     StateFlags
	STATE_FLAG_SELECTED     StateFlags
	STATE_FLAG_INSENSITIVE  StateFlags
	STATE_FLAG_INCONSISTENT StateFlags
	STATE_FLAG_FOCUSED      StateFlags
	STATE_FLAG_BACKDROP     StateFlags
)

// TargetFlags is a representation of GTK's GtkTargetFlags.
type TargetFlags int

var (
	TARGET_SAME_APP     TargetFlags
	TARGET_SAME_WIDGET  TargetFlags
	TARGET_OTHER_APP    TargetFlags
	TARGET_OTHER_WIDGET TargetFlags
)

// ToolbarStyle is a representation of GTK's GtkToolbarStyle.
type ToolbarStyle int

var (
	TOOLBAR_ICONS      ToolbarStyle
	TOOLBAR_TEXT       ToolbarStyle
	TOOLBAR_BOTH       ToolbarStyle
	TOOLBAR_BOTH_HORIZ ToolbarStyle
)

// TreeModelFlags is a representation of GTK's GtkTreeModelFlags.
type TreeModelFlags int

var (
	TREE_MODEL_ITERS_PERSIST TreeModelFlags
	TREE_MODEL_LIST_ONLY     TreeModelFlags
)

// WindowPosition is a representation of GTK's GtkWindowPosition.
type WindowPosition int

var (
	WIN_POS_NONE             WindowPosition
	WIN_POS_CENTER           WindowPosition
	WIN_POS_MOUSE            WindowPosition
	WIN_POS_CENTER_ALWAYS    WindowPosition
	WIN_POS_CENTER_ON_PARENT WindowPosition
)

// WindowType is a representation of GTK's GtkWindowType.
type WindowType int

var (
	WINDOW_TOPLEVEL WindowType
	WINDOW_POPUP    WindowType
)

// WrapMode is a representation of GTK's GtkWrapMode.
type WrapMode int

var (
	WRAP_NONE      WrapMode
	WRAP_CHAR      WrapMode
	WRAP_WORD      WrapMode
	WRAP_WORD_CHAR WrapMode
)

// LevelBarMode is a representation of GTK's GtkLevelBarMode.
type LevelBarMode int

var (
	LEVEL_BAR_MODE_CONTINUOUS LevelBarMode
	LEVEL_BAR_MODE_DISCRETE   LevelBarMode
)

type StyleProviderPriority int

var (
	STYLE_PROVIDER_PRIORITY_FALLBACK    StyleProviderPriority
	STYLE_PROVIDER_PRIORITY_THEME       StyleProviderPriority
	STYLE_PROVIDER_PRIORITY_SETTINGS    StyleProviderPriority
	STYLE_PROVIDER_PRIORITY_APPLICATION StyleProviderPriority
	STYLE_PROVIDER_PRIORITY_USER        StyleProviderPriority
)

// Convenient map for Columns and values (See ListStore, TreeStore)
type Cols map[int]interface{}

var (
	ALIGN_BASELINE Align
)

// RevealerTransitionType is a representation of GTK's GtkRevealerTransitionType.
type RevealerTransitionType int

var (
	REVEALER_TRANSITION_TYPE_NONE        RevealerTransitionType
	REVEALER_TRANSITION_TYPE_CROSSFADE   RevealerTransitionType
	REVEALER_TRANSITION_TYPE_SLIDE_RIGHT RevealerTransitionType
	REVEALER_TRANSITION_TYPE_SLIDE_LEFT  RevealerTransitionType
	REVEALER_TRANSITION_TYPE_SLIDE_UP    RevealerTransitionType
	REVEALER_TRANSITION_TYPE_SLIDE_DOWN  RevealerTransitionType
)

// StackTransitionType is a representation of GTK's GtkStackTransitionType.
type StackTransitionType int

var (
	STACK_TRANSITION_TYPE_NONE             StackTransitionType
	STACK_TRANSITION_TYPE_CROSSFADE        StackTransitionType
	STACK_TRANSITION_TYPE_SLIDE_RIGHT      StackTransitionType
	STACK_TRANSITION_TYPE_SLIDE_LEFT       StackTransitionType
	STACK_TRANSITION_TYPE_SLIDE_UP         StackTransitionType
	STACK_TRANSITION_TYPE_SLIDE_DOWN       StackTransitionType
	STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT StackTransitionType
	STACK_TRANSITION_TYPE_SLIDE_UP_DOWN    StackTransitionType
)

var (
	STATE_FLAG_DIR_LTR StateFlags
	STATE_FLAG_DIR_RTL StateFlags
)

var (
	LEVEL_BAR_OFFSET_LOW  string
	LEVEL_BAR_OFFSET_HIGH string
)
