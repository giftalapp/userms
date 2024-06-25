package middleware

import (
	"net"
	"net/http"
	"strings"

	"github.com/giftalapp/userms/config"
	"golang.org/x/time/rate"
)

// https://stackoverflow.com/a/37897238
func GetIP(r *http.Request) string {

	remoteIP := ""
	// the default is the originating ip. but we try to find better options because this is almost
	// never the right IP
	if parts := strings.Split(r.RemoteAddr, ":"); len(parts) == 2 {
		remoteIP = parts[0]
	}
	// If we have a forwarded-for header, take the address from there
	if xff := strings.Trim(r.Header.Get("X-Forwarded-For"), ","); len(xff) > 0 {
		addrs := strings.Split(xff, ",")
		lastFwd := addrs[len(addrs)-1]
		if ip := net.ParseIP(lastFwd); ip != nil {
			remoteIP = ip.String()
		}
		// parse X-Real-Ip header
	} else if xri := r.Header.Get("X-Real-Ip"); len(xri) > 0 {
		if ip := net.ParseIP(xri); ip != nil {
			remoteIP = ip.String()
		}
	}

	return remoteIP

}

type RateLimitter struct {
	handler http.Handler
	bucket  *RateLimiterBucket
}

func NewRateLimitter(handler http.Handler) *RateLimitter {
	bucket := NewRateLimitterBucket(
		rate.Limit(config.Env.RateLimitInversedSeconds),
		config.Env.RateLimitRequests,
	)

	return &RateLimitter{
		handler: handler,
		bucket:  bucket,
	}
}

func (rl *RateLimitter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	limiter := rl.bucket.GetLimiter(GetIP(r))

	if !limiter.Allow() {
		http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
		return
	}

	rl.handler.ServeHTTP(w, r)
}
