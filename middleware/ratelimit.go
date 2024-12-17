package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/barealek/chatapp/types"
)

type OptFunction func(*RateLimitOptions)

type RateLimitOptions struct {
	MaxRequests int
	Window      time.Duration
}

func defaultOptions() *RateLimitOptions {
	return &RateLimitOptions{
		MaxRequests: 5,
		Window:      1 * time.Second,
	}
}

func WithMaxRequests(max int) OptFunction {
	return func(o *RateLimitOptions) {
		o.MaxRequests = max
	}
}

func WithWindow(window time.Duration) OptFunction {
	return func(o *RateLimitOptions) {
		o.Window = window
	}
}

type visitor struct {
	sync.RWMutex
	lastAccess time.Time
	requests   int
}

// HAS TO BE AFTER AUTH MIDDLEWARE
func NewRateLimiter(opts ...OptFunction) func(http.Handler) http.Handler {
	options := defaultOptions()
	for _, opt := range opts {
		opt(options)
	}

	var visitors = make(map[string]*visitor)
	var mtx sync.Mutex

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u := r.Context().Value(ContextKeyUser).(types.User)

			mtx.Lock()
			v, exists := visitors[u.ID]
			if !exists {
				v = &visitor{
					lastAccess: time.Now(),
					requests:   1,
				}
				visitors[u.ID] = v
				mtx.Unlock()
			} else {
				mtx.Unlock()
				v.Lock()
				if time.Since(v.lastAccess) > options.Window {
					v.requests = 0
					v.lastAccess = time.Now()
				}
				v.requests++
				v.Unlock()
			}

			if v.requests > options.MaxRequests {
				http.Error(w, "too many requests", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
