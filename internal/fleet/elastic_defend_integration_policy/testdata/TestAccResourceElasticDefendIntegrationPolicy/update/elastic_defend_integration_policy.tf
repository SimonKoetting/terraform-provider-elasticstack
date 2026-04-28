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
  description         = "Updated description"
  agent_policy_id     = elasticstack_fleet_agent_policy.test.policy_id
  enabled             = false
  integration_version = "8.14.0"
  preset              = "EDRComplete"

  policy = {
    windows = {
      events = {
        process = true
        network = true
        file    = false
        dns     = false
      }
      malware = {
        mode          = "detect"
        blocklist     = false
        notify_user   = false
        on_write_scan = false
      }
      ransomware = {
        mode = "detect"
      }
      memory_protection = {
        mode = "prevent"
      }
      behavior_protection = {
        mode               = "detect"
        reputation_service = false
      }
      logging = {
        file = "error"
      }
      advanced = {
        agent = {
          connection_delay = 75
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
    mac = {
      events = {
        process = true
        network = false
        file    = false
      }
      malware = {
        mode = "detect"
      }
      memory_protection = {
        mode = "detect"
      }
      behavior_protection = {
        mode               = "prevent"
        reputation_service = false
      }
      logging = {
        file = "error"
      }
      advanced = {
        agent = {
          connection_delay = 35
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
    linux = {
      events = {
        process      = true
        network      = false
        file         = false
        session_data = false
        tty_io       = true
      }
      malware = {
        mode      = "prevent"
        blocklist = false
      }
      memory_protection = {
        mode = "detect"
      }
      behavior_protection = {
        mode               = "prevent"
        reputation_service = false
      }
      logging = {
        file = "error"
      }
      advanced = {
        agent = {
          connection_delay = 20
        }
        alerts = {
          hash = {
            md5  = false
            sha1 = true
          }
        }
      }
    }
  }
}
