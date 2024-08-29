package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ravidhavlesha/go-multiproto-kvs/pkg/kvstore"
)

type HTTPServer struct {
	address string
	kvStore *kvstore.KVStore
}

// NewHTTPServer initializes a new HTTP server
func NewHTTPServer(address string, kvStore *kvstore.KVStore) *HTTPServer {
	return &HTTPServer{address: address, kvStore: kvStore}
}

// Start starts a HTTP server that listens to incoming connections.
func (server *HTTPServer) Start() error {
	http.HandleFunc("/get", server.handleGet)
	http.HandleFunc("/set", server.handleSet)
	http.HandleFunc("/delete", server.handleDelete)

	log.Printf("HTTP server started on %s", server.address)
	return http.ListenAndServe(server.address, nil)
}

// handleGet handles /get request
func (server *HTTPServer) handleGet(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing 'key' query parameter", http.StatusBadRequest)
		return
	}

	value, exists := server.kvStore.Get(key)
	if !exists {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, value)
}

// handleSet handles /set request
func (server *HTTPServer) handleSet(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	if key == "" || value == "" {
		http.Error(w, "Missing 'key' or 'value' query parameter", http.StatusBadRequest)
		return
	}

	err := server.kvStore.Set(key, value)
	if err != nil {
		http.Error(w, "Error setting value", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

// handleDelete handles /delete request
func (server *HTTPServer) handleDelete(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing 'key' query parameter", http.StatusBadRequest)
		return
	}

	err := server.kvStore.Delete(key)
	if err != nil {
		http.Error(w, "Error deleting value", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
