package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// ResponseReport struct that holds data of reports
type ResponseReport struct {
	Error      bool   `json:"error"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       struct {
		LastChecked  time.Time `json:"lastChecked"`
		Covid19Stats []struct {
			City       string `json:"city"`
			Province   string `json:"province"`
			Country    string `json:"country"`
			LastUpdate string `json:"lastUpdate"`
			KeyID      string `json:"keyId"`
			Confirmed  int    `json:"confirmed"`
			Deaths     int    `json:"deaths"`
			Recovered  int    `json:"recovered"`
		} `json:"covid19Stats"`
	} `json:"data"`
}

// GetSpecificReport handler for get reports
func GetSpecificReport(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["country"]
	url := "https://covid-19-coronavirus-statistics.p.rapidapi.com/v1/stats?country=" + key
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-host", "covid-19-coronavirus-statistics.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "523669ef40msh869770cdf015e26p1c8d52jsn51883614dbfa")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	s, err := formatStatistics([]byte(body))

	fmt.Println(s, err)
	fmt.Fprint(w, s.Data.Covid19Stats)

}

func formatStatistics(body []byte) (*ResponseReport, error) {
	var s = new(ResponseReport)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("Whoops", err)
	}
	return s, err
}
