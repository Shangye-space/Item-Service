FROM golang:latest
WORKDIR $GOPATH/src/github.com/Shangye-space/Item-Service
COPY ./ .
RUN apt-get update
RUN go get "github.com/gorilla/mux"
RUN GOOS=linux GOARCH=386 go build -ldflags="-w -s" -v
RUN cp Item-Service /
