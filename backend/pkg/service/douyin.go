package service

import (
	"encoding/json"
	"fmt"
	"getnews/pkg/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// FetchDouyinHot fetches and parses Douyin hot search data into model.News format
func FetchDouyinHot() ([]model.News, error) {
	// Step 1: Fetch the guest cookie for Douyin Web
	cookie, err := fetchDouyinGuestCookie()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch guest cookie: %v", err)
	}

	// Step 2: Define the Douyin API URL for hot search
	douyinURL := "https://www.douyin.com/aweme/v1/web/hot/search/list/?device_platform=webapp&aid=6383&channel=channel_pc_web&detail_list=1"

	// Create a new HTTP request with the cookie
	req, err := http.NewRequest("GET", douyinURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers, including the guest cookie and User-Agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	req.Header.Set("Cookie", cookie)

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Douyin hot search: %v", err)
	}
	defer resp.Body.Close()

	// Read and validate the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read Douyin response: %v", err)
	}

	// Parse the JSON response
	var result struct {
		StatusCode int `json:"status_code"`
		Data       struct {
			WordList []struct {
				Word      string `json:"word"`
				HotValue  int    `json:"hot_value"`
				Position  int    `json:"position"`
				GroupID   string `json:"group_id"`
				WordCover struct {
					URLList []string `json:"url_list"`
				} `json:"word_cover"`
			} `json:"word_list"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse Douyin JSON data: %v", err)
	}

	// Check if result contains valid data
	if len(result.Data.WordList) == 0 {
		return nil, fmt.Errorf("no hot search data found in Douyin response")
	}

	// Convert parsed data into []model.News format
	var newsList []model.News
	for _, item := range result.Data.WordList {
		encodedQuery := url.QueryEscape(item.Word)
		baseURL := "https://www.douyin.com/search/"
		finalURL := fmt.Sprintf("%s%s", baseURL, encodedQuery)

		newsList = append(newsList, model.News{
			Title:     item.Word,
			Link:      finalURL,
			Rank:      item.Position,
			Source:    "Douyin",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}
	return newsList, nil
}

// fetchDouyinGuestCookie fetches the guest cookie from the specified API
func fetchDouyinGuestCookie() (string, error) {
	// Define the API URL to fetch the guest cookie
	cookieAPIURL := "https://api.tikhub.io/api/v1/douyin/web/fetch_douyin_web_guest_cookie"

	// Create the request
	req, err := http.NewRequest("GET", cookieAPIURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create cookie request: %v", err)
	}

	// Set headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	req.Header.Set("Authorization", "Bearer YOUR_BEARER_TOKEN") // replace YOUR_BEARER_TOKEN with your actual token

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to get guest cookie: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read guest cookie response: %v", err)
	}

	// Parse the response JSON to extract the cookie
	var result struct {
		Cookie string `json:"cookie"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to parse cookie JSON data: %v", err)
	}

	// Check if cookie is empty
	if result.Cookie == "" {
		return "", fmt.Errorf("fetched cookie is empty")
	}

	return result.Cookie, nil
}
