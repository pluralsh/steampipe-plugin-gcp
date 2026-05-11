package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	"github.com/pluralsh/steampipe-plugin-gcp/gcp"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: gcp.Plugin})
}
