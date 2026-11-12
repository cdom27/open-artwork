package artwork

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	defaultPage     = 1
	defaultPageSize = 12
	maxPageSize     = 12
)

// GetSearchParams parses and validates artwork list query parameters.
func GetSearchParams(c *gin.Context) (SearchParams, error) {
	page, err := queryPositiveInt(c, "page", defaultPage, 0)
	if err != nil {
		return SearchParams{}, err
	}

	pageSize, err := queryPositiveInt(c, "pageSize", defaultPageSize, maxPageSize)
	if err != nil {
		return SearchParams{}, err
	}

	sort := Sort(c.Query("sort"))
	if sort == "" {
		sort = SortRelevance
	}

	if !isValidSort(sort) {
		return SearchParams{}, fmt.Errorf("invalid sort %q", sort)
	}

	return SearchParams{
		Q:          c.Query("q"),
		Page:       page,
		PageSize:   pageSize,
		Sort:       sort,
		From:       queryDateBound(c, "from"),
		To:         queryDateBound(c, "to"),
		ObjectType: c.Query("objectType"),
		Medium:     c.Query("medium"),
		Culture:    c.Query("culture"),
	}, nil
}

// queryDateBound returns nil for a missing or invalid bound so the other
// date filter can still be applied.
func queryDateBound(c *gin.Context, key string) *int64 {
	value, exists := c.GetQuery(key)
	if !exists {
		return nil
	}

	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil
	}

	return &parsed
}

func queryPositiveInt(c *gin.Context, key string, defaultValue, maxValue int) (int, error) {
	value, exists := c.GetQuery(key)
	if !exists {
		return defaultValue, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil || parsed < 1 {
		return 0, fmt.Errorf("%s must be a positive integer", key)
	}
	if maxValue > 0 && parsed > maxValue {
		return maxValue, nil
	}

	return parsed, nil
}

func isValidSort(sort Sort) bool {
	switch sort {
	case SortRelevance, SortTitle, SortNewest, SortOldest, SortArtistAsc, SortArtistDesc:
		return true
	default:
		return false
	}
}
