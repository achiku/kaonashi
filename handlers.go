package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-zoo/bone"

	"golang.org/x/net/context"
)

func getNoteHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	noteID := bone.GetValue(r, "id")
	db := ctx.Value(ctxKeyDB).(*DB)
	note, err := getNote(db, noteID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("getNoteHandler: %s", err)
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			res := MessageResponse{Data: StatusMessage{Message: "not found"}}
			json.NewEncoder(w).Encode(res)
			return
		}
		log.Printf("getNoteHandler: %s", err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		res := MessageResponse{Data: StatusMessage{Message: "error"}}
		json.NewEncoder(w).Encode(res)
		return
	}
	w.Header().Set("Content-type", "application/json")
	res := NoteResponse{Data: note}
	json.NewEncoder(w).Encode(res)
	return
}

func getNoteTitlesHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	db := ctx.Value(ctxKeyDB).(*DB)
	noteTitles, err := getNoteTitles(db)
	if err != nil {
		log.Printf("getNoteTitlesHandler: %s", err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		res := MessageResponse{Data: StatusMessage{Message: "error"}}
		json.NewEncoder(w).Encode(res)
		return
	}
	w.Header().Set("Content-type", "application/json")
	res := NoteTitlesResponse{Data: noteTitles}
	json.NewEncoder(w).Encode(res)
	return
}

func createNoteHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var noteRequest NoteRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&noteRequest)
	if err != nil {
		log.Printf("createNoteHandler: %s", err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		res := MessageResponse{Data: StatusMessage{Message: "error"}}
		json.NewEncoder(w).Encode(res)
		return
	}

	note := Note{
		Title:   noteRequest.Data.Title,
		Body:    noteRequest.Data.Body,
		Created: time.Now().String(),
		Updated: time.Now().String(),
	}
	db := ctx.Value(ctxKeyDB).(*DB)
	err = createNote(db, note)
	if err != nil {
		log.Printf("createNoteHandler: %s", err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		res := MessageResponse{Data: StatusMessage{Message: "error"}}
		json.NewEncoder(w).Encode(res)
		return
	}
	w.Header().Set("Content-type", "application/json")
	res := MessageResponse{Data: StatusMessage{Message: "created"}}
	json.NewEncoder(w).Encode(res)
	return
}

func deleteNoteHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	noteID := bone.GetValue(r, "id")
	db := ctx.Value(ctxKeyDB).(*DB)
	err := deleteNote(db, noteID)
	if err != nil {
		log.Printf("deleteNoteHandler: %s", err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		res := MessageResponse{Data: StatusMessage{Message: "error"}}
		json.NewEncoder(w).Encode(res)
		return
	}
	w.Header().Set("Content-type", "application/json")
	res := MessageResponse{Data: StatusMessage{Message: "deleted"}}
	json.NewEncoder(w).Encode(res)
	return
}

func updateNoteHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var noteRequest NoteRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&noteRequest)
	if err != nil {
		log.Printf("updateNoteHandler: %s", err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		res := MessageResponse{Data: StatusMessage{Message: "error"}}
		json.NewEncoder(w).Encode(res)
		return
	}

	noteID, _ := strconv.Atoi(bone.GetValue(r, "id"))
	noteRequest.Data.ID = noteID
	db := ctx.Value(ctxKeyDB).(*DB)
	err = updateNote(db, noteRequest.Data)
	if err != nil {
		log.Printf("getNoteHandler: %s", err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		res := MessageResponse{Data: StatusMessage{Message: "error"}}
		json.NewEncoder(w).Encode(res)
		return
	}
	w.Header().Set("Content-type", "application/json")
	res := MessageResponse{Data: StatusMessage{Message: "updated"}}
	json.NewEncoder(w).Encode(res)
	return
}
