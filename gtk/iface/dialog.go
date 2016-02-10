package iface

type Dialog interface {
	Window

	AddActionWidget(Widget, ResponseType)
	AddButton(string, ResponseType) (Button, error)
	GetContentArea() (Box, error)
	GetResponseForWidget(Widget) ResponseType
	GetWidgetForResponse(ResponseType) (Widget, error)
	Response(ResponseType)
	Run() int
	SetDefaultResponse(ResponseType)
	SetResponseSensitive(ResponseType, bool)
} // end of Dialog

func AssertDialog(_ Dialog) {}
