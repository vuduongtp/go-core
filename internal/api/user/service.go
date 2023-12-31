package user

import (
	"context"

	"github.com/vuduongtp/go-core/internal/model"
	"github.com/vuduongtp/go-core/pkg/rbac"
	dbutil "github.com/vuduongtp/go-core/pkg/util/db"

	"gorm.io/gorm"
)

// New creates new user application service
func New(db *gorm.DB, udb MyDB, rbacSvc rbac.Intf, cr Crypter) *User {
	return &User{db: db, udb: udb, rbac: rbacSvc, cr: cr}
}

// User represents user application service
type User struct {
	db   *gorm.DB
	udb  MyDB
	rbac rbac.Intf
	cr   Crypter
}

// MyDB represents user repository interface
type MyDB interface {
	dbutil.Intf
	FindByUsername(context.Context, *gorm.DB, string) (*model.User, error)
}

// Crypter represents security interface
type Crypter interface {
	CompareHashAndPassword(hasedPwd string, rawPwd string) bool
	HashPassword(string) string
}
