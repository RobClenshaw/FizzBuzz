package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
)

var count = 0
var hostName string = os.Getenv("HOSTNAME")
var divisor string = os.Getenv("DIVISOR")
var outputPhrase string = os.Getenv("OUTPUT_PHRASE")

func main() {
	http.HandleFunc("/data/", handleData)
	http.HandleFunc("/ready", handleReady)
	http.ListenAndServe(":80", nil)
}

func handleData(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Path[6:len(r.URL.Path)]
	number, err := strconv.Atoi(input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Expected a positive integer, but got "+input)
		return
	}

	divisor, _ := strconv.Atoi(divisor)

	payload := outputPayload{
		OutputString: getOutputString(number, divisor, outputPhrase),
		Host:         hostName,
	}

	output, _ := json.Marshal(payload)

	io.WriteString(w, string(output))

	count++
}

func handleReady(w http.ResponseWriter, r *http.Request) {
	if count > 0 {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getOutputString(input int, divisor int, outputPhrase string) string {
	if input%divisor == 0 {
		return outputPhrase
	}

	return ""
}

type outputPayload struct {
	OutputString string
	Host         string
}
