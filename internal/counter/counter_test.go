package counter_test

import (
	"github.com/hitesharora1997/hitcounter/internal/counter"
	"testing"
	"time"
)

var mockTime time.Time

func mockTimeProvider() time.Time {
	return mockTime
}

func TestRequestCounter_RecordAndCount(t *testing.T) {
	mockTime = time.Now()
	counter := counter.NewRequestCounter(mockTimeProvider)

	for i := 0; i < 5; i++ {
		mockTime = mockTime.Add(10 * time.Second) // advance time by 10 seconds
		count := counter.RecordAndCount()
		if count != i+1 {
			t.Errorf("Expected count %d, got %d at time %v", i+1, count, mockTime)
		}
	}

	mockTime = mockTime.Add(61 * time.Second)
	count := counter.RecordAndCount()

	t.Logf("Current mock time: %v", mockTime)
	t.Logf("Request times: %v", counter.RequestTimes)

	if count != 1 {
		t.Errorf("Expected count to reset to 1, got %d", count)
	}

	if counter.TotalRequests != 6 {
		t.Errorf("Expected total requests to be 6, got %d", counter.TotalRequests)
	}
}
