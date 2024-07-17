# Prerequisite
- Go wire CLI ([Link](https://github.com/google/wire?tab=readme-ov-file))
- Docker ([Link](https://docs.docker.com/desktop/install/mac-install/#system-requirements))
  

# Getting started
### Generate ORM scripts
```shell
  cd ./backend && go generate ./ent
``` 

### Generate DI script
>  We use go-wire generation tool to generate DI scripts

Command format:

```shell
  wire gen ./cmd/${service}/${framework}

  # service: auth/user/......etc
  # framework: gin/lambda/......etc
```

example:
```shell
  wire gen ./cmd/auth/gin
```

### Development
```shell
docker compose -f ./docker/docker-compose.yml up -d
```
> Currently, we use `CompileDaemon` for hot reload, but if there are changes relate to ORM/DI, we should manually generate the scripts.


# Todo
- [x] User Provider Table
- [ ] Primary JWT token and refresh token
- [ ] Facebook provider
- [ ] Google provider