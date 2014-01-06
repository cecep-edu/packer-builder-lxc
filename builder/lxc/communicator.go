package lxc

import (
	"io"
	"log"
	"os/exec"
	"syscall"

	"github.com/mitchellh/packer/packer"
)

// Communicator.
type Communicator struct {
	ContainerDir  string
	ContainerName string
	HostDir       string
}

// Start.
func (c *Communicator) Start(remote *packer.RemoteCmd) error {
	cmd := exec.Command("lxc-attach", "-n", c.ContainerName, "--", remote.Command)
	// Start the command
	log.Printf("Executing in container %s: %#v", c.ContainerName, remote.Command)
	if err := cmd.Start(); err != nil {
		log.Printf("Error executing: %s", err)
		remote.SetExited(254)
		return nil
	}
	err := cmd.Wait()
	if exitErr, ok := err.(*exec.ExitError); ok {
		exitStatus := 1
		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			exitStatus = status.ExitStatus()
		}
		remote.SetExited(exitStatus)
		return nil
	}
	return nil
}

// Upload.
func (c *Communicator) Upload(dst string, src io.Reader) error {
	panic("not implemented")
}

// UploadDir.
func (c *Communicator) UploadDir(dst string, src string, exclude []string) error {
	panic("not implemented")
}

// Download.
func (c *Communicator) Download(src string, w io.Writer) error {
	panic("not implemented")
}
