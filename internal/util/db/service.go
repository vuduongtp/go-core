package dbutil

import (
	dbutil "github.com/vuduongtp/go-core/pkg/util/db"
	"github.com/vuduongtp/go-logadapter"

	_ "gorm.io/driver/postgres" // DB adapter
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// New creates new database connection to the database server
func New(dbType, dbPsn string, enableLog bool) (*gorm.DB, error) {
	config := new(gorm.Config)
	db, err := dbutil.New(dbType, dbPsn, config)
	if err != nil {
		return nil, err
	}

	if enableLog {
		db.Logger = logadapter.NewGormLogger().LogMode(logger.Info)
	} else {
		db.Logger = logadapter.NewGormLogger().LogMode(logger.Silent)
	}

	return db, nil
}

// NewDB creates new DB instance
func NewDB(model interface{}) *dbutil.DB {
	return &dbutil.DB{Model: model}
}
