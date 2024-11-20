
provider "ksql" {
  # Configuration options
}

resource "ksql_stream" "test" {
   name = "test1"
   query = "(key VARCHAR KEY) WITH (KAFKA_TOPIC='topic', VALUE_FORMAT='JSON');"
}

