package model

import "time"

type Project struct {
	UpdateTime         time.Time `json:"update_time"`
	OwnerName          string    `json:"owner_name"`
	Name               string    `json:"name"`
	Deleted            bool      `json:"deleted"`
	OwnerID            int       `json:"owner_id"`
	RepoCount          int       `json:"repo_count"`
	ChartCount         int       `json:"chart_count"`
	CreationTime       time.Time `json:"creation_time"`
	Togglable          bool      `json:"togglable"`
	CurrentUserRoleID  int       `json:"current_user_role_id"`
	CurrentUserRoleIDs []int     `json:"current_user_role_ids"`
	CveAllowlist       struct {
		UpdateTime time.Time `json:"update_time"`
		Items      []struct {
			CveID string `json:"cve_id"`
		} `json:"items"`
		ProjectID    int       `json:"project_id"`
		CreationTime time.Time `json:"creation_time"`
		ID           int       `json:"id"`
		ExpiresAt    int       `json:"expires_at"`
	} `json:"cve_allowlist"`
	ProjectID  int `json:"project_id"`
	RegistryID int `json:"registry_id"`
	Metadata   struct {
		EnableContentTrust   string `json:"enable_content_trust"`
		Public               string `json:"public"`
		AutoScan             string `json:"auto_scan"`
		Severity             string `json:"severity"`
		RetentionID          string `json:"retention_id"`
		ReuseSysCveAllowlist string `json:"reuse_sys_cve_allowlist"`
		PreventVul           string `json:"prevent_vul"`
	} `json:"metadata"`
}
