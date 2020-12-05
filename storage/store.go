// Package storage -
package storage

// Store -
type Store interface {
	GetRandomProverb() (proverb string, translation string, err error)
	GetRandomPlacename() (proverb string, translation string, err error)
}
