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
			return os.MkdirAll(fm.DestDir, 0755)
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
				// overwrite (delete) existing directory here
			}
			return err
		}
	}
	return nil
}
