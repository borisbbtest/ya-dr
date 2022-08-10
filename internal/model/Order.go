package model

import "time"

const (
	StatusNew        = "NEW"
	StatusProcessing = "PROCESSING"
	StatusInvalid    = "INVALID"
	StatusProcessed  = "PROCESSED"
)

type DataOrders struct {
	Number      string    `json:"number"`
	Status      string    `json:"status"`
	Accrual     *int      `json:"accrual,omitempty"` //nullable
	Person      string    `json:"-"`
	Uploaded_at time.Time `json:"uploaded_at"`
}
