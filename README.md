# kaonashi

[![Build Status](https://travis-ci.org/achiku/kaonashi.svg?branch=master)](https://travis-ci.org/achiku/kaonashi)


## Description

Headless note server with just-enough set of RESTful APIs


## Why created

I wanted a simple note app for programmers who can create whatever the UI they want.


## Instalation

```
go get -u github.com/achiku/kaonashi/cmd/kaonashi
```


## Start service

```
$ kaonashi -init
$ kaonashi -d &
```


## What it can do

- POST a note
- PUT a note
- DELETE a note
- GET note titles
- GET a note
- GET search notes
