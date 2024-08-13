package main

import (
	"fmt"
	"net/http"
)

// Handler is a function that handles an HTTP request.
type Handler func(w http.ResponseWriter, r *http.Request) error // Handler type is a function that takes http.ResponseWriter and *http.Request as arguments and returns an error

// App is a struct that represents an HTTP application.
type App struct {
	mux *http.ServeMux
}

// NewApp creates a new App.
func NewApp() *App {
	return &App{
		mux: http.NewServeMux(),
	}
}

// Handle registers a new handler for a given pattern.
func (a *App) Handle(pattern string, handler Handler) {
	a.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			http.Error(w, fmt.Sprintf("an error occurred: %v", err), http.StatusInternalServerError)
		}
	})
}

// ServeHTTP implements the http.Handler interface.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

// ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
func (a *App) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, a)
}

// ListenAndServeTLS listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
func (a *App) ListenAndServeTLS(addr, certFile, keyFile string) error {
	return http.ListenAndServeTLS(addr, certFile, keyFile, a)
}

// function main
func main() {
	// Create a new App
	app := NewApp()

	// Register a handler for the "/" pattern
	app.Handle("/", func(w http.ResponseWriter, r *http.Request) error {
		_, err := fmt.Fprintln(w, "Hello, World!")
		return err
	})

	// Register a handler for the "/" pattern
	app.Handle("/test", func(w http.ResponseWriter, r *http.Request) error {
		_, err := fmt.Fprintln(w, "Welcome to test page!")
		return err
	})

	// Listen and serve on port 8080
	if err := app.ListenAndServe(":8080"); err != nil {
		fmt.Printf("server error: %v\n", err)
	}
}
