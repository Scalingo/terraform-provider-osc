package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/remijouannet/terraform-provider-osc/osc"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: osc.Provider,
	})
}
