package main

import (
	"github.com/fuku2014/terraform-provider-takenokonosato/takenokonosato"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: takenokonosato.Provider,
	})
}
