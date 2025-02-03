package stats

import (
	"log"
	"os"
	"path/filepath"
)

func Init() {
	statsDirectory := os.Getenv("STATS_DIRECTORY")
	statsFile := os.Getenv("STATS_FILE")
	statsPath := filepath.Join(statsDirectory, statsFile)

	if _, err := os.Stat(statsPath); os.IsNotExist(err) {
		log.Println("📂 Stats directory not found. Creating directory...")
		if err := os.MkdirAll(statsDirectory, 0755); err != nil {
			log.Fatalf("Failed to create directory/ directory: %v", err)
		}
	}
}
