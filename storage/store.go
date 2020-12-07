// Package storage -
package storage

// Store -
type Store interface {
	GetRandomProverb() (proverb, translation, explanation string, err error)
	GetRandomPlacename() (proverb, translation, explanation string, err error)
}
