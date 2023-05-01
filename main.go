package main

import (
	"net/http"
)

func AdminHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Admin page"))
	}
}

func RequestMethodGetMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
			return
		}

		// jika method adalah get
		next.ServeHTTP(w, r)
	}) 
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("role") != "ADMIN" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Role not authorized"))
			return
		} 

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/admin", AdminHandler())

	handler := RequestMethodGetMiddleware(AdminMiddleware(mux))

	http.ListenAndServe("localhost:8080", handler)
}
