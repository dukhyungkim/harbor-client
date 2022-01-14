package model

import (
	"fmt"
	"strings"
)

type ListTagsParams struct {
	Query               string
	Sort                string
	Page                int64
	PageSize            int64
	WithSignature       bool
	WithImmutableStatus bool
}

func NewListTagsParams() *ListTagsParams {
	return &ListTagsParams{
		Page:     1,
		PageSize: 10,
	}
}

func (p *ListTagsParams) ToParamString() string {
	params := []string{"?"}

	if p.Query != "" {
		params = append(params, fmt.Sprintf("q=%s", p.Query))
	}

	if p.Page != 0 {
		params = append(params, fmt.Sprintf("page=%d", p.Page))
	}

	if p.PageSize != 0 {
		params = append(params, fmt.Sprintf("page_size=%d", p.PageSize))
	}

	if p.Sort != "" {
		params = append(params, fmt.Sprintf("sort=%s", p.Sort))
	}

	if p.WithSignature {
		params = append(params, fmt.Sprintf("%t", p.WithSignature))
	}

	if p.WithImmutableStatus {
		params = append(params, fmt.Sprintf("%t", p.WithImmutableStatus))
	}

	return strings.Join(params, "&")
}
