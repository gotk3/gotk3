// Same copyright and license as the rest of the files in this project

package glib_test

import (
	"math"
	"testing"

	"github.com/gotk3/gotk3/glib"
)

func TestVariantGetInt(t *testing.T) {
	t.Run("int16", func(t *testing.T) {
		expected := int16(math.MinInt16)
		variant := glib.VariantFromInt16(expected)
		actual, err := variant.GetInt()
		if err != nil {
			t.Error("Unexpected error:", err.Error())
		}
		if int64(expected) != actual {
			t.Error("Expected", expected, "got", actual)
		}
	})

	t.Run("int32", func(t *testing.T) {
		expected := int32(math.MinInt32)
		variant := glib.VariantFromInt32(expected)
		actual, err := variant.GetInt()
		if err != nil {
			t.Error("Unexpected error:", err.Error())
		}
		if int64(expected) != actual {
			t.Error("Expected", expected, "got", actual)
		}
	})

	t.Run("int64", func(t *testing.T) {
		expected := int64(math.MinInt64)
		variant := glib.VariantFromInt64(expected)
		actual, err := variant.GetInt()
		if err != nil {
			t.Error("Unexpected error:", err.Error())
		}
		if expected != actual {
			t.Error("Expected", expected, "got", actual)
		}
	})

	t.Run("other type", func(t *testing.T) {
		variant := glib.VariantFromUint64(987)
		_, err := variant.GetInt()
		if err == nil {
			t.Error("expected error, did not get one")
		}
	})
}

func TestVariantGetUint(t *testing.T) {
	t.Run("byte", func(t *testing.T) {
		expected := uint8(math.MaxUint8)
		variant := glib.VariantFromByte(expected)
		actual, err := variant.GetUint()
		if err != nil {
			t.Error("Unexpected error:", err.Error())
		}
		if uint64(expected) != actual {
			t.Error("Expected", expected, "got", actual)
		}
	})

	t.Run("int16", func(t *testing.T) {
		expected := uint16(math.MaxUint16)
		variant := glib.VariantFromUint16(expected)
		actual, err := variant.GetUint()
		if err != nil {
			t.Error("Unexpected error:", err.Error())
		}
		if uint64(expected) != actual {
			t.Error("Expected", expected, "got", actual)
		}
	})

	t.Run("int32", func(t *testing.T) {
		expected := uint32(math.MaxUint32)
		variant := glib.VariantFromUint32(expected)
		actual, err := variant.GetUint()
		if err != nil {
			t.Error("Unexpected error:", err.Error())
		}
		if uint64(expected) != actual {
			t.Error("Expected", expected, "got", actual)
		}
	})

	t.Run("int64", func(t *testing.T) {
		expected := uint64(math.MaxUint64)
		variant := glib.VariantFromUint64(expected)
		actual, err := variant.GetUint()
		if err != nil {
			t.Error("Unexpected error:", err.Error())
		}
		if expected != actual {
			t.Error("Expected", expected, "got", actual)
		}
	})

	t.Run("other type", func(t *testing.T) {
		variant := glib.VariantFromInt64(987)
		_, err := variant.GetUint()
		if err == nil {
			t.Error("expected error, did not get one")
		}
	})
}

func TestVariantType(t *testing.T) {
	variant := glib.VariantFromBoolean(true)
	variantType := variant.Type()
	if !glib.VariantTypeEqual(glib.VARIANT_TYPE_BOOLEAN, variantType) {
		t.Error("Expected", glib.VARIANT_TYPE_BOOLEAN, "got", variantType)
	}
}

func TestVariantBool(t *testing.T) {
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
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			variant := glib.VariantFromBoolean(tC.value)
			actual := variant.GetBoolean()
			if tC.value != actual {
				t.Error("Expected", tC.value, "got", actual)
			}
		})
	}
}

func TestVariantString(t *testing.T) {
	testCases := []struct {
		desc  string
		value string
	}{
		{
			desc:  "Plain string",
			value: "Simple Data",
		},
		{
			desc:  "String with special characters",
			value: "Üö@/Data",
		},
		{
			desc:  "Empty String",
			value: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			variant := glib.VariantFromString(tC.value)
			actual := variant.GetString()
			if tC.value != actual {
				t.Error("Expected", tC.value, "got", actual)
			}
		})
	}
}

func TestVariantVariant(t *testing.T) {

	boxed := glib.VariantFromString("I'm in a box")

	variant := glib.VariantFromVariant(boxed)

	actual := variant.GetVariant()
	if boxed.Native() != actual.Native() {
		t.Error("Expected", boxed.Native(), "got", actual.Native())
	}
}
