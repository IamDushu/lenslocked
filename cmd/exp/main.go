package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//Exercises
// Ex.1 Add new fields,
// Ex.2 Experiment diff data types
// Ex.3 Add nested data and access
// Ex.4 Try to add if/else inside template

type User struct {
	Name       string
	Bio        string
	Age        int
	Experience float64
	Education  map[string]string
	Likes      []string
	Friends    []User
}

func main() {

	r := chi.NewRouter()

	r.Get("/template", func(w http.ResponseWriter, r *http.Request) {

		t, err := template.ParseFiles("hello.gohtml")
		if err != nil {
			panic(err)
		}

		myLikes := []string{}
		myLikes = append(myLikes, "PaniPuri", "Perugu")

		user1 := User{
			Name: "Chaithanya",
		}

		user2 := User{
			Name: "Dushyanth",
		}

		user := User{
			Name:       "Dhana",
			Bio:        "<script>alert(`You're hacked`)</script>",
			Age:        122,
			Experience: 1.2,
			Education: map[string]string{
				"10":     "Vidyanikethan",
				"12":     "Chaithanya-Kazana",
				"Degree": "SVU",
			},
			Likes:   myLikes[:1],
			Friends: []User{user1, user2},
		}

		err = t.Execute(w, user)
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {

		t, err := template.ParseFiles("hello.gohtml")
		if err != nil {
			panic(err)
		}

		myLikes := []string{}
		myLikes = append(myLikes, "PaniPuri", "Perugu")

		user1 := User{
			Name: "Chaithanya",
		}

		user2 := User{
			Name: "Dushyanth",
		}

		user := User{
			// Name:       "Dhana",
			Bio:        "<script>alert(`You're hacked`)</script>",
			Age:        122,
			Experience: 1.2,
			Education: map[string]string{
				"10":     "Vidyanikethan",
				"12":     "Chaithanya-Kazana",
				"Degree": "SVU",
			},
			Likes:   myLikes[:1],
			Friends: []User{user1, user2},
		}

		err = t.Execute(w, user)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":3000", nil) //instead of nil if you use 'r' then chi get's used
}
