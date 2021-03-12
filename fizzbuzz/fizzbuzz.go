package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	fizzHost := os.Getenv("FIZZ_SERVICE_HOST")
	buzzHost := os.Getenv("BUZZ_SERVICE_HOST")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		input := r.URL.Path[1:len(r.URL.Path)]
		number, err := strconv.Atoi(input)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Expected a positive integer, but got "+input)
			return
		}

		fizzResponse, fizzHost, err := getResponseFromHost(fizzHost, number)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Error talking to the fizz service")
			return
		}

		buzzResponse, buzzHost, err := getResponseFromHost(buzzHost, number)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Error talking to the buzz service")
			return
		}

		io.WriteString(w, fizzResponse+buzzResponse+"\n")
		io.WriteString(w, fizzHost+"\n")
		io.WriteString(w, buzzHost)
	})

	http.ListenAndServe(":8080", nil)
}

func getResponseFromHost(service string, number int) (string, string, error) {
	url := fmt.Sprintf("http://%s:8080/data/%d", service, number)

	response, err := http.Get(url)

	if err != nil {
		return "", "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return "", "", err
	}

	var data map[string]interface{}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return "", "", err
	}

	outputString := data["OutputString"].(string)
	host := data["Host"].(string)

	return outputString, host, nil
}
