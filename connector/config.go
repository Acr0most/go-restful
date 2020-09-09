package connector

import "time"

type Config struct {
	MaxRetries int
	IntervalMs time.Duration
}

func (t *Config) Compact() {
	if t.MaxRetries <= 0 {
		t.MaxRetries = 20
	}

	if t.IntervalMs <= 0 {
		t.IntervalMs = 3000
	}
}
