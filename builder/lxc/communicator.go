package lxc

import (
	"io"

	"github.com/mitchellh/packer/packer"
)

// Communicator.
type Communicator struct {
	ContainerDir  string
	ContainerName string
	HostDir       string
}

// Start.
func (c *Communicator) Start(cmd *packer.RemoteCmd) error {
	panic("not implemented")
}

// Upload.
func (c *Communicator) Upload(dst string, r io.Reader) error {
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
