package github

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"twin/internal/utils"
)

type GitHubResponse struct {
	DefaultBranch string `json:"default_branch"`
}

func VerifyBranchName(userName string, repoName string, branchName string) (bool, error) {
	response, _ := http.Get("https://api.github.com/repos/" + userName + "/" + repoName + "/branches" + branchName)
	if response.StatusCode == http.StatusNotFound {
		return false, utils.ErrBranchDoesNotExist
	}
	return true, nil
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
