package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDataSourceMwsNetworkConnectivityConfigTest(t *testing.T) {
	loadWorkspaceEnv(t)
	if isGcp(t) {
		skipf(t)("GCP not supported")
	}
	var sourceRegion string
	if isAzure(t) {
		sourceRegion = "eastus2"
	} else if isAws(t) {
		sourceRegion = "us-east-2"
	}
	AccountLevel(t,
		Step{
			Template: fmt.Sprintf(`
			resource "databricks_mws_network_connectivity_config" "this" {
				name = "tf-{var.RANDOM}"
				region = "%s"
			}	

			data "databricks_mws_network_connectivity_config" "this" {
			  depends_on = [databricks_mws_network_connectivity_config.this]
			  name = databricks_mws_network_connectivity_config.this.name
			}`, sourceRegion),
			Check: func(s *terraform.State) error {
				r, ok := s.RootModule().Resources["data.databricks_mws_network_connectivity_config.this"]
				if !ok {
					return fmt.Errorf("data not found in state")
				}
				name := r.Primary.Attributes["name"]
				if name == "" {
					return fmt.Errorf("name is empty: %v", r.Primary.Attributes)
				}
				expect := sourceRegion
				region := r.Primary.Attributes["region"]
				if region != expect {
					return fmt.Errorf("incorrect region. expected: %v, received: %v",
						expect, region)
				}
				return nil
			},
		})
}
