package lxc

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

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
	wrappedRemoteCmd := fmt.Sprintf("lxc-attach -n %s -- /bin/sh -c '%s'", c.ContainerName, remote.Command)
	cmd := exec.Command("/bin/sh", "-c", wrappedRemoteCmd)
	cmd.Stdout = remote.Stdout
	cmd.Stderr = remote.Stderr

	// Start the command
	log.Printf("Executing in container %s: %#v", c.ContainerName, wrappedRemoteCmd)
	if err := cmd.Start(); err != nil {
		log.Printf("Error executing: %s", err)
		remote.SetExited(254)
		return nil
	}

	// run this in a goroutine so we don't block.
	go func() {
		exitStatus := 0
		err := cmd.Wait()
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitStatus = 1
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				exitStatus = status.ExitStatus()
			}
		}
		remote.SetExited(exitStatus)
	}()

	return nil
}

// Upload.
func (c *Communicator) Upload(dst string, r io.Reader) error {
	tempDest := filepath.Join(c.HostDir, filepath.Base(dst))
	f, err := os.Create(tempDest)
	if err != nil {
		return err
	}

	log.Printf("Uploading to container dir: %s\n", dst)
	//defer os.Remove(f.Name())
	io.Copy(f, r)
	f.Close()

	cmd := &packer.RemoteCmd{
		Command: fmt.Sprintf("cp /packer-files/%s %s", filepath.Base(dst), dst)}

	time.Sleep(time.Second * 5)
	if err := c.Start(cmd); err != nil {
		return err
	}

	// Wait for the copy to complete
	cmd.Wait()
	if cmd.ExitStatus != 0 {
		return fmt.Errorf("Upload failed with non-zero exit status: %d", cmd.ExitStatus)
	}

	log.Printf("Upload complete\n")
	return nil
}

// UploadDir.
func (c *Communicator) UploadDir(dst string, src string, exclude []string) error {
	panic("not implemented")
}

// Download.
func (c *Communicator) Download(src string, w io.Writer) error {
	panic("not implemented")
}
