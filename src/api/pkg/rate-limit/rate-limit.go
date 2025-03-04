package ratelimit

import (
	"fmt"
	"net/http"
	"priagram/src/api/utils"

	"golang.org/x/time/rate"
)

type nextFunc func(w http.ResponseWriter, r *http.Request)

var limiter = rate.NewLimiter(10, 1) // 10 requests per minute, burst 1

func RateLimitMiddleware(next nextFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("max request capacity reached, wait and try again later"))
			return
		}
		next(w, r)
	})
}
