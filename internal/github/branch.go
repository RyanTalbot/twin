package github

import "net/http"

type GitHubResponse struct {
	DefaultBranch string `json:"default_branch"`
}

func VerifyBranchName(userName string, repoName string, branchName string) bool {
	response, err := http.Get("https://api.github.com/repos/" + userName + "/" + repoName + "/branches" + branchName)
	if err != nil {
		panic(err)
	}
	if response.StatusCode == http.StatusNotFound {
		return false
	}
	return true
}
