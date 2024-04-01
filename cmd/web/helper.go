package server

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"sync"
)

func (app *application) addDefaultData(r *http.Request, td *templateData) *templateData {

	if td == nil {
		td = &templateData{}
	}

	td.Flash = app.sessions.PopString(r, "flash")

	return td
}

func (app *application) render(w http.ResponseWriter, r *http.Request,
	name string, td *templateData) {

	mutex := sync.RWMutex{}
	mutex.Lock()
	ts, ok := app.templateCache[name]
	mutex.Unlock()

	if !ok {
		app.serverError(w, fmt.Errorf("the template %s does not exist", name))
	}

	buff := new(bytes.Buffer)

	err := ts.Execute(buff, app.addDefaultData(r, td))

	if err != nil {
		app.serverError(w, err)
		return
	}

	buff.WriteTo(w)
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s \n\n %s", err.Error(), debug.Stack())

	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

