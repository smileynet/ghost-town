package testing

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TempDir creates a temporary directory for testing
func TempDir(t *testing.T) string {
	t.Helper()

	dir, err := os.MkdirTemp("", "ghost-town-*")
	require.NoError(t, err, "Failed to create temp directory")

	t.Cleanup(func() {
		os.RemoveAll(dir)
	})

	return dir
}

// TempFile creates a temporary file for testing
func TempFile(t *testing.T, content string) string {
	t.Helper()

	dir := TempDir(t)
	file := filepath.Join(dir, "test.txt")

	err := os.WriteFile(file, []byte(content), 0644)
	require.NoError(t, err, "Failed to create temp file")

	return file
}

// AssertEqual wraps testify assertion with better error message
func AssertEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	assert.Equal(t, expected, actual, msgAndArgs...)
}

// RequireEqual wraps testify require with better error message
func RequireEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	require.Equal(t, expected, actual, msgAndArgs...)
}

// AssertNoError wraps testify assertion with better error message
func AssertNoError(t *testing.T, err error, msgAndArgs ...interface{}) {
	t.Helper()
	assert.NoError(t, err, msgAndArgs...)
}

// RequireNoError wraps testify require with better error message
func RequireNoError(t *testing.T, err error, msgAndArgs ...interface{}) {
	t.Helper()
	require.NoError(t, err, msgAndArgs...)
}

// Contains checks if a string slice contains a string
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// SkipIntegration skips test if SKIP_INTEGRATION env var is set
func SkipIntegration(t *testing.T) {
	t.Helper()

	if os.Getenv("SKIP_INTEGRATION") == "true" {
		t.Skip("Skipping integration test (SKIP_INTEGRATION=true)")
	}
}

// SkipShort skips test if running short tests
func SkipShort(t *testing.T) {
	t.Helper()

	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}
}

// TableTest is a helper for table-driven tests
type TableTest struct {
	Name     string
	Input    interface{}
	Expected interface{}
	Error     bool
}

// RunTableTest runs a table-driven test
func RunTableTest(t *testing.T, tests []TableTest, testFunc func(t *testing.T, tt TableTest)) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			testFunc(t, tt)
		})
	}
}
