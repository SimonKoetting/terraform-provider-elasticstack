variable "policy_name" {
  description = "The integration policy name"
  type        = string
}

provider "elasticstack" {
  elasticsearch {}
  kibana {}
}

resource "elasticstack_fleet_agent_policy" "test" {
  name      = "${var.policy_name}-agent-policy"
  namespace = "default"
}

resource "elasticstack_fleet_elastic_defend_integration_policy" "test" {
  name                = var.policy_name
  namespace           = "default"
  agent_policy_id     = elasticstack_fleet_agent_policy.test.policy_id
  enabled             = true
  integration_version = "8.14.0"
  preset              = "EDRComplete"

  policy = {
    windows = {
      events = {
        process = true
        network = true
        file    = true
        dns     = true
      }
      malware = {
        mode          = "prevent"
        blocklist     = true
        notify_user   = true
        on_write_scan = true
      }
      ransomware = {
        mode = "prevent"
      }
      memory_protection = {
        mode = "detect"
      }
      behavior_protection = {
        mode               = "prevent"
        reputation_service = true
      }
      logging = {
        file = "info"
      }
      advanced = {
        agent = {
          connection_delay = 90
        }
        alerts = {
          cloud_lookup = true
          hash = {
            md5  = true
            sha1 = false
          }
        }
      }
    }
    mac = {
      events = {
        process = true
        file    = true
      }
      malware = {
        mode = "prevent"
      }
      memory_protection = {
        mode = "prevent"
      }
      behavior_protection = {
        mode               = "detect"
        reputation_service = true
      }
      logging = {
        file = "warning"
      }
      advanced = {
        agent = {
          connection_delay = 45
        }
        alerts = {
          cloud_lookup = false
          hash = {
            md5  = false
            sha1 = true
          }
        }
      }
    }
    linux = {
      events = {
        process      = true
        network      = true
        file         = true
        session_data = true
        tty_io       = false
      }
      malware = {
        mode      = "detect"
        blocklist = true
      }
      memory_protection = {
        mode = "prevent"
      }
      behavior_protection = {
        mode               = "detect"
        reputation_service = true
      }
      logging = {
        file = "warning"
      }
      advanced = {
        agent = {
          connection_delay = 30
        }
        alerts = {
          hash = {
            md5  = true
            sha1 = true
          }
        }
      }
    }
  }
}
