package utils

import (
	"fmt"
	"strings"
)

func SplitGitHubURL(tail []string) (
	source map[string]string,
	err error,
) {
	if len(tail) == 0 {
		fmt.Println("A GitHub URL must be specified")
		return nil, err
	}

	source = make(map[string]string)

	source["repoURL"] = tail[0]
	splitURL := strings.Split(source["repoURL"], "/")
	source["userName"] = splitURL[3]
	source["repoName"] = splitURL[4]

	return source, nil
}
