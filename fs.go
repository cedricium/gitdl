package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type FileManager struct {
	DestDir string
}

func (fm *FileManager) ValidateDestDir() error {
	dstInfo, err := os.Stat(fm.DestDir)
	if err != nil {
		if os.IsNotExist(err) {
			// if destination does not exist, skip validation since the given
			// destination will be created in `fm.MoveFiles`. if created here,
			// `os.Rename` operation will fail due to an existing newpath
			return nil
		}
		return err
	}
	if !dstInfo.IsDir() {
		return fmt.Errorf("%s is not a directory", fm.DestDir)
	}
	return nil
}

func (fm *FileManager) CreateTempDir() (string, error) {
	tmpdir, err := os.MkdirTemp("", "gitdl-*")
	if err != nil {
		return "", err
	}
	return tmpdir, nil
}

func (fm *FileManager) MoveFiles(sources []string, tempDir string) error {
	for _, s := range sources {
		_, ford := filepath.Split(s)
		src := filepath.Join(tempDir, s)
		dst := filepath.Join(fm.DestDir, ford)
		if err := os.Rename(src, dst); err != nil {
			if errors.Is(err, os.ErrExist) {
				// if options.Force:
				// 	— delete existing directory dst
				// 	- return nil to continue with overwrite
				// else:
				return fmt.Errorf("directory with content at `%s' would be overwritten, skipping…", dst)
			}
			return err
		}
	}
	return nil
}
