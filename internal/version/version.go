package version

import "fmt"

var (
	Version   = "dev"
	BuildTime = "unknown"
	Commit    = "unknown"
)

func GetVersion() string {
	return Version
}

func GetBuildInfo() string {
	return fmt.Sprintf("Version: %s\nBuild: %s\nCommit: %s", Version, BuildTime, Commit)
}

func String() string {
	return GetVersion()
}
