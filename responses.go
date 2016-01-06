package main

type NoteResponse struct {
	Data Note `json:"data"`
}

type NoteTitlesResponse struct {
	Data []NoteTitle `json:"data"`
}

type NotesResponse struct {
	Data []Note `json:"data"`
}

type StatusMessage struct {
	Message string `json:"message"`
}

type MessageResponse struct {
	Data StatusMessage `json:"data"`
}
