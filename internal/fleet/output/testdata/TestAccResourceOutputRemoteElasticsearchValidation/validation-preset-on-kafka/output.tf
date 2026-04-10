provider "elasticstack" {
  elasticsearch {}
  kibana {}
}

resource "elasticstack_fleet_output" "test_output" {
  name      = "Kafka Output ${var.policy_name}"
  output_id = "${var.policy_name}-kafka-output"
  type      = "kafka"
  preset    = "throughput"

  hosts = [
    "kafka:9092",
  ]
}
