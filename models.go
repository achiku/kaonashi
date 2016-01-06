package main

import "time"

type Note struct {
	Id      int64     `json:"id" db:id`
	Title   string    `json:"title" db:title`
	Body    string    `json:"body" db:body`
	Created time.Time `json:"created" db:created`
	Updated time.Time `json:"updated" db:updated`
}

type NoteTitle struct {
	Id      int64     `json:"id"`
	Title   string    `json:"title"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
