package lxc

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

type StepCreateBindMount struct {
	SharedDir       string
	fstabPath       string
	fstabBackupPath string
}

func (s *StepCreateBindMount) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	containerPath := state.Get("container_path").(string)
	tempDir := state.Get("temp_dir").(string)

	ui.Say("Creating shared directory ...")

	// Create the /packer-files directory in the container path
	sharedDir := filepath.Join(containerPath, "rootfs", "/packer-files")
	err := os.Mkdir(sharedDir, 755)

	if err != nil {
		err := fmt.Errorf("Error creating shared directory: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	// Add bind mount to the lxc config
	fstabPath := filepath.Join(containerPath, "fstab")
	fstabBackupPath := filepath.Join(containerPath, "fstab-backup")
	os.Rename(fstabPath, fstabBackupPath)

	fstabBackup, err := os.Open(fstabBackupPath)
	if err != nil {
		err := fmt.Errorf("Error creating shared directory: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	defer fstabBackup.Close()

	fstab, err := os.Create(fstabPath)
	if err != nil {
		err := fmt.Errorf("Error creating shared directory: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	defer fstab.Close()

	scanner := bufio.NewScanner(fstab)
	for scanner.Scan() {
		fstab.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		err := fmt.Errorf("Error creating shared directory: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	entry := tempDir + " packer-files none bind 0 0\n"
	fstab.WriteString(entry)

	s.SharedDir = sharedDir
	s.fstabPath = fstabPath
	s.fstabBackupPath = fstabBackupPath

	return multistep.ActionContinue
}

func (s *StepCreateBindMount) Cleanup(state multistep.StateBag) {}
