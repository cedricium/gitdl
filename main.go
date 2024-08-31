/*
Gitdl downloads files and directories locally from a remote git repository.

Usage:

	gitdl REPO SOURCE[…] DEST_DIR

The process (read: happy path) for fetching git blobs (files) and trees
(directories) and persisting those locally is described as follows:

 1. a temporary directory is created locally
 2. given target REPO is partially cloned to the temporary directory
 3. sparse checkout is enabled with only given SOURCE object(s) configured
 4. SOURCE object(s) are extracted and written to temporary directory using `git checkout`
 5. target's SOURCE files/directories copied to specified DEST_DIR
 6. temporary directory cleaned up/removed

Refer to git's [`sparse-checkout` documentation] and this [GitHub blog post] for
more information on the performance impact of the commands used.

[`sparse-checkout` documentation]: https://git-scm.com/docs/git-sparse-checkout#Documentation/git-sparse-checkout.txt
[GitHub blog post]: https://github.blog/2020-01-17-bring-your-monorepo-down-to-size-with-sparse-checkout/
*/
package main

import (
	"errors"
	"fmt"
	"os"
)

const usage string = `gitdl v0.0.3
Download files and directories locally from a remote git repository.

USAGE:
	gitdl REPO SOURCE[…] DEST_DIR

ARGUMENTS:
	REPO		short GitHub repo path "<owner>/<repo>" (e.g. 'cedricium/gitdl')
	SOURCE		file(s) and/or director(y|ies) from REPO wanting to download
	DEST_DIR	local destination directory where SOURCE args are copied to
`

func main() {
	config, err := ParseArgs(os.Args[1:])
	if err != nil {
		if errors.Is(err, ErrTooFewArguments) {
			fmt.Print(usage)
		} else {
			fmt.Printf("gitdl: %v\n", err)
		}
		os.Exit(1)
	}

	fm := FileManager{DestDir: config.DestDir}
	if err := fm.ValidateDestDir(); err != nil {
		fmt.Printf("gitdl: %v\n", err)
		os.Exit(1)
	}
	tempDir, err := fm.CreateTempDir()
	if err != nil {
		fmt.Printf("gitdl: %v\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(tempDir)

	git := Git{
		RepoURL: "https://github.com/" + config.Repo + ".git",
		TempDir: tempDir,
	}
	if err := git.Clone(); err != nil {
		fmt.Printf("gitdl: %v\n", err)
		os.Exit(1)
	}
	if err := git.SparseCheckout(config.Sources); err != nil {
		fmt.Printf("gitdl: %v\n", err)
		os.Exit(1)
	}
	if err := git.Checkout(); err != nil {
		fmt.Printf("gitdl: %v\n", err)
		os.Exit(1)
	}

	if err := fm.MoveFiles(config.Sources, tempDir); err != nil {
		fmt.Printf("gitdl: %v\n", err)
		os.Exit(1)
	}
}
