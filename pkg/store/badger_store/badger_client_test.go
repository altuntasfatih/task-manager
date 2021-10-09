package badger_store

import (
	"github.com/altuntasfatih/task-manager/pkg/custom"
	"github.com/altuntasfatih/task-manager/pkg/models"
	"github.com/altuntasfatih/task-manager/pkg/store"
	"github.com/rs/xid"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var badgerStore store.ReaderWriterRemover

func init() {
	var err error
	badgerStore, err = NewClient(true)
	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_Write(t *testing.T) {
	guid := xid.New().String()

	newUser := &models.User{Id: guid, Email: "altuntasfatih42@gmail.com", FirstName: "Fatih", LastName: "Altuntaş"}
	err := badgerStore.CreateUser(guid, newUser)
	require.Nil(t, err)

	actualUser, err := badgerStore.GetUser(guid)

	require.Nil(t, err)
	require.Equal(t, newUser, actualUser)
}

func TestClient_Get(t *testing.T) {
	guid := xid.New().String()
	_, err := badgerStore.GetUser(guid)

	require.Equal(t, err, custom.ErrUserNotFound)
}

func TestClient_Remove(t *testing.T) {
	guid := xid.New().String()
	newUser := &models.User{Id: guid, Email: "altuntasfatih42@gmail.com", FirstName: "Fatih", LastName: "Altuntaş"}
	err := badgerStore.CreateUser(guid, newUser)
	require.Nil(t, err)

	err = badgerStore.DeleteUser(guid)
	require.Nil(t, err)

	_, err = badgerStore.GetUser(guid)

	require.Equal(t, err, custom.ErrUserNotFound)
}
