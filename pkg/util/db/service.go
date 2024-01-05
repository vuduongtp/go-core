package dbutil

import (
	"github.com/imdatngo/gowhere"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// New creates new database connection to the database server
func New(dialect, dbPsn string, cfg *gorm.Config) (db *gorm.DB, err error) {
	switch dialect {
	case "mysql":
		gowhere.DefaultConfig.Dialect = gowhere.DialectMySQL
		db, err = gorm.Open(mysql.Open(dbPsn), cfg)
		if err != nil {
			return nil, err
		}
	case "postgres":
		gowhere.DefaultConfig.Dialect = gowhere.DialectPostgreSQL
		db, err = gorm.Open(postgres.Open(dbPsn), cfg)
		if err != nil {
			return nil, err
		}
	case "sqlite3":
		gowhere.DefaultConfig.Dialect = gowhere.DialectMySQL
		db, err = gorm.Open(sqlite.Open(dbPsn), cfg)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
