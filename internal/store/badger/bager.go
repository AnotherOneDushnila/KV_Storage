package badger


import (
	"log"
	"fmt"
	"github.com/dgraph-io/badger/v4"
)


type BadgerStore struct {
	db   *badger.DB
	path string
}


func New(path string) *BadgerStore {
	db, err := badger.Open(badger.DefaultOptions(path))

	if err != nil {
		log.Fatal(err)
	}

	return &BadgerStore{db: db, path: path}
}


func (b *BadgerStore) Close() error {
    return b.db.Close()
}

func (b *BadgerStore) Exists(collection, key string) bool {
	_, err := b.Get(collection, key)

	if err != nil {
		return false
	}

	return true
}

func (b *BadgerStore) Put(collection, key string, value []byte) error {
	if b.Exists(collection, key) {
		return fmt.Errorf("key already exists")
	}

	bytesKey := []byte(collection + "/" + key)

	err := b.db.Update((func(txn *badger.Txn) error {
		return txn.Set(bytesKey, value)
	}))

	if err != nil {
		return err
	}
	return nil
}


func (b *BadgerStore) Get(collection, key string) ([]byte, error) {
	bytesKey := []byte(collection + "/" + key)

	var res []byte
    err := b.db.View(func(txn *badger.Txn) error {
        item, err := txn.Get(bytesKey)
        if err != nil {
            return err
        }
        res, err = item.ValueCopy(nil)
        return err
    })

    if err != nil {
        return nil, err
    }

    return res, nil
}


func (b *BadgerStore) Delete(collection, key string) error {
	bytesKey := []byte(collection + "/" + key)

	err := b.db.Update((func(txn *badger.Txn) error {
		return txn.Delete(bytesKey)
	}))

	if err != nil {
		return err
	}

	return nil
}