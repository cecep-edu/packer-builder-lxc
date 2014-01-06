package lxc

// Driver.
type Driver interface {
	// CloneContainer clones a LXC container.
	CloneContainer(*ContainerConfig) error

	// DestroyContainer
	DestroyContainer(name string) error

	// StartContainer starts a LXC container.
	StartContainer(name string) error

	// StopContainer stops a LXC container.
	StopContainer(name string) error

	// Verify verifies that the driver can run.
	Verify() error
}

// ContainerConfig is the configuration used to start a container.
type ContainerConfig struct {
	NewContainerName  string
	OrigContainerName string
}
