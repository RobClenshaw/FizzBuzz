package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		input := r.URL.Path[1:len(r.URL.Path)]
		number, err := strconv.Atoi(input)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Expected a positive integer, but got "+input)
			return
		}

		fizzResponse, fizzHost, err := getResponseFromHost("fizz", number)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Error talking to the fizz service")
			return
		}

		buzzResponse, buzzHost, err := getResponseFromHost("buzz", number)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Error talking to the buzz service")
			return
		}

		response := fizzResponse + buzzResponse
		if response == "" {
			response = input
		}
		io.WriteString(w, response+"\n")
		io.WriteString(w, fizzHost+"\n")
		io.WriteString(w, buzzHost)
	})

	http.ListenAndServe(":80", nil)
}

func getResponseFromHost(service string, number int) (string, string, error) {
	url := fmt.Sprintf("http://%s:80/data/%d", service, number)

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
