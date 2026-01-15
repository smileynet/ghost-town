package integration_test

import (
	"os"
	"os/exec"
	"testing"
)

func TestBuildBinary(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION") == "true" {
		t.Skip("Skipping integration tests")
	}

	// Test that we can build the binary
	cmd := exec.Command("go", "build", "-o", "/tmp/ghost-town-test", "./cmd/ghost-town")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to build binary: %v\nOutput: %s", err, output)
	}

	// Clean up
	os.Remove("/tmp/ghost-town-test")
}

func TestBinaryVersion(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION") == "true" {
		t.Skip("Skipping integration tests")
	}

	// Build binary first
	cmd := exec.Command("go", "build", "-o", "/tmp/ghost-town-test", "./cmd/ghost-town")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove("/tmp/ghost-town-test")

	// Test version command
	cmd = exec.Command("/tmp/ghost-town-test", "version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run version command: %v", err)
	}

	if len(output) == 0 {
		t.Error("Version command should produce output")
	}
}
