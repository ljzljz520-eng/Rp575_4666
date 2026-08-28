package store

import (
	"encoding/json"
	"errors"
	"go.etcd.io/bbolt"
	"supplierhub/internal/model"
	"sync"
)

var buckets = [][]byte{[]byte("suppliers"), []byte("permissions"), []byte("inbounds"), []byte("qualities"), []byte("settlements"), []byte("audits"), []byte("sessions")}

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, x := tx.CreateBucketIfNotExists(b); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
	}
	return s, e
}
func (s *Store) Close() error { return s.db.Close() }
func (s *Store) Put(bucket, key string, v any) error {
	raw, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), raw) })
}
func (s *Store) Get(bucket, key string, v any) error {
	return s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return errors.New("bucket missing")
		}
		raw := b.Get([]byte(key))
		if raw == nil {
			return errors.New("not found")
		}
		return json.Unmarshal(raw, v)
	})
}
func (s *Store) Delete(bucket, key string) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Delete([]byte(key)) })
}
func (s *Store) List(bucket string) ([][]byte, error) {
	var out [][]byte
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucket)).ForEach(func(k, v []byte) error {
			if v != nil {
				out = append(out, append([]byte(nil), v...))
			}
			return nil
		})
	})
	return out, e
}

var _ = model.Supplier{}
