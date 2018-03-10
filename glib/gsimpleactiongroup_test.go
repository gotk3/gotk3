package glib

import (
	"testing"
)

func TestSimpleActionGroupNew(t *testing.T) {
	sag := SimpleActionGroupNew()
	if sag == nil {
		t.Error("SimpleActionGroupNew returned nil")
	}

	if sag.IActionGroup == nil {
		t.Error("Embedded IActionGroup is nil")
	}
	if sag.IActionMap == nil {
		t.Error("Embedded IActionGroup is nil")
	}
}

func TestSimpleActionGroup_AddAction_RemoveAction_HasAction(t *testing.T) {
	sag := SimpleActionGroupNew()
	if sag == nil {
		t.Error("SimpleActionGroup returned nil")
	}

	// Check before: empty
	hasAction := sag.HasAction("nope")
	if hasAction {
		t.Error("Action group contained unexpected action 'nope'")
	}
	hasAction = sag.HasAction("yepp")
	if hasAction {
		t.Error("Action group contained unexpected action 'yepp'")
	}

	// Add a new action
	act := SimpleActionNew("yepp", nil)
	if act == nil {
		t.Error("SimpleActionNew returned nil")
	}
	sag.AddAction(act)

	// Check that it exists
	hasAction = sag.HasAction("nope")
	if hasAction {
		t.Error("Action group contained unexpected action 'nope'")
	}
	hasAction = sag.HasAction("yepp")
	if !hasAction {
		t.Error("Action group did not contain action 'yepp' after adding it")
	}

	// Remove the action again
	sag.RemoveAction("yepp")

	// Check that it was removed
	hasAction = sag.HasAction("nope")
	if hasAction {
		t.Error("Action group contained unexpected action 'nope'")
	}
	hasAction = sag.HasAction("yepp")
	if hasAction {
		t.Error("Action group contained unexpected action 'yepp'")
	}

	// NoFail check: removing a non-existing action
	sag.RemoveAction("yepp")
	sag.RemoveAction("nope")
}
