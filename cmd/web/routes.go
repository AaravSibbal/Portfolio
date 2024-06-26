package server

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) Routes() http.Handler {

	standardMiddleware := alice.New(app.logRequest, app.recoverPanic, app.secureHeaders) 
	dynamicMiddleware := alice.New(app.sessions.Enable)
	
	mux := pat.New()

	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/contact", dynamicMiddleware.ThenFunc(app.contactForm))
	mux.Post("/contact", dynamicMiddleware.ThenFunc(app.contact))
	mux.Get("/projects", dynamicMiddleware.ThenFunc(app.projects))
	mux.Get("/ping", dynamicMiddleware.ThenFunc(app.pong))

	fileServer := http.FileServer(http.Dir("ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
