package version

import "testing"

func TestGetVersion(t *testing.T) {
	v := GetVersion()
	if v == "" {
		t.Error("Version should not be empty")
	}
}

func TestGetBuildInfo(t *testing.T) {
	info := GetBuildInfo()
	if info == "" {
		t.Error("Build info should not be empty")
	}
}

func TestString(t *testing.T) {
	v := String()
	if v != Version {
		t.Errorf("String() = %v, want %v", v, Version)
	}
}
