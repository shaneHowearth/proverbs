// Package postgresstore -
package postgresstore

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jackc/pgx"
)

type pg struct {
	dbConn *pgx.ConnPool
}

// NewPGStore -
// nolint:golint
func NewPGStore(host, database, user, password string) (*pg, error) {
	pgxConfig := pgx.ConnConfig{
		Host:     host,
		Database: database,
		User:     user,
		Password: password,
	}
	pgxConnPoolConfig := pgx.ConnPoolConfig{ConnConfig: pgxConfig, MaxConnections: 3, AfterConnect: nil, AcquireTimeout: 0}
	conn, err := pgx.NewConnPool(pgxConnPoolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create db connection with error %w", err)
	}
	p := pg{
		dbConn: conn,
	}
	return &p, nil
}

// GetProverb -
func (pg *pg) GetRandomProverb() (proverb, translation, explanation string, err error) {
	rowCount, err := pg.getRowCount("proverb")
	if err != nil {
		return "", "", "", fmt.Errorf("unable to get rowcount for proverb table with error %w", err)
	}

	// select a row
	id := pg.getRandomID(rowCount)
	err = pg.dbConn.QueryRow("select maori_name, translation, explanation from proverb where id = $1", id).Scan(&proverb, &translation, &explanation)
	if err != nil {
		return "", "", "", fmt.Errorf("unable to query proverb with error %w", err)
	}

	return proverb, translation, explanation, nil
}

// getRowCount -
func (pg *pg) getRowCount(table string) (int64, error) {
	var n int64
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s", pgx.Identifier.Sanitize([]string{table}))
	err := pg.dbConn.QueryRow(query).Scan(&n)
	return n, err
}

// getRandomID -
func (pg *pg) getRandomID(rowcount int64) int64 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Int63n(rowcount + 1)
}

// GetPlacename -
func (pg *pg) GetRandomPlacename() (placename, translation, explanation string, err error) {
	rowCount, err := pg.getRowCount("placename")
	if err != nil {
		return "", "", "", fmt.Errorf("unable to get rowcount for placename table with error %w", err)
	}

	// select a row
	id := pg.getRandomID(rowCount)
	err = pg.dbConn.QueryRow("select maori_name, translation, explanation from placename where id = $1", id).Scan(&placename, &translation, &explanation)
	if err != nil {
		return "", "", "", fmt.Errorf("unable to query placename with error %w", err)
	}

	return placename, translation, explanation, nil
}
