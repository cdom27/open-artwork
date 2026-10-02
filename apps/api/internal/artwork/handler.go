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

func artworkPreviews(c *gin.Context) {
	params, err := GetSearchParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	previewsResponse, err := database.FindArtworks(
		string(params.Sort),
		params.Page,
		params.PageSize,
		params.From,
		params.To,
		params.ObjectType,
		params.Medium,
		params.Culture,
		params.Q,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Could not find artworks",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Artworks found",
		"data":    ToArtworkPreviews(previewsResponse),
	})
}

func RouteV1(r *gin.Engine) {
	v1 := r.Group("v0/api")

	artworkV1 := v1.Group("artworks")
	{
		artworkV1.GET("", artworkPreviews)
		artworkV1.GET(":slug", artworkBySlug)
	}
}
