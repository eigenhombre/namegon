FROM golang:1.19

RUN apt-get install -y make

WORKDIR /work

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make test
