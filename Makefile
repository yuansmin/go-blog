BINARY=hello.out

all: build

build:
	go build -o $(BINARY) -v

run:
	go build -o $(BINARY) -v
	./$(BINARY)

