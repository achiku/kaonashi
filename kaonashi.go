package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/go-zoo/bone"
	"github.com/rs/xhandler"
	"golang.org/x/net/context"
)

const (
	ctxKeyConfig = "config"
	ctxKeyDB     = "db"
)

func run(confPath string) {
	var appConfig *AppConfig
	var err error
	if confPath != "" {
		appConfig, err = NewAppConfig(confPath)
		if err != nil {
			log.Fatalf("failed to load config: %s", err)
		}
	} else {
		appConfig = NewDefaultConfig()
	}
	db, err := NewDB(appConfig)
	if err != nil {
		log.Fatalf("failed to open database: %s", err)
	}

	// set up root context
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

func main() {
	confPath := flag.String("c", "", "configuration file path for kaonashi")
	d := flag.Bool("d", false, "whether or not to launch in the background(like a daemon)")
	flag.Parse()
	if *d {
		cmd := exec.Command(os.Args[0],
			"-c", *confPath,
		)
		serr, err := cmd.StderrPipe()
		if err != nil {
			log.Fatalln(err)
		}
		err = cmd.Start()
		if err != nil {
			log.Fatalln(err)
		}
		s, err := ioutil.ReadAll(serr)
		s = bytes.TrimSpace(s)
		if bytes.HasPrefix(s, []byte("addr: ")) {
			fmt.Println(string(s))
			cmd.Process.Release()
		} else {
			log.Printf("unexpected response from kaonashi: `%s` error: `%v`\n", s, err)
			cmd.Process.Kill()
		}
	}
	run(*confPath)
}
