# Explore HashiCorp with Docker Compose

- This blueprint uses docker-compose to demonstrate a cluster of Nomad, Consul.
- The cluster contains 3 machines: `north`, `middle` and `south`. They are defined in the [docker-compose.yml](docker-compose.yml) 

## Up

```bash
$ make up
```

## Down

```bash
$ make down
```

## Access

```text
Consul: http://localhost:8500
Nomad: http://localhost:4646
```

## Import

```bash
$ make import
```

## Deploy

```bash
$ make deploy
```
