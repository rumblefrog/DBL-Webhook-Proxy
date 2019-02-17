package schema

import "time"

type Votes struct {
	ID            uint64
	Attempts      uint8
	NextRun       uint64
	Completed     bool
	Endpoint      string
	HMAC          bool
	Authorization string
	Payload       string
	CreatedAt     time.Time
}
