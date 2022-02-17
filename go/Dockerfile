FROM golang:1.17.6-alpine3.14

# ENV APP_NAME benchmarks
# ENV PORT 8080
ENV secret ciao
#${SECRET}
#${APP_NAME}

EXPOSE 8080
WORKDIR /go/src/benchmarks
COPY go.mod go.sum /go/src/benchmarks/
RUN go mod download

COPY *.go /go/src/benchmarks/
RUN go build main.go

CMD ./main

EXPOSE 8080