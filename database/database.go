package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDB() (*pgx.Conn, error) {
	connStr := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func createTable(conn *pgx.Conn) error {
	_, err := conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS playing_with_neon(id SERIAL PRIMARY KEY, name TEXT NOT NULL, value REAL);")
	return err
}

// insertData populates the playing_with_neon table with sample data.
func insertData(conn *pgx.Conn) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO playing_with_neon(name, value) SELECT LEFT(md5(i::TEXT), 10), random() FROM generate_series(1, 10) s(i);")
	return err
}

// fetchAndPrintData retrieves and prints the data from the playing_with_neon table.
func fetchAndPrintData(conn *pgx.Conn) error {
	rows, err := conn.Query(context.Background(), "SELECT * FROM playing_with_neon")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int32
		var name string
		var value float32
		if err := rows.Scan(&id, &name, &value); err != nil {
			return err
		}
		fmt.Printf("%d | %s | %f\n", id, name, value)
	}
	return nil
}
