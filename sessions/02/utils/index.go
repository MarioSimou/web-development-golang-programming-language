package utils

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"../models"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func ParseTemplates(root string) (*template.Template, error) {
	tpl := template.New("")

	cb := func(path string, info os.FileInfo, e error) error {
		if strings.Contains(path, ".gohtml") {
			if _, e := tpl.ParseFiles(path); e != nil {
				return e
			}
		}
		return e
	}

	if e := filepath.Walk(root, cb); e != nil {
		return tpl, e
	}

	return tpl, nil
}

func SetCookieMiddl(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("URL: ", r.URL.Path)
		if c, _ := r.Cookie("sid"); c == nil {
			if cookieSid, e := uuid.NewV4(); e == nil {
				sid := cookieSid.String()
				cookie := http.Cookie{
					Name:   "sid",
					Value:  sid,
					MaxAge: 3600,
				}

				// Sets the HTTP Cookie header in the request
				r.Header.Set("Cookie", cookie.String())
				// Sets the HTTP Set-Cookie header in the response
				http.SetCookie(w, &cookie)
			}
		}

		next(w, r, p)
	}
}

func CreateUser(r *http.Request) (models.User, bool) {
	var user models.User

	if userId, e := uuid.NewV4(); e == nil {
		id := userId.String()
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		hashedPassword, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if e != nil {
			return user, false
		}

		return models.User{
			Id:       id,
			Username: username,
			Email:    email,
			Password: hashedPassword,
		}, true
	}

	return user, false
}

func GetSessionId(r *http.Request) string {
	var cookie string
	if header := r.Header.Get("Cookie"); header != "" {
		cookie = strings.Split(strings.Split(header, ";")[1], "=")[1]
	}

	if header := r.Header.Get("Set-Cookie"); header != "" {
		cookie = strings.Split(strings.Split(header, ";")[0], "=")[1]
	}

	fmt.Println("COOKIE: ", cookie)
	return cookie
}

func GetUser(id string, storage map[string]models.User) models.User {
	if user, ok := storage[id]; ok {
		return user
	}
	return models.User{}
}
