package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceService() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceCreate,
		Read:   resourceServiceRead,
		Update: resourceServiceUpdate,
		Delete: resourceServiceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	serviceName := d.Get("name").(string)
	action := d.Get("action").(string)

	_, err := client.makeRequest("POST", fmt.Sprintf("service/%s?action=%s", serviceName, action), nil)
	if err != nil {
		return err
	}
	d.SetId(serviceName)
	return resourceServiceRead(d, meta)
}

func resourceServiceRead(d *schema.ResourceData, meta interface{}) error {
	// Implement service read logic
	return nil
}

func resourceServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	// Implement service update logic
	return nil
}

func resourceServiceDelete(d *schema.ResourceData, meta interface{}) error {
	// Implement service delete logic
	return nil
}
