package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// command line flags
	addr := flag.String("addr", ":443", "HTTP Address")
	tls := flag.Bool("tls", false, "Use TLS server")
	dsn := flag.String("dsn", "web:pass@/deldrone?parseTime=true", "Database DSN")
	flag.Parse()

	// open connection to the database
	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping() // ping to check if connection is established
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
	if *tls {
		err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	} else {
		err = srv.ListenAndServe()
	}
	log.Fatal(err)

}
