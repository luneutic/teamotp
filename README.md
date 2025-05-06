# TeamOTP
## Quick Start
Create the config file
```sh
mkdir data
nano data/config.toml
```

with content:
```toml
# IP of domain controller
LdapHost = "169.254.42.10"

# Domain name
LdapDomain = "mydomain.com"
```

_Internal Use_: Create a self signed certificate.
Move them to `data/cert.pem` and `data/key.pem`.
```sh
mkcert <internal IP>
```

Build and run container.
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
