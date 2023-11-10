package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var c int = 0

type User struct {
	Nom    string
	Prenom string
	Anniv  string
	Sexe   string
}

var data = User{}

func main() {
	temp, err := template.ParseGlob("./temp/*.html")
	if err != nil {
		fmt.Println(fmt.Sprintf("ERREUR => %s", err.Error()))
		return
	}
	type Mentor struct {
		FirstName string
		LastName  string
		Age       int
		Sexe      bool
	}
	type PageVariable struct {
		Titre    string
		Nom      string
		Filiere  string
		Niveau   int
		Nombre   int
		Liste    string
		Etudiant []Mentor
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		Users := []Mentor{{"Cyril", "RODRIGUES", 22, true},
			{"Kheir-eddine", "MEDERREG", 22, false},
			{"Alan", "PHILIPIERT", 26, true}}
		dataPage := PageVariable{"Information de la promotion", "Mentor'ac", "informatique", 5, 3, "Liste des Etudiants", Users}
		temp.ExecuteTemplate(w, "promo", dataPage)
	})

	type Display struct {
		Cpt  int
		Trou bool
	}
	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		c++
		var laverite bool
		if c%2 == 0 {
			laverite = false
		} else {
			laverite = true
		}
		var dataPage = Display{c, laverite}
		temp.ExecuteTemplate(w, "change", dataPage)
	})

	http.HandleFunc("/user/init", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "init", nil)
	})

	http.HandleFunc("/user/treatment", func(w http.ResponseWriter, r *http.Request) {
		data = User{
			r.FormValue("nom"),
			r.FormValue("prenom"),
			r.FormValue("date"),
			r.FormValue("sexe")}
		fmt.Println(data)
		http.Redirect(w, r, "/user/display", http.StatusMovedPermanently)
	})

	http.HandleFunc("/user/display", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "display", data)
	})

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8080", nil)
}
