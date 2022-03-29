package controller

import (
	"github.com/beefsack/go-rate"
)

var rateLimiter *rate.RateLimiter

func init() {
	rateLimiter = rate.New(100, 1000)
}
