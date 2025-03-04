package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"priagram/src/api/config/headers"
	"priagram/src/api/dtos"
	"priagram/src/api/utils"
	"priagram/src/pkg/lexer"
	"priagram/src/pkg/lexer/formatter"
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
	headers.SetupHeaders(&w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to read request body"))
		return
	}

	defer r.Body.Close()

	var prismaRequest PrismaRequest

	err = json.Unmarshal(body, &prismaRequest)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to read json file"))
		return
	}

	if prismaRequest.Source == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing `source` in json"))
		return
	}

	tokens, err := lexer.Tokenize(prismaRequest.Source)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	formatedData, relations := lexer.Format(tokens)

	utils.WriteJSON(w, http.StatusOK, formatter.NewDiagram(formatedData, relations))
}
