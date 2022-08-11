package model

import "time"

const (
	StatusNew        = "NEW"
	StatusProcessing = "PROCESSING"
	StatusInvalid    = "INVALID"
	StatusProcessed  = "PROCESSED"
)

type DataOrder struct {
	Number     string    `json:"number"`
	Status     string    `json:"status"`
	Accrual    *float64  `json:"accrual,omitempty"` //nullable
	Person     string    `json:"-"`
	UploadedAt time.Time `json:"uploaded_at"`
}
