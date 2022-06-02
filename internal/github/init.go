package github

import (
	"os"
	"os/exec"
)

func InitializeRepository(path string) error {
	gitPath, _ := exec.LookPath("git")

	command := exec.Cmd{
		Path:   gitPath,
		Args:   []string{path, "init", path},
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
	}

	err := command.Run()
	return err
}
