package routes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"priagram/src/api/dtos"
	"priagram/src/pkg/lexer"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	healthCheck := dtos.NewhealthCheck(http.StatusOK, "application working correctly")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(healthCheck)
}

type PrismaRequest struct {
	Source string `json:"source"`
}

func PrismaToDiagram(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var prismaRequest PrismaRequest
	err = json.Unmarshal(body, &prismaRequest)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if prismaRequest.Source == "" {
		http.Error(w, "Invalid JSON -> missing source in body", http.StatusBadRequest)
	}

	log.Printf("Received: %+v\n", prismaRequest)

	tokens := lexer.Tokenize(prismaRequest.Source)
	formatedData := lexer.Format(tokens)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(formatedData)
}
