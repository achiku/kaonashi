# kaonashi

[![Build Status](https://travis-ci.org/achiku/kaonashi.svg?branch=master)](https://travis-ci.org/achiku/kaonashi)


## Description

Headless local note app with just-enough set of RESTful APIs


## Why created

- There are variety of note apps out there, but the fundamental functionality is not so different among them. And Evernote sucks.
- It's getting easier and easier to create UI for multiple devices.
- Programmers and desingers are kind of people who can create whatever suitable for their own use. 
- It might be interesting to have solid note backend with simple RESTful APIs, and let people to create what is valuable for their use case.


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
