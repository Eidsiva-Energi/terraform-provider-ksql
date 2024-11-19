package ksql

import (
	"log"
	"testing"

	"github.com/Mongey/terraform-provider-kafka/kafka"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func setup() error {
	kafkaConfig := &kafka.Config{
		BootstrapServers: &[]string{"localhost:9092"},
		Timeout:          900,
	}
	kAdmin, err := kafka.NewClient(kafkaConfig)
	if err == nil {
		topic := kafka.Topic{
			Name:              "topic",
			Partitions:        1,
			ReplicationFactor: 1,
		}
		err = kAdmin.CreateTopic(topic)

		if err != nil {
			log.Printf("[ERROR] Creating Topic: %v", err)
			return err
		}
	} else {
		log.Printf("[ERROR] Unable to create client: %s", err)
	}
	return err
}

func TestAccOrderResource(t *testing.T) {
	err := setup()
	if err != nil {
		log.Printf("[ERROR] Unable to setup topic for test: %s", err)
	}

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
resource "ksql_stream" "test" {
   name = "test1"
   query = "(key VARCHAR KEY) WITH (KAFKA_TOPIC='topic', VALUE_FORMAT='JSON');"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ksql_stream.test", "name", "test1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
