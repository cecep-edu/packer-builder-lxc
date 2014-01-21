package lxc

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/mitchellh/packer/packer"
)

// LXCDriver.
type LXCDriver struct {
	Ui  packer.Ui
	Tpl *packer.ConfigTemplate
}

func (l *LXCDriver) CreateTarball(src, dest string) error {
	args := []string{"-c", "-S", "--numeric-owner", "-f", dest, "-C", src, "rootfs"}
	return exec.Command("tar", args...).Run()
}

// CloneContainer.
func (l *LXCDriver) CloneContainer(c *ContainerConfig) error {
	args := []string{
		"--orig", c.OrigContainerName,
		"--new", c.NewContainerName,
		"--backingstore", "dir",
	}
	cmd := exec.Command("lxc-clone", args...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	l.Ui.Message(fmt.Sprintf("Running lxc-clone %s", strings.Join(args, " ")))
	// Clone the container
	log.Printf("Cloning the container...")
	if err := cmd.Start(); err != nil {
		return err
	}
	io.Copy(os.Stdout, stdout)
	io.Copy(os.Stderr, stderr)

	log.Printf("Waiting for container to finish cloning")
	if err := cmd.Wait(); err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			err = fmt.Errorf("lxc-clone exited with a non-zero exit status\n")
		}
		return err
	}
	return nil
}

// DestroyContainer.
func (l *LXCDriver) DestroyContainer(name string) error {
	return exec.Command("lxc-destroy", "-f", "-n", name).Run()
}

// StartContainer.
func (l *LXCDriver) StartContainer(name string) error {
	return exec.Command("lxc-start", "-d", "-n", name).Run()
}

// StopContainer.
func (l *LXCDriver) StopContainer(name string) error {
	return exec.Command("lxc-stop", "-n", name).Run()
}

// Verify.
func (l *LXCDriver) Verify() error {
	if _, err := exec.LookPath("lxc-version"); err != nil {
		return err
	}
	return nil
}
