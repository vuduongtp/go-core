package auth

import (
	"context"
	"net/http"

	"github.com/vuduongtp/go-core/internal/model"

	"github.com/labstack/echo/v4"
)

// HTTP represents auth http service
type HTTP struct {
	svc Service
}

// Service represents auth service interface
type Service interface {
	Authenticate(context.Context, Credentials) (*model.AuthToken, error)
	RefreshToken(context.Context, RefreshTokenData) (*model.AuthToken, error)
}

// NewHTTP creates new auth http service
func NewHTTP(svc Service, e *echo.Echo) {
	h := HTTP{svc}

	e.POST("/login", h.login)
	e.POST("/refresh-token", h.refreshToken)
}

// Credentials represents login request data
type Credentials struct {
	Username string `json:"username" validate:"required" example:"superadmin"`
	Password string `json:"password" validate:"required" example:"superadmin123!@#"`
}

// RefreshTokenData represents refresh token request data
type RefreshTokenData struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// @Summary		Logs in user by username and password
// @Description	Logs in user by username and password
// @Accept			json
// @Produce		json
// @Tags			auth
// @ID				authLogin
// @Param			request	body		auth.Credentials	true	"Credentials"
// @Success		200		{object}	model.AuthToken
// @Failure		401		{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		500		{object}	swaggerutil.SwaggErrDetailsResp
// @Router			/login [post]
func (h *HTTP) login(c echo.Context) error {
	r := Credentials{}
	if err := c.Bind(&r); err != nil {
		return err
	}
	resp, err := h.svc.Authenticate(c.Request().Context(), r)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary		Refresh access token
// @Description	Refresh access token
// @Accept			json
// @Produce		json
// @Tags			auth
// @ID				authRefreshToken
// @Param			request	body		auth.RefreshTokenData	true	"RefreshTokenData"
// @Success		200		{object}	model.AuthToken
// @Failure		401		{object}	swaggerutil.SwaggErrDetailsResp
// @Failure		500		{object}	swaggerutil.SwaggErrDetailsResp
// @Router			/refresh-token [post]
func (h *HTTP) refreshToken(c echo.Context) error {
	r := RefreshTokenData{}
	if err := c.Bind(&r); err != nil {
		return err
	}
	resp, err := h.svc.RefreshToken(c.Request().Context(), r)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
