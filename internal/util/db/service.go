package dbutil

import (
	dbutil "github.com/vuduongtp/go-core/pkg/util/db"

	"github.com/imdatngo/gowhere"
	_ "gorm.io/driver/postgres" // DB adapter
	"gorm.io/gorm"
)

// New creates new database connection to the database server
func New(dbPsn string, enableLog bool) (*gorm.DB, error) {
	gowhere.DefaultConfig.Dialect = gowhere.DialectMySQL
	config := new(gorm.Config)
	return dbutil.New("postgres", dbPsn, config)
}

// NewDB creates new DB instance
func NewDB(model interface{}) *dbutil.DB {
	return &dbutil.DB{Model: model}
}
