package model

type WebhookEvent struct {
	EventType string `json:"event_type"`
	Events    []struct {
		Project     string `json:"project"`
		RepoName    string `json:"repo_name"`
		Tag         string `json:"tag"`
		FullName    string `json:"full_name"`
		TriggerTime int64  `json:"trigger_time"`
		ImageId     string `json:"image_id"`
		ProjectType string `json:"project_type"`
	} `json:"events"`
}
