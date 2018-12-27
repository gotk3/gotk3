package gtk

import "testing"

func TestActionableImplementsIActionable(t *testing.T) {
	var cut interface{}
	cut = &Actionable{}
	_, ok := cut.(IActionable)

	if !ok {
		t.Error("Actionable does not implement IActionable")
		return
	}
}

// TestGetSetActionName tests the getter and setter for action name
// using a button, as we need an actual instance implementing Actionable.
func TestGetSetActionName(t *testing.T) {
	cut, err := ButtonNew()
	if err != nil {
		t.Fatal("Error creating button", err.Error())
	}

	expected := "app.stuff"
	cut.SetActionName(expected)

	actual, err := cut.GetActionName()
	if err != nil {
		t.Fatal("Error getting action name", err.Error())
	}

	if expected != actual {
		t.Fatalf("Expected %s got %s", expected, actual)
	}
}
