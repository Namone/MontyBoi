package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type MontyBoi struct {
	Location   string
	Status     string
	ServerName string
}

var (
	location string
	status   string
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	// Don't forget to omit the trailing /!
	serverName := ":" + port
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		data := MontyBoi{
			Location:   location,
			Status:     status,
			ServerName: serverName,
		}

		tmpl.Execute(w, data)
	})

	http.HandleFunc("/update/location", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
			postValue, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}

			location = string(postValue)
			fmt.Printf("Updated Monty's location to: %s\n", postValue)
		}

	})

	http.HandleFunc("/update/status", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
			postValue, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}

			status = string(postValue)
			fmt.Printf("Updated Monty's status to: %s\n", status)
		}
	})

	http.HandleFunc("/src/css/primary/montyboi.css", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]
		data, err := ioutil.ReadFile(string(path))

		fmt.Print("Loading " + path + "\n")
		if err == nil {
			w.Header().Add("Content-Type", "text/css")
			w.Write(data)
		} else {
			w.WriteHeader(404)
			w.Write([]byte("404 My Friend - " + http.StatusText(404)))
		}
	})

	http.HandleFunc("/src/js/montyboi.js", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]
		data, err := ioutil.ReadFile(string(path))

		fmt.Print("Loading " + path + "\n")
		if err == nil {
			w.Header().Add("Content-Type", "text/javascript")
			w.Write(data)
		} else {
			w.WriteHeader(404)
			w.Write([]byte("404 My Friend - " + http.StatusText(404)))
		}
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
