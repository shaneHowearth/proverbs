// Package storage -
package storage

import (
	"strconv"
	"time"
)

// Store -
type Store interface {
	GetRandomProverb() (proverb, translation, explanation string, err error)
	GetRandomPlacename() (placename, translation, explanation string, err error)
}

func isMorning() bool {
	now := time.Now()
	newLayout := "15:04"
	check, _ := time.Parse(newLayout, strconv.Itoa(now.Hour())+":"+strconv.Itoa(now.Minute()))
	start, _ := time.Parse(newLayout, "23:59")
	end, _ := time.Parse(newLayout, "11:59")
	start, end = start.UTC(), end.UTC()
	if start.After(end) {
		start, end = end, start

	}
	check = check.Local()
	return !check.Before(start) && !check.After(end)

}

// GetContent -
func GetContent(s Store) (content, translation, explanation string, err error) {
	if isMorning() {
		return s.GetRandomProverb()
	}
	return s.GetRandomPlacename()
}
