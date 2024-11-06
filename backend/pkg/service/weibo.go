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

// FetchWeiboHot fetches and parses Weibo hot search data into model.News format
func FetchWeiboHot() ([]model.News, error) {
	weiboUrl := "https://weibo.com/ajax/side/hotSearch"

	// 创建请求
	req, err := http.NewRequest("GET", weiboUrl, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
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
		encodedQuery := url.QueryEscape(item.Link)
		baseURL := "https://s.weibo.com/weibo"
		finalURL := fmt.Sprintf("%s?q=%s&typeall=1&suball=1&Refer=g", baseURL, encodedQuery)
		newsList = append(newsList, model.News{
			Title: item.Word,
			Link:  finalURL,
			// 因为微博的排序是从0开始的
			Rank:      item.Rank + 1,
			Source:    "Weibo",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}
	return newsList, nil
}
