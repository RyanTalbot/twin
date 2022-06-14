package utils

import "fmt"

var (
	ErrBranchDoesNotExist = fmt.Errorf("ERR: The repo branch provided does not exist\n")
)
