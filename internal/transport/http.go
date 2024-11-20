package transport

import (
	"encoding/json"
	"example/go-blog-api/internal/todo"
	"log"
	"net/http"
)

type TodoItem struct {
	Item string `json:"item"`
}

type Server struct {
	mux *http.ServeMux
}

func NewServer(todoSvc *todo.Service) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {

		todoItems, err := todoSvc.GetAll()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(todoItems)

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

		err = todoSvc.Add(t.Item)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusCreated)
		return
	})

	mux.HandleFunc("GET /search", func(writer http.ResponseWriter, request *http.Request) {
		query := request.URL.Query().Get("q")

		if query == "" {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		results, err := todoSvc.Search(query)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		b, err := json.Marshal(results)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = writer.Write(b)
		if err != nil {
			log.Println(err)
			return
		}
	})

	return &Server{
		mux: mux,
	}

}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mux)
}
