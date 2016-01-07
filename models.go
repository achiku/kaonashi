package main

import "github.com/guregu/null"

// Note struct
type Note struct {
	ID      int64       `json:"id"`
	Title   string      `json:"title"`
	Body    null.String `json:"body"`
	Created string      `json:"created"`
	Updated string      `json:"updated"`
}

// NoteTitle struct
type NoteTitle struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

// NoteRequest struct
type NoteRequest struct {
	Data Note `json:"data"`
}

// NoteResponse struct
type NoteResponse struct {
	Data Note `json:"data"`
}

// NoteTitlesResponse struct
type NoteTitlesResponse struct {
	Data []NoteTitle `json:"data"`
}

// NotesResponse struct
type NotesResponse struct {
	Data []Note `json:"data"`
}

// StatusMessage struct
type StatusMessage struct {
	Message string `json:"message"`
}

// MessageResponse struct
type MessageResponse struct {
	Data StatusMessage `json:"data"`
}
