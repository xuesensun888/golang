package models

import "time"

type FaultReport struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`       // 必填
	Description string    `json:"description"` // 必填
	Severity    string    `json:"severity"`    // 可选
	Reporter    string    `json:"reporter"`    // 必填
	Timestamp   time.Time `json:"timestamp"`
}
