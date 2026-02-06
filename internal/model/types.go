package model

import "time"

type File struct {
	ID         string
	Filename   string
	Size       int64
	StorageKey string
	CreatedAt  time.Time
}
