package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type DownloadOptions struct {
	repo    string
	source  []string
	destdir string
}

func download(do DownloadOptions) error {
	fi, err := os.Stat(do.destdir)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		// TODO: should we set default DEST_DIR to CWD here?
		err = fmt.Errorf("%s is not a directory", do.destdir)
		return err
	}

	tmpdir, err := os.MkdirTemp("", "gitdl-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpdir)

	repoUrl := "https://github.com/" + do.repo + ".git"
	cmd := exec.Command("git", "clone", "--depth=1", "--filter=blob:none", "--no-checkout", repoUrl, tmpdir)
	if err = cmd.Run(); err != nil {
		return err
	}

	args := append([]string{"sparse-checkout", "set", "--no-cone"}, do.source...)
	cmd = exec.Command("git", args...)
	cmd.Dir = tmpdir
	if err = cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "checkout", "HEAD")
	cmd.Dir = tmpdir
	if err = cmd.Run(); err != nil {
		return err
	}

	for _, s := range do.source {
		_, ford := filepath.Split(s)
		src := filepath.Join(tmpdir, s)
		dst := filepath.Join(do.destdir, ford)
		if err = os.Rename(src, dst); err != nil {
			fmt.Printf("gitdl: %v\n", err)
			continue
		}
	}

	return nil
}
