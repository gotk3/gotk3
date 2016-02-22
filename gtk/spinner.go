package gtk

type Spinner interface {
	Widget

	Start()
	Stop()
} // end of Spinner

func AssertSpinner(_ Spinner) {}
