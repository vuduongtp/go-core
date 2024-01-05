package user

import (
	"context"

	"github.com/vuduongtp/go-core/internal/model"
	dbutil "github.com/vuduongtp/go-core/pkg/util/db"

	"gorm.io/gorm"
)

// NewDB returns a new user database instance
func NewDB() *DB {
	return &DB{dbutil.NewDB(model.User{})}
}

// DB represents the client for user table
type DB struct {
	*dbutil.DB
}

// FindByUsername queries for single user by username
func (d *DB) FindByUsername(ctx context.Context, db *gorm.DB, uname string) (*model.User, error) {
	rec := new(model.User)
	if err := d.View(ctx, db, rec, "username = ?", uname); err != nil {
		return nil, err
	}
	return rec, nil
}

// FindByRefreshToken queries for single user by refresh token
func (d *DB) FindByRefreshToken(ctx context.Context, db *gorm.DB, token string) (*model.User, error) {
	rec := new(model.User)
	if err := d.View(ctx, db, rec, "refresh_token = ?", token); err != nil {
		return nil, err
	}
	return rec, nil
}
