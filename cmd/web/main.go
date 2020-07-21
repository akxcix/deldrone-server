package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/iamadarshk/deldrone-server/pkg/models/mysql"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template
	sessionStore  *sessions.CookieStore
	customers     *mysql.CustomerModel
	vendors       *mysql.VendorModel
	listings      *mysql.ListingModel
}

func main() {
	// command line flags
	addr := flag.String("addr", ":443", "HTTP Address")
	tls := flag.Bool("tls", false, "Use TLS server")
	dsn := flag.String("dsn", "web:pass@/deldrone?parseTime=true", "Database DSN")
	authkey := flag.String("authkey", "super-secret-key", "Authentication key for sessions. Use 32 or 64 bytes.")
	encryptionkey := flag.String("encryptionkey", "super-secret-key", "Encryption key for sessions. Use 16, 24 or 32 bytes")
	flag.Parse()

	// different loggers to seperate informative logs and error logs
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	// sessions
	store := sessions.NewCookieStore(
		[]byte(*authkey),
		[]byte(*encryptionkey),
	)
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 6, // 12 hours
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	}

	// open connection to the database
	db, err := connectDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// template cache
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		log.Fatal(err)
	}

	// application dependencies
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: templateCache,
		sessionStore:  store,
		customers:     &mysql.CustomerModel{DB: db},
		vendors:       &mysql.VendorModel{DB: db},
		listings:      &mysql.ListingModel{DB: db},
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

func connectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping() // ping to check if connection is established
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
