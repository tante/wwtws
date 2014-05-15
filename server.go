package main

import (
    "bufio"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "math/rand"
    "net/http"
    "os"
    "time"
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
        quote := RandomQuote()

        type QuoteInformation struct {
            Quote string
            Tech  string
        }

        info := QuoteInformation{quote, req.PostForm.Get("name")}

        r.HTML(200, "greeting", info)
    })

    m.Run()
}

func Filereader(Filename string) []string {
    file, _ := os.Open(Filename)
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines
}

func RandomQuote() string {
    quotes := Filereader("quotes.txt")
    // without seeding rand kinda sucks.
    rand.Seed(time.Now().Unix())
    return quotes[rand.Intn(len(quotes))]
}
