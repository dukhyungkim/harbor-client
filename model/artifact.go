package model

import "time"

type Artifact struct {
	Size         int       `json:"size"`
	PushTime     time.Time `json:"push_time"`
	ScanOverview struct {
		AdditionalProp1 struct {
			Scanner struct {
				Version string `json:"version"`
				Vendor  string `json:"vendor"`
				Name    string `json:"name"`
			} `json:"scanner"`
			StartTime  time.Time `json:"start_time"`
			ScanStatus string    `json:"scan_status"`
			Summary    struct {
				Fixable int `json:"fixable"`
				Total   int `json:"total"`
				Summary struct {
					High     int `json:"High"`
					Critical int `json:"Critical"`
				} `json:"summary"`
			} `json:"summary"`
			CompletePercent int       `json:"complete_percent"`
			EndTime         time.Time `json:"end_time"`
			Duration        int       `json:"duration"`
			ReportID        string    `json:"report_id"`
			Severity        string    `json:"severity"`
		} `json:"additionalProp1"`
		AdditionalProp2 struct {
			Scanner struct {
				Version string `json:"version"`
				Vendor  string `json:"vendor"`
				Name    string `json:"name"`
			} `json:"scanner"`
			StartTime  time.Time `json:"start_time"`
			ScanStatus string    `json:"scan_status"`
			Summary    struct {
				Fixable int `json:"fixable"`
				Total   int `json:"total"`
				Summary struct {
					High     int `json:"High"`
					Critical int `json:"Critical"`
				} `json:"summary"`
			} `json:"summary"`
			CompletePercent int       `json:"complete_percent"`
			EndTime         time.Time `json:"end_time"`
			Duration        int       `json:"duration"`
			ReportID        string    `json:"report_id"`
			Severity        string    `json:"severity"`
		} `json:"additionalProp2"`
		AdditionalProp3 struct {
			Scanner struct {
				Version string `json:"version"`
				Vendor  string `json:"vendor"`
				Name    string `json:"name"`
			} `json:"scanner"`
			StartTime  time.Time `json:"start_time"`
			ScanStatus string    `json:"scan_status"`
			Summary    struct {
				Fixable int `json:"fixable"`
				Total   int `json:"total"`
				Summary struct {
					High     int `json:"High"`
					Critical int `json:"Critical"`
				} `json:"summary"`
			} `json:"summary"`
			CompletePercent int       `json:"complete_percent"`
			EndTime         time.Time `json:"end_time"`
			Duration        int       `json:"duration"`
			ReportID        string    `json:"report_id"`
			Severity        string    `json:"severity"`
		} `json:"additionalProp3"`
	} `json:"scan_overview"`
	Tags []struct {
		RepositoryID int       `json:"repository_id"`
		Name         string    `json:"name"`
		PushTime     time.Time `json:"push_time"`
		PullTime     time.Time `json:"pull_time"`
		Signed       bool      `json:"signed"`
		ID           int       `json:"id"`
		Immutable    bool      `json:"immutable"`
		ArtifactID   int       `json:"artifact_id"`
	} `json:"tags"`
	PullTime time.Time `json:"pull_time"`
	Labels   []struct {
		UpdateTime   time.Time `json:"update_time"`
		Description  string    `json:"description"`
		Color        string    `json:"color"`
		CreationTime time.Time `json:"creation_time"`
		Deleted      bool      `json:"deleted"`
		Scope        string    `json:"scope"`
		ProjectID    int       `json:"project_id"`
		ID           int       `json:"id"`
		Name         string    `json:"name"`
	} `json:"labels"`
	References []struct {
		Platform struct {
			Os           string   `json:"os"`
			Variant      string   `json:"variant"`
			Architecture string   `json:"architecture"`
			OsFeatures   []string `json:"'os.features'"`
			OsVersion    string   `json:"'os.version'"`
		} `json:"platform"`
		ChildDigest string   `json:"child_digest"`
		Urls        []string `json:"urls"`
		ParentID    int      `json:"parent_id"`
		ChildID     int      `json:"child_id"`
		Annotations struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"annotations"`
	} `json:"references"`
	ManifestMediaType string `json:"manifest_media_type"`
	ExtraAttrs        struct {
		AdditionalProp1 struct {
		} `json:"additionalProp1"`
		AdditionalProp2 struct {
		} `json:"additionalProp2"`
		AdditionalProp3 struct {
		} `json:"additionalProp3"`
	} `json:"extra_attrs"`
	ID            int    `json:"id"`
	Digest        string `json:"digest"`
	Icon          string `json:"icon"`
	RepositoryID  int    `json:"repository_id"`
	AdditionLinks struct {
		AdditionalProp1 struct {
			Href     string `json:"href"`
			Absolute bool   `json:"absolute"`
		} `json:"additionalProp1"`
		AdditionalProp2 struct {
			Href     string `json:"href"`
			Absolute bool   `json:"absolute"`
		} `json:"additionalProp2"`
		AdditionalProp3 struct {
			Href     string `json:"href"`
			Absolute bool   `json:"absolute"`
		} `json:"additionalProp3"`
	} `json:"addition_links"`
	MediaType   string `json:"media_type"`
	ProjectID   int    `json:"project_id"`
	Type        string `json:"type"`
	Annotations struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"annotations"`
}
