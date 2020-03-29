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
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/all", GetAllPatients).Methods(("GET"))
	r.HandleFunc("/addpatient", newPatientHandler).Methods("GET")
	r.HandleFunc("/reports/{country}", GetSpecificReport).Methods("GET")
	r.HandleFunc("/patient/{id}", GetSpecificPatient).Methods("GET")
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
