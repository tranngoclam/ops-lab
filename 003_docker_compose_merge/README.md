# Use YAML Merge in Docker Compose

- [YAML Merge](https://yaml.org/type/merge.html) in docker-compose file helps to reuse shared configurations between docker containers
- How to define and apply shared configs:

```yaml
x-srv-config: &ref
//

<<: *ref
``` 

## Up

```bash
$ make up
WARNING: Native build is an experimental feature and could change at any time
Creating network "003_docker_compose_merge_default" with the default driver
Creating 003_docker_compose_merge_bar_1 ... done
Creating 003_docker_compose_merge_baz_1 ... done
Attaching to 003_docker_compose_merge_bar_1, 003_docker_compose_merge_baz_1
bar_1  | HOSTNAME=d6e6801cd5ac
bar_1  | SHLVL=1
bar_1  | HOME=/root
bar_1  | PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
bar_1  | BAR=bar
bar_1  | FOO=foo
bar_1  | PWD=/
baz_1  | HOSTNAME=a42be1aedd5e
baz_1  | SHLVL=1
baz_1  | HOME=/root
baz_1  | PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
baz_1  | FOO=foo
baz_1  | BAZ=baz
baz_1  | PWD=/
003_docker_compose_merge_bar_1 exited with code 0
003_docker_compose_merge_baz_1 exited with code 0
```

## Down

```bash
$ make down
Removing 003_docker_compose_merge_bar_1 ... done
Removing 003_docker_compose_merge_baz_1 ... done
Removing network 003_docker_compose_merge_default
```
