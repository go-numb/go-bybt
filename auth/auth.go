package auth

import "sync"

type Config struct {
	mux sync.RWMutex

	Key string
}

func New(key string) *Config {
	if len(key) == 0 {
		return nil
	}

	return &Config{
		Key: key,
	}
}
