/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

/*
required parameteres

https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line
github token for modifying release

# dependencies
* requires a build environment (make artifact-release must work)
* github.com/itchio/gothub

*/

const (
	// TODO figure this out based on directory name
	repository   = "cluster-api-provider-aws"
	artifactsDir = "out"
)

func main() {
	// TODO consider moving all these things out to some config struct so other
	// providers can reuse this release
	expectedFiles := []string{
		// TODO these aren't examples. rename them.
		"cluster-api-provider-aws-examples.tar",
		"clusterawsadm-darwin-amd64",
		"clusterawsadm-linux-amd64",
		"clusterctl-darwin-amd64",
		"clusterctl-linux-amd64",
	}

	// TODO it would be ideal if we could release major/minor/patch and have it
	// automatically bump the latest tag git finds
	// until then, hard code it
	remote := os.Args[1] // default to origin probably
	version := os.Args[2]

	// Build all the release binaries
	RunCommand("make", "release-artifacts")

	// Creates git tag
	RunCommand("git", "tag", "--force", version)

	// Push git tag
	RunCommand("git", "push", remote, version)

	// create draft release
	// TODO: not everyone names the remotes after the username
	RunCommand("gothub", "release", "--tag", version, "--user", remote, "--repo", repository, "--draft")

	// attach tarball of yaml and binaries for systems to the release
	for _, file := range expectedFiles {
		RunCommand("gothub", "upload", "--tag", version, "--user", remote, "--repo", repository, "--file", path.Join(artifactsDir, file), "--name", file)
	}

	// TODO: automate writing the release notes
	// TODO: something something docker container
	// TODO: send an email or at least print out the contents of an email to
	// send and who it's going to.
}

// RunCommand runs a shell command and prints a log of what it's doing and when
// it finishes a step.
// TODO the error handling is bad and I should feel bad.
func RunCommand(cmd string, args ...string) {
	fmt.Printf("-> %v %v ", cmd, args)
	command := exec.Command(cmd, args...)
	err := command.Run()
	if err != nil {
		fmt.Printf("\nError: %v\n", err)
		out, err := command.CombinedOutput()
		if err != nil {
			panic(fmt.Sprintf("failed to get combined output: %v", err))
		}
		fmt.Println(string(out))
		os.Exit(1)
	}
}
