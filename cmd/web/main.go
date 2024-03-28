package main

import (
	"crypto/tls"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golangcollege/sessions"
	"github.com/joho/godotenv"
)

type application struct {
	templateCache map[string]*template.Template
	errorLog      *log.Logger
	infoLog       *log.Logger
	sessions      *sessions.Session
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	secret := flag.String("secret", "-j-&qeotIeCgF&w_qJwOM^jYniD6J11K", "Secret")

	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errLogger := log.New(os.Stdout, "ERROR:\t", log.Ltime|log.Ldate|log.Lshortfile)

	err := godotenv.Load(".env")

	if err != nil {
		errLogger.Fatal("Error loading evvironment variable")
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true
	session.SameSite = http.SameSiteStrictMode

	newTemplateCache, err := newTemplateCache("ui/html")

	if err != nil {
		errLogger.Fatal(err)
	}

	app := application{
		sessions:      session,
		templateCache: newTemplateCache,
		infoLog:       infoLog,
		errorLog:      errLogger,
	}

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
		InsecureSkipVerify: true,
	}

	srv := &http.Server{
		Addr:         *addr,
		TLSConfig:    tlsConfig,
		ErrorLog:     errLogger,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,

		Handler: app.routes(),
	}

	err = srv.ListenAndServeTLS("tls/cert.pem", "tls/key.pem")

	log.Fatal(err)
}
