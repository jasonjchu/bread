FROM golang:latest

RUN mkdir -p go/src/github.com/jasonjchu/bread
WORKDIR /go/src/github.com/jasonjchu/bread

COPY . /go/src/github.com/jasonjchu/bread

RUN apt-get update
RUN apt-get install go-dep
RUN dep ensure

EXPOSE 8080

CMD ["go", "run", "/go/src/github.com/jasonjchu/bread/cmd/bread/main.go"]
