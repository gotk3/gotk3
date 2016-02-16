package glib

type Application interface {
	Object

	Activate()
	GetApplicationID() string
	GetDbusObjectPath() string
	GetFlags() ApplicationFlags
	GetInactivityTimeout() uint
	GetIsRegistered() bool
	GetIsRemote() bool
	Hold()
	MarkBusy()
	Quit()
	Release()
	Run([]string) int
	SendNotification(string, Notification)
	SetApplicationID(string)
	SetDefault()
	SetFlags(ApplicationFlags)
	SetInactivityTimeout(uint)
	UnmarkBusy()
	WithdrawNotification(string)
} // end of Application

func AssertApplication(_ Application) {}
