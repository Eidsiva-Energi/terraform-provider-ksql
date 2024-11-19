package ksql

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// providerConfig is a shared configuration to combine with the actual
	// test configuration so the HashiCups client is properly configured.
	// It is also possible to use the HASHICUPS_ environment variables instead,
	// such as updating the Makefile and running the testing through that tool.
	providerConfig = `
provider "ksql" {
  url     = "http://localhost:8088"
}
`
)

var testAccProviders map[string]func() (*schema.Provider, error)
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = testAccProvidersFactory(testAccProvider)
}

func testAccProvidersFactory(provider *schema.Provider) map[string]func() (*schema.Provider, error) {
	return map[string]func() (*schema.Provider, error){
		"ksql": func() (*schema.Provider, error) {
			return provider, nil
		},
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("KSQLDB_URL"); v == "" {
		t.Fatal("KSQLDB_URL must be set for acceptance tests")
	}
}
