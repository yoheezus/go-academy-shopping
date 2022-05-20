FROM golang:alpine3.15 AS builder

RUN apk update
RUN apk add curl
RUN mkdir -p /go/src/github.com/go-academy-shopping
RUN adduser -D -u 1000 appuser

WORKDIR /go/src/github.com/go-academy-shopping

COPY . .
# GOOS=linux and GOARCH=amd64 required for it to build
RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/academy-api

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/academy-api /go/bin/academy-api

EXPOSE 8080

ENTRYPOINT [ "/go/bin/academy-api" ] 

