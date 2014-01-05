package lxc

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packager/common"
	"github.com/mitchellh/packager/packer"
)

// The unique ID for this builder.
const BuilderId = "kelseyhightower.lxc"

// Builder represents a Packer Builder.
type Builder struct {
	config config
	runner multistep.Runner
}

// config holds the lxc builder configuration settings.
type config struct {
	SSHUsername     string `mapstructure:"ssh_username"`
	SSHPassword     string `mapstructure:"ssh_password"`
	SSHPort         uint   `mapstructure:"ssh_port"`
	RawSSHTimeout   string `mapstructure:"ssh_timeout"`
	RawStateTimeout string `mapstructure:"state_timeout"`
}

// Prepare processes the build configuration parameters.
func (b *Builder) Prepare(raws ...interface{}) ([]string, error) {
	// Load the packer config.
	md, err := common.DecodeConfig(&b.config, raws...)
	if err != nil {
		return nil, err
	}
	b.config.tpl, err = packer.NewConfigTemplate()
	if err != nil {
		return nil, err
	}
	b.config.tpl.UserVars = b.config.PackerUserVars

	errs := common.CheckUnusedConfig(md)
	// Collect errors if any.
	if err := common.CheckUnusedConfig(md); err != nil {
		return nil, err
	}

	// Set defaults.
	if b.config.SSHUsername == "" {
		b.config.SSHUsername = "root"
	}
	if b.config.SSHPort == 0 {
		b.config.SSHPort = 22
	}
	return nil, err
}

// Run executes a lxc Packer build and returns a packer.Artifact
// representing a lxc container image.
func (b *Builder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
	// Not yet.
}
