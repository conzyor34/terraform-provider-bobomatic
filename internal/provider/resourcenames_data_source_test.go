package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourcenamesDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + `data "bobomatic_resourcenames" "test" {
					bu          = "bu"
					purpose     = "purpose"
					region      = "region"
					geography   = "geography"
					client      = "client"
					program     = "program"
					environment = "environment"
					index       = 1
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of Resourcenames returned
					resource.TestCheckResourceAttr("data.bobomatic_resourcenames.test", "resourcegroup", "bu-purpose-geography-client-program-environment-rg-001"),
					resource.TestCheckResourceAttr("data.bobomatic_resourcenames.test", "subscription", "bu-purpose-geography-client-program-environment-sub-001"),
				),
			},
		},
	})
}
