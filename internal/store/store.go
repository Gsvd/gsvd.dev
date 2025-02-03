package store

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var (
	instance *sql.DB
)

func Init() *sql.DB {
	sqliteDirectory := os.Getenv("SQLITE_DIRECTORY")
	sqliteFile := os.Getenv("SQLITE_FILE")
	sqlitePath := filepath.Join(sqliteDirectory, sqliteFile)

	if _, err := os.Stat(sqlitePath); os.IsNotExist(err) {
		log.Println("📂 SQLite file not found. Creating database...")
		if err := os.MkdirAll(sqliteDirectory, 0755); err != nil {
			log.Fatalf("Failed to create data/ directory: %v", err)
		}
	}

	var err error
	instance, err = sql.Open("sqlite", sqlitePath)
	if err != nil {
		log.Fatalf("SQLite connection error: %v", err)
	}

	if err = instance.Ping(); err != nil {
		log.Fatalf("SQLite ping error: %v", err)
	}

	createTables(instance)

	log.Println("✅ SQLite initialized")

	return instance
}

func Get() *sql.DB {
	if instance == nil {
		log.Fatal("🔴 SQLite not initialized. Call Init() first.")
	}

	return instance
}

func createTables(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(16) NOT NULL DEFAULT 'Anonymous',
		post_id INTEGER NOT NULL,
		comment TEXT NOT NULL,
		approved BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}
	log.Println("✅ Tables created (or already exist)")
}
