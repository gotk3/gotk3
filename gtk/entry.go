package gtk

type Entry interface {
	Widget

	GetActivatesDefault() bool
	GetAlignment() float32
	GetBuffer() (EntryBuffer, error)
	GetCompletion() (EntryCompletion, error)
	GetCurrentIconDragSource() int
	GetCursorHAdjustment() (Adjustment, error)
	GetHasFrame() bool
	GetIconActivatable(EntryIconPosition) bool
	GetIconAtPos(int, int) int
	GetIconName(EntryIconPosition) (string, error)
	GetIconSensitive(EntryIconPosition) bool
	GetIconStorageType(EntryIconPosition) ImageType
	GetIconTooltipMarkup(EntryIconPosition) (string, error)
	GetIconTooltipText(EntryIconPosition) (string, error)
	GetInputHints() InputHints
	GetInputPurpose() InputPurpose
	GetInvisibleChar() rune
	GetLayoutOffsets() (int, int)
	GetMaxLength() int
	GetOverwriteMode() bool
	GetPlaceholderText() (string, error)
	GetProgressFraction() float64
	GetProgressPulseStep() float64
	GetText() (string, error)
	GetTextLength() uint16
	GetVisibility() bool
	GetWidthChars() int
	LayoutIndexToTextIndex(int) int
	ProgressPulse()
	ResetIMContext()
	SetActivatesDefault(bool)
	SetAlignment(float32)
	SetBuffer(EntryBuffer)
	SetCompletion(EntryCompletion)
	SetCursorHAdjustment(Adjustment)
	SetHasFrame(bool)
	SetIconActivatable(EntryIconPosition, bool)
	SetIconFromIconName(EntryIconPosition, string)
	SetIconSensitive(EntryIconPosition, bool)
	SetIconTooltipMarkup(EntryIconPosition, string)
	SetIconTooltipText(EntryIconPosition, string)
	SetInputHints(InputHints)
	SetInputPurpose(InputPurpose)
	SetInvisibleChar(rune)
	SetMaxLength(int)
	SetOverwriteMode(bool)
	SetPlaceholderText(string)
	SetProgressFraction(float64)
	SetProgressPulseStep(float64)
	SetText(string)
	SetVisibility(bool)
	SetWidthChars(int)
	TextIndexToLayoutIndex(int) int
	UnsetInvisibleChar()
} // end of Entry

func AssertEntry(_ Entry) {}
