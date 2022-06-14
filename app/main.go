package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"rmtcapp/api"
)

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	publicData := make([]interface{}, len(data))
	for i, d := range data {
		publicData[i] = api.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}

func jobsHandler(jobs []interface{}) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, jobs)
	}
	return http.HandlerFunc(fn)
}

func okHandler(msg string) http.Handler {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	fn := func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte(msg + hostname))
	}
	return http.HandlerFunc(fn)
}

func main() {

	var jobs = []interface{}{
		api.Jobs{
			Name:         "Product Designer",
			Company:      "Partnero",
			ContractType: "FullTime - Remote",
			Salary:       60000,
		},
		api.Jobs{
			Name:         "SRE",
			Company:      "The Remote Company",
			ContractType: "FullTime - Remote",
			Salary:       80000,
		},
		api.Jobs{
			Name:         "Sr.Laravel and Vue.js developer",
			Company:      "MailerLite",
			ContractType: "FullTime - Remote",
			Salary:       80000,
		},
	}

	j := jobsHandler(jobs)
	ok := okHandler("Api operational: ")

	mux := http.NewServeMux()
	mux.Handle("/", ok)
	mux.Handle("/jobs", j)
	http.ListenAndServe(":8080", mux)

}
