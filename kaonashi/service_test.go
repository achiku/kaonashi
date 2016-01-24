package kaonashi

import (
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/guregu/null"
)

func createTestData(db *DB) error {
	return nil
}

func setup() (*Kaonashi, error) {
	configPath := "./conf/config_test.toml"
	k, err := NewKaonashi(configPath)
	if err != nil {
		return k, err
	}
	k.InitDB()
	notes := []Note{
		Note{
			Title:   "title01",
			Body:    null.NewString("body01", true),
			Created: time.Now().String(),
			Updated: time.Now().String(),
		},
		Note{
			Title:   "title01",
			Body:    null.NewString("body01", true),
			Created: time.Now().String(),
			Updated: time.Now().String(),
		},
	}
	for _, note := range notes {
		_, err := k.db.NamedExec(`
		INSERT INTO note (title, body, created, updated) VALUES
		(:title, :body, :created, :updated)
		`, note)
		if err != nil {
			log.Fatalf("%s", err)
			return k, err
		}
	}
	return k, nil
}

func TestGetNoteTitles(t *testing.T) {
	k, err := setup()
	if err != nil {
		t.Fatalf("failed to init kaonashi: %s", err)
	}
	noteTitles, err := getNoteTitles(k.db)
	if err != nil {
		t.Errorf("failed: %s", err)
	}
	if len(noteTitles) != 2 {
		t.Errorf("expectedd 2 but got %d", len(noteTitles))
		t.Logf("%v", noteTitles)
	}
}

func TestGetNote(t *testing.T) {
	k, err := setup()
	if err != nil {
		t.Fatalf("failed to init kaonashi: %s", err)
	}
	note, err := getNote(k.db, "1")
	if err != nil {
		t.Errorf("failed: %s", err)
	}
	if note.Title != "title01" {
		t.Errorf("expected title01 but got %s", note.Title)
	}
}

func TestCreateNote(t *testing.T) {
	k, err := setup()
	if err != nil {
		t.Fatalf("failed to init kaonashi: %s", err)
	}

	note := Note{
		Title:   "new note",
		Body:    null.NewString("new body", true),
		Created: time.Now().String(),
		Updated: time.Now().String(),
	}
	if err := createNote(k.db, note); err != nil {
		t.Errorf("failed: %s", err)
	}
	noteTitles, _ := getNoteTitles(k.db)
	if len(noteTitles) != 3 {
		t.Errorf("expected 3 but got %d", len(noteTitles))
	}
}

func TestUpdateNote(t *testing.T) {
	k, err := setup()
	if err != nil {
		t.Fatalf("failed to init kaonashi: %s", err)
	}

	note := Note{
		ID:      1,
		Title:   "updated note",
		Body:    null.NewString("new body", true),
		Created: time.Now().String(),
		Updated: time.Now().String(),
	}

	if err = updateNote(k.db, note); err != nil {
		t.Errorf("failed: %s", err)
	}
	updatedNote, _ := getNote(k.db, strconv.Itoa(note.ID))
	if updatedNote.Title != "updated note" {
		t.Errorf("expected 'updated note' but got %s", updatedNote.Title)
	}

}

func TestDeleteNote(t *testing.T) {
	k, err := setup()
	if err != nil {
		t.Fatalf("failed to init kaonashi: %s", err)
	}

	if err := deleteNote(k.db, "1"); err != nil {
		t.Errorf("failed: %s", err)
	}
	noteTitles, _ := getNoteTitles(k.db)
	if len(noteTitles) != 1 {
		t.Errorf("expected 1 but got %d", len(noteTitles))
	}
}
