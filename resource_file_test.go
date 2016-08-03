package main

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestResourceFile(t *testing.T) {
	provider := Provider()
	providers := map[string]terraform.ResourceProvider{
		"filesystem": provider,
	}
	c1 := `
		resource "filesystem_file" "test" {
			path = "/tmp/test.txt"
			contents = "alpha"
		}
	`
	c2 := `
		resource "filesystem_file" "test" {
			path = "/tmp/test.txt"
			contents = "beta"
		}
	`
	resource.UnitTest(t, resource.TestCase{
		Providers: providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: c1,
				Check:  resource.TestCheckResourceAttr("filesystem_file.test", "contents", "alpha"),
			},
			resource.TestStep{
				Config:      c2,
				ExpectError: regexp.MustCompile(`FAIL`),
				Check:       resource.TestCheckResourceAttr("filesystem_file.test", "contents", "alpha"),
			},
		},
	})
}
