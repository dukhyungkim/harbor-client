package model

import "time"

type Tag struct {
	ID           int       `json:"id"`
	RepositoryID int       `json:"repository_id"`
	ArtifactID   int       `json:"artifact_id"`
	Name         string    `json:"name"`
	PushTime     time.Time `json:"push_time"`
	PullTime     time.Time `json:"pull_time"`
	Signed       bool      `json:"signed"`
	Immutable    bool      `json:"immutable"`
}
