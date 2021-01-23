# Scaling and Load Balancing with HAProxy

- This blueprint uses docker-compose to scale a container to multiple containers and put HAProxy as a load balancer in front of them.
- If a container is down, we can spin up a new one, and the load is still balanced.
- `docker-compose up --scale SERVICE=NUM` is the command to scale a service to ${NUM} different containers.  

## Up

```bash
$ make up
...
server_3   | 2020/05/25 14:54:38 [info] server is listening on :8000
server_2   | 2020/05/25 14:54:38 [info] server is listening on :8000
server_1   | 2020/05/25 14:54:38 [info] server is listening on :8000
server_5   | 2020/05/25 14:54:38 [info] server is listening on :8000
server_4   | 2020/05/25 14:54:38 [info] server is listening on :8000
haproxy_1  | [NOTICE] 145/145438 (1) : New worker #1 (6) forked
```

## Test

> This test makes curl command to the server through the load balancer every second and prints the response to console, which is the hostname of container.

```bash
$ make test
c1452817b456
45a69813b21b
ea75b072f47a
7705f40eec57
5b17d63340e1
c1452817b456
45a69813b21b
ea75b072f47a
7705f40eec57
5b17d63340e1
```

> Logs from docker-compose up command will have more entries

```bash
...
server_2   | 2021/01/23 11:40:32 [info] server received request
haproxy_1  | <134>Jan 23 11:40:32 haproxy[9]: 192.168.0.1:62612 [23/Jan/2021:11:40:32.651] proxy server/server1 0/0/0/1/1 200 137 - - ---- 1/1/0/0/0 0/0 {e8616fab-dff0-493c-9983-ebf7250c} "GET / HTTP/1.1" 1
server_4   | 2021/01/23 11:40:33 [info] server received request
haproxy_1  | <134>Jan 23 11:40:33 haproxy[9]: 192.168.0.1:62624 [23/Jan/2021:11:40:33.669] proxy server/server2 0/0/0/1/1 200 137 - - ---- 1/1/0/0/0 0/0 {e8616fab-dff0-493c-9983-ebf7250c} "GET / HTTP/1.1" 1
server_5   | 2021/01/23 11:40:34 [info] server received request
haproxy_1  | <134>Jan 23 11:40:34 haproxy[9]: 192.168.0.1:62634 [23/Jan/2021:11:40:34.689] proxy server/server3 0/0/0/1/1 200 137 - - ---- 1/1/0/0/0 0/0 {e8616fab-dff0-493c-9983-ebf7250c} "GET / HTTP/1.1" 1
server_3   | 2021/01/23 11:40:35 [info] server received request
haproxy_1  | <134>Jan 23 11:40:35 haproxy[9]: 192.168.0.1:62646 [23/Jan/2021:11:40:35.708] proxy server/server4 0/0/0/1/1 200 137 - - ---- 1/1/0/0/0 0/0 {e8616fab-dff0-493c-9983-ebf7250c} "GET / HTTP/1.1" 1
server_1   | 2021/01/23 11:40:36 [info] server received request
haproxy_1  | <134>Jan 23 11:40:36 haproxy[9]: 192.168.0.1:62656 [23/Jan/2021:11:40:36.730] proxy server/server5 0/0/0/1/1 200 137 - - ---- 1/1/0/0/0 0/0 {e8616fab-dff0-493c-9983-ebf7250c} "GET / HTTP/1.1" 1
```

## Down

```bash
$ make down
Removing 001_haproxy_scaling_containers_server_4  ... done
Removing 001_haproxy_scaling_containers_server_5  ... done
Removing 001_haproxy_scaling_containers_server_1  ... done
Removing 001_haproxy_scaling_containers_server_2  ... done
Removing 001_haproxy_scaling_containers_server_3  ... done
Removing 001_haproxy_scaling_containers_haproxy_1 ... done
Removing network 001_haproxy_scaling_containers_default
Removing image 001_haproxy_scaling_containers_server
```

## References:

- [HAProxy Log Customization](https://www.haproxy.com/blog/haproxy-log-customization/)
- [How to handle HAProxy log format](https://uala.io/how-to-handle-haproxy-log-format/)
