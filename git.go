package main

import "os/exec"

type Git struct {
	RepoURL string
	TempDir string
}

func (g *Git) Clone() error {
	cmd := exec.Command("git", "clone", "--depth=1", "--filter=blob:none", "--no-checkout", g.RepoURL, g.TempDir)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (g *Git) SparseCheckout(sources []string) error {
	args := append([]string{"sparse-checkout", "set", "--no-cache"}, sources...)
	cmd := exec.Command("git", args...)
	cmd.Dir = g.TempDir
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (g *Git) Checkout() error {
	cmd := exec.Command("git", "checkout", "HEAD")
	cmd.Dir = g.TempDir
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
