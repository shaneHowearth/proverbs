// Package publish - provide an interface that needs to be satisfied in order to
// publish to a given platform.
package publish

// Publisher -
type Publisher interface {
	PublishContent(string) error
}
