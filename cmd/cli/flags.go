package main

import (
	"flag"
	"fmt"
	"os"
	"twin/internal/github"
	"twin/internal/utils"
)

func Flags() (
	flags map[string]string,
	noInitPointer *bool,
	err error,
) {
	branchNamePointer := flag.String("branch", "", "The name of the branch you want")
	outPointer := flag.String("out", "", "Destination path of the project")
	noInitPointer = flag.Bool("no-init", false, "If set, will not auto init")
	flag.Parse()

	tail := flag.Args()

	flags = make(map[string]string)

	source, _ := utils.SplitGitHubURL(tail)
	flags["repoURL"] = source["repoURL"]
	flags["userName"] = source["userName"]
	flags["repoName"] = source["repoName"]

	if *outPointer == "" {
		flags["destinationPath"] = source["repoName"]
	} else {
		flags["destinationPath"] = *outPointer
	}

	if *branchNamePointer == "" {
		mainBranchName, err := github.GetMainBranchName(flags["userName"], flags["repoName"])
		if err != nil {
			panic(err)
		}
		flags["branchName"] = mainBranchName
	} else {
		existsBranch := github.VerifyBranchName(flags["userName"], flags["repoName"], *branchNamePointer)
		if !existsBranch {
			fmt.Printf("The branch %v not exists for repo %v\n", *branchNamePointer, flags["repoURL"])
			os.Exit(1)
		}
		flags["branchName"] = *branchNamePointer
	}
	return flags, noInitPointer, nil
}
