package ports

import (
	"time"
)

type CacheStore interface {
	Get(string) (any, bool)
	Set(key string, value any) error
	SetWithTTL(key string, value any, duration time.Duration) error
	Delete(key string) error
	Flush() error
}
