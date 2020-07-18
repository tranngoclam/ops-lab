# Increase log verbosity
log_level = "DEBUG"

# Setup data dir
data_dir = "/data/server"

# Give the agent a unique name. Defaults to hostname
name = "server"

# Enable the server
server {
  enabled = true

  # Self-elect, should be 3 or 5 for production
  bootstrap_expect = 1
}
