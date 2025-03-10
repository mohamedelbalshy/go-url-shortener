package url

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"go-url-shortener/database"
)

// Service interface
type URLServiceInterface interface {
	GenerateShortURL(longURL string) (*ShortenedURL, error)
	GetLongURL(shortURL string) (*ShortenedURL, error)
}

type URLService struct {
	repo *URLRepository
}

func NewURLService(repo *URLRepository) URLServiceInterface {
	return &URLService{repo: repo}
}

// Generate a random short URL
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// Shorten URL and store in Redis
func (s *URLService) GenerateShortURL(longURL string) (*ShortenedURL, error) {
	shortURL := generateRandomString(6)
	url := &ShortenedURL{
		ShortURL: shortURL,
		LongURL:  longURL,
	}

	err := s.repo.CreateURL(url)
	if err != nil {
		return nil, err
	}

	// Store mapping in Redis
	ttl, _ := strconv.Atoi(os.Getenv("REDIS_TTL"))
	err = database.RedisClient.Set(context.Background(), shortURL, longURL, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		fmt.Println("Failed to cache URL in Redis:", err)
	}

	return url, nil
}

// Get URL from Redis or fallback to DB
func (s *URLService) GetLongURL(shortURL string) (*ShortenedURL, error) {
	// Try Redis first
	longURL, err := database.RedisClient.Get(context.Background(), shortURL).Result()
	if err == nil {
		return &ShortenedURL{ShortURL: shortURL, LongURL: longURL}, nil
	}

	// Fallback to PostgreSQL if not in Redis
	url, err := s.repo.GetURLByShort(shortURL)
	if err != nil {
		return nil, errors.New("URL not found")
	}

	// Cache in Redis for future requests
	ttl, _ := strconv.Atoi(os.Getenv("REDIS_TTL"))
	err = database.RedisClient.Set(context.Background(), shortURL, url.LongURL, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		fmt.Println("Failed to cache URL in Redis:", err)
	}

	return url, nil
}
