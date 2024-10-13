package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"reverse-proxy/internal/core/ports"
	"sync"
)

var postgresDB *sql.DB

var (
	_    ports.IRepo = (*PostgreSQL)(nil)
	once sync.Once
)

type PostgreSQL struct{}

func NewPostgreSQL() ports.IRepo {
	return &PostgreSQL{}
}

func (p *PostgreSQL) StartConnection() error {
	var err error

	dsn := fmt.Sprintf("host=db port=5432 user=postgres password=postgres dbname=reverse_proxy sslmode=disable")

	once.Do(func() {
		postgresDB, err = sql.Open("pgx", dsn)
		if err != nil {
			log.Fatalf("error to connect with database: %v", err)
		}
	})

	return err
}

func (p *PostgreSQL) GetAllBlockedIPs() ([]string, error) {
	var ips []string
	rows, err := postgresDB.Query(`SELECT ip FROM blocked_ips`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ip string
		if err := rows.Scan(&ip); err != nil {
			return nil, err
		}
		ips = append(ips, ip)
	}

	return ips, nil
}
