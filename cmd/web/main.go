package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var logo = "" +
	"_____   _           _      _____ _                 _   _       \n" +
	"|  __ \\| |         | |    / ____| |               | | (_)      \n" +
	"| |  | | |     __ _| |__ | |    | | ___  _   _  __| |  _  ___  \n" +
	"| |  | | |    / _` | '_ \\| |    | |/ _ \\| | | |/ _` | | |/ _ \\ \n" +
	"| |__| | |___| (_| | |_) | |____| | (_) | |_| | (_| |_| | (_) |\n" +
	"|_____/|______\\__,_|_.__/ \\_____|_|\\___/ \\__,_|\\__,_(_)_|\\___/ \n"

func main() {
	fmt.Println(logo)

	addr := flag.String("addr", "127.0.0.1:4000", "HTTP Network Address")
	flag.Parse()

	iLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	eLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := application{
		errorLog: eLog,
		infoLog:  iLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: eLog,
		Handler:  app.routes(),
	}

	iLog.Printf("Starting at address %s", *addr)
	err := srv.ListenAndServe()
	eLog.Fatal(err)
}
