package handler

import (
	"context"
	"fmt"
	"net/http"
	"personal_page/model"
	"personal_page/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

// Login 用户登录接口
// @Summary 用于用户登录
// @Description 完成用户登陆操作分发access_token和refresh_token
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param {object} body model.LoginRequest true "登录参数"
// @Success 200 {object} model.LoginResponse
// @Router /user/login [post]
func (h *UserHandler) Login(ctx *gin.Context) {
	var loginRequest model.LoginRequest
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		ctx.JSON(http.StatusOK, model.LoginResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	result, err := h.usecase.Login(context.Background(), loginRequest.Username, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, model.LoginResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.LoginResponse{
		Status:       "ok",
		Message:      "",
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	})
}

// Auth 验证token
// @Summary 验证用户身份
// @Description 使用access_token验证用户身份
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "访问令牌"
// @Success 200 {object} string
// @Router /user/auth [post]
func (h *UserHandler) Auth(ctx *gin.Context) {
	accessToken := ctx.GetHeader("Authorization")
	if accessToken == "" {
		ctx.JSON(http.StatusOK, "no token")
		return
	}

	_, err := h.usecase.CheckAccessToken(context.Background(), accessToken)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "success")
}

// RefreshToken 更新访问令牌
// @Summary 更新访问令牌
// @Description 使用refresh_token更新访问令牌
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "refresh_token"
// @Success 200 {object} string
// @Router /user/refresh [post]
func (h *UserHandler) RefreshToken(ctx *gin.Context) {
	refreshToken := ctx.GetHeader("Authorization")
	if refreshToken == "" {
		ctx.JSON(http.StatusUnauthorized, "no token")
		return
	}

	accessToken, err := h.usecase.RefreshAccessToken(context.Background(), refreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, accessToken)
}

func NewUserHandler(uc usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: uc,
	}
}
