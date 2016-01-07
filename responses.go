package main

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
