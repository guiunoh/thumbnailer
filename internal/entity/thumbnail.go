package entity

import (
	"time"
)

type Thumbnail struct {
	ID     ID        `json:"id" gorm:"primaryKey"`
	Type   string    `json:"type"`
	Data   []byte    `json:"data"`
	Expiry time.Time `json:"expiry"`
}

func NewThumbnail(data []byte, imageType string) *Thumbnail {
	const expiry = 3 * time.Minute
	return &Thumbnail{
		ID:     NewID(),
		Data:   data,
		Type:   imageType,
		Expiry: time.Now().Add(expiry),
	}
}
