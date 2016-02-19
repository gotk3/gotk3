package gtk

type Statusbar interface {
	Box

	GetContextId(string) uint
	GetMessageArea() (Box, error)
	Pop(uint)
	Push(uint, string) uint
} // end of Statusbar

func AssertStatusbar(_ Statusbar) {}
