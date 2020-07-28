FROM golang:1.14.6-buster AS builder

WORKDIR /go/src/bench-api

COPY ./ ./

RUN ls -la cmd

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/benchapi -ldflags="-s -w" cmd/benchapi/main.go

FROM scratch

COPY --from=builder /go/bin/benchapi /benchapi

EXPOSE 8080

CMD ["/benchapi"]