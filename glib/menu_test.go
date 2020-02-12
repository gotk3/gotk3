package glib_test

import (
	"testing"

	"github.com/gotk3/gotk3/glib"
)

func TestGetSetAttributeValueCustomBool(t *testing.T) {
	testCases := []struct {
		desc  string
		value bool
	}{
		{
			desc:  "true",
			value: true,
		},
		{
			desc:  "false",
			value: false,
		},
	}

	menuItem := glib.MenuItemNew()

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			variant := glib.VariantFromBoolean(tC.value)

			menuItem.SetAttributeValue("custom-bool-attribute", variant)
			actual := menuItem.GetAttributeValue("custom-bool-attribute", glib.VARIANT_TYPE_BOOLEAN)

			if !actual.IsType(glib.VARIANT_TYPE_BOOLEAN) {
				t.Error("Expected value of type", glib.VARIANT_TYPE_BOOLEAN, "got", actual.Type())
			}

			if tC.value != actual.GetBoolean() {
				t.Error("Expected", tC.value, "got", actual)
			}
		})
	}
}

func TestUnsetLabel(t *testing.T) {
	menuItem := glib.MenuItemNewWithLabel("unit_label")

	menuItem.UnsetLabel()

	value := menuItem.GetAttributeValue(glib.MENU_ATTRIBUTE_LABEL, glib.VARIANT_TYPE_STRING)
	actual := value.GetString()

	if "" != actual {
		t.Error("Expected empty string, got", actual)
	}
}

func TestSetLabel(t *testing.T) {
	menuItem := glib.MenuItemNewWithLabel("unit_label")

	expected := "New Label"
	menuItem.SetLabel(expected)

	value := menuItem.GetAttributeValue(glib.MENU_ATTRIBUTE_LABEL, glib.VARIANT_TYPE_STRING)
	actual := value.GetString()

	if expected != actual {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestSetDetailedAction(t *testing.T) {
	menuItem := glib.MenuItemNewWithAction("unit_action")

	expected := "new-action"
	menuItem.SetDetailedAction(expected)

	value := menuItem.GetAttributeValue(glib.MENU_ATTRIBUTE_ACTION, glib.VARIANT_TYPE_STRING)
	actual := value.GetString()

	if expected != actual {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestSetActionAndTargetValue(t *testing.T) {
	menuItem := glib.MenuItemNew()

	t.Run("Action, Value", func(t *testing.T) {
		expectedValue := glib.VariantFromString("Hello!")
		expected := "act1"
		menuItem.SetActionAndTargetValue(expected, expectedValue)

		// Check target value
		actualValue := menuItem.GetAttributeValue(glib.MENU_ATTRIBUTE_TARGET, glib.VARIANT_TYPE_STRING)
		if expectedValue.Native() != actualValue.Native() {
			t.Error("Expected", expectedValue.Native(), "got", actualValue.Native())
		}

		// Check action value
		actualAction := menuItem.GetAttributeValue(glib.MENU_ATTRIBUTE_ACTION, glib.VARIANT_TYPE_STRING).GetString()
		if expected != actualAction {
			t.Error("Expected", expected, "got", actualAction)
		}
	})
	t.Run("Action, Null Value", func(t *testing.T) {
		expected := "act2"
		menuItem.SetActionAndTargetValue(expected, nil)

		// Check target value
		actualValue := menuItem.GetAttributeValue(glib.MENU_ATTRIBUTE_TARGET, glib.VARIANT_TYPE_STRING)
		if actualValue != nil {
			t.Error("Expected nil value got", actualValue.Native())
		}

		// Check action value
		actualAction := menuItem.GetAttributeValue(glib.MENU_ATTRIBUTE_ACTION, glib.VARIANT_TYPE_STRING).GetString()
		if expected != actualAction {
			t.Error("Expected", expected, "got", actualAction)
		}
	})
}

func TestUnsetAction(t *testing.T) {
	menuItem := glib.MenuItemNew()

	initialValue := glib.VariantFromString("Hello!")
	initial := "act1"
	menuItem.SetActionAndTargetValue(initial, initialValue)

	menuItem.UnsetAction()

	// Check target value
	actualValue := menuItem.GetAttributeValue(glib.MENU_ATTRIBUTE_TARGET, glib.VARIANT_TYPE_STRING)
	if actualValue != nil {
		t.Error("Expected nil value got", actualValue.Native())
	}

	// Check action value
	actualAction := menuItem.GetAttributeValue(glib.MENU_ATTRIBUTE_ACTION, glib.VARIANT_TYPE_STRING)
	if actualAction != nil {
		t.Error("Expected nil action got", actualAction.Native())
	}
}
