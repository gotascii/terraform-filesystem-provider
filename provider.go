package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider is the adapter for terraform, that gives access to all the resources
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: resourcesMap(),
	}
}

func resourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"filesystem_file": ResourceLocalFile(),
	}
}
