# TeamOTP
## Quick Start
Create config file:
```toml
# app_data/config.toml

Port = 8081
LdapHost = "IP of domain controller"
LdapDomain = "mydomain.com"
```


```sh
docker-compose up --build -d
```

## Building
```sh
# Local binary
go build -o teamotp.bin ./cmd/teamotp

# Docker container
DOCKER_BUILDKIT=1 docker build -t teamotp .
```
