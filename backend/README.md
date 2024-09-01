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

### Open API docs
To generate the docs by service format:
```shell
  swag init -o cmd/{service_name}/docs -d ./cmd/{service_name}/gin,./controller/{service_name}

  # auth: swag init -o cmd/auth/docs -d ./cmd/auth/gin ./controller/auth
  # user: swag init -o cmd/user/docs -d ./cmd/user/gin ./controller/user
```

The endpoint format:
```shell
  http://localhost:{service_port}/swagger/index.html

  # auth: http://localhost:10001/swagger/index.html
  # user: http://localhost:10002/swagger/index.html
  # ... etc
```


# Todo
- [X] Integrate gin-swagger 
- [ ] Implement deployment by Terraform
- [ ] Build common commands by makefile for services 