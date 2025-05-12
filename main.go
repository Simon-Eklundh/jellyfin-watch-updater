package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type watchStatus struct {
	Items []struct {
		UserData struct {
			Played         bool    `json:"Played"`
			LastPlayedDate *string `json:"LastPlayedDate"`
		}
		Name string `json:"Name"`
		Id   string `json:"Id"`
	}
}

func main() {
	// Log the start of the program
	fmt.Println("starting jellyfin-watched")
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	// Define the API endpoint
	baseURL := os.Getenv("JELLYFIN_BASE_URL") + "/Items"

	// Create a URL object and add query parameters
	reqURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	query := reqURL.Query()
	query.Set("userId", os.Getenv("JELLYFIN_USER_ID"))
	query.Set("filters", "isPlayed,isNotFolder")
	query.Set("recursive", "true")
	reqURL.RawQuery = query.Encode() // Encode the query parameters

	// Create a new HTTP request
	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add headers to the request
	req.Header.Set("Authorization", "MediaBrowser Token="+os.Getenv("JELLYFIN_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer resp.Body.Close()
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}
	fmt.Println("Status Code:", resp.StatusCode)
	var test watchStatus
	err2 := json.Unmarshal(body, &test)
	if err2 != nil {
		fmt.Println("Error unmarshalling JSON:", err2)
		return
	}
	var watchedCount = 0

	for _, item := range test.Items {
		if item.UserData.LastPlayedDate == nil {
			setWatched(item)
			watchedCount++
		}
	}
	fmt.Printf("%d items set as watched with date\n", watchedCount)
}

// todo add the api calls for setting the item status to watched
func setWatched(item struct {
	UserData struct {
		Played         bool    `json:"Played"`
		LastPlayedDate *string `json:"LastPlayedDate"`
	}
	Name string `json:"Name"`
	Id   string `json:"Id"`
}) {
	// Define the API endpoint
	baseURL := fmt.Sprintf(os.Getenv("JELLYFIN_BASE_URL")+"/UserItems/%s/UserData", item.Id)
	// Create a URL object and add query parameters
	reqURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	// Add query parameters
	query := reqURL.Query()
	query.Set("userId", os.Getenv("JELLYFIN_USER_ID"))
	reqURL.RawQuery = query.Encode() // Encode the query parameters

	// Create the request body
	reqBody := fmt.Sprintf(`{"Played":true, "LastPlayedDate":"%s", "PlayCount": 1, "PlaybackPositionTicks": 0 }`, time.Now().UTC().Format(time.RFC3339))
	// Create a new HTTP request
	req, err := http.NewRequest("POST", reqURL.String(), strings.NewReader(reqBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add headers to the request
	req.Header.Set("Authorization", "MediaBrowser Token="+os.Getenv("JELLYFIN_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer resp.Body.Close()
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println("Error setting item as watched:", item.Name)
		fmt.Println("Status Code:", resp.StatusCode)
		fmt.Println("Response Body:", string(body))
	}

}
