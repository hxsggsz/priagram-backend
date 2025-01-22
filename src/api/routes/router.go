package routes

import "net/http"

func InitializeRoutes() {
	http.HandleFunc("/api/healthcheck", HealthCheck)
	http.HandleFunc("/api/prisma", PrismaToDiagram)
}
