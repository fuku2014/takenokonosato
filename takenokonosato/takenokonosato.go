package takenokonosato

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a schema.Provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"kinoko": {
				Description: "If true like kinokonoyama",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"takenokonosato_ichigoazi": NewIchigo(),
		},
	}
}
