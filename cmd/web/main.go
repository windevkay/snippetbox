package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog	*log.Logger
	infoLog		*log.Logger
}

func main() {
	// runtime flags
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	// levelled logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// app instance
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	// override http defaults e.g. ErrorLog
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}