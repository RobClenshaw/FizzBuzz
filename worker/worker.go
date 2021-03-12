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

	http.HandleFunc("/data/", func(w http.ResponseWriter, r *http.Request) {
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
	})

	http.ListenAndServe(":8080", nil)
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
