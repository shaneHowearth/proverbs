// Package postgresstore -
package postgresstore

import (
	"fmt"
	"math/rand"
	"time"

	"database/sql"

	_ "github.com/lib/pq" //nolint:golint
)

type pg struct {
	dbConn *sql.DB
}

// NewPGStore -
// nolint:golint
func NewPGStore(connStr string) (*pg, error) {
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse DATABASE_URL error %v", err)
	}

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
	rowCount, err := pg.getRowCount("proverbs")
	if err != nil {
		return "", "", "", fmt.Errorf("unable to get rowcount for proverbs table with error %w", err)
	}

	// select a row
	id := pg.getRandomID(rowCount)
	var tmpProverb, tmpTranslation, tmpExplanation sql.NullString
	err = pg.dbConn.QueryRow("select maori_name, translation, explanation from proverbs where id = $1", id).Scan(&tmpProverb, &tmpTranslation, &tmpExplanation)
	if err != nil {
		return "", "", "", fmt.Errorf("unable to query proverb with error %w", err)
	}

	return tmpProverb.String, tmpTranslation.String, tmpExplanation.String, nil
}

// getRowCount -
func (pg *pg) getRowCount(table string) (int64, error) {
	var n int64
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s", table)
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
	rowCount, err := pg.getRowCount("placenames")
	if err != nil {
		return "", "", "", fmt.Errorf("unable to get rowcount for placename table with error %w", err)
	}

	// select a row
	id := pg.getRandomID(rowCount)
	var tmpPlacename, tmpTranslation, tmpExplanation sql.NullString
	err = pg.dbConn.QueryRow("select maori_name, translation, explanation from placenames where id = $1", id).Scan(&tmpPlacename, &tmpTranslation, &tmpExplanation)
	if err != nil {
		return "", "", "", fmt.Errorf("unable to query placename with error %w", err)
	}

	return tmpPlacename.String, tmpTranslation.String, tmpExplanation.String, nil
}
