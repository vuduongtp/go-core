package country

import (
	"context"
	"net/http"

	"github.com/vuduongtp/go-core/internal/model"
	"github.com/vuduongtp/go-core/pkg/rbac"
	"github.com/vuduongtp/go-core/pkg/server"
	dbutil "github.com/vuduongtp/go-core/pkg/util/db"
	structutil "github.com/vuduongtp/go-core/pkg/util/struct"
)

// Custom errors
var (
	ErrCountryNotFound    = server.NewHTTPError(http.StatusBadRequest, "COUNTRY_NOTFOUND", "Country not found")
	ErrCountryNameExisted = server.NewHTTPValidationError("Country name already exists")
)

// Create creates a new country
func (s *Country) Create(ctx context.Context, authUsr *model.AuthUser, data CreationData) (*model.Country, error) {
	if err := s.enforce(authUsr, model.ActionCreateAll); err != nil {
		return nil, err
	}

	if existed, err := s.cdb.Exist(ctx, s.db, map[string]interface{}{"name": data.Name}); err != nil || existed {
		return nil, ErrCountryNameExisted.SetInternal(err)
	}

	rec := &model.Country{
		Name:      data.Name,
		Code:      data.Code,
		PhoneCode: data.PhoneCode,
	}
	if err := s.cdb.Create(ctx, s.db, rec); err != nil {
		return nil, server.NewHTTPInternalError("Error creating country").SetInternal(err)
	}

	return rec, nil
}

// View returns single country
func (s *Country) View(ctx context.Context, authUsr *model.AuthUser, id int) (*model.Country, error) {
	if err := s.enforce(authUsr, model.ActionViewAll); err != nil {
		return nil, err
	}

	rec := new(model.Country)
	if err := s.cdb.View(ctx, s.db, rec, id); err != nil {
		return nil, ErrCountryNotFound.SetInternal(err)
	}

	return rec, nil
}

// List returns list of countrys
func (s *Country) List(ctx context.Context, authUsr *model.AuthUser, lq *dbutil.ListQueryCondition, count *int64) ([]*model.Country, error) {
	if err := s.enforce(authUsr, model.ActionViewAll); err != nil {
		return nil, err
	}

	var data []*model.Country
	if err := s.cdb.List(ctx, s.db, &data, lq, count); err != nil {
		return nil, server.NewHTTPInternalError("Error listing country").SetInternal(err)
	}

	return data, nil
}

// Update updates country information
func (s *Country) Update(ctx context.Context, authUsr *model.AuthUser, id int, data UpdateData) (*model.Country, error) {
	if err := s.enforce(authUsr, model.ActionUpdateAll); err != nil {
		return nil, err
	}

	if existed, err := s.cdb.Exist(ctx, s.db, map[string]interface{}{"name": data.Name, "id__notexact": id}); err != nil || existed {
		return nil, ErrCountryNameExisted.SetInternal(err)
	}

	// optimistic update
	updates := structutil.ToMap(data)
	if err := s.cdb.Update(ctx, s.db, updates, id); err != nil {
		return nil, server.NewHTTPInternalError("Error updating country").SetInternal(err)
	}

	rec := new(model.Country)
	if err := s.cdb.View(ctx, s.db, rec, id); err != nil {
		return nil, ErrCountryNotFound.SetInternal(err)
	}

	return rec, nil
}

// Delete deletes a country
func (s *Country) Delete(ctx context.Context, authUsr *model.AuthUser, id int) error {
	if err := s.enforce(authUsr, model.ActionDeleteAll); err != nil {
		return err
	}

	if existed, err := s.cdb.Exist(ctx, s.db, id); err != nil || !existed {
		return ErrCountryNotFound.SetInternal(err)
	}

	if err := s.cdb.Delete(ctx, s.db, id); err != nil {
		return server.NewHTTPInternalError("Error deleting country").SetInternal(err)
	}

	return nil
}

// enforce checks user permission to perform the action
func (s *Country) enforce(authUsr *model.AuthUser, action string) error {
	if !s.rbac.Enforce(authUsr.Role, model.ObjectCountry, action) {
		return rbac.ErrForbiddenAction
	}
	return nil
}
