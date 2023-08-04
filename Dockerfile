FROM golang:1.19

RUN apt-get install -y make

ENV GOPATH /go
WORKDIR /work

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN make test namegon
