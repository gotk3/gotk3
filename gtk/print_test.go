package gtk

import (
	"testing"
)

func init() {
	Init(nil)
}

// TestPageSetup tests creating and manipulating PageSetup
func TestPageSetup(t *testing.T) {
	_, err := PageSetupNew()
	if err != nil {
		t.Error(err)
	}
}

// TestPaperSize tests creating and manipulating PaperSize
func TestPaperSize(t *testing.T) {
	_, err := PaperSizeNew(PAPER_NAME_A4)
	if err != nil {
		t.Error(err)
	}
}

// TestPrintContext tests creating and manipulating PrintContext

// TestPrintOperation tests creating and manipulating PrintOperation
func TestPrintOperation(t *testing.T) {
	_, err := PrintOperationNew()
	if err != nil {
		t.Error(err)
	}
}

// TestPrintOperationPreview tests creating and manipulating PrintOperationPreview

// TestPrintSettings tests creating and manipulating PrintSettings
func TestPrintSettings(t *testing.T) {
	settings, err := PrintSettingsNew()
	if err != nil {
		t.Error(err)
	}

	settings.Set("Key1", "String1")
	settings.SetBool("Key2", true)
	settings.Set("Key3", "String2")
	settings.SetInt("Key4", 2)

	settings.ForEach(func(key, value string, ptr uintptr) {
	}, 0)
}

// TestPrintContext tests creating and manipulating PrintContext
