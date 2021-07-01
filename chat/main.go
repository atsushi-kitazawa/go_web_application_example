package main

import (
	"flag"
	"log"
	"net/http"
	_ "os"
	"path/filepath"
	"sync"
	"text/template"

	_ "github.com/atsushi-kitazawa/go_web_application_example/trace"
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
    t.templ.Execute(w, r)
}

func main() {
    var addr = flag.String("addr", ":8080", "address of application")
    flag.Parse()
    r := newRoom()
    //r.tracer = trace.New(os.Stdout)
    http.Handle("/chat", MustAuth(&templateHander{filename: "top.html"}))
    http.Handle("/login",&templateHander{filename: "login.html"})
    http.HandleFunc("/auth/",loginHander)
    http.Handle("/room", r)

    go r.run()

    log.Println("web server started. port:", *addr)
    if err := http.ListenAndServe(*addr, nil); err != nil {
	log.Fatal("ListenAndServe:", err)
    }
    log.Println("web server stopped.")
}
