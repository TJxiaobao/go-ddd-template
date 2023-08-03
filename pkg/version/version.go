package version

import "fmt"

var (
	GitRevision = "UNKNOWN"
	BuildTime   = "UNKNOWN"
	BuildUrl    = "UNKNOWN"

	Info = fmt.Sprintf("Git revision: %s\nBuild time: %s\nBuild Url: %s\n", GitRevision, BuildTime, BuildUrl)
)
