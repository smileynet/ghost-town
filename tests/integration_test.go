package integration_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	testinghelpers "github.com/smileynet/ghost-town/internal/testing"
)

func TestBuildBinary(t *testing.T) {
	testinghelpers.SkipIntegration(t)

	buildPath := "/tmp/ghost-town-test"

	cmd := exec.Command("go", "build", "-o", buildPath, "../cmd/ghost-town")
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "Failed to build binary: %v\nOutput: %s", err, output)

	require.FileExists(t, buildPath, "Binary should exist after build")

	defer os.Remove(buildPath)
}

func TestBinaryVersion(t *testing.T) {
	testinghelpers.SkipIntegration(t)

	buildPath := "/tmp/ghost-town-test"

	cmd := exec.Command("go", "build", "-o", buildPath, "../cmd/ghost-town")
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "Failed to build binary: %v\nOutput: %s", err, output)
	defer os.Remove(buildPath)

	cmd = exec.Command(buildPath, "version")
	output, err = cmd.CombinedOutput()
	require.NoError(t, err, "Failed to run version command: %v", err)
	assert.NotEmpty(t, string(output), "Version command should produce output")
}

func TestBinaryHelp(t *testing.T) {
	testinghelpers.SkipIntegration(t)

	buildPath := "/tmp/ghost-town-test"

	cmd := exec.Command("go", "build", "-o", buildPath, "../cmd/ghost-town")
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "Failed to build binary: %v\nOutput: %s", err, output)
	defer os.Remove(buildPath)

	// Binary with no args prints usage and exits with status 1
	cmd = exec.Command(buildPath)
	output, err = cmd.CombinedOutput()

	// We expect an exit error (status 1) when run without args
	assert.Error(t, err, "Binary should exit with error when run without args")

	helpOutput := string(output)
	assert.NotEmpty(t, helpOutput, "Help should produce output")
	assert.Contains(t, helpOutput, "ghost-town", "Help should contain program name")
	assert.Contains(t, helpOutput, "Usage:", "Help should contain usage section")
}

func TestBinaryExecution(t *testing.T) {
	testinghelpers.SkipIntegration(t)

	buildPath := "/tmp/ghost-town-test"

	cmd := exec.Command("go", "build", "-o", buildPath, "../cmd/ghost-town")
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "Failed to build binary: %v\nOutput: %s", err, output)
	defer os.Remove(buildPath)

	cmd = exec.Command(buildPath, "invalid-command")
	output, err = cmd.CombinedOutput()

	assert.Error(t, err, "Invalid command should return error")
	assert.Contains(t, string(output), "Unknown command", "Error message should indicate unknown command")
}

func TestBuildWithLdflags(t *testing.T) {
	testinghelpers.SkipIntegration(t)

	buildPath := "/tmp/ghost-town-test"

	cmd := exec.Command("go", "build", "-ldflags", "-X main.Version=1.0.0", "-o", buildPath, "../cmd/ghost-town")
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "Failed to build binary with ldflags: %v\nOutput: %s", err, output)
	defer os.Remove(buildPath)

	require.FileExists(t, buildPath, "Binary should exist after build with ldflags")
}
