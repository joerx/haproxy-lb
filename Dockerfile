FROM golang:1.13 AS builder
WORKDIR /src
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /src/app /usr/local/bin
CMD ["/usr/local/bin/app"]
