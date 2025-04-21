package main

import (
	"encoding/json"
	"fmt"
	"os"
	"review-service/config"
	"review-service/internal/db"
	"review-service/internal/filescanner"
	"review-service/internal/logger"
	"review-service/internal/parser"

	//"review-service/internal/s3client"
	"sync"
)

type State struct {
	ProcessedFiles map[string]bool `json:"processed_files"`
}

func ensureProcessedDir() {
	if _, err := os.Stat("processed"); os.IsNotExist(err) {
		_ = os.Mkdir("processed", 0755)
	}
}
func loadState() State {
	ensureProcessedDir()

	data, err := os.ReadFile("processed/state.json")
	if err != nil {
		// file doesn't exist yet — create empty map
		return State{ProcessedFiles: make(map[string]bool)}
	}

	var state State
	err = json.Unmarshal(data, &state)
	if err != nil || state.ProcessedFiles == nil {
		state.ProcessedFiles = make(map[string]bool)
	}
	return state
}

func saveState(state State) {
	ensureProcessedDir()

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		logger.Error("Failed to marshal state: " + err.Error())
		return
	}

	err = os.WriteFile("./processed/state.json", data, 0644)
	if err != nil {
		logger.Error("Failed to write state file: " + err.Error())
	} else {
		logger.Info("State saved to processed/state.json")
	}
}

func main() {
	cfg := config.Load()
	state := loadState()

	/*s3svc, err := s3client.New(cfg.AWSRegion, cfg.S3Bucket, cfg.S3Prefix)
	if err != nil {
		logger.Error("Failed to connect to S3: " + err.Error())
		return
	}*/

	store, err := db.NewStore(cfg.DatabaseURL)
	if err != nil {
		logger.Error("Failed to connect to DB: " + err.Error())
		return
	}

	/*files, err := s3svc.ListJLFiles()
	if err != nil {
		logger.Error("Failed to list S3 files: " + err.Error())
		return
	}*/

	scanner := filescanner.New("./data")
	files, err := scanner.ListJLFiles()
	var wg sync.WaitGroup
	jobs := make(chan string, len(files))

	for i := 0; i < cfg.WorkerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for key := range jobs {
				if state.ProcessedFiles[key] {
					continue
				}

				//data, err := s3svc.DownloadFile(key)
				data, err := scanner.ReadFile(key)
				if err != nil {
					logger.Warn("Failed to download " + key)
					continue
				}

				reviews := parser.ParseJL(data)
				for _, r := range reviews {
					if err := store.InsertReview(r); err != nil {
						logger.Warn(fmt.Sprintf("Failed to insert review: %d", r.HotelReviewID))
					}
				}
				state.ProcessedFiles[key] = true
			}
		}()
	}

	for _, f := range files {
		jobs <- f
	}
	close(jobs)
	wg.Wait()

	saveState(state)
	logger.Info("Processing complete.")
}
