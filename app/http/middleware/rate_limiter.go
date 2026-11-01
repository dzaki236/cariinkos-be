package middleware

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/http/middleware"
)

type RateLimiter struct {
	throttle http.Middleware
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		throttle: middleware.Throttle("global"),
	}
}
func (m *RateLimiter) Signature() string {
	return "RateLimiter"
}

func (m *RateLimiter) Handle(ctx http.Context) {
	m.throttle.Handle(ctx)
}
