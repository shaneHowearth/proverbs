// Package storage -
package storage

// Store -
type Store interface {
	GetProverb() (proverb string, translation string, err error)
	GetPlacename() (proverb string, translation string, err error)
}
