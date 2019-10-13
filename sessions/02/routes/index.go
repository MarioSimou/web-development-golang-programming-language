package routes

import (
	"html/template"
	"net/http"

	"../models"
	"../utils"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type Message struct {
	Type    string
	Content string
}

type Payload struct {
	Message Message
	Data    string
	User    models.User
}

func IndexGet(tpl *template.Template, sessions map[string]models.User, db map[string]models.User) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		user := utils.GetUser(utils.GetSessionId(r), sessions)
		tpl.ExecuteTemplate(w, "index.gohtml", Payload{User: user})
	}
}

func RegisterGet(tpl *template.Template, sessions map[string]models.User, db map[string]models.User) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		tpl.ExecuteTemplate(w, "register.gohtml", nil)
	}
}

func RegisterPost(tpl *template.Template, sessions map[string]models.User, db map[string]models.User) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		cookie, _ := r.Cookie("sid")

		// checks if the email exists in the database - 409 Conflict
		if user := utils.GetUser(r.FormValue("email"), db); user.Id != "" {
			http.Error(w, "The user has already been registered", http.StatusConflict)
			return
		}

		if user, ok := utils.CreateUser(r); ok {
			// stores the user in sessions
			sessions[cookie.Value] = user
			// stores the user in the database
			db[user.Email] = user

			tpl.ExecuteTemplate(w, "index.gohtml", Payload{
				Message: Message{
					Type:    "success",
					Content: "Welcome to our page. A user with the name " + user.Email + " has been correctly created.",
				},
				User: user,
			})
			return
		}
	}
}

func LoginGet(tpl *template.Template, sessions map[string]models.User, db map[string]models.User) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		tpl.ExecuteTemplate(w, "login.gohtml", nil)
	}
}

func LoginPost(tpl *template.Template, sessions map[string]models.User, db map[string]models.User) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		email := r.FormValue("email")
		password := r.FormValue("password")
		cookie, _ := r.Cookie("sid")

		if user := utils.GetUser(email, db); user.Id != "" {
			doesMatch := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
			if doesMatch == nil {
				db[user.Email] = user
				sessions[cookie.Value] = user

				tpl.ExecuteTemplate(w, "index.gohtml", Payload{
					Message: Message{
						Type:    "success",
						Content: "Successful sign in",
					},
					User: user,
				})
				return
			}

			http.Error(w, "The user password does not match", http.StatusUnauthorized)
			return
		}

		tpl.ExecuteTemplate(w, "register.gohtml", nil)
	}
}

func LogoutGet(tpl *template.Template, sessions map[string]models.User, db map[string]models.User) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if cookie, _ := r.Cookie("sid"); cookie != nil {
			// destroys the session
			if _, ok := sessions[cookie.Value]; ok {
				delete(sessions, cookie.Value)
			}
			// destroys cookie
			cookie.MaxAge = -1
			http.SetCookie(w, cookie)

			tpl.ExecuteTemplate(w, "index.gohtml", Payload{
				Message: Message{
					Type:    "success",
					Content: "Successful sign out",
				},
			})
			return
		}

		http.Error(w, "No session or cookie was identified during logout", http.StatusInternalServerError)
	}
}
