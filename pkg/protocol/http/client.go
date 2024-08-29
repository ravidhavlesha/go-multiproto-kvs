package http

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// HTTPClient represents a HTTP client.
type HTTPClient struct {
	address string
	client  *http.Client
}

// NewHTTPClient initializes a new HTTP client
func NewHTTPClient(address string) *HTTPClient {
	return &HTTPClient{address: address, client: &http.Client{}}
}

func (client *HTTPClient) Get(key string) (string, error) {
	// Use url.Values to safely encode parameters
	params := url.Values{}
	params.Add("key", key)

	// Construct the full URL with query parameters
	u, err := url.Parse(client.address)
	if err != nil {
		return "", fmt.Errorf("invalid address: %w", err)
	}
	u.Path = "/get"
	u.RawQuery = params.Encode()

	// Create a GET request
	request, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to create SET request: %w", err)
	}

	// Send the request and handle the response
	response, err := client.client.Do(request)
	if err != nil {
		return "", fmt.Errorf("failed to send SET request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return "", fmt.Errorf("error response from server:%s", string(body))
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	return string(body), nil
}

func (client *HTTPClient) Set(key, value string) error {
	// Use url.Values to safely encode parameters
	params := url.Values{}
	params.Add("key", key)
	params.Add("value", value)

	// Construct the full URL with query parameters
	u, err := url.Parse(client.address)
	if err != nil {
		return fmt.Errorf("invalid address: %w", err)
	}
	u.Path = "/set"
	u.RawQuery = params.Encode()

	// Create a POST request
	request, err := http.NewRequest(http.MethodPost, u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create SET request: %w", err)
	}

	// Send the request and handle the response
	response, err := client.client.Do(request)
	if err != nil {
		return fmt.Errorf("failed to send SET request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return fmt.Errorf("SET request failed with status %d: %s", response.StatusCode, string(body))
	}

	return nil
}

func (client *HTTPClient) Delete(key string) error {
	// Use url.Values to safely encode parameters
	params := url.Values{}
	params.Add("key", key)

	// Construct the full URL with query parameters
	u, err := url.Parse(client.address)
	if err != nil {
		return fmt.Errorf("invalid address: %w", err)
	}
	u.Path = "/delete"
	u.RawQuery = params.Encode()

	// Create a DELETE request
	request, err := http.NewRequest(http.MethodDelete, u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create DELETE request: %w", err)
	}

	// Send the request and handle the response
	response, err := client.client.Do(request)
	if err != nil {
		return fmt.Errorf("failed to send DELETE request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return fmt.Errorf("DELETE request failed with status %d: %s", response.StatusCode, string(body))
	}

	return nil
}
