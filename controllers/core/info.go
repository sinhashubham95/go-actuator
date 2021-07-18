package core

import (
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/flags"
	"runtime"
)

var (
	BuildStamp       string
	GitCommitAuthor  string
	GitCommitID      string
	GitCommitTime    string
	GitPrimaryBranch string
	GitURL           string
	HostName         string
	Username         string
)

// GetInfo is used to get the basic information for the running application
func GetInfo() map[string]interface{} {
	return map[string]interface{}{
		commons.Application: map[string]string{
			commons.Env:     flags.Env(),
			commons.Name:    flags.Name(),
			commons.Version: flags.Version(),
		},
		commons.Git: map[string]string{
			commons.BuildStamp:       BuildStamp,
			commons.GitCommitAuthor:  GitCommitAuthor,
			commons.GitCommitID:      GitCommitID,
			commons.GitCommitTime:    GitCommitTime,
			commons.GitPrimaryBranch: GitPrimaryBranch,
			commons.GitURL:           GitURL,
			commons.HostName:         HostName,
			commons.Username:         Username,
		},
		commons.Runtime: map[string]interface{}{
			commons.Arch:           runtime.GOARCH,
			commons.OS:             runtime.GOOS,
			commons.Port:           flags.Port(),
			commons.RuntimeVersion: runtime.Version(),
		},
	}
}
