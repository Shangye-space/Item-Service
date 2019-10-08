FROM golang:1.11 as builder
WORKDIR $GOPATH/src/github.com/Shangye-space/Item-Service
COPY ./ .
RUN GOOS=linux GOARCH=386 go build -ldflags="-w -s" -v
RUN cp Item-Service /

FROM alpine:latest
COPY --from=builder /Item-Service /
CMD ["/Item-Service"]