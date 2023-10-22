package controller

import (
	"base-gin-go/models"
	"base-gin-go/pkg/response"
	"base-gin-go/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserServicer
}

func NewUserController(s services.UserServicer) *UserController {
	return &UserController{service: s}
}

// Login godoc
//
//	@Summary		login
//	@Description	login
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			json	body		models.RequestLogin	true	"Body"
//	@Success		200		{object}	models.Response
//	@Failure		400		{object}	models.ErrorResponse
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/login [post]
func (c *UserController) Login(ctx *gin.Context) {
	req := models.RequestLogin{}
	if !bindJSON(ctx, &req) {
		return
	}
	token, err := c.service.Login(ctx, req)
	if checkError(ctx, err) {
		return
	}
	response.OK(ctx, gin.H{
		"token": "Bearer " + token,
	})
}

// GetUser godoc
//
//	@Summary		Get a user
//	@Description	Get a user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	models.User
//	@Failure		400	{object}	models.ErrorResponse
//	@Failure		404	{object}	models.ErrorResponse
//	@Failure		500	{object}	models.ErrorResponse
//	@Security		Bearer
//	@Router			/users/{id} [get]
func (c *UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if checkError(ctx, err) {
		return
	}
	user, err := c.service.GetUser(ctx, id)
	if checkError(ctx, err) {
		return
	}
	response.OK(ctx, user)
}
