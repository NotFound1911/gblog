package web

import (
	"github.com/NotFound1911/gblog/internal/server"
	"github.com/NotFound1911/gblog/internal/user/domain"
	"github.com/NotFound1911/gblog/internal/user/errs"
	"github.com/NotFound1911/gblog/internal/user/service"
	ijwt "github.com/NotFound1911/gblog/internal/user/web/jwt"
	"github.com/NotFound1911/mserver"
	regexp "github.com/dlclark/regexp2"
)

const (
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	bizLogin             = "login"
)

type UserHandler struct {
	ijwt.Handler
	emailRexExp    *regexp.Regexp
	passwordRexExp *regexp.Regexp
	svc            service.UserService
}

func NewUserHandler(svc service.UserService,
	hdl ijwt.Handler) *UserHandler {
	return &UserHandler{
		emailRexExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		svc:            svc,
		Handler:        hdl,
	}
}
func (u *UserHandler) SignUp(ctx *mserver.Context, req SignUpReq) (server.Result, error) {
	isEmail, err := u.emailRexExp.MatchString(req.Email)
	if err != nil {
		return server.Result{
			Code: errs.UserInvalidInput,
			Msg:  "系统错误",
		}, err
	}
	if !isEmail {
		return server.Result{
			Code: errs.UserInvalidInput,
			Msg:  "非法邮箱格式",
		}, nil
	}
	if req.Password != req.ConfirmPassword {
		return server.Result{
			Code: errs.UserInvalidInput,
			Msg:  "两次输入的密码不相等",
		}, nil
	}
	isPassword, err := u.passwordRexExp.MatchString(req.Password)
	if err != nil {
		return server.Result{
			Code: errs.UserInternalServerError,
			Msg:  "系统错误",
		}, err
	}
	if !isPassword {
		return server.Result{
			Code: errs.UserInvalidInput,
			Msg:  "密码必须包含字母、数字、特殊字符,并且不少于八位",
		}, nil
	}
	err = u.svc.Signup(ctx.GetRequest().Context(), domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	switch err {
	case nil:
		return server.Result{
			Msg: "OK",
		}, nil
	case service.ErrDuplicateEmail:
		return server.Result{
			Code: errs.UserDuplicateEmail,
			Msg:  "邮箱冲突",
		}, nil
	default:
		return server.Result{
			Code: errs.UserInternalServerError,
			Msg:  "系统错误",
		}, err
	}
}
func (u *UserHandler) RegisterRoutes(core *mserver.Core) {

	ug := core.Group("/users")
	ug.Post("/signup", server.WrapBody(u.SignUp))
}
