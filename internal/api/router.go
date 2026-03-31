package api

import "net/http"

// func SetupRoutes : It connects URLs -> handler
func SetupRoutes() {
	http.HandleFunc("/put", PutHandler)
	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/delete", DeleteHandler)
}
