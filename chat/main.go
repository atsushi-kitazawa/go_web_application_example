package main

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"sync"
)

type templateHander struct {
    once sync.Once
    filename string
    templ *template.Template
}

func (t *templateHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    t.once.Do(func() {
	t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
    })
    t.templ.Execute(w, nil)
}

func main() {
    //http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request) {
    //    w.Write([]byte(`
    //    <html>
    //        <head>
    //    	<title>chat</title>
    //        </head>
    //        <body>
    //    	let's play chat!!
    //        </body>
    //    </html>
    //    `))
    //})

    http.Handle("/", &templateHander{filename: "top.html"})

    if err := http.ListenAndServe(":8080", nil); err != nil {
	log.Fatal("ListenAndServe:", err)
    }
}
