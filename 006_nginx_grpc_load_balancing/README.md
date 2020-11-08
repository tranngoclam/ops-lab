# Load Balancing with NGINX and gRPC

- This blueprint uses docker-compose to spin up a fleet of application containers and put NGINX as a load balancer in front of them.
- NGINX uses Docker DNS to discover up-and-running application containers, if any of them goes down, Docker DNS will remove it from the registry.

## Up

```bash
$ make up
...
Attaching to 006_nginx_grpc_load_balancing_nginx_1, 006_nginx_grpc_load_balancing_server_4, 006_nginx_grpc_load_balancing_server_3, 006_nginx_grpc_load_balancing_server_1, 006_nginx_grpc_load_balancing_server_5, 006_nginx_grpc_load_balancing_server_2
nginx_1   | /docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
nginx_1   | /docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
nginx_1   | /docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
nginx_1   | 10-listen-on-ipv6-by-default.sh: error: can not modify /etc/nginx/conf.d/default.conf (read-only file system?)
nginx_1   | /docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
nginx_1   | /docker-entrypoint.sh: Configuration complete; ready for start up
```

## Test

> This test makes curl command to the server through the load balancer every second and prints the response to console, which is the hostname of container.

```bash
$ go run client/main -addr 127.0.0.1:80
Hello grpc from 56d0f207eb7a
Hello grpc from d26faa8f87e7
Hello grpc from 7f9dfcab15e2
...
```

> Logs from docker-compose up command will have more entries

```bash
...
server_1  | 2020-11-08T04:12:33.205Z	DEBUG	server/main.go:44	Hello grpc from 33616ddaa759
server_1  | 2020-11-08T04:12:33.205Z	INFO	zap/options.go:203	finished unary call with code OK	{"grpc.start_time": "2020-11-08T04:12:33Z", "system": "grpc", "span.kind": "server", "grpc.service": "helloworld.Greeter", "grpc.method": "SayHello", "grpc.code": "OK", "grpc.time_ms": 0.24500000476837158}
nginx_1   | 172.28.0.1 - - [08/Nov/2020:04:12:33 +0000] "POST /helloworld.Greeter/SayHello HTTP/2.0" 200 35 "-" "grpc-go/1.33.2" "-"
...
```

## Down

```bash
$ make down
```
