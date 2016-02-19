package glib

type Notification interface {
	Object

	AddButton(string, string)
	SetBody(string)
	SetDefaultAction(string)
	SetTitle(string)
} // end of Notification

func AssertNotification(_ Notification) {}
