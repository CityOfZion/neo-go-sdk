package neo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/CityOfZion/neo-go-sdk/neo/models/request"
)

func executeRequest(method string, bodyParameters []interface{}, nodeURI string, model interface{}) error {
	var body []byte
	var err error

	if bodyParameters == nil {
		body, err = request.NewBody(method)
		if err != nil {
			return err
		}
	} else {
		body, err = request.NewBodyWithParameters(method, bodyParameters)
		if err != nil {
			return err
		}
	}

	ioBody := bytes.NewReader(body)

	request, err := http.NewRequest("POST", nodeURI, ioBody)
	if err != nil {
		return err
	}

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode > 200 {
		return fmt.Errorf(
			"Non-200 status code returned from NEO node, got: '%d'",
			response.StatusCode,
		)
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &model)
	if err != nil {
		return err
	}

	return nil
}
