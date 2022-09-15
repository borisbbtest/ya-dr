package model

import (
	"sync"
	"time"
)

type DataUser struct {
	Login            string      `json:"login"`
	Password         string      `json:"password"`
	SessionExpiredAt time.Time   `json:"-"`
	ID               int         `json:"-"`
	LocalMutex       *sync.Mutex `json:"-"`
}
