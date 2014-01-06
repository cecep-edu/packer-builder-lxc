# LXC Builder

Type: `lxc`

The `lxc` builder is able to create new images for use with LXC.

## Status

Broken. Should be ready for testing 1/13/2014.

## TODO

* Finish work on LXC communicator -- proxy packer remote commands through lxc-attach.

## Building

Clone this repository into `$GOPATH/src/github.com/kelseyhightower/packer-builder-lxc`.  Then build the `packer-builder-lxc` binary:

```
cd $GOPATH/src/github.com/kelseyhightower/packer-builder-lxc
go get
go build
```

Copy the results to the Packer install directory.

```
cp packer-builder-lxc /usr/local/packer/packer-builder-lxc
```

## Configure

Enable the googlecompute builder in `~/.packerconfig`

```
{
  "builders": {
    "lxc": "/usr/local/packer/packer-builder-lxc"
  }
}
```

> See [configure Packer](http://www.packer.io/docs/other/core-configuration.html) for more info.

## Basic Example

```JSON
{
  "builders": [{
    "type": "lxc",
    "image": "ubuntu-base",
    "export_path": "/tmp/ubuntu-01062014.tar"
  }]
}
```

## Configuration Reference

The reference of available configuration options is listed below.

### Required parameters:

* `image` (string) - The LXC container to clone. This container must already exist.
* `export_path` (string) - The path where the final container will be exported as a tar file.
