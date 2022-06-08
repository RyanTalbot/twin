package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"twin/internal/github"
	"twin/pkg/fileutil"
	"twin/pkg/sysutil"
)

func main() {
	gitInstalled := sysutil.IsGitInstalled()

	if gitInstalled == false {
		fmt.Println("Git not found")
	}

	prefs, init, _ := Flags()

	tempDir := ".twin-temp"
	err := os.Mkdir(".twin-temp", 0755)
	if err != nil {
		panic(err)
	}

	fileURL := "https://github.com/" + prefs["userName"] + "/" + prefs["repoName"] + "/archive/refs/heads/" + prefs["branchName"] + ".zip"
	zipString := []string{tempDir, "zip.zip"}
	zipFilePath := strings.Join(zipString, "/")
	fileutil.DownloadFile(zipFilePath, fileURL)

	unzipString := []string{tempDir, "unzipped"}
	unzipPath := strings.Join(unzipString, "/")
	fileutil.Unzip(zipFilePath, unzipPath)

	files, err := ioutil.ReadDir(unzipPath)
	if err != nil {
		panic(err)
	}
	repoDir := files[0]

	repoDirString := []string{unzipPath, repoDir.Name()}
	repoDirPath := strings.Join(repoDirString, "/")
	err = os.Rename(repoDirPath, "./"+prefs["destinationPath"])
	if err != nil {
		panic(err)
	}

	defer os.RemoveAll(tempDir)

	if gitInstalled && *init == false {
		github.InitializeRepository(prefs["destinationPath"])
	}
}
