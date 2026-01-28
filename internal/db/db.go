package db

import (
    "fmt"
    "log"
    "os"
    "time"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Connect() *gorm.DB {
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname,
    )

    var db *gorm.DB
    var err error

    // Retry loop
    for i := 0; i < 10; i++ {
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err == nil {
            log.Println("[info] database connected")
            return db
        }
        log.Println("[warn] database not ready, retrying in 2s...")
        time.Sleep(2 * time.Second)
    }

    log.Fatalf("[error] failed to connect to database after retries: %v", err)
    return nil
}