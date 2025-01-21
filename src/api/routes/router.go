package routes

import "net/http"

func InitializeRoutes() {
	http.HandleFunc("/api/healthcheck", HealthCheck)
}
