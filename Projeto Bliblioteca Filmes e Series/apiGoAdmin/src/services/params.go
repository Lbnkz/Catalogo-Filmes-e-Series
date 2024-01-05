package services

import (
	"net/http"
	"strconv"
)

func GetLimit(r *http.Request) int {
	limitStr := r.URL.Query().Get("limit")
	var limit int
	if limitStr == "" {
		limit = 1000
	} else {
		limit, _ = strconv.Atoi(limitStr)
	}

	return limit
}

func GetOffset(r *http.Request) int {
	offsetStr := r.URL.Query().Get("offset")
	var offset int
	if offsetStr == "" {
		offset = 0
	} else {
		offset, _ = strconv.Atoi(offsetStr)
	}

	return offset
}
