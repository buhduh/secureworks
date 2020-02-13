package sqllite

import (
	"secureworks/constants"
	"testing"
)

const (
	DUMMY_DB string = "/root/dummydb"
)

func TestCreateDB(t *testing.T) {
	err := CreateDatabase(constants.CREATE_DB_SQL, DUMMY_DB)
	if err != nil {
		t.Fatalf(
			"Should not have failed creating database, got error '%s'.",
			err,
		)
	}
}
