package persistence_test

import (
	"encoding/json"
	"github.com/hitesharora1997/hitcounter/internal/counter"
	"github.com/hitesharora1997/hitcounter/pkg/persistence"
	"os"
	"reflect"
	"testing"
)

func TestSaveData(t *testing.T) {
	requestCounter := &counter.RequestCounter{
		RequestTimes: []int64{123456789, 987654321},
	}

	tempFile, err := os.CreateTemp("", "requestCounter")
	if err != nil {
		t.Fatalf("Unable to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	persistence.SaveData(tempFile.Name(), requestCounter)

	data, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Unable to read temp file: %v", err)
	}

	var savedData struct {
		RequestTimes []int64 `json:"requestTimes"`
	}
	if err := json.Unmarshal(data, &savedData); err != nil {
		t.Fatalf("Error unmarshalling data: %v", err)
	}

	if !reflect.DeepEqual(savedData.RequestTimes, requestCounter.RequestTimes) {
		t.Errorf("Saved data does not match counter state")
	}
}

func TestRestore(t *testing.T) {
	mockData := struct {
		RequestTimes []int64 `json:"requestTimes"`
	}{
		RequestTimes: []int64{123456789, 987654321},
	}
	mockDataBytes, _ := json.Marshal(mockData)

	tempFile, err := os.CreateTemp("", "requestCounter")
	if err != nil {
		t.Fatalf("Unable to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	if _, err := tempFile.Write(mockDataBytes); err != nil {
		t.Fatalf("Unable to write to temp file: %v", err)
	}

	requestCounter := &counter.RequestCounter{}
	if err := persistence.Restore(tempFile.Name(), requestCounter); err != nil {
		t.Fatalf("Restore failed: %v", err)
	}

	if !reflect.DeepEqual(requestCounter.RequestTimes, mockData.RequestTimes) {
		t.Errorf("Restored data does not match mock data")
	}
}
