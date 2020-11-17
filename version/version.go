package version

// Step 1: 创建version包，设置变量和Get Version的方法

import (
	"fmt"
	"runtime"
)

var (
	Version   string
	GitCommit string // GIT_COMMIT=$(git rev-list -1 HEAD)
	BuildDate string // BUILD_DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
	GoVersion = runtime.Version()
	Platform  = runtime.GOOS + "/" + runtime.GOARCH
)

func Get() string {
	return fmt.Sprintf("{Version: %s, GitCommit: %s, "+
		"BuildDate: %s, GoVersion: %s, Platform: %s}", Version,
		GitCommit, BuildDate, GoVersion, Platform)
}
