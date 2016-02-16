package gtk

type Calendar interface {
	Widget

	ClearMarks()
	GetDate() (uint, uint, uint)
	GetDayIsMarked(uint) bool
	GetDetailHeightRows() int
	GetDetailWidthChars() int
	GetDisplayOptions() CalendarDisplayOptions
	MarkDay(uint)
	SelectDay(uint)
	SelectMonth(uint, uint)
	SetDetailHeightRows(int)
	SetDetailWidthChars(int)
	SetDisplayOptions(CalendarDisplayOptions)
	UnmarkDay(uint)
} // end of Calendar

func AssertCalendar(_ Calendar) {}
