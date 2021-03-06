package osc

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccAWSCallerIdentity_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsCallerIdentityConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsCallerIdentityAccountId("data.aws_caller_identity.current"),
				),
			},
		},
	})
}

func testAccCheckAwsCallerIdentityAccountId(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find AccountID resource: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Account Id resource ID not set.")
		}

		expected := testAccProvider.Meta().(*AWSClient).accountid
		if rs.Primary.Attributes["account_id"] != expected {
			return fmt.Errorf("Incorrect Account ID: expected %q, got %q", expected, rs.Primary.Attributes["account_id"])
		}

		return nil
	}
}

const testAccCheckAwsCallerIdentityConfig_basic = `
data "aws_caller_identity" "current" { }
`
