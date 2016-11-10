package main

import (
	"github.com/docker/machine/drivers/packet"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(packet.NewDriver("", ""))
}
