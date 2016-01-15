VERSION=0.0.1

build:
	go build -o bin/kaonashi

gom:
	go get -u github.com/mattn/gom

bundle:
	gom install

fmt:
	go fmt ./...

check:
	gom test ./kaonashi

clean:
	rm -rf bin/*
