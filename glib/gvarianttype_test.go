// Same copyright and license as the rest of the files in this project

package glib_test

import (
	"testing"

	"github.com/gotk3/gotk3/glib"
)

func TestVariantTypeEqual(t *testing.T) {
	testCases := []struct {
		desc     string
		type1    *glib.VariantType
		type2    *glib.VariantType
		expected bool
	}{
		{
			desc:     "bool == bool constants",
			type1:    glib.VARIANT_TYPE_BOOLEAN,
			type2:    glib.VARIANT_TYPE_BOOLEAN,
			expected: true,
		},
		{
			desc:     "bool != string constants",
			type1:    glib.VARIANT_TYPE_BOOLEAN,
			type2:    glib.VARIANT_TYPE_STRING,
			expected: false,
		},
		{
			desc:     "bool == bool dynamic",
			type1:    glib.VARIANT_TYPE_BOOLEAN,
			type2:    glib.VariantTypeNew("b"),
			expected: true,
		},
		{
			desc:     "bool != string dynamic",
			type1:    glib.VariantTypeNew("b"),
			type2:    glib.VariantTypeNew("s"),
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := glib.VariantTypeEqual(tC.type1, tC.type2)
			if tC.expected != actual {
				t.Error("Expected", tC.expected, "got", actual)
			}
		})
	}
}

func TestVariantTypeStringIsValid(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected bool
	}{
		{
			desc:     "String",
			input:    "s",
			expected: true,
		},
		{
			desc:     "Boolean",
			input:    "b",
			expected: true,
		},
		{
			desc:     "Tuple of String and Boolean",
			input:    "(sb)",
			expected: true,
		},
		{
			desc:     "Junk",
			input:    "r{{sb}",
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := glib.VariantTypeStringIsValid(tC.input)
			if tC.expected != actual {
				t.Error("Expected", tC.expected, "got", actual)
			}
		})
	}
}

func TestVariantIsSubtypeOf(t *testing.T) {
	testCases := []struct {
		desc      string
		type1     *glib.VariantType
		superType *glib.VariantType
		expected  bool
	}{
		{
			desc:      "bool is not a supertype",
			type1:     glib.VARIANT_TYPE_STRING,
			superType: glib.VARIANT_TYPE_BOOLEAN,
			expected:  false,
		},
		{
			desc:      "a* is supertype of as",
			type1:     glib.VariantTypeNew("as"),
			superType: glib.VariantTypeNew("a*"),
			expected:  true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := tC.type1.IsSubtypeOf(tC.superType)
			if tC.expected != actual {
				t.Error("Expected", tC.expected, "got", actual)
			}
		})
	}
}
