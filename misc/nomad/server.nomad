job "server" {
  datacenters = [
    "dc1"
  ]

  group "server" {
    count = 1

    constraint {
      operator = "distinct_hosts"
      value    = "true"
    }

    task "server" {
      driver = "docker"

      env {
        PORT = "${NOMAD_PORT_http}"
        NODE_IP = "${NOMAD_IP_http}"
      }

      config {
        image = "docker.io/library/server:v1.0.0"
      }

      resources {
        network {
          mbits = 10
          port "http" {}
        }
      }

      service {
        name = "server"
        port = "http"

        check {
          type = "http"
          path = "/"
          interval = "2s"
          timeout = "2s"
        }
      }
    }
  }
}
