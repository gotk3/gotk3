package gtk

type Revealer interface {
	Bin

	GetChildRevealed() bool
	GetRevealChild() bool
	GetTransitionDuration() uint
	GetTransitionType() RevealerTransitionType
	SetRevealChild(bool)
	SetTransitionDuration(uint)
	SetTransitionType(RevealerTransitionType)
} // end of Revealer

func AssertRevealer(_ Revealer) {}
