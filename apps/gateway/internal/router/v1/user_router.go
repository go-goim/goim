package v1

import (
	"github.com/gin-gonic/gin"

	userv1 "github.com/go-goim/core/api/user/v1"
	"github.com/go-goim/core/apps/gateway/internal/service"
	"github.com/go-goim/core/pkg/mid"
	"github.com/go-goim/core/pkg/request"
	"github.com/go-goim/core/pkg/response"
	"github.com/go-goim/core/pkg/router"
)

type UserRouter struct {
	router.Router
}

func NewUserRouter() *UserRouter {
	return &UserRouter{
		Router: &router.BaseRouter{},
	}
}

func (r *UserRouter) Load(router *gin.RouterGroup) {
	friend := NewFriendRouter()
	friend.Load(router.Group("/friend", mid.AuthJwtCookie))

	router.POST("/query", mid.AuthJwtCookie, r.queryUser)
	router.POST("/login", r.login)
	router.POST("/register", r.register)
	router.POST("/update", mid.AuthJwtCookie, r.updateUserInfo)
}

type QueryRequestForSwagger struct {
	// Email and Phone only one can be set
	Email *string `json:"email" example:"user1@example.com"`
	Phone *string `json:"phone" example:"13800138000"`
}

// @Summary 查询用户信息
// @Description 查询用户信息
// @Tags [gateway]用户
// @Accept json
// @Produce json
// @Param   req body QueryRequestForSwagger true "req"
// @Success 200 {object} response.Response{data=userv1.User}
// @Failure 200 {object} response.Response
// @Router /gateway/v1/user/query [post]
func (r *UserRouter) queryUser(c *gin.Context) {
	var req = &userv1.QueryUserRequest{}
	if err := c.ShouldBindWith(req, &request.PbJSONBinding{}); err != nil {
		response.ErrorResp(c, err)
		return
	}

	user, err := service.GetUserService().QueryUserInfo(mid.GetContext(c), req)
	if err != nil {
		response.ErrorResp(c, err)
		return
	}

	response.SuccessResp(c, user)
}

type LoginRequestForSwagger struct {
	// Email and Phone only one can be set
	Email     *string `json:"email" example:"user1@example.com"`
	Phone     *string `json:"phone" example:"13800138000"`
	Password  string  `json:"password" example:"123456"`
	LoginType int     `json:"loginType" example:"0"`
}

// @Summary 登录
// @Description 用户登录
// @Tags [gateway]用户
// @Accept json
// @Produce json
// @Param   req body LoginRequestForSwagger true "req"
// @Success 200 {object} response.Response{data=userv1.User}
// @Header  200 {string} Authorization "Bearer "
// @Router /gateway/v1/user/login [post]
func (r *UserRouter) login(c *gin.Context) {
	var req = &userv1.UserLoginRequest{}
	if err := c.ShouldBindWith(req, &request.PbJSONBinding{}); err != nil {
		response.ErrorResp(c, err)
		return
	}

	user, err := service.GetUserService().Login(mid.GetContext(c), req)
	if err != nil {
		response.ErrorResp(c, err)
		return
	}

	if err = mid.SetJwtToHeader(c, user.Uid); err != nil {
		response.ErrorResp(c, err)
		return
	}

	response.SuccessResp(c, user)
}

// @Summary 注册
// @Description 用户注册
// @Tags [gateway]用户
// @Accept json
// @Produce json
// @Param   req body userv1.CreateUserRequest true "req"
// @Success 200 {object} userv1.User
// @Failure 200 {object} response.Response
// @Router /gateway/v1/user/register [post]
func (r *UserRouter) register(c *gin.Context) {
	var req = &userv1.CreateUserRequest{}
	if err := c.ShouldBindWith(req, &request.PbJSONBinding{}); err != nil {
		response.ErrorResp(c, err)
		return
	}

	user, err := service.GetUserService().Register(mid.GetContext(c), req)
	if err != nil {
		response.ErrorResp(c, err)
		return
	}

	response.SuccessResp(c, user)
}

// @Summary 更新用户信息
// @Description 更新用户信息
// @Tags [gateway]用户
// @Accept json
// @Produce json
// @Param   req body userv1.UpdateUserRequest true "req"
// @Success 200 {object} userv1.User
// @Failure 200 {object} response.Response
// @Router /gateway/v1/user/update [post]
func (r *UserRouter) updateUserInfo(c *gin.Context) {
	var req = &userv1.UpdateUserRequest{}
	if err := c.ShouldBindWith(req, &request.PbJSONBinding{}); err != nil {
		response.ErrorResp(c, err)
		return
	}

	user, err := service.GetUserService().UpdateUser(mid.GetContext(c), req)
	if err != nil {
		response.ErrorResp(c, err)
		return
	}

	response.SuccessResp(c, user)
}
