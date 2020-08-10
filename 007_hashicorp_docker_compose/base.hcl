region = "vietnam"

data_dir = "/tmp/nomad/"

bind_addr = "0.0.0.0"

addresses {
  http = "{{ GetPrivateIP }}"
  rpc = "{{ GetPrivateIP }}"
  serf = "{{ GetPrivateIP }}"
}

advertise {
  http = "{{ GetPrivateIP }}:4646"
  rpc = "{{ GetPrivateIP }}:4647"
  serf = "{{ GetPrivateIP }}:4648"
}

log_level = "INFO"

enable_debug = true

consul {
  # Consult agent's HTTP Address
  address = "127.0.0.1:8500"

  # The service name to register the server and client with Consul.
  server_service_name = "nomad"
  client_service_name = "nomad-client"

  # Auth info for http access
  #auth = user:password

  # Advertise Nomad services to Consul
  # Enables automatically registering the services
  auto_advertise = true

  # Enables the servers and clients bootstrap using Consul
  server_auto_join = true
  client_auto_join = true
}
