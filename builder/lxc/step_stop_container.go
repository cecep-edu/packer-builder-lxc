package lxc

import (
	"fmt"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

type StepStopContainer struct{}

func (s *StepStopContainer) Run(state multistep.StateBag) multistep.StepAction {
	driver := state.Get("driver").(Driver)
	ui := state.Get("ui").(packer.Ui)
	containerName := state.Get("container_name").(string)

	ui.Say("Stoping the container...")

	// Start the container.
	err := driver.StopContainer(containerName)
	if err != nil {
		err := fmt.Errorf("Error starting the container: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (s *StepStopContainer) Cleanup(state multistep.StateBag) {}
