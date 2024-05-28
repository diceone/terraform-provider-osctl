package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceStats() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceStatsRead,

		Schema: map[string]*schema.Schema{
			"ram_usage": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disk_usage": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// Add more fields as needed
		},
	}
}

func dataSourceStatsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	resp, err := client.makeRequest("GET", "stats", nil)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}

	d.Set("ram_usage", result["ram_usage"])
	d.Set("disk_usage", result["disk_usage"])
	// Set more fields as needed

	d.SetId("osctl_stats")
	return nil
}
