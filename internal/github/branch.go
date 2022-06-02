package github

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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

func GetMainBranchName(userName string, repoName string) (string, error) {
	response, err := http.Get("https://api.github.com/repos/" + userName + "/" + repoName)
	if err != nil {
		return "", err
	}

	var object GitHubResponse

	data, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(data, &object)
	if err != nil {
		return "", err
	}

	return object.DefaultBranch, nil
}
