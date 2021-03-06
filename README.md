# kaonashi 

[![Build Status](https://travis-ci.org/achiku/kaonashi.svg?branch=master)](https://travis-ci.org/achiku/kaonashi)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://travis-ci.org/achiku/kaonashi/LICENSE)

Headless local note app with just-enough set of RESTful APIs


## Why created

There are wide range of note apps out there, but the fundamental functionality is not so different among them. Programmers and desingers are kind of people who are good at creating whatever tools/UIs suitable for their own use case, and it has became substantially easier to create UI for multiple devices as variety of tools evolved in the past few years. This go gettable binary, called kaonashi, will get your back. You can create and keep polishing valuable note UIs for your own use case on top of RESTful APIs that should be just-enough for simple note app.


## Example UIs

- https://github.com/achiku/vim-kaonashi
- https://github.com/ideyuta/kao


## Instalation

```
go get -u github.com/achiku/kaonashi/cmd/kaonashi
```

## Create database

```
$ kaonashi -init
```

## Start service

```
$ kaonashi -d &
```


## What it can do

Interacting with kaonashi using curl.

#### POST a note

```
curl -v -H "Accept: application/json" \
    -H "Content-type: application/json" \
    -X POST \
    -d '{"data":{"title":"my first note","body":"note body"}}' \
    http://localhost:8080/note
```

```json
{
    "data": {
        "message": "created"
    }
}
```

#### PUT a note

```
curl -v -H "Accept: application/json" \
    -H "Content-type: application/json" \
    -X PUT \
    -d '{"data":{"title":"my first note","body":"note body updated"}}' \
    http://localhost:8080/note/1
```

```json
{
    "data": {
        "message": "updated"
    }
}
```


#### DELETE a note

```
curl -v -H "Accept: application/json" \
    -H "Content-type: application/json" \
    -X DELETE \
    http://localhost:8080/note/1
```

```json
{
    "data": {
        "message": "deleted"
    }
}
```


#### GET note titles

```
curl -v -H "Accept: application/json" \
    -H "Content-type: application/json" \
    -X GET \
    http://localhost:8080/note
```

```json
{
    "data": [
        {
            "created": "2016-01-09 22:51:10.860726046 +0900 JST",
            "id": 2,
            "title": "title 02",
            "updated": "2016-01-09 22:51:10.861432579 +0900 JST"
        },
        {
            "created": "2016-01-09 22:20:45.814969118 +0900 JST",
            "id": 1,
            "title": "title 01",
            "updated": "2016-01-09 22:20:45.816734399 +0900 JST"
        }
    ]
}
```

#### GET a note

```
curl -v -H "Accept: application/json" \
    -H "Content-type: application/json" \
    -X GET \
    http://localhost:8080/note/2
```

```json
{
    "data": {
        "body": "hello, kaonashi!",
        "created": "2016-01-09 22:51:10.860726046 +0900 JST",
        "id": 2,
        "title": "title01",
        "updated": "2016-01-09 22:51:10.861432579 +0900 JST"
    }
}

```

## Contribute

Pull requests for new features, bug fixes, and suggestions are welcome!

### Install gom

This project is using [gom](https://github.com/mattn/gom) for dependency management.

```
$ go get -u github.com/mattn/gom
$ gom install
```

### Test

```
$ go test -v
```


### Write code and auto-reload

Using [wbs](https://github.com/achiku/wbs), it will be really easy to code and reload.

```
$ cd cmd/kaonashi
$ gom exec wbs -c wbs.toml
```
