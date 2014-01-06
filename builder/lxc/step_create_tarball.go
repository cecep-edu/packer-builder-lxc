package lxc

import (
	"fmt"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

type StepCreateTarball struct{}

func (s *StepCreateTarball) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*Config)
	driver := state.Get("driver").(Driver)
	ui := state.Get("ui").(packer.Ui)
	containerName := state.Get("container_name").(string)

	ui.Say("Creating the container tar...")

	// Start the container.
	err := driver.CreateTarball("/var/lib/lxc/"+containerName, config.ExportPath)
	if err != nil {
		err := fmt.Errorf("Error creating the tarball: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (s *StepCreateTarball) Cleanup(state multistep.StateBag) {}
