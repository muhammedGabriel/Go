package main

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := newRouter()
	http.ListenAndServe(":8081", r)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", homePage).Methods("GET")
	r.HandleFunc("/all", getAllPatients).Methods(("GET"))
	r.HandleFunc("/addpatient", newPatientHandler).Methods("GET")
	r.HandleFunc("/reports", ServeHTTP).Methods("GET")
	r.HandleFunc("/patient/{name}", getSpecificPatient).Methods("GET")
	return r
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home page requested")
	fmt.Fprintf(w, "This is my home page")
}

func newPatientHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new patient is about to be added")
	fmt.Fprintf(w, "New Patient Record")
}
