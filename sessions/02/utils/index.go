package utils

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
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
		if c, _ := r.Cookie("sid"); c == nil {
			if cookieSid, e := uuid.NewV4(); e == nil {
				sid := cookieSid.String()
				cookie := http.Cookie{
					Name:   "sid",
					Value:  sid,
					MaxAge: 3600,
				}

				r.Header.Set("Set-Cookie", cookie.String())
				http.SetCookie(w, &cookie) // sets cookie to writable
			}
		}

		next(w, r, p)
	}
}
