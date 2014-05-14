package main

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "net/http"
)

func main() {
    m := martini.Classic()
    // use render for templating
    m.Use(render.Renderer())

    // index
    m.Get("/", func(r render.Render) {

        r.HTML(200, "hello", "")
    })

    m.Post("/", func(r render.Render, req *http.Request) {
        req.ParseForm()
        r.HTML(200, "greeting", req.PostForm.Get("name"))
    })

    m.Run()
}
