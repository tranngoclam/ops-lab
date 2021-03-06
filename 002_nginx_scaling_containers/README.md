# Scaling and Load Balancing with NGINX

- This blueprint uses docker-compose to scale a container to multiple containers and put NGINX as a load balancer in front of them.
- If a container is down, we can spin up a new one, and the load is still balanced.
- `docker-compose up --scale SERVICE=NUM` is the command to scale a service to ${NUM} different containers.  

## Up

```bash
$ make up
...
server_1  | 2020/05/28 08:00:24 [info] server is listening on :8000
server_3  | 2020/05/28 08:00:24 [info] server is listening on :8000
server_2  | 2020/05/28 08:00:23 [info] server is listening on :8000
server_4  | 2020/05/28 08:00:23 [info] server is listening on :8000
server_5  | 2020/05/28 08:00:23 [info] server is listening on :8000
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
server_2   | 2020/05/25 14:54:39 [info] server received request
server_4   | 2020/05/25 14:54:40 [info] server received request
server_1   | 2020/05/25 14:54:41 [info] server received request
server_3   | 2020/05/25 14:54:42 [info] server received request
server_5   | 2020/05/25 14:54:43 [info] server received request
server_2   | 2020/05/25 14:54:44 [info] server received request
server_4   | 2020/05/25 14:54:45 [info] server received request
server_1   | 2020/05/25 14:54:46 [info] server received request
server_3   | 2020/05/25 14:54:47 [info] server received request
server_5   | 2020/05/25 14:54:48 [info] server received request
```

## Down

```bash
$ make down
Removing 002_nginx_scaling_containers_server_4  ... done
Removing 002_nginx_scaling_containers_server_5  ... done
Removing 002_nginx_scaling_containers_server_1  ... done
Removing 002_nginx_scaling_containers_server_2  ... done
Removing 002_nginx_scaling_containers_server_3  ... done
Removing 002_nginx_scaling_containers_haproxy_1 ... done
Removing network 002_nginx_scaling_containers_default
Removing image 002_nginx_scaling_containers_server
```
