package request

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Action  string           `json:"action"`
	RefType string           `json:"ref_type"`
	Ref     string           `json:"ref"`
	Commits []map[string]any `json:"commits"`
}

type Repo struct {
	Name string `json:"name"`
}

type GithubUserEvent struct {
	Type    string  `json:"type"`
	Repo    Repo    `json:"repo"`
	Payload Payload `json:"payload"`
}

func FetchGithubUserEvents(username string) ([]GithubUserEvent, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var e []GithubUserEvent
	if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
		return nil, err
	}

	return e, nil
}
