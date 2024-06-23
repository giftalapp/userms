package middleware

/*
We will use x/time/rate Go package which provides
a token bucket rate-limiter algorithm. rate#Limiter
controls how frequently events are allowed to happen.
It implements a “token bucket” of size “b”,
initially full and refilled at rate “r” tokens per second.
Informally, in any large enough time interval,
the Limiter limits the rate to r tokens per second,
with a maximum burst size of “b” events.
*/

import (
	"sync"

	"golang.org/x/time/rate"
)

type RateLimiterBucket struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

// Creates a new bucket that only accepts up to b requests every 1/r seconds per ip
func NewRateLimitterBucket(r rate.Limit, b int) *RateLimiterBucket {
	i := &RateLimiterBucket{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}

	return i
}

// AddIP creates a new rate limiter and adds it to the ips map,
// using the IP address as the key
func (i *RateLimiterBucket) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)

	i.ips[ip] = limiter

	return limiter
}

// GetLimiter returns the rate limiter for the provided IP address if it exists.
// Otherwise calls AddIP to add IP address to the map
func (i *RateLimiterBucket) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]

	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}

	i.mu.Unlock()

	return limiter
}
