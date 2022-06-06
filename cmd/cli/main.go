package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"twin/internal/github"
	"twin/pkg/fileutil"
	"twin/pkg/sysutil"
)

func main() {
	var destinationPath string
	var branchName string

	gitInstalled := sysutil.IsGitInstalled()

	if gitInstalled == false {
		fmt.Println("Git not found")
	}

	branchNamePointer := flag.String("branch", "", "The name of the branch you want")
	outPointer := flag.String("out", "", "Destination path of the project")
	shouldNotInitPointer := flag.Bool("no-init", false, "If set, will not auto init")
	flag.Parse()

	tail := flag.Args()

	if len(tail) == 0 {
		fmt.Println("A GitHub URL must be specified")
		return
	}

	repoURL := tail[0]
	splitGitURL := strings.Split(repoURL, "/")
	userName := splitGitURL[3]
	repoName := splitGitURL[4]

	if *outPointer == "" {
		destinationPath = repoName
	} else {
		destinationPath = *outPointer
	}

	if *branchNamePointer == "" {
		mainBranchName, err := github.GetMainBranchName(userName, repoName)
		if err != nil {
			panic(err)
		}
		branchName = mainBranchName
	} else {
		existsBranch := github.VerifyBranchName(userName, repoName, *branchNamePointer)
		if !existsBranch {
			fmt.Printf("The branch %v not exists for repo %v\n", *branchNamePointer, repoURL)
			os.Exit(1)
		}
		branchName = *branchNamePointer
	}

	tempDir := ".twin-temp"
	err := os.Mkdir(".twin-temp", 0755)
	if err != nil {
		panic(err)
	}

	fileURL := "https://github.com/" + userName + "/" + repoName + "/archive/refs/heads/" + branchName + ".zip"
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
	err = os.Rename(repoDirPath, "./"+destinationPath)
	if err != nil {
		panic(err)
	}

	defer os.RemoveAll(tempDir)

	if gitInstalled && *shouldNotInitPointer == false {
		github.InitializeRepository(destinationPath)
	}
}
