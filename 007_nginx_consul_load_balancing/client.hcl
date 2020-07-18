# Increase log verbosity
log_level = "DEBUG"

# Setup data dir
data_dir = "/data/client"

# Give the agent a unique name. Defaults to hostname
name = "client"

# Enable the client
client {
  enabled = true

  # For demo assume we are talking to server1. For production,
  # this should be like "nomad.service.consul:4647" and a system
  # like Consul used for service discovery.
  servers = [
    "10.64.0.2:4647"
  ]
}

# Modify our port to avoid a collision with server1
ports {
  http = 5656
}
