package main

import (
	"fmt"
	"os"

	"github.com/smileynet/ghost-town/internal/cli"
	"github.com/smileynet/ghost-town/internal/version"
)

var (
	Version   = version.Version
	BuildTime = version.BuildTime
	Commit    = version.Commit
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "version", "-v", "--version":
		fmt.Println(version.GetBuildInfo())
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("ghost-town - kiro-cli tool based on Gas Town")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  ghost-town <command> [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  version      Show version information")
	fmt.Println()
	fmt.Println("For more information, run:")
	fmt.Println("  ghost-town <command> --help")
}
