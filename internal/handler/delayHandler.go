package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/JoaoDallagnol/API-Delay-Simulator-for-Performance-Testing.git/internal/model"
	"github.com/JoaoDallagnol/API-Delay-Simulator-for-Performance-Testing.git/internal/service"
)

func DelayHandler(w http.ResponseWriter, r *http.Request) {
	service.FixedDelay(5 * time.Second)
	response := model.Message{Message: "This response was delayed by 5 seconds"}
	jsonResponse(w, response, http.StatusOK)
}

func UnstableHandler(w http.ResponseWriter, r *http.Request) {
	if service.RandomFailure(0.3) {
		response := model.Message{Error: "Internal Server Error"}
		jsonResponse(w, response, http.StatusInternalServerError)
		return
	}
	response := model.Message{Message: "Success"}
	jsonResponse(w, response, http.StatusOK)
}

func CustomDelayHandler(w http.ResponseWriter, r *http.Request) {
	delayParam := r.URL.Query().Get("delay")
	if delayParam == "" {
		delayParam = "2"
	}

	delay, err := strconv.Atoi(delayParam)
	if err != nil || delay < 0 {
		http.Error(w, service.ErrInvalidDelay.Error(), http.StatusBadRequest)
		return
	}

	if err := service.CustomDelay(delay); err != nil {
		http.Error(w, service.ErrInvalidDelay.Error(), http.StatusBadRequest)
		return
	}

	response := model.Message{Message: "Response delayed by " + strconv.Itoa(delay) + " seconds"}
	jsonResponse(w, response, http.StatusOK)
}

func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
