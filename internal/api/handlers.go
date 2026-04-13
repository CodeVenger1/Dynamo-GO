package api

import (
	"fmt"
	"net/http"
)

// func PutHandler: Handles write requests
func PutHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	if key == "" || value == "" {
		http.Error(w, "missing key or value", http.StatusBadRequest)
		return
	}

	success := coord.Put(key, value)

	if !success {
		http.Error(w, "write failed", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Stored successfully\n")
}

// func GetHandler: Handles read requests
func GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	if key == "" {
		http.Error(w, "missing key", http.StatusBadRequest)
		return
	}

	val, ok := coord.Get(key)

	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, val)
}

// func DeleteHandler: Handles Delete requests
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}
	coord.Delete(key)
	fmt.Fprintf(w, "Deleted key: %s\n", key)
}
