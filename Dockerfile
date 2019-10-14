FROM arm32v7/golang:1.8
FROM golang:1.12-alpine AS builder

RUN apk add --no-cache git

RUN adduser -D -u 1000 appuser

WORKDIR $GOPATH/src/github.com/domgoodwin/go-api/

COPY ./go.mod ./
COPY ./go.sum ./
# Using go get.
RUN GO111MODULE=on go get -d -v

COPY . .
RUN GO111MODULE=on GOOS=linux GOARCH=386 go build -o /go/bin/go-api

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /go/bin/go-api /go/bin/go-api

USER 1000

EXPOSE 8080

LABEL version="0.9.1" maintainer="d0m182goodwin@gmail.com"

ENTRYPOINT ["/go/bin/go-api"]
