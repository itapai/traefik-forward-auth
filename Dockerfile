FROM golang:alpine as builder
RUN mkdir /build 
COPY . /build
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o traefik-forward-auth cmd/server/main.go

# Build the smallest image
FROM scratch
COPY --from=builder /build/traefik-forward-auth ./
ENTRYPOINT ["./traefik-forward-auth"]


