FROM golang:1.15 as builder
WORKDIR /go/src/github.com/takahiro-hayakawa/user-api-server
ENV CGO_ENABLED=0 \
GOOS=linux \
GOARCH=amd64
COPY . .
RUN go build

FROM alpine:3.12 as runtime
WORKDIR /usr/local/bin
COPY --from=builder /go/src/github.com/takahiro-hayakawa/user-api-server/user-api-server /usr/local/bin/user-api-server
CMD ["./user-api-server"]
