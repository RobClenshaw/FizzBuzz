package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	hostName := os.Getenv("HOSTNAME")
	divisor := os.Getenv("DIVISOR")
	outputPhrase := os.Getenv("OUTPUT_PHRASE")
	count := 0

	http.HandleFunc("/data/", func(w http.ResponseWriter, r *http.Request) {
		handleData(w, r, hostName, divisor, outputPhrase, &count)
	})
	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) { handleReady(w, r, &count) })
	http.ListenAndServe(":80", nil)
}

func handleData(w http.ResponseWriter, r *http.Request, hostName string, divisor string, outputPhrase string, count *int) {
	input := r.URL.Path[6:len(r.URL.Path)]
	number, err := strconv.Atoi(input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Expected a positive integer, but got "+input)
		return
	}

	intDivisor, _ := strconv.Atoi(divisor)

	payload := outputPayload{
		OutputString: getOutputString(number, intDivisor, outputPhrase),
		Host:         hostName,
	}

	output, _ := json.Marshal(payload)

	io.WriteString(w, string(output))

	*count++
}

func handleReady(w http.ResponseWriter, r *http.Request, count *int) {
	if *count > 0 {
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
