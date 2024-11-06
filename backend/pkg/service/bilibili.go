package service

import (
	"encoding/json"
	"fmt"
	"getnews/pkg/model"
	"io/ioutil"
	"net/http"
	"time"
)

// FetchBilibiliPopular fetches popular videos from Bilibili and parses the data into model.News format
func FetchBilibiliPopular() ([]model.News, error) {
	// Define the Bilibili popular API URL
	bilibiliURL := "https://api.bilibili.com/x/web-interface/popular?ps=50&pn=1"

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", bilibiliURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Referer", "https://www.bilibili.com/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from Bilibili API: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read Bilibili response: %v", err)
	}

	// Define the JSON structure to match Bilibili's response format
	var result struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			List []struct {
				Link    string `json:"short_link_v2"`
				Title   string `json:"title"`
				Pubdate int64  `json:"pubdate"`
			} `json:"list"`
		} `json:"data"`
	}

	// Unmarshal JSON response into result struct
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse Bilibili JSON data: %v", err)
	}

	// Check if the API returned an error
	if result.Code != 0 {
		return nil, fmt.Errorf("API error: %s", result.Message)
	}

	// Transform the data into the model.News format
	var newsList []model.News
	for i, item := range result.Data.List {
		timestamp := time.Unix(item.Pubdate, 0)

		newsList = append(newsList, model.News{
			Title:     item.Title,
			Link:      item.Link,
			Rank:      i + 1,
			Source:    "BiliBili",
			CreatedAt: timestamp,
			UpdatedAt: time.Now(),
		})
	}
	return newsList, nil
}
