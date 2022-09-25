package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	uptimerobotapi "github.com/lapras-inc/terraform-provider-uptimerobot/internal/provider/api"
)

func dataSourceAlertContact() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlertContactRead,
		Schema: map[string]*schema.Schema{
			"friendly_name": {Optional: true, Type: schema.TypeString},
			"id":            {Computed: true, Type: schema.TypeString},
			"type":          {Computed: true, Type: schema.TypeString},
			"status":        {Computed: true, Type: schema.TypeString},
			"value":         {Optional: true, Type: schema.TypeString},
		},
	}
}

func dataSourceAlertContactRead(d *schema.ResourceData, m interface{}) error {
	alertContacts, err := m.(uptimerobotapi.UptimeRobotApiClient).GetAlertContacts()
	if err != nil {
		return err
	}

	friendlyName := d.Get("friendly_name").(string)

	var alertContact uptimerobotapi.AlertContact

	for _, a := range alertContacts {
		if friendlyName != "" && a.FriendlyName == friendlyName {
			alertContact = a
			break
		}
	}
	if alertContact == (uptimerobotapi.AlertContact{}) {
		return fmt.Errorf("failed to find alert contact by name %s", friendlyName)
	}

	d.SetId(alertContact.ID)

	d.Set("friendly_name", alertContact.FriendlyName)
	d.Set("type", alertContact.Type)
	d.Set("status", alertContact.Status)
	d.Set("value", alertContact.Value)

	return nil
}
