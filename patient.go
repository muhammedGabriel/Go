package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// type MedicalStatus struct {
// 	Temperature   float32 `json:"temperature"`
// 	bloodPressure int     `json:"description"`
// }
type Patient struct {
	Id          int     `json:"id`
	Name        string  `json:"name"`
	Age         int     `json:"age"`
	Temperature float32 `json:"temperature"`
}

type Patients []Patient

func getAllPatients(w http.ResponseWriter, r *http.Request) {
	patients := Patients{
		Patient{Id: 110, Name: "Test Pateint", Age: 22, Temperature: 37},
		Patient{Id: 120, Name: "Test Pateint2", Age: 23, Temperature: 38},
	}
	fmt.Println("All Patients request")
	json.NewEncoder(w).Encode(patients)
}

func getSpecificPatient(w http.ResponseWriter, r *http.Request) {
	patients := Patients{
		Patient{Id: 110, Name: "Test_Pateint", Age: 22, Temperature: 37},
		Patient{Id: 120, Name: "Test Pateint2", Age: 23, Temperature: 38},
	}
	// we have to pass the request to the mux variable mux.Vars(request)
	// to extract the variables

	vars := mux.Vars(r) // this will return a map with var:[userInput]
	fmt.Println(vars)   //map[name:ahmad] so i need the value of the name key
	key := vars["name"]

	for _, patient := range patients {
		if patient.Name == key {
			json.NewEncoder(w).Encode(patient)

		}
	}
}
