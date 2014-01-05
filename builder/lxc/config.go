package lxc

import (
	"time"

	"github.com/mitchellh/packer/common"
	"github.com/mitchellh/packer/packer"
)

// config holds the lxc builder configuration settings.
type Config struct {
	common.PackerConfig `mapstructure:",squash"`

	SSHUsername     string `mapstructure:"ssh_username"`
	SSHPassword     string `mapstructure:"ssh_password"`
	SSHPort         uint   `mapstructure:"ssh_port"`
	RawSSHTimeout   string `mapstructure:"ssh_timeout"`
	RawStateTimeout string `mapstructure:"state_timeout"`

	sshTimeout   time.Duration
	stateTimeout time.Duration
	tpl          *packer.ConfigTemplate
}

func NewConfig(raws ...interface{}) (*Config, []string, error) {
	c := new(Config)
	md, err := common.DecodeConfig(c, raws...)
	if err != nil {
		return nil, nil, err
	}

	c.tpl, err = packer.NewConfigTemplate()
	if err != nil {
		return nil, nil, err
	}
	c.tpl.UserVars = c.PackerUserVars

	// Prepare the rrors
	errs := common.CheckUnusedConfig(md)

	// Set defaults.
	if c.SSHUsername == "" {
		c.SSHUsername = "root"
	}
	if c.SSHPort == 0 {
		c.SSHPort = 22
	}

	// Check for any errors.
	if errs != nil && len(errs.Errors) > 0 {
		return nil, nil, errs
	}

	return c, nil, nil
}
