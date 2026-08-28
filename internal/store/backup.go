package store

import (
	"errors"
	"go.etcd.io/bbolt"
)

func (s *Store) Healthy() error {
	if s == nil || s.db == nil {
		return errors.New("closed")
	}
	return s.db.View(func(*bbolt.Tx) error { return nil })
}
