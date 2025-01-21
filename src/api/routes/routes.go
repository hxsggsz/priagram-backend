package routes

import (
	"encoding/json"
	"net/http"
	"priagram/src/api/dtos"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	healthCHeck := dtos.NewhealthCheck(http.StatusOK, "application working correctly")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(healthCHeck)
}
