package routes

import (
	"fmt"
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

func createCookie(r *http.Request) (*http.Cookie, bool) {
	var cookie http.Cookie
	if cookieUUID, e := uuid.NewV4(); e == nil {
		sid := cookieUUID.String()

		cookie = http.Cookie{
			Name:   "sid",
			Value:  sid,
			MaxAge: 3600,
		}

		return &cookie, true
	}

	return &cookie, false
}

func RegisterPost(tpl *template.Template, sessions map[string]string, db map[string]models.User) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		c, _ := r.Cookie("sid")
		fmt.Println("C: ", c)
		if c != nil {
			if _, ok := sessions[c.Name]; ok {
				http.Redirect(w, r, "/", http.StatusPermanentRedirect)
				return
			}
			// if a cookie exists delete it and recreate another one
			c.MaxAge = -1
			http.SetCookie(w, c)
		}

		if cookie, ok := createCookie(r); ok {
			if user, ok := createUser(r); ok {
				// Set-Cookie HTTP header
				http.SetCookie(w, cookie)

				sessions[cookie.Value] = user.Id
				db[user.Id] = user

				fmt.Println(db)

				tpl.ExecuteTemplate(w, "index.gohtml", "Succesful registration")
				return
			}
		}
	}
}
