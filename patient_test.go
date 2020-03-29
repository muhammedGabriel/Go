package main 

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	_"github.com/gorilla/mux"

)


func TestGetAllPAtients(t *testing.T){
	// we have prepeared the request 
	req,err := http.NewRequest("GET","/all",nil)
	// we have handled the error
	if err != nil{
		fmt.Println("Error in wrapping a request")
		t.Fatal(err)
	}
	// we need the response recorder
	rr := httptest.NewRecorder()
	handler:=http.HandlerFunc(GetAllPatients)
	handler.ServeHTTP(rr,req)
	status:= rr.Code

	if status != http.StatusOK {
		t.Errorf("Handler return Wrong status, Expected %v got %v",http.StatusOK, status)
	}

}



func TestGetSpecificPatient (t *testing.T){
	// Wrapping the request
	req,err:=http.NewRequest("GET","/patient/{id}",nil)
	//error handling
	if err!=nil{
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "110")
	req.URL.RawQuery = q.Encode()
	//creating response recorder
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetSpecificPatient)
	handler.ServeHTTP(rr, req)

	expected:= `{"ID":110,"name":"Test_Pateint","age":22,"temperature":37}`
	if rr.Body.String() != expected {
		t.Errorf("Unexpected response, Expected %v |||| but got %v",expected, rr.Body.String())
	}

}