package lorawan_provider_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"machine_max_deveui_generator/deveui_service/domain_models"
	"machine_max_deveui_generator/lorawan_provider_client/client_models"
	"net/http"
	"time"
)



const baseURL = "http://europe-west1-machinemax-dev-d524.cloudfunctions.net/"

type Client struct {}

func NewDefaultClient() *Client {
	return new(Client)
}

func (c *Client) RegisterDevEUI(eui *domain_models.DevEUI) (*domain_models.DevEUI, error) {
	input := client_models.RegisterDevEUIInput{DevEUI: eui.ID}
	err := c.doPost("sensor-onboarding-sample", input)
	if err != nil {
		return nil, err
	}

	return eui, nil
}

func (c *Client) doPost(path string, body interface{}) error {
	return c.doRequest("POST", path, body)
}

func (c *Client) doRequest(method string, path string, body interface{}) error{
	reqURL := baseURL + path

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Create requests
	req, err := http.NewRequest(method, reqURL, bytes.NewBuffer(bodyJson))
	if err != nil {
		return fmt.Errorf("could not execute request! #1 (%s)", err.Error())
	}

	req.Header.Set("Content-type", "application/json")

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	switch  {
	case resp.StatusCode == http.StatusUnprocessableEntity:
		return client_models.ErrDevEUIAlreadyExists
	case resp.StatusCode >= 400:
		return fmt.Errorf("server responded with statuscode: %d", resp.StatusCode)
	default:
		return nil
	}
}