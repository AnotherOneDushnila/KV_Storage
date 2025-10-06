package store

type Store interface {
	Put(collection, key string, value []byte) error
	Get(collection, key string) ([]byte, error)
	Delete(collection, key string) error
}