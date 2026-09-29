package artwork

import (
	"api/internal/artwork/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func artworkBySlug(c *gin.Context) {
	artwork, constituents, err := database.FindArtworkBySlug(c.Param("slug"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Could not find artwork",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Artwork found",
		"data":    ToArtworkResponse(artwork, constituents),
	})
}

func RouteV1(r *gin.Engine) {
	v1 := r.Group("v1/api")

	artworkV1 := v1.Group("artworks")
	{
		artworkV1.GET(":slug", artworkBySlug)
	}
}
