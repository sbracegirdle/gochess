package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Singe global game for now
var game *Game

func until(count int) (slice []int) {
	for i := 0; i < count; i++ {
		slice = append(slice, i)
	}
	return
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"until": until,
		"mod":   func(i, j int) int { return i % j },
		"add":   func(i, j int) int { return i + j },
		"sub":   func(i, j int) int { return i - j },
		"split": func(s string) []string { return strings.Split(s, "") },
	}).ParseFiles("chess.html"))
	err := tmpl.ExecuteTemplate(w, "body", game)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func boardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"until": until,
		"mod":   func(i, j int) int { return i % j },
		"add":   func(i, j int) int { return i + j },
		"sub":   func(i, j int) int { return i - j },
		"split": func(s string) []string { return strings.Split(s, "") },
	}).ParseFiles("chess.html"))
	err := tmpl.ExecuteTemplate(w, "board", game.Board) // Only return the board component

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// FormValue move
	move := r.FormValue("move")

	source, target, err := notationToCoordinates(move)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = game.MovePiece(source[0], source[1], target[0], target[1])

	if err != nil {
		// TODO respond with form error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func startServer() {
	// Create single global game for now
	game = NewGame("Player 1", "Player 2")

	r := mux.NewRouter()
	r.HandleFunc("/", gameHandler)
	r.HandleFunc("/move", moveHandler).Methods("POST")
	r.HandleFunc("/board", boardHandler).Methods("GET") // Add this line

	// TODO render history of moves
	// TODO handle multiple games
	http.ListenAndServe(":8080", r)
}
