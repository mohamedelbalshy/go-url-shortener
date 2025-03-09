package url

import (
	"go-url-shortener/database"
)

type URLRepository struct{}

func NewURLRepository() *URLRepository {
	return &URLRepository{}
}

func (r *URLRepository) CreateURL(url *ShortenedURL) error {
	return database.DB.Create(url).Error
}

func (r *URLRepository) GetURLByShort(shortURL string) (*ShortenedURL, error) {
	var url ShortenedURL
	err := database.DB.Where("short_url = ?", shortURL).First(&url).Error
	if err != nil {
		return nil, err
	}
	return &url, nil
}
