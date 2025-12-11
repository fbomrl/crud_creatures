package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func EnsureDBAndMigrate(user, pass, host, dbName, migrationsDir string) error {
	masterDSN := fmt.Sprintf("sqlserver://%s:%s@%s?database=master", user, pass, host)
	db, err := sql.Open("sqlserver", masterDSN)
	if err != nil {
		return fmt.Errorf("open master connection: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createSQL := fmt.Sprintf("IF DB_ID(N'%s') IS NULL CREATE DATABASE [%s];", dbName, dbName)
	if _, err := db.ExecContext(ctx, createSQL); err != nil {
		return fmt.Errorf("create database: %w", err)
	}

	dbURL := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", user, pass, host, dbName)
	m, err := migrate.New("file://"+migrationsDir, dbURL)
	if err != nil {
		return fmt.Errorf("migrate new: %w", err)
	}
	defer func() {
		_, _ = m.Close()
	}()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrate up: %w", err)
	}

	return nil
}
