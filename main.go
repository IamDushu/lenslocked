package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	bio := "<script>alert(`You're hacked`)</script>"
	w.Header().Set("content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my great site!</h1><p>"+ bio +"</p>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:hey.dushyanth@gmail.com\">hey.dushyanth@gmail.com</a></p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(
		w, 
		`
		<h1>FAQ Page</h1>
		<h2>Q: Is there a free version?</h2>
		<p style="font-size: 22px">A: Yes! We offer a free trail for 30 days on any paid plans.</p>
		<h2>Q: What are your support hours?</h2>
		<p style="font-size: 22px">A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends.</p>
		<h2>Q: How do I contact support?</h2>
		<p style="font-size: 22px">A: Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a></p>
		`,
	)
}

func paramHandler (w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>Your Id is %s</h1>`, param)
}

func main() {
	
	r := chi.NewRouter()

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/article/{id}", paramHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
  