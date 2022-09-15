package model

import (
	"time"
)

type Wallet struct {
	Person      int       `json:"-"`
	Order       string    `json:"order"`
	Sum         float32   `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}
