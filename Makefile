VERSION=0.0.1

gom:
	go get -u github.com/mattn/gom

bundle:
	gom install

fmt:
	go fmt ./...

check:
	gom test

clean:
	rm -rf bin/*
