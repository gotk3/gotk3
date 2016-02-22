package cairo

type ErrorStatus interface {
	Error() string
} // end of ErrorStatus

func AssertErrorStatus(_ ErrorStatus) {}
