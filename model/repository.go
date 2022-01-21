package model

import "time"

type Repository struct {
	UpdateTime    time.Time `json:"update_time"`
	Description   string    `json:"description"`
	PullCount     int       `json:"pull_count"`
	CreationTime  time.Time `json:"creation_time"`
	ArtifactCount int       `json:"artifact_count"`
	ProjectID     int       `json:"project_id"`
	ID            int       `json:"id"`
	Name          string    `json:"name"`
}
