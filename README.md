# kaonashi [![Build Status](https://travis-ci.org/achiku/kaonashi.svg?branch=master)](https://travis-ci.org/achiku/kaonashi) [![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://travis-ci.org/achiku/kaonashi/LICENSE)


## Description

Headless local note app with just-enough set of RESTful APIs


## Why created

There are variety of note apps out there, but the fundamental functionality is not so different among them. Programmers and desingers are kind of people who are good at creating whatever tools/UIs suitable for their own use case. This go gettable binary get your back, and you can create and keep polishing valuable note UIs for your own use case on top of RESTful APIs that should be enough for usual note app.


## Example UIs

- https://github.com/achiku/vim-kaonashi


## Instalation

```
go get -u github.com/achiku/kaonashi
```

## Create database

```
$ kaonashi -init
```

## Start service

```
$ nohup kaonashi -d &
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
