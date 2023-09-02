package models

import "time"

type Comments struct {
	Content   string     `json:"content"`
	CommentBy string     `json:"commentBy"`
	CreatedAt *time.Time `json:"createdAt"`
}
