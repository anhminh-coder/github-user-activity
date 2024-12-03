package activity

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GithubActivity struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Repo    `json:"repo"`
	Payload `json:"payload"`
}

type Repo struct {
	Name string `json:"name"`
}

type Payload struct {
	Action  string   `json:"action,omitempty"`
	Ref     string   `json:"ref,omitempty"`
	RefType string   `json:"ref_type,omitempty"`
	Commits []Commit `json:"commits,omitempty"`
}

type Commit struct {
	Message string `json:"message"`
}

func FetchGithubActivity(username string) (*[]GithubActivity, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, ErrInternalServer
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrUserNotFound
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ErrFetchingData(resp.StatusCode)
	}

	defer resp.Body.Close()
	var activities []GithubActivity
	err = json.NewDecoder(resp.Body).Decode(&activities)
	if err != nil {
		return nil, ErrInternalServer
	}

	return &activities, nil
}

func DisplayActivities(username string, activities *[]GithubActivity) {
	fmt.Printf("GitHub Activities for User: %s\n", username)
	fmt.Println("-----------------------------------")
	if activities == nil || len(*activities) == 0 {
		fmt.Println("No activities found.")
		return
	}

	for _, event := range *activities {
		var action string
		switch event.Type {
		case "PushEvent":
			commitCount := len(event.Payload.Commits)
			action = fmt.Sprintf("Pushed %d commit(s) to %s", commitCount, event.Repo.Name)
		case "IssuesEvent":
			action = fmt.Sprintf("%s an issue in %s", event.Payload.Action, event.Repo.Name)
		case "WatchEvent":
			action = fmt.Sprintf("Starred %s", event.Repo.Name)
		case "ForkEvent":
			action = fmt.Sprintf("Forked %s", event.Repo.Name)
		case "CreateEvent":
			action = fmt.Sprintf("Created %s in %s", event.Payload.RefType, event.Repo.Name)
		default:
			action = fmt.Sprintf("%s in %s", event.Type, event.Repo.Name)
		}
		fmt.Println(action)
	}
}
