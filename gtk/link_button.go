package gtk

type LinkButton interface {
	Button

	GetUri() string
	SetUri(string)
} // end of LinkButton

func AssertLinkButton(_ LinkButton) {}
