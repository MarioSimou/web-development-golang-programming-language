package main

import (
	"html/template"
	"log"
	"net/http"

	"./models"
	"./routes"
	"./utils"
	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(utils.ParseTemplates("./views"))
}

func main() {
	sessions := make(map[string]string)
	db := make(map[string]models.User)

	router := httprouter.New()
	router.GET("/", utils.SetCookieMiddl(routes.IndexGet(tpl, sessions, db)))
	router.GET("/register", utils.SetCookieMiddl(routes.RegisterGet(tpl, sessions, db)))
	router.POST("/register", utils.SetCookieMiddl(routes.RegisterPost(tpl, sessions, db)))

	log.Fatal(http.ListenAndServe(":8080", router))
}
