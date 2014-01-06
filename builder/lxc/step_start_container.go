package lxc

import (
	"fmt"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

type StepStartContainer struct{}

func (s *StepStartContainer) Run(state multistep.StateBag) multistep.StepAction {
	driver := state.Get("driver").(Driver)
	ui := state.Get("ui").(packer.Ui)
	containerName := state.Get("container_name").(string)

	ui.Say("Starting the container...")

	// Start the container.
	err := driver.StartContainer(containerName)
	if err != nil {
		err := fmt.Errorf("Error starting the container: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (s *StepStartContainer) Cleanup(state multistep.StateBag) {}
