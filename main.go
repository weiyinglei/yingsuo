package main

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
)

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	//get classic martini
	m := martini.Classic()

	//middlewares
	//use gzip
	m.Use(gzip.All(gzip.Options{
		CompressionLevel: gzip.BestCompression,
	}))
	//use sessions
	store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(sessions.Sessions("my_session", store))
	//use render
	opt := render.Options{
		Directory:  "./view",
		Layout:     "layout",
		Extensions: []string{".html"},
	}
	m.Use(render.Renderer(opt))

	//router
	m.Get("/", func(r render.Render) {
		r.HTML(http.StatusOK, "index", nil)
	})
	m.Get("/user/:userid", func(r render.Render, p martini.Params) {
		var retData struct {
			Id string
		}
		retData.Id = p["userid"]
		// r.HTML(http.StatusOK, "user", retData)
		r.JSON(http.StatusOK, retData)
	})
	m.Post("/user/login", binding.Bind(User{}), func(r render.Render, user User, session sessions.Session) {
		if user.Username == "admin" {
			session.Set("username", user.Username)
			r.Redirect("/welcome", http.StatusOK)
		} else {
			r.HTML(http.StatusNotFound, "error", nil)
		}
	})
	m.Get("/welcome", func(r render.Render, session sessions.Session) {
		v := session.Get("username")
		var retData struct {
			Username string
		}
		retData.Username = v.(string)
		r.HTML(http.StatusOK, "welcome", retData)
	})

	//run on port 3000
	m.Run()
}
