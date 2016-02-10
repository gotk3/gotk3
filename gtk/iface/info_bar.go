package iface

type InfoBar interface {
	Box

	AddActionWidget(Widget, ResponseType)
	AddButton(string, ResponseType)
	GetActionArea() (Widget, error)
	GetContentArea() (Box, error)
	GetMessageType() MessageType
	GetShowCloseButton() bool
	SetDefaultResponse(ResponseType)
	SetMessageType(MessageType)
	SetResponseSensitive(ResponseType, bool)
	SetShowCloseButton(bool)
} // end of InfoBar

func AssertInfoBar(_ InfoBar) {}
