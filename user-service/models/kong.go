package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

// Structs for Kong API response
type KongConsumer struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type KongAPIKey struct {
	Key string `gorm:"uniqueIndex:uidx_users_api_key" json:"key" `
	ID  string `json:"id"`
}

var kongAdminURL string

// init to connect with kong
func init() {
	// Use dedicated env var for Kong Admin; fallback to local default
	kongAdminURL = os.Getenv("KONG_ADMIN_URL")
	if kongAdminURL == "" {
		kongAdminURL = "http://localhost:8001"
	}
	slog.Info(kongAdminURL)
	resp, err := http.Get(kongAdminURL)
	if err != nil {
		slog.Error("Failed to connect to Kong Admin", "error", err)
	}
	if resp.StatusCode != 200 {
		slog.Error("Failed to connect to Kong Admin", "status", resp.Status)
		// exit with code 0
		os.Exit(-1)
	}
	defer resp.Body.Close()
	slog.Info("Connected to Kong Admin")

}

// create a kong consumer
func CreateKongConsumer(email string) (*KongConsumer, error) {
	payload := map[string]string{"email": email}
	body, _ := json.Marshal(payload)
	slog.Info("CreateKongConsumer route hit")
	resp, err := http.Post(kongAdminURL+"/consumers", "application/json", bytes.NewBuffer(body))
	// print resp
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 409 {
		// Consumer already exists (duplicate email)
		slog.Info("Consumer already exists")
		return &KongConsumer{Username: email}, nil
	}

	if resp.StatusCode >= 300 {
		slog.Info("Failed to create consumer")
		return nil, fmt.Errorf("failed to create consumer, status: %s", resp.Status)
	}

	var consumer KongConsumer
	if err := json.NewDecoder(resp.Body).Decode(&consumer); err != nil {
		slog.Info("Failed to decode consumer")
		return nil, err
	}
	slog.Info(consumer.ID)
	slog.Info(consumer.Username)

	return &consumer, nil
}

// create new api key for the consumer
func CreateAPIKey(consumerID string) (*KongAPIKey, error) {
	url := fmt.Sprintf("%s/consumers/%s/key-auth", kongAdminURL, consumerID)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		slog.ErrorContext(nil, "Error creating request", "error", err)

	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.ErrorContext(nil, "Error sending request", "error", err)

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.ErrorContext(nil, "Error reading response", "error", err)

	}

	// Define a struct to parse the response
	var result struct {
		Key string `json:"key"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		slog.ErrorContext(nil, "Error parsing JSON", "error", err)

	}

	fmt.Println("Extracted key:", result.Key)
	var key KongAPIKey
	key.Key = result.Key
	key.ID = consumerID
	return &key, nil
}

// no main here; this package provides Kong helper functions
