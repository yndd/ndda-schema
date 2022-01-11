package main

import (
	"github.com/openconfig/ygot/ygot"
	schema "github.com/yndd/ndda-schema/pkg/nddaschema"
)

func main() {

	d := schema.Device{}
	d.Interface["eth1"] = &schema.NddaInterface_Interface{
		Name: ygot.String("eth1"),
		Config: &schema.NddaInterface_Interface_Config{
			Kind: schema.NddaCommon_InterfaceKind_INTERFACE,
		},
	}

	// Create a new interface called "eth0"
	i, _ := d.NewInterface("eth0")
	i.Config.Kind = schema.NddaCommon_InterfaceKind_INTERFACE
}
