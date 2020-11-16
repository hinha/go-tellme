package rest

import (
	"fmt"
	"github.com/coocood/freecache"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"go-tellme/internal/constants/model"
	"go-tellme/internal/constants/vo"
	"go-tellme/internal/module/client"
	"net/http"
)

type webHandler struct {
	usecase client.Usecase
	cache   *freecache.Cache
}

var (
	emptyField = "Email dan password tidak boleh kosong"
	wrongType  = "Email atau password tidak terdaftar"
)

func (u *webHandler) PageHome(ctx *gin.Context) {
	data := map[string]interface{}{
		"title":     "Home",
		"nav":       gin.H{"name": "Sign Out", "link": "/logout"},
		"is_logged": true,
	}

	_, err := ctx.Cookie("session")
	if err != nil {
		data["nav"] = gin.H{
			"name": "Sign In",
			"link": "/login",
		}
	}

	gintemplate.HTML(ctx, http.StatusOK, "index", data)
}

func (u *webHandler) PageAbout(ctx *gin.Context) {
	data := map[string]interface{}{
		"title": "About",
		"nav": gin.H{
			"name": "Sign Out",
			"link": "/logout",
		},
		"is_logged": true,
	}
	gintemplate.HTML(ctx, http.StatusOK, "about", data)
}

func (u *webHandler) PageLogin(ctx *gin.Context) {
	data := map[string]interface{}{
		"title":      "Sign In",
		"form_title": "Sign In",
		"nav": gin.H{
			"name": "Sign Up",
			"link": "/signup",
		},
	}

	gintemplate.HTML(ctx, http.StatusOK, "login", data)
}

func (u *webHandler) PageLoginPerform(ctx *gin.Context) {
	data := map[string]interface{}{
		"title":      "Sign In",
		"form_title": "Sign In",
		"nav": gin.H{
			"name": "Sign Up",
			"link": "/signup",
		},
	}

	var form vo.LoginForm

	if err := ctx.ShouldBind(&form); err != nil {
		//	error func
		data["error"] = gin.H{"message": emptyField}
		gintemplate.HTML(ctx, http.StatusBadRequest, "login", data)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	userID, err := u.usecase.Login(form.Email, form.Password)
	if err != nil {
		data["error"] = gin.H{"message": wrongType}
		gintemplate.HTML(ctx, http.StatusBadRequest, "login", data)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := u.usecase.GetToken(userID)
	if err != nil {
		data["error"] = gin.H{"message": wrongType}
		gintemplate.HTML(ctx, http.StatusBadRequest, "login", data)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.SetCookie("session", token, 3600, "/", "", false, true)
	ctx.Redirect(http.StatusFound, "/")
}

func (u *webHandler) PageSignup(ctx *gin.Context) {
	data := map[string]interface{}{
		"title":      "Sign Up",
		"form_title": "Sign Up",
		"nav": gin.H{
			"name": "Sign In",
			"link": "/login",
		},
	}

	gintemplate.HTML(ctx, http.StatusOK, "register", data)
}

func (u *webHandler) PageSignupPerform(ctx *gin.Context) {
	data := map[string]interface{}{
		"title":      "Sign Up",
		"form_title": "Sign Up",
		"nav": gin.H{
			"name": "Sign In",
			"link": "/login",
		},
	}

	var form vo.RegisterForm

	if err := ctx.ShouldBind(&form); err != nil {
		//	error func
		data["error"] = gin.H{"message": emptyField}
		gintemplate.HTML(ctx, http.StatusBadRequest, "register", data)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if form.Email == "" || form.Password == "" {
		data["error"] = gin.H{"message": emptyField}
		gintemplate.HTML(ctx, http.StatusBadRequest, "register", data)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if len(form.Email) < 5 || len(form.Email) > 40 {
		data["error"] = gin.H{"message": "Email minimal 5 sampai 40 maksimal karakter"}
		gintemplate.HTML(ctx, http.StatusBadRequest, "register", data)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userID, err := u.usecase.Register(&model.User{
		Email:    form.Email,
		Password: form.Password,
	})
	if err != nil {
		data["error"] = gin.H{
			"message": "Email sudah terdaftar",
		}
		gintemplate.HTML(ctx, http.StatusBadRequest, "register", data)
		return
	}

	token, err := u.usecase.GetToken(userID)
	fmt.Println(err)
	ctx.SetCookie("session", token, 3600, "/", "", false, true)

	data["success"] = gin.H{
		"message": "Success",
	}

	gintemplate.HTML(ctx, http.StatusOK, "register", data)
}

func (u *webHandler) PageLogout(ctx *gin.Context) {
	ctx.SetCookie("session", "", -1, "/", "", false, true)
	ctx.Redirect(http.StatusFound, "/")
}

func NewHandlerWeb(usecase client.Usecase, cache *freecache.Cache) client.Handler {
	return &webHandler{usecase: usecase, cache: cache}
}
