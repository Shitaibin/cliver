package version

// Step 1: 创建version包，设置变量和Get Version的方法

import (
	"fmt"
	"runtime"
)

var (
	Version    string
	GitVersion string
	GitCommit  string
	BuildDate  string
	GoVersion  = runtime.Version()
	Platform   = runtime.GOOS + "/" + runtime.GOARCH
)

func Get() string {
	return fmt.Sprintf("{Version: %s, GitVersion: %s, GitCommit: %s, "+
		"BuildDate: %s, GoVersion: %s, Platform: %s}", Version, GitVersion,
		GitCommit, BuildDate, GoVersion, Platform)
}
