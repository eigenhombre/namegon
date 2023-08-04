.PHONY: all install

all: namegon

# Update namegon if any go file in any subdirectory changes:
namegon: $(shell find . -type f -name '*.go')
	go build -o namegon cmd/namegon/main.go

clean:
	rm -f namegon

test:
	go test -v

install:
	go install .

docker:
	docker build -t namegon .
