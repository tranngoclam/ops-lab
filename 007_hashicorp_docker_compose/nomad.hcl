region = "vietnam"

data_dir = "/var/lib/nomad/"

bind_addr = "127.0.0.1"

addresses {
  http = "__IP_ADDRESS__"
  rpc = "__IP_ADDRESS__"
  serf = "__IP_ADDRESS__"
}

advertise {
  http = "__IP_ADDRESS__:4646"
  rpc = "__IP_ADDRESS__:4647"
  serf = "__IP_ADDRESS__:4648"
}

log_level = "DEBUG"

enable_debug = true

consul {
  # Consult agent's HTTP Address
  address = "127.0.0.1:8500"

  # The service name to register the server and client with Consul.
  server_service_name = "nomad"
  client_service_name = "nomad-client"

  # Advertise Nomad services to Consul
  # Enables automatically registering the services
  auto_advertise = true

  # Enables the servers and clients bootstrap using Consul
  server_auto_join = true
  client_auto_join = true
}
