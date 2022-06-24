package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", getNameParamHandler).Methods(http.MethodGet)
	router.HandleFunc("/bad", getBadHandler).Methods(http.MethodGet)
	router.HandleFunc("/data", postDataHandler).Methods(http.MethodPost)
	router.HandleFunc("/headers", postHeaderSumHandler).Methods(http.MethodPost)
	router.NotFoundHandler = http.HandlerFunc(defaultHandler)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func postHeaderSumHandler(w http.ResponseWriter, r *http.Request) {
	a := r.Header.Values("a")
	b := r.Header.Values("b")
	aInt, err := strconv.Atoi(a[0])
	if err != nil {
		log.Fatal(err)
	}
	bInt, err := strconv.Atoi(b[0])
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("a+b", fmt.Sprintf("%d", aInt+bInt))
	w.WriteHeader(http.StatusOK)
}

func postDataHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "I got message:\n%s", body)
}

func getBadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func getNameParamHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %v!", vars["PARAM"])
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
