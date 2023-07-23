package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"rpsweb/rps"
	"strconv"
)

type Player struct {
	Name string
}

var player Player

const (
	templateDir  = "templates/"
	templateBase = templateDir + "base.html"
)

func Index(w http.ResponseWriter, r *http.Request) {

	restartValues()
	renderPage(w, "index.html", nil)

}

func Game(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al parsear el formulario", http.StatusInternalServerError)
			log.Print(err)
			return
		}

		player.Name = r.Form.Get("name")
	}

	if player.Name == "" {
		http.Redirect(w, r, "/newgame", http.StatusFound)
	}

	renderPage(w, "game.html", player)
}

func Play(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func About(w http.ResponseWriter, r *http.Request) {
	restartValues()
	renderPage(w, "about.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	restartValues()
	renderPage(w, "newgame.html", nil)
}

func renderPage(w http.ResponseWriter, page string, data interface{}) {
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))
	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		log.Print(err)
		return
	}
}

func restartValues() {
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0
}
