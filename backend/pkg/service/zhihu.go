package service

import (
	"encoding/json"
	"fmt"
	"getnews/pkg/model"
	"io/ioutil"
	"net/http"
	"time"
)

// FetchZhihuHot fetches and parses Zhihu hot search data into model.News format
func FetchZhihuHot() ([]model.News, error) {
	zhihuURL := "https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total"

	// 创建请求
	req, err := http.NewRequest("GET", zhihuURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// 设置User-Agent以避免403错误
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read Zhihu response: %v", err)
	}

	// 定义一个结构体匹配知乎API返回的JSON结构
	var result struct {
		Data []struct {
			Target struct {
				ID            int    `json:"id"`
				Title         string `json:"title"`
				URL           string `json:"url"`
				AnswerCount   int    `json:"answer_count"`
				FollowerCount int    `json:"follower_count"`
				Excerpt       string `json:"excerpt"`
			} `json:"target"`
			DetailText string `json:"detail_text"` // 热度信息
		} `json:"data"`
	}

	// 解析JSON响应
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse Zhihu JSON data: %v", err)
	}

	// 转换为[]model.News格式
	var newsList []model.News
	for index, item := range result.Data {
		// 解析链接
		finalURL := fmt.Sprintf("https://www.zhihu.com/question/%d", item.Target.ID)

		// 构造model.News对象
		newsList = append(newsList, model.News{
			Title:     item.Target.Title,
			Link:      finalURL,
			Rank:      index + 1, // 按照顺序排列
			Source:    "Zhihu",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}
	return newsList, nil
}
