package service

import (
	"encoding/json"
	"fmt"
	"getnews/pkg/model"
	"io/ioutil"
	"net/http"
	"time"
)

// FetchWeiboHot fetches and parses Weibo hot search data into model.News format
func FetchWeiboHot() ([]model.News, error) {
	// Define the Weibo API URL or endpoint
	url := "https://weibo.com/ajax/side/hotSearch" // Replace with the actual API URL

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Weibo hot search: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read Weibo response: %v", err)
	}

	// Define a struct to match the JSON structure
	var result struct {
		Ok   int `json:"ok"`
		Data struct {
			// 解析realtime字段
			Realtime []struct {
				Word string `json:"word"`        // Title of the hot search topic
				Note string `json:"note"`        // Description of the hot search topic
				Num  int    `json:"num"`         // Popularity number
				Rank int    `json:"rank"`        // Rank of the hot search topic
				Link string `json:"word_scheme"` // Use word_scheme as link
			} `json:"realtime"`
		} `json:"data"`
	}

	// Parse the JSON response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse Weibo JSON data: %v", err)
	}

	// Convert parsed data into []model.News format
	var newsList []model.News
	for _, item := range result.Data.Realtime {
		// 避免Rank重复
		if item.Link == "" {
			continue
		}
		newsList = append(newsList, model.News{
			Title:     item.Word,
			Link:      item.Link,
			Rank:      item.Rank,
			Source:    "Weibo",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}
	return newsList, nil
}
