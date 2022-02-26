job "nginx" {
  datacenters = [
    "dc1"
  ]
  type = "service"

  update {
    max_parallel = 1
    min_healthy_time = "10s"
    healthy_deadline = "2m"
    progress_deadline = "5m"
    auto_revert = true
  }

  group "nginx" {
    count = 3

    constraint {
      operator = "distinct_hosts"
      value = "true"
    }

    restart {
      attempts = 1
      interval = "5m"
      delay = "25s"
      mode = "delay"
    }

    task "nginx" {
      driver = "docker"

      template {
        data = <<EOF
upstream backend {
{{ range service "demo-webapp" }}
  server {{ .Address }}:{{ .Port }};
{{ else }}server 127.0.0.1:65535; # force a 502
{{ end }}
}

server {
   listen 80;

   location / {
      proxy_pass http://backend;
   }
}
EOF

        destination = "local/load-balancer.conf"
        change_mode = "signal"
        change_signal = "SIGHUP"
      }

      template {
        data = <<EOH
A={{ key "A" }}
B={{ key "B" }}
C={{ key "C" }}
D={{ key "D" }}
E={{ key "E" }}
F={{ key "F" }}
G={{ key "G" }}
H={{ key "H" }}
I={{ key "I" }}
J={{ key "J" }}
K={{ key "K" }}
L={{ key "L" }}
M={{ key "M" }}
N={{ key "N" }}
O={{ key "O" }}
P={{ key "P" }}
Q={{ key "Q" }}
R={{ key "R" }}
S={{ key "S" }}
S={{ key "T" }}
S={{ key "U" }}
V={{ key "V" }}
W={{ key "W" }}
X={{ key "X" }}
Y={{ key "Y" }}
Z={{ key "Z" }}
EOH
        destination = "etc/env"
        env = true
      }

      config {
        image = "nginx:1.19.1-alpine"

        port_map {
          http = 80
        }

        volumes = [
          "local:/etc/nginx/conf.d",
        ]
      }

      resources {
        cpu = 256
        memory = 256

        network {
          mbits = 10
          port "http" {}
        }
      }

      service {
        name = "nginx"
        port = "http"
        tags = [
          "nginx"
        ]

        check {
          type = "tcp"
          interval = "1s"
          timeout = "5s"
        }
      }
    }
  }
}
