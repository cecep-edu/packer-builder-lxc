package lxc

import (
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/common"
)

type StepProvision struct{}

func (s *StepProvision) Run(state multistep.StateBag) multistep.StepAction {
	containerName := state.Get("container_name").(string)
	tempDir := state.Get("temp_dir").(string)

	// Create the communicator that talks to LXC.
	comm := &Communicator{
		ContainerName: containerName,
		ContainerDir:  "/var/lib/lxc/" + containerName + "/rootfs",
		HostDir:       tempDir,
	}

	prov := common.StepProvision{Comm: comm}
	return prov.Run(state)
}

func (s *StepProvision) Cleanup(state multistep.StateBag) {}
