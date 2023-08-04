.PHONY: all install

all: test

# Update namegon if any go file in any subdirectory changes:
namegon: *.go cmd/namegon/*.go
	go build -o namegon cmd/namegon/main.go

clean:
	rm -f namegon

test:
	go test -v

install:
	go install .

docker:
	docker build -t namegon .
