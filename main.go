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
)

type watchStatus struct {
    Items []struct {
        UserData struct {
            Played bool `json:"Played"`
        }
        Name string `json:"Name"`
        Id   string `json:"Id"`
    }
}

func main() {
    // Define the API endpoint
    baseURL := os.Getenv("JELLYFIN_BASE_URL")+ "/Items"

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
	req.Header.Set("Authorization", "MediaBrowser Token="+ os.Getenv("JELLYFIN_TOKEN"))
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

    for _, item := range test.Items {
        setWatched(item)
    }
}

// todo add the api calls for setting the item status to watched
func setWatched(item struct {
    UserData struct {
        Played bool `json:"Played"`
    }
    Name string `json:"Name"`
    Id   string `json:"Id"`
}) {
    // Define the API endpoint
    baseURL := fmt.Sprintf(os.Getenv("JELLYFIN_BASE_URL") + "/UserItems/%s/UserData", item.Id)
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
    reqBody := fmt.Sprintf(`{"Played":true, "LastPlayedDate":"%s"}`, time.Now().Format("2006-01-02T15:04:05Z"))

    // Create a new HTTP request
    req, err := http.NewRequest("POST", reqURL.String(), strings.NewReader(reqBody))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    // Add headers to the request
    req.Header.Set("Authorization", "MediaBrowser Token=" + os.Getenv("JELLYFIN_TOKEN"))
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
    fmt.Println("Response Body:", string(body))
}