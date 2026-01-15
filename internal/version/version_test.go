package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	v := GetVersion()
	assert.NotEmpty(t, v, "Version should not be empty")
}

func TestGetBuildInfo(t *testing.T) {
	info := GetBuildInfo()
	assert.NotEmpty(t, info, "Build info should not be empty")
}

func TestString(t *testing.T) {
	v := String()
	assert.Equal(t, Version, v, "String() should return Version")
}

func TestVersionFormat(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"has version value", "not empty"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expected == "not empty" {
				assert.NotEmpty(t, Version, "Version should not be empty")
			}
		})
	}
}

func TestVersionConsistency(t *testing.T) {
	originalVersion := Version
	originalBuildTime := BuildTime

	assert.NotEmpty(t, originalVersion, "Version should be set")
	assert.NotEmpty(t, originalBuildTime, "BuildTime should be set")

	info := GetBuildInfo()
	assert.Contains(t, info, originalVersion, "Build info should contain version")
	assert.Contains(t, info, originalBuildTime, "Build info should contain build time")
}
