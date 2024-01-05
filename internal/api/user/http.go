package user

import (
	"context"
	"net/http"
	"strings"

	"github.com/vuduongtp/go-core/internal/model"
	"github.com/vuduongtp/go-core/pkg/server"
	dbutil "github.com/vuduongtp/go-core/pkg/util/db"
	httputil "github.com/vuduongtp/go-core/pkg/util/http"

	"github.com/labstack/echo/v4"
)

// HTTP represents user http service
type HTTP struct {
	svc  Service
	auth model.Auth
}

// Service represents user application interface
type Service interface {
	Create(context.Context, *model.AuthUser, CreationData) (*model.User, error)
	View(context.Context, *model.AuthUser, int) (*model.User, error)
	List(context.Context, *model.AuthUser, *dbutil.ListQueryCondition, *int64) ([]*model.User, error)
	Update(context.Context, *model.AuthUser, int, UpdateData) (*model.User, error)
	Delete(context.Context, *model.AuthUser, int) error
	Me(context.Context, *model.AuthUser) (*model.User, error)
	ChangePassword(context.Context, *model.AuthUser, PasswordChangeData) error
}

// NewHTTP creates new user http service
func NewHTTP(svc Service, auth model.Auth, eg *echo.Group) {
	h := HTTP{svc, auth}

	eg.POST("", h.create)
	eg.GET("/:id", h.view)
	eg.GET("", h.list)
	eg.PATCH("/:id", h.update)
	eg.DELETE("/:id", h.delete)
	eg.GET("/me", h.me)
	eg.PATCH("/me/password", h.changePassword)
}

// CreationData contains user data from json request
type CreationData struct {
	Username  string `json:"username" validate:"required,min=3"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Mobile    string `json:"mobile" validate:"required,mobile"`
	Role      string `json:"role" validate:"required"`
	Blocked   bool   `json:"blocked"`
}

// UpdateData contains user data from json request
type UpdateData struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Email     *string `json:"email,omitempty" validate:"omitempty,email"`
	Mobile    *string `json:"mobile,omitempty" validate:"omitempty,mobile"`
	Role      *string `json:"role,omitempty"`
	Blocked   *bool   `json:"blocked,omitempty"`
}

// PasswordChangeData contains password change request
type PasswordChangeData struct {
	OldPassword        string `json:"old_password" validate:"required"`
	NewPassword        string `json:"new_password" validate:"required,min=8"`
	NewPasswordConfirm string `json:"new_password_confirm" validate:"required,eqfield=NewPassword"`
}

// ListResp contains list of users and current page number response
type ListResp struct {
	Data       []*model.User `json:"data"`
	TotalCount int64         `json:"total_count"`
}

// @Security		BearerToken
// @Summary		Creates new user
// @Description	The new user
// @Accept			json
// @Produce		json
// @Tags			users
// @ID				usersCreate
// @Param			request		body		user.CreationData	true	"CreationData"
// @Success		200			{object}	model.User
// @Failure		400			{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		401			{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		403			{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		500			{object}	swaggerutil.SwaggErrDetailsResp
// @Router			/v1/users	[post]
func (h *HTTP) create(c echo.Context) error {
	r := CreationData{}
	if err := c.Bind(&r); err != nil {
		return err
	}
	r.Email = strings.TrimSpace(r.Email)
	r.FirstName = strings.TrimSpace(r.FirstName)
	r.LastName = strings.TrimSpace(r.LastName)
	r.Mobile = strings.TrimSpace(strings.Replace(r.Mobile, " ", "", -1))
	r.Role = strings.TrimSpace(r.Role)

	if err := validateRole(&r.Role); err != nil {
		return err
	}

	resp, err := h.svc.Create(c.Request().Context(), h.auth.User(c), r)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

// @Security		BearerToken
// @Summary		Returns a single user
// @Description	Returns a single user
// @Accept			json
// @Produce		json
// @Tags			users
// @ID				usersView
// @Param			id				path		int	true	"User ID"
// @Success		200				{object}	model.User
// @Failure		400				{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		401				{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		403				{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		404				{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		500				{object}	swaggerutil.SwaggErrDetailsResp
// @Router			/v1/users/{id}	[get]
func (h *HTTP) view(c echo.Context) error {
	id, err := httputil.ReqID(c)
	if err != nil {
		return err
	}
	resp, err := h.svc.View(c.Request().Context(), h.auth.User(c), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

// @Security		BearerToken
// @Summary		Get list user
// @Description	Get list user
// @Accept			json
// @Produce		json
// @Tags			users
// @ID				usersList
// @Param			q			query		swagger.ListRequest	false	"QueryListRequest"
// @Success		200			{object}	user.ListResp
// @Failure		400			{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		401			{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		403			{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		500			{object}	swaggerutil.SwaggErrDetailsResp
// @Router			/v1/users	[get]
func (h *HTTP) list(c echo.Context) error {
	lq, err := httputil.ReqListQuery(c)
	if err != nil {
		return err
	}
	var count int64 = 0
	resp, err := h.svc.List(c.Request().Context(), h.auth.User(c), lq, &count)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ListResp{resp, count})
}

// @Security		BearerToken
// @Summary		Updates user information
// @Description	Updates user information
// @Accept			json
// @Produce		json
// @Tags			users
// @ID				usersUpdate
// @Param			id				path		int				true	"User ID"
// @Param			request			body		user.UpdateData	true	"UpdateData"
// @Success		200				{object}	model.User
// @Failure		400				{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		401				{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		403				{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		404				{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		500				{object}	swaggerutil.SwaggErrDetailsResp
// @Router			/v1/users/{id}	[patch]
func (h *HTTP) update(c echo.Context) error {
	id, err := httputil.ReqID(c)
	if err != nil {
		return err
	}
	r := UpdateData{}
	if err := c.Bind(&r); err != nil {
		return err
	}
	r.Email = httputil.TrimSpacePointer(r.Email)
	r.FirstName = httputil.TrimSpacePointer(r.FirstName)
	r.LastName = httputil.TrimSpacePointer(r.LastName)
	r.Mobile = httputil.RemoveSpacePointer(r.Mobile)
	r.Role = httputil.RemoveSpacePointer(r.Role)

	if err := validateRole(r.Role); err != nil {
		return err
	}

	resp, err := h.svc.Update(c.Request().Context(), h.auth.User(c), id, r)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

// @Security		BearerToken
// @Summary		Deletes an user
// @Description	Deletes an user
// @Accept			json
// @Produce		json
// @Tags			users
// @ID				usersDelete
// @Param			id				path		int	true	"User ID"
// @Success		200				{object}	swaggerutil.SwaggOKResp
// @Failure		400				{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		401				{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		403				{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		404				{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		500				{object}	swaggerutil.SwaggErrDetailsResp
// @Router			/v1/users/{id}	[delete]
func (h *HTTP) delete(c echo.Context) error {
	id, err := httputil.ReqID(c)
	if err != nil {
		return err
	}
	if err := h.svc.Delete(c.Request().Context(), h.auth.User(c), id); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

// @Security		BearerToken
// @Summary		Returns authenticated user
// @Description	Returns authenticated user
// @Accept			json
// @Produce		json
// @Tags			users
// @ID				usersMe
// @Success		200				{object}	model.User
// @Failure		400				{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		401				{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		403				{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		404				{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		500				{object}	swaggerutil.SwaggErrDetailsResp
// @Router			/v1/users/me	[get]
func (h *HTTP) me(c echo.Context) error {
	resp, err := h.svc.Me(c.Request().Context(), h.auth.User(c))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

// @Security		BearerToken
// @Summary		Changes authenticated user password
// @Description	Changes authenticated user password
// @Accept			json
// @Produce		json
// @Tags			users
// @ID				usersChangePwd
// @Param			request				body		user.PasswordChangeData	true	"PasswordChangeData"
// @Success		200					{object}	swaggerutil.SwaggOKResp
// @Failure		400					{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		401					{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		403					{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		404					{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		500					{object}	swaggerutil.SwaggErrDetailsResp
// @Router			/v1/users/password	[get]
func (h *HTTP) changePassword(c echo.Context) error {
	r := PasswordChangeData{}
	if err := c.Bind(&r); err != nil {
		return err
	}
	if err := h.svc.ChangePassword(c.Request().Context(), h.auth.User(c), r); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func validateRole(input *string) error {
	if input == nil {
		return nil
	}

	validRole := false
	for _, role := range model.AvailableRoles {
		if role == *input {
			validRole = true
			break
		}
	}
	if !validRole {
		return server.NewHTTPValidationError("Invalid role")
	}
	return nil
}
