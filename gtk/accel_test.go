// Same copyright and license as the rest of the files in this project

package gtk

import "testing"

func Test_AccelGroup_Locking(t *testing.T) {
	ag, _ := AccelGroupNew()
	if ag.IsLocked() {
		t.Error("A newly created AccelGroup should not be locked")
	}

	ag.Lock()

	if !ag.IsLocked() {
		t.Error("A locked AccelGroup should report being locked")
	}

	ag.Unlock()

	if ag.IsLocked() {
		t.Error("An unlocked AccelGroup should report being unlocked")
	}
}

func Test_AcceleratorParse(t *testing.T) {
	l, r := AcceleratorParse("<Shift><Alt>F1")
	if l != 65470 {
		t.Errorf("Expected parsed key to equal %d but was %d", 65470, l)
	}
	if r != 9 {
		t.Errorf("Expected parsed mods to equal %d but was %d", 9, r)
	}
}
