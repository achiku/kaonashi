package kaonashi

import (
	"database/sql"
	"log"
	"time"
)

func getNote(db *DB, noteID string) (Note, error) {
	note := Note{}
	err := db.Get(&note, `
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
		if err == sql.ErrNoRows {
			log.Printf(
				"getNote: no rows round for id: %s", noteID)
			return note, err
		}
		log.Printf("getNote: %s", err)
		return note, err
	}
	return note, nil
}

func createNote(db *DB, note Note) error {
	_, err := db.NamedExec(`
	INSERT INTO note (title, body, created, updated)
	VALUES (:title, :body, :created, :updated)
	`, note)
	if err != nil {
		log.Printf("createNote: %s", err)
		return err
	}
	return nil
}

func getNoteTitles(db *DB) ([]NoteTitle, error) {
	noteTitles := []NoteTitle{}
	err := db.Select(&noteTitles, `
	SELECT
	  id
	  ,title
	  ,updated
	  ,created
    FROM note
	ORDER BY updated DESC
	`)
	if err != nil {
		log.Printf("getNoteTitles: %s", err)
		return noteTitles, err
	}
	return noteTitles, nil
}

func deleteNote(db *DB, noteID string) error {
	_, err := db.Exec(`
	DELETE FROM note WHERE id = ?
	`, noteID)
	if err != nil {
		log.Printf("deleteNote: %s", err)
		return err
	}
	return nil
}

func updateNote(db *DB, note Note) error {
	_, err := db.Exec(`
	UPDATE note 
	SET
	  title = ?
	  ,body = ?
	  ,updated = ?
	WHERE id = ?
	`, note.Title, note.Body, time.Now().String(), note.ID)
	if err != nil {
		log.Printf("updateNote: %s", err)
		return err
	}
	return nil
}

func createTables(db *DB) error {
	_, err := db.Exec(`
	CREATE TABLE note (
	  id INTEGER PRIMARY KEY AUTOINCREMENT
	  ,title TEXT
	  ,body TEXT
	  ,created TEXT NOT NULL
	  ,updated TEXT NOT NULL
	)
	`)
	if err != nil {
		log.Printf("createTables: %s", err)
		return err
	}
	return nil
}
