package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	_ "os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/atsushi-kitazawa/go_web_application_example/credential"
	"github.com/atsushi-kitazawa/go_web_application_example/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
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

    // load credentail
    cred := credential.Load()
    t := trace.New(os.Stdout)
    t.Trace(cred.SecurityKey)
    t.Trace(cred.ClientId)
    t.Trace(cred.Secret)
    gomniauth.SetSecurityKey(cred.SecurityKey)
    gomniauth.WithProviders(
      google.New(cred.ClientId, cred.Secret, "http://localhost:8080/auth/callback/google"),
    )

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
