package model

import (
	"time"
)

type Session struct {
	UserID       string
	SessionID    string
	TimeAccessed time.Time
	MaxAge       int64
}
