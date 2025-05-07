# TeamOTP
## Quick Start
Create the config file
```sh
mkdir data
nano data/config.toml
```

with content:
```toml
LdapHost = "169.254.42.10"
LdapDomain = "mydomain.com"
```

_Internal Use_: Create a self signed certificate.
Move them to `data/cert.pem` and `data/key.pem`.
```sh
mkcert <internal IP>
```

Create and run `docker-compose.yml`:
```yml
services:
  teamotp:
    image: ghcr.io/luneutic/teamotp:0.1.0
    container_name: teamotp
    ports:
      - 6443:6443
    volumes:
      - ./data:/app/data
```

## Building
```sh
# Local binary
go build -o teamotp.bin ./cmd/teamotp

# Or run directly
go run ./cmd/teamotp

# Docker container
DOCKER_BUILDKIT=1 docker build -t teamotp .
```
