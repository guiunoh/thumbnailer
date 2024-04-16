package thumbnail

import (
	"time"

	"github.com/google/uuid"
)

type Thumbnail struct {
	ID     uuid.UUID `json:"id" gorm:"primaryKey"`
	Type   string    `json:"type"`
	Data   []byte    `json:"data"`
	Expiry time.Time `json:"expiry"`
}
