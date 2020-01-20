package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"timetracker/internal/handlers"
	"timetracker/internal/models"
)

func main() {
	// Config
	addr := flag.String("addr", ":4000", "HTTP network address")
	mode := flag.String("mode", "debug", "Use 'release' for production")
	dsn := flag.String("dsn", "postgres://postgres@localhost:54320/pentotimetracker?sslmode=disable", "Postgres data source")
	dbMigrationDir := flag.String("db.migration.dir", "", "Directory containing the DB migration files, if not provided migration will be skipped")
	flag.Parse()

	// DB setup
	db := configureDB(dsn, dbMigrationDir)
	recordsStore := models.NewRecordDBStore(db)

	// Web server setup
	gin.SetMode(*mode)
	r := gin.Default()
	// CORS for localhost
	r.Use(options)

	// Routes setup
	r.GET("/records", handlers.GetRecords(recordsStore))
	r.POST("/records", handlers.AddRecord(recordsStore))

	r.Run(*addr)
}
