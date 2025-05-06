# TeamOTP
## Quick Start
Create config file
```sh
mkdir data
nano data/config.toml
```

Add the following content
```toml
Port = 8081
LdapHost = "IP of domain controller"
LdapDomain = "mydomain.com"
```

Build and run container
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
