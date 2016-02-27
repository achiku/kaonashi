package kaonashi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/xhandler"
	"github.com/rs/xmux"
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
	mux := xmux.New()
	mux.GET("/note", xhandler.HandlerFuncC(getNoteTitlesHandler))
	mux.GET("/note/:id", xhandler.HandlerFuncC(getNoteHandler))
	mux.DELETE("/note/:id", xhandler.HandlerFuncC(deleteNoteHandler))
	mux.PUT("/note/:id", xhandler.HandlerFuncC(updateNoteHandler))
	mux.POST("/note", xhandler.HandlerFuncC(createNoteHandler))

	fmt.Printf("starting kaonashi using port: %s\n", k.config.ServerPort)
	if err := http.ListenAndServe(":"+k.config.ServerPort, c.HandlerCtx(rootCtx, mux)); err != nil {
		log.Fatal(err)
	}
}
