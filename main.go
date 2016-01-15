package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/achiku/kaonashi/kaonashi"
)

func main() {
	confPath := flag.String("c", "", "configuration file path for kaonashi")
	deamonFlag := flag.Bool("d", false, "whether or not to launch in the background(like a daemon)")
	initDatabase := flag.Bool("init", false, "initialize database schemas")
	flag.Parse()

	if *initDatabase {
		kaonashi.Init(*confPath)
		os.Exit(0)
	}
	if *deamonFlag {
		log.Println("starting deamon")
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
		if bytes.HasPrefix(s, []byte("starting kaonashi")) {
			fmt.Println(string(s))
			cmd.Process.Release()
		} else {
			log.Printf("unexpected response from kaonashi: `%s` error: `%v`\n", s, err)
			cmd.Process.Kill()
		}
	}
	kaonashi.Run(*confPath)
}
