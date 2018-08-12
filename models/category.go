package models

import "time"

type Category struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ParentId    int64     `json:"parent_id"`
	SortOrder   int64     `json:"sort_order"`
	CreatedAt   time.Time `json:"created_at"`
}
