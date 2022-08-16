package model

type DataBalance struct {
	Withdrawn      *float32 `json:"withdrawn"`
	CurrentAccrual *float32 `json:"current"` //nullable
	Person         *int     `json:"-"`
}
