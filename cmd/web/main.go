package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// command line flags
	addr := flag.String("addr", ":4000", "HTTP Address")
	flag.Parse()

	// different loggers to seperate informative logs and error logs
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	// application dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// server settings
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem") // TODO: Shift to HTTPS server
	log.Fatal(err)

}
