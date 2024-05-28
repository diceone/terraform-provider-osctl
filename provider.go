package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
				Sensitive: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"osctl_service": resourceService(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"osctl_stats": dataSourceStats(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		ApiUrl:  d.Get("api_url").(string),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
	}
	return config.Client()
}
