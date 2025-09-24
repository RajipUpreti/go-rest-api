package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func init() {
	// Use dedicated env var for Kong Admin; fallback to local default
	kongAdminURL = os.Getenv("KONG_ADMIN_URL")
	if kongAdminURL == "" {
		kongAdminURL = "http://localhost:8001"
	}
}

func CreateKongConsumer(username string) (*KongConsumer, error) {
	payload := map[string]string{"username": username}
	body, _ := json.Marshal(payload)
	slog.Info("CreateKongConsumer route hit")
	resp, err := http.Post(kongAdminURL+"/consumers", "application/json", bytes.NewBuffer(body))
	// print resp
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 409 {
		// Consumer already exists (duplicate username)
		slog.Info("Consumer already exists")
		return &KongConsumer{Username: username}, nil
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
	key, err := CreateAPIKey(consumer.ID)
	if err != nil {
		slog.Info("Failed to create API key")
		return nil, err
	}
	slog.Info(key.Key)
	return &consumer, nil
}

// create new api key for the consumer
// http://localhost:8001/consumers/99142762-e64d-484f-bbb6-f9e6bb18e3d5/key-auth
func CreateAPIKey(consumerID string) (*KongAPIKey, error) {
	slog.Info("CreateAPIKey route hit")
	slog.Info(consumerID)
	url := fmt.Sprintf("%s/consumers/%s/key-auth", kongAdminURL, consumerID)
	slog.Info(url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	// Define a struct to parse the response
	var result struct {
		Key string `json:"key"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	fmt.Println("Extracted key:", result.Key)
	var key KongAPIKey
	key.Key = result.Key
	key.ID = consumerID
	return &key, nil
}

// no main here; this package provides Kong helper functions
