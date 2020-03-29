package main

import (
	"encoding/json"
	"strconv"

	"fmt"

	"net/http"

	"github.com/gorilla/mux"

	_ "strconv"
)

//Patient patient struct which will hold registered patients
type Patient struct {
	ID          int     `json:"id`
	Name        string  `json:"name"`
	Age         int     `json:"age"`
	Temperature float32 `json:"temperature"`
}

//Patients slices of patients
type Patients []Patient

func getAllPatients(w http.ResponseWriter, r *http.Request) {
	patients := Patients{
		Patient{ID: 110, Name: "Test Pateint", Age: 22, Temperature: 37},
		Patient{ID: 120, Name: "Test Pateint2", Age: 23, Temperature: 38},
	}
	fmt.Println("All Patients request")
	json.NewEncoder(w).Encode(patients)
}

func getSpecificPatient(w http.ResponseWriter, r *http.Request) {
	patients := Patients{
		Patient{ID: 110, Name: "Test_Pateint", Age: 22, Temperature: 37},
		Patient{ID: 120, Name: "Test Pateint2", Age: 23, Temperature: 38},
	}
	// we have to pass the request to the mux variable mux.Vars(request)
	// to extract the variables

	vars := mux.Vars(r) // this will return a map with var:[userInput]
	fmt.Println(vars)   //map[id:111] so i need the value of the id key
	key := vars["id"]
	temp, _ := strconv.Atoi(key)

	for _, patient := range patients {
		if patient.ID == temp {
			json.NewEncoder(w).Encode(patient)

		}
	}
}
