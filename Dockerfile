###
# Build Stage
###

FROM golang:1.24-alpine AS builder
WORKDIR /app

# Install dependencies
RUN apk add --no-cache build-base sqlite-dev

# Cache go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Build with CGO (required for SQLite)
RUN CGO_ENABLED=1 go build -o teamotp.bin ./cmd/teamotp

###
# Final Stage
###

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/teamotp.bin .
COPY --from=builder /app/web ./web
EXPOSE 6443
ENTRYPOINT ["./teamotp.bin"]
