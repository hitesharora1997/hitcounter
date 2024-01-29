package persistence

import (
	"encoding/json"
	"github.com/hitesharora1997/hitcounter/internal/counter"
	"log"
	"os"
)

func SaveData(dataFile string, counter *counter.RequestCounter) {
	counter.Lock()
	defer counter.Unlock()

	savedData := struct {
		RequestTimes []int64 `json:"requestTimes"`
	}{
		RequestTimes: counter.RequestTimes,
	}

	data, err := json.Marshal(savedData)
	if err != nil {
		log.Println("Error marshalling the data:", err)
		return
	}

	if err = os.WriteFile(dataFile, data, 0644); err != nil {
		log.Println("Error writing file:", err)
	}
}

func Restore(dataFile string, counter *counter.RequestCounter) error {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return err
	}

	counter.Lock()
	defer counter.Unlock()

	var savedData struct {
		RequestTimes []int64 `json:"requestTimes"`
	}
	if err := json.Unmarshal(data, &savedData); err != nil {
		return err
	}

	counter.RequestTimes = savedData.RequestTimes
	return nil
}
