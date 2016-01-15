# kaonashi

[![Build Status](https://travis-ci.org/achiku/kaonashi.svg?branch=master)](https://travis-ci.org/achiku/kaonashi)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://travis-ci.org/achiku/kaonashi/LICENSE)


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

Interacting with kaonashi using [httpie](https://github.com/jkbrzt/httpie).

#### POST a note

```
http POST http://localhost:8080/note data:='{"title": "title01", "body": "hello, kaonashi!"}'
HTTP/1.1 200 OK
Content-Length: 31
Content-Type: application/json
Date: Sat, 09 Jan 2016 13:46:17 GMT

{
    "data": {
        "message": "created"
    }
}
```

#### PUT a note

```
http PUT http://localhost:8080/note/1 data:='{"title": "title01", "body": "hello, kaonashi! This is achiku."}'
HTTP/1.1 200 OK
Content-Length: 31
Content-Type: application/json
Date: Sat, 09 Jan 2016 13:48:23 GMT

{
    "data": {
        "message": "updated"
    }
}
```


#### DELETE a note

```
http DELETE http://localhost:8080/note/8
HTTP/1.1 200 OK
Content-Length: 31
Content-Type: application/json
Date: Sat, 09 Jan 2016 13:49:10 GMT

{
    "data": {
        "message": "deleted"
    }
}
```


#### GET note titles

```
http GET http://localhost:8080/note
HTTP/1.1 200 OK
Content-Length: 272
Content-Type: application/json
Date: Sat, 09 Jan 2016 13:51:12 GMT

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
http GET http://localhost:8080/note/2
HTTP/1.1 200 OK
Content-Length: 166
Content-Type: application/json
Date: Sat, 09 Jan 2016 13:51:53 GMT

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
