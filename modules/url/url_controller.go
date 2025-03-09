package url

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShortenURLRequest defines the request body for shortening a URL
type ShortenURLRequest struct {
	LongURL string `json:"long_url"`
}

// URLController handles URL shortening and redirection
type URLController struct {
	service *URLService
}

func NewURLController(service *URLService) *URLController {
	return &URLController{service: service}
}

// ShortenURL creates a short URL from a long URL
// @Summary Shorten a URL
// @Description Generates a short URL for a given long URL
// @Tags URL
// @Accept json
// @Produce json
// @Param request body ShortenURLRequest true "Long URL"
// @Success 201 {object} url.ShortenedURL  // ✅ Fully qualified reference
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /url/shorten [post]
func (c *URLController) ShortenURL(ctx *gin.Context) {
	var request ShortenURLRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	url, err := c.service.GenerateShortURL(request.LongURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Could not create short URL"})
		return
	}

	ctx.JSON(http.StatusCreated, url)
}

// RedirectToLongURL redirects from a short URL to the original URL
// @Summary Redirect to long URL
// @Description Fetches long URL from short URL and redirects
// @Tags URL
// @Param short_url path string true "Short URL"
// @Success 302
// @Failure 404 {object} ErrorResponse
// @Router /url/{short_url} [get]
func (c *URLController) RedirectToLongURL(ctx *gin.Context) {
	shortURL := ctx.Param("short_url")
	url, err := c.service.GetLongURL(shortURL)
	if err != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{Error: "Short URL not found"})
		return
	}

	ctx.Redirect(http.StatusFound, url.LongURL)
}
