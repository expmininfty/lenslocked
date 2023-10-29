package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to this awsome site</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch...</p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>FAQ page</h1><p>faq page...</p>")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	http.ListenAndServe(":3000", r)
}
