package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *webHandler) EnsureIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := u.usecase.VerifyToken(c); err != nil {
			c.SetCookie("session", "", -1, "/", "", false, true)
			return
		}
	}
}

func (u *webHandler) EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {

		// jika sudah login redirect home
		_, err := c.Cookie("session")
		if err == nil {
			c.Redirect(http.StatusFound, "/")
			c.AbortWithStatus(http.StatusFound)
			return
		}

	}

}

// if the user is already logged in
func (u *webHandler) EnsureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {

		// If there's no error or if the token is not empty
		_, err := c.Cookie("session")
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			c.AbortWithStatus(http.StatusFound)
			return
		}

		if err := u.usecase.VerifyToken(c); err != nil {
			c.SetCookie("session", "", -1, "/", "", false, true)
			c.Redirect(http.StatusFound, "/login")
			c.AbortWithStatus(http.StatusFound)
			return
		}
	}
}
