package server

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/hitesharora1997/hitcounter/internal/counter"
	"github.com/hitesharora1997/hitcounter/pkg/persistence"
)

type Server struct {
	sync.Mutex
	dataFile    string
	counter     *counter.RequestCounter
	stopChannel chan struct{} // Channel to signal stopping of periodic saving
}

func NewServer(dataFile string, timeProvider func() time.Time) *Server {
	s := &Server{
		dataFile:    dataFile,
		counter:     counter.NewRequestCounter(timeProvider),
		stopChannel: make(chan struct{}),
	}
	if err := persistence.Restore(s.dataFile, s.counter); err != nil {
		log.Printf("Error restoring data: %v\n", err)
	} else {
		log.Println("Data successfully restored")
	}
	go s.persistData()
	return s
}

func (s *Server) persistData() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-s.stopChannel:
			persistence.SaveData(s.dataFile, s.counter) // Final save before exiting
			return
		case <-ticker.C:
			persistence.SaveData(s.dataFile, s.counter)
		}
	}
}

func (s *Server) Stop() {
	close(s.stopChannel)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	count := s.counter.RecordAndCount()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Total request in the last 60 seconds: " + strconv.Itoa(count)))
}
