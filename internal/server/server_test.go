package server_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hitesharora1997/hitcounter/internal/server"
)

var mockTime time.Time

func mockTimeProvider() time.Time {
	return mockTime
}

func TestServerPersistenceAndRestart(t *testing.T) {
	dataFile := "./testdata/counter.json"

	if err := os.MkdirAll("./testdata", 0755); err != nil {
		t.Fatalf("Failed to create testdata directory: %v", err)
	}

	mockTime = time.Date(2024, 01, 29, 0, 0, 0, 0, time.UTC)
	srv := server.NewServer(dataFile, mockTimeProvider)
	defer srv.Stop()

	for i := 0; i < 5; i++ {
		req, _ := http.NewRequest("GET", "/", nil)
		srv.ServeHTTP(httptest.NewRecorder(), req)
		mockTime = mockTime.Add(10 * time.Second) // advance time by 10 seconds
	}

	newSrv := server.NewServer(dataFile, mockTimeProvider)
	defer newSrv.Stop()

	mockTime = mockTime.Add(50 * time.Second) // simulate 50 seconds passing

	req, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	newSrv.ServeHTTP(response, req)
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	respStr := string(respBody)
	prefix := "Total request in the last 60 seconds: "
	countStr := strings.TrimPrefix(respStr, prefix)
	count, err := strconv.Atoi(strings.TrimSpace(countStr))
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	if count >= 5 {
		t.Errorf("Expected less than 5 requests after time passage, got %d", count)
	}
}
