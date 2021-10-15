package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/", handleRedirect)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		log.Println("request must have url parameter.")
	}
	handleCookie(w, r)
	printRequest(r)
	http.Redirect(w, r, url, 302)
}

func handleCookie(w http.ResponseWriter, r *http.Request) {
	uuidCookie, err := r.Cookie("uuid")
	if err != nil {
		log.Printf("cookie error: %v\n", err)
	}
	if uuidCookie == nil || uuidCookie.Value == "" {
		if uuidCookie == nil {
			uuidCookie = &http.Cookie{
				Name: "uuid",
				// TODO: expire
			}
		}
		uuidCookie.Value = uuid.NewString()
	}
	http.SetCookie(w, uuidCookie)
}

func printRequest(r *http.Request) {
	printHeaders(r)
	printCookies(r)
	printQueries(r)
}

func printHeaders(r *http.Request) {
	fmt.Println("Headers: start")
	for k, v := range r.Header {
		fmt.Printf("\t%s: %s\n", k, v)
	}
	fmt.Println("Headers: end")
}

func printCookies(r *http.Request) {
	fmt.Println(("Cookeis: start"))
	cookies := r.Cookies()
	if cookies != nil {
		for _, v := range cookies {
			fmt.Println("\t" + v.String())
		}
	}
	fmt.Println(("Cookeis: end"))
}

func printQueries(r *http.Request) {
	fmt.Println(("Queries: start"))
	queries := r.URL.Query()
	if queries != nil {
		for k, v := range r.URL.Query() {
			fmt.Printf("\t%s: %s\n", k, v)
		}
	}
	fmt.Println(("Queries: end"))
}
