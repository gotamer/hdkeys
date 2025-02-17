package build

import "fmt"

var (
	AppName = "NoName"
	Version = "dev"
	CommitHash = "n/a"
    BuildTime = "n/a"
    UserName= "n/a"
)

func Info() string {
	return fmt.Sprintf("%s %s-%s (%s) by %s", AppName, Version, CommitHash, BuildTime, UserName)
}
