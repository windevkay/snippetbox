package main

import (
	"fmt"
	"net/http"
)

// Set OWASP Security Headers on responses
func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin") 
		w.Header().Set("X-Content-Type-Options", "nosniff") 
		w.Header().Set("X-Frame-Options", "deny") 
		w.Header().Set("X-XSS-Protection", "0")
		
		next.ServeHTTP(w, r) 
	})
}

// Log information on incoming requests
func (app *application) logRequest(next http.Handler) http.Handler { 
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		 
		next.ServeHTTP(w, r)
	}) 
}

// Detect panics and send appropriate response to users
func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// defer this from happening until a panic DOES happen
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}