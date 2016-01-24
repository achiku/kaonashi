package kaonashi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/rs/xhandler"
	"golang.org/x/net/context"
)

const (
	ctxKeyConfig = "config"
	ctxKeyDB     = "db"
)

// Kaonashi struct
type Kaonashi struct {
	config *AppConfig
	db     *DB
}

// NewKaonashi create tables
func NewKaonashi(confPath string) (*Kaonashi, error) {
	var (
		k         *Kaonashi
		appConfig *AppConfig
		err       error
	)
	if confPath != "" {
		appConfig, err = NewAppConfig(confPath)
		if err != nil {
			log.Fatalf("failed to load config: %s", err)
			return k, err
		}
	} else {
		appConfig = NewAppDefaultConfig()
	}
	db, err := NewDB(appConfig)
	if err != nil {
		log.Fatalf("failed to open database: %s", err)
		return k, err
	}
	k = &Kaonashi{
		config: appConfig,
		db:     db,
	}
	return k, nil
}

// InitDB create tables
func (k *Kaonashi) InitDB() {
	err := createTables(k.db)
	if err != nil {
		log.Fatalf("failed to create table: %s", err)
	}
}

// Run kaonashi
func (k *Kaonashi) Run() {
	// set up root context
	rootCtx := context.Background()
	rootCtx = context.WithValue(rootCtx, ctxKeyConfig, k.config)
	rootCtx = context.WithValue(rootCtx, ctxKeyDB, k.db)

	// middleware chaining
	c := xhandler.Chain{}
	c.Use(recoverMiddleware)
	c.Use(loggingMiddleware)
	c.UseC(xhandler.CloseHandler)

	// application routing
	mux := bone.New()
	mux.Get("/note", c.HandlerCtx(rootCtx, xhandler.HandlerFuncC(getNoteTitlesHandler)))
	mux.Get("/note/:id", c.HandlerCtx(rootCtx, xhandler.HandlerFuncC(getNoteHandler)))
	mux.Delete("/note/:id", c.HandlerCtx(rootCtx, xhandler.HandlerFuncC(deleteNoteHandler)))
	mux.Put("/note/:id", c.HandlerCtx(rootCtx, xhandler.HandlerFuncC(updateNoteHandler)))
	mux.Post("/note", c.HandlerCtx(rootCtx, xhandler.HandlerFuncC(createNoteHandler)))
	fmt.Printf("starting kaonashi using port: %s\n", k.config.ServerPort)
	if err := http.ListenAndServe(":"+k.config.ServerPort, mux); err != nil {
		log.Fatal(err)
	}
}
