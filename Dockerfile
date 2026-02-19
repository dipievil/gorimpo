FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG APP_VERSION=dev

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-X main.Version=${APP_VERSION}" \
    -o /bin/gorimpo ./cmd/gorimpo/main.go

FROM ubuntu:jammy

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /bin/gorimpo .

# future: playwright
# RUN ./gorimpo install-browsers

CMD ["./gorimpo"]