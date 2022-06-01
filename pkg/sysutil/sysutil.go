package sysutil

import "os/exec"

func IsGitInstalled() bool {
	path, _ := exec.LookPath("git")
	if path != "" {
		return true
	}
	return false
}
