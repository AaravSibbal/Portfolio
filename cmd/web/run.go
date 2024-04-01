package server

import (
	"crypto/tls"
	"flag"
	"fmt"
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

var secret = flag.String("secret", "-j-&qeotIeCgF&w_qJwOM^jYniD6J11K", "Secret")

var session = makeSessionSecure(sessions.New([]byte(*secret))) 
var infoLog = log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
var errLogger = log.New(os.Stdout, "ERROR:\t", log.Ltime|log.Ldate|log.Lshortfile)

var app = application{
	sessions:      session,
	templateCache: getTemplateCache(),
	infoLog:       infoLog,
	errorLog:      errLogger,
}

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	addr := fmt.Sprintf("0.0.0.0:%s", port)

	err := godotenv.Load(".env")

	if err != nil {
		errLogger.Fatal("Error loading environment variable")
	}


	if err != nil {
		errLogger.Fatal(err)
	}

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
		InsecureSkipVerify:       true,
	}

	srv := &http.Server{
		Addr:         addr,
		TLSConfig:    tlsConfig,
		ErrorLog:     errLogger,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      app.Routes(),
	}

	err = srv.ListenAndServe()

	log.Fatal(err)
}

func getTemplateCache() map[string]*template.Template {
	templateCache, err := newTemplateCache("ui/html")
	if err != nil {
		log.Fatal("There was an error loading the template cache")
		return nil
	}
	return templateCache
}

func makeSessionSecure(session *sessions.Session) *sessions.Session{
	session.Lifetime = 12 * time.Hour
	session.Secure = true
	session.SameSite = http.SameSiteStrictMode

	return session
}
