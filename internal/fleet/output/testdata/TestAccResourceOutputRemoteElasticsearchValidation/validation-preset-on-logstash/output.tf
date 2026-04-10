provider "elasticstack" {
  elasticsearch {}
  kibana {}
}

resource "elasticstack_fleet_output" "test_output" {
  name      = "Logstash Output ${var.policy_name}"
  output_id = "${var.policy_name}-logstash-output"
  type      = "logstash"
  preset    = "throughput"

  hosts = [
    "logstash:5044",
  ]
}
