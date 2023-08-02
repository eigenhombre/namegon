.PHONY: all install

all: namegon

namegon: main.go
	go build .

clean:
	rm -f namegon

test:
	go test -v

install:
	go install .

docker:
	docker build -t namegon .
