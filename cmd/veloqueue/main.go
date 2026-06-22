package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env %s", err)
	}
	host := os.Getenv("HOST")
	password := os.Getenv("PASSWORD")
	user := os.Getenv("USER")
	database := os.Getenv("DATABASE")
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s", user, password, host, database)
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatalf("Error opening sql. Error: %s", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to database. Ping: %s", err)
	} else {
		log.Println("Database connected Successfully")
	}

	query := "SELECT * FROM jobs"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error fetching rows %s", err)
	} else {
		cols, err := rows.Columns()
		if err != nil {
			log.Fatalf("No columns %s", err)
		}
		lenCols := len(cols)
		var jobsTable [][]any
		for rows.Next() {
			rowValues := make([]any, lenCols)
			scanArgs := make([]interface{}, lenCols)
			for i := range rowValues {
				scanArgs[i] = &rowValues[i]
			}
			if err := rows.Scan(scanArgs...); err != nil {
				log.Fatalf("Unable to extract column values %s", err)
			}

			jobsTable = append(jobsTable, rowValues)
		}
		rows.Close()
		w := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)
		for _, row := range jobsTable {
			for _, val := range row {
				switch v := val.(type) {
				case nil:
					fmt.Fprintf(w, "NULL\t")
				case []byte:
					fmt.Fprintf(w, "s\t")
				case string:
					fmt.Fprintf(w, "%s\t", v)
				case int64:
					fmt.Fprintf(w, "%d\t", v)
				case time.Time:
					fmt.Fprintf(w, "%s\t", v.Format("2006-01-02 15:04:05"))
				default:
					fmt.Fprintf(w, "%v\t", v)
				}
			}
			fmt.Fprintln(w)
		}
		w.Flush()
	}

}
