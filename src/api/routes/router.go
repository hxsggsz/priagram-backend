package routes

import (
	"net/http"
	ratelimit "priagram/src/api/pkg/rate-limit"
)

func InitializeRoutes() {
	http.HandleFunc("/api/healthcheck", HealthCheck)
	http.HandleFunc("/api/prisma", ratelimit.RateLimitMiddleware(PrismaToDiagram))
}
