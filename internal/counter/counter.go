package counter

import (
	"sync"
	"time"
)

type RequestCounter struct {
	sync.Mutex
	TotalRequests int64
	RequestTimes  []int64
	TimeProvider  func() int64
}

func NewRequestCounter(timeProvider func() time.Time) *RequestCounter {
	var tp func() int64
	if timeProvider == nil {
		tp = func() int64 { return time.Now().Unix() }
	} else {
		tp = func() int64 { return timeProvider().Unix() }
	}
	return &RequestCounter{
		TimeProvider: tp,
	}
}

func (rc *RequestCounter) RecordAndCount() int {
	rc.Lock()
	defer rc.Unlock()

	now := rc.TimeProvider()
	rc.TotalRequests++
	rc.RequestTimes = append(rc.RequestTimes, now)

	cutoff := now - 60
	var index int
	for i, t := range rc.RequestTimes {
		if t > cutoff {
			index = i
			break
		}
	}
	rc.RequestTimes = rc.RequestTimes[index:]

	return len(rc.RequestTimes)
}
