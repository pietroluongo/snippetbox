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
	// Parse flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Setup Loggers & handlers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// Run
	infoLog.Printf("Starting server on port %s", *addr)
	err := server.ListenAndServe()
	errorLog.Fatal(err)

}
