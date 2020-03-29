package main 

import (
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
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
	handler:=http.HandlerFunc(getAllPatients)
	handler.ServeHTTP(rr,req)
	status:= rr.Code

	if status != http.StatusOK {
		t.Errorf("Handler return Wrong status, Expected %v got %v",http.StatusOK, status)
	}

}