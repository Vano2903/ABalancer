FROM golang:1.17.6-alpine3.14

EXPOSE 8080
WORKDIR /go/src/admin
COPY go.mod go.sum /go/src/admin/
RUN go mod download

COPY *.html /go/src/admin/

COPY *.go /go/src/admin/
RUN go build *.go

CMD ./execs

EXPOSE 8080