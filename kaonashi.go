package main

import (
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

func main() {
	// set up root context
	appConfig, err := NewAppConfig("./conf/config_devel.toml")
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}
	db, err := NewDB(appConfig)
	if err != nil {
		log.Fatalf("failed to open database: %s", err)
	}
	rootCtx := context.Background()
	rootCtx = context.WithValue(rootCtx, ctxKeyConfig, appConfig)
	rootCtx = context.WithValue(rootCtx, ctxKeyDB, db)

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
	if err := http.ListenAndServe(":"+appConfig.ServerPort, mux); err != nil {
		log.Fatal(err)
	}
}
