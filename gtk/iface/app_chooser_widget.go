package iface

type AppChooserWidget interface {
	Box

	GetDefaultText() (string, error)
	GetShowAll() bool
	GetShowDefault() bool
	GetShowFallback() bool
	GetShowOther() bool
	GetShowRecommended() bool
	SetDefaultText(string)
	SetShowAll(bool)
	SetShowDefault(bool)
	SetShowFallback(bool)
	SetShowOther(bool)
	SetShowRecommended(bool)
} // end of AppChooserWidget

func AssertAppChooserWidget(_ AppChooserWidget) {}
