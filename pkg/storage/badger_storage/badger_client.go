package badger_storage

import (
	"bytes"
	"encoding/gob"
	"github.com/altuntasfatih/task-manager/pkg/custom"
	"github.com/altuntasfatih/task-manager/pkg/models"
	"github.com/altuntasfatih/task-manager/pkg/storage"
	"github.com/dgraph-io/badger/v3"
	"log"
)

var dbPath = "./storage/"

type client struct {
	db *badger.DB
}

func NewClient(inMemory bool) (storage.ReaderWriterRemover, error) {
	var opt badger.Options
	if inMemory {
		opt = badger.DefaultOptions("").WithInMemory(inMemory)
	} else {
		opt = badger.DefaultOptions(dbPath)
	}
	db, err := badger.Open(opt)
	if err != nil {
		return nil, err
	}
	return &client{db: db}, nil
}

func (c *client) GetUser(id string) (*models.User, error) {
	var user models.User
	return &user, c.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id))

		if err != nil {
			if err == badger.ErrKeyNotFound {
				return custom.ErrUserNotFound
			}
			return err
		}

		return parseItem(item, &user)
	})
}
func (c *client) GetAllUsers() ([]*models.User, error) {
	users := make([]*models.User, 0)
	return users, c.db.View(func(txn *badger.Txn) error {

		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			var user models.User
			if err := parseItem(item, &user); err == nil {
				users = append(users, &user)
			}
		}
		return nil
	})
}

func (c *client) CreateUser(id string, value *models.User) error {

	var val bytes.Buffer
	e := gob.NewEncoder(&val)
	if err := e.Encode(value); err != nil {
		panic(err)
	}
	return c.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(id), val.Bytes())
		return err
	})
}

func (c *client) UpdateUser(id string, value *models.User) error {
	//todo come back later
	return c.CreateUser(id, value)
}

func (c *client) DeleteUser(id string) error {
	return c.db.Update(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(id))
		if err != nil {
			return custom.ErrUserNotFound
		}
		return txn.Delete([]byte(id))
	})
}

func parseItem(item *badger.Item, value *models.User) error {

	val, err := item.ValueCopy(nil)
	if err != nil {
		return err
	}
	d := gob.NewDecoder(bytes.NewReader(val))
	if err := d.Decode(value); err != nil {
		return err
	}
	log.Printf("Decoded Struct from badger_storage :  [%v] \n", value)
	return nil
}
