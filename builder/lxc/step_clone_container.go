package lxc

import (
	"fmt"
	"path/filepath"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/common/uuid"
	"github.com/mitchellh/packer/packer"
)

type StepCloneContainer struct {
	containerName string
}

func (s *StepCloneContainer) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*Config)
	driver := state.Get("driver").(Driver)
	ui := state.Get("ui").(packer.Ui)

	ui.Say(fmt.Sprintf("Cloning the %s container...", config.Image))

	containerName := fmt.Sprintf("packer-%s", uuid.TimeOrderedUUID())
	cloneConfig := ContainerConfig{
		NewContainerName:  containerName,
		OrigContainerName: config.Image,
	}

	// Clone the container.
	err := driver.CloneContainer(&cloneConfig)
	if err != nil {
		err := fmt.Errorf("Error cloning container: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	s.containerName = containerName
	state.Put("container_path", filepath.Join("/var/lib/lxc", s.containerName))
	state.Put("container_name", s.containerName)
	ui.Message(fmt.Sprintf("Container Name: %s", s.containerName))
	return multistep.ActionContinue
}

func (s *StepCloneContainer) Cleanup(state multistep.StateBag) {
	if s.containerName == "" {
		return
	}

	driver := state.Get("driver").(Driver)
	ui := state.Get("ui").(packer.Ui)
	ui.Say("Destorying the container...")
	driver.DestroyContainer(s.containerName)
}
