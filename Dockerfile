FROM golang:1.7.4

RUN mkdir /app

ADD . /go/src/goblin
WORKDIR /go/src/goblin

RUN go install .

ENTRYPOINT /go/bin/goblin

EXPOSE 8001
