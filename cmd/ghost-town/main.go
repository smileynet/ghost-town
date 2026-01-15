package main

import (
	"fmt"
	"os"

	"github.com/smileynet/ghost-town/internal/version"
)

var (
	Version   = version.Version
	BuildTime = version.BuildTime
	Commit    = version.Commit
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println(version.GetBuildInfo())
		os.Exit(0)
	}

	fmt.Println("ghost-town - kiro-cli tool based on Gas Town")
	fmt.Printf("Version: %s\n", version.GetVersion())
	fmt.Println("Early development - coming soon!")
	os.Exit(0)
}
