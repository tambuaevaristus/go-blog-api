package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type TodoItem struct {
	Item string `json:"item"`
}

func main() {

	var todos = make([]string, 0)
	 mux :=http.NewServeMux()

	 mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {
		b,err := json.Marshal(todos)

		if err != nil {
			log.Println(err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
      
	 })

	 mux.HandleFunc("POST /todo", func(w http.ResponseWriter, r *http.Request) {

		var t TodoItem
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}


		todos = append(todos, t.Item)
		w.WriteHeader(http.StatusCreated)
		return
	 })
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}