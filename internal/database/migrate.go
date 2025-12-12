package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func EnsureDBAndMigrate(user, pass, host, instance, dbName, migrationsDir string) error {
	if strings.Contains(host, "\\") {
		parts := strings.SplitN(host, "\\", 2)
		host = parts[0]
		instance = parts[1]
	}

	// --- MONTAGEM DO DSN ---
	var masterDSN string
	if instance != "" {
		masterDSN = fmt.Sprintf(
			"sqlserver://%s:%s@%s?database=master&instance=%s&encrypt=disable",
			user, pass, host, instance,
		)
	} else {
		masterDSN = fmt.Sprintf(
			"sqlserver://%s:%s@%s?database=master&encrypt=disable",
			user, pass, host,
		)
	}

	db, err := sql.Open("sqlserver", masterDSN)
	if err != nil {
		return fmt.Errorf("open master connection: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Cria database se não existir
	createSQL := fmt.Sprintf("IF DB_ID(N'%s') IS NULL CREATE DATABASE [%s];", dbName, dbName)
	if _, err := db.ExecContext(ctx, createSQL); err != nil {
		return fmt.Errorf("create database: %w", err)
	}

	// DSN final
	var dbURL string
	if instance != "" {
		dbURL = fmt.Sprintf(
			"sqlserver://%s:%s@%s?database=%s&instance=%s&encrypt=disable",
			user, pass, host, dbName, instance,
		)
	} else {
		dbURL = fmt.Sprintf(
			"sqlserver://%s:%s@%s?database=%s&encrypt=disable",
			user, pass, host, dbName,
		)
	}

	m, err := migrate.New("file://"+migrationsDir, dbURL)
	if err != nil {
		return fmt.Errorf("migrate new: %w", err)
	}
	defer func() { _, _ = m.Close() }()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrate up: %w", err)
	}

	return nil
}
