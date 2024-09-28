package web

import (
	"errors"
	"strconv"
	"strings"
	"time"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/claim"

	"github.com/gin-gonic/gin"
)

func ParseStringArrayQueryValue(value string, delimiter string) []string {
	parsedValue := strings.Split(value, delimiter)
	if len(parsedValue) == 1 && parsedValue[0] == "" {
		return []string{}
	}

	return parsedValue
}

func ParseUint64QueryValue(value string) uint64 {
	uint64Value, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0
	}

	return uint64Value
}

func ParseTimeQueryValue(value string, layout string) time.Time {
	timeValue, err := time.Parse(layout, value)
	if err != nil {
		return time.Time{}
	}

	return timeValue
}

func ParseBoolQueryValue(value string) bool {
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}

	return boolValue
}

func ParseBoolPointerQueryValue(value string) *bool {
	boolValue := ParseBoolQueryValue(value)
	if boolValue {
		return &boolValue
	}

	return nil
}

func GetClaimsFromContext(c *gin.Context) (*domain.Claims, error) {
	claims, exists := c.Get("claims")
	if !exists {
		return nil, errors.New("no claims found in context")
	}

	userClaims, ok := claims.(*domain.Claims)
	if !ok {
		return nil, errors.New("invalid claims type")
	}

	return userClaims, nil
}
