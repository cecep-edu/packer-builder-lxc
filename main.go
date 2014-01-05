package main

import (
	"github.com/kelseyhightower/packer-builder-lxc/buider/lxc"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {
	plugin.ServeBuilder(new(lxc.Builder))
}
