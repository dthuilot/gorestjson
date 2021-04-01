hello:
	echo "Hello"

build:
	go build -o bin/main cmd/*

run:
	go run cmd/*