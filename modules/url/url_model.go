package url

import "time"

// ShortenedURL represents the shortened URL model
type ShortenedURL struct {
	ID        uint       `json:"id"`                   // Primary Key
	CreatedAt time.Time  `json:"created_at"`           // Timestamp
	UpdatedAt time.Time  `json:"updated_at"`           // Timestamp
	DeletedAt *time.Time `json:"deleted_at,omitempty"` // Nullable Timestamp
	ShortURL  string     `gorm:"uniqueIndex" json:"short_url"`
	LongURL   string     `gorm:"not null" json:"long_url"`
}
