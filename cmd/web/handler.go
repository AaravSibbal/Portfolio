package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/AaravSibbal/Portfolio/pkg/forms"
	"github.com/AaravSibbal/Portfolio/pkg/mailing"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.tmpl", nil)
}

func (app *application) contactForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "contact.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	fmt.Println("This is something")
	// if err != nil {
	// 	app.clientError(w, http.StatusBadRequest)
	// 	app.errorLog.Output(2, "something just happened")
	// }

	form := forms.New(r.PostForm)
	form.Required("name", "email", "phone", "discription")
	form.RequiredLength("phone", 10)
	form.MinLength("discription", 20)
	form.MaxLength("discription", 500)
	form.MatchesPattern("email", forms.EmailRX)

	if !form.Valid() {
		app.render(w, r, "contact.page.tmpl", &templateData{
			Form: form,
		})
	}

	err = mailing.Mail(r.FormValue("name"), r.FormValue("email"),
		r.FormValue("phone"), r.FormValue("discription"),
		os.Getenv("GMAIL"), os.Getenv("PASSWORD"))

	if err != nil {
		app.serverError(w, err)
	}

	app.sessions.Put(r, "flash", "Your message was sent successfully")

	http.Redirect(w, r, "/contact", http.StatusSeeOther)
}

func (app *application) projects(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "projects.page.tmpl", nil)
}

func (app *application) pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
