package gtk

import "testing"

func TestCheckVersion(t *testing.T) {
	err := CheckVersion(GetMajorVersion(), GetMinorVersion(), GetMicroVersion())
	if err != nil {
		t.Error(err)
	}

	err = CheckVersion(GetMajorVersion(), GetMinorVersion(), GetMicroVersion()-1)
	if err != nil {
		t.Error(err)
	}

	err = CheckVersion(GetMajorVersion(), GetMinorVersion(), GetMicroVersion()+1)
	if err == nil {
		t.Error("Expected to fail when an more recent version is expected")
	}
}
