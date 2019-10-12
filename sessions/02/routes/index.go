package routes

import (
	"html/template"
	"net/http"

	"../models"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

func IndexGet(tpl *template.Template, sessions map[string]string, db map[string]models.User) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		tpl.ExecuteTemplate(w, "index.gohtml", nil)
	}
}

func RegisterGet(tpl *template.Template, sessions map[string]string, db map[string]models.User) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		tpl.ExecuteTemplate(w, "register.gohtml", nil)
	}
}

func createUser(r *http.Request) (models.User, bool) {
	var user models.User

	if userId, e := uuid.NewV4(); e == nil {
		id := userId.String()
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		return models.User{
			Id:       id,
			Username: username,
			Email:    email,
			Password: password,
		}, true
	}

	return user, false
}
func RegisterPost(tpl *template.Template, sessions map[string]string, db map[string]models.User) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		cookie, _ := r.Cookie("sid")

		if _, ok := sessions[cookie.Value]; ok {
			tpl.ExecuteTemplate(w, "index.gohtml", "No reason to register")
			return
		}

		if user, ok := createUser(r); ok {
			// Set-Cookie HTTP header
			http.SetCookie(w, cookie)
			sessions[cookie.Value] = user.Id
			db[user.Id] = user
			tpl.ExecuteTemplate(w, "index.gohtml", "Succesful registration")
			return
		}
	}
}
