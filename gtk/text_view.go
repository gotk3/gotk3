package gtk

type TextView interface {
	Container

	GetAcceptsTab() bool
	GetBuffer() (TextBuffer, error)
	GetCursorVisible() bool
	GetEditable() bool
	GetIndent() int
	GetInputHints() InputHints
	GetInputPurpose() InputPurpose
	GetJustification() Justification
	GetLeftMargin() int
	GetOverwrite() bool
	GetPixelsAboveLines() int
	GetPixelsBelowLines() int
	GetPixelsInsideWrap() int
	GetRightMargin() int
	GetWrapMode() WrapMode
	SetAcceptsTab(bool)
	SetBuffer(TextBuffer)
	SetCursorVisible(bool)
	SetEditable(bool)
	SetIndent(int)
	SetInputHints(InputHints)
	SetInputPurpose(InputPurpose)
	SetJustification(Justification)
	SetLeftMargin(int)
	SetOverwrite(bool)
	SetPixelsAboveLines(int)
	SetPixelsBelowLines(int)
	SetPixelsInsideWrap(int)
	SetRightMargin(int)
	SetWrapMode(WrapMode)
} // end of TextView

func AssertTextView(_ TextView) {}
