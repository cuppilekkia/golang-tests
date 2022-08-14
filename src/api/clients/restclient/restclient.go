package restclient

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {

	jsonBytes, err := json.Marshal(body)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}

	return client.Do(request)
}
