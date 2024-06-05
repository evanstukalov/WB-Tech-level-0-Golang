package api

import (
	"encoding/json"
	"github.com/evanstukalov/wildberries_internship_l0/internal/cache"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	Cache cache.Cache
}

func NewServer(cache cache.Cache) *Server {
	return &Server{
		Cache: cache,
	}
}

func (s *Server) initializeRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", s.healthCheckHandler).Methods("GET")
	r.HandleFunc("/orders/{id}", s.getOrderHandler).Methods("GET")
	return r
}

func (s Server) StartHTTPServer() {
	r := s.initializeRouter()
	log.Println("Starting HTTP server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("HTTP server failed to start: %v", err)
	}
}

func (s Server) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (s Server) getOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	order, err := s.Cache.Get(id)
	if err != nil {
		http.Error(w, "Failed to retrieve order", http.StatusInternalServerError)
		return
	}

	if order == nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	response, err := json.Marshal(order)
	if err != nil {
		http.Error(w, "Failed to marshal order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
