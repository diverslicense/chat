package main

import (
    "fmt"
    "log"
    "net/http"
    "net/url"
    "time"

    //"github.com/gorilla/mux"
)

type Status struct {
    Status string `json:"status"`
}

type Username struct {
    Username string `json:"username"`
}

func main() {

    cookie := http.Cookie{}

    // func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) 
    // HandleFunc registers the handler function for the given pattern in the DefaultServeMux 
    // A ResponseWriter interface is used by an HTTP handler to construct an HTTP response
    // A Request (struct) represents an HTTP request received by a server or to be sent by a client
    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        param, _ := url.ParseQuery(r.URL.RawQuery)
        username := param["username"][0]

        cookie = http.Cookie{
            Name: "username", 
            Value: username, 
            Path: "/",
            Domain: "localhost",
            Expires: time.Now().AddDate(0, 0, 2),
        }
        http.SetCookie(w, &cookie)
        s := Status{}
        if cookie.Value != "" {
            s.Status = "success"
        } else {
            s.Status = "error: user not logged in"
        }
        fmt.Fprintf(w, s.Status + " " + cookie.Value)
        //fmt.Println("inside of /login handler function, name: " + cookie.Name + "; value: " + cookie.Value + "; expires: " + cookie.Expires.String())
    })

    //fmt.Println("outside in main, name: " + cookie.Name + "; value: " + cookie.Value + "; expires: " + cookie.Expires.String())

    http.HandleFunc("/whatsmyname", func(w http.ResponseWriter, r *http.Request) {
        c, err := r.Cookie("username")
        if err != nil {
            log.Print(err)
        }

        fmt.Println(c.Value)
        //fmt.Println("inside of /whatsmyname handler function, name: " + cookie.Name + "; value: " + cookie.Value + "; expires: " + cookie.Expires.String())
    })

    http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
        cooki, err := r.Cookie("username")
        if err != nil {
            log.Print(err)
        }
        cooki.Path = "/login"
        cooki.Domain = "localhost"
        cooki.MaxAge = -100
        // fmt.Println("inside of /logout handler function, name: " + cookie.Name + "; value: " + cookie.Value + "; expires: " + cookie.Expires.String())
        // fmt.Println("inside of /logout handler function, name: " + cooki.Name + "; value: " + cooki.Value + "; expires: " + cooki.Expires.String())
        fmt.Fprint(w, cookie)
    })

    // func ListenAndServe(addr string, handler Handler) error
    // ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle
    // requests on incoming connections.
    // ListenAndServe starts an HTTP server with a given address and handler.
    // The handler is usually nil, which means to use DefaultServeMux.
    // Handle and HandleFunc add handlers to DefaultServeMux.
    log.Fatal(http.ListenAndServe(":8080", nil))
}