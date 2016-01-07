package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-zoo/bone"

	"golang.org/x/net/context"
)

func getNote(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	noteID := bone.GetValue(r, "id")
	db := ctx.Value(ctxKeyDB).(*DB)

	note := Note{}
	err := db.Select(&note, `
	SELECT
	  id
	  ,title
	  ,body
	  ,updated
	  ,created
    FROM note
	WHERE id = $1
	`, noteID)
	if err != nil {
		log.Printf("%s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	res := NoteResponse{Data: note}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func getNoteTitles(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// db := ctx.Value("db").(*DB)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode("getNoteList")
}

func createNote(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// db := ctx.Value("db").(*DB)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode("createNote")
}

func deleteNote(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// noteId := bone.GetValue(r, "id")
	// db := ctx.Value("db").(*DB)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode("deleteNote")
}

func updateNote(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// noteId := bone.GetValue(r, "id")
	// db := ctx.Value("db").(*DB)
	// if note.Id == 0 {
	// 	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	// 	return
	// }
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode("updateNote")
}
