package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"slices"
	"time"
)

const (
	apiBase   = "https://soybooru.com/api/danbooru/post/index.xml"
	limit     = 20
	userAgent = "SoyBooruAutoPosterPrototype/0.1"
)

func fetchPosts() (*Posts, error) {
	url := fmt.Sprintf("%s?page=1&limit=%d", apiBase, limit)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	var posts Posts
	if err := xml.NewDecoder(resp.Body).Decode(&posts); err != nil {
		return nil, err
	}

	return &posts, nil
}

func fetchNewPosts(lastMaxID int) ([]Post, int, error) {
	posts, err := fetchPosts()
	if err != nil {
		return nil, 0, err
	}

	newMaxID := lastMaxID
	var newPosts []Post
	for _, p := range slices.Backward(posts.Items) {
		if p.ID > lastMaxID {
			newPosts = append(newPosts, p)
		}
		if p.ID > newMaxID {
			newMaxID = p.ID
		}
	}

	return newPosts, newMaxID, nil
}
