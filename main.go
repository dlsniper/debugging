package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	//"debugging/subpkg"

	"debugging/subpkg"

	"github.com/gorilla/mux"
)

const (
	readTimeout = 50
	writeTimeout = 100
	idleTimeout = 1200
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	returnStatus := http.StatusOK
	message := fmt.Sprintf("Hello %s with status %d!", r.UserAgent(), returnStatus)
	w.WriteHeader(returnStatus)
	w.Write([]byte(message))
}

func main() {
	serverAddress := ":8080"
	l :=log.New(os.Stdout, "sample-srv ", log.LstdFlags | log.Lshortfile)
	m := mux.NewRouter()

	m.HandleFunc("/fact", factHandler)
	m.HandleFunc("/", indexHandler)

	srv := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
		IdleTimeout:  idleTimeout * time.Second,
		Handler:      m,
	}

	l.Println("server started")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func factHandler(w http.ResponseWriter, r *http.Request) {
	times := r.URL.Query().Get("times")
	t, err := strconv.ParseInt(times, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "got error: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "the result is: %d", subpkg.Factorial(t))
}
