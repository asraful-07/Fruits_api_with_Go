package product

import (
	"fmt"
	"fruits-api/utils"
	"net/http"
	"strconv"
	"sync"
)

var ctn int64

// Get all fruits
func (h *Handler) GetFruits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqQuery := r.URL.Query()

	pageAsStr  := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")
	
	// parse query params
	page, err := strconv.Atoi(pageAsStr)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(limitAsStr)
	if err != nil || limit <= 0 {
		limit = 4
	}

	// Service → Repo → DB clean architecture maintain 
	fruits, err := h.svc.List(int64(page), int64(limit)) 
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to fetch fruits list")
		return
	}

	var wg sync.WaitGroup

    wg.Add(1)
	go func () {
	defer wg.Done()
		ctn1, err := h.svc.Count()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to fetch fruits list")
		return
	}
	 ctn = ctn1
	}()

    wg.Add(1)
	go func () {
	defer wg.Done()
		ctn2, err := h.svc.Count()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to fetch fruits list")
		return
	}
	fmt.Println(ctn2)
	}()

    wg.Add(1)
	go func () {
	defer wg.Done()
		ctn3, err := h.svc.Count()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to fetch fruits list")
		return
	}
	fmt.Println(ctn3)
	}()

	// time.Sleep(8 * time.Second)
	wg.Wait()

	utils.SendPage(w, fruits, int64(page), int64(limit), ctn)
}
