package utils

import (
	"encoding/json"
	"net/http"
)

type Pagination struct {
    Page       int64 `json:"page"`
    Limit      int64 `json:"limit"`
    TotalItems int64 `json:"totalItems"`
    TotalPages int64 `json:"totalPages"`
}

type PaginatedData struct {
    Data       interface{} `json:"data"`
    Pagination Pagination  `json:"pagination"`
}

func SendPageData(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func SendPage(w http.ResponseWriter, data interface{}, page, limit, cnt int64) {
	paginatedData := PaginatedData{
		Data: data,
		Pagination: Pagination{
			Page:       page,
			Limit:      limit,
			TotalItems: cnt,
			TotalPages: (cnt + limit - 1) / limit, 
		},
	}

	SendPageData(w, http.StatusOK, paginatedData)
}