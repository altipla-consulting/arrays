package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

var (
	integers64Sess   sqlbuilder.Database
	integers64Models db.Collection
)

type integers64Model struct {
	ID  int64      `db:"id,omitempty"`
	Foo Integers64 `db:"foo"`
}

func initIntegers64DB(t *testing.T) {
	cnf := &mysql.ConnectionURL{
		User:     "dev-user",
		Password: "dev-password",
		Host:     "localhost",
		Database: "test",
		Options: map[string]string{
			"charset":   "utf8mb4",
			"collation": "utf8mb4_bin",
			"parseTime": "true",
		},
	}
	var err error
	integers64Sess, err = mysql.Open(cnf)
	require.Nil(t, err)

	_, err = integers64Sess.Exec(`DROP TABLE IF EXISTS arrays_test`)
	require.Nil(t, err)

	_, err = integers64Sess.Exec(`
    CREATE TABLE arrays_test (
      id INT(11) NOT NULL AUTO_INCREMENT,
      foo JSON,

      PRIMARY KEY(id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
  `)
	require.Nil(t, err)

	integers64Models = integers64Sess.Collection("arrays_test")

	require.Nil(t, integers64Models.Truncate())
}

func finishIntegers64DB() {
	integers64Sess.Close()
}

func TestLoadSaveIntegers64(t *testing.T) {
	initIntegers64DB(t)
	defer finishIntegers64DB()

	model := new(integers64Model)
	require.Nil(t, integers64Models.InsertReturning(model))

	require.EqualValues(t, model.ID, 1)

	other := new(integers64Model)
	require.Nil(t, integers64Models.Find(1).One(other))
}

func TestLoadSaveIntegers64WithContent(t *testing.T) {
	initIntegers64DB(t)
	defer finishIntegers64DB()

	model := &integers64Model{
		Foo: []int64{3, 4},
	}
	require.Nil(t, integers64Models.InsertReturning(model))

	other := new(integers64Model)
	require.Nil(t, integers64Models.Find(1).One(other))

	require.Equal(t, other.Foo, Integers64{3, 4})
	require.EqualValues(t, other.Foo, []int64{3, 4})
}
