package main

import "time"

// Note struct
type Note struct {
	ID      int64     `json:"id"`
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// NoteTitle struct
type NoteTitle struct {
	ID      int64     `json:"id"`
	Title   string    `json:"title"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
